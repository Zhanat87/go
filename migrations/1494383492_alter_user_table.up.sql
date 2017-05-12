CREATE TYPE auth_type AS ENUM ('app', 'facebook', 'twitter', 'google', 'github');
ALTER TABLE "public"."user" ADD COLUMN provider auth_type NOT NULL DEFAULT 'app';
ALTER TABLE "public"."user" ADD COLUMN provider_id character varying(100);

ALTER TABLE "public"."user" DROP CONSTRAINT IF EXISTS constraint_unique_username;
ALTER TABLE "public"."user" DROP CONSTRAINT IF EXISTS constraint_unique_email;

ALTER TABLE "public"."user" ADD CONSTRAINT constraint_unique_username_by_provider UNIQUE (username, provider);
ALTER TABLE "public"."user" ADD CONSTRAINT constraint_unique_email_by_provider UNIQUE (email, provider);