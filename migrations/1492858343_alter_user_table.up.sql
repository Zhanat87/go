ALTER TABLE "public"."user" DROP COLUMN deleted_at;
ALTER TABLE "public"."user" ALTER COLUMN avatar DROP NOT NULL;
ALTER TABLE "public"."user" ALTER COLUMN phones DROP NOT NULL;