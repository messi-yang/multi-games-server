ALTER TABLE units ADD COLUMN info_snapshot jsonb;

DO $$
DECLARE
  portalUnitInfoRow RECORD;
BEGIN
  FOR portalUnitInfoRow IN SELECT * FROM portal_unit_infos LOOP
    if portalUnitInfoRow.target_unit_id IS NULL THEN
        UPDATE units
        SET info_snapshot = jsonb_set(info_snapshot, '{target_unit_id}', 'null')
        WHERE id = portalUnitInfoRow.id;
    ELSE
        UPDATE units
        SET info_snapshot = COALESCE(
            jsonb_set(
                info_snapshot,
                '{target_unit_id}',
                to_jsonb(portalUnitInfoRow.target_unit_id)
            ),
            jsonb_build_object('target_unit_id', portalUnitInfoRow.target_unit_id)
        )
        WHERE id = portalUnitInfoRow.id;
    END IF;
  END LOOP;
END $$;

