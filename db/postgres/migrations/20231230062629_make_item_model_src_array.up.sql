ALTER TABLE items ADD COLUMN model_sources VARCHAR(150)[] DEFAULT '{}' NOT NULL;

DO $$
DECLARE
  itemRow RECORD;
BEGIN
  FOR itemRow IN SELECT * FROM items LOOP
    UPDATE items
        SET model_sources = ARRAY[itemRow.model_src]
        WHERE id = itemRow.id;
  END LOOP;
END $$;

ALTER TABLE items DROP COLUMN model_src;
