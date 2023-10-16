UPDATE units SET info_snapshot = to_jsonb('null'::jsonb) WHERE info_snapshot IS NULL;

ALTER TABLE units ALTER COLUMN info_snapshot SET NOT NULL;
