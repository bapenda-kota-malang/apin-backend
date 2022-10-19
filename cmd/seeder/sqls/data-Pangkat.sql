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
-- Data for Name: Pangkat; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Pangkat" ("Id", "Kode", "Nama", "Golongan", "Ruang") FROM stdin;
1	I/A	JURU MUDA	1	A
2	I/B	JURU MUDA TINGKAT I	1	B
3	I/C	JURU MUDA	1	C
4	I/D	JURU TINGKAT I	1	D
5	II/A	PENGATUR MUDA	2	A
6	II/B	PENGATUR MUDA TINGKAT I	2	B
7	II/C	PENGATUR	2	C
8	II/D	PENGATUR TINGKAT I	2	D
9	III/A	PENATA MUDA	3	A
10	III/B	PENATA MUDA TINGKAT I	3	B
11	III/C	PENATA	3	C
12	III/D	PENATA TINGKAT I	3	D
13	IV/A	PEMBINA	4	A
14	IV/B	PEMBINA TINGKAT I	4	B
15	IV/C	PEMBINA UTAMA MUDA	4	C
16	IV/D	PEMBINA UTAMA MADYA	4	D
17	IV/E	PEMBINA UTAMA	4	E
\.


--
-- Name: Pangkat_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Pangkat_Id_seq"', 17, true);


--
-- PostgreSQL database dump complete
--

