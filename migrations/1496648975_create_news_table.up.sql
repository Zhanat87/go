CREATE SEQUENCE "news_id_seq" START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

create table news_real(
    id bigint not null default nextval('news_id_seq'::regclass),
    category_id int not null,
    author character varying not null,
    rate int not null,
    title character varying,
    text text
);
ALTER TABLE ONLY "news_real" ADD CONSTRAINT "pk_news_real" PRIMARY KEY ("id");
ALTER TABLE ONLY "news_real" ADD CONSTRAINT "news_category_id_fkey" FOREIGN KEY (category_id)
REFERENCES "public"."category"(id) ON DELETE CASCADE NOT DEFERRABLE;

create table news_1 (
    check (category_id = 1)
) inherits (news_real);
create table news_2 (
    check (category_id = 2)
) inherits (news_real);
create table news_other (
    check (category_id NOT IN (1, 2))
) inherits (news_real);

-- block attempts to insert on the actual parent table
CREATE FUNCTION news_nope() RETURNS TRIGGER LANGUAGE plpgsql
AS $f$
  BEGIN
    raise exception 'insert on wrong table, can not save to news_real table directly';
    RETURN NULL;
  END;
$f$;
CREATE TRIGGER news_nope BEFORE INSERT ON news_real EXECUTE PROCEDURE news_nope();
 
-- create the view, which will be used for all access to the table:
CREATE VIEW news AS SELECT * FROM news_real;
 
-- need to copy any defaults from news to the view
ALTER VIEW news ALTER COLUMN id SET DEFAULT NEXTVAL('news_id_seq');

-- this is the actual partition insert trigger:
CREATE FUNCTION news_partition() RETURNS TRIGGER LANGUAGE plpgsql
AS $f$
  BEGIN
    IF NEW.category_id = 1 THEN
        INSERT INTO news_1 SELECT NEW.*;
    ELSIF NEW.category_id = 2 THEN
        INSERT INTO news_2 SELECT NEW.*;
    ELSE
        INSERT INTO news_other SELECT NEW.*;
    END IF;
    RETURN NEW;
  END;
$f$;

CREATE TRIGGER news_partition instead OF INSERT ON news
  FOR each ROW EXECUTE PROCEDURE news_partition();

-- create rule news_insert_to_1 as on insert to news where (category_id = 1) returning id
-- do instead insert into news_1 values (new.*);
-- create rule news_insert_to_2 as on insert to news where (category_id = 2) returning id
-- do instead insert into news_2 values (new.*);

-- note: need when absent chief table
-- https://gist.github.com/RhodiumToad/b82aac9aa4e3fbdda967d89b1e418aa4 - not work
-- https://wiki.postgresql.org/wiki/INSERT_RETURNING_vs_Partitioning
-- https://gist.github.com/copiousfreetime/59067 - need replace with rule
-- CREATE OR REPLACE FUNCTION news_insert_trigger()
-- RETURNS TRIGGER AS $$
-- DECLARE
--     r news%rowtype;
-- BEGIN
--     IF NEW.category_id = 1 THEN
--         INSERT INTO news_1 VALUES (NEW.*) RETURNING * INTO r;
--     ELSIF NEW.category_id = 2 THEN
--         INSERT INTO news_2 VALUES (NEW.*) RETURNING * INTO r;
--     ELSE
-- --         INSERT INTO news VALUES (NEW.*); recursion was
-- --         RAISE EXCEPTION 'Category id out of range. Fix the news_insert_trigger() function!', NEW.category_id;
--         INSERT INTO news_other VALUES (NEW.*) RETURNING * INTO r;
--     END IF;
-- --     RETURN NULL;
-- --     RETURN NEW.id;
-- --     RETURN NEW;
--     RETURN r; -- note: temporary decision, need without duplicate
-- END;
-- $$
-- LANGUAGE plpgsql;
--
-- -- Trigger to invoke the insert trigger
-- CREATE TRIGGER insert_news_trigger
--     BEFORE INSERT ON news
--     FOR EACH ROW EXECUTE PROCEDURE news_insert_trigger();
--
-- -- Trigger function to delete from the master table after the insert
-- CREATE OR REPLACE FUNCTION news_delete_master() RETURNS trigger
--     AS $$
-- DECLARE
--     r news%rowtype;
-- BEGIN
--     DELETE FROM ONLY news where id = new.id returning * into r;
--     RETURN r;
-- end;
-- $$
-- LANGUAGE plpgsql;
--
-- -- Create the after insert trigger
-- create trigger after_insert_news_trigger
--     after insert on news
--     for each row
-- execute procedure news_delete_master();

create index news_real_rate_idx on news_real(rate);
-- create index news_rate_idx on news(rate);
create index news_1_rate_idx on news_1(rate);
create index news_2_rate_idx on news_2(rate);
create index news_other_rate_idx on news_other(rate);

insert into news (category_id, title, author, rate, text)
values (1, 'news #1', 'john', 1, 'text 1');
insert into news_2 (category_id, title, author, rate, text)
values (2, 'news #2', 'doe', 1, 'text 2');
insert into news (category_id, title, author, rate, text)
values (3, 'news #3', 'author 3', 1, 'text 3');
insert into news_1 (category_id, title, author, rate, text)
values (1, 'news #4', 'author 4', 1, 'text 4');
insert into news (category_id, title, author, rate, text)
values (1, 'news #5', 'author 5', 1, 'text 5'),
(2, 'news #6', 'author 6', 1, 'text 6'),
(3, 'news #7', 'author 7', 1, 'text 7');