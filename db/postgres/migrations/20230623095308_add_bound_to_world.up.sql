ALTER TABLE worlds
    ADD COLUMN bound_from_x INT DEFAULT -50,
    ADD COLUMN bound_from_z INT DEFAULT -50,
    ADD COLUMN bound_to_x INT DEFAULT 50,
    ADD COLUMN bound_to_z INT DEFAULT 50;
