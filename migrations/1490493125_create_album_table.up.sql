CREATE SEQUENCE album_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."album" (
    "id" integer DEFAULT nextval('album_id_seq'),
    "title" character varying(160),
    "artist_id" integer,
    CONSTRAINT "album_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "album_artist_id_fkey" FOREIGN KEY (artist_id) REFERENCES artist(id) ON DELETE CASCADE NOT DEFERRABLE
) WITH (oids = false);