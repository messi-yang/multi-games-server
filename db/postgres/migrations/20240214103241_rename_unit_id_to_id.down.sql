ALTER TABLE units DROP CONSTRAINT units_pkey;

ALTER TABLE units ADD CONSTRAINT units_unique_info_id UNIQUE (id);

ALTER TABLE units RENAME COLUMN id TO info_id;
