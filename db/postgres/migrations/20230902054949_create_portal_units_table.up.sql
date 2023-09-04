CREATE TABLE portal_units (
	id uuid PRIMARY KEY,
	target_pos_x INT,
	target_pos_z INT
);

ALTER TABLE units ADD COLUMN linked_unit_id uuid;
