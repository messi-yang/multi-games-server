ALTER TABLE units
ADD COLUMN color VARCHAR(7);

UPDATE units
SET info_snapshot = 'null'::jsonb,
    color = info_snapshot->>'color'
WHERE type = 'color';

DROP TABLE color_unit_infos;
