ALTER TABLE posts 
ADD COLUMN created_at timestamp,
ADD COLUMN updated_at timestamp,
ADD COLUMN deleted_at timestamp;

CREATE INDEX idx_deleted_at ON posts (deleted_at) WHERE deleted_at IS NOT NULL;