ALTER TYPE unit_type RENAME TO old_unit_type;

CREATE TYPE unit_type AS ENUM ('static', 'portal');

ALTER TABLE units ADD COLUMN temp_type unit_type NOT NULL DEFAULT 'static';
DELETE FROM units WHERE type = 'fence';
UPDATE units
SET temp_type = CASE
   WHEN type = 'static' THEN 'static'::unit_type
   WHEN type = 'portal' THEN 'portal'::unit_type
END;
ALTER TABLE units DROP COLUMN type;
ALTER TABLE units RENAME COLUMN temp_type TO type;

ALTER TABLE items ADD COLUMN temp_compatible_unit_type unit_type NOT NULL DEFAULT 'static';
DELETE FROM items WHERE compatible_unit_type = 'fence';
UPDATE items
SET temp_compatible_unit_type = CASE
   WHEN compatible_unit_type = 'static' THEN 'static'::unit_type
   WHEN compatible_unit_type = 'portal' THEN 'portal'::unit_type
END;
ALTER TABLE items DROP COLUMN compatible_unit_type;
ALTER TABLE items RENAME COLUMN temp_compatible_unit_type TO compatible_unit_type;


DROP TYPE old_unit_type;
