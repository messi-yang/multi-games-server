ALTER TABLE users ADD COLUMN friendly_name VARCHAR (20) NOT NULL DEFAULT '';

DO $$
DECLARE
  userRow RECORD;
BEGIN
  FOR userRow IN SELECT * FROM users LOOP
    UPDATE users SET friendly_name = userRow.username WHERE
        id = userRow.id;
  END LOOP;
END $$;
