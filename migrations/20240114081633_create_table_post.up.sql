CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.posts (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_uuid UUID NOT NULL,
    caption TEXT,
    image_url VARCHAR,
    location JSONB
);
CREATE INDEX IF NOT EXISTS idx_user_uuid ON posts (user_uuid);

ALTER TABLE public.posts OWNER TO "user";
GRANT ALL ON TABLE public.posts TO "user";