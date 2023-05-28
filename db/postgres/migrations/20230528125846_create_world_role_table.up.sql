CREATE TYPE world_role_name AS ENUM ('admin');
CREATE TABLE IF NOT EXISTS world_roles (
    id uuid PRIMARY KEY,
	world_id uuid NOT NULL,
    user_id uuid NOT NULL,
    name world_role_name NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_world FOREIGN KEY(world_id) REFERENCES worlds(id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    UNIQUE (world_id, user_id)
);
CREATE INDEX world_roles_world_id_user_id ON world_roles (world_id, user_id);

DO $$
DECLARE
  worldRow RECORD;
BEGIN
  FOR worldRow IN SELECT * FROM worlds LOOP
    INSERT INTO world_roles (
        id, world_id, user_id, name, created_at, updated_at
    ) VALUES (
        gen_random_uuid(), worldRow.id, worldRow.user_id, 'admin', NOW(), NOW()
    );
  END LOOP;
END $$;
