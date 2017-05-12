ALTER TABLE "public"."user" DROP COLUMN provider;
ALTER TABLE "public"."user" DROP COLUMN provider_id;
DROP TYPE auth_type;

ALTER TABLE "public"."user" ADD CONSTRAINT constraint_unique_username UNIQUE (username);
ALTER TABLE "public"."user" ADD CONSTRAINT constraint_unique_email UNIQUE (email);

ALTER TABLE "public"."user" DROP CONSTRAINT IF EXISTS constraint_unique_username_by_provider;
ALTER TABLE "public"."user" DROP CONSTRAINT IF EXISTS constraint_unique_email_by_provider;
