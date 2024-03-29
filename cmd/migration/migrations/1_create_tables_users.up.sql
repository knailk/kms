CREATE TYPE "UserRole" AS ENUM ('admin', 'student', 'teacher', 'chef', 'driver');
CREATE TABLE IF NOT EXISTS public.users
(
    username text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default",
    role "UserRole",
    full_name text COLLATE pg_catalog."default",
    gender text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default",
    birth_date timestamp with time zone,
    phone_number text COLLATE pg_catalog."default",
    picture_url text COLLATE pg_catalog."default",
    address text COLLATE pg_catalog."default",
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    is_deleted boolean DEFAULT false,
    CONSTRAINT users_pkey PRIMARY KEY (username),
    CONSTRAINT uni_users_email UNIQUE (email)
);
CREATE INDEX idx_role_id ON users (role);

CREATE TYPE "MessageType" AS ENUM ('text', 'image', 'video', 'audio', 'file', 'link', 'voice', 'sticker');
-----------------------------------------------------------------------------------------------------------------
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
