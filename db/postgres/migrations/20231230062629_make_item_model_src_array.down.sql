ALTER TABLE items ADD COLUMN model_src VARCHAR(255) DEFAULT '' NOT NULL;

DO $$
DECLARE
  itemRow RECORD;
BEGIN
  FOR itemRow IN SELECT * FROM items LOOP
    UPDATE items
        SET model_src = model_sources[0]
        WHERE id = itemRow.id;
  END LOOP;
END $$;

ALTER TABLE items DROP COLUMN model_sources;
