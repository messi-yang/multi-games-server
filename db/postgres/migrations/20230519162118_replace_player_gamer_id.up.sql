ALTER TABLE players DROP CONSTRAINT fk_gamer;

DO $$
DECLARE
  playerRow RECORD;
  gamerRow RECORD;
BEGIN
  FOR playerRow IN SELECT * FROM players LOOP
    SELECT * INTO gamerRow FROM gamers WHERE id = playerRow.gamer_id;
    UPDATE players SET gamer_id = gamerRow.user_id WHERE gamer_id = playerRow.gamer_id;
  END LOOP;
END $$;

ALTER TABLE players RENAME gamer_id to user_id;

ALTER TABLE players ADD CONSTRAINT fk_players_user FOREIGN KEY(user_id) REFERENCES users(id);
