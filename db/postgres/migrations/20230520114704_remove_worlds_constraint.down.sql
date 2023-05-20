ALTER TABLE gamers DROP COLUMN worlds_count_limit;
ALTER TABLE gamers DROP COLUMN worlds_count;
ALTER TABLE worlds ADD CONSTRAINT worlds_user_id_key UNIQUE (user_id);

