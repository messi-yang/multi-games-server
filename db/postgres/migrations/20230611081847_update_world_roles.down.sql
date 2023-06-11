ALTER TABLE user_world_roles ALTER COLUMN world_role TYPE varchar(255);

DROP TYPE world_role;
CREATE TYPE world_role_name AS ENUM ('admin');

UPDATE user_world_roles SET world_role = 'admin' WHERE world_role = 'owner';
ALTER TABLE user_world_roles ALTER COLUMN world_role TYPE world_role_name USING world_role::world_role_name;
