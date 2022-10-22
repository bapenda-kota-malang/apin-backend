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
-- Data for Name: TarifReklame; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."TarifReklame" ("Id", "JenisMasa", "JenisReklame", "DasarPengenaan", "KlasifikasiJalan_Id", "MasaPajak", "Tarif") FROM stdin;
31	1	Kendaraan dan Sejenisnya	Unit	\N	\N	840000.00
32	2	Udara atau Balon	Unit	\N	Bulan	6000000.00
33	2	Slide Film	Unit	\N	Bulan	1260000.00
34	2	Apung	Unit	\N	Bulan	6000000.00
35	2	Flagchain	Luas	\N	Bulan	27000.00
36	3	Spanduk, Umbul-umbul, V.Banner, Sun Screen	Luas	\N	Bulan	120000.00
37	3	Spanduk, Umbul-umbul, V.Banner, Sun Screen	Luas	\N	Minggu	30000.00
38	3	Spanduk, Umbul-umbul, V.Banner, Sun Screen	Luas	\N	Hari	6000.00
39	3	Baliho	Luas	\N	Bulan	315000.00
40	3	Baliho	Luas	\N	Minggu	78750.00
41	3	Baliho	Luas	\N	Hari	15750.00
42	3	Tenda	Unit	\N	Bulan	120000.00
43	3	Tenda	Unit	\N	Minggu	30000.00
44	3	Tenda	Unit	\N	Hari	6000.00
45	3	Pengarah Jalur	Unit	\N	Hari	7500.00
46	4	Peragaan/Pameran, Demo	Unit	\N	\N	750000.00
47	4	Poster, Stiker dan Sejenisnya	Luas	\N	\N	45000.00
48	4	Selebaran, Leaflet/Brosur Berwarna	Unit	\N	\N	2400.00
49	4	Selebaran, Leaflet/Brosur Tidak Berwarna	Unit	\N	\N	900.00
54	1	Kendaraan Dan Sejenisnya	Luas	\N	\N	840000.00
1	1	Billboard, Tembok/Tugu, Shop Panel, Letter, Neon Sign, Prismatek	Luas	1	\N	480000.00
2	1	Billboard, Tembok/Tugu, Shop Panel, Letter, Neon Sign, Prismatek	Luas	2	\N	405000.00
3	1	Billboard, Tembok/Tugu, Shop Panel, Letter, Neon Sign, Prismatek	Luas	3	\N	367500.00
4	1	Billboard, Tembok/Tugu, Shop Panel, Letter, Neon Sign, Prismatek	Luas	4	\N	345000.00
5	1	Megatron, TV Media	Luas	1	\N	3900000.00
6	1	Megatron, TV Media	Luas	2	\N	3750000.00
7	1	Megatron, TV Media	Luas	3	\N	3600000.00
8	1	Megatron, TV Media	Luas	4	\N	3450000.00
9	1	Display Board	Luas	1	\N	3000000.00
10	1	Display Board	Luas	2	\N	2760000.00
11	1	Display Board	Luas	3	\N	2610000.00
12	1	Display Board	Luas	4	\N	2520000.00
13	1	Bando Jalan, JPO, Taman Gantung	Luas	1	\N	1050000.00
14	1	Bando Jalan, JPO, Taman Gantung	Luas	2	\N	960000.00
15	1	Bando Jalan, JPO, Taman Gantung	Luas	3	\N	870000.00
16	1	Bando Jalan, JPO, Taman Gantung	Luas	4	\N	750000.00
17	1	Neon Box	Luas	1	\N	750000.00
18	1	Neon Box	Luas	2	\N	705000.00
19	1	Neon Box	Luas	3	\N	660000.00
20	1	Neon Box	Luas	4	\N	615000.00
21	1	Mini Jumbo Board, Bus Shelter	Luas	1	\N	330000.00
22	1	Mini Jumbo Board, Bus Shelter	Luas	2	\N	285000.00
23	1	Mini Jumbo Board, Bus Shelter	Luas	3	\N	255000.00
24	1	Mini Jumbo Board, Bus Shelter	Luas	4	\N	240000.00
25	1	Rombong	Luas	1	\N	570000.00
26	1	Rombong	Luas	2	\N	480000.00
27	1	Rombong	Luas	3	\N	450000.00
28	1	Rombong	Luas	4	\N	405000.00
50	1	Billboard Disinari	Luas	1	\N	900000.00
51	1	Billboard Disinari	Luas	2	\N	810000.00
52	1	Billboard Disinari	Luas	3	\N	720000.00
53	1	Billboard Disinari	Luas	4	\N	630000.00
\.


--
-- Name: TarifReklame_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."TarifReklame_Id_seq"', 1, false);


--
-- PostgreSQL database dump complete
--

