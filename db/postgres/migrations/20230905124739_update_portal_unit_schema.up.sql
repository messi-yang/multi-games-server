ALTER TABLE units DROP COLUMN linked_unit_id;

DROP TABLE portal_units;

CREATE TABLE portal_units (
    world_id uuid NOT NULL,
    pos_x integer NOT NULL,
    pos_z integer NOT NULL,
    item_id uuid NOT NULL,
    direction integer NOT NULL,
	target_pos_x INT,
	target_pos_z INT,
    CONSTRAINT fk_world FOREIGN KEY(world_id) REFERENCES worlds(id),
    CONSTRAINT fk_item FOREIGN KEY(item_id) REFERENCES items(id),
    CONSTRAINT portal_unit_world_id_pos_x_pos_z UNIQUE (world_id, pos_x, pos_z)
);
