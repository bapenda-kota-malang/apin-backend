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
-- Data for Name: Rekening; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Rekening" ("Id", "Parent_Id", "Tipe", "Kelompok", "Jenis", "Objek", "Rincian", "Sub1", "Sub2", "Sub3", "Kode", "Nama", "Level", "KodeBaru", "KodeJenisPajak", "KodeVaJatim", "KodeBilling", "KodeJenisUsaha", "JenisUsaha") FROM stdin;
10096	\N	4								4	Pendapatan Daerah	1	4	\N	\N	\N	\N	\N
10097	10096	4	1							4.1	Pendapatan Asli Daerah	2	41	\N	\N	\N	\N	\N
10099	10097	4	1	1						4.1.1	Hasil Pajak Daerah	3	411	\N	\N	\N	\N	\N
10195	\N	1								1	Sumber Dana	1	\N	\N	\N	\N	\N	\N
10196	10195	1	1							1.1	Sumber Dana	2	\N	\N	\N	\N	\N	\N
10197	10196	1	1	1						1.1.1	Sumber dana	3	\N	\N	\N	\N	\N	\N
10198	10197	1	1	1	01					1.1.1.01	Sumber Dana	4	\N	\N	\N	\N	\N	\N
10199	10198	1	1	1	01	01				1.1.1.01.01	Kas Daerah	5	\N	\N	\N	\N	\N	\N
10200	10198	1	1	1	01	02				1.1.1.01.02	Pendapatan Asli Daerah	5	\N	\N	\N	\N	\N	\N
10201	10197	1	1	1	03					1.1.1.03	Kas di Bendahara Pengeluaran	4	\N	\N	\N	\N	\N	\N
10202	10201	1	1	1	03	32				1.1.1.03.32	Kas di Bendahara Penerimaan Dinas Pendapatan Daerah	5	\N	\N	\N	\N	\N	\N
10203	10200	1	1	1	01	02	01			1.1.1.01.02.01	Pendapatan Asli Daerah	6	\N	\N	\N	\N	\N	\N
10204	10200	1	1	1	01	02	32			1.1.1.01.02.32	Kas di Bendahara Penerimaan Dinas Pendapatan Daerah	6	\N	\N	\N	\N	\N	\N
10104	10100	4	1	1	01	04				4.1.1.01.04	Hotel Bintang Tiga	5	4110104	\N	\N	\N	211	Hotel Berbintang
10105	10100	4	1	1	01	05				4.1.1.01.05	Hotel Bintang Dua	5	4110105	\N	\N	\N	221	Hotel Melati
10106	10100	4	1	1	01	06				4.1.1.01.06	Hotel Bintang Satu	5	4110106	\N	\N	\N	\N	\N
10107	10100	4	1	1	01	07				4.1.1.01.07	Hotel Melati	5	4110107	\N	\N	\N	221	Hotel Melati
10108	10100	4	1	1	01	08				4.1.1.01.08	Hotel Melati Dua	5	4110108	\N	\N	\N	\N	\N
10109	10100	4	1	1	01	09				4.1.1.01.09	Hotel Melati Satu	5	4110109	\N	\N	\N	\N	\N
10137	10100	4	1	1	01	10				4.1.1.01.10	Motel	5	4110110	\N	\N	\N	\N	\N
10138	10100	4	1	1	01	11				4.1.1.01.11	Cottage	5	4110111	\N	\N	\N	\N	\N
10139	10100	4	1	1	01	12				4.1.1.01.12	Losmen/Rumah Penginapan/Pesanggrahan/Hostel/Rumah Kos	5	4110112	\N	\N	\N	\N	\N
10140	10100	4	1	1	01	13				4.1.1.01.13	Wisma Pariwisata	5	4110113	\N	\N	\N	221	Hotel Melati
10141	10100	4	1	1	01	14				4.1.1.01.14	Gubuk Pariwisata	5	4110114	\N	\N	\N	\N	\N
10110	10099	4	1	1	02					4.1.1.02	Pajak Restoran	4	41102	B	14605	35730020000	\N	\N
10112	10110	4	1	1	02	02				4.1.1.02.02	Rumah Makan	5	4110202	\N	\N	\N	101	Rumah Makan
10113	10110	4	1	1	02	03				4.1.1.02.03	Kafetaria/Cafe	5	4110203	\N	\N	\N	103	Cafe
10114	10110	4	1	1	02	04				4.1.1.02.04	Kantin	5	4110204	\N	\N	\N	\N	\N
10115	10110	4	1	1	02	05				4.1.1.02.05	Katering	5	4110205	\N	\N	\N	\N	\N
10116	10110	4	1	1	02	06				4.1.1.02.06	Warung Makan	5	4110206	\N	\N	\N	\N	\N
10143	10110	4	1	1	02	07				4.1.1.02.08	Bar	5	4110208	\N	\N	\N	\N	\N
10144	10110	4	1	1	02	08				4.1.1.02.09	Jasa Boga	5	4110209	\N	\N	\N	\N	\N
10117	10099	4	1	1	03					4.1.1.03	Pajak Hiburan	4	41103	c	14606	35730030000	\N	\N
10118	10117	4	1	1	03	01				4.1.1.03.01	Tontonan Film/Bioskop	5	4110301	\N	\N	\N	302	Film/VCD/Bioskop
10120	10117	4	1	1	03	03				4.1.1.03.03	Kontes Kecantikan	5	4110303	\N	\N	\N	\N	\N
10121	10117	4	1	1	03	04				4.1.1.03.04	Kontes Binaraga	5	4110304	\N	\N	\N	\N	\N
10122	10117	4	1	1	03	05				4.1.1.03.05	Pameran	5	4110305	\N	\N	\N	\N	\N
10123	10117	4	1	1	03	06				4.1.1.03.06	Diskotik	5	4110306	\N	\N	\N	309	Diskotik/Klab Malam/Bar
10145	10117	4	1	1	03	07				4.1.1.03.07	Karaoke	5	4110307	\N	\N	\N	307	Karaoke
10146	10117	4	1	1	03	08				4.1.1.03.08	Klub Malam	5	4110308	\N	\N	\N	\N	\N
10147	10117	4	1	1	03	09				4.1.1.03.09	Sirkus/Akrobat/Sulap	5	4110309	\N	\N	\N	\N	\N
10148	10117	4	1	1	03	10				4.1.1.03.10	Permainan Bilyar	5	4110310	\N	\N	\N	303	Bilyard
10149	10117	4	1	1	03	11				4.1.1.03.11	Permainan Golf	5	4110311	\N	\N	\N	\N	\N
10150	10117	4	1	1	03	12				4.1.1.03.12	Permainan Bowling	5	4110312	\N	\N	\N	308	Bowling
10151	10117	4	1	1	03	13				4.1.1.03.13	Pacuan Kuda	5	4110313	\N	\N	\N	\N	\N
10152	10117	4	1	1	03	14				4.1.1.03.14	Balap Kendaraan Motor	5	4110314	\N	\N	\N	\N	\N
10155	10117	4	1	1	03	17				4.1.1.03.17	Mandi Uap/spa	5	4110317	\N	\N	\N	\N	\N
10156	10117	4	1	1	03	18				4.1.1.03.18	Pusat Kebugaran	5	4110318	\N	\N	\N	312	Pusat Kebugaran
10157	10117	4	1	1	03	19				4.1.1.03.19	Pertandingan Olaharaga	5	4110319	\N	\N	\N	311	Pertandingan Olahraga
10124	10099	4	1	1	04					4.1.1.04	PAJAK REKLAME	4	41104	D	14607	35730040000	\N	\N
10125	10124	4	1	1	04	01				4.1.1.04.01	Reklame Papan/Billboard,Videotron,Megatron	5	4110401	\N	\N	\N	401	Pajak Reklame
10126	10124	4	1	1	04	02				4.1.1.04.02	Reklame Kain (SunScreen)	5	4110402	\N	\N	\N	401	Pajak Reklame
10127	10124	4	1	1	04	03				4.1.1.04.03	Reklame Melekat/stiker	5	4110403	\N	\N	\N	401	Pajak Reklame
10128	10124	4	1	1	04	04				4.1.1.04.04	Reklame Selebaran	5	4110404	\N	\N	\N	401	Pajak Reklame
10130	10124	4	1	1	04	06				4.1.1.04.06	Reklame Udara	5	4110406	\N	\N	\N	401	Pajak Reklame
10158	10124	4	1	1	04	07				4.1.1.04.07	Reklame Apung	5	4110407	\N	\N	\N	401	Pajak Reklame
10159	10124	4	1	1	04	08				4.1.1.04.08	Reklame Suara	5	4110408	\N	\N	\N	401	Pajak Reklame
10160	10124	4	1	1	04	09				4.1.1.04.09	Reklame Film/Slide	5	4110409	\N	\N	\N	401	Pajak Reklame
10161	10124	4	1	1	04	10				4.1.1.04.10	Reklame Peragaan	5	4110410	\N	\N	\N	401	Pajak Reklame
10132	10131	4	1	1	05	01				4.1.1.05.01	Pajak Penerangan Jalan PLN	5	4110501	\N	\N	\N	600	PPJ. PLN
10205	10131	4	1	1	05	02				4.1.1.05.02	Pajak Penerangan Jalan Non PLN/Jalan Dihasilkan Sendiri	5	4110502	\N	\N	\N	600	PPJ. Non PLN
10188	10099	4	1	1	07					4.1.1.07	Pajak Parkir	4	41107	F	14609	35730060000	\N	\N
10189	10188	4	1	1	07	01				4.1.1.07.01	Pajak Parkir	5	4110701	\N	\N	\N	400	Pajak Parkir
10190	10099	4	1	1	08					4.1.1.08	Pajak Air Tanah	4	41108	G	14610	35730050000	\N	\N
10191	10190	4	1	1	08	01				4.1.1.08.01	Pajak Air Tanah	5	4110801	\N	\N	\N	700	Pajak Air Tanah
10101	10100	4	1	1	01	01				4.1.1.01.01	Hotel Bintang Lima Berlian	5	4110101	\N	\N	\N	\N	\N
10100	10099	4	1	1	01					4.1.1.01	Pajak Hotel	4	41101	A	14604	35730010000	\N	\N
10102	10100	4	1	1	01	02				4.1.1.01.02	Hotel Bintang Lima	5	4110102	\N	\N	\N	211	Hotel Berbintang
10103	10100	4	1	1	01	03				4.1.1.01.03	Hotel Bintang Empat	5	4110103	\N	\N	\N	211	Hotel Berbintang
10207	10100	4	1	1	01	15				4.1.1.01.15	Guest House	5	4110115	\N	\N	\N	221	Hotel Melati
10208	10100	4	1	1	01	16				4.1.1.01.16	Rumah Kos	5	4110116	\N	\N	\N	221	Hotel Melati
10111	10110	4	1	1	02	01				4.1.1.02.01	Restoran	5	4110201	\N	\N	\N	102	Restoran
10209	10112	4	1	1	02	02	01			4.1.1.02.02.01	Catering	6	4110202	\N	\N	\N	101	Rumah Makan
10119	10117	4	1	1	03	02				4.1.1.03.02	Pagelaran Kesenian/Musik/Tari/Busana	5	4110302	\N	\N	\N	310	Pagelaran Kesenian/Musik/Tari/Busana
10210	10145	4	1	1	03	07	01			4.1.1.03.07.01	KARAOKE KELUARGA	6	4110307	\N	\N	\N	307	Karaoke
10211	10145	4	1	1	03	07	02			4.1.1.03.07.02	KARAOKE NON KELUARGA	6	4110307	\N	\N	\N	307	Karaoke
10153	10117	4	1	1	03	15				4.1.1.03.15	Permainan Ketangkasan/Time Zone/Mandi Bola,dan Sejenisnya	5	4110315	\N	\N	\N	305	Permainan Ketangkasan
10154	10117	4	1	1	03	16				4.1.1.03.16	Panti Pijat/Refleksi	5	4110316	\N	\N	\N	306	Panti Pijat/Refleksi/Spa
10206	10117	4	1	1	03	20				4.1.1.03.20	Taman Rekreasi	5	4110320	\N	\N	\N	304	Taman Rekreasi
10212	10117	4	1	1	03	21				4.1.1.03.21	Karaoke Non	5	4110321	\N	\N	\N	307	Karaoke
10213	10117	4	1	1	03	22				4.1.1.03.22	Rental VCD	5	4110322	\N	\N	\N	302	Film/VCD/Bioskop
10129	10124	4	1	1	04	05				4.1.1.04.05	Reklame Berjalan	5	4110405	\N	\N	\N	401	Pajak Reklame
10131	10099	4	1	1	05					4.1.1.05	PAJAK PENERANGAN JALAN	4	41105	E	14608	35730090000	\N	\N
10214	10124	4	1	1	04	11				4.1.1.04.11	Reklame Flaghtchain	5	4110411	\N	\N	\N	401	Pajak Reklame
\.


--
-- Name: Rekening_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Rekening_Id_seq"', 1, false);


--
-- PostgreSQL database dump complete
--

