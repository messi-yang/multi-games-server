CREATE TABLE IF NOT EXISTS players (
    id uuid NOT NULL,
    gamer_id uuid,
	world_id uuid NOT NULL,
    name VARCHAR (50) NOT NULL,
	pos_x INT NOT NULL,
	pos_z INT NOT NULL,
    direction INT NOT NULL,
    held_item_id uuid NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_gamer FOREIGN KEY(gamer_id) REFERENCES gamers(id),
    CONSTRAINT fk_world FOREIGN KEY(world_id) REFERENCES worlds(id),
    CONSTRAINT fk_item FOREIGN KEY(held_item_id) REFERENCES items(id)
);
CREATE INDEX players_world_id_pos_x_pos_z ON players (world_id, pos_x, pos_z);
