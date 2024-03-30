ALTER TABLE portal_unit_infos
ADD COLUMN target_pos_x INT,
ADD COLUMN target_pos_z INT;

DO $$
DECLARE
  portalUnitInfoRow RECORD;
  targetPortalUnitRow RECORD;
BEGIN
  FOR portalUnitInfoRow IN SELECT * FROM portal_unit_infos LOOP
    if portalUnitInfoRow.target_unit_id IS NULL THEN
        UPDATE units
        SET info_snapshot = jsonb_set(info_snapshot, '{targetPosition}', 'null')
        WHERE id = portalUnitInfoRow.id;
    ELSE
        SELECT * INTO targetPortalUnitRow FROM units WHERE id = portalUnitInfoRow.target_unit_id;

        UPDATE portal_unit_infos
        SET target_pos_x = targetPortalUnitRow.pos_x, target_pos_z = targetPortalUnitRow.pos_z
        WHERE id = portalUnitInfoRow.id;

        UPDATE units
        SET info_snapshot = jsonb_set(
            info_snapshot,
            '{targetPosition}',
            jsonb_build_object(
                'x', targetPortalUnitRow.pos_x,
                'z', targetPortalUnitRow.pos_z
            )
        )
        WHERE id = portalUnitInfoRow.id;
    END IF;
  END LOOP;
END $$;
