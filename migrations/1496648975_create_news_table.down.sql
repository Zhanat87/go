-- https://stackoverflow.com/questions/35338711/cannot-drop-table-users-because-other-objects-depend-on-it
DROP TABLE IF EXISTS "news" CASCADE;
DROP TABLE IF EXISTS "news_partition_1" CASCADE;
DROP TABLE IF EXISTS "news_partition_2" CASCADE;
DROP TABLE IF EXISTS "news_partition_other" CASCADE;
DROP SEQUENCE IF EXISTS news_id_seq;