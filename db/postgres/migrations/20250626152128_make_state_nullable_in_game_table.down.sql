UPDATE games SET state = '{}' WHERE state IS NULL;
ALTER TABLE games ALTER COLUMN state SET NOT NULL;