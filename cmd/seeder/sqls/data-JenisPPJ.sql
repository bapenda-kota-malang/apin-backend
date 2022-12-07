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
-- Data for Name: JenisPPJ; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."JenisPPJ" ("Id", "Tahun", "Jenis", "TarifPersen", "CreatedAt", "UpdatedAt", "DeletedAt") FROM stdin;
1	2022	SOSIAL	0	\N	\N	\N
2	2022	RUMAH TANGGA	7	\N	\N	\N
3	2022	BISNIS	5	\N	\N	\N
4	2022	INDUSTRI	3	\N	\N	\N
5	2022	PEMERINTAH	0	\N	\N	\N
6	\N	LAIN LAIN	0	\N	\N	\N
\.


--
-- Name: JenisPPJ_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."JenisPPJ_Id_seq"', 6, false);


--
-- PostgreSQL database dump complete
--

