CREATE SEQUENCE "news_id_seq" START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

create table news(
    id bigint not null default nextval('news_id_seq'::regclass),
    category_id int not null,
    author character varying not null,
    rate int not null,
    title character varying,
    text text
);
ALTER TABLE ONLY "news" ADD CONSTRAINT "pk_news" PRIMARY KEY ("id");
ALTER TABLE ONLY "news" ADD CONSTRAINT "news_category_id_fkey" FOREIGN KEY (category_id)
REFERENCES "public"."category"(id) ON DELETE CASCADE NOT DEFERRABLE;

create table news_1 (
    check (category_id = 1)
) inherits (news);
create table news_2 (
    check (category_id = 2)
) inherits (news);

create rule news_insert_to_1 as on insert to news where (category_id = 1)
do instead insert into news_1 values (new.*);
create rule news_insert_to_2 as on insert to news where (category_id = 2)
do instead insert into news_2 values (new.*);

create index news_rate_idx on news(rate);
create index news_1_rate_idx on news_1(rate);
create index news_2_rate_idx on news_2(rate);

insert into news (category_id, title, author, rate)
values (1, 'news #1', 'john', 1);
insert into news_2 (category_id, title, author, rate)
values (2, 'news #2', 'doe', 1);
insert into news (category_id, title, author, rate)
values (3, 'news #3', 'author 3', 1);
insert into news_1 (category_id, title, author, rate)
values (1, 'news #4', 'author 4', 1);
insert into news (category_id, title, author, rate)
values (1, 'news #5', 'author 5', 1), (2, 'news #6', 'author 6', 1), (3, 'news #7', 'author 7', 1);