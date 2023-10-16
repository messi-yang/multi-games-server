ALTER TABLE units ALTER COLUMN info_snapshot DROP NOT NULL;

UPDATE units SET info_snapshot = NULL WHERE info_snapshot = to_jsonb('null'::jsonb);
