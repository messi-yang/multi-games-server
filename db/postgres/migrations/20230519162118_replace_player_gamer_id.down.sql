ALTER TABLE players DROP CONSTRAINT fk_players_user;

DO $$
DECLARE
  playerRow RECORD;
  gamerRow RECORD;
BEGIN
  FOR playerRow IN SELECT * FROM players LOOP
    SELECT * INTO gamerRow FROM gamers WHERE user_id = playerRow.user_id;
    UPDATE players SET user_id = gamerRow.id WHERE user_id = playerRow.user_id;
  END LOOP;
END $$;

ALTER TABLE players RENAME user_id to gamer_id;

ALTER TABLE players ADD CONSTRAINT fk_gamer FOREIGN KEY(gamer_id) REFERENCES gamers(id); 
