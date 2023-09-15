ALTER TABLE portal_units RENAME TO portal_unit_infos;

CREATE INDEX portal_unit_infos_world_id_target_pos_x_target_pos_z
    ON portal_unit_infos (world_id, target_pos_x, target_pos_z);

ALTER TABLE portal_unit_infos ADD COLUMN id uuid PRIMARY KEY DEFAULT gen_random_uuid();

ALTER TABLE units ADD COLUMN info_id uuid;

DO $$
DECLARE
  portalUnitInfoRow RECORD;
BEGIN
  FOR portalUnitInfoRow IN SELECT * FROM portal_unit_infos LOOP
    UPDATE portal_unit_infos SET id = gen_random_uuid() WHERE
        world_id = portalUnitInfoRow.world_id
        AND pos_x = portalUnitInfoRow.pos_x
        AND pos_z = portalUnitInfoRow.pos_z;
  END LOOP;
END $$;

DO $$
DECLARE
  portalUnitInfoRow RECORD;
BEGIN
  FOR portalUnitInfoRow IN SELECT * FROM portal_unit_infos LOOP
    UPDATE units SET info_id = portalUnitInfoRow.id WHERE
        world_id = portalUnitInfoRow.world_id
        AND pos_x = portalUnitInfoRow.pos_x
        AND pos_z = portalUnitInfoRow.pos_z;
  END LOOP;
END $$;

ALTER TABLE portal_unit_infos DROP COLUMN pos_x, DROP COLUMN pos_z, DROP COLUMN item_id, DROP COLUMN direction;
