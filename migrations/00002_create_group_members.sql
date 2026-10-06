-- +goose Up
CREATE TABLE group_members (
    group_id UUID REFERENCES groups(id) NOT NULL,
    user_id UUID NOT NULL,
    role VARCHAR(100) CHECK(
        (role = 'OWNER')
        OR (role = 'ADMIN')
        OR (role = 'MEMBER')
    ) NOT NULL,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(group_id, user_id)
);
CREATE INDEX idx_group_members_user_id ON group_members(user_id);
-- +goose Down
DROP TABLE group_members;
