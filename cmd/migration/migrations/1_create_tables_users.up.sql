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

CREATE TABLE IF NOT EXISTS public.chat_sessions
(
    id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    is_deleted bigint,
    CONSTRAINT chat_sessions_pkey PRIMARY KEY (id)
)
-----------------------------------------------------------------------------------------------------------------

