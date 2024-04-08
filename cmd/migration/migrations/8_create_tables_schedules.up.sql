CREATE TABLE IF NOT EXISTS public.schedules
(
    id text COLLATE pg_catalog."default" NOT NULL,
    class_id text COLLATE pg_catalog."default",
    from_time timestamp with time zone,
    to_time timestamp with time zone,
    action text COLLATE pg_catalog."default",
    date bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT schedules_pkey PRIMARY KEY (id),
    CONSTRAINT fk_classes_schedules FOREIGN KEY (class_id)
        REFERENCES public.classes (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);