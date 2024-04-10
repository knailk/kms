
CREATE TABLE IF NOT EXISTS public.check_in_outs
(
    id text COLLATE pg_catalog."default" NOT NULL,
    username text COLLATE pg_catalog."default",
    action "CheckInOutAction",
    date bigint,
    class_id text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT check_in_outs_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_check_in_outs_class_id
    ON public.check_in_outs USING btree
    (class_id COLLATE pg_catalog."default" ASC NULLS LAST);

CREATE INDEX IF NOT EXISTS idx_check_in_outs_date
    ON public.check_in_outs USING btree
    (date ASC NULLS LAST);

CREATE INDEX IF NOT EXISTS idx_check_in_outs_username
    ON public.check_in_outs USING btree
    (username COLLATE pg_catalog."default" ASC NULLS LAST);