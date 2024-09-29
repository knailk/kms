CREATE TABLE IF NOT EXISTS public.user_classes
(
    id text COLLATE pg_catalog."default" NOT NULL,
    username text COLLATE pg_catalog."default",
    class_id text COLLATE pg_catalog."default",
    status text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT user_classes_pkey PRIMARY KEY (id),
    CONSTRAINT fk_classes_user FOREIGN KEY (class_id)
        REFERENCES public.classes (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_classes_username_class_id_unique
    ON public.user_classes USING btree
    (username COLLATE pg_catalog."default" ASC NULLS LAST, class_id COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;