ALTER TABLE "public"."user" RENAME password TO password_hash;
ALTER TABLE "public"."user" ADD COLUMN password_reset_token character varying(100) UNIQUE;