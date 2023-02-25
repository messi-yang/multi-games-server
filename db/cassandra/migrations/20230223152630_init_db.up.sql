CREATE TABLE IF NOT EXISTS units (
  game_id uuid,
  pos_x int,
  pos_z int,
  item_id int,
  PRIMARY KEY((game_id, pos_x), pos_z)
) WITH CLUSTERING ORDER BY (pos_z ASC);
