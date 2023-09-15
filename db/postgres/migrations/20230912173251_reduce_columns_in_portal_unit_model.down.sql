ALTER TABLE portal_unit_infos ADD COLUMN pos_x integer NOT NULL DEFAULT 0;
ALTER TABLE portal_unit_infos ADD COLUMN pos_z integer NOT NULL DEFAULT 0;
ALTER TABLE portal_unit_infos ADD COLUMN item_id uuid NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000';
ALTER TABLE portal_unit_infos ADD COLUMN direction integer NOT NULL DEFAULT 0;

DO $$
DECLARE
  unitRow RECORD;
BEGIN
  FOR unitRow IN SELECT * FROM units LOOP
    UPDATE portal_unit_infos SET
        pos_x = unitRow.pos_x,
        pos_z = unitRow.pos_z,
        item_id = unitRow.item_id,
        direction = unitRow.direction
    WHERE
        id = unitRow.info_id;
  END LOOP;
END $$;

ALTER TABLE portal_unit_infos ADD CONSTRAINT fk_item FOREIGN KEY(item_id) REFERENCES items(id);
ALTER TABLE portal_unit_infos ADD CONSTRAINT portal_unit_world_id_pos_x_pos_z UNIQUE (world_id, pos_x, pos_z);


ALTER TABLE units DROP COLUMN info_id;

ALTER TABLE portal_unit_infos DROP COLUMN id;

DROP INDEX portal_unit_infos_world_id_target_pos_x_target_pos_z;

ALTER TABLE portal_unit_infos RENAME TO portal_units;
