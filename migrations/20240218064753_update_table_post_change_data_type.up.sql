ALTER TABLE posts
ALTER COLUMN image_url TYPE JSONB
USING jsonb_build_array(image_url)::JSONB;

