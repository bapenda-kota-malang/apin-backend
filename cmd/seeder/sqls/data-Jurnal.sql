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
-- Data for Name: Jurnal; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Jurnal" ("Id", "Kode", "Nama") FROM stdin;
7	TBP-OA	Jurnal TBP Official Assesment
8	STS	Jurnal untuk modul STS
9	SKPD-PB	Jurnal untuk modul SKP
10	TBP-SA	Jurnal TBP Self Assesment
11	SKPDKB	Pengakuan Piutang Pendapatan Pada Pendapatan LO
\.


--
-- Name: Jurnal_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Jurnal_Id_seq"', 11, false);


--
-- PostgreSQL database dump complete
--

