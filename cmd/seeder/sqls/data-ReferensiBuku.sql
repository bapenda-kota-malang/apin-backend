--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0
-- Dumped by pg_dump version 13.0

-- Started on 2023-02-07 13:02:07

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE "apin-new";
--
-- TOC entry 4294 (class 1262 OID 125576)
-- Name: apin-new; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "apin-new" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';


ALTER DATABASE "apin-new" OWNER TO postgres;

\connect -reuse-previous=on "dbname='apin-new'"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4287 (class 0 OID 202142)
-- Dependencies: 613
-- Data for Name: ReferensiBuku; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."ReferensiBuku" (id, "Kode", "ThnAwal", "ThnAkhir", "NilaiMin", "NilaiMax", "Id", "LuasMinTipe", "LuasMaxTipe", "FaktorPembagiTipe", "Status", "CreatedAt", "UpdatedAt", "DeletedAt") FROM stdin;
1	1	2000	2020	0	100000	1	\N	\N	\N	\N	\N	\N	\N
2	2	2000	2020	100001	500000	2	\N	\N	\N	\N	\N	\N	\N
3	3	2000	2020	500001	2000000	3	\N	\N	\N	\N	\N	\N	\N
4	4	2000	2020	2000001	5000000	4	\N	\N	\N	\N	\N	\N	\N
5	5	2000	2020	5000001	999999999999999	5	\N	\N	\N	\N	\N	\N	\N
\.


--
-- TOC entry 4297 (class 0 OID 0)
-- Dependencies: 659
-- Name: ReferensiBuku_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."ReferensiBuku_Id_seq"', 5, true);


--
-- TOC entry 4298 (class 0 OID 0)
-- Dependencies: 612
-- Name: ReferensiBuku_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."ReferensiBuku_id_seq"', 1, false);


-- Completed on 2023-02-07 13:02:07

--
-- PostgreSQL database dump complete
--

