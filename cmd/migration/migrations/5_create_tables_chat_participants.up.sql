CREATE TABLE IF NOT EXISTS public.chat_participants
(
    id text COLLATE pg_catalog."default" NOT NULL,
    chat_session_id text COLLATE pg_catalog."default",
    username text COLLATE pg_catalog."default",
    is_owner boolean,
    created_at timestamp with time zone,
    is_deleted bigint,
    CONSTRAINT chat_participants_pkey PRIMARY KEY (id),
    CONSTRAINT fk_chat_sessions_chat_participants FOREIGN KEY (chat_session_id)
        REFERENCES public.chat_sessions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE UNIQUE INDEX IF NOT EXISTS "uniqueIndex"
    ON public.chat_participants USING btree
    (chat_session_id COLLATE pg_catalog."default" ASC NULLS LAST, username COLLATE pg_catalog."default" ASC NULLS LAST);