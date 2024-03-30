ALTER TABLE portal_unit_infos
ADD COLUMN target_unit_id uuid,
ADD CONSTRAINT portal_unit_info_fk_target_unit_id FOREIGN KEY(target_unit_id) REFERENCES units(id);

DO $$
DECLARE
  portalUnitInfoRow RECORD;
  targetPortalUnitRow RECORD;
BEGIN
  FOR portalUnitInfoRow IN SELECT * FROM portal_unit_infos LOOP
    if portalUnitInfoRow.target_pos_x IS NULL THEN
        UPDATE units
        SET info_snapshot = jsonb_set(info_snapshot, '{target_unit_id}', 'null')
        WHERE id = portalUnitInfoRow.id;
    ELSE
        SELECT * INTO targetPortalUnitRow FROM units WHERE pos_x = portalUnitInfoRow.target_pos_x AND pos_z = portalUnitInfoRow.target_pos_z;

        UPDATE portal_unit_infos
        SET target_unit_id = targetPortalUnitRow.id
        WHERE id = portalUnitInfoRow.id;

        UPDATE units
        SET info_snapshot = jsonb_set(info_snapshot, '{target_unit_id}', to_jsonb(targetPortalUnitRow.id))
        WHERE id = portalUnitInfoRow.id;
    END IF;
  END LOOP;
END $$;
