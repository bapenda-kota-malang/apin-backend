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
-- Data for Name: Provinsi; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Provinsi" ("Id", "Kode", "Nama") FROM stdin;
1	11	ACEH
2	12	SUMATERA UTARA
3	13	SUMATERA BARAT
4	14	RIAU
5	15	JAMBI
6	16	SUMATERA SELATAN
7	17	BENGKULU
8	18	LAMPUNG
9	19	KEPULAUAN BANGKA BELITUNG
10	21	KEPULAUAN RIAU
11	31	DKI JAKARTA
12	32	JAWA BARAT
13	33	JAWA TENGAH
14	34	DI YOGYAKARTA
15	35	JAWA TIMUR
16	36	BANTEN
17	51	BALI
18	52	NUSA TENGGARA BARAT
19	53	NUSA TENGGARA TIMUR
20	61	KALIMANTAN BARAT
21	62	KALIMANTAN TENGAH
22	63	KALIMANTAN SELATAN
23	64	KALIMANTAN TIMUR
24	65	KALIMANTAN UTARA
25	71	SULAWESI UTARA
26	72	SULAWESI TENGAH
27	73	SULAWESI SELATAN
28	74	SULAWESI TENGGARA
29	75	GORONTALO
30	76	SULAWESI BARAT
31	81	MALUKU
32	82	MALUKU UTARA
33	91	PAPUA BARAT
34	94	PAPUA
\.


--
-- Name: Provinsi_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Provinsi_Id_seq"', 34, true);


--
-- PostgreSQL database dump complete
--

