ALTER TABLE chat ALTER COLUMN created_at SET NOT NULL;
ALTER TABLE worlds ALTER COLUMN created_at SET NOT NULL;
ALTER TABLE world_admins ALTER COLUMN created_at SET NOT NULL;
ALTER TABLE characters ALTER COLUMN created_at SET NOT NULL;