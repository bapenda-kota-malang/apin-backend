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
-- Data for Name: Jabatan; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Jabatan" ("Id", "Kode", "Nama") FROM stdin;
1	\N	BENDAHARA PENERIMAAN
2	\N	KASI PELAYANAN, PENGAWASAN DAN PENYELESAIAN SENGKETA
3	\N	KASI PEMBUKUAN DAN PELAPORAN
4	\N	KASI PENAGIHAN PAJAK BUMI DAN BANGUNAN
5	\N	KASI PENAGIHAN PAJAK DAERAH LAINNYA
6	\N	KASI PENDAFTARAN
7	\N	KASI PENDATAAN, PENILAIAN DAN PENETAPAN
8	\N	KASI PENETAPAN
9	\N	KASI PENGELOLAAN BENDA BERHARGA
10	\N	KASI PENGEMBANGAN POTENSI
11	\N	KASI PENGOLAHAN DATA
12	\N	KASI PENYELESAIAN KEBERATAN PAJAK DAERAH LAINNYA
13	\N	KEPALA BIDANG PAJAK DAERAH
14	\N	KEPALA BIDANG PEMBUKUAN DAN PENGEMBANGAN POTENSI
15	\N	KEPALA BIDANG PENAGIHAN DAN PEMERIKSAAN
16	\N	KEPALA BIDANG PENDATAAN, PENDAFTARAN, DAN PENETAPAN
17	\N	KEPALA BIDANG PENGEMBANGAN POTENSI
18	\N	KEPALA DINAS
19	\N	KEPALA SUB BIDANG PAJAK DAERAH II
20	\N	SEKRETARIS
21	\N	SUBAG PERENCANAAN DAN KEUANGAN
22	\N	SUBAG SUNGRAM
23	\N	SUBAG UMUM
\.


--
-- Name: Jabatan_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Jabatan_Id_seq"', 23, true);


--
-- PostgreSQL database dump complete
--

