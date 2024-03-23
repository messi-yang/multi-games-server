ALTER TABLE units
ADD COLUMN dimension_width INT NOT NULL DEFAULT 1,
ADD COLUMN dimension_depth INT NOT NULL DEFAULT 1;

DO $$
DECLARE
  itemRow RECORD;
BEGIN
  FOR itemRow IN SELECT * FROM items LOOP
    UPDATE units
    SET dimension_width = itemRow.dimension_width, dimension_depth = itemRow.dimension_depth
    WHERE item_id = itemRow.id;
  END LOOP;
END $$;
