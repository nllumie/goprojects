-- Table: public.students

-- DROP TABLE IF EXISTS public.students;

CREATE TABLE IF NOT EXISTS public.students
(
    id integer NOT NULL DEFAULT nextval('students_id_seq'::regclass),
    firstname text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default",
    mobile text COLLATE pg_catalog."default",
    gender text COLLATE pg_catalog."default",
    group_id integer,
    birth_date date,
    CONSTRAINT students_pkey PRIMARY KEY (id),
    CONSTRAINT fk_students_group FOREIGN KEY (group_id)
        REFERENCES public.groups (group_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.students
    OWNER to postgres;
