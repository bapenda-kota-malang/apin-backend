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
-- Data for Name: HargaDasarAir; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."HargaDasarAir" ("Id", "Peruntukan", "BatasBawah", "BatasAtas", "TarifMataAir", "TarifBukanMataAir") FROM stdin;
26	NON NIAGA	11001	999999	125.00	125.00
27	PDAM	0	999999	125.00	125.00
10	NON NIAGA	0	50	255.00	225.00
11	NON NIAGA	51	500	295.00	265.00
12	NON NIAGA	501	1000	330.00	300.00
13	NON NIAGA	1001	2500	370.00	340.00
14	NON NIAGA	2501	11000	405.00	375.00
15	NIAGA	0	50	380.00	350.00
16	NIAGA	51	500	480.00	450.00
17	NIAGA	501	1000	580.00	550.00
18	NIAGA	1001	2500	680.00	650.00
19	NIAGA	2500	999999	780.00	750.00
20	INDUSTRI	0	50	1580.00	1550.00
21	INDUSTRI	51	500	2280.00	2250.00
22	INDUSTRI	501	1000	2980.00	2950.00
23	INDUSTRI	1001	2500	3680.00	3650.00
25	INDUSTRI	2501	999999	4380.00	4350.00
\.


--
-- Name: HargaDasarAir_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."HargaDasarAir_Id_seq"', 27, false);


--
-- PostgreSQL database dump complete
--

