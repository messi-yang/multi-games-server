CREATE TABLE occupied_positions (
	world_id uuid NOT NULL,
	pos_x INT NOT NULL,
	pos_z INT NOT NULL,
    unit_id uuid NOT NULL,
    FOREIGN KEY (world_id) REFERENCES worlds(id),
    FOREIGN KEY (unit_id) REFERENCES units(id),
    CONSTRAINT occupied_positions_unique_world_id_pos_x_pos_z UNIQUE (world_id, pos_x, pos_z)
);

DO $$
DECLARE
  unitRow RECORD;
BEGIN
  FOR unitRow IN SELECT * FROM units LOOP
    INSERT INTO occupied_positions (world_id, pos_x, pos_z, unit_id)
    VALUES (unitRow.world_id, unitRow.pos_x, unitRow.pos_z, unitRow.id);
  END LOOP;
END $$;

ALTER TABLE items
ADD COLUMN dimension_width INT NOT NULL DEFAULT 1,
ADD COLUMN dimension_depth INT NOT NULL DEFAULT 1;
