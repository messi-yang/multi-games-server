-- Add missing world foreign key to link_unit_infos
ALTER TABLE link_unit_infos ADD FOREIGN KEY (world_id) REFERENCES worlds(id);
ALTER TABLE link_unit_infos ADD PRIMARY KEY (id);
ALTER TABLE link_unit_infos ALTER COLUMN url SET NOT NULL;

CREATE TABLE embed_unit_infos (
    id uuid PRIMARY KEY,
    world_id uuid NOT NULL,
    embed_code VARCHAR (2048),
    FOREIGN KEY (world_id) REFERENCES worlds(id)
);
