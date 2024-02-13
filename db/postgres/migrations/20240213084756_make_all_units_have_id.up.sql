DO $$
DECLARE
  unitRow RECORD;
BEGIN
  FOR unitRow IN SELECT * FROM units WHERE info_id IS NULL LOOP
    UPDATE units
        SET info_id = gen_random_uuid()
        WHERE world_id = unitRow.world_id AND pos_x = unitRow.pos_x AND pos_z = unitRow.pos_z;
  END LOOP;
END $$;


ALTER TABLE units ADD CONSTRAINT units_unique_info_id UNIQUE (info_id);

ALTER TABLE units ALTER COLUMN info_id SET NOT NULL;
