-- +goose Up
CREATE UNIQUE INDEX uq_group_members_one_owner_per_group ON group_members (group_id)
WHERE role = 'OWNER';
-- +goose Down
DROP INDEX uq_group_members_one_owner_per_group;
