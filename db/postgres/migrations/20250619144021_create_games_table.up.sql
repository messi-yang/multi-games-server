CREATE TABLE games (
    id UUID PRIMARY KEY,
    room_id UUID REFERENCES rooms(id),
    name VARCHAR(64) NOT NULL,
    started BOOLEAN NOT NULL,
    selected BOOLEAN NOT NULL,
    state JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
