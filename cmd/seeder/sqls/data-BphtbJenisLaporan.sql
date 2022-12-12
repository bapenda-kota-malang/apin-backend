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
-- Data for Name: BphtbJenisLaporan; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."BphtbJenisLaporan" ("Id", "Kode", "Nama") FROM stdin;
1	01	Jual Beli
2	02	Tukar Menukar
3	03	Hibah
4	04	Hibah Wasiat
5	05	Waris
6	06	Pemasukan dalam perseroan / badan hukum lainnya
7	07	Pemisahan hak yang mengakibatkan peralihan
8	08	Penunjukan pemberian dalam lelang
9	09	Pelaksanaan putusan hakim yang mempunyai kekuatan hukum tetap
10	10	Penggabungan usaha
11	11	Pelebaran usaha
12	12	Pemekaran usaha
13	13	Hadiah
14	14	Perolehan Hak Rumah Sederhana Sehat dan HSS melalui KPR bersubsidi
15	15	Pemberian hak baru
16	16	Pemberian hak baru sebagai kelanjutan pelepasan hak
17	17	Pemberian hak baru diluar pelepasan hak
18	22	Jenis Baru
\.


--
-- Name: BphtbJenisLaporan_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."BphtbJenisLaporan_Id_seq"', 18, true);


--
-- PostgreSQL database dump complete
--

