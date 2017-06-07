GRANT ALL PRIVILEGES ON DATABASE go_restful TO postgres;

CREATE SEQUENCE "news_replication_id_seq" START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

create table news_replication(
    id bigint not null default nextval('news_replication_id_seq'::regclass),
    category_id int not null,
    author character varying not null,
    rate int not null,
    title character varying,
    text text
);

ALTER TABLE ONLY "news_replication" ADD CONSTRAINT "pk_news_replication" PRIMARY KEY ("id");

create index news_replication_rate_idx on news_replication(rate);

insert into news_replication (category_id, title, author, rate, text)
values (1, 'slave news #1', 'slave author 1', 1, 'slave text 1'),
(2, 'slave news #2', 'slave author 2', 1, 'slave text 2'),
(3, 'slave news #3', 'slave author 3', 1, 'slave text 3'),
(4, 'slave news #4', 'slave author 4', 1, 'slave text 4');