CREATE TABLE color_unit_infos (
    id uuid PRIMARY KEY,
    world_id uuid NOT NULL,
    color VARCHAR (7),
    FOREIGN KEY (world_id) REFERENCES worlds(id)
);

INSERT INTO color_unit_infos (id, world_id, color)
SELECT id, world_id, color
FROM units
WHERE type = 'color';

UPDATE units
SET info_snapshot = jsonb_build_object('color', color),
    color = NULL
WHERE type = 'color';

ALTER TABLE units
DROP COLUMN color;
