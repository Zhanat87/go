ALTER TABLE "public"."user" RENAME password_hash TO password;
ALTER TABLE "public"."user" DROP COLUMN password_reset_token;