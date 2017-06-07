-- CREATE DATABASE go_restful;
-- go_restful created while created container, because it was in environments in docker-compose.yml
GRANT ALL PRIVILEGES ON DATABASE go_restful TO postgres;

CREATE SEQUENCE "news_shard_id_seq" START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

create table news_shard_other(
    id bigint not null,
    category_id int not null constraint category_id_check check (category_id not in (1, 2)),
    author character varying not null,
    rate int not null,
    title character varying,
    text text
);

create index news_shard_other_rate_idx on news_shard_other(rate);

create extension postgres_fdw;

-- for first shard:
create server news_1_server foreign data wrapper postgres_fdw
options (host '172.18.0.5', port '5432', dbname 'go_restful');
create user mapping for postgres server news_1_server
options (user 'postgres', password 'postgres');

create foreign table news_1_shard(
    id bigint not null,
    category_id int not null,
    author character varying not null,
    rate int not null,
    title character varying,
    text text
)
server news_1_server
options (schema_name 'public', table_name 'news_shard_1');

-- for second shard:
create server news_2_server foreign data wrapper postgres_fdw
options (host '172.18.0.6', port '5432', dbname 'go_restful');
create user mapping for postgres server news_2_server
options (user 'postgres', password 'postgres');

create foreign table news_2_shard(
    id bigint not null,
    category_id int not null,
    author character varying not null,
    rate int not null,
    title character varying,
    text text
)
server news_2_server
options (schema_name 'public', table_name 'news_shard_2');

-- create view
create view news_shard as
    select * from news_1_shard
    union all
    select * from news_2_shard
    union all
    select * from news_shard_other;

ALTER VIEW news_shard ALTER COLUMN id SET DEFAULT NEXTVAL('news_shard_id_seq');

-- note: work, but need returning
-- create rule news_shard_insert as on insert to news_shard do instead nothing;
-- create rule news_shard_update as on update to news_shard do instead nothing;
-- create rule news_shard_delete as on delete to news_shard do instead nothing;

-- create rule news_shard_1_insert as on insert to news_shard where category_id=1
-- do instead insert into news_1_shard values (new.*);
-- create rule news_shard_2_insert as on insert to news_shard where category_id=2
-- do instead insert into news_2_shard values (new.*);
-- create rule news_shard_other_insert as on insert to news_shard where category_id not in (1, 2)
-- do instead insert into news_shard_other values (new.*);

-- https://www.postgresql.org/docs/current/static/plpgsql-trigger.html
-- this is the actual shard insert trigger:
CREATE FUNCTION news_shard_insert() RETURNS TRIGGER LANGUAGE plpgsql
AS $f$
  BEGIN
    IF NEW.category_id = 1 THEN
        INSERT INTO news_1_shard SELECT NEW.*;
    ELSIF NEW.category_id = 2 THEN
        INSERT INTO news_2_shard SELECT NEW.*;
    ELSE
        INSERT INTO news_shard_other SELECT NEW.*;
    END IF;
    RETURN NEW;
  END;
$f$;
CREATE TRIGGER news_shard_insert instead OF INSERT ON news_shard
  FOR each ROW EXECUTE PROCEDURE news_shard_insert();

-- this is the actual shard update trigger:
CREATE FUNCTION news_shard_update() RETURNS TRIGGER LANGUAGE plpgsql
AS $f$
  BEGIN
    IF OLD.category_id = 1 THEN
        IF NEW.category_id = 1 THEN
            UPDATE news_1_shard SET category_id = NEW.category_id, author = NEW.author,
                title = NEW.title, text = NEW.text, rate = NEW.rate WHERE id = OLD.id;
        ELSE
            DELETE FROM news_1_shard WHERE id = OLD.id;
            IF NEW.category_id = 2 THEN
                INSERT INTO news_2_shard (id, category_id, author, title, text, rate)
                    VALUES (OLD.id, NEW.category_id, NEW.author, NEW.title, NEW.text, NEW.rate);
            ELSE
                INSERT INTO news_shard_other (id, category_id, author, title, text, rate)
                    VALUES (OLD.id, NEW.category_id, NEW.author, NEW.title, NEW.text, NEW.rate);
            END IF;
        END IF;
    ELSIF OLD.category_id = 2 THEN
        IF NEW.category_id = 2 THEN
            UPDATE news_2_shard SET category_id = NEW.category_id, author = NEW.author,
                title = NEW.title, text = NEW.text, rate = NEW.rate WHERE id = OLD.id;
        ELSE
            DELETE FROM news_2_shard WHERE id = OLD.id;
            IF NEW.category_id = 1 THEN
                INSERT INTO news_1_shard (id, category_id, author, title, text, rate)
                    VALUES (OLD.id, NEW.category_id, NEW.author, NEW.title, NEW.text, NEW.rate);
            ELSE
                INSERT INTO news_shard_other (id, category_id, author, title, text, rate)
                    VALUES (OLD.id, NEW.category_id, NEW.author, NEW.title, NEW.text, NEW.rate);
            END IF;
        END IF;
    ELSE
        IF NEW.category_id NOT IN (1, 2) THEN
            UPDATE news_shard_other SET category_id = NEW.category_id, author = NEW.author,
                title = NEW.title, text = NEW.text, rate = NEW.rate WHERE id = OLD.id;
        ELSE
            DELETE FROM news_shard_other WHERE id = OLD.id;
            IF NEW.category_id = 1 THEN
                INSERT INTO news_1_shard (id, category_id, author, title, text, rate)
                    VALUES (OLD.id, NEW.category_id, NEW.author, NEW.title, NEW.text, NEW.rate);
            ELSIF NEW.category_id = 2 THEN
                INSERT INTO news_2_shard (id, category_id, author, title, text, rate)
                    VALUES (OLD.id, NEW.category_id, NEW.author, NEW.title, NEW.text, NEW.rate);
            END IF;
        END IF;
    END IF;
    RETURN NEW;
  END;
$f$;
CREATE TRIGGER news_shard_update instead OF UPDATE ON news_shard
  FOR each ROW EXECUTE PROCEDURE news_shard_update();

-- this is the actual shard delete trigger:
CREATE FUNCTION news_shard_delete() RETURNS TRIGGER LANGUAGE plpgsql
AS $f$
  BEGIN
    IF OLD.category_id = 1 THEN
        DELETE FROM news_1_shard WHERE id = OLD.id;
    ELSIF OLD.category_id = 2 THEN
        DELETE FROM news_2_shard WHERE id = OLD.id;
    ELSE
        DELETE FROM news_shard_other WHERE id = OLD.id;
    END IF;
    RETURN NULL;
  END;
$f$;
CREATE TRIGGER news_shard_delete instead OF DELETE ON news_shard
  FOR each ROW EXECUTE PROCEDURE news_shard_delete();

-- note: need sleep/wait before shards postgres servers will starts
-- https://www.if-not-true-then-false.com/2010/postgresql-sleep-function-pg_sleep-postgres-delay-execution/
SELECT pg_sleep(20);
insert into news_shard (category_id, title, author, rate, text)
values (1, 'news #1', 'author 1', 1, 'text 1'),
(2, 'news #2', 'author 2', 1, 'text 2'),
(3, 'news #3', 'author 3', 1, 'text 3'),
(4, 'news #4', 'author 4', 1, 'text 4');