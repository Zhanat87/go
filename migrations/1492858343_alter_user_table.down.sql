ALTER TABLE "public"."user" ADD COLUMN deleted_at timestamp;
ALTER TABLE "public"."user" ALTER COLUMN avatar SET NOT NULL;
ALTER TABLE "public"."user" ALTER COLUMN phones SET NOT NULL;