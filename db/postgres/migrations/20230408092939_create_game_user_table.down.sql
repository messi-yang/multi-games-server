-- Drop foreign key to game_users table
ALTER TABLE worlds DROP CONSTRAINT fk_game_user;

-- Replace game_user_id to its corresponding user_id based on game_users table
DO $$
DECLARE
  worldRow RECORD;
  gameUserRow RECORD;
BEGIN
  FOR worldRow IN SELECT * FROM worlds LOOP
    SELECT * INTO gameUserRow FROM game_users WHERE id = worldRow.game_user_id;
    UPDATE worlds SET game_user_id = gameUserRow.user_id WHERE game_user_id = worldRow.game_user_id;
  END LOOP;
END $$;

-- Rename game_user_id to user_id
ALTER TABLE worlds RENAME game_user_id to user_id;

-- Add foreign key to users table
ALTER TABLE worlds ADD CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id);

-- Drop game_users table
DROP TABLE game_users;
