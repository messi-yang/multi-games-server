ALTER TABLE user_world_roles ALTER COLUMN world_role TYPE varchar(255);

DROP TYPE world_role_name;
CREATE TYPE world_role AS ENUM ('owner', 'admin', 'editor', 'viewer');

UPDATE user_world_roles SET world_role = 'owner' WHERE world_role = 'admin';
ALTER TABLE user_world_roles ALTER COLUMN world_role TYPE world_role USING world_role::world_role;
