CREATE TYPE "CheckInOutAction" AS ENUM ('check_in', 'check_out');

CREATE TABLE
    IF NOT EXISTS public.check_in_outs (
        id text COLLATE pg_catalog."default" NOT NULL,
        user_class_id text COLLATE pg_catalog."default",
        action "CheckInOutAction",
        date bigint,
        created_at timestamp
        with
            time zone,
            updated_at timestamp
        with
            time zone,
            CONSTRAINT check_in_outs_pkey PRIMARY KEY (id),
            CONSTRAINT fk_user_classes_check_in_out FOREIGN KEY (user_class_id) REFERENCES public.user_classes (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION,
            CONSTRAINT fk_user_classes_check_in_outs FOREIGN KEY (user_class_id) REFERENCES public.user_classes (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION
    );

ALTER TABLE IF EXISTS public.check_in_outs OWNER to admin;

-- Index: idx_check_in_outs_date
-- DROP INDEX IF EXISTS public.idx_check_in_outs_date;
CREATE INDEX IF NOT EXISTS idx_check_in_outs_date ON public.check_in_outs USING btree (date ASC NULLS LAST);

-- Index: user_date_unique
-- DROP INDEX IF EXISTS public.user_date_unique;
CREATE UNIQUE INDEX IF NOT EXISTS user_date_unique ON public.check_in_outs USING btree (
    user_class_id COLLATE pg_catalog."default" ASC NULLS LAST,
    date ASC NULLS LAST
);