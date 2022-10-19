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
-- Data for Name: TarifJambong; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."TarifJambong" ("Id", "JenisReklame", "Nominal") FROM stdin;
1	Rombong	12375.00
2	Bando Jalan	24375.00
3	Reklame JPO	28125.00
6	Megatron/Videotron	28125.00
\.


--
-- Name: TarifJambong_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."TarifJambong_Id_seq"', 1, false);


--
-- PostgreSQL database dump complete
--

