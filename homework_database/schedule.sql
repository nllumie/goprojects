-- Table: public.schedule

-- DROP TABLE IF EXISTS public.schedule;

CREATE TABLE IF NOT EXISTS public.schedule
(
    schedule_id integer NOT NULL DEFAULT nextval('schedule_schedule_id_seq'::regclass),
    group_id integer,
    subject character varying(50) COLLATE pg_catalog."default",
    time_slot character varying(30) COLLATE pg_catalog."default",
    CONSTRAINT schedule_pkey PRIMARY KEY (schedule_id),
    CONSTRAINT schedule_group_id_fkey FOREIGN KEY (group_id)
        REFERENCES public.groups (group_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.schedule
    OWNER to postgres;
