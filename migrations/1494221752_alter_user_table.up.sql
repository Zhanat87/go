UPDATE "public"."user" SET phones = NULL WHERE phones='';
ALTER TABLE "public"."user" ALTER COLUMN phones TYPE json USING phones::JSON;