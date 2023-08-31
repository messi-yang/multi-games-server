DROP INDEX unit_type;

ALTER TABLE units DROP COLUMN "type";

DROP INDEX item_compatible_unit_type;

ALTER TABLE items DROP COLUMN compatible_unit_type;

DROP TYPE unit_type;
