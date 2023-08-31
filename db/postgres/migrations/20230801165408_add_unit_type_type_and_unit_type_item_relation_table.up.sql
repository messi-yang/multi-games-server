CREATE TYPE unit_type AS ENUM ('static', 'portal');

ALTER TABLE items ADD COLUMN compatible_unit_type unit_type NOT NULL DEFAULT 'static';

CREATE INDEX item_compatible_unit_type ON items (compatible_unit_type);

ALTER TABLE units ADD COLUMN "type" unit_type NOT NULL DEFAULT 'static';

CREATE INDEX unit_type ON units ("type");
