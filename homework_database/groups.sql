-- Table: public.groups

-- DROP TABLE IF EXISTS public.groups;

CREATE TABLE IF NOT EXISTS public.groups
(
    group_id integer NOT NULL DEFAULT nextval('groups_group_id_seq'::regclass),
    group_name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT groups_pkey PRIMARY KEY (group_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.groups
    OWNER to postgres;
