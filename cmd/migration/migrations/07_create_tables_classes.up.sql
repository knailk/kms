CREATE TABLE IF NOT EXISTS public.classes
(
    id text COLLATE pg_catalog."default" NOT NULL,
    teacher_id text COLLATE pg_catalog."default",
    driver_id text COLLATE pg_catalog."default",
    from_date bigint,
    to_date bigint,
    class_name text COLLATE pg_catalog."default",
    status text,
    description text,
    age_group bigint,
    price numeric,
    currency text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    is_deleted bigint,
    CONSTRAINT classes_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_classes_teacher_id
    ON public.classes USING btree
    (teacher_id COLLATE pg_catalog."default" ASC NULLS LAST);