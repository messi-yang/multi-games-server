ALTER TABLE units ALTER COLUMN info_id DROP NOT NULL;

ALTER TABLE units DROP CONSTRAINT units_unique_info_id;

UPDATE units SET info_id = NULL WHERE type = 'static' OR type = 'fence';
