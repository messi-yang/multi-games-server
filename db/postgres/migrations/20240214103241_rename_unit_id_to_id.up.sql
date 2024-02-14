ALTER TABLE units RENAME COLUMN info_id TO id;

ALTER TABLE units DROP CONSTRAINT units_unique_info_id;

ALTER TABLE units ADD PRIMARY KEY (id);
