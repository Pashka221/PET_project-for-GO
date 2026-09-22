CREATE TABLE chats (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    is_group BOOLEAN DEFAULT false NOT NULL,
    name VARCHAR(50) CHECK(is_group = false OR name IS NOT NULL),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);