
CREATE TABLE IF NOT EXISTS public.user_classes
(
    username text COLLATE pg_catalog."default" NOT NULL,
    class_id text COLLATE pg_catalog."default" NOT NULL,
    status text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT user_classes_pkey PRIMARY KEY (username, class_id),
    CONSTRAINT fk_classes_user FOREIGN KEY (class_id)
        REFERENCES public.classes (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);