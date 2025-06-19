ALTER TABLE worlds RENAME TO rooms;

ALTER TABLE world_members RENAME TO room_members;
ALTER TABLE room_members RENAME COLUMN world_id TO room_id;

ALTER TABLE world_accounts RENAME TO game_accounts;
ALTER TABLE game_accounts RENAME COLUMN worlds_count TO rooms_count;
ALTER TABLE game_accounts RENAME COLUMN worlds_count_limit TO rooms_count_limit;
