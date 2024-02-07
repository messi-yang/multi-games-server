DO $$
DECLARE
  linkUnitInfoRow RECORD;
BEGIN
  FOR linkUnitInfoRow IN SELECT * FROM link_unit_infos LOOP
    UPDATE units
        SET info_snapshot = jsonb_build_object('url', linkUnitInfoRow.url)
        WHERE info_id = linkUnitInfoRow.id;
  END LOOP;
END $$;
