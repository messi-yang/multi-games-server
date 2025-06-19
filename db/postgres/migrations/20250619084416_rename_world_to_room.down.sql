ALTER TABLE game_accounts RENAME TO world_accounts;
ALTER TABLE world_accounts RENAME COLUMN rooms_count TO worlds_count;
ALTER TABLE world_accounts RENAME COLUMN rooms_count_limit TO worlds_count_limit;

ALTER TABLE room_members RENAME TO world_members;
ALTER TABLE world_members RENAME COLUMN room_id TO world_id;

ALTER TABLE rooms RENAME TO worlds;