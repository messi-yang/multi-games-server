CREATE TABLE link_unit_infos (
    world_id uuid NOT NULL,
    url VARCHAR (2048),
    id uuid NOT NULL
);

CREATE TABLE unit_types (
    name VARCHAR(20) PRIMARY KEY
);

INSERT INTO unit_types (name)
    VALUES
        ('static'),
        ('portal'),
        ('fence'),
        ('link');

ALTER TABLE units
    ALTER COLUMN "type"
    TYPE VARCHAR(20),
    ALTER COLUMN "type"
    SET DEFAULT 'static',
    ADD CONSTRAINT units_type_unit_types_name
	FOREIGN KEY ("type") REFERENCES unit_types(name);
    
ALTER TABLE items
    ALTER COLUMN "compatible_unit_type"
    TYPE VARCHAR(20),
    ALTER COLUMN "compatible_unit_type"
    SET DEFAULT 'static',
    ADD CONSTRAINT items_compatible_unit_type_unit_types_name
	FOREIGN KEY ("compatible_unit_type") REFERENCES unit_types(name);

DROP TYPE unit_type;
