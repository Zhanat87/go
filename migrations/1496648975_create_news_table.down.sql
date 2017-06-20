-- https://stackoverflow.com/questions/35338711/cannot-drop-table-users-because-other-objects-depend-on-it
DROP TABLE IF EXISTS "news_real" CASCADE;
DROP TABLE IF EXISTS "news_partition_1" CASCADE;
DROP TABLE IF EXISTS "news_partition_2" CASCADE;
DROP TABLE IF EXISTS "news_partition_other" CASCADE;
DROP SEQUENCE IF EXISTS news_id_seq;
DROP FUNCTION IF EXISTS news_nope();
DROP FUNCTION IF EXISTS news_partition_insert();
DROP FUNCTION IF EXISTS news_partition_update();
DROP FUNCTION IF EXISTS news_partition_delete();