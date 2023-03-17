CREATE TABLE IF NOT EXISTS units (
	world_id uuid NOT NULL,
	pos_x INT NOT NULL,
	pos_z INT NOT NULL,
    item_id uuid NOT NULL,
    direction INT NOT NULL,
    CONSTRAINT fk_world FOREIGN KEY(world_id) REFERENCES worlds(id),
    CONSTRAINT fk_item FOREIGN KEY(item_id) REFERENCES items(id),
    CONSTRAINT unique_world_id_pos_x_pos_z UNIQUE (world_id, pos_x, pos_z)
);
CREATE INDEX unit_world_id_pos_x_pos_z ON units (world_id, pos_x, pos_z);
