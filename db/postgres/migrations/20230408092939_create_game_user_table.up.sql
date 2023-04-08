-- Create game_users table
CREATE TABLE game_users (
	id uuid PRIMARY KEY,
	user_id uuid UNIQUE NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);

-- Insert records into game_users table according to users table
DO $$
DECLARE
  userRow RECORD;
BEGIN
  FOR userRow IN SELECT * FROM users LOOP
  	INSERT INTO game_users (id, user_id, created_at, updated_at) VALUES (gen_random_uuid(), userRow.id, NOW(), NOW());
  END LOOP;
END $$;

-- Drop foreign key to users table
ALTER TABLE worlds DROP CONSTRAINT fk_user;

-- Replace user_id to its corresponding game_user_id based on game_users table
DO $$
DECLARE
  worldRow RECORD;
  gameUserRow RECORD;
BEGIN
  FOR worldRow IN SELECT * FROM worlds LOOP
    SELECT * INTO gameUserRow FROM game_users WHERE user_id = worldRow.user_id;
    -- RAISE NOTICE 'The value of my_variable is %', gameUserRow.id;
    UPDATE worlds SET user_id = gameUserRow.id WHERE user_id = worldRow.user_id;
  END LOOP;
END $$;

-- Rename user_id to game_user_id
ALTER TABLE worlds RENAME user_id to game_user_id;

-- Add foreign key to game_users table
ALTER TABLE worlds ADD CONSTRAINT fk_game_user FOREIGN KEY(game_user_id) REFERENCES game_users(id); 
