--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-2.pgdg20.04+2)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-2.pgdg20.04+2)

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
-- Data for Name: Jpb; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Jpb" ("Id", "Kode", "Nama") FROM stdin;
1	00	TANAH DIBANGUN
2	01	PERUMAHAN
3	02	PERKANTORAN SWASTA
4	03	PABRIK
5	04	TOKO/APOTIK/PASAR/RUKO
6	05	RUMAH SAKIT/KLINIK
7	06	OLAH RAGA/REKREASI
8	07	HOTEL/WISMA
9	08	BENGKEL/GUDANG/PERTANIAN
10	09	GEDUNG PEMERINTAH
11	10	LAIN-LAIN
12	11	BANGUNAN TIDAK KENA PAJAK
13	12	BANGUNAN PARKIR
14	13	APARTEMEN
15	14	POMPA BENSIN
16	15	TANGKI MINYAK
17	16	GEDUNG SEKOLAH
18	50	TANAH KOSONG
\.


--
-- Name: Jpb_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Jpb_Id_seq"', 18, true);


--
-- PostgreSQL database dump complete
--

