--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-1.pgdg20.04+1)

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
-- Data for Name: TarifJambongRek; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."TarifJambongRek" ("Id", "Tipe", "BatasBawah", "BatasAtas", "Nominal") FROM stdin;
2	RB	0	5	6375.00
3	RB	5.1	8	10500.00
4	RB	8.1	-	12375.00
5	RNB	0	5	10500.00
7	RNB	5.1	-	12375.00
8	Insidentil	\N	\N	2.50
\.


--
-- Name: TarifJambongRek_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."TarifJambongRek_Id_seq"', 1, false);


--
-- PostgreSQL database dump complete
--

