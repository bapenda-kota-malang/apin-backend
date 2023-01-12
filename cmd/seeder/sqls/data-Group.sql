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
-- Data for Name: Group; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Group" ("Id", "Name", "Description", "Status") FROM stdin;
1	Administrator	admin full akses menu	1
3	Pendaftaran		1
4	Pendataan		1
5	Penetapan	penetapan 	1
6	Bendahara		1
38	Loket AT Hiburan	Pajak Air Tanah dan Hiburan	1
68	Penagihan	penagihan	1
40	Loket Hotel Parkir	Pajak Hotel dan Parkir	1
41	Loket Resto PPJ	Pajak Restoran dan PPJ	1
69	UPT	UPT KECAMATAN	1
43	Reklame	Pajak Reklame	1
44	Kasir	SSPD Jambong	1
106	Revoked	Ex Pemegang Acc Simpatda	1
46	Pendataan2	Kepala Seksi Pendataan	1
123	ReklameHotel	ReklameHotel	1
126	read only reklame	read only reklame	1
142	Resto	Resto	1
155	SKPDKB	SKPDKB	1
170	KPK	KPK	1
175	Pembukuan	Pembukuan	1
178	ReklameBendahara	ReklameBendahara	1
179	daftarbenda	Pendaftaranbendahargita	1
180	Stafpenetapan	Stafpenetapan	1
184	kabidp3	kabidp3	1
189	RO_Tax	ReadOnly SPTSKPSSPD	1
194	ReklameRead	ReklameRead	1
195	PPJ_Parkir_Readonly	PPJ_Parkir_Readonly	1
196	KASUBIDPD2	KASUBIDPD2	1
200	PendaftaranRO		1
204	LoketFull	LoketFull	1
213	Inspektorat	Inspektorat Readonly	1
\.


--
-- Name: Group_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Group_Id_seq"', 213, false);


--
-- PostgreSQL database dump complete
--

