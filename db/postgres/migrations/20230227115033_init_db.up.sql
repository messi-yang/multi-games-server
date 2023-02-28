CREATE TABLE IF NOT EXISTS users (
	id uuid PRIMARY KEY,
	email_address VARCHAR (255) UNIQUE NOT NULL,
	username VARCHAR (50) UNIQUE NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX user_id ON users (id);
CREATE INDEX user_email_address ON users (id);
CREATE INDEX user_username ON users (id);

CREATE TABLE IF NOT EXISTS worlds (
	id uuid PRIMARY KEY,
	user_id uuid UNIQUE NOT NULL,
    name VARCHAR (50) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);
CREATE INDEX world_id ON worlds (id);
CREATE INDEX world_user_id ON worlds (user_id);

CREATE TABLE IF NOT EXISTS items (
	id uuid PRIMARY KEY,
	name VARCHAR (50) UNIQUE NOT NULL,
	traversable boolean NOT NULL,
	model_src VARCHAR (255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX item_id ON items (id);
CREATE INDEX item_name ON items (name);
