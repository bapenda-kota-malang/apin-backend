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
-- Data for Name: KlasifikasiJalan; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."KlasifikasiJalan" ("Id", "Kode") FROM stdin;
1	A
2	B
3	C
4	D
\.


--
-- Name: KlasifikasiJalan_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."KlasifikasiJalan_Id_seq"', 4, true);


--
-- PostgreSQL database dump complete
--

