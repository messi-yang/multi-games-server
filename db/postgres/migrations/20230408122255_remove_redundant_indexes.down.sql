CREATE INDEX user_id ON users (id);
CREATE INDEX user_email_address ON users (id);
CREATE INDEX user_username ON users (id);

CREATE INDEX world_id ON worlds (id);
CREATE INDEX world_user_id ON worlds (game_user_id);

CREATE INDEX item_id ON items (id);
CREATE INDEX item_name ON items (name);
