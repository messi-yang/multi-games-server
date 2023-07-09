CREATE TABLE players (
    id uuid NOT NULL,
    user_id uuid,
    world_id uuid NOT NULL,
    name character varying(50) NOT NULL,
    pos_x integer NOT NULL,
    pos_z integer NOT NULL,
    direction integer NOT NULL,
    held_item_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);

ALTER TABLE ONLY players
    ADD CONSTRAINT players_pkey PRIMARY KEY (id);

CREATE INDEX players_world_id_pos_x_pos_z ON players USING btree (world_id, pos_x, pos_z);

ALTER TABLE ONLY players
    ADD CONSTRAINT fk_item FOREIGN KEY (held_item_id) REFERENCES items(id);

ALTER TABLE ONLY players
    ADD CONSTRAINT fk_players_user FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE ONLY players
    ADD CONSTRAINT fk_world FOREIGN KEY (world_id) REFERENCES worlds(id);
