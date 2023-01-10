--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Ubuntu 14.6-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.6 (Ubuntu 14.6-1.pgdg20.04+1)

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
-- Data for Name: JenisPajak; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."JenisPajak" ("Id", "Kode", "Uraian", "Rekening_Id", "KodeBilling", "KodeVAJatim") FROM stdin;
1	A	Pajak Hotel / Kos / Sejenis	10100	35730010000	14604
2	B	Pajak Restoran	10110	35730020000	14605
3	C	Pajak Hiburan	10117	35730030000	14606
4	D	Pajak Reklame	10124	35730040000	14607
5	G	Pajak Air Tanah	10190	35730050000	14610
6	F	Pajak Parkir	10188	35730060000	14609
9	E	Pajak Penerangan Jalan	10131	35730090000	14608
\.


--
-- Name: JenisPajak_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."JenisPajak_Id_seq"', 9, false);


--
-- PostgreSQL database dump complete
--

