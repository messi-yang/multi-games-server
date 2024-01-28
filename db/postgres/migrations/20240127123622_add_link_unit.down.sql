CREATE TYPE unit_type AS ENUM ('static', 'portal', 'fence');

ALTER TABLE units
    DROP CONSTRAINT units_type_unit_types_name,
	ALTER COLUMN "type"
		DROP DEFAULT,
    ALTER COLUMN "type"
        TYPE unit_type USING type::unit_type,
    ALTER COLUMN "type"
        SET DEFAULT 'static'::unit_type;


ALTER TABLE items
    DROP CONSTRAINT items_compatible_unit_type_unit_types_name,
	ALTER COLUMN "compatible_unit_type"
		DROP DEFAULT,
    ALTER COLUMN "compatible_unit_type"
        TYPE unit_type USING compatible_unit_type::unit_type,
    ALTER COLUMN "compatible_unit_type"
        SET DEFAULT 'static'::unit_type;

DROP TABLE unit_types;

DROP TABLE link_unit_infos;
