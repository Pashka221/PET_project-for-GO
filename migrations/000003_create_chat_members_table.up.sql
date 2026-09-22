CREATE TABLE chat_members (
    chat_id UUID REFERENCES chats(id) ON DELETE CASCADE NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    joined_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    PRIMARY KEY (chat_id, user_id)
);