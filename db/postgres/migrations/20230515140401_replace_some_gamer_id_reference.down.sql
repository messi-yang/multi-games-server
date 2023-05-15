-- Drop foreign key to gamers table
ALTER TABLE worlds DROP CONSTRAINT fk_user;

-- Replace user_id to its corresponding gamer_id based on gamers table
DO $$
DECLARE
  worldRow RECORD;
  gamerRow RECORD;
BEGIN
  FOR worldRow IN SELECT * FROM worlds LOOP
    SELECT * INTO gamerRow FROM gamers WHERE user_id = worldRow.user_id;
    UPDATE worlds SET user_id = gamerRow.id WHERE user_id = worldRow.user_id;
  END LOOP;
END $$;

-- Rename user_id to game_id
ALTER TABLE worlds RENAME user_id to gamer_id;

-- Add foreign key to users table
ALTER TABLE worlds ADD CONSTRAINT fk_game_user FOREIGN KEY(gamer_id) REFERENCES gamers(id); 
