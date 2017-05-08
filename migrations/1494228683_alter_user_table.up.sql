ALTER TABLE "public"."user" ADD CONSTRAINT constraint_unique_username UNIQUE (username);
ALTER TABLE "public"."user" ADD CONSTRAINT constraint_unique_email UNIQUE (email);