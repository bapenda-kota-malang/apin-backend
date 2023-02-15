--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0
-- Dumped by pg_dump version 13.0

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
-- Data for Name: ReferensiBuku; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."ReferensiBuku" (Id, "Kode", "ThnAwal", "ThnAkhir", "NilaiMin", "NilaiMax", "Id", "LuasMinTipe", "LuasMaxTipe", "FaktorPembagiTipe", "Status", "CreatedAt", "UpdatedAt", "DeletedAt") FROM stdin;
1	1	2000	2020	0	100000	1	\N	\N	\N	\N	\N	\N	\N
2	2	2000	2020	100001	500000	2	\N	\N	\N	\N	\N	\N	\N
3	3	2000	2020	500001	2000000	3	\N	\N	\N	\N	\N	\N	\N
4	4	2000	2020	2000001	5000000	4	\N	\N	\N	\N	\N	\N	\N
5	5	2000	2020	5000001	999999999999999	5	\N	\N	\N	\N	\N	\N	\N
\.


--
-- Name: ReferensiBuku_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."ReferensiBuku_Id_seq"', 5, true);


--
-- Name: ReferensiBuku_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."ReferensiBuku_Id_seq"', 1, false);

--
-- PostgreSQL database dump complete
--

