CREATE SEQUENCE artist_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."artist" (
    "id" integer DEFAULT nextval('artist_id_seq') NOT NULL,
    "name" character varying(120) NOT NULL
) WITH (oids = false);