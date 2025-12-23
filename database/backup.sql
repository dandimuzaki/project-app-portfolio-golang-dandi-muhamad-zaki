--
-- PostgreSQL database dump
--

\restrict B4fs6ym9bAYkXU9HLiYzOybejsRwXRxANxgLgMe20RYgTZ4bcZLl6gJWEGpnsKq

-- Dumped from database version 18.0
-- Dumped by pg_dump version 18.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: likes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.likes (
    id integer NOT NULL,
    user_id integer NOT NULL,
    project_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.likes OWNER TO postgres;

--
-- Name: likes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.likes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.likes_id_seq OWNER TO postgres;

--
-- Name: likes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.likes_id_seq OWNED BY public.likes.id;


--
-- Name: projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.projects (
    id integer NOT NULL,
    user_id integer NOT NULL,
    title text NOT NULL,
    description text,
    url text,
    image text,
    tags text[],
    tech_stack text[],
    is_published boolean DEFAULT true,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.projects OWNER TO postgres;

--
-- Name: projects_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.projects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.projects_id_seq OWNER TO postgres;

--
-- Name: projects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.projects_id_seq OWNED BY public.projects.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(200) NOT NULL,
    email character varying(200) NOT NULL,
    password character varying(200) NOT NULL,
    avatar text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    github text,
    linkedin text,
    cv_link text,
    description text,
    phone_number text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: likes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes ALTER COLUMN id SET DEFAULT nextval('public.likes_id_seq'::regclass);


--
-- Name: projects id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects ALTER COLUMN id SET DEFAULT nextval('public.projects_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.likes (id, user_id, project_id, created_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.projects (id, user_id, title, description, url, image, tags, tech_stack, is_published, created_at, updated_at, deleted_at) FROM stdin;
7	1	ShuttleTime — Smart Badminton Court Reservation System	ShuttleTime is a web-based reservation system that simplifies the process of booking badminton courts for players, clubs, and sports centers. Users can check real-time court availability, select time slots, and make reservations without manual coordination or phone calls. The application prevents double bookings by implementing strict scheduling logic and validation.\r\nAdministrators can manage courts, pricing, operational hours, and reservations from a centralized dashboard. ShuttleTime improves operational efficiency for court owners while providing a frictionless booking experience for players. This project highlights real-world scheduling logic, date-time handling, role separation, and transaction consistency.	https://badminton-app-kohl.vercel.app/	http://localhost:8080/public/uploads/1_0_WhatsApp Image 2025-12-23 at 21.36.33.jpeg	\N	{Golang,NextJS,Typescript,Gin,PostgreSQL}	t	2025-12-23 21:58:45.924187	2025-12-23 21:59:59.484816	\N
4	1	test			http://localhost:8080/public/uploads/1_4_anete-lusina-zwsHjakE_iI-unsplash.jpg	{""}	{Golang,PostgreSQL}	t	2025-12-22 23:01:02.299856	2025-12-23 16:07:17.677752	2025-12-23 22:00:52.689646
3	1	test				{inventory,golang,cli}	{golang,django,laravel}	t	2025-12-22 22:06:31.984505	2025-12-22 22:06:31.984505	2025-12-22 22:46:24.869398
2	1	test	aaa	https://youtube.com/shorts/1ZR_1NxwRHs?si=AWpxisKu0ipWSk8I		{golang,laravel,nextjs}	{golang,laravel,reactjs}	t	2025-12-22 22:01:28.963177	2025-12-22 22:42:58.071909	2025-12-22 22:49:31.649835
5	1	Agripacul — Farm-to-Table Agricultural E-Commerce Platform	Agripacul is a farm-to-table e-commerce platform designed to connect local farmers directly with consumers and small businesses. The platform focuses on transparency, product traceability, and fair pricing by reducing unnecessary middlemen in the agricultural supply chain. Users can browse fresh produce, view detailed product descriptions including origin and harvest date, and place orders seamlessly through an intuitive shopping experience.\r\n\r\nOn the seller side, farmers can manage product listings, update stock availability, and track incoming orders in real time. Agripacul is built with scalability in mind, supporting future features such as logistics integration, payment gateways, and product recommendations. The project demonstrates full-stack e-commerce fundamentals including authentication, role-based access, CRUD operations, and responsive UI design.	https://agripacul.vercel.app/	http://localhost:8080/public/uploads/1_5_WhatsApp Image 2025-12-23 at 11.29.44.jpeg	\N	{ReactJS,NodeJS,ExpressJS,MongoDB,Cloudinary}	t	2025-12-23 05:54:16.807951	2025-12-23 11:50:49.368271	\N
1	1	Inventory Management App with Golang CLI	Test	<nil>	http://localhost:8080/public/uploads/1_1_WhatsApp Image 2025-12-13 at 21.32.52.jpeg	\N	{Golang}	t	2025-12-20 15:26:28.626892	2025-12-23 12:08:43.807683	\N
6	2	MaterialKu	Mobile-view web app that has purpose to become an e-commerce application. The app was built using React and Node framework. Its main functionality is to manage user requests for easy transactions. This e-commerce is especially built for material product used for building. 	https://github.com/dandimuzaki	http://localhost:8080/public/uploads/2_0_QuickDBD-export.png	\N	{ReactJS,NodeJS}	t	2025-12-23 17:30:15.42018	2025-12-23 17:30:15.42018	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, avatar, created_at, updated_at, deleted_at, github, linkedin, cv_link, description, phone_number) FROM stdin;
2	Firizky Ardiansyah	firizky@example.com	password	http://localhost:8080/public/uploads/2_014_Rizka Kalsani.png	2025-12-22 14:46:33.972914	2025-12-23 17:32:56.595108	\N	https://github.com/dandimuzaki	https://github.com/dandimuzaki	\N	Full Stack Developer	081245332187
1	Dandi Muhamad Zaki	dandimuzaki@gmail.com	password	http://localhost:8080/public/uploads/1_CD-108.jpg	2025-12-19 23:15:22.493208	2025-12-23 22:01:48.30732	\N	https://github.com/dandimuzaki/	https://id.linkedin.com/in/dandimuzaki	http://localhost:8080/public/uploads/1_CV_Dandi Muhamad Zaki_30-11-2025.pdf	IT & Software Engineer with a strong foundation in full-stack development, machine learning, and data-driven problem-solving. Graduated with a 3.90 GPA from Institut Teknologi Bandung with experience in building scalable web systems, secure API solutions, and user-friendly interfaces. Proven adaptability through cross-domain projects in entrepreneurship, e-commerce, and sustainability. Skilled in collaboration and leadership, with hands-on experience leading student organizations and startup initiatives. Passionate about driving digital transformation by delivering secure, efficient, and innovative IT solutions.	085117388153
\.


--
-- Name: likes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.likes_id_seq', 1, false);


--
-- Name: projects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.projects_id_seq', 7, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: likes likes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_pkey PRIMARY KEY (id);


--
-- Name: likes likes_user_id_project_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_user_id_project_id_key UNIQUE (user_id, project_id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: likes likes_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- Name: likes likes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

\unrestrict B4fs6ym9bAYkXU9HLiYzOybejsRwXRxANxgLgMe20RYgTZ4bcZLl6gJWEGpnsKq

