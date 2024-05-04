-- Table: public.user_requesteds

-- DROP TABLE IF EXISTS public.user_requesteds;

CREATE TABLE IF NOT EXISTS public.user_requesteds
(
    id text COLLATE pg_catalog."default" NOT NULL,
    username text COLLATE pg_catalog."default",
    full_name text COLLATE pg_catalog."default",
    parent_name text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default",
    phone_number text COLLATE pg_catalog."default",
    birth_date timestamp with time zone,
    gender text COLLATE pg_catalog."default",
    status text COLLATE pg_catalog."default",
    class_id text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT user_requesteds_pkey PRIMARY KEY (id)
);