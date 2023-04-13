
CREATE TABLE public."Users" (
    "Id" integer NOT NULL,
    "Name" text NOT NULL
);



CREATE SEQUENCE public."Users_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Users_Id_seq" OWNED BY public."Users"."Id";



CREATE TABLE public.labels (
    "Id" integer NOT NULL,
    name text NOT NULL
);



CREATE SEQUENCE public."labels_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."labels_Id_seq" OWNED BY public.labels."Id";



CREATE TABLE public.tasks (
    id integer NOT NULL,
    "Opened" bigint,
    "Closed" bigint DEFAULT 0,
    author_id integer DEFAULT 0,
    assigned_id integer,
    title text NOT NULL,
    content text
);



CREATE SEQUENCE public."tasks_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



ALTER SEQUENCE public."tasks_Id_seq" OWNED BY public.tasks.id;



CREATE TABLE public.tasks_labels (
    task_id integer,
    label_id integer
);


ALTER TABLE ONLY public."Users" ALTER COLUMN "Id" SET DEFAULT nextval('public."Users_Id_seq"'::regclass);


ALTER TABLE ONLY public.labels ALTER COLUMN "Id" SET DEFAULT nextval('public."labels_Id_seq"'::regclass);

ALTER TABLE ONLY public.tasks ALTER COLUMN id SET DEFAULT nextval('public."tasks_Id_seq"'::regclass);


INSERT INTO public."Users" VALUES (2, 'Jon');
INSERT INTO public."Users" VALUES (3, 'Дима');
INSERT INTO public."Users" VALUES (1, 'Ken');



INSERT INTO public.labels VALUES (1, 'Обучение');
INSERT INTO public.labels VALUES (2, 'Разработка');


INSERT INTO public.tasks VALUES (2, 123456, NULL, 1, 2, 'test', 'разработать');
INSERT INTO public.tasks VALUES (3, 2345678, 23569999, 3, 2, 'test', 'изучить');
INSERT INTO public.tasks VALUES (4, 222222, 5555555, 2, 3, 'test1', NULL);
INSERT INTO public.tasks VALUES (7, NULL, 0, 1, 3, 'Тестовая Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (8, NULL, 0, 1, 3, 'Тестовая Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (9, NULL, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (10, NULL, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (11, 1681368189, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (12, 1681368985, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (13, 1681369128, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (14, 1681369923, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (15, 1681370230, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (16, 1681370390, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (17, 1681370456, 0, 1, 3, 'Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (18, 1681370490, 0, 1, 3, 'Очень Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (19, 1681370610, 0, 1, 3, 'Совсем  очень-очень Важная Задача', 'Изучить pgx');
INSERT INTO public.tasks VALUES (20, 1681370651, 0, 1, 3, 'Совсем  очень-очень Важная Задача', 'Изучить pgx');



INSERT INTO public.tasks_labels VALUES (3, 1);
INSERT INTO public.tasks_labels VALUES (2, 2);
INSERT INTO public.tasks_labels VALUES (8, 1);



SELECT pg_catalog.setval('public."Users_Id_seq"', 3, true);


SELECT pg_catalog.setval('public."labels_Id_seq"', 2, true);


SELECT pg_catalog.setval('public."tasks_Id_seq"', 22, true);


ALTER TABLE ONLY public."Users"
    ADD CONSTRAINT "id_Users_pk" PRIMARY KEY ("Id");



ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT id_tasks_pk PRIMARY KEY (id);



ALTER TABLE ONLY public.labels
    ADD CONSTRAINT labels_pk PRIMARY KEY ("Id");



ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT "assigned_tasks_FK" FOREIGN KEY (assigned_id) REFERENCES public."Users"("Id");



ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT "author_tasks_FK" FOREIGN KEY (author_id) REFERENCES public."Users"("Id");



ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT "tasks_FK" FOREIGN KEY (assigned_id) REFERENCES public."Users"("Id");



ALTER TABLE ONLY public.tasks_labels
    ADD CONSTRAINT "tasks_labels_FK" FOREIGN KEY (task_id) REFERENCES public.tasks(id);



ALTER TABLE ONLY public.tasks_labels
    ADD CONSTRAINT "tasks_labels_FK_1" FOREIGN KEY (label_id) REFERENCES public.labels("Id");



