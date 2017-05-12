INSERT INTO "public"."user" ("id", "username", "email", "password_hash", "full_name", "status", "created_at", "updated_at") VALUES
(2,	'test2',	'test2@test.com',
'$2a$10$YOGE3lBg7SXbhEa8kr8B3OBFimlWLrytjad8VquOFWBYIVY1UP.xa', -- 'pass' hash
  'test full name 2', 1, '2017-04-22 13:17:30', '2017-04-22 13:17:30');

ALTER SEQUENCE user_id_seq RESTART WITH 3;