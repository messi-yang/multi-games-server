ALTER TABLE units ADD COLUMN info_snapshot JSONB;

DO $$
DECLARE
  portalUnitInfoRow RECORD;
BEGIN
  FOR portalUnitInfoRow IN SELECT * FROM portal_unit_infos LOOP
    IF portalUnitInfoRow.target_pos_x IS NOT NULL THEN
        UPDATE units
            SET info_snapshot = jsonb_build_object(
                'targetPosition',
                jsonb_build_object(
                    'x', portalUnitInfoRow.target_pos_x,
                    'z', portalUnitInfoRow.target_pos_z
                )
            )
            WHERE info_id = portalUnitInfoRow.id;
    ELSE
        UPDATE units
            SET info_snapshot = jsonb_build_object(
                'targetPosition',
                NULL
            )
            WHERE info_id = portalUnitInfoRow.id;
    END IF;
  END LOOP;
END $$;
