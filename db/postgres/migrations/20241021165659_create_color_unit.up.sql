CREATE TABLE color_unit_infos (
    id uuid PRIMARY KEY,
    world_id uuid NOT NULL,
    color VARCHAR (7),
    FOREIGN KEY (world_id) REFERENCES worlds(id)
);
