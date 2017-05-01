CREATE SEQUENCE artist_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 277 CACHE 1;

CREATE TABLE "public"."artist" (
    "id" integer DEFAULT nextval('artist_id_seq') NOT NULL UNIQUE,
    "name" character varying(120) NOT NULL,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp
) WITH (oids = false);