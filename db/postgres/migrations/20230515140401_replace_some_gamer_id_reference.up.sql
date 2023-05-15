-- Drop foreign key to gamers table
ALTER TABLE worlds DROP CONSTRAINT fk_game_user;

-- Replace gamer_id to its corresponding user_id based on gamers table
DO $$
DECLARE
  worldRow RECORD;
  gamerRow RECORD;
BEGIN
  FOR worldRow IN SELECT * FROM worlds LOOP
    SELECT * INTO gamerRow FROM gamers WHERE id = worldRow.gamer_id;
    UPDATE worlds SET gamer_id = gamerRow.user_id WHERE gamer_id = worldRow.gamer_id;
  END LOOP;
END $$;

-- Rename game_id to user_id
ALTER TABLE worlds RENAME gamer_id to user_id;

-- Add foreign key to users table
ALTER TABLE worlds ADD CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id);
