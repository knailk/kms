
CREATE TYPE "MessageType" AS ENUM ('text', 'image', 'video', 'audio', 'file', 'link', 'voice', 'sticker');

CREATE TABLE IF NOT EXISTS public.chat_messages
(
    id text COLLATE pg_catalog."default" NOT NULL,
    chat_session_id text COLLATE pg_catalog."default",
    sender text COLLATE pg_catalog."default",
    message text COLLATE pg_catalog."default",
    type "MessageType",
    is_read boolean,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    is_deleted bigint,
    CONSTRAINT chat_messages_pkey PRIMARY KEY (id),
    CONSTRAINT fk_chat_sessions_chat_messages FOREIGN KEY (chat_session_id)
        REFERENCES public.chat_sessions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
-----------------------------------------------------------------------------------------------------------------