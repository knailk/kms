CREATE TABLE IF NOT EXISTS public.chat_sessions
(
    id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    is_deleted bigint,
    CONSTRAINT chat_sessions_pkey PRIMARY KEY (id)
);
