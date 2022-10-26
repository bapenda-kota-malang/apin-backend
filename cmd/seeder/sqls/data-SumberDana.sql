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
-- Data for Name: SumberDana; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."SumberDana" ("Id", "Rekening_Id", "Nama", "Bank", "NoRekeningBank", "UpdatedAt") FROM stdin;
4	10203	PAD - 0041 - 000266	Bank Jatim	0041 - 000266	2016-04-27 03:57:22+07
\.


--
-- Name: SumberDana_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."SumberDana_Id_seq"', 4, false);


--
-- PostgreSQL database dump complete
--

