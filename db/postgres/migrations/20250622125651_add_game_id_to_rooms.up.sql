ALTER TABLE rooms ADD COLUMN current_game_id UUID;
ALTER TABLE rooms ADD CONSTRAINT fk_rooms_games FOREIGN KEY (current_game_id) REFERENCES games(id);

ALTER TABLE games DROP COLUMN selected;
