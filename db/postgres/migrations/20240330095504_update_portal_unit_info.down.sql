UPDATE units
SET info_snapshot = info_snapshot - 'target_unit_id'
WHERE type = 'portal';

ALTER TABLE portal_unit_infos DROP COLUMN target_unit_id;
