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
-- Data for Name: Menu; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Menu" ("Id", "Parent_Id", "Title", "Url", "Active") FROM stdin;
1	0	Beranda	#	1
2	0	Data Dasar	#	1
3	0	Pendaftaran	#	1
4	0	Bendahara	#	1
5	0	Laporan	laporan/	1
6	0	Setelan	#	1
7	2	Pejabat Daerah	pejabat_daerah/	1
8	2	Satuan Kerja	dinas/	1
9	2	Kecamatan	kecamatan/	1
11	2	Rekening	rekening/	1
12	2	Anggaran	anggaran/	1
13	61	Tarif Air Tanah	koefisien_air/	1
14	61	Jenis Usaha	jenis_usaha/	1
15	3	Wajib Pajak/Retribusi	pendaftaran/	1
16	30	Pajak Reklame	pajak_reklame/daftar_sa/	0
17	4	Pembayaran Self Assessment	bayar_sa/	0
18	34	Pajak Reklame	pajak_reklame/daftar_oa/	1
19	32	Penetapan OA	penetapan_oa/	1
20	\N	Pembayaran ke BKP	pembayaran/	1
21	4	Penyetoran Ke Kas Daerah	penyetoran/	0
22	4	Surat Tanda Setoran	sts/	1
23	4	Pembayaran Angsuran	angsuran/	0
24	4	Pembayaran Lain-lain	bayar_lain/	0
25	\N	Pemeriksaan Pajak/Retribusi	#	1
71	0	Penagihan	#	1
27	6	Group	group/	1
28	6	User	group/user/	1
29	6	Menu	group/menu/	1
30	0	SPTPD SA	#	1
32	0	SKPD/SKPDKB	#	1
33	4	Pembayaran Official Assessment	bayar_oa/	0
34	0	Penetapan OA	#	1
35	30	Pajak Hotel	pajak_hotel/daftar_sa/	1
36	30	Pajak Restoran	pajak_restoran/daftar_sa/	1
37	30	Pajak Hiburan	pajak_hiburan/daftar_sa/	1
38	30	Pajak Parkir	pajak_parkir/daftar_sa/	1
64	2	Sumber Dana	sumber_dana/	1
40	34	Pajak Air Tanah	pajak_air/daftar_oa/	1
41	61	Konfigurasi Pajak	master_pajak/	1
42	30	Pajak Air Tanah	pajak_air/daftar_sa/	1
43	34	Pajak Hiburan	pajak_hiburan/daftar_oa/	1
44	34	Pajak Hotel	pajak_hotel/daftar_oa/	1
65	61	Tarif Jambong	tarif_jambong/	1
46	34	Pajak Restoran	pajak_restoran/daftar_oa/	1
47	34	Pajak Parkir	pajak_parkir/daftar_oa/	1
66	0	JamBong	#	1
50	0	FITURE	#	1
51	50	Pemeriksaan Pajak	feature/pemeriksaan	0
52	50	Report Aging	feature/aging	0
53	50	Reminder	reminder/	1
54	32	Penetapan SA	penetapan_sa/	1
55	50	Surat Teguran	surat_teguran/	1
56	4	Tanda Bukti Pembayaran	tbp/	1
57	2	Bendahara	bendahara/	1
58	2	Jurnal	jurnal/	1
48	30	Pajak Penerangan Jalan	pajak_penerangan/daftar_sa/	1
49	34	Pajak Penerangan Jalan	pajak_penerangan/daftar_oa/	1
59	61	Jenis Pajak	jenis_pajak/	1
60	61	Koefisien Reklame	koefisien_reklame/	1
61	0	Konfigurasi	#	1
67	66	Jaminan Bongkar	jaminan_bongkar/	1
68	66	Proses JamBong	proses_jaminan_bongkar/	1
69	0	Utilitas	#	1
70	69	Transfer Pendapatan	transfer_pendapatan/	1
72	71	Penagihan Pajak	penagihan/	1
73	0	ANGSURAN	#	1
75	73	Angsuran Denda	angsuran	1
76	0	Penagihan Pajak	penagihan/	1
82	50	MASAL OA (2%)	massal_oa/	1
85	0	MASAL OA		1
84	50	Masal SA 2%	massal_sa/	1
86	85	AIR TANAH	massal_oa_air/	1
87	85	PPJ	massal_oa_jalan/	1
88	85	Parkir	massal_oa_parkir/	1
89	85	Restoran	massal_oa_restoran/	1
90	50	Copy SPT	salin/	1
92	0	Perijinan	#	1
93	92	Reklame OA	potensi_reklame/daftar_oa/	1
\.


--
-- Name: Menu_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Menu_Id_seq"', 1, false);


--
-- PostgreSQL database dump complete
--

