CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."user" (
    "id" integer DEFAULT nextval('user_id_seq') UNIQUE,
    "username" character varying(100),
    "email" character varying(100),
    "password" character varying(72),
    "avatar" character varying(100) NOT NULL,
    "full_name" character varying(100),
    "phones" character varying(100) NOT NULL,
    "status" smallint,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp
) WITH (oids = false);