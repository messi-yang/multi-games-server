ALTER TABLE portal_unit_infos
DROP COLUMN target_pos_x,
DROP COLUMN target_pos_z;

UPDATE units
SET info_snapshot = info_snapshot - 'targetPosition'
WHERE type = 'portal';
