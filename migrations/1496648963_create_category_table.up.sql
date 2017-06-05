CREATE SEQUENCE category_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."category" (
    "id" integer DEFAULT nextval('category_id_seq') UNIQUE,
    "title" character varying(100) UNIQUE
) WITH (oids = false);

INSERT INTO "category" ("title") VALUES ('sport'), ('culture'), ('health');