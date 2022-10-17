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
-- Data for Name: Skpd; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Skpd" ("Id", "Kode", "Nama", "Alamat", "Telp") FROM stdin;
112	1.20.08	Badan Pelayanan Pajak Daerah	Malang	\N
\.


--
-- Name: Skpd_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Skpd_Id_seq"', 1, false);


--
-- PostgreSQL database dump complete
--

