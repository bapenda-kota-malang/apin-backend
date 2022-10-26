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
-- Data for Name: Npwpd; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."Npwpd" ("Id", "ObjekPajak_Id", "Golongan", "Nomor", "Npwp", "TanggalPengukuhan", "TanggalNpwpd", "Npwpd", "Status", "TanggalTutup", "TanggalBuka", "JenisPajak", "Skpd_Id", "Rekening_Id", "Nama", "User_Id", "OmsetOp", "Genset", "AirTanah", "TanggalMulaiUsaha", "LuasBangunan", "JamBukaUsaha", "JamTutupUsaha", "Pengunjung", "FotoKtp", "SuratIzinUsaha", "LainLain", "FotoObjek", "JalurRegistrasi", "VerifiedAt", "VendorEtax_Id", "CreatedAt", "UpdatedAt", "DeletedAt") FROM stdin;
2902	\N	1	1308	\N	\N	2016-12-01 00:00:00+07	1308..62.221	2	2019-04-01 00:00:00+07	\N	SA	\N	10107	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2991	\N	1	1598	\N	\N	2017-03-27 00:00:00+07	1598.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4244	\N	1	755	\N	\N	2014-12-01 00:00:00+07	0755..65.221	2	2017-01-30 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4218	\N	1	1741	\N	\N	2016-09-01 00:00:00+07	1741..62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1107	\N	2	1131	\N	\N	2015-03-05 00:00:00+07	1131..65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
909	\N	2	1704	\N	\N	2013-12-19 00:00:00+07	1704..62.221	2	2017-05-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3398	\N	1	528	\N	\N	2017-03-01 00:00:00+07	'0528.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2767	\N	2	338	\N	\N	2016-01-28 00:00:00+07	0338..62.102	1	\N	2020-06-01 00:00:00+07	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2168	\N	2	1271	\N	\N	2011-02-11 00:00:00+07	1271...62.102	2	2017-08-28 00:00:00+07	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2734	\N	2	675	\N	\N	2015-11-09 00:00:00+07	0675..61.102	2	2017-10-18 00:00:00+07	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1736	\N	2	1271	\N	\N	2011-02-07 00:00:00+07	1271.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1698	\N	2	744	\N	\N	2011-02-08 00:00:00+07	0744..62.101	2	2017-09-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1850	\N	2	193	\N	\N	2011-02-07 00:00:00+07	0193..62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1955	\N	2	241	\N	\N	2011-02-08 00:00:00+07	0241..62.101	2	2018-02-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1972	\N	2	186	\N	\N	2011-02-08 00:00:00+07	0186..62.101	2	2020-01-08 00:00:00+07	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2073	\N	2	190	\N	\N	2011-02-09 00:00:00+07	0190..64.101	2	2019-09-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2092	\N	2	56	\N	\N	2011-02-09 00:00:00+07	0056..61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2102	\N	2	742	\N	\N	2011-02-09 00:00:00+07	0742..62.101	2	2017-09-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2104	\N	2	794	\N	\N	2011-02-09 00:00:00+07	0794..62.101	2	2017-09-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2173	\N	2	516	\N	\N	2011-02-11 00:00:00+07	0516..65.101	2	2019-01-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2190	\N	2	1377	\N	\N	2011-02-22 00:00:00+07	1377..62.101	2	2017-04-04 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2199	\N	2	327	\N	\N	2011-03-23 00:00:00+07	0327..61.101	1	2017-10-23 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2386	\N	2	1346	\N	\N	2013-05-01 00:00:00+07	1346..62.101	2	2017-08-14 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2392	\N	2	1690	\N	\N	2013-08-13 00:00:00+07	1690..62.101	2	2017-10-27 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2425	\N	2	240	\N	\N	2014-02-28 00:00:00+07	0240..62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2757	\N	2	749	\N	\N	2016-01-08 00:00:00+07	0749..61.101	2	2018-07-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2793	\N	2	833	\N	\N	2016-03-04 00:00:00+07	0833..62.101	2	2019-11-05 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3574	\N	1	1599	\N	\N	2017-03-01 00:00:00+07	1599..62.101	2	2017-08-15 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3353	\N	1	602	\N	\N	2016-04-04 00:00:00+07	0602..65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3363	\N	1	483	\N	\N	2017-01-02 00:00:00+07	0483..64.101	2	2017-03-27 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3400	\N	1	1103	\N	\N	2016-12-06 00:00:00+07	1103..65.101	2	\N	\N	SA	\N	10112	solusiti	\N	\N	f	f	2021-01-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3405	\N	1	926	\N	\N	2017-02-01 00:00:00+07	0926..65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3411	\N	1	1994	\N	\N	2017-02-01 00:00:00+07	1994..61.101	2	2019-03-28 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3436	\N	1	1608	\N	\N	2017-04-12 00:00:00+07	1608.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3565	\N	1	1596	\N	\N	2017-03-01 00:00:00+07	1596..65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3585	\N	1	519	\N	\N	2017-03-01 00:00:00+07	0519.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4085	\N	1	1021	\N	\N	2007-09-01 00:00:00+07	1021..62.101	2	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4124	\N	1	1117	\N	\N	2007-09-01 00:00:00+07	1117..62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4127	\N	1	96	\N	\N	2015-12-01 00:00:00+07	0096..101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4129	\N	1	269	\N	\N	2008-10-01 00:00:00+07	0269..62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4236	\N	1	1023	\N	\N	2011-06-01 00:00:00+07	1023..62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4249	\N	1	1286	\N	\N	2009-12-01 00:00:00+07	1286..62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4252	\N	1	129	\N	\N	2013-10-01 00:00:00+07	0129..62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4294	\N	1	487	\N	\N	2009-03-02 00:00:00+07	0487..65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4295	\N	1	488	\N	\N	2009-02-02 00:00:00+07	0488..65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4300	\N	1	505	\N	\N	2009-10-01 00:00:00+07	0505..65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4301	\N	1	498	\N	\N	2012-12-03 00:00:00+07	0498..65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4303	\N	1	491	\N	\N	2012-07-02 00:00:00+07	0491..65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4306	\N	1	156	\N	\N	2012-05-01 00:00:00+07	0156..65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4309	\N	1	32	\N	\N	2004-06-01 00:00:00+07	0032..65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4312	\N	1	190	\N	\N	2009-06-01 00:00:00+07	0190..64.101	2	2018-07-18 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2803	\N	2	1437	\N	\N	2016-03-21 00:00:00+07	1437..65.103	2	2018-02-28 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6876	\N	1	258265	\N	\N	2021-11-04 00:00:00+07	2582.65.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-04 00:00:00+07	\N	10:00	21:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2364	\N	2	234	\N	\N	2012-12-20 00:00:00+07	0234..63.103	2	2018-03-01 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3993	\N	1	176	\N	\N	2009-03-01 00:00:00+07	0176..65.303	2	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4000	\N	1	413	\N	\N	1998-01-01 00:00:00+07	0413..62.303	2	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3992	\N	1	176	\N	\N	2009-03-01 00:00:00+07	0176..65.303	2	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3988	\N	1	45	\N	\N	2015-09-12 00:00:00+07	0045..64.303	1	2017-09-13 00:00:00+07	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4038	\N	1	153	\N	\N	2007-03-01 00:00:00+07	0153..64.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3563	\N	1	1649	\N	\N	2017-05-09 00:00:00+07	1649.62.305	2	2017-05-04 00:00:00+07	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3588	\N	1	87	\N	\N	2017-05-10 00:00:00+07	0087..61.306	2	2017-10-09 00:00:00+07	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3754	\N	1	1	\N	\N	2017-05-31 00:00:00+07	0001.00.600	1	\N	\N	OA	\N	10132	Amalia Mahardani	\N	\N	f	f	2021-12-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1639	\N	2	468	\N	\N	2010-02-25 00:00:00+07	0468..62.600	2	\N	\N	SA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3981	\N	2	1027	\N	\N	2014-06-01 00:00:00+07	1027..62.600	2	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1440	\N	2	488	\N	\N	2011-02-17 00:00:00+07	0488..61.400	2	2018-02-26 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
464	\N	2	341	\N	\N	2011-02-23 00:00:00+07	0341....62	2	2020-07-01 00:00:00+07	\N	SA	\N	10211	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3889	\N	1	363	\N	\N	2017-07-17 00:00:00+07	0363..63.700	1	2018-06-28 00:00:00+07	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3562	\N	1	323	\N	\N	2017-05-09 00:00:00+07	0323..64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3297	\N	1	1026	\N	\N	2016-10-03 00:00:00+07	1026.65.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
412	\N	2	639	\N	\N	2015-01-13 00:00:00+07	0639..65.700	1	\N	2020-06-01 00:00:00+07	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
409	\N	2	684	\N	\N	2014-03-04 00:00:00+07	0684..61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3471	\N	1	10	\N	\N	2017-03-01 00:00:00+07	0010.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5104	\N	1	2029	\N	\N	2019-03-15 00:00:00+07	2029.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
907	\N	2	1375	\N	\N	2011-02-17 00:00:00+07	1375.62.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1307	\N	2	1349	\N	\N	2015-10-09 00:00:00+07	1349.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1308	\N	2	1351	\N	\N	2015-10-09 00:00:00+07	1351.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6895	\N	1	1093	\N	\N	2021-11-10 00:00:00+07	1093.61.211	1	\N	\N	SA	\N	10102	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
574	\N	2	341	\N	\N	2011-02-21 00:00:00+07	0341..62.211	1	\N	\N	SA	\N	10104	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
501	\N	2	337	\N	\N	2011-02-17 00:00:00+07	337.62.211	1	\N	2020-06-01 00:00:00+07	SA	\N	10102	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1344	\N	2	463	\N	\N	2016-01-07 00:00:00+07	0463.61.211	1	\N	\N	SA	\N	10103	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
503	\N	2	90	\N	\N	2011-02-17 00:00:00+07	0090.65.211	1	\N	\N	SA	\N	10103	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
576	\N	2	1625	\N	\N	2012-02-14 00:00:00+07	1625.62.211	1	\N	\N	SA	\N	10103	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
590	\N	2	1665	\N	\N	2013-02-27 00:00:00+07	1665.62.211	1	\N	\N	SA	\N	10103	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
593	\N	2	674	\N	\N	2013-04-16 00:00:00+07	0674.61.211	1	\N	\N	SA	\N	10103	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
596	\N	2	1687	\N	\N	2013-10-01 00:00:00+07	1687.62.211	1	\N	\N	SA	\N	10103	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
597	\N	2	684	\N	\N	2013-11-12 00:00:00+07	0684.61.211	1	\N	\N	SA	\N	10103	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
605	\N	2	1706	\N	\N	2013-12-24 00:00:00+07	1706.62.211	1	\N	\N	SA	\N	10103	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
502	\N	2	89	\N	\N	2011-02-17 00:00:00+07	0089.65.211	1	\N	\N	SA	\N	10103	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1346	\N	2	1952	\N	\N	2016-01-08 00:00:00+07	1952.62.211	1	\N	\N	SA	\N	10103	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
767	\N	2	699	\N	\N	2014-08-26 00:00:00+07	0699.61.211	1	\N	\N	SA	\N	10103	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
610	\N	2	685	\N	\N	2014-01-13 00:00:00+07	0685.61.211	1	\N	\N	SA	\N	10104	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
828	\N	2	1763	\N	\N	2014-10-10 00:00:00+07	1763.62.211	1	\N	\N	SA	\N	10104	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4210	\N	1	2210	\N	\N	2017-10-23 00:00:00+07	2210.62.211	1	\N	\N	SA	\N	10104	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
512	\N	2	339	\N	\N	2011-02-17 00:00:00+07	0339.62.211	1	\N	\N	SA	\N	10104	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5355	\N	1	2126	\N	\N	2019-05-27 00:00:00+07	2126.65.211	1	\N	\N	SA	\N	10104	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
511	\N	2	338	\N	\N	2011-02-17 00:00:00+07	0338.62.211	1	\N	\N	SA	\N	10104	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
510	\N	2	454	\N	\N	2011-02-17 00:00:00+07	0454.62.211	1	\N	2020-05-01 00:00:00+07	SA	\N	10104	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4506	\N	1	658	\N	\N	2018-04-13 00:00:00+07	0658.64.211	1	\N	\N	SA	\N	10104	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
509	\N	2	383	\N	\N	2011-02-17 00:00:00+07	0383.62.211	1	\N	2020-05-01 00:00:00+07	SA	\N	10104	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
508	\N	2	539	\N	\N	2011-02-17 00:00:00+07	0539.62.211	1	\N	\N	SA	\N	10104	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
507	\N	2	538	\N	\N	2011-02-17 00:00:00+07	0538.62.211	1	\N	\N	SA	\N	10104	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
506	\N	2	341	\N	\N	2011-02-17 00:00:00+07	0341.62.211	1	\N	\N	SA	\N	10104	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
505	\N	2	282	\N	\N	2011-02-17 00:00:00+07	0282.62.211	1	\N	\N	SA	\N	10104	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1384	\N	2	1965	\N	\N	2016-04-08 00:00:00+07	1965.62.211	1	\N	\N	SA	\N	10104	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
504	\N	2	281	\N	\N	2011-02-17 00:00:00+07	0281.62.211	1	\N	\N	SA	\N	10104	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
514	\N	2	334	\N	\N	2011-02-17 00:00:00+07	0334.62.221	1	\N	\N	SA	\N	10105	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5114	\N	1	2033	\N	\N	2019-03-18 00:00:00+07	2033.65.221	1	\N	\N	SA	\N	10105	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1303	\N	2	1918	\N	\N	2015-10-07 00:00:00+07	1918.62.221	1	\N	\N	SA	\N	10105	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6447	\N	1	2492	\N	\N	2021-04-20 00:00:00+07	2492.65.221	1	\N	\N	SA	\N	10105	Amalia Mahardani	\N	\N	f	f	2021-04-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
516	\N	2	370	\N	\N	2011-02-17 00:00:00+07	0370.62.221	1	\N	\N	SA	\N	10105	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
513	\N	2	333	\N	\N	2011-02-17 00:00:00+07	0333.62.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10105	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
603	\N	2	661	\N	\N	2013-12-13 00:00:00+07	0661.65.221	1	\N	\N	SA	\N	10105	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
585	\N	2	639	\N	\N	2013-01-08 00:00:00+07	0639.65.221	1	\N	\N	SA	\N	10105	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
543	\N	2	76	\N	\N	2011-02-17 00:00:00+07	0076.61.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
538	\N	2	349	\N	\N	2011-02-17 00:00:00+07	0349.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
522	\N	2	375	\N	\N	2011-02-17 00:00:00+07	0375.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
592	\N	2	1668	\N	\N	2013-03-20 00:00:00+07	1668.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
521	\N	2	361	\N	\N	2011-02-17 00:00:00+07	0361.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1331	\N	2	1930	\N	\N	2015-11-10 00:00:00+07	1930.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
527	\N	2	357	\N	\N	2011-02-17 00:00:00+07	0357.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
545	\N	2	63	\N	\N	2011-02-17 00:00:00+07	0063.65.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
540	\N	2	382	\N	\N	2011-02-17 00:00:00+07	0382.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
546	\N	2	1201	\N	\N	2011-02-17 00:00:00+07	1201.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
529	\N	2	72	\N	\N	2011-02-17 00:00:00+07	0072.61.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1354	\N	2	1957	\N	\N	2016-02-10 00:00:00+07	1957.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6779	\N	1	2739	\N	\N	2021-09-23 00:00:00+07	2739.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2021-09-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
539	\N	2	354	\N	\N	2011-02-17 00:00:00+07	0354.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
550	\N	2	350	\N	\N	2011-02-17 00:00:00+07	0350.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
594	\N	2	384	\N	\N	2011-02-17 00:00:00+07	0384.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
551	\N	2	64	\N	\N	2011-02-17 00:00:00+07	0064.65.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
549	\N	2	73	\N	\N	2011-02-17 00:00:00+07	0073.61.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
517	\N	2	381	\N	\N	2011-02-17 00:00:00+07	0381.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
530	\N	2	366	\N	\N	2011-02-17 00:00:00+07	0366.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5899	\N	1	820	\N	\N	2019-11-11 00:00:00+07	0820.64.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
524	\N	2	359	\N	\N	2011-02-17 00:00:00+07	0359.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5903	\N	1	2559	\N	\N	2019-11-12 00:00:00+07	2559.62.221	1	\N	\N	SA	\N	10107	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
553	\N	2	80	\N	\N	2011-02-17 00:00:00+07	0080.61.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
526	\N	2	1199	\N	\N	2011-02-17 00:00:00+07	1199.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
556	\N	2	77	\N	\N	2011-02-17 00:00:00+07	0077.61.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
589	\N	2	1663	\N	\N	2013-02-13 00:00:00+07	1663.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
591	\N	2	1667	\N	\N	2013-03-05 00:00:00+07	1667.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6433	\N	1	2687	\N	\N	2021-04-14 00:00:00+07	2687.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2021-04-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
536	\N	2	74	\N	\N	2011-02-17 00:00:00+07	0074.61.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
528	\N	2	36	\N	\N	2011-02-17 00:00:00+07	0036.64.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
544	\N	2	82	\N	\N	2011-02-17 00:00:00+07	0082.61.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2875	\N	1	26	\N	\N	2011-02-17 00:00:00+07	0026.63.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
560	\N	2	368	\N	\N	2011-02-17 00:00:00+07	0368.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
523	\N	2	78	\N	\N	2011-02-17 00:00:00+07	0078.61.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
542	\N	2	343	\N	\N	2011-02-17 00:00:00+07	0343.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
541	\N	2	367	\N	\N	2011-02-17 00:00:00+07	0367.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
525	\N	2	334	\N	\N	2011-02-17 00:00:00+07	0334.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
573	\N	2	345	\N	\N	2011-02-18 00:00:00+07	0345.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
825	\N	2	703	\N	\N	2014-10-09 00:00:00+07	0703.61.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
533	\N	2	358	\N	\N	2011-02-17 00:00:00+07	0358.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
555	\N	2	61	\N	\N	2011-02-17 00:00:00+07	0061.65.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
570	\N	2	376	\N	\N	2011-02-17 00:00:00+07	0376.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
578	\N	2	1639	\N	\N	2012-04-23 00:00:00+07	1639.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
579	\N	2	621	\N	\N	2012-04-25 00:00:00+07	0621.65.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1352	\N	2	1949	\N	\N	2016-02-09 00:00:00+07	1949.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1430	\N	2	2004	\N	\N	2016-09-09 00:00:00+07	2004.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
535	\N	2	374	\N	\N	2011-02-17 00:00:00+07	0374.62.221	1	\N	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1418	\N	2	780	\N	\N	2016-07-29 00:00:00+07	0780.61.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
559	\N	2	365	\N	\N	2011-02-17 00:00:00+07	0365.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
520	\N	2	75	\N	\N	2011-02-17 00:00:00+07	0075.61.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
532	\N	2	353	\N	\N	2011-02-17 00:00:00+07	0353.62.221	2	2018-05-01 00:00:00+07	\N	SA	\N	10107	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
656	\N	2	713	\N	\N	2014-02-28 00:00:00+07	0713.65.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
534	\N	2	372	\N	\N	2011-02-17 00:00:00+07	0372.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
837	\N	2	1679	\N	\N	2014-10-16 00:00:00+07	1679.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
531	\N	2	363	\N	\N	2011-02-17 00:00:00+07	0363.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
572	\N	2	377	\N	\N	2011-02-18 00:00:00+07	0377.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
519	\N	2	62	\N	\N	2011-02-17 00:00:00+07	0062.65.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
537	\N	2	356	\N	\N	2011-02-17 00:00:00+07	0356.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6426	\N	1	606	\N	\N	2021-04-12 00:00:00+07	0606.63.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2021-04-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
518	\N	2	364	\N	\N	2011-02-17 00:00:00+07	0364.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
552	\N	2	346	\N	\N	2011-02-17 00:00:00+07	0346.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
584	\N	2	369	\N	\N	2012-12-18 00:00:00+07	0369.64.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
548	\N	2	1357	\N	\N	2011-02-17 00:00:00+07	1357.62.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2021-07-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
547	\N	2	360	\N	\N	2011-02-17 00:00:00+07	0360.62.221	1	\N	\N	SA	\N	10107	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1429	\N	2	1490	\N	\N	2016-09-09 00:00:00+07	1490.65.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
564	\N	2	733	\N	\N	2011-02-17 00:00:00+07	0733.62.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
571	\N	2	241	\N	\N	2011-02-17 00:00:00+07	0241.61.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
563	\N	2	378	\N	\N	2011-02-17 00:00:00+07	0378.62.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1267	\N	2	1305	\N	\N	2015-09-01 00:00:00+07	1305.65.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
568	\N	2	380	\N	\N	2011-02-17 00:00:00+07	0380.62.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
561	\N	2	369	\N	\N	2011-02-17 00:00:00+07	0369.62.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1046	\N	2	1817	\N	\N	2015-02-10 00:00:00+07	1817.62.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
566	\N	2	379	\N	\N	2011-02-17 00:00:00+07	0379.62.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3408	\N	1	2010	\N	\N	2017-01-02 00:00:00+07	2010.62.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
581	\N	2	635	\N	\N	2012-10-10 00:00:00+07	0635.65.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
575	\N	2	570	\N	\N	2011-03-08 00:00:00+07	0570.65.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1015	\N	2	1050	\N	\N	2015-02-02 00:00:00+07	1050.65.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
677	\N	2	1727	\N	\N	2014-03-14 00:00:00+07	1727.62.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1431	\N	2	306	\N	\N	2016-09-09 00:00:00+07	0306.63.221	1	\N	2020-05-01 00:00:00+07	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1385	\N	2	1972	\N	\N	2016-04-08 00:00:00+07	1972.62.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1340	\N	2	1395	\N	\N	2015-12-08 00:00:00+07	1395.65.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3440	\N	1	2114	\N	\N	2017-03-01 00:00:00+07	2114.62.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1345	\N	2	1938	\N	\N	2016-01-08 00:00:00+07	1938.62.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1020	\N	2	93	\N	\N	2015-02-04 00:00:00+07	0093.65.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
562	\N	2	81	\N	\N	2011-02-17 00:00:00+07	0081.61.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5569	\N	1	2466	\N	\N	2019-08-02 00:00:00+07	2466.62.221	1	\N	\N	SA	\N	10140	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
567	\N	2	352	\N	\N	2011-02-17 00:00:00+07	0352.62.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4367	\N	1	2225	\N	\N	2018-01-29 00:00:00+07	2225.62.221	1	\N	\N	SA	\N	10140	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5441	\N	1	954	\N	\N	2019-06-26 00:00:00+07	0954.61.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5924	\N	1	2347	\N	\N	2019-11-20 00:00:00+07	2347.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5930	\N	1	2348	\N	\N	2019-11-22 00:00:00+07	2348.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5931	\N	1	2570	\N	\N	2019-11-22 00:00:00+07	2570.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3368	\N	1	1536	\N	\N	2017-01-02 00:00:00+07	1536.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1049	\N	2	1079	\N	\N	2015-02-11 00:00:00+07	1079.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4413	\N	1	631	\N	\N	2018-02-28 00:00:00+07	0631.64.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1411	\N	2	1985	\N	\N	2016-05-10 00:00:00+07	1985.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1412	\N	2	1993	\N	\N	2016-06-13 00:00:00+07	1993.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5935	\N	1	2349	\N	\N	2019-11-27 00:00:00+07	2349.65.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5964	\N	1	2354	\N	\N	2019-12-16 00:00:00+07	2354.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5963	\N	1	1006	\N	\N	2019-12-16 00:00:00+07	1006.61.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5939	\N	1	2350	\N	\N	2019-12-02 00:00:00+07	2350.65.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4532	\N	1	877	\N	\N	2018-04-25 00:00:00+07	0877.61.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3366	\N	1	2068	\N	\N	2017-03-01 00:00:00+07	2068.62.221	1	\N	2020-05-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6239	\N	1	586	\N	\N	2020-10-14 00:00:00+07	0586.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2020-10-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5020	\N	1	1979	\N	\N	2019-02-22 00:00:00+07	1979.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6237	\N	1	584	\N	\N	2020-10-13 00:00:00+07	0584.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2020-10-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3978	\N	1	2187	\N	\N	2017-09-07 00:00:00+07	2187.62.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5294	\N	1	2105	\N	\N	2019-05-15 00:00:00+07	2105.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3975	\N	1	1652	\N	\N	2017-09-07 00:00:00+07	1652.65.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
655	\N	2	1716	\N	\N	2014-02-27 00:00:00+07	1716.62.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5022	\N	1	2338	\N	\N	2019-02-25 00:00:00+07	2338.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5011	\N	1	1970	\N	\N	2019-02-21 00:00:00+07	1970.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5767	\N	1	2311	\N	\N	2019-10-04 00:00:00+07	2311.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4934	\N	1	2320	\N	\N	2019-01-24 00:00:00+07	2320.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4933	\N	1	898	\N	\N	2019-01-24 00:00:00+07	0898.61.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4932	\N	1	2319	\N	\N	2019-01-24 00:00:00+07	2319.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
580	\N	2	1648	\N	\N	2012-06-29 00:00:00+07	1648.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
582	\N	2	1655	\N	\N	2012-10-25 00:00:00+07	1655.62.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3657	\N	1	1609	\N	\N	2017-05-17 00:00:00+07	1609.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
583	\N	2	1654	\N	\N	2012-11-14 00:00:00+07	1654.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4918	\N	1	1926	\N	\N	2018-12-28 00:00:00+07	1926.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1108	\N	2	249	\N	\N	2015-03-05 00:00:00+07	0249.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
663	\N	2	1714	\N	\N	2014-03-05 00:00:00+07	1714.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3707	\N	1	2125	\N	\N	2017-04-27 00:00:00+07	2125.62.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5642	\N	1	2242	\N	\N	2019-08-22 00:00:00+07	2242.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5641	\N	1	2241	\N	\N	2019-08-22 00:00:00+07	2241.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-08-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6874	\N	1	928	\N	\N	2021-11-04 00:00:00+07	0928.64.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-11-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5631	\N	1	2238	\N	\N	2019-08-15 00:00:00+07	2238.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
586	\N	2	1660	\N	\N	2013-01-21 00:00:00+07	1660.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5776	\N	1	2314	\N	\N	2019-10-04 00:00:00+07	2314.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5625	\N	1	2234	\N	\N	2019-08-14 00:00:00+07	2234.65.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
587	\N	2	1661	\N	\N	2013-01-21 00:00:00+07	1661.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5777	\N	1	2315	\N	\N	2019-10-04 00:00:00+07	2315.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1080	\N	2	1820	\N	\N	2015-02-26 00:00:00+07	1820.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6246	\N	1	2632	\N	\N	2020-10-14 00:00:00+07	2632.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2020-10-14 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4084	\N	1	2200	\N	\N	2017-09-26 00:00:00+07	2200.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
588	\N	2	1657	\N	\N	2013-01-30 00:00:00+07	1657.62.221	2	2020-03-01 00:00:00+07	\N	SA	\N	10207	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6278	\N	1	2450	\N	\N	2020-11-25 00:00:00+07	2450.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2020-11-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6279	\N	1	592	\N	\N	2020-11-26 00:00:00+07	0592.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2020-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5227	\N	1	2084	\N	\N	2019-04-18 00:00:00+07	2084.65.221	1	\N	2020-03-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5743	\N	1	2301	\N	\N	2019-09-25 00:00:00+07	2301.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5744	\N	1	2302	\N	\N	2019-09-25 00:00:00+07	2302.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2898	\N	1	1542	\N	\N	2017-04-03 00:00:00+07	1542.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2899	\N	1	1984	\N	\N	2017-03-10 00:00:00+07	1984.62.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1355	\N	2	1414	\N	\N	2016-02-11 00:00:00+07	1414.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3977	\N	1	573	\N	\N	2017-09-07 00:00:00+07	0573.64.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4988	\N	1	460	\N	\N	2019-02-15 00:00:00+07	0460.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
598	\N	2	372	\N	\N	2013-11-13 00:00:00+07	0372.64.221	2	2016-11-01 00:00:00+07	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1136	\N	2	1827	\N	\N	2015-03-23 00:00:00+07	1827.62.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4682	\N	1	2280	\N	\N	2018-08-30 00:00:00+07	2280.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4986	\N	1	2332	\N	\N	2019-02-15 00:00:00+07	2332.62.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4683	\N	1	2281	\N	\N	2018-08-30 00:00:00+07	2281.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6824	\N	1	2565	\N	\N	2021-10-18 00:00:00+07	2565.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5770	\N	1	2501	\N	\N	2019-10-04 00:00:00+07	2501.62.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4985	\N	1	1952	\N	\N	2019-02-14 00:00:00+07	1952.65.221	2	2019-02-15 00:00:00+07	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1118	\N	2	722	\N	\N	2015-03-11 00:00:00+07	0722.61.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4726	\N	1	709	\N	\N	2018-09-13 00:00:00+07	0709.64.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7074	\N	1	2798	\N	\N	2022-01-11 00:00:00+07	2798.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5680	\N	1	2259	\N	\N	2019-09-12 00:00:00+07	2259.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2963	\N	1	398	\N	\N	2017-04-03 00:00:00+07	0398.64.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3972	\N	1	367	\N	\N	2017-09-04 00:00:00+07	0367.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5232	\N	1	763	\N	\N	2019-04-23 00:00:00+07	0763.64.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6432	\N	1	2488	\N	\N	2021-04-14 00:00:00+07	2488.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-04-14 00:00:00+07	\N	00:00	23:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5749	\N	1	524	\N	\N	2019-09-26 00:00:00+07	0524.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5687	\N	1	2265	\N	\N	2019-09-17 00:00:00+07	2265.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
686	\N	2	729	\N	\N	2014-03-19 00:00:00+07	0729.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4293	\N	1	1675	\N	\N	2017-12-07 00:00:00+07	1675.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
619	\N	2	1712	\N	\N	2014-02-07 00:00:00+07	1712.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4392	\N	1	868	\N	\N	2018-02-06 00:00:00+07	0868.61.221	1	\N	2020-01-01 00:00:00+07	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5181	\N	1	921	\N	\N	2019-04-04 00:00:00+07	0921.61.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7378	\N	1	2703	\N	\N	2022-04-07 00:00:00+07	2703.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5771	\N	1	527	\N	\N	2019-10-04 00:00:00+07	0527.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6521	\N	1	2517	\N	\N	2021-06-30 00:00:00+07	2517.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-06-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5688	\N	1	2266	\N	\N	2019-09-17 00:00:00+07	2266.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1255	\N	2	1290	\N	\N	2015-08-05 00:00:00+07	1290.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
577	\N	2	1631	\N	\N	2012-02-27 00:00:00+07	1631.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
955	\N	2	1800	\N	\N	2014-12-04 00:00:00+07	1800.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3441	\N	1	2074	\N	\N	2017-03-01 00:00:00+07	2074.62.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4275	\N	1	2211	\N	\N	2017-11-20 00:00:00+07	2211.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
742	\N	2	379	\N	\N	2014-06-17 00:00:00+07	0379.64.221	1	\N	2020-08-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4270	\N	1	379	\N	\N	2017-11-09 00:00:00+07	0379.63.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6666	\N	1	622	\N	\N	2021-08-18 00:00:00+07	0622.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-08-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3703	\N	1	2119	\N	\N	2017-04-13 00:00:00+07	2119.62.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6112	\N	1	2401	\N	\N	2020-07-01 00:00:00+07	2401.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
978	\N	2	720	\N	\N	2015-01-16 00:00:00+07	0720.61.221	2	2019-01-01 00:00:00+07	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7468	\N	1	1027	\N	\N	2022-06-15 00:00:00+07	1027.64.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3432	\N	1	2099	\N	\N	2017-03-01 00:00:00+07	2099.62.221	2	2017-08-01 00:00:00+07	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5671	\N	1	2483	\N	\N	2019-09-10 00:00:00+07	2483.62.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5696	\N	1	2269	\N	\N	2019-09-18 00:00:00+07	2269.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5780	\N	1	2318	\N	\N	2019-10-07 00:00:00+07	2318.65.221	1	\N	2020-05-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3545	\N	1	2086	\N	\N	2017-04-01 00:00:00+07	2086.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6467	\N	1	2698	\N	\N	2021-05-06 00:00:00+07	2698.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-05-06 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
565	\N	2	1316	\N	\N	2011-02-17 00:00:00+07	1316.62.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5615	\N	1	2228	\N	\N	2019-08-12 00:00:00+07	2228.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6682	\N	1	626	\N	\N	2021-08-25 00:00:00+07	0626.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-08-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4159	\N	1	1664	\N	\N	2017-10-10 00:00:00+07	1664.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5210	\N	1	2373	\N	\N	2019-04-11 00:00:00+07	2373.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5517	\N	1	2203	\N	\N	2019-07-18 00:00:00+07	2203.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5819	\N	1	2526	\N	\N	2019-10-16 00:00:00+07	2526.62.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5769	\N	1	2500	\N	\N	2019-10-04 00:00:00+07	2500.62.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6487	\N	1	611	\N	\N	2021-06-03 00:00:00+07	0611.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-06-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5821	\N	1	2527	\N	\N	2019-10-16 00:00:00+07	2527.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5772	\N	1	528	\N	\N	2019-10-04 00:00:00+07	0528.63.221	1	\N	2020-05-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7463	\N	1	2903	\N	\N	2022-06-14 00:00:00+07	2903.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5835	\N	1	533	\N	\N	2019-10-17 00:00:00+07	0533.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5836	\N	1	534	\N	\N	2019-10-17 00:00:00+07	0534.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6039	\N	1	2382	\N	\N	2020-02-24 00:00:00+07	2382.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7321	\N	1	693	\N	\N	2022-03-15 00:00:00+07	0693.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-03-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5773	\N	1	529	\N	\N	2019-10-04 00:00:00+07	0529.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4956	\N	1	1941	\N	\N	2019-02-08 00:00:00+07	1941.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5699	\N	1	519	\N	\N	2019-09-19 00:00:00+07	0519.63.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5768	\N	1	2499	\N	\N	2019-10-04 00:00:00+07	2499.62.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5175	\N	1	2060	\N	\N	2019-04-04 00:00:00+07	2060.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4385	\N	1	1682	\N	\N	2018-02-05 00:00:00+07	1682.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
569	\N	2	551	\N	\N	2011-02-17 00:00:00+07	0551.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5839	\N	1	2530	\N	\N	2019-10-17 00:00:00+07	2530.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1371	\N	2	750	\N	\N	2016-03-15 00:00:00+07	0750.61.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4706	\N	1	700	\N	\N	2018-09-12 00:00:00+07	0700.64.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
908	\N	2	1680	\N	\N	2013-08-28 00:00:00+07	1680.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5775	\N	1	2313	\N	\N	2019-10-04 00:00:00+07	2313.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3379	\N	1	2075	\N	\N	2017-02-01 00:00:00+07	2075.62.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6509	\N	2	897	\N	\N	2021-06-21 00:00:00+07	0897.64.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-06-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5842	\N	1	815	\N	\N	2019-10-18 00:00:00+07	0815.64.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5843	\N	1	988	\N	\N	2019-10-18 00:00:00+07	0988.61.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5845	\N	1	2326	\N	\N	2019-10-21 00:00:00+07	2326.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5848	\N	1	989	\N	\N	2019-10-23 00:00:00+07	0989.61.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4777	\N	1	888	\N	\N	2018-10-02 00:00:00+07	0888.61.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6511	\N	1	2513	\N	\N	2021-06-23 00:00:00+07	2513.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-06-23 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5997	\N	1	2367	\N	\N	2020-01-13 00:00:00+07	2367.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5879	\N	1	2334	\N	\N	2019-11-04 00:00:00+07	2334.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2021-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5900	\N	1	821	\N	\N	2019-11-11 00:00:00+07	0821.64.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5592	\N	1	2220	\N	\N	2019-08-08 00:00:00+07	2220.65.221	1	\N	\N	SA	\N	10207	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5911	\N	1	2343	\N	\N	2019-11-14 00:00:00+07	2343.65.221	1	\N	\N	SA	\N	10207	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1285	\N	2	1331	\N	\N	2015-09-22 00:00:00+07	1331.65.221	1	\N	\N	SA	\N	10207	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1253	\N	2	1900	\N	\N	2015-07-27 00:00:00+07	1900.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1254	\N	2	1289	\N	\N	2015-07-31 00:00:00+07	1289.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1256	\N	2	1293	\N	\N	2015-08-11 00:00:00+07	1293.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1257	\N	2	1295	\N	\N	2015-08-19 00:00:00+07	1295.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1258	\N	2	824	\N	\N	2015-08-25 00:00:00+07	0824.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1259	\N	2	1297	\N	\N	2015-08-25 00:00:00+07	1297.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1260	\N	2	1298	\N	\N	2015-08-25 00:00:00+07	1298.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1261	\N	2	1300	\N	\N	2015-08-27 00:00:00+07	1300.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1262	\N	2	1302	\N	\N	2015-08-28 00:00:00+07	1302.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1263	\N	2	1052	\N	\N	2015-08-31 00:00:00+07	1052.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1264	\N	2	1905	\N	\N	2015-08-31 00:00:00+07	1905.62.221	1	2017-12-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1265	\N	2	1303	\N	\N	2015-08-31 00:00:00+07	1303.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1266	\N	2	1304	\N	\N	2015-08-31 00:00:00+07	1304.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1268	\N	2	1306	\N	\N	2015-09-02 00:00:00+07	1306.65.221	2	2017-06-06 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1269	\N	2	1307	\N	\N	2015-09-02 00:00:00+07	1307.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1270	\N	2	1608	\N	\N	2015-09-02 00:00:00+07	1608.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1271	\N	2	1309	\N	\N	2015-09-03 00:00:00+07	1309.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1272	\N	2	1310	\N	\N	2015-09-03 00:00:00+07	1310.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1273	\N	2	1311	\N	\N	2015-09-03 00:00:00+07	1311.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1274	\N	2	1912	\N	\N	2015-09-03 00:00:00+07	1912.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1275	\N	2	1913	\N	\N	2015-09-04 00:00:00+07	1913.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1276	\N	2	1316	\N	\N	2015-09-08 00:00:00+07	1316.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1277	\N	2	1317	\N	\N	2015-09-10 00:00:00+07	1317.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3979	\N	1	1653	\N	\N	2017-09-11 00:00:00+07	1653.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1278	\N	2	1915	\N	\N	2015-09-11 00:00:00+07	1915.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1279	\N	2	1319	\N	\N	2015-09-11 00:00:00+07	1319.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1280	\N	2	1320	\N	\N	2015-09-14 00:00:00+07	1320.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1281	\N	2	1322	\N	\N	2015-09-14 00:00:00+07	1322.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1282	\N	2	1325	\N	\N	2015-09-15 00:00:00+07	1325.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1283	\N	2	1328	\N	\N	2015-09-21 00:00:00+07	1328.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1284	\N	2	1330	\N	\N	2015-09-22 00:00:00+07	1330.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1286	\N	2	1917	\N	\N	2015-09-23 00:00:00+07	1917.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1287	\N	2	442	\N	\N	2015-09-25 00:00:00+07	0442.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1288	\N	2	1333	\N	\N	2015-09-28 00:00:00+07	1333.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1289	\N	2	1334	\N	\N	2015-09-28 00:00:00+07	1334.65.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1290	\N	2	1335	\N	\N	2015-09-28 00:00:00+07	1335.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1291	\N	2	1339	\N	\N	2015-10-02 00:00:00+07	1339.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1292	\N	2	1340	\N	\N	2015-10-02 00:00:00+07	1340.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1293	\N	2	1341	\N	\N	2015-10-05 00:00:00+07	1341.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1294	\N	2	1321	\N	\N	2015-10-05 00:00:00+07	1321.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1295	\N	2	1342	\N	\N	2015-10-05 00:00:00+07	1342.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1296	\N	2	1343	\N	\N	2015-10-05 00:00:00+07	1343.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1297	\N	2	1344	\N	\N	2015-10-05 00:00:00+07	1344.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1298	\N	2	1324	\N	\N	2015-10-06 00:00:00+07	1324.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1299	\N	2	1315	\N	\N	2015-10-07 00:00:00+07	1315.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1300	\N	2	444	\N	\N	2015-10-07 00:00:00+07	0444.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1301	\N	2	1337	\N	\N	2015-10-07 00:00:00+07	1337.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1302	\N	2	1318	\N	\N	2015-10-07 00:00:00+07	1318.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1304	\N	2	1347	\N	\N	2015-10-08 00:00:00+07	1347.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5423	\N	1	497	\N	\N	2019-06-19 00:00:00+07	0497.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1305	\N	2	1348	\N	\N	2015-10-09 00:00:00+07	1348.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1306	\N	2	1346	\N	\N	2015-10-09 00:00:00+07	1346.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1309	\N	2	1350	\N	\N	2015-10-09 00:00:00+07	1350.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1310	\N	2	1352	\N	\N	2015-10-09 00:00:00+07	1352.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1311	\N	2	1294	\N	\N	2015-10-12 00:00:00+07	1294.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1312	\N	2	446	\N	\N	2015-10-12 00:00:00+07	0446.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1313	\N	2	448	\N	\N	2015-10-12 00:00:00+07	0448.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1314	\N	2	1323	\N	\N	2015-10-12 00:00:00+07	1323.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1315	\N	2	447	\N	\N	2015-10-12 00:00:00+07	0447.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1316	\N	2	1353	\N	\N	2015-10-12 00:00:00+07	1353.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1317	\N	2	1354	\N	\N	2015-10-12 00:00:00+07	1354.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1318	\N	2	1355	\N	\N	2015-10-12 00:00:00+07	1355.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1319	\N	2	752	\N	\N	2015-10-13 00:00:00+07	0752.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1320	\N	2	1357	\N	\N	2015-10-26 00:00:00+07	1357.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1321	\N	2	455	\N	\N	2015-10-27 00:00:00+07	0455.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1322	\N	2	1361	\N	\N	2015-11-02 00:00:00+07	1361.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1323	\N	2	1364	\N	\N	2015-11-03 00:00:00+07	1364.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1324	\N	2	1356	\N	\N	2015-11-03 00:00:00+07	1356.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1325	\N	2	1312	\N	\N	2015-11-04 00:00:00+07	1312.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1326	\N	2	1929	\N	\N	2015-11-05 00:00:00+07	1929.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1327	\N	2	449	\N	\N	2015-11-09 00:00:00+07	0449.64.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-07-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4504	\N	1	1720	\N	\N	2018-04-13 00:00:00+07	1720.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1328	\N	2	1359	\N	\N	2015-11-09 00:00:00+07	1359.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1329	\N	2	1360	\N	\N	2015-11-09 00:00:00+07	1360.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1330	\N	2	1390	\N	\N	2015-11-10 00:00:00+07	1390.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1332	\N	2	1391	\N	\N	2015-11-12 00:00:00+07	1391.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1333	\N	2	1392	\N	\N	2015-11-12 00:00:00+07	1392.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1334	\N	2	1393	\N	\N	2015-11-12 00:00:00+07	1393.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1335	\N	2	1394	\N	\N	2015-11-13 00:00:00+07	1394.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1336	\N	2	1953	\N	\N	2015-11-18 00:00:00+07	1953.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1337	\N	2	1373	\N	\N	2015-11-20 00:00:00+07	1373.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1338	\N	2	1374	\N	\N	2015-11-23 00:00:00+07	1374.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1339	\N	2	1376	\N	\N	2015-11-24 00:00:00+07	1376.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1341	\N	2	1935	\N	\N	2015-12-10 00:00:00+07	1935.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1342	\N	2	1398	\N	\N	2015-12-22 00:00:00+07	1398.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1343	\N	2	764	\N	\N	2015-12-29 00:00:00+07	0764.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1347	\N	2	1397	\N	\N	2016-01-12 00:00:00+07	1397.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1348	\N	2	1408	\N	\N	2016-01-15 00:00:00+07	1408.65.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1349	\N	2	1409	\N	\N	2016-02-03 00:00:00+07	1409.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1350	\N	2	1410	\N	\N	2016-02-09 00:00:00+07	1410.65.221	1	2017-06-08 00:00:00+07	2019-08-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1351	\N	2	1411	\N	\N	2016-02-09 00:00:00+07	1411.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3980	\N	1	2188	\N	\N	2017-09-11 00:00:00+07	2188.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4502	\N	1	657	\N	\N	2018-04-12 00:00:00+07	0657.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1353	\N	2	465	\N	\N	2016-02-10 00:00:00+07	0465.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4501	\N	1	2251	\N	\N	2018-04-12 00:00:00+07	2251.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1356	\N	2	1415	\N	\N	2016-02-15 00:00:00+07	1415.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4500	\N	1	2250	\N	\N	2018-04-12 00:00:00+07	2250.62.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1357	\N	2	1416	\N	\N	2016-02-15 00:00:00+07	1416.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1358	\N	2	1418	\N	\N	2016-02-18 00:00:00+07	1418.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1359	\N	2	1419	\N	\N	2016-02-18 00:00:00+07	1419.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1360	\N	2	1420	\N	\N	2016-02-19 00:00:00+07	1420.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1361	\N	2	1421	\N	\N	2016-02-22 00:00:00+07	1421.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1362	\N	2	1250	\N	\N	2016-02-22 00:00:00+07	1250.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1363	\N	2	1417	\N	\N	2016-02-22 00:00:00+07	1417.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1364	\N	2	1422	\N	\N	2016-02-23 00:00:00+07	1422.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1365	\N	2	1960	\N	\N	2016-02-23 00:00:00+07	1960.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1366	\N	2	1423	\N	\N	2016-02-23 00:00:00+07	1423.65.221	2	2020-07-01 00:00:00+07	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2020-09-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1367	\N	2	1425	\N	\N	2016-03-01 00:00:00+07	1425.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1368	\N	2	297	\N	\N	2016-03-07 00:00:00+07	0297.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1369	\N	2	298	\N	\N	2016-03-10 00:00:00+07	0298.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1370	\N	2	1396	\N	\N	2016-03-14 00:00:00+07	1396.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4497	\N	1	1718	\N	\N	2018-04-11 00:00:00+07	1718.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1372	\N	2	1432	\N	\N	2016-03-16 00:00:00+07	1432.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1373	\N	2	1434	\N	\N	2016-03-17 00:00:00+07	1434.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1374	\N	2	1435	\N	\N	2016-03-21 00:00:00+07	1435.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1375	\N	2	1436	\N	\N	2016-03-21 00:00:00+07	1436.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1376	\N	2	1439	\N	\N	2016-03-24 00:00:00+07	1439.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1377	\N	2	1440	\N	\N	2016-03-24 00:00:00+07	1440.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1378	\N	2	1441	\N	\N	2016-03-24 00:00:00+07	1441.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1379	\N	2	1443	\N	\N	2016-03-31 00:00:00+07	1443.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1380	\N	2	1442	\N	\N	2016-04-01 00:00:00+07	1442.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1381	\N	2	1970	\N	\N	2016-04-04 00:00:00+07	1970.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1382	\N	2	1445	\N	\N	2016-04-04 00:00:00+07	1445.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1383	\N	2	1446	\N	\N	2016-04-04 00:00:00+07	1446.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1386	\N	2	1448	\N	\N	2016-04-11 00:00:00+07	1448.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4496	\N	1	1717	\N	\N	2018-04-11 00:00:00+07	1717.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1387	\N	2	1449	\N	\N	2016-04-11 00:00:00+07	1449.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1388	\N	2	1450	\N	\N	2016-04-14 00:00:00+07	1450.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1389	\N	2	1453	\N	\N	2016-04-18 00:00:00+07	1453.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1390	\N	2	1378	\N	\N	2016-04-20 00:00:00+07	1378.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1391	\N	2	1456	\N	\N	2016-04-25 00:00:00+07	1456.65.221	1	\N	2020-09-15 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1392	\N	2	1462	\N	\N	2016-04-27 00:00:00+07	1462.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1393	\N	2	1452	\N	\N	2016-04-27 00:00:00+07	1452.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1394	\N	2	1460	\N	\N	2016-04-28 00:00:00+07	1460.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1395	\N	2	1468	\N	\N	2016-04-28 00:00:00+07	1468.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3982	\N	1	2189	\N	\N	2017-09-11 00:00:00+07	2189.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4495	\N	1	656	\N	\N	2018-04-10 00:00:00+07	0656.64.221	2	2018-07-01 00:00:00+07	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1396	\N	2	1469	\N	\N	2016-04-28 00:00:00+07	1469.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1397	\N	2	1458	\N	\N	2016-04-28 00:00:00+07	1458.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1398	\N	2	1463	\N	\N	2016-04-28 00:00:00+07	1463.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1399	\N	2	1470	\N	\N	2016-04-28 00:00:00+07	1470.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1400	\N	2	1466	\N	\N	2016-04-28 00:00:00+07	1466.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1401	\N	2	1457	\N	\N	2016-04-28 00:00:00+07	1457.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1402	\N	2	1461	\N	\N	2016-04-28 00:00:00+07	1461.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4493	\N	1	1716	\N	\N	2018-04-10 00:00:00+07	1716.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4492	\N	1	1715	\N	\N	2018-04-09 00:00:00+07	1715.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1403	\N	2	1467	\N	\N	2016-04-28 00:00:00+07	1467.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3466	\N	1	538	\N	\N	2017-03-01 00:00:00+07	0538.64.221	1	\N	2020-03-06 00:00:00+07	SA	\N	10208	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1404	\N	2	1471	\N	\N	2016-04-28 00:00:00+07	1471.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1405	\N	2	1459	\N	\N	2016-04-29 00:00:00+07	1459.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1406	\N	2	1465	\N	\N	2016-05-02 00:00:00+07	1465.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4476	\N	1	653	\N	\N	2018-03-27 00:00:00+07	0653.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4475	\N	1	1713	\N	\N	2018-03-26 00:00:00+07	1713.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1407	\N	2	1474	\N	\N	2016-05-04 00:00:00+07	1474.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1408	\N	2	1475	\N	\N	2016-05-09 00:00:00+07	1475.65.221	1	\N	2020-10-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1409	\N	2	1476	\N	\N	2016-05-09 00:00:00+07	1476.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1410	\N	2	1478	\N	\N	2016-05-09 00:00:00+07	1478.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1413	\N	2	1485	\N	\N	2016-06-16 00:00:00+07	1485.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1414	\N	2	1486	\N	\N	2016-07-12 00:00:00+07	1486.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1415	\N	2	1497	\N	\N	2016-07-25 00:00:00+07	1497.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1416	\N	2	1498	\N	\N	2016-07-25 00:00:00+07	1498.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4474	\N	1	652	\N	\N	2018-03-26 00:00:00+07	0652.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1417	\N	2	782	\N	\N	2016-07-26 00:00:00+07	0782.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1419	\N	2	1483	\N	\N	2016-08-02 00:00:00+07	1483.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1420	\N	2	1505	\N	\N	2016-08-08 00:00:00+07	1505.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1421	\N	2	1506	\N	\N	2016-08-09 00:00:00+07	1506.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1422	\N	2	1509	\N	\N	2016-08-22 00:00:00+07	1509.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4465	\N	1	1709	\N	\N	2018-03-22 00:00:00+07	1709.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1423	\N	2	1510	\N	\N	2016-09-01 00:00:00+07	1510.65.221	1	\N	2018-12-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1424	\N	2	1504	\N	\N	2016-09-05 00:00:00+07	1504.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1425	\N	2	1511	\N	\N	2016-09-06 00:00:00+07	1511.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1426	\N	2	1512	\N	\N	2016-09-07 00:00:00+07	1512.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1427	\N	2	2012	\N	\N	2016-09-07 00:00:00+07	2012.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4472	\N	1	1712	\N	\N	2018-03-26 00:00:00+07	1712.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1428	\N	2	1514	\N	\N	2016-09-08 00:00:00+07	1514.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4470	\N	1	650	\N	\N	2018-03-26 00:00:00+07	0650.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1432	\N	2	1516	\N	\N	2016-09-09 00:00:00+07	1516.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4469	\N	1	1710	\N	\N	2018-03-26 00:00:00+07	1710.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1433	\N	2	2016	\N	\N	2016-09-15 00:00:00+07	2016.62.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1434	\N	2	1521	\N	\N	2016-09-21 00:00:00+07	1521.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1435	\N	2	479	\N	\N	2016-10-04 00:00:00+07	0479.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1436	\N	2	1524	\N	\N	2016-10-05 00:00:00+07	1524.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1437	\N	2	1519	\N	\N	2016-10-06 00:00:00+07	1519.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1438	\N	2	1515	\N	\N	2016-10-07 00:00:00+07	1515.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1439	\N	2	1482	\N	\N	2016-10-10 00:00:00+07	1482.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5171	\N	1	2057	\N	\N	2019-04-04 00:00:00+07	2057.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5207	\N	1	2073	\N	\N	2019-04-11 00:00:00+07	2073.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7483	\N	1	2910	\N	\N	2022-06-21 00:00:00+07	2910.62.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5119	\N	1	2038	\N	\N	2019-03-18 00:00:00+07	2038.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4468	\N	1	649	\N	\N	2018-03-23 00:00:00+07	0649.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5319	\N	1	2114	\N	\N	2019-05-21 00:00:00+07	2114.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4455	\N	1	1707	\N	\N	2018-03-20 00:00:00+07	1707.65.221	1	\N	2020-01-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7477	\N	1	2736	\N	\N	2022-06-16 00:00:00+07	2736.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7471	\N	1	2734	\N	\N	2022-06-15 00:00:00+07	2734.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7466	\N	1	2733	\N	\N	2022-06-15 00:00:00+07	2733.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7473	\N	1	2735	\N	\N	2022-06-16 00:00:00+07	2735.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5172	\N	1	2058	\N	\N	2019-04-04 00:00:00+07	2058.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5211	\N	1	2075	\N	\N	2019-04-11 00:00:00+07	2075.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5215	\N	1	2078	\N	\N	2019-04-12 00:00:00+07	2078.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5216	\N	1	2079	\N	\N	2019-04-12 00:00:00+07	2079.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5111	\N	1	2031	\N	\N	2019-03-18 00:00:00+07	2031.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7454	\N	1	2730	\N	\N	2022-06-10 00:00:00+07	2730.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5191	\N	1	2067	\N	\N	2019-04-08 00:00:00+07	2067.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5217	\N	1	2080	\N	\N	2019-04-12 00:00:00+07	2080.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5213	\N	1	472	\N	\N	2019-04-12 00:00:00+07	0472.63.221	2	2019-12-04 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4443	\N	1	639	\N	\N	2018-03-14 00:00:00+07	0639.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5208	\N	1	2074	\N	\N	2019-04-11 00:00:00+07	2074.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7446	\N	1	2728	\N	\N	2022-06-07 00:00:00+07	2728.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7445	\N	1	2727	\N	\N	2022-06-07 00:00:00+07	2727.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5060	\N	1	2001	\N	\N	2019-03-05 00:00:00+07	2001.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7444	\N	1	2900	\N	\N	2022-06-07 00:00:00+07	2900.62.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5102	\N	1	2027	\N	\N	2019-03-14 00:00:00+07	2027.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7439	\N	1	2726	\N	\N	2022-06-06 00:00:00+07	2726.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7437	\N	1	2725	\N	\N	2022-06-03 00:00:00+07	2725.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5100	\N	1	2025	\N	\N	2019-03-14 00:00:00+07	2025.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5292	\N	1	2385	\N	\N	2019-05-15 00:00:00+07	2385.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5099	\N	1	2024	\N	\N	2019-03-14 00:00:00+07	2024.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5206	\N	1	471	\N	\N	2019-04-11 00:00:00+07	0471.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5139	\N	1	2046	\N	\N	2019-03-25 00:00:00+07	2046.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5214	\N	1	2077	\N	\N	2019-04-12 00:00:00+07	2077.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5226	\N	1	2083	\N	\N	2019-04-16 00:00:00+07	2083.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5092	\N	1	2018	\N	\N	2019-03-13 00:00:00+07	2018.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5090	\N	1	2017	\N	\N	2019-03-12 00:00:00+07	2017.65.221	1	\N	2019-08-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4439	\N	1	637	\N	\N	2018-03-13 00:00:00+07	0637.64.221	1	\N	2019-08-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4437	\N	1	2238	\N	\N	2018-03-12 00:00:00+07	2238.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5089	\N	1	2016	\N	\N	2019-03-12 00:00:00+07	2016.65.221	1	\N	2019-08-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5329	\N	1	2115	\N	\N	2019-05-22 00:00:00+07	2115.65.221	1	\N	2020-09-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5331	\N	1	2116	\N	\N	2019-05-22 00:00:00+07	2116.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7418	\N	1	2720	\N	\N	2022-05-24 00:00:00+07	2720.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4429	\N	1	1699	\N	\N	2018-03-08 00:00:00+07	1699.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5098	\N	1	2023	\N	\N	2019-03-14 00:00:00+07	2023.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4424	\N	1	869	\N	\N	2018-03-05 00:00:00+07	0869.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5228	\N	1	2085	\N	\N	2019-04-18 00:00:00+07	2085.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3989	\N	1	1654	\N	\N	2017-09-12 00:00:00+07	1654.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5436	\N	1	2170	\N	\N	2019-06-24 00:00:00+07	2170.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5160	\N	1	2052	\N	\N	2019-03-28 00:00:00+07	2052.65.221	1	\N	2020-11-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5145	\N	1	2047	\N	\N	2019-03-26 00:00:00+07	2047.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3502	\N	1	2124	\N	\N	2017-03-01 00:00:00+07	2124.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7402	\N	1	2885	\N	\N	2022-04-25 00:00:00+07	2885.62.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-04-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7401	\N	1	2714	\N	\N	2022-04-25 00:00:00+07	2714.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-04-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5233	\N	1	473	\N	\N	2019-04-24 00:00:00+07	0473.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5212	\N	1	2076	\N	\N	2019-04-11 00:00:00+07	2076.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5267	\N	1	2093	\N	\N	2019-05-07 00:00:00+07	2093.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5097	\N	1	2022	\N	\N	2019-03-14 00:00:00+07	2022.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5234	\N	1	2086	\N	\N	2019-04-24 00:00:00+07	2086.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4422	\N	1	632	\N	\N	2018-03-01 00:00:00+07	0632.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5571	\N	1	2468	\N	\N	2019-08-05 00:00:00+07	2468.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5096	\N	1	2021	\N	\N	2019-03-14 00:00:00+07	2021.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5095	\N	1	2020	\N	\N	2019-03-13 00:00:00+07	2020.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5333	\N	1	2117	\N	\N	2019-05-22 00:00:00+07	2117.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5077	\N	1	2013	\N	\N	2019-03-11 00:00:00+07	2013.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5076	\N	1	2012	\N	\N	2019-03-11 00:00:00+07	2012.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5311	\N	1	2113	\N	\N	2019-05-17 00:00:00+07	2113.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5075	\N	1	2011	\N	\N	2019-03-11 00:00:00+07	2011.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5074	\N	1	2010	\N	\N	2019-03-11 00:00:00+07	2010.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5269	\N	1	764	\N	\N	2019-05-07 00:00:00+07	0764.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5572	\N	1	2469	\N	\N	2019-08-05 00:00:00+07	2469.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5438	\N	1	2172	\N	\N	2019-06-25 00:00:00+07	2172.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5334	\N	1	2118	\N	\N	2019-05-22 00:00:00+07	2118.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5148	\N	1	2048	\N	\N	2019-03-27 00:00:00+07	2048.65.221	2	2020-03-01 00:00:00+07	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5166	\N	1	2054	\N	\N	2019-04-02 00:00:00+07	2054.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7348	\N	1	1160	\N	\N	2022-03-23 00:00:00+07	1160.61.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-03-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5167	\N	1	2360	\N	\N	2019-04-02 00:00:00+07	2360.62.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5270	\N	1	2094	\N	\N	2019-05-08 00:00:00+07	2094.65.221	2	2020-08-01 00:00:00+07	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5185	\N	1	2064	\N	\N	2019-04-04 00:00:00+07	2064.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7338	\N	1	2870	\N	\N	2022-03-21 00:00:00+07	2870.62.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-03-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5065	\N	1	2007	\N	\N	2019-03-06 00:00:00+07	2007.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5064	\N	1	2006	\N	\N	2019-03-06 00:00:00+07	2006.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5283	\N	1	2098	\N	\N	2019-05-09 00:00:00+07	2098.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5134	\N	1	2043	\N	\N	2019-03-21 00:00:00+07	2043.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5062	\N	1	2004	\N	\N	2019-03-05 00:00:00+07	2004.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7300	\N	1	2683	\N	\N	2022-03-10 00:00:00+07	2683.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-03-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5061	\N	1	2002	\N	\N	2019-03-05 00:00:00+07	2002.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5133	\N	1	2042	\N	\N	2019-03-21 00:00:00+07	2042.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5128	\N	1	754	\N	\N	2019-03-20 00:00:00+07	0754.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5247	\N	1	2089	\N	\N	2019-05-02 00:00:00+07	2089.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5248	\N	1	2090	\N	\N	2019-05-02 00:00:00+07	2090.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5249	\N	1	2091	\N	\N	2019-05-03 00:00:00+07	2091.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3771	\N	1	1630	\N	\N	2017-06-06 00:00:00+07	1630.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5168	\N	1	2361	\N	\N	2019-04-02 00:00:00+07	2361.62.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5058	\N	1	1999	\N	\N	2019-03-04 00:00:00+07	1999.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5021	\N	1	1980	\N	\N	2019-02-22 00:00:00+07	1980.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5169	\N	1	2055	\N	\N	2019-04-02 00:00:00+07	2055.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5120	\N	1	2039	\N	\N	2019-03-18 00:00:00+07	2039.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5019	\N	1	1978	\N	\N	2019-02-22 00:00:00+07	1978.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5277	\N	1	2382	\N	\N	2019-05-09 00:00:00+07	2382.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5018	\N	1	1974	\N	\N	2019-02-22 00:00:00+07	1974.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7262	\N	1	2674	\N	\N	2022-02-16 00:00:00+07	2674.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-02-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5016	\N	1	1973	\N	\N	2019-02-22 00:00:00+07	1973.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5015	\N	1	1972	\N	\N	2019-02-22 00:00:00+07	1972.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5051	\N	1	1994	\N	\N	2019-02-28 00:00:00+07	1994.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
634	\N	2	679	\N	\N	2014-02-06 00:00:00+07	0679.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5170	\N	1	2056	\N	\N	2019-04-04 00:00:00+07	2056.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5282	\N	1	2097	\N	\N	2019-05-09 00:00:00+07	2097.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5013	\N	1	1971	\N	\N	2019-02-21 00:00:00+07	1971.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5010	\N	1	1969	\N	\N	2019-02-21 00:00:00+07	1969.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5008	\N	1	1968	\N	\N	2019-02-21 00:00:00+07	1968.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5007	\N	1	1967	\N	\N	2019-02-21 00:00:00+07	1967.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5006	\N	1	1966	\N	\N	2019-02-21 00:00:00+07	1966.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7248	\N	1	996	\N	\N	2022-02-09 00:00:00+07	0996.64.221	1	2022-04-04 00:00:00+07	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-02-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7225	\N	1	2663	\N	\N	2022-02-07 00:00:00+07	2663.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4405	\N	1	1692	\N	\N	2018-02-20 00:00:00+07	1692.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3572	\N	1	1618	\N	\N	2017-04-03 00:00:00+07	1618.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5005	\N	1	1965	\N	\N	2019-02-21 00:00:00+07	1965.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5004	\N	1	1964	\N	\N	2019-02-21 00:00:00+07	1964.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2020-10-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4997	\N	1	1958	\N	\N	2019-02-20 00:00:00+07	1958.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7213	\N	1	988	\N	\N	2022-02-03 00:00:00+07	0988.64.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-02-03 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4404	\N	1	628	\N	\N	2018-02-19 00:00:00+07	0628.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4403	\N	1	627	\N	\N	2018-02-19 00:00:00+07	0627.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5048	\N	1	1991	\N	\N	2019-02-28 00:00:00+07	1991.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7169	\N	1	981	\N	\N	2022-01-24 00:00:00+07	0981.64.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-01-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4989	\N	1	1953	\N	\N	2019-02-15 00:00:00+07	1953.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
595	\N	2	385	\N	\N	2014-09-25 00:00:00+07	0385.64.221	1	2017-06-08 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4987	\N	1	902	\N	\N	2019-02-15 00:00:00+07	0902.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
599	\N	2	386	\N	\N	2014-09-29 00:00:00+07	0386.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4407	\N	1	629	\N	\N	2018-02-21 00:00:00+07	0629.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
600	\N	2	387	\N	\N	2014-10-29 00:00:00+07	0387.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5056	\N	1	1998	\N	\N	2019-03-01 00:00:00+07	1998.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
601	\N	2	388	\N	\N	2014-12-01 00:00:00+07	0388.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
602	\N	2	389	\N	\N	2014-12-02 00:00:00+07	0389.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5343	\N	1	2120	\N	\N	2019-05-23 00:00:00+07	2120.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
604	\N	2	390	\N	\N	2014-12-02 00:00:00+07	0390.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
606	\N	2	392	\N	\N	2014-12-11 00:00:00+07	0392.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5055	\N	1	1997	\N	\N	2019-03-01 00:00:00+07	1997.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5344	\N	1	2121	\N	\N	2019-05-23 00:00:00+07	2121.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5296	\N	1	2107	\N	\N	2019-05-15 00:00:00+07	2107.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4399	\N	1	1689	\N	\N	2018-02-09 00:00:00+07	1689.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4397	\N	1	1688	\N	\N	2018-02-08 00:00:00+07	1688.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
608	\N	2	646	\N	\N	2013-08-28 00:00:00+07	0646.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
609	\N	2	649	\N	\N	2014-01-08 00:00:00+07	0649.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
611	\N	2	651	\N	\N	2014-08-21 00:00:00+07	0651.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
612	\N	2	654	\N	\N	2014-02-03 00:00:00+07	0654.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
613	\N	2	656	\N	\N	2013-11-26 00:00:00+07	0656.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
614	\N	2	658	\N	\N	2014-02-17 00:00:00+07	0658.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
615	\N	2	659	\N	\N	2014-02-17 00:00:00+07	0659.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5054	\N	1	1996	\N	\N	2019-03-01 00:00:00+07	1996.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2950	\N	1	2032	\N	\N	2017-01-02 00:00:00+07	2032.62.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-03-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
616	\N	2	660	\N	\N	2013-12-05 00:00:00+07	0660.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
617	\N	2	662	\N	\N	2013-11-26 00:00:00+07	0662.65.221	1	\N	2019-09-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
618	\N	2	663	\N	\N	2013-11-26 00:00:00+07	0663.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4396	\N	1	1687	\N	\N	2018-02-08 00:00:00+07	1687.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5042	\N	1	1989	\N	\N	2019-02-28 00:00:00+07	1989.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4978	\N	1	1950	\N	\N	2019-02-14 00:00:00+07	1950.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4977	\N	1	1949	\N	\N	2019-02-14 00:00:00+07	1949.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
620	\N	2	664	\N	\N	2014-01-08 00:00:00+07	0664.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
621	\N	2	665	\N	\N	2014-01-09 00:00:00+07	0665.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
622	\N	2	666	\N	\N	2014-01-09 00:00:00+07	0666.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4976	\N	1	1948	\N	\N	2019-02-14 00:00:00+07	1948.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5041	\N	1	1988	\N	\N	2019-02-28 00:00:00+07	1988.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
623	\N	2	667	\N	\N	2014-01-15 00:00:00+07	0667.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5348	\N	1	2123	\N	\N	2019-05-24 00:00:00+07	2123.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5039	\N	1	1987	\N	\N	2019-02-27 00:00:00+07	1987.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5350	\N	1	491	\N	\N	2019-05-24 00:00:00+07	0491.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4971	\N	1	1945	\N	\N	2019-02-13 00:00:00+07	1945.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
624	\N	2	668	\N	\N	2014-12-08 00:00:00+07	0668.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2951	\N	1	1566	\N	\N	2017-01-02 00:00:00+07	1566.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4970	\N	1	1944	\N	\N	2019-02-13 00:00:00+07	1944.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4391	\N	1	626	\N	\N	2018-02-05 00:00:00+07	0626.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4390	\N	1	1686	\N	\N	2018-02-05 00:00:00+07	1686.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
625	\N	2	669	\N	\N	2014-02-21 00:00:00+07	0669.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
626	\N	2	670	\N	\N	2014-02-21 00:00:00+07	0670.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5036	\N	1	1985	\N	\N	2019-02-27 00:00:00+07	1985.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
627	\N	2	671	\N	\N	2014-01-16 00:00:00+07	0671.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4389	\N	1	1685	\N	\N	2018-02-05 00:00:00+07	1685.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5034	\N	1	1984	\N	\N	2019-02-27 00:00:00+07	1984.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
628	\N	2	672	\N	\N	2014-01-16 00:00:00+07	0672.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5032	\N	1	1983	\N	\N	2019-02-27 00:00:00+07	1983.65.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5031	\N	1	1982	\N	\N	2019-02-26 00:00:00+07	1982.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
629	\N	2	673	\N	\N	2014-02-17 00:00:00+07	0673.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4968	\N	1	1943	\N	\N	2019-02-13 00:00:00+07	1943.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4388	\N	1	1684	\N	\N	2018-02-05 00:00:00+07	1684.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
630	\N	2	674	\N	\N	2014-02-17 00:00:00+07	0674.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
631	\N	2	676	\N	\N	2014-01-24 00:00:00+07	0676.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
632	\N	2	677	\N	\N	2014-02-05 00:00:00+07	0677.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
633	\N	2	678	\N	\N	2014-02-05 00:00:00+07	0678.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
635	\N	2	680	\N	\N	2014-02-11 00:00:00+07	0680.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
636	\N	2	681	\N	\N	2014-02-11 00:00:00+07	0681.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
637	\N	2	682	\N	\N	2014-02-11 00:00:00+07	0682.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
638	\N	2	683	\N	\N	2014-02-11 00:00:00+07	0683.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6941	\N	1	2610	\N	\N	2021-11-25 00:00:00+07	2610.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-11-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
639	\N	2	684	\N	\N	2014-02-12 00:00:00+07	0684.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
640	\N	2	685	\N	\N	2014-02-12 00:00:00+07	0685.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4958	\N	1	1942	\N	\N	2019-02-11 00:00:00+07	1942.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4387	\N	1	1683	\N	\N	2018-02-05 00:00:00+07	1683.65.221	2	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
641	\N	2	687	\N	\N	2014-02-17 00:00:00+07	0687.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4952	\N	1	1940	\N	\N	2019-02-07 00:00:00+07	1940.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4951	\N	1	1939	\N	\N	2019-02-07 00:00:00+07	1939.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4949	\N	1	1938	\N	\N	2019-02-06 00:00:00+07	1938.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4384	\N	1	1681	\N	\N	2018-02-05 00:00:00+07	1681.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
642	\N	2	688	\N	\N	2014-02-17 00:00:00+07	0688.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
643	\N	2	689	\N	\N	2014-02-26 00:00:00+07	0689.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3573	\N	1	1619	\N	\N	2017-04-03 00:00:00+07	1619.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6925	\N	1	2602	\N	\N	2021-11-22 00:00:00+07	2602.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-11-22 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4948	\N	1	1937	\N	\N	2019-02-06 00:00:00+07	1937.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4946	\N	1	1935	\N	\N	2019-01-31 00:00:00+07	1935.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4383	\N	1	625	\N	\N	2018-02-05 00:00:00+07	0625.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4945	\N	1	739	\N	\N	2019-01-31 00:00:00+07	0739.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4944	\N	1	1934	\N	\N	2019-01-31 00:00:00+07	1934.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
644	\N	2	689	\N	\N	2014-03-07 00:00:00+07	0689.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
645	\N	2	690	\N	\N	2014-03-14 00:00:00+07	0690.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
646	\N	2	691	\N	\N	2014-02-19 00:00:00+07	0691.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
647	\N	2	692	\N	\N	2014-02-19 00:00:00+07	0692.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4943	\N	1	1933	\N	\N	2019-01-31 00:00:00+07	1933.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
648	\N	2	693	\N	\N	2014-04-03 00:00:00+07	0693.61.221	2	2019-05-01 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4382	\N	1	1680	\N	\N	2018-02-05 00:00:00+07	1680.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
649	\N	2	693	\N	\N	2014-02-19 00:00:00+07	0693.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4381	\N	1	2229	\N	\N	2018-02-05 00:00:00+07	2229.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4380	\N	1	1679	\N	\N	2018-02-05 00:00:00+07	1679.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4379	\N	1	624	\N	\N	2018-02-02 00:00:00+07	0624.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
650	\N	2	694	\N	\N	2014-02-20 00:00:00+07	0694.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4942	\N	1	1932	\N	\N	2019-01-30 00:00:00+07	1932.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4377	\N	1	623	\N	\N	2018-01-31 00:00:00+07	0623.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
651	\N	2	695	\N	\N	2014-02-20 00:00:00+07	0695.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
652	\N	2	696	\N	\N	2014-02-20 00:00:00+07	0696.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
653	\N	2	697	\N	\N	2014-02-21 00:00:00+07	0697.65.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
654	\N	2	698	\N	\N	2014-02-21 00:00:00+07	0698.65.221	1	2017-08-02 00:00:00+07	2020-10-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4935	\N	1	1930	\N	\N	2019-01-25 00:00:00+07	1930.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6898	\N	1	2591	\N	\N	2021-11-10 00:00:00+07	2591.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4376	\N	1	622	\N	\N	2018-01-31 00:00:00+07	0622.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4375	\N	1	621	\N	\N	2018-01-31 00:00:00+07	0621.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
657	\N	2	699	\N	\N	2014-02-21 00:00:00+07	0699.65.221	1	2017-08-02 00:00:00+07	2020-10-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5374	\N	1	2131	\N	\N	2019-06-11 00:00:00+07	2131.65.221	1	\N	2019-08-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4374	\N	1	620	\N	\N	2018-01-31 00:00:00+07	0620.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
658	\N	2	700	\N	\N	2014-02-21 00:00:00+07	0700.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4373	\N	1	619	\N	\N	2018-01-31 00:00:00+07	0619.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4372	\N	1	618	\N	\N	2018-01-31 00:00:00+07	0618.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
659	\N	2	701	\N	\N	2014-02-21 00:00:00+07	0701.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4371	\N	1	617	\N	\N	2018-01-31 00:00:00+07	0617.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
660	\N	2	702	\N	\N	2014-02-21 00:00:00+07	0702.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4917	\N	1	1925	\N	\N	2018-12-28 00:00:00+07	1925.65.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
661	\N	2	703	\N	\N	2014-02-21 00:00:00+07	0703.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
662	\N	2	704	\N	\N	2014-02-21 00:00:00+07	0704.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
664	\N	2	705	\N	\N	2014-06-19 00:00:00+07	0705.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4368	\N	1	1678	\N	\N	2018-01-30 00:00:00+07	1678.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
665	\N	2	705	\N	\N	2014-02-21 00:00:00+07	0705.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
666	\N	2	706	\N	\N	2014-06-19 00:00:00+07	0706.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4877	\N	1	447	\N	\N	2018-12-03 00:00:00+07	0447.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
667	\N	2	707	\N	\N	2014-06-20 00:00:00+07	0707.61.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
668	\N	2	708	\N	\N	2014-02-24 00:00:00+07	0708.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4359	\N	1	616	\N	\N	2018-01-24 00:00:00+07	0616.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4355	\N	1	615	\N	\N	2018-01-22 00:00:00+07	0615.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4861	\N	1	1902	\N	\N	2018-11-27 00:00:00+07	1902.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
669	\N	2	709	\N	\N	2014-02-25 00:00:00+07	0709.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4354	\N	1	614	\N	\N	2018-01-22 00:00:00+07	0614.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5385	\N	1	2140	\N	\N	2019-06-12 00:00:00+07	2140.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
670	\N	2	710	\N	\N	2014-02-27 00:00:00+07	0710.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
671	\N	2	712	\N	\N	2014-09-12 00:00:00+07	0712.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4353	\N	1	613	\N	\N	2018-01-22 00:00:00+07	0613.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
672	\N	2	712	\N	\N	2014-02-28 00:00:00+07	0712.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
673	\N	2	714	\N	\N	2014-03-04 00:00:00+07	0714.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5386	\N	1	2141	\N	\N	2019-06-12 00:00:00+07	2141.65.221	1	\N	\N	SA	\N	10208	solusiti	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
674	\N	2	715	\N	\N	2014-03-04 00:00:00+07	0715.65.221	2	2016-12-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
675	\N	2	716	\N	\N	2014-03-04 00:00:00+07	0716.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5387	\N	1	2142	\N	\N	2019-06-12 00:00:00+07	2142.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5388	\N	1	2143	\N	\N	2019-06-12 00:00:00+07	2143.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2909	\N	1	486	\N	\N	2016-12-01 00:00:00+07	0486.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2910	\N	1	484	\N	\N	2016-12-01 00:00:00+07	0484.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2911	\N	1	485	\N	\N	2017-03-10 00:00:00+07	0485.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4830	\N	1	1894	\N	\N	2018-10-29 00:00:00+07	1894.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5389	\N	1	2144	\N	\N	2019-06-12 00:00:00+07	2144.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4825	\N	1	2293	\N	\N	2018-10-23 00:00:00+07	2293.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
676	\N	2	717	\N	\N	2014-03-04 00:00:00+07	0717.65.221	1	2016-12-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2941	\N	1	1591	\N	\N	2017-03-17 00:00:00+07	1591.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
678	\N	2	718	\N	\N	2014-03-05 00:00:00+07	0718.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5454	\N	1	2185	\N	\N	2019-06-27 00:00:00+07	2185.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4824	\N	1	1892	\N	\N	2018-10-18 00:00:00+07	1892.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4823	\N	1	1891	\N	\N	2018-10-18 00:00:00+07	1891.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2923	\N	1	1590	\N	\N	2017-02-01 00:00:00+07	1590.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2924	\N	1	1589	\N	\N	2017-02-01 00:00:00+07	1589.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
679	\N	2	719	\N	\N	2014-03-07 00:00:00+07	0719.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4822	\N	1	1890	\N	\N	2018-10-18 00:00:00+07	1890.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2946	\N	1	1558	\N	\N	2017-02-01 00:00:00+07	1558.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2947	\N	1	1559	\N	\N	2017-02-01 00:00:00+07	1559.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2948	\N	1	505	\N	\N	2017-02-01 00:00:00+07	0505.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5457	\N	1	2186	\N	\N	2019-06-27 00:00:00+07	2186.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5459	\N	1	2187	\N	\N	2019-06-27 00:00:00+07	2187.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4821	\N	1	1889	\N	\N	2018-10-18 00:00:00+07	1889.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
680	\N	2	720	\N	\N	2014-03-07 00:00:00+07	0720.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2962	\N	1	309	\N	\N	2016-12-01 00:00:00+07	0309.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2964	\N	1	496	\N	\N	2017-02-01 00:00:00+07	0496.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2965	\N	1	1543	\N	\N	2017-02-01 00:00:00+07	1543.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2966	\N	1	1563	\N	\N	2017-01-02 00:00:00+07	1563.65.221	1	2017-06-08 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2967	\N	1	2083	\N	\N	2017-02-01 00:00:00+07	2083.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
681	\N	2	721	\N	\N	2014-03-07 00:00:00+07	0721.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5460	\N	1	2188	\N	\N	2019-06-27 00:00:00+07	2188.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4820	\N	1	1888	\N	\N	2018-10-18 00:00:00+07	1888.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
682	\N	2	722	\N	\N	2014-03-07 00:00:00+07	0722.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5467	\N	1	2190	\N	\N	2019-07-02 00:00:00+07	2190.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
683	\N	2	725	\N	\N	2014-05-30 00:00:00+07	0725.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2974	\N	1	1597	\N	\N	2017-02-01 00:00:00+07	1597.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4819	\N	1	1887	\N	\N	2018-10-18 00:00:00+07	1887.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6181	\N	1	2425	\N	\N	2020-09-07 00:00:00+07	2425.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	3500000	\N	\N	2020-08-09 00:00:00+07	200	00.00	00.00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4310	\N	1	608	\N	\N	2017-12-12 00:00:00+07	0608.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2980	\N	1	1595	\N	\N	2017-02-01 00:00:00+07	1595.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6179	\N	1	2423	\N	\N	2020-09-01 00:00:00+07	2423.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4818	\N	1	1886	\N	\N	2018-10-18 00:00:00+07	1886.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
684	\N	2	727	\N	\N	2014-03-12 00:00:00+07	0727.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
685	\N	2	728	\N	\N	2014-03-12 00:00:00+07	0728.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4299	\N	1	1676	\N	\N	2017-12-08 00:00:00+07	1676.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6169	\N	1	1029	\N	\N	2020-08-19 00:00:00+07	1029.61.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2020-12-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4817	\N	1	1885	\N	\N	2018-10-18 00:00:00+07	1885.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4816	\N	1	1884	\N	\N	2018-10-18 00:00:00+07	1884.65.221	2	2018-11-01 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4292	\N	1	607	\N	\N	2017-12-06 00:00:00+07	0607.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
687	\N	2	730	\N	\N	2014-03-19 00:00:00+07	0730.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4815	\N	1	1883	\N	\N	2018-10-17 00:00:00+07	1883.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4814	\N	1	1882	\N	\N	2018-10-17 00:00:00+07	1882.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
688	\N	2	731	\N	\N	2014-03-24 00:00:00+07	0731.65.221	1	\N	2018-08-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
689	\N	2	732	\N	\N	2014-03-21 00:00:00+07	0732.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
690	\N	2	733	\N	\N	2014-03-21 00:00:00+07	0733.65.221	2	2019-05-01 00:00:00+07	2019-10-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4813	\N	1	1881	\N	\N	2018-10-17 00:00:00+07	1881.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4287	\N	1	606	\N	\N	2017-12-04 00:00:00+07	0606.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4812	\N	1	1880	\N	\N	2018-10-17 00:00:00+07	1880.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4286	\N	1	605	\N	\N	2017-12-04 00:00:00+07	0605.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4811	\N	1	1879	\N	\N	2018-10-17 00:00:00+07	1879.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4285	\N	1	604	\N	\N	2017-12-04 00:00:00+07	0604.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4810	\N	1	1878	\N	\N	2018-10-17 00:00:00+07	1878.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4284	\N	1	603	\N	\N	2017-12-04 00:00:00+07	0603.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4809	\N	1	1877	\N	\N	2018-10-17 00:00:00+07	1877.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4281	\N	1	2024	\N	\N	2016-11-01 00:00:00+07	2024.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
691	\N	2	734	\N	\N	2014-03-21 00:00:00+07	0734.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4808	\N	1	1876	\N	\N	2018-10-17 00:00:00+07	1876.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5486	\N	1	500	\N	\N	2019-07-09 00:00:00+07	0500.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4277	\N	1	602	\N	\N	2017-11-22 00:00:00+07	0602.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4807	\N	1	1875	\N	\N	2018-10-17 00:00:00+07	1875.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4806	\N	1	1874	\N	\N	2018-10-17 00:00:00+07	1874.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5488	\N	1	501	\N	\N	2019-07-10 00:00:00+07	0501.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4271	\N	1	599	\N	\N	2017-11-09 00:00:00+07	0599.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4263	\N	1	1672	\N	\N	2017-11-06 00:00:00+07	1672.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4254	\N	1	1669	\N	\N	2017-11-01 00:00:00+07	1669.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4805	\N	1	1873	\N	\N	2018-10-17 00:00:00+07	1873.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4804	\N	1	1871	\N	\N	2018-10-16 00:00:00+07	1871.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4243	\N	1	755	\N	\N	2014-12-01 00:00:00+07	0755.65.221	2	2017-01-30 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4803	\N	1	1870	\N	\N	2018-10-16 00:00:00+07	1870.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
692	\N	2	735	\N	\N	2014-03-26 00:00:00+07	0735.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6105	\N	1	2400	\N	\N	2020-06-25 00:00:00+07	2400.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4214	\N	1	1667	\N	\N	2017-10-23 00:00:00+07	1667.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
693	\N	2	738	\N	\N	2014-10-22 00:00:00+07	0738.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
694	\N	2	739	\N	\N	2014-04-10 00:00:00+07	0739.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6098	\N	1	1021	\N	\N	2020-06-12 00:00:00+07	1021.61.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4211	\N	1	1666	\N	\N	2017-10-23 00:00:00+07	1666.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4802	\N	1	1869	\N	\N	2018-10-16 00:00:00+07	1869.65.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4188	\N	1	1139	\N	\N	2016-02-01 00:00:00+07	1139.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4164	\N	1	1544	\N	\N	2017-10-11 00:00:00+07	1544.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5504	\N	1	502	\N	\N	2019-07-16 00:00:00+07	0502.63.221	1	2019-08-21 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5506	\N	1	2199	\N	\N	2019-07-16 00:00:00+07	2199.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5507	\N	1	2200	\N	\N	2019-07-16 00:00:00+07	2200.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4152	\N	1	1663	\N	\N	2017-10-09 00:00:00+07	1663.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5513	\N	1	2435	\N	\N	2019-07-17 00:00:00+07	2435.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4135	\N	1	1662	\N	\N	2017-09-29 00:00:00+07	1662.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
695	\N	2	740	\N	\N	2014-04-04 00:00:00+07	0740.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4134	\N	1	586	\N	\N	2017-09-28 00:00:00+07	0586.64.221	1	\N	2020-04-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6060	\N	1	2604	\N	\N	2020-03-12 00:00:00+07	2604.62.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
696	\N	2	741	\N	\N	2014-04-14 00:00:00+07	0741.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
697	\N	2	742	\N	\N	2014-04-16 00:00:00+07	0742.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
698	\N	2	743	\N	\N	2014-04-22 00:00:00+07	0743.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4133	\N	1	585	\N	\N	2017-09-27 00:00:00+07	0585.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4132	\N	1	1	\N	\N	2017-09-27 00:00:00+07	001.64.221	2	2019-08-29 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4072	\N	1	582	\N	\N	2017-09-20 00:00:00+07	0582.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4071	\N	1	1660	\N	\N	2017-09-20 00:00:00+07	1660.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4054	\N	1	581	\N	\N	2017-09-20 00:00:00+07	0581.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5528	\N	1	2206	\N	\N	2019-07-22 00:00:00+07	2206.65.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6038	\N	1	2381	\N	\N	2020-02-24 00:00:00+07	2381.65.221	1	\N	\N	SA	\N	10208	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6036	\N	1	2380	\N	\N	2020-02-20 00:00:00+07	2380.65.221	1	\N	\N	SA	\N	10208	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4801	\N	1	1868	\N	\N	2018-10-16 00:00:00+07	1868.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
699	\N	2	744	\N	\N	2014-04-23 00:00:00+07	0744.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4800	\N	1	1867	\N	\N	2018-10-16 00:00:00+07	1867.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
700	\N	2	745	\N	\N	2014-04-23 00:00:00+07	0745.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4053	\N	1	580	\N	\N	2017-09-20 00:00:00+07	0580.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4051	\N	1	1659	\N	\N	2017-09-20 00:00:00+07	1659.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4050	\N	1	1658	\N	\N	2017-09-19 00:00:00+07	1658.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4049	\N	1	1657	\N	\N	2017-09-19 00:00:00+07	1657.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6021	\N	1	2374	\N	\N	2020-01-30 00:00:00+07	2374.65.221	1	\N	\N	SA	\N	10208	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4047	\N	1	578	\N	\N	2017-09-18 00:00:00+07	0578.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4046	\N	1	577	\N	\N	2017-09-18 00:00:00+07	0577.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4044	\N	1	1655	\N	\N	2017-09-18 00:00:00+07	1655.65.221	1	\N	2020-10-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4086	\N	1	583	\N	\N	2017-09-26 00:00:00+07	0583.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4799	\N	1	1866	\N	\N	2018-10-16 00:00:00+07	1866.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5578	\N	1	789	\N	\N	2019-08-05 00:00:00+07	0789.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4798	\N	1	1865	\N	\N	2018-10-16 00:00:00+07	1865.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4001	\N	1	2194	\N	\N	2017-09-13 00:00:00+07	2194.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4797	\N	1	1864	\N	\N	2018-10-16 00:00:00+07	1864.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
701	\N	2	748	\N	\N	2014-04-30 00:00:00+07	0748.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4796	\N	1	1863	\N	\N	2018-10-16 00:00:00+07	1863.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4795	\N	1	719	\N	\N	2018-10-16 00:00:00+07	0719.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5983	\N	1	2363	\N	\N	2020-01-07 00:00:00+07	2363.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5982	\N	1	833	\N	\N	2020-01-07 00:00:00+07	0833.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5979	\N	1	2362	\N	\N	2020-01-06 00:00:00+07	2362.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4794	\N	1	718	\N	\N	2018-10-16 00:00:00+07	0718.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4793	\N	1	717	\N	\N	2018-10-16 00:00:00+07	0717.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5552	\N	1	505	\N	\N	2019-07-25 00:00:00+07	0505.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3890	\N	1	1640	\N	\N	2017-07-19 00:00:00+07	1640.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5971	\N	1	2358	\N	\N	2019-12-27 00:00:00+07	2358.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5553	\N	1	506	\N	\N	2019-07-25 00:00:00+07	0506.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
702	\N	2	749	\N	\N	2014-05-12 00:00:00+07	0749.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5968	\N	1	2357	\N	\N	2019-12-23 00:00:00+07	2357.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3856	\N	1	1636	\N	\N	2017-07-10 00:00:00+07	1636.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
703	\N	2	750	\N	\N	2014-05-16 00:00:00+07	0750.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5559	\N	1	2212	\N	\N	2019-07-29 00:00:00+07	2212.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5565	\N	1	2216	\N	\N	2019-08-01 00:00:00+07	2216.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3146	\N	1	1600	\N	\N	2017-02-01 00:00:00+07	1600.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3147	\N	1	1537	\N	\N	2016-12-01 00:00:00+07	1537.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5955	\N	1	1003	\N	\N	2019-12-10 00:00:00+07	1003.61.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5953	\N	1	2351	\N	\N	2019-12-09 00:00:00+07	2351.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5566	\N	1	2217	\N	\N	2019-08-01 00:00:00+07	2217.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5567	\N	1	2218	\N	\N	2019-08-01 00:00:00+07	2218.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5568	\N	1	507	\N	\N	2019-08-01 00:00:00+07	0507.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
704	\N	2	752	\N	\N	2014-05-21 00:00:00+07	0752.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
976	\N	2	1804	\N	\N	2015-01-14 00:00:00+07	1804.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
977	\N	2	1003	\N	\N	2015-01-15 00:00:00+07	1003.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
979	\N	2	250	\N	\N	2015-01-19 00:00:00+07	0250.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
980	\N	2	251	\N	\N	2015-01-19 00:00:00+07	0251.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
981	\N	2	1006	\N	\N	2015-01-19 00:00:00+07	1006.65.221	2	2016-12-07 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
982	\N	2	1005	\N	\N	2015-01-19 00:00:00+07	1005.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
983	\N	2	1007	\N	\N	2015-01-19 00:00:00+07	1007.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
984	\N	2	1009	\N	\N	2015-01-19 00:00:00+07	1009.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
985	\N	2	1010	\N	\N	2015-01-19 00:00:00+07	1010.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
986	\N	2	1011	\N	\N	2015-01-19 00:00:00+07	1011.65.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-08-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
987	\N	2	1012	\N	\N	2015-01-20 00:00:00+07	1012.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
988	\N	2	1013	\N	\N	2015-01-20 00:00:00+07	1013.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
989	\N	2	401	\N	\N	2015-01-20 00:00:00+07	0401.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
990	\N	2	1015	\N	\N	2015-01-20 00:00:00+07	1015.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
991	\N	2	1004	\N	\N	2015-01-22 00:00:00+07	1004.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
992	\N	2	1019	\N	\N	2015-01-23 00:00:00+07	1019.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
993	\N	2	1020	\N	\N	2015-01-23 00:00:00+07	1020.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
994	\N	2	1021	\N	\N	2015-01-26 00:00:00+07	1021.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
995	\N	2	1022	\N	\N	2015-01-26 00:00:00+07	1022.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
996	\N	2	1023	\N	\N	2015-01-26 00:00:00+07	1023.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
997	\N	2	1026	\N	\N	2015-01-27 00:00:00+07	1026.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
998	\N	2	1025	\N	\N	2015-01-27 00:00:00+07	1025.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
999	\N	2	1024	\N	\N	2015-01-27 00:00:00+07	1024.65.221	2	2015-03-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1000	\N	2	1027	\N	\N	2015-01-27 00:00:00+07	1027.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1001	\N	2	1028	\N	\N	2015-01-27 00:00:00+07	1028.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1002	\N	2	403	\N	\N	2015-01-27 00:00:00+07	0403.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1003	\N	2	1029	\N	\N	2015-01-27 00:00:00+07	1029.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1004	\N	2	1031	\N	\N	2015-01-27 00:00:00+07	1031.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1005	\N	2	1030	\N	\N	2015-01-27 00:00:00+07	1030.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1006	\N	2	1036	\N	\N	2015-01-28 00:00:00+07	1036.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1007	\N	2	1037	\N	\N	2015-01-28 00:00:00+07	1037.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1008	\N	2	1032	\N	\N	2015-01-29 00:00:00+07	1032.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1009	\N	2	1038	\N	\N	2015-01-29 00:00:00+07	1038.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1010	\N	2	1041	\N	\N	2015-02-02 00:00:00+07	1041.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1011	\N	2	1043	\N	\N	2015-02-02 00:00:00+07	1043.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1012	\N	2	1046	\N	\N	2015-02-02 00:00:00+07	1046.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1013	\N	2	1048	\N	\N	2015-02-02 00:00:00+07	1048.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1014	\N	2	1049	\N	\N	2015-02-02 00:00:00+07	1049.65.221	2	2016-12-06 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1016	\N	2	1051	\N	\N	2015-02-02 00:00:00+07	1051.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1017	\N	2	1035	\N	\N	2015-02-03 00:00:00+07	1035.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4522	\N	1	1726	\N	\N	2018-04-18 00:00:00+07	1726.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1018	\N	2	1039	\N	\N	2015-02-03 00:00:00+07	1039.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1019	\N	2	961	\N	\N	2015-02-03 00:00:00+07	0961.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1021	\N	2	253	\N	\N	2015-02-04 00:00:00+07	0253.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1022	\N	2	1055	\N	\N	2015-02-04 00:00:00+07	1055.65.221	1	\N	2018-09-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1023	\N	2	1016	\N	\N	2015-02-04 00:00:00+07	1016.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1024	\N	2	1056	\N	\N	2015-02-04 00:00:00+07	1056.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1025	\N	2	1054	\N	\N	2015-02-04 00:00:00+07	1054.65.221	2	2017-03-17 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1026	\N	2	1057	\N	\N	2015-02-04 00:00:00+07	1057.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1027	\N	2	1058	\N	\N	2015-02-04 00:00:00+07	1058.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1028	\N	2	1059	\N	\N	2015-02-04 00:00:00+07	1059.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1029	\N	2	1060	\N	\N	2015-02-04 00:00:00+07	1060.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1030	\N	2	1061	\N	\N	2015-02-04 00:00:00+07	1061.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1031	\N	2	1062	\N	\N	2015-02-05 00:00:00+07	1062.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1032	\N	2	1063	\N	\N	2015-02-05 00:00:00+07	1063.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1033	\N	2	1064	\N	\N	2015-02-05 00:00:00+07	1064.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1034	\N	2	1065	\N	\N	2015-02-06 00:00:00+07	1065.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1035	\N	2	1068	\N	\N	2015-02-09 00:00:00+07	1068.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1036	\N	2	1069	\N	\N	2015-02-09 00:00:00+07	1069.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1037	\N	2	1070	\N	\N	2015-02-09 00:00:00+07	1070.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1038	\N	2	1073	\N	\N	2015-02-10 00:00:00+07	1073.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1039	\N	2	1017	\N	\N	2015-02-10 00:00:00+07	1017.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1040	\N	2	1072	\N	\N	2015-02-10 00:00:00+07	1072.65.221	1	\N	2019-11-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1041	\N	2	1071	\N	\N	2015-02-10 00:00:00+07	1071.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1042	\N	2	1074	\N	\N	2015-02-10 00:00:00+07	1074.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1043	\N	2	1075	\N	\N	2015-02-10 00:00:00+07	1075.65.221	2	2018-10-29 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1044	\N	2	405	\N	\N	2015-02-10 00:00:00+07	0405.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1045	\N	2	1077	\N	\N	2015-02-10 00:00:00+07	1077.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1047	\N	2	1078	\N	\N	2015-02-11 00:00:00+07	1078.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1048	\N	2	1034	\N	\N	2015-02-11 00:00:00+07	1034.65.221	1	\N	2020-05-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1050	\N	2	1080	\N	\N	2015-02-11 00:00:00+07	1080.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1051	\N	2	1081	\N	\N	2015-02-11 00:00:00+07	1081.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1052	\N	2	1082	\N	\N	2015-02-11 00:00:00+07	1082.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1053	\N	2	1083	\N	\N	2015-02-11 00:00:00+07	1083.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1054	\N	2	1084	\N	\N	2015-02-11 00:00:00+07	1084.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1055	\N	2	1085	\N	\N	2015-02-12 00:00:00+07	1085.65.221	2	2018-12-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1056	\N	2	1086	\N	\N	2015-02-12 00:00:00+07	1086.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1057	\N	2	1087	\N	\N	2015-02-12 00:00:00+07	1087.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1058	\N	2	1088	\N	\N	2015-02-12 00:00:00+07	1088.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1059	\N	2	1053	\N	\N	2015-02-13 00:00:00+07	1053.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1060	\N	2	1089	\N	\N	2015-02-13 00:00:00+07	1089.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1061	\N	2	1090	\N	\N	2015-02-13 00:00:00+07	1090.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1062	\N	2	1091	\N	\N	2015-02-13 00:00:00+07	1091.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1063	\N	2	1092	\N	\N	2015-02-16 00:00:00+07	1092.65.221	1	2017-08-08 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1064	\N	2	1093	\N	\N	2015-02-16 00:00:00+07	1093.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1065	\N	2	1094	\N	\N	2015-02-16 00:00:00+07	1094.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1066	\N	2	1095	\N	\N	2015-02-16 00:00:00+07	1095.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1067	\N	2	1096	\N	\N	2015-02-16 00:00:00+07	1096.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1068	\N	2	1097	\N	\N	2015-02-17 00:00:00+07	1097.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1069	\N	2	1098	\N	\N	2015-02-18 00:00:00+07	1098.65.221	1	\N	2019-11-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1070	\N	2	1099	\N	\N	2015-02-18 00:00:00+07	1099.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1071	\N	2	998	\N	\N	2015-02-18 00:00:00+07	0998.65.221	1	2017-08-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1072	\N	2	989	\N	\N	2015-02-18 00:00:00+07	0989.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1073	\N	2	1101	\N	\N	2015-02-20 00:00:00+07	1101.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1074	\N	2	1042	\N	\N	2015-02-20 00:00:00+07	1042.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1075	\N	2	1103	\N	\N	2015-02-20 00:00:00+07	1103.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1076	\N	2	1047	\N	\N	2015-02-24 00:00:00+07	1047.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1077	\N	2	1104	\N	\N	2015-02-24 00:00:00+07	1104.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1078	\N	2	1105	\N	\N	2015-02-25 00:00:00+07	1105.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1079	\N	2	1100	\N	\N	2015-02-25 00:00:00+07	1100.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1081	\N	2	1107	\N	\N	2015-02-26 00:00:00+07	1107.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1082	\N	2	1106	\N	\N	2015-02-26 00:00:00+07	1106.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1083	\N	2	1108	\N	\N	2015-02-26 00:00:00+07	1108.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1084	\N	2	409	\N	\N	2015-02-27 00:00:00+07	0409.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1085	\N	2	1033	\N	\N	2015-02-27 00:00:00+07	1033.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1086	\N	2	1109	\N	\N	2015-02-27 00:00:00+07	1109.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1087	\N	2	1110	\N	\N	2015-03-02 00:00:00+07	1110.65.221	1	\N	2020-08-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1088	\N	2	1112	\N	\N	2015-03-02 00:00:00+07	1112.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1089	\N	2	1111	\N	\N	2015-03-02 00:00:00+07	1111.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1090	\N	2	1115	\N	\N	2015-03-02 00:00:00+07	1115.65.221	1	\N	2018-09-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1091	\N	2	1114	\N	\N	2015-03-02 00:00:00+07	1114.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1092	\N	2	1116	\N	\N	2015-03-03 00:00:00+07	1116.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1093	\N	2	1118	\N	\N	2015-03-03 00:00:00+07	1118.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1094	\N	2	1117	\N	\N	2015-03-03 00:00:00+07	1117.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1095	\N	2	1119	\N	\N	2015-03-03 00:00:00+07	1119.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1096	\N	2	1120	\N	\N	2015-03-03 00:00:00+07	1120.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1097	\N	2	1121	\N	\N	2015-03-03 00:00:00+07	1121.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1098	\N	2	1828	\N	\N	2015-03-04 00:00:00+07	1828.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1099	\N	2	1018	\N	\N	2015-03-04 00:00:00+07	1018.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1100	\N	2	1124	\N	\N	2015-03-04 00:00:00+07	1124.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1101	\N	2	1125	\N	\N	2015-03-04 00:00:00+07	1125.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1102	\N	2	1126	\N	\N	2015-03-04 00:00:00+07	1126.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1103	\N	2	1127	\N	\N	2015-03-04 00:00:00+07	1127.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1104	\N	2	1128	\N	\N	2015-03-04 00:00:00+07	1128.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1105	\N	2	1129	\N	\N	2015-03-05 00:00:00+07	1129.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1106	\N	2	1130	\N	\N	2015-03-05 00:00:00+07	1130.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1109	\N	2	1132	\N	\N	2015-03-06 00:00:00+07	1132.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1110	\N	2	1133	\N	\N	2015-03-09 00:00:00+07	1133.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1111	\N	2	1134	\N	\N	2015-03-10 00:00:00+07	1134.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1112	\N	2	1136	\N	\N	2015-03-10 00:00:00+07	1136.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1113	\N	2	1137	\N	\N	2015-03-10 00:00:00+07	1137.65.221	2	2015-06-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1114	\N	2	1139	\N	\N	2015-03-11 00:00:00+07	1139.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1115	\N	2	1140	\N	\N	2015-03-11 00:00:00+07	1140.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1116	\N	2	1141	\N	\N	2015-03-11 00:00:00+07	1141.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1117	\N	2	1142	\N	\N	2015-03-11 00:00:00+07	1142.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1119	\N	2	1143	\N	\N	2015-03-11 00:00:00+07	1143.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1120	\N	2	1135	\N	\N	2015-03-12 00:00:00+07	1135.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1121	\N	2	1835	\N	\N	2015-03-12 00:00:00+07	1835.62.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1122	\N	2	1144	\N	\N	2015-03-12 00:00:00+07	1144.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1123	\N	2	1145	\N	\N	2015-03-12 00:00:00+07	1145.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1124	\N	2	1146	\N	\N	2015-03-13 00:00:00+07	1146.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1125	\N	2	1147	\N	\N	2015-03-16 00:00:00+07	1147.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1126	\N	2	1148	\N	\N	2015-03-17 00:00:00+07	1148.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1127	\N	2	1149	\N	\N	2015-03-17 00:00:00+07	1149.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1128	\N	2	1151	\N	\N	2015-03-17 00:00:00+07	1151.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1129	\N	2	1150	\N	\N	2015-03-17 00:00:00+07	1150.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1130	\N	2	1841	\N	\N	2015-03-18 00:00:00+07	1841.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1131	\N	2	1153	\N	\N	2015-03-18 00:00:00+07	1153.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1132	\N	2	1154	\N	\N	2015-03-18 00:00:00+07	1154.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1133	\N	2	1155	\N	\N	2015-03-18 00:00:00+07	1155.65.221	1	\N	2018-10-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1134	\N	2	1158	\N	\N	2015-03-20 00:00:00+07	1158.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1135	\N	2	1152	\N	\N	2015-03-20 00:00:00+07	1152.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1137	\N	2	412	\N	\N	2015-03-24 00:00:00+07	0412.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1138	\N	2	1113	\N	\N	2015-03-24 00:00:00+07	1113.65.221	2	2017-03-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1139	\N	2	1842	\N	\N	2015-03-24 00:00:00+07	1842.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1140	\N	2	1160	\N	\N	2015-03-26 00:00:00+07	1160.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1141	\N	2	1122	\N	\N	2015-03-26 00:00:00+07	1122.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1142	\N	2	1123	\N	\N	2015-03-26 00:00:00+07	1123.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1143	\N	2	1162	\N	\N	2015-03-27 00:00:00+07	1162.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1144	\N	2	1163	\N	\N	2015-03-30 00:00:00+07	1163.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1145	\N	2	1164	\N	\N	2015-04-01 00:00:00+07	1164.65.221	1	2017-06-08 00:00:00+07	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1146	\N	2	1166	\N	\N	2015-04-02 00:00:00+07	1166.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1147	\N	2	1168	\N	\N	2015-04-06 00:00:00+07	1168.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1148	\N	2	1169	\N	\N	2015-04-07 00:00:00+07	1169.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1149	\N	2	1171	\N	\N	2015-04-07 00:00:00+07	1171.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1150	\N	2	1170	\N	\N	2015-04-08 00:00:00+07	1170.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3970	\N	1	1649	\N	\N	2017-08-31 00:00:00+07	1649.65.221	1	\N	2019-08-02 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1151	\N	2	411	\N	\N	2015-04-08 00:00:00+07	0411.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1152	\N	2	1067	\N	\N	2015-04-08 00:00:00+07	1067.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1153	\N	2	1172	\N	\N	2015-04-09 00:00:00+07	1172.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1154	\N	2	1173	\N	\N	2015-04-10 00:00:00+07	1173.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1155	\N	2	1176	\N	\N	2015-04-14 00:00:00+07	1176.65.221	2	2017-08-08 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1156	\N	2	1174	\N	\N	2015-04-14 00:00:00+07	1174.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1157	\N	2	1177	\N	\N	2015-04-15 00:00:00+07	1177.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1158	\N	2	1157	\N	\N	2015-04-15 00:00:00+07	1157.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1159	\N	2	1175	\N	\N	2015-04-15 00:00:00+07	1175.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1160	\N	2	1179	\N	\N	2015-04-15 00:00:00+07	1179.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1161	\N	2	1178	\N	\N	2015-04-15 00:00:00+07	1178.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2943	\N	1	1571	\N	\N	2017-02-01 00:00:00+07	1571.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1163	\N	2	1180	\N	\N	2015-04-16 00:00:00+07	1180.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1164	\N	2	1181	\N	\N	2015-04-16 00:00:00+07	1181.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1165	\N	2	1182	\N	\N	2015-04-17 00:00:00+07	1182.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1166	\N	2	1183	\N	\N	2015-04-20 00:00:00+07	1183.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1167	\N	2	1184	\N	\N	2015-04-20 00:00:00+07	1184.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1168	\N	2	1185	\N	\N	2015-04-20 00:00:00+07	1185.65.221	2	2017-01-01 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1169	\N	2	1186	\N	\N	2015-04-20 00:00:00+07	1186.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1170	\N	2	1187	\N	\N	2015-04-20 00:00:00+07	1187.65.221	1	2017-06-07 00:00:00+07	2019-10-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1171	\N	2	1189	\N	\N	2015-04-20 00:00:00+07	1189.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1172	\N	2	1190	\N	\N	2015-04-20 00:00:00+07	1190.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1173	\N	2	1191	\N	\N	2015-04-20 00:00:00+07	1191.65.221	1	\N	2020-09-15 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1174	\N	2	1192	\N	\N	2015-04-20 00:00:00+07	1192.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1175	\N	2	1188	\N	\N	2015-04-20 00:00:00+07	1188.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1176	\N	2	1193	\N	\N	2015-04-20 00:00:00+07	1193.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1177	\N	2	1194	\N	\N	2015-04-20 00:00:00+07	1194.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1178	\N	2	1195	\N	\N	2015-04-20 00:00:00+07	1195.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1179	\N	2	1197	\N	\N	2015-04-20 00:00:00+07	1197.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1180	\N	2	1196	\N	\N	2015-04-20 00:00:00+07	1196.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1181	\N	2	1198	\N	\N	2015-04-21 00:00:00+07	1198.65.221	1	2017-05-08 00:00:00+07	2018-10-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1182	\N	2	1199	\N	\N	2015-04-21 00:00:00+07	1199.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1183	\N	2	1200	\N	\N	2015-04-21 00:00:00+07	1200.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1184	\N	2	1201	\N	\N	2015-04-21 00:00:00+07	1201.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1185	\N	2	1202	\N	\N	2015-04-21 00:00:00+07	1202.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1186	\N	2	416	\N	\N	2015-04-21 00:00:00+07	0416.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1187	\N	2	1203	\N	\N	2015-04-21 00:00:00+07	1203.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1188	\N	2	1204	\N	\N	2015-04-21 00:00:00+07	1204.65.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1189	\N	2	1156	\N	\N	2015-04-22 00:00:00+07	1156.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1190	\N	2	1205	\N	\N	2015-04-22 00:00:00+07	1205.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1191	\N	2	1206	\N	\N	2015-04-23 00:00:00+07	1206.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1192	\N	2	1207	\N	\N	2015-04-23 00:00:00+07	1207.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1193	\N	2	1208	\N	\N	2015-04-24 00:00:00+07	1208.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1194	\N	2	727	\N	\N	2015-04-28 00:00:00+07	0727.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1195	\N	2	1219	\N	\N	2015-05-04 00:00:00+07	1219.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1196	\N	2	1220	\N	\N	2015-05-04 00:00:00+07	1220.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1197	\N	2	1221	\N	\N	2015-05-04 00:00:00+07	1221.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1198	\N	2	1222	\N	\N	2015-05-04 00:00:00+07	1222.65.221	1	2018-06-01 00:00:00+07	2019-09-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1199	\N	2	1214	\N	\N	2015-05-04 00:00:00+07	1214.65.221	1	\N	\N	SA	\N	10208	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1200	\N	2	1223	\N	\N	2015-05-04 00:00:00+07	1223.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1201	\N	2	1855	\N	\N	2015-05-04 00:00:00+07	1855.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1202	\N	2	1224	\N	\N	2015-05-04 00:00:00+07	1224.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1203	\N	2	1226	\N	\N	2015-05-04 00:00:00+07	1226.65.221	3	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1204	\N	2	1227	\N	\N	2015-05-05 00:00:00+07	1227.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1205	\N	2	1230	\N	\N	2015-05-06 00:00:00+07	1230.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1206	\N	2	1856	\N	\N	2015-05-06 00:00:00+07	1856.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1207	\N	2	1232	\N	\N	2015-05-06 00:00:00+07	1232.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1208	\N	2	1235	\N	\N	2015-05-06 00:00:00+07	1235.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4519	\N	1	1723	\N	\N	2018-04-18 00:00:00+07	1723.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1209	\N	2	419	\N	\N	2015-05-07 00:00:00+07	0419.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1210	\N	2	1237	\N	\N	2015-05-07 00:00:00+07	1237.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1211	\N	2	420	\N	\N	2015-05-07 00:00:00+07	0420.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1212	\N	2	1239	\N	\N	2015-05-08 00:00:00+07	1239.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1213	\N	2	1240	\N	\N	2015-05-11 00:00:00+07	1240.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1214	\N	2	987	\N	\N	2015-05-12 00:00:00+07	0987.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1215	\N	2	1242	\N	\N	2015-05-13 00:00:00+07	1242.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1216	\N	2	1243	\N	\N	2015-05-13 00:00:00+07	1243.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1217	\N	2	1245	\N	\N	2015-05-15 00:00:00+07	1245.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1218	\N	2	278	\N	\N	2015-05-19 00:00:00+07	0278.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1219	\N	2	430	\N	\N	2015-05-20 00:00:00+07	0430.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1220	\N	2	1253	\N	\N	2015-05-20 00:00:00+07	1253.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1221	\N	2	1866	\N	\N	2015-05-20 00:00:00+07	1866.62.221	1	2019-02-01 00:00:00+07	\N	SA	\N	10208	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1222	\N	2	1878	\N	\N	2015-05-27 00:00:00+07	1878.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1223	\N	2	1258	\N	\N	2015-05-28 00:00:00+07	1258.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1224	\N	2	432	\N	\N	2015-05-28 00:00:00+07	0432.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1225	\N	2	1259	\N	\N	2015-05-28 00:00:00+07	1259.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1226	\N	2	1881	\N	\N	2015-05-29 00:00:00+07	1881.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1227	\N	2	1260	\N	\N	2015-06-01 00:00:00+07	1260.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1228	\N	2	1263	\N	\N	2015-06-05 00:00:00+07	1263.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1229	\N	2	1887	\N	\N	2015-06-05 00:00:00+07	1887.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1230	\N	2	433	\N	\N	2015-06-08 00:00:00+07	0433.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1231	\N	2	1888	\N	\N	2015-06-08 00:00:00+07	1888.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1232	\N	2	1264	\N	\N	2015-06-09 00:00:00+07	1264.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1233	\N	2	1265	\N	\N	2015-06-10 00:00:00+07	1265.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1234	\N	2	1266	\N	\N	2015-06-11 00:00:00+07	1266.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1235	\N	2	1268	\N	\N	2015-06-11 00:00:00+07	1268.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3973	\N	1	1650	\N	\N	2017-09-05 00:00:00+07	1650.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1236	\N	2	1269	\N	\N	2015-06-15 00:00:00+07	1269.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1237	\N	2	1270	\N	\N	2015-06-15 00:00:00+07	1270.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1238	\N	2	1271	\N	\N	2015-06-15 00:00:00+07	1271....65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1239	\N	2	1273	\N	\N	2015-06-16 00:00:00+07	1273.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1240	\N	2	1274	\N	\N	2015-06-17 00:00:00+07	1274.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3974	\N	1	1651	\N	\N	2017-09-05 00:00:00+07	1651.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1241	\N	2	1275	\N	\N	2015-06-17 00:00:00+07	1275.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1242	\N	2	1276	\N	\N	2015-06-17 00:00:00+07	1276.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1243	\N	2	1277	\N	\N	2015-06-17 00:00:00+07	1277.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1244	\N	2	1278	\N	\N	2015-06-17 00:00:00+07	1278.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1245	\N	2	1280	\N	\N	2015-06-18 00:00:00+07	1280.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1246	\N	2	434	\N	\N	2015-06-22 00:00:00+07	0434.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1247	\N	2	435	\N	\N	2015-06-26 00:00:00+07	0435.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1248	\N	2	1282	\N	\N	2015-06-30 00:00:00+07	1282.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1249	\N	2	1283	\N	\N	2015-07-03 00:00:00+07	1283.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1250	\N	2	1898	\N	\N	2015-07-08 00:00:00+07	1898.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1251	\N	2	1284	\N	\N	2015-07-09 00:00:00+07	1284.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1252	\N	2	788	\N	\N	2015-07-13 00:00:00+07	0788.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4595	\N	1	1755	\N	\N	2018-07-09 00:00:00+07	1755.65.221	1	\N	2018-08-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
791	\N	2	857	\N	\N	2014-10-09 00:00:00+07	0857.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
792	\N	2	858	\N	\N	2014-10-09 00:00:00+07	0858.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
793	\N	2	860	\N	\N	2014-10-09 00:00:00+07	0860.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
794	\N	2	861	\N	\N	2014-10-09 00:00:00+07	0861.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4589	\N	1	674	\N	\N	2018-07-05 00:00:00+07	0674.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
795	\N	2	862	\N	\N	2014-10-13 00:00:00+07	0862.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3946	\N	1	1642	\N	\N	2017-08-23 00:00:00+07	1642.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4588	\N	1	1751	\N	\N	2018-07-03 00:00:00+07	1751.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4581	\N	1	1749	\N	\N	2018-06-05 00:00:00+07	1749.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
796	\N	2	863	\N	\N	2014-10-13 00:00:00+07	0863.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
797	\N	2	864	\N	\N	2014-10-13 00:00:00+07	0864.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
798	\N	2	865	\N	\N	2014-10-14 00:00:00+07	0865.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
799	\N	2	866	\N	\N	2014-10-15 00:00:00+07	0866.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4579	\N	1	1748	\N	\N	2018-06-04 00:00:00+07	1748.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
800	\N	2	867	\N	\N	2014-10-15 00:00:00+07	0867.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
801	\N	2	868	\N	\N	2014-10-16 00:00:00+07	0868.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
802	\N	2	869	\N	\N	2014-10-16 00:00:00+07	0869.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4578	\N	1	1747	\N	\N	2018-06-04 00:00:00+07	1747.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4577	\N	1	1746	\N	\N	2018-06-04 00:00:00+07	1746.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
803	\N	2	870	\N	\N	2014-10-16 00:00:00+07	0870.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
804	\N	2	871	\N	\N	2014-10-16 00:00:00+07	0871.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
557	\N	2	239	\N	\N	2014-03-12 00:00:00+07	0239.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
558	\N	2	246	\N	\N	2014-11-03 00:00:00+07	0246.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
805	\N	2	872	\N	\N	2014-10-16 00:00:00+07	0872.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
806	\N	2	873	\N	\N	2014-10-16 00:00:00+07	0873.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
807	\N	2	875	\N	\N	2014-10-21 00:00:00+07	0875.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
808	\N	2	876	\N	\N	2014-10-22 00:00:00+07	0876.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
809	\N	2	877	\N	\N	2014-10-22 00:00:00+07	0877.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
810	\N	2	878	\N	\N	2014-10-23 00:00:00+07	0878.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4560	\N	1	1742	\N	\N	2018-05-14 00:00:00+07	1742.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4559	\N	1	2259	\N	\N	2018-05-14 00:00:00+07	2259.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4558	\N	1	1741	\N	\N	2018-05-14 00:00:00+07	1741.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4557	\N	1	2258	\N	\N	2018-05-14 00:00:00+07	2258.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4556	\N	1	1740	\N	\N	2018-05-09 00:00:00+07	1740.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4555	\N	1	667	\N	\N	2018-05-08 00:00:00+07	0667.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4553	\N	1	1738	\N	\N	2018-05-07 00:00:00+07	1738.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4549	\N	1	418	\N	\N	2018-05-07 00:00:00+07	0418.63.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
811	\N	2	879	\N	\N	2014-10-23 00:00:00+07	0879.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
812	\N	2	880	\N	\N	2014-10-24 00:00:00+07	0880.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4548	\N	1	417	\N	\N	2018-05-07 00:00:00+07	0417.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4547	\N	1	416	\N	\N	2018-05-07 00:00:00+07	0416.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
813	\N	2	881	\N	\N	2014-10-24 00:00:00+07	0881.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4546	\N	1	415	\N	\N	2018-05-07 00:00:00+07	0415.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
814	\N	2	882	\N	\N	2014-10-27 00:00:00+07	0882.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4545	\N	1	414	\N	\N	2018-05-07 00:00:00+07	0414.63.221	1	\N	2020-09-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
815	\N	2	883	\N	\N	2014-10-27 00:00:00+07	0883.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
816	\N	2	884	\N	\N	2014-11-12 00:00:00+07	0884.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4544	\N	1	1737	\N	\N	2018-05-04 00:00:00+07	1737.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
817	\N	2	885	\N	\N	2014-10-27 00:00:00+07	0885.65.221	1	\N	2019-10-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4541	\N	1	666	\N	\N	2018-05-03 00:00:00+07	0666.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
818	\N	2	886	\N	\N	2014-10-28 00:00:00+07	0886.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
819	\N	2	887	\N	\N	2014-10-28 00:00:00+07	0887.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4540	\N	1	413	\N	\N	2018-05-02 00:00:00+07	0413.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
820	\N	2	888	\N	\N	2014-10-28 00:00:00+07	0888.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5413	\N	1	2157	\N	\N	2019-06-17 00:00:00+07	2157.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3949	\N	1	1644	\N	\N	2017-08-24 00:00:00+07	1644.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
821	\N	2	889	\N	\N	2014-10-28 00:00:00+07	0889.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
822	\N	2	890	\N	\N	2014-10-28 00:00:00+07	0890.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6254	\N	1	2446	\N	\N	2020-10-21 00:00:00+07	2446.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2020-10-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
823	\N	2	891	\N	\N	2014-10-29 00:00:00+07	0891.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4535	\N	1	2257	\N	\N	2018-04-26 00:00:00+07	2257.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4533	\N	1	1733	\N	\N	2018-04-26 00:00:00+07	1733.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
824	\N	2	892	\N	\N	2014-10-29 00:00:00+07	0892.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
826	\N	2	893	\N	\N	2014-10-30 00:00:00+07	0893.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
827	\N	2	894	\N	\N	2014-10-30 00:00:00+07	0894.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
829	\N	2	895	\N	\N	2014-10-31 00:00:00+07	0895.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
830	\N	2	896	\N	\N	2014-10-31 00:00:00+07	0896.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
831	\N	2	898	\N	\N	2014-11-03 00:00:00+07	0898.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
832	\N	2	901	\N	\N	2014-11-03 00:00:00+07	0901.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
833	\N	2	902	\N	\N	2014-11-20 00:00:00+07	0902.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
834	\N	2	903	\N	\N	2014-11-04 00:00:00+07	0903.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
835	\N	2	904	\N	\N	2014-11-04 00:00:00+07	0904.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
836	\N	2	906	\N	\N	2014-11-17 00:00:00+07	0906.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
838	\N	2	907	\N	\N	2014-11-04 00:00:00+07	0907.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
839	\N	2	908	\N	\N	2014-11-06 00:00:00+07	0908.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
840	\N	2	909	\N	\N	2014-11-05 00:00:00+07	0909.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
841	\N	2	910	\N	\N	2014-11-05 00:00:00+07	0910.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
842	\N	2	911	\N	\N	2014-11-06 00:00:00+07	0911.65.221	2	2019-10-01 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
843	\N	2	912	\N	\N	2014-11-11 00:00:00+07	0912.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
844	\N	2	913	\N	\N	2014-11-11 00:00:00+07	0913.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
845	\N	2	914	\N	\N	2014-11-12 00:00:00+07	0914.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
846	\N	2	916	\N	\N	2014-11-13 00:00:00+07	0916.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
847	\N	2	917	\N	\N	2014-11-13 00:00:00+07	0917.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
848	\N	2	918	\N	\N	2014-11-13 00:00:00+07	0918.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
849	\N	2	919	\N	\N	2014-11-13 00:00:00+07	0919.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
850	\N	2	920	\N	\N	2014-11-14 00:00:00+07	0920.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
851	\N	2	921	\N	\N	2014-11-14 00:00:00+07	0921.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
852	\N	2	922	\N	\N	2014-11-17 00:00:00+07	0922.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
853	\N	2	923	\N	\N	2014-11-17 00:00:00+07	0923.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
854	\N	2	924	\N	\N	2014-11-17 00:00:00+07	0924.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
855	\N	2	929	\N	\N	2014-11-24 00:00:00+07	0929.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
856	\N	2	930	\N	\N	2014-11-24 00:00:00+07	0930.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
857	\N	2	931	\N	\N	2014-12-18 00:00:00+07	0931.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
858	\N	2	931	\N	\N	2014-11-24 00:00:00+07	0931.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
859	\N	2	932	\N	\N	2014-11-24 00:00:00+07	0932.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
860	\N	2	933	\N	\N	2014-11-25 00:00:00+07	0933.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
861	\N	2	934	\N	\N	2014-11-25 00:00:00+07	0934.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
862	\N	2	935	\N	\N	2014-11-25 00:00:00+07	0935.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
863	\N	2	936	\N	\N	2014-11-26 00:00:00+07	0936.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
864	\N	2	937	\N	\N	2014-11-26 00:00:00+07	0937.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
865	\N	2	938	\N	\N	2014-11-26 00:00:00+07	0938.65.221	1	\N	2019-08-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
866	\N	2	939	\N	\N	2014-11-27 00:00:00+07	0939.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
867	\N	2	940	\N	\N	2014-11-27 00:00:00+07	0940.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
868	\N	2	941	\N	\N	2014-11-27 00:00:00+07	0941.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
869	\N	2	942	\N	\N	2014-11-27 00:00:00+07	0942.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
870	\N	2	943	\N	\N	2014-11-27 00:00:00+07	0943.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
871	\N	2	944	\N	\N	2014-11-27 00:00:00+07	0944.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
872	\N	2	945	\N	\N	2014-11-27 00:00:00+07	0945.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-03-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
873	\N	2	946	\N	\N	2014-11-27 00:00:00+07	0946.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
874	\N	2	947	\N	\N	2014-11-27 00:00:00+07	0947.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
875	\N	2	948	\N	\N	2014-11-27 00:00:00+07	0948.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
876	\N	2	949	\N	\N	2014-11-28 00:00:00+07	0949.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
877	\N	2	950	\N	\N	2014-12-02 00:00:00+07	0950.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
878	\N	2	951	\N	\N	2014-11-28 00:00:00+07	0951.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
879	\N	2	952	\N	\N	2014-11-28 00:00:00+07	0952.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
880	\N	2	953	\N	\N	2014-11-28 00:00:00+07	0953.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
881	\N	2	955	\N	\N	2014-12-03 00:00:00+07	0955.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
882	\N	2	956	\N	\N	2014-12-03 00:00:00+07	0956.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
883	\N	2	957	\N	\N	2014-12-03 00:00:00+07	0957.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
884	\N	2	958	\N	\N	2014-12-03 00:00:00+07	0958.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
885	\N	2	959	\N	\N	2014-12-04 00:00:00+07	0959.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
886	\N	2	962	\N	\N	2014-12-05 00:00:00+07	0962.65.221	2	2018-05-08 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
887	\N	2	963	\N	\N	2014-12-08 00:00:00+07	0963.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
888	\N	2	964	\N	\N	2014-12-11 00:00:00+07	0964.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
889	\N	2	965	\N	\N	2014-12-11 00:00:00+07	0965.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
890	\N	2	966	\N	\N	2014-12-12 00:00:00+07	0966.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
891	\N	2	968	\N	\N	2014-12-19 00:00:00+07	0968.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
892	\N	2	969	\N	\N	2014-12-15 00:00:00+07	0969.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
893	\N	2	970	\N	\N	2014-12-16 00:00:00+07	0970.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
894	\N	2	971	\N	\N	2014-12-16 00:00:00+07	0971.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
895	\N	2	972	\N	\N	2014-12-18 00:00:00+07	0972.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
896	\N	2	973	\N	\N	2014-12-18 00:00:00+07	0973.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
897	\N	2	974	\N	\N	2014-12-18 00:00:00+07	0974.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
898	\N	2	975	\N	\N	2014-12-18 00:00:00+07	0975.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
899	\N	2	976	\N	\N	2014-12-19 00:00:00+07	0976.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
900	\N	2	978	\N	\N	2014-12-22 00:00:00+07	0978.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
901	\N	2	979	\N	\N	2014-12-22 00:00:00+07	0979.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
902	\N	2	980	\N	\N	2014-12-22 00:00:00+07	0980.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
903	\N	2	981	\N	\N	2014-12-23 00:00:00+07	0981.65.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
904	\N	2	982	\N	\N	2014-12-24 00:00:00+07	0982.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
905	\N	2	983	\N	\N	2014-12-29 00:00:00+07	0983.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
906	\N	2	984	\N	\N	2014-12-29 00:00:00+07	0984.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3961	\N	1	1645	\N	\N	2017-08-28 00:00:00+07	1645.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
910	\N	2	1711	\N	\N	2014-02-11 00:00:00+07	1711.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
911	\N	2	1713	\N	\N	2014-03-17 00:00:00+07	1713.62.221	1	2017-06-09 00:00:00+07	2020-08-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
912	\N	2	1715	\N	\N	2014-02-18 00:00:00+07	1715.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
913	\N	2	1717	\N	\N	2014-03-20 00:00:00+07	1717.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
914	\N	2	1721	\N	\N	2014-03-06 00:00:00+07	1721.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
915	\N	2	1722	\N	\N	2014-03-06 00:00:00+07	1722.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
916	\N	2	1723	\N	\N	2014-03-07 00:00:00+07	1723.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
917	\N	2	1724	\N	\N	2014-03-07 00:00:00+07	1724.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
918	\N	2	1728	\N	\N	2014-03-18 00:00:00+07	1728.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
919	\N	2	1729	\N	\N	2014-03-19 00:00:00+07	1729.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
920	\N	2	1730	\N	\N	2014-03-20 00:00:00+07	1730.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
921	\N	2	1731	\N	\N	2014-03-25 00:00:00+07	1731.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
922	\N	2	1732	\N	\N	2014-03-20 00:00:00+07	1732.62.221	1	\N	2020-01-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
923	\N	2	1733	\N	\N	2014-03-20 00:00:00+07	1733.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
924	\N	2	1734	\N	\N	2014-03-25 00:00:00+07	1734.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
925	\N	2	1736	\N	\N	2014-03-20 00:00:00+07	1736.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
926	\N	2	1737	\N	\N	2014-03-21 00:00:00+07	1737.62.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
927	\N	2	1738	\N	\N	2014-03-24 00:00:00+07	1738.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
928	\N	2	1739	\N	\N	2014-03-24 00:00:00+07	1739.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
929	\N	2	1740	\N	\N	2014-03-24 00:00:00+07	1740.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
930	\N	2	1741	\N	\N	2014-03-24 00:00:00+07	1741.62.221	2	2016-12-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
931	\N	2	1742	\N	\N	2014-03-25 00:00:00+07	1742.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
932	\N	2	1743	\N	\N	2014-03-26 00:00:00+07	1743.62.221	2	2017-08-03 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
933	\N	2	1744	\N	\N	2014-03-27 00:00:00+07	1744.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
934	\N	2	1745	\N	\N	2014-03-27 00:00:00+07	1745.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
935	\N	2	1746	\N	\N	2014-03-28 00:00:00+07	1746.62.221	1	2017-06-06 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
936	\N	2	1747	\N	\N	2014-04-01 00:00:00+07	1747.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
937	\N	2	1748	\N	\N	2014-04-02 00:00:00+07	1748.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
938	\N	2	1749	\N	\N	2014-04-07 00:00:00+07	1749.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
939	\N	2	1753	\N	\N	2014-04-21 00:00:00+07	1753.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
940	\N	2	1754	\N	\N	2014-04-28 00:00:00+07	1754.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
941	\N	2	1755	\N	\N	2014-04-30 00:00:00+07	1755.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
942	\N	2	1757	\N	\N	2014-05-19 00:00:00+07	1757.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
943	\N	2	1761	\N	\N	2014-06-19 00:00:00+07	1761.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
944	\N	2	1769	\N	\N	2014-07-14 00:00:00+07	1769.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
945	\N	2	1776	\N	\N	2014-09-01 00:00:00+07	1776.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
946	\N	2	1778	\N	\N	2014-09-10 00:00:00+07	1778.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
947	\N	2	1781	\N	\N	2014-09-17 00:00:00+07	1781.62.221	2	2021-04-01 00:00:00+07	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-07-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
948	\N	2	1782	\N	\N	2014-09-25 00:00:00+07	1782.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
949	\N	2	1784	\N	\N	2014-10-02 00:00:00+07	1784.62.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
950	\N	2	1786	\N	\N	2014-10-16 00:00:00+07	1786.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
951	\N	2	1787	\N	\N	2014-10-16 00:00:00+07	1787.62.221	1	\N	2020-07-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
952	\N	2	1789	\N	\N	2014-10-16 00:00:00+07	1789.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
953	\N	2	1790	\N	\N	2014-10-23 00:00:00+07	1790.62.221	1	\N	2020-09-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
954	\N	2	1792	\N	\N	2014-10-31 00:00:00+07	1792.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
956	\N	2	1805	\N	\N	2014-12-22 00:00:00+07	1805.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
957	\N	2	1807	\N	\N	2014-12-29 00:00:00+07	1807.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
958	\N	2	399	\N	\N	2015-01-02 00:00:00+07	0399.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
959	\N	2	988	\N	\N	2015-01-05 00:00:00+07	0988.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
960	\N	2	990	\N	\N	2015-01-05 00:00:00+07	0990.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
961	\N	2	721	\N	\N	2015-01-05 00:00:00+07	0721.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
962	\N	2	991	\N	\N	2015-01-05 00:00:00+07	0991.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
963	\N	2	992	\N	\N	2015-01-05 00:00:00+07	0992.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
964	\N	2	996	\N	\N	2015-01-05 00:00:00+07	0996.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
965	\N	2	995	\N	\N	2015-01-05 00:00:00+07	0995.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
966	\N	2	997	\N	\N	2015-01-05 00:00:00+07	0997.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
967	\N	2	999	\N	\N	2015-01-07 00:00:00+07	0999.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
968	\N	2	396	\N	\N	2015-01-07 00:00:00+07	0396.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
969	\N	2	395	\N	\N	2015-01-07 00:00:00+07	0395.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
970	\N	2	394	\N	\N	2015-01-07 00:00:00+07	0394.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
971	\N	2	393	\N	\N	2015-01-07 00:00:00+07	0393.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
972	\N	2	400	\N	\N	2015-01-08 00:00:00+07	0400.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
973	\N	2	967	\N	\N	2015-01-12 00:00:00+07	0967.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
974	\N	2	1001	\N	\N	2015-01-12 00:00:00+07	1001.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5583	\N	1	508	\N	\N	2019-08-08 00:00:00+07	0508.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3827	\N	1	559	\N	\N	2017-06-15 00:00:00+07	0559.64.221	2	2019-07-01 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3826	\N	1	558	\N	\N	2017-06-15 00:00:00+07	0558.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5584	\N	1	509	\N	\N	2019-08-08 00:00:00+07	0509.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5586	\N	1	510	\N	\N	2019-08-08 00:00:00+07	0510.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6721	\N	1	633	\N	\N	2021-09-10 00:00:00+07	0633.63.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-09-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5588	\N	1	2219	\N	\N	2019-08-08 00:00:00+07	2219.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3816	\N	1	1631	\N	\N	2017-06-13 00:00:00+07	1631.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
975	\N	2	986	\N	\N	2015-01-13 00:00:00+07	0986.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3810	\N	1	556	\N	\N	2017-06-13 00:00:00+07	0556.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
705	\N	2	753	\N	\N	2014-05-26 00:00:00+07	0753.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5915	\N	1	2344	\N	\N	2019-11-14 00:00:00+07	2344.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
706	\N	2	754	\N	\N	2014-05-22 00:00:00+07	0754.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
708	\N	2	756	\N	\N	2014-05-22 00:00:00+07	0756.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5907	\N	1	2341	\N	\N	2019-11-13 00:00:00+07	2341.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3804	\N	1	848	\N	\N	2017-06-09 00:00:00+07	0848.61.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5904	\N	1	2340	\N	\N	2019-11-13 00:00:00+07	2340.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5593	\N	1	2221	\N	\N	2019-08-08 00:00:00+07	2221.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5596	\N	1	968	\N	\N	2019-08-09 00:00:00+07	0968.61.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5390	\N	1	2145	\N	\N	2019-06-12 00:00:00+07	2145.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4786	\N	1	715	\N	\N	2018-10-10 00:00:00+07	0715.64.221	1	\N	2020-01-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
709	\N	2	757	\N	\N	2014-05-26 00:00:00+07	0757.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5889	\N	1	2338	\N	\N	2019-11-07 00:00:00+07	2338.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5887	\N	1	2337	\N	\N	2019-11-07 00:00:00+07	2337.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
710	\N	2	758	\N	\N	2014-05-30 00:00:00+07	0758.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3792	\N	1	555	\N	\N	2017-06-07 00:00:00+07	0555.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5880	\N	1	818	\N	\N	2019-11-05 00:00:00+07	0818.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
711	\N	2	759	\N	\N	2014-05-30 00:00:00+07	0759.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5877	\N	1	817	\N	\N	2019-10-31 00:00:00+07	0817.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5876	\N	1	816	\N	\N	2019-10-31 00:00:00+07	0816.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4785	\N	1	714	\N	\N	2018-10-10 00:00:00+07	0714.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
712	\N	2	760	\N	\N	2014-06-04 00:00:00+07	0760.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5872	\N	1	2332	\N	\N	2019-10-31 00:00:00+07	2332.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4779	\N	1	1856	\N	\N	2018-10-03 00:00:00+07	1856.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4776	\N	1	1854	\N	\N	2018-10-02 00:00:00+07	1854.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5612	\N	1	2225	\N	\N	2019-08-12 00:00:00+07	2225.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4775	\N	1	1853	\N	\N	2018-10-02 00:00:00+07	1853.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4774	\N	1	1852	\N	\N	2018-10-02 00:00:00+07	1852.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4773	\N	1	1851	\N	\N	2018-10-02 00:00:00+07	1851.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5864	\N	1	2331	\N	\N	2019-10-29 00:00:00+07	2331.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3424	\N	1	494	\N	\N	2016-11-01 00:00:00+07	0494.64.221	3	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4772	\N	1	1850	\N	\N	2018-10-02 00:00:00+07	1850.65.221	1	\N	2020-08-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5613	\N	1	2226	\N	\N	2019-08-12 00:00:00+07	2226.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4771	\N	1	1849	\N	\N	2018-10-02 00:00:00+07	1849.65.221	2	2020-03-17 00:00:00+07	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5854	\N	1	537	\N	\N	2019-10-24 00:00:00+07	0537.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5853	\N	1	536	\N	\N	2019-10-23 00:00:00+07	0536.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4770	\N	1	1848	\N	\N	2018-10-02 00:00:00+07	1848.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4769	\N	1	1847	\N	\N	2018-10-02 00:00:00+07	1847.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5850	\N	1	2538	\N	\N	2019-10-23 00:00:00+07	2538.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4768	\N	1	1846	\N	\N	2018-10-02 00:00:00+07	1846.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4767	\N	1	1845	\N	\N	2018-10-02 00:00:00+07	1845.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4766	\N	1	1844	\N	\N	2018-10-02 00:00:00+07	1844.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5844	\N	1	535	\N	\N	2019-10-21 00:00:00+07	0535.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4765	\N	1	1843	\N	\N	2018-10-02 00:00:00+07	1843.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4764	\N	1	1842	\N	\N	2018-10-02 00:00:00+07	1842.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4763	\N	1	1840	\N	\N	2018-10-02 00:00:00+07	1840.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4762	\N	1	1839	\N	\N	2018-10-02 00:00:00+07	1839.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4761	\N	1	1838	\N	\N	2018-10-02 00:00:00+07	1838.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5820	\N	1	2325	\N	\N	2019-10-16 00:00:00+07	2325.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4760	\N	1	1837	\N	\N	2018-10-01 00:00:00+07	1837.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4759	\N	1	1836	\N	\N	2018-10-01 00:00:00+07	1836.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5818	\N	1	2525	\N	\N	2019-10-15 00:00:00+07	2525.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5817	\N	1	814	\N	\N	2019-10-15 00:00:00+07	0814.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5816	\N	1	2324	\N	\N	2019-10-15 00:00:00+07	2324.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4758	\N	1	1835	\N	\N	2018-10-01 00:00:00+07	1835.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4757	\N	1	1834	\N	\N	2018-10-01 00:00:00+07	1834.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4756	\N	1	1833	\N	\N	2018-10-01 00:00:00+07	1833.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4755	\N	1	1832	\N	\N	2018-10-01 00:00:00+07	1832.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4754	\N	1	1831	\N	\N	2018-10-01 00:00:00+07	1831.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4753	\N	1	1830	\N	\N	2018-10-01 00:00:00+07	1830.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4752	\N	1	712	\N	\N	2018-10-01 00:00:00+07	0712.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4751	\N	1	1829	\N	\N	2018-09-28 00:00:00+07	1829.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5614	\N	1	2227	\N	\N	2019-08-12 00:00:00+07	2227.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4746	\N	1	1828	\N	\N	2018-09-27 00:00:00+07	1828.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4745	\N	1	1827	\N	\N	2018-09-27 00:00:00+07	1827.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4744	\N	1	1826	\N	\N	2018-09-27 00:00:00+07	1826.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4743	\N	1	1825	\N	\N	2018-09-26 00:00:00+07	1825.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4742	\N	1	1824	\N	\N	2018-09-24 00:00:00+07	1824.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
713	\N	2	761	\N	\N	2014-06-04 00:00:00+07	0761.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
714	\N	2	762	\N	\N	2014-06-04 00:00:00+07	0762.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
715	\N	2	763	\N	\N	2014-06-04 00:00:00+07	0763.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
716	\N	2	764	\N	\N	2014-06-04 00:00:00+07	0764.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
717	\N	2	765	\N	\N	2014-06-04 00:00:00+07	0765.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4740	\N	1	1823	\N	\N	2018-09-24 00:00:00+07	1823.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5789	\N	1	532	\N	\N	2019-10-10 00:00:00+07	0532.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5787	\N	1	531	\N	\N	2019-10-09 00:00:00+07	0531.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4739	\N	1	1822	\N	\N	2018-09-24 00:00:00+07	1822.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
718	\N	2	766	\N	\N	2014-06-04 00:00:00+07	0766.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4736	\N	1	1820	\N	\N	2018-09-19 00:00:00+07	1820.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4730	\N	1	1818	\N	\N	2018-09-17 00:00:00+07	1818.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
719	\N	2	768	\N	\N	2014-06-05 00:00:00+07	0768.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4729	\N	1	1817	\N	\N	2018-09-17 00:00:00+07	1817.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
720	\N	2	769	\N	\N	2014-06-11 00:00:00+07	0769.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5761	\N	1	526	\N	\N	2019-10-02 00:00:00+07	0526.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4727	\N	1	710	\N	\N	2018-09-13 00:00:00+07	0710.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5616	\N	1	2229	\N	\N	2019-08-13 00:00:00+07	2229.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5754	\N	1	985	\N	\N	2019-09-30 00:00:00+07	0985.61.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5753	\N	1	2306	\N	\N	2019-09-26 00:00:00+07	2306.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5752	\N	1	2305	\N	\N	2019-09-26 00:00:00+07	2305.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5617	\N	1	2230	\N	\N	2019-08-13 00:00:00+07	2230.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5747	\N	1	2304	\N	\N	2019-09-25 00:00:00+07	2304.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5745	\N	1	2303	\N	\N	2019-09-25 00:00:00+07	2303.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5779	\N	1	2317	\N	\N	2019-10-07 00:00:00+07	2317.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
721	\N	2	770	\N	\N	2014-06-12 00:00:00+07	0770.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5778	\N	1	2316	\N	\N	2019-10-04 00:00:00+07	2316.65.221	1	\N	2020-09-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5742	\N	1	2300	\N	\N	2019-09-25 00:00:00+07	2300.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5741	\N	1	2299	\N	\N	2019-09-25 00:00:00+07	2299.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5619	\N	1	2231	\N	\N	2019-08-14 00:00:00+07	2231.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-06-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5623	\N	1	2233	\N	\N	2019-08-14 00:00:00+07	2233.65.221	1	\N	2020-06-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5737	\N	1	2295	\N	\N	2019-09-24 00:00:00+07	2295.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5628	\N	1	2235	\N	\N	2019-08-15 00:00:00+07	2235.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
722	\N	2	771	\N	\N	2014-06-16 00:00:00+07	0771.65.221	1	\N	2019-09-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5734	\N	1	2293	\N	\N	2019-09-24 00:00:00+07	2293.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5733	\N	1	983	\N	\N	2019-09-24 00:00:00+07	0983.61.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5629	\N	1	2236	\N	\N	2019-08-15 00:00:00+07	2236.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
723	\N	2	772	\N	\N	2014-06-24 00:00:00+07	0772.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
724	\N	2	773	\N	\N	2014-06-26 00:00:00+07	0773.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3758	\N	1	2151	\N	\N	2017-06-02 00:00:00+07	2151.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5637	\N	1	2240	\N	\N	2019-08-21 00:00:00+07	2240.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3695	\N	1	524	\N	\N	2017-04-07 00:00:00+07	0524.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3694	\N	1	517	\N	\N	2017-04-06 00:00:00+07	0517.64.221	2	2017-08-15 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3693	\N	1	513	\N	\N	2017-04-05 00:00:00+07	0513.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3659	\N	1	1611	\N	\N	2017-07-17 00:00:00+07	1611.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5764	\N	1	2308	\N	\N	2019-10-03 00:00:00+07	2308.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5765	\N	1	2309	\N	\N	2019-10-03 00:00:00+07	2309.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3608	\N	1	1621	\N	\N	2017-04-03 00:00:00+07	1621.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3335	\N	1	1565	\N	\N	2017-02-01 00:00:00+07	1565.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3654	\N	1	529	\N	\N	2017-03-01 00:00:00+07	0529.64.221	1	2017-10-09 00:00:00+07	2019-05-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3337	\N	1	1546	\N	\N	2017-01-02 00:00:00+07	1546.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4725	\N	1	436	\N	\N	2018-09-13 00:00:00+07	0436.63.221	2	2019-04-01 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
725	\N	2	774	\N	\N	2014-06-30 00:00:00+07	0774.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3341	\N	1	1533	\N	\N	2017-02-01 00:00:00+07	1533.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4724	\N	1	432	\N	\N	2018-09-13 00:00:00+07	0432.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
726	\N	2	775	\N	\N	2014-09-19 00:00:00+07	0775.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4723	\N	1	431	\N	\N	2018-09-13 00:00:00+07	0431.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4722	\N	1	2284	\N	\N	2018-09-13 00:00:00+07	2284.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4721	\N	1	1816	\N	\N	2018-09-13 00:00:00+07	1816.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4720	\N	1	1814	\N	\N	2018-09-13 00:00:00+07	1814.65.221	2	2019-08-14 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3347	\N	1	2076	\N	\N	2017-03-01 00:00:00+07	2076.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4715	\N	1	708	\N	\N	2018-09-12 00:00:00+07	0708.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3649	\N	1	1585	\N	\N	2017-03-09 00:00:00+07	1585.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4714	\N	1	707	\N	\N	2018-09-12 00:00:00+07	0707.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3590	\N	1	1587	\N	\N	2017-03-01 00:00:00+07	1587.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5706	\N	1	2492	\N	\N	2019-09-20 00:00:00+07	2492.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3350	\N	1	511	\N	\N	2017-03-01 00:00:00+07	0511.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3351	\N	1	512	\N	\N	2017-03-01 00:00:00+07	0512.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4713	\N	1	706	\N	\N	2018-09-12 00:00:00+07	0706.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4712	\N	1	705	\N	\N	2018-09-12 00:00:00+07	0705.64.221	2	2018-09-12 00:00:00+07	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4711	\N	1	704	\N	\N	2018-09-12 00:00:00+07	0704.64.221	1	\N	2019-09-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3358	\N	1	1541	\N	\N	2017-01-02 00:00:00+07	1541.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3359	\N	1	811	\N	\N	2017-03-01 00:00:00+07	0811.64.221	2	2018-09-05 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4719	\N	1	1812	\N	\N	2018-09-13 00:00:00+07	1812.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4710	\N	1	703	\N	\N	2018-09-12 00:00:00+07	0703.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3361	\N	1	1562	\N	\N	2017-03-01 00:00:00+07	1562.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
727	\N	2	776	\N	\N	2014-07-07 00:00:00+07	0776.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
728	\N	2	777	\N	\N	2014-07-07 00:00:00+07	0777.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-12-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3643	\N	1	544	\N	\N	2017-04-03 00:00:00+07	0544.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3367	\N	1	514	\N	\N	2017-03-01 00:00:00+07	0514.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3369	\N	1	516	\N	\N	2017-03-01 00:00:00+07	0516.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3757	\N	1	2152	\N	\N	2017-06-02 00:00:00+07	2152.62.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5701	\N	1	980	\N	\N	2019-09-19 00:00:00+07	0980.61.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5660	\N	1	2250	\N	\N	2019-09-05 00:00:00+07	2250.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3373	\N	1	518	\N	\N	2017-02-01 00:00:00+07	0518.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
729	\N	2	778	\N	\N	2014-07-07 00:00:00+07	0778.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
730	\N	2	779	\N	\N	2014-07-07 00:00:00+07	0779.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4709	\N	1	702	\N	\N	2018-09-12 00:00:00+07	0702.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3620	\N	1	1624	\N	\N	2017-05-15 00:00:00+07	1624.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5774	\N	1	2312	\N	\N	2019-10-04 00:00:00+07	2312.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
731	\N	2	780	\N	\N	2014-07-07 00:00:00+07	0780.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2022-03-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4707	\N	1	701	\N	\N	2018-09-12 00:00:00+07	0701.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
732	\N	2	782	\N	\N	2014-07-08 00:00:00+07	0782.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3391	\N	1	1604	\N	\N	2017-03-01 00:00:00+07	1604.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3395	\N	1	527	\N	\N	2017-03-01 00:00:00+07	0527.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3399	\N	1	530	\N	\N	2017-03-01 00:00:00+07	0530.64.221	1	\N	2018-10-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
733	\N	2	784	\N	\N	2014-07-18 00:00:00+07	0784.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
734	\N	2	785	\N	\N	2014-08-08 00:00:00+07	0785.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5661	\N	1	2251	\N	\N	2019-09-05 00:00:00+07	2251.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
735	\N	2	790	\N	\N	2014-08-21 00:00:00+07	0790.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
736	\N	2	791	\N	\N	2014-08-21 00:00:00+07	0791.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6486	\N	1	2506	\N	\N	2021-06-03 00:00:00+07	2506.65.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-06-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4705	\N	1	699	\N	\N	2018-09-12 00:00:00+07	0699.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4704	\N	1	698	\N	\N	2018-09-12 00:00:00+07	0698.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
737	\N	2	792	\N	\N	2014-08-21 00:00:00+07	0792.65.221	2	2017-03-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3416	\N	1	1606	\N	\N	2017-03-01 00:00:00+07	1606.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
738	\N	2	793	\N	\N	2014-08-21 00:00:00+07	0793.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5664	\N	1	2252	\N	\N	2019-09-09 00:00:00+07	2252.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3420	\N	1	1568	\N	\N	2017-02-01 00:00:00+07	1568.65.221	2	2018-08-01 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3421	\N	1	1529	\N	\N	2016-11-01 00:00:00+07	1529.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
739	\N	2	794	\N	\N	2014-08-21 00:00:00+07	0794.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3423	\N	1	1578	\N	\N	2017-03-01 00:00:00+07	1578.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
740	\N	2	795	\N	\N	2014-08-21 00:00:00+07	0795.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5673	\N	1	2256	\N	\N	2019-09-11 00:00:00+07	2256.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
741	\N	2	796	\N	\N	2014-08-25 00:00:00+07	0796.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3533	\N	1	675	\N	\N	2016-12-01 00:00:00+07	0675.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4703	\N	1	697	\N	\N	2018-09-12 00:00:00+07	0697.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5391	\N	1	2146	\N	\N	2019-06-12 00:00:00+07	2146.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3940	\N	1	570	\N	\N	2017-08-16 00:00:00+07	0570.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3437	\N	1	1579	\N	\N	2017-03-01 00:00:00+07	1579.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3941	\N	1	569	\N	\N	2017-08-16 00:00:00+07	0569.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5689	\N	1	2267	\N	\N	2019-09-17 00:00:00+07	2267.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4718	\N	1	1813	\N	\N	2018-09-13 00:00:00+07	1813.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3443	\N	1	1596	\N	\N	2017-03-01 00:00:00+07	1596.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3524	\N	1	1564	\N	\N	2017-04-28 00:00:00+07	1564.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4717	\N	1	1811	\N	\N	2018-09-13 00:00:00+07	1811.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4700	\N	1	696	\N	\N	2018-09-12 00:00:00+07	0696.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
743	\N	2	797	\N	\N	2014-09-01 00:00:00+07	0797.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
744	\N	2	798	\N	\N	2014-08-29 00:00:00+07	0798.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3463	\N	1	1612	\N	\N	2017-03-01 00:00:00+07	1612.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5686	\N	1	2264	\N	\N	2019-09-17 00:00:00+07	2264.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3506	\N	1	1614	\N	\N	2017-03-01 00:00:00+07	1614.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4716	\N	1	1800	\N	\N	2018-09-13 00:00:00+07	1800.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5684	\N	1	2263	\N	\N	2019-09-16 00:00:00+07	2263.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
745	\N	2	799	\N	\N	2014-09-01 00:00:00+07	0799.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3512	\N	1	1308	\N	\N	2016-12-01 00:00:00+07	1308.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
746	\N	2	800	\N	\N	2014-09-01 00:00:00+07	0800.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
747	\N	2	801	\N	\N	2014-09-25 00:00:00+07	0801.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
748	\N	2	803	\N	\N	2014-09-08 00:00:00+07	0803.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3518	\N	1	1526	\N	\N	2016-12-01 00:00:00+07	1526.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
749	\N	2	804	\N	\N	2014-09-17 00:00:00+07	0804.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5679	\N	1	2486	\N	\N	2019-09-12 00:00:00+07	2486.62.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5677	\N	1	2258	\N	\N	2019-09-11 00:00:00+07	2258.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
750	\N	2	808	\N	\N	2014-09-09 00:00:00+07	0808.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4692	\N	1	1804	\N	\N	2018-09-05 00:00:00+07	1804.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
751	\N	2	809	\N	\N	2014-09-09 00:00:00+07	0809.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4691	\N	1	1803	\N	\N	2018-09-04 00:00:00+07	1803.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
752	\N	2	810	\N	\N	2014-09-09 00:00:00+07	0810.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4690	\N	1	1802	\N	\N	2018-09-04 00:00:00+07	1802.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
753	\N	2	812	\N	\N	2014-09-09 00:00:00+07	0812.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
754	\N	2	813	\N	\N	2014-09-09 00:00:00+07	0813.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
755	\N	2	814	\N	\N	2014-09-10 00:00:00+07	0814.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
756	\N	2	819	\N	\N	2014-09-15 00:00:00+07	0819.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
757	\N	2	820	\N	\N	2014-09-16 00:00:00+07	0820.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
758	\N	2	821	\N	\N	2014-09-16 00:00:00+07	0821.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4689	\N	1	1801	\N	\N	2018-09-04 00:00:00+07	1801.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
759	\N	2	825	\N	\N	2014-09-17 00:00:00+07	0825.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4685	\N	1	1799	\N	\N	2018-08-30 00:00:00+07	1799.65.221	2	2019-07-22 00:00:00+07	2020-01-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4684	\N	1	1798	\N	\N	2018-08-30 00:00:00+07	1798.65.221	1	\N	2020-01-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
760	\N	2	826	\N	\N	2014-10-23 00:00:00+07	0826.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
761	\N	2	827	\N	\N	2014-09-18 00:00:00+07	0827.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
762	\N	2	828	\N	\N	2014-09-18 00:00:00+07	0828.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
763	\N	2	829	\N	\N	2014-09-18 00:00:00+07	0829.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5394	\N	1	2147	\N	\N	2019-06-13 00:00:00+07	2147.65.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
764	\N	2	830	\N	\N	2014-09-19 00:00:00+07	0830.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
765	\N	2	831	\N	\N	2014-09-22 00:00:00+07	0831.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
766	\N	2	832	\N	\N	2014-09-23 00:00:00+07	0832.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
768	\N	2	833	\N	\N	2014-09-23 00:00:00+07	0833.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4681	\N	1	1797	\N	\N	2018-08-30 00:00:00+07	1797.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
769	\N	2	834	\N	\N	2014-09-23 00:00:00+07	0834.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6375	\N	1	601	\N	\N	2021-02-04 00:00:00+07	0601.63.221	1	\N	\N	SA	\N	10208	Amalia Mahardani	\N	\N	f	f	2021-02-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
770	\N	2	835	\N	\N	2014-09-23 00:00:00+07	0835.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
771	\N	2	836	\N	\N	2014-09-23 00:00:00+07	0836.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
772	\N	2	837	\N	\N	2014-09-23 00:00:00+07	0837.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
773	\N	2	838	\N	\N	2014-09-25 00:00:00+07	0838.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4679	\N	1	1795	\N	\N	2018-08-29 00:00:00+07	1795.65.221	1	\N	2020-08-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
774	\N	2	839	\N	\N	2014-09-25 00:00:00+07	0839.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4678	\N	1	1794	\N	\N	2018-08-29 00:00:00+07	1794.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
775	\N	2	840	\N	\N	2014-09-26 00:00:00+07	0840.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
776	\N	2	841	\N	\N	2014-09-29 00:00:00+07	0841.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4677	\N	1	1793	\N	\N	2018-08-29 00:00:00+07	1793.65.221	1	\N	2019-09-01 00:00:00+07	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4676	\N	1	1792	\N	\N	2018-08-29 00:00:00+07	1792.65.221	1	\N	2020-09-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4675	\N	1	1791	\N	\N	2018-08-29 00:00:00+07	1791.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5395	\N	1	495	\N	\N	2019-06-13 00:00:00+07	0495.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4665	\N	1	2277	\N	\N	2018-08-21 00:00:00+07	2277.62.221	1	\N	2020-07-11 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4664	\N	1	429	\N	\N	2018-08-21 00:00:00+07	0429.63.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4663	\N	1	690	\N	\N	2018-08-16 00:00:00+07	0690.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
777	\N	2	842	\N	\N	2014-09-30 00:00:00+07	0842.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2722	\N	2	1916	\N	\N	2015-10-13 00:00:00+07	1916.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5396	\N	1	496	\N	\N	2019-06-13 00:00:00+07	0496.63.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4659	\N	1	688	\N	\N	2018-08-13 00:00:00+07	0688.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
778	\N	2	843	\N	\N	2014-09-30 00:00:00+07	0843.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
779	\N	2	844	\N	\N	2014-10-02 00:00:00+07	0844.65.221	1	\N	2018-09-01 00:00:00+07	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
780	\N	2	845	\N	\N	2014-09-26 00:00:00+07	0845.65.221	1	\N	2020-08-01 00:00:00+07	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
781	\N	2	846	\N	\N	2014-10-01 00:00:00+07	0846.65.221	2	2020-06-01 00:00:00+07	\N	SA	\N	10208	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4656	\N	1	1786	\N	\N	2018-08-09 00:00:00+07	1786.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4655	\N	1	687	\N	\N	2018-08-09 00:00:00+07	0687.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
782	\N	2	847	\N	\N	2014-10-02 00:00:00+07	0847.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4652	\N	1	686	\N	\N	2018-08-09 00:00:00+07	0686.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
783	\N	2	848	\N	\N	2014-10-02 00:00:00+07	0848.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5399	\N	1	950	\N	\N	2019-06-13 00:00:00+07	0950.61.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4643	\N	1	1780	\N	\N	2018-08-06 00:00:00+07	1780.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
784	\N	2	849	\N	\N	2014-10-03 00:00:00+07	0849.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4636	\N	1	1777	\N	\N	2018-08-02 00:00:00+07	1777.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
785	\N	2	850	\N	\N	2014-10-06 00:00:00+07	0850.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4635	\N	1	1776	\N	\N	2018-08-02 00:00:00+07	1776.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4634	\N	1	1775	\N	\N	2018-08-02 00:00:00+07	1775.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4633	\N	1	1774	\N	\N	2018-08-02 00:00:00+07	1774.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4628	\N	1	1771	\N	\N	2018-07-30 00:00:00+07	1771.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
786	\N	2	851	\N	\N	2014-10-06 00:00:00+07	0851.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4627	\N	1	683	\N	\N	2018-07-30 00:00:00+07	0683.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5401	\N	1	773	\N	\N	2019-06-13 00:00:00+07	0773.64.221	1	\N	\N	SA	\N	10208	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4619	\N	1	1769	\N	\N	2018-07-24 00:00:00+07	1769.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4618	\N	1	679	\N	\N	2018-07-24 00:00:00+07	0679.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
787	\N	2	852	\N	\N	2014-10-07 00:00:00+07	0852.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
788	\N	2	853	\N	\N	2014-10-09 00:00:00+07	0853.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
789	\N	2	855	\N	\N	2014-10-09 00:00:00+07	0855.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4606	\N	1	1763	\N	\N	2018-07-12 00:00:00+07	1763.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3944	\N	1	571	\N	\N	2017-08-21 00:00:00+07	0571.64.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4599	\N	1	1759	\N	\N	2018-07-11 00:00:00+07	1759.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4597	\N	1	1757	\N	\N	2018-07-09 00:00:00+07	1757.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4598	\N	1	1758	\N	\N	2018-07-11 00:00:00+07	1758.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4596	\N	1	1756	\N	\N	2018-07-09 00:00:00+07	1756.65.221	2	2018-08-27 00:00:00+07	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
790	\N	2	856	\N	\N	2014-10-09 00:00:00+07	0856.65.221	1	\N	\N	SA	\N	10208	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5954	\N	1	1002	\N	\N	2019-12-09 00:00:00+07	1002.61.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7441	\N	1	711	\N	\N	2022-06-06 00:00:00+07	0711.63.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-06-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5220	\N	1	2374	\N	\N	2019-04-12 00:00:00+07	2374.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7455	\N	1	2731	\N	\N	2022-06-10 00:00:00+07	2731.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-06-10 00:00:00+07	\N	11:00	20:30	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4447	\N	1	2243	\N	\N	2018-03-14 00:00:00+07	2243.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5943	\N	1	2574	\N	\N	2019-12-05 00:00:00+07	2574.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5942	\N	1	550	\N	\N	2019-12-03 00:00:00+07	0550.63.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5960	\N	1	2576	\N	\N	2019-12-12 00:00:00+07	2576.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-09-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7462	\N	1	2902	\N	\N	2022-06-14 00:00:00+07	2902.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-06-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4451	\N	1	1705	\N	\N	2018-03-15 00:00:00+07	1705.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3851	\N	2	357	\N	\N	2017-07-07 00:00:00+07	0357.63.102	2	2020-06-01 00:00:00+07	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-02-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2800	\N	2	1962	\N	\N	2016-03-10 00:00:00+07	1962.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5975	\N	1	2578	\N	\N	2019-12-30 00:00:00+07	2578.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3895	\N	1	850	\N	\N	2017-07-24 00:00:00+07	0850.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5115	\N	1	2034	\N	\N	2019-03-18 00:00:00+07	2034.65.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2790	\N	2	1427	\N	\N	2016-03-01 00:00:00+07	1427.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3896	\N	1	2166	\N	\N	2017-07-24 00:00:00+07	2166.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7481	\N	1	2909	\N	\N	2022-06-20 00:00:00+07	2909.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-06-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5920	\N	1	2563	\N	\N	2019-11-19 00:00:00+07	2563.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6926	\N	1	2603	\N	\N	2021-11-23 00:00:00+07	2603.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3809	\N	2	2160	\N	\N	2017-06-13 00:00:00+07	2160.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7484	\N	1	2738	\N	\N	2022-06-22 00:00:00+07	2738.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	09:00	21:00	80	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4792	\N	1	2292	\N	\N	2018-10-16 00:00:00+07	2292.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2426	\N	2	1727	\N	\N	2014-03-14 00:00:00+07	1727.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3897	\N	1	2170	\N	\N	2017-07-24 00:00:00+07	2170.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3894	\N	2	2169	\N	\N	2017-07-24 00:00:00+07	2169.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2417	\N	2	1708	\N	\N	2014-01-15 00:00:00+07	1708.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4658	\N	1	2276	\N	\N	2018-08-10 00:00:00+07	2276.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2781	\N	2	1940	\N	\N	2016-02-10 00:00:00+07	1940.62.102	2	2020-02-01 00:00:00+07	\N	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6508	\N	1	2512	\N	\N	2021-06-21 00:00:00+07	2512.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-06-21 00:00:00+07	\N	10:00	21:00	7	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3914	\N	1	359	\N	\N	2017-07-26 00:00:00+07	0359.63.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3915	\N	1	1634	\N	\N	2017-07-27 00:00:00+07	1634.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3916	\N	1	1637	\N	\N	2017-07-27 00:00:00+07	1637.65.102	2	2020-10-01 00:00:00+07	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2020-10-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5898	\N	1	2558	\N	\N	2019-11-11 00:00:00+07	2558.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2416	\N	2	685	\N	\N	2014-01-13 00:00:00+07	0685.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5896	\N	1	2557	\N	\N	2019-11-11 00:00:00+07	2557.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2406	\N	2	684	\N	\N	2013-11-12 00:00:00+07	0684.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6500	\N	1	2710	\N	\N	2021-06-15 00:00:00+07	2710.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-06-15 00:00:00+07	\N	09:00	23:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2780	\N	2	1941	\N	\N	2016-02-10 00:00:00+07	1941.62.102	2	2020-04-27 00:00:00+07	\N	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5597	\N	1	791	\N	\N	2019-08-09 00:00:00+07	0791.64.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2404	\N	2	653	\N	\N	2013-10-30 00:00:00+07	0653.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2779	\N	2	1958	\N	\N	2016-02-10 00:00:00+07	1958.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5888	\N	1	2335	\N	\N	2019-11-07 00:00:00+07	2335.65.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6342	\N	1	2653	\N	\N	2021-01-12 00:00:00+07	2653.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-01-12 00:00:00+07	48	11:00	21:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5886	\N	1	2555	\N	\N	2019-11-07 00:00:00+07	2555.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3384	\N	2	1555	\N	\N	2017-01-02 00:00:00+07	1555.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2401	\N	2	1700	\N	\N	2013-10-17 00:00:00+07	1700.62.102	1	\N	2020-06-01 00:00:00+07	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2399	\N	2	1694	\N	\N	2013-10-01 00:00:00+07	1694.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5994	\N	1	2583	\N	\N	2020-01-09 00:00:00+07	2583.62.102	1	\N	\N	SA	\N	10111	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2396	\N	2	647	\N	\N	2013-09-16 00:00:00+07	0647.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2777	\N	2	1950	\N	\N	2016-02-09 00:00:00+07	1950.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2382	\N	2	675	\N	\N	2013-04-16 00:00:00+07	0675.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2376	\N	2	1666	\N	\N	2013-02-27 00:00:00+07	1666.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4784	\N	1	438	\N	\N	2018-10-09 00:00:00+07	0438.63.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2366	\N	2	370	\N	\N	2013-01-09 00:00:00+07	0370.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2766	\N	2	1942	\N	\N	2016-01-20 00:00:00+07	1942.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3917	\N	2	2173	\N	\N	2017-07-27 00:00:00+07	2173.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6936	\N	1	2606	\N	\N	2021-11-24 00:00:00+07	2606.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-11-24 00:00:00+07	\N	08:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2365	\N	2	640	\N	\N	2013-01-08 00:00:00+07	0640.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2354	\N	2	668	\N	\N	2012-07-16 00:00:00+07	0668.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2335	\N	2	622	\N	\N	2012-04-25 00:00:00+07	0622.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2322	\N	2	283	\N	\N	2012-02-16 00:00:00+07	0283.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2321	\N	2	1627	\N	\N	2012-02-14 00:00:00+07	1627.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2756	\N	2	1939	\N	\N	2016-01-08 00:00:00+07	1939.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1743	\N	2	226	\N	\N	2011-02-07 00:00:00+07	0226.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2929	\N	1	2023	\N	\N	2017-03-16 00:00:00+07	2023.62.102	2	2017-04-10 00:00:00+07	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2755	\N	2	464	\N	\N	2016-01-07 00:00:00+07	0464.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2751	\N	2	459	\N	\N	2015-12-10 00:00:00+07	0459.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2748	\N	2	1362	\N	\N	2015-12-04 00:00:00+07	1362.65.102	2	2017-03-01 00:00:00+07	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2285	\N	2	600	\N	\N	2011-08-25 00:00:00+07	0600.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1744	\N	2	778	\N	\N	2011-02-07 00:00:00+07	0778.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4508	\N	1	660	\N	\N	2018-04-13 00:00:00+07	0660.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5548	\N	1	2456	\N	\N	2019-07-24 00:00:00+07	2456.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6472	\N	1	2700	\N	\N	2021-05-10 00:00:00+07	2700.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-05-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6966	\N	1	2779	\N	\N	2021-12-07 00:00:00+07	2779.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-12-07 00:00:00+07	\N	09:00	20:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3925	\N	2	2177	\N	\N	2017-08-02 00:00:00+07	2177.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1745	\N	2	239	\N	\N	2011-02-07 00:00:00+07	0239.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3926	\N	1	2178	\N	\N	2017-08-02 00:00:00+07	2178.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3930	\N	1	2180	\N	\N	2017-08-03 00:00:00+07	2180.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4511	\N	1	2252	\N	\N	2018-04-16 00:00:00+07	2252.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1746	\N	2	217	\N	\N	2011-02-07 00:00:00+07	0217.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6022	\N	1	1010	\N	\N	2020-02-04 00:00:00+07	1010.61.102	1	\N	\N	SA	\N	10111	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2084	\N	2	371	\N	\N	2011-02-09 00:00:00+07	0371.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6769	\N	1	2545	\N	\N	2021-09-20 00:00:00+07	2545.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-09-20 00:00:00+07	\N	00:00	12:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6042	\N	1	2599	\N	\N	2020-02-27 00:00:00+07	2599.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-07-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2739	\N	2	1921	\N	\N	2015-11-11 00:00:00+07	1921.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4610	\N	1	1764	\N	\N	2018-07-16 00:00:00+07	1764.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6771	\N	1	2737	\N	\N	2021-09-21 00:00:00+07	2737.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-09-21 00:00:00+07	\N	11:00	20:00	2	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1680	\N	2	574	\N	\N	2014-04-02 00:00:00+07	0574.65.102	1	\N	2018-07-01 00:00:00+07	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4077	\N	1	2197	\N	\N	2017-09-22 00:00:00+07	2197.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4539	\N	1	1734	\N	\N	2018-05-02 00:00:00+07	1734.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2162	\N	2	1271	\N	\N	2011-02-11 00:00:00+07	1271.....62.102	2	2017-08-28 00:00:00+07	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4607	\N	1	676	\N	\N	2018-07-12 00:00:00+07	0676.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2738	\N	2	1363	\N	\N	2015-11-10 00:00:00+07	1363.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6948	\N	1	2775	\N	\N	2021-11-30 00:00:00+07	2775.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	11:00	20:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6460	\N	1	2694	\N	\N	2021-04-30 00:00:00+07	2694.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-04-30 00:00:00+07	\N	10:00	22:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2188	\N	2	269	\N	\N	2011-02-22 00:00:00+07	0269.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1751	\N	2	238	\N	\N	2011-02-07 00:00:00+07	0238.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3971	\N	2	2186	\N	\N	2017-09-04 00:00:00+07	2186.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1722	\N	2	813	\N	\N	2011-02-07 00:00:00+07	0813.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1723	\N	2	813	\N	\N	2011-02-07 00:00:00+07	0813.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4750	\N	1	2291	\N	\N	2018-09-27 00:00:00+07	2291.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6972	\N	1	2780	\N	\N	2021-12-08 00:00:00+07	2780.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-12-08 00:00:00+07	\N	10:00	18:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4749	\N	2	2290	\N	\N	2018-09-27 00:00:00+07	2290.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1724	\N	2	1202	\N	\N	2011-02-07 00:00:00+07	1202.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2724	\N	2	758	\N	\N	2015-10-16 00:00:00+07	0758.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1725	\N	2	335	\N	\N	2011-02-07 00:00:00+07	0335.62.102	1	\N	2020-06-01 00:00:00+07	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1726	\N	2	322	\N	\N	2011-02-07 00:00:00+07	0322.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6445	\N	1	2691	\N	\N	2021-04-15 00:00:00+07	2691.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-04-15 00:00:00+07	\N	08:00	22:00	90	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1727	\N	2	91	\N	\N	2011-02-07 00:00:00+07	0091.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2723	\N	2	751	\N	\N	2015-10-13 00:00:00+07	0751.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1728	\N	2	385	\N	\N	2011-02-07 00:00:00+07	0385.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4169	\N	1	375	\N	\N	2017-10-17 00:00:00+07	0375.63.102	2	2019-12-24 00:00:00+07	\N	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4782	\N	1	713	\N	\N	2018-10-05 00:00:00+07	0713.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4991	\N	1	2333	\N	\N	2019-02-15 00:00:00+07	2333.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2640	\N	2	1890	\N	\N	2015-06-12 00:00:00+07	1890.62.102	2	2018-11-27 00:00:00+07	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6091	\N	1	566	\N	\N	2020-06-08 00:00:00+07	0566.63.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1729	\N	2	373	\N	\N	2011-02-07 00:00:00+07	0373.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4591	\N	2	1752	\N	\N	2018-07-06 00:00:00+07	1752.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6988	\N	1	943	\N	\N	2021-12-14 00:00:00+07	0943.64.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-12-14 00:00:00+07	\N	\N	\N	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1697	\N	2	403	\N	\N	2011-02-07 00:00:00+07	0403.62.102	1	\N	2020-05-01 00:00:00+07	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1730	\N	2	348	\N	\N	2011-02-07 00:00:00+07	0348.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2706	\N	2	1911	\N	\N	2015-09-04 00:00:00+07	1911.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6994	\N	1	2786	\N	\N	2021-12-16 00:00:00+07	2786.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-12-16 00:00:00+07	\N	12:00	19:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6099	\N	1	2609	\N	\N	2020-06-15 00:00:00+07	2609.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2696	\N	2	1296	\N	\N	2015-08-20 00:00:00+07	1296.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1699	\N	2	234	\N	\N	2011-02-08 00:00:00+07	0234.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2690	\N	2	1851	\N	\N	2015-07-31 00:00:00+07	1851.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1731	\N	2	355	\N	\N	2011-02-07 00:00:00+07	0355.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1976	\N	2	1382	\N	\N	2011-02-08 00:00:00+07	1382.62.102	1	\N	2020-06-01 00:00:00+07	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-07-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2686	\N	2	1854	\N	\N	2015-07-27 00:00:00+07	1854.62.102	1	\N	\N	OA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1733	\N	2	283	\N	\N	2011-02-07 00:00:00+07	0283.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1734	\N	2	540	\N	\N	2011-02-07 00:00:00+07	0540.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4269	\N	1	1673	\N	\N	2017-11-09 00:00:00+07	1673.65.102	1	\N	2019-11-01 00:00:00+07	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2020-10-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6423	\N	1	2686	\N	\N	2021-04-08 00:00:00+07	2686.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-04-08 00:00:00+07	\N	11:00	21:00	40-50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2983	\N	1	2030	\N	\N	2015-11-02 00:00:00+07	2030.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5458	\N	1	2414	\N	\N	2019-06-27 00:00:00+07	2414.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1735	\N	2	342	\N	\N	2011-02-07 00:00:00+07	0342.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6848	\N	1	2762	\N	\N	2021-10-25 00:00:00+07	2762.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-10-25 00:00:00+07	\N	10:00	21:00	16	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1737	\N	2	340	\N	\N	2011-02-07 00:00:00+07	0340.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1738	\N	2	73	\N	\N	2011-02-07 00:00:00+07	0073.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2930	\N	1	2046	\N	\N	2017-03-16 00:00:00+07	2046.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1739	\N	2	1187	\N	\N	2011-02-07 00:00:00+07	1187.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7212	\N	1	2838	\N	\N	2022-02-03 00:00:00+07	2838.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-02-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4851	\N	1	892	\N	\N	2018-11-23 00:00:00+07	0892.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7220	\N	1	2840	\N	\N	2022-02-04 00:00:00+07	2840.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-02-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4983	\N	1	901	\N	\N	2019-02-14 00:00:00+07	0901.61.102	2	2020-06-01 00:00:00+07	\N	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2624	\N	2	1853	\N	\N	2015-05-22 00:00:00+07	1853.62.102	2	2017-05-23 00:00:00+07	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4864	\N	1	1903	\N	\N	2018-11-29 00:00:00+07	1903.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7108	\N	1	677	\N	\N	2022-01-13 00:00:00+07	0677.63.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5626	\N	1	2476	\N	\N	2019-08-14 00:00:00+07	2476.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7223	\N	1	2662	\N	\N	2022-02-04 00:00:00+07	2662.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-02-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4876	\N	1	2304	\N	\N	2018-12-03 00:00:00+07	2304.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3931	\N	1	2181	\N	\N	2017-08-04 00:00:00+07	2181.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7265	\N	1	2675	\N	\N	2022-02-16 00:00:00+07	2675.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-02-16 00:00:00+07	\N	11:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2592	\N	2	1852	\N	\N	2015-05-11 00:00:00+07	1852.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2591	\N	2	410	\N	\N	2015-05-11 00:00:00+07	0410.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5250	\N	1	2092	\N	\N	2019-05-03 00:00:00+07	2092.65.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6872	\N	1	2580	\N	\N	2021-11-03 00:00:00+07	2580.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-11-03 00:00:00+07	\N	10:00	21:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7286	\N	1	2678	\N	\N	2022-03-02 00:00:00+07	2678.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-03-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3735	\N	1	2148	\N	\N	2017-05-23 00:00:00+07	2148.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7294	\N	1	690	\N	\N	2022-03-08 00:00:00+07	0690.63.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-03-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7320	\N	1	2861	\N	\N	2022-03-15 00:00:00+07	2861.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-03-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5640	\N	1	2479	\N	\N	2019-08-21 00:00:00+07	2479.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4889	\N	1	2306	\N	\N	2018-12-06 00:00:00+07	2306.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4915	\N	1	2313	\N	\N	2018-12-20 00:00:00+07	2313.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7324	\N	1	2863	\N	\N	2022-03-16 00:00:00+07	2863.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-03-16 00:00:00+07	\N	10:00	21:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2567	\N	2	418	\N	\N	2015-05-07 00:00:00+07	0418.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2566	\N	2	3	\N	\N	2015-05-07 00:00:00+07	0003.00.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2843	\N	2	1902	\N	\N	2016-08-05 00:00:00+07	1902.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3704	\N	2	2120	\N	\N	2017-04-13 00:00:00+07	2120.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5271	\N	1	2095	\N	\N	2019-05-08 00:00:00+07	2095.65.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4925	\N	1	2316	\N	\N	2019-01-14 00:00:00+07	2316.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2564	\N	2	1236	\N	\N	2015-05-06 00:00:00+07	1236.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2562	\N	2	728	\N	\N	2015-05-06 00:00:00+07	0728.61.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5067	\N	1	2350	\N	\N	2019-03-06 00:00:00+07	2350.62.102	1	\N	2020-06-01 00:00:00+07	SA	\N	10111	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7332	\N	1	2869	\N	\N	2022-03-17 00:00:00+07	2869.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-03-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5342	\N	1	939	\N	\N	2019-05-23 00:00:00+07	0939.61.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1740	\N	2	229	\N	\N	2011-02-07 00:00:00+07	0229.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7142	\N	1	2812	\N	\N	2022-01-20 00:00:00+07	2812.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-01-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6897	\N	1	1095	\N	\N	2021-11-10 00:00:00+07	1095.61.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7339	\N	1	2693	\N	\N	2022-03-22 00:00:00+07	2693.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-03-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7352	\N	1	2695	\N	\N	2022-03-28 00:00:00+07	2695.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-03-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7359	\N	1	1014	\N	\N	2022-03-29 00:00:00+07	1014.64.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-03-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2548	\N	2	1213	\N	\N	2015-04-28 00:00:00+07	1213.65.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7361	\N	1	2698	\N	\N	2022-04-04 00:00:00+07	2698.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7369	\N	1	2876	\N	\N	2022-04-05 00:00:00+07	2876.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-04-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7372	\N	1	2878	\N	\N	2022-04-06 00:00:00+07	2878.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-04-06 00:00:00+07	\N	11:00	19:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1741	\N	2	217	\N	\N	2011-02-07 00:00:00+07	0217.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7375	\N	1	1163	\N	\N	2022-04-07 00:00:00+07	1163.61.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5136	\N	1	2358	\N	\N	2019-03-22 00:00:00+07	2358.62.102	1	\N	\N	SA	\N	10111	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2522	\N	2	1836	\N	\N	2015-03-12 00:00:00+07	1836.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7409	\N	1	2886	\N	\N	2022-05-12 00:00:00+07	2886.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-05-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7405	\N	1	705	\N	\N	2022-04-28 00:00:00+07	0705.63.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-04-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7413	\N	1	2718	\N	\N	2022-05-19 00:00:00+07	2718.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-05-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7414	\N	1	2888	\N	\N	2022-05-19 00:00:00+07	2888.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-05-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6369	\N	1	2473	\N	\N	2021-02-01 00:00:00+07	2473.65.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2021-02-01 00:00:00+07	\N	10:00	20:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7416	\N	1	2890	\N	\N	2022-05-23 00:00:00+07	2890.62.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-05-23 00:00:00+07	\N	10:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3491	\N	1	2040	\N	\N	2017-03-01 00:00:00+07	2040.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2809	\N	2	1965	\N	\N	2016-04-08 00:00:00+07	1965.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2489	\N	2	1798	\N	\N	2014-12-19 00:00:00+07	1798.62.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2487	\N	2	248	\N	\N	2014-12-15 00:00:00+07	0248.63.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4440	\N	1	638	\N	\N	2018-03-14 00:00:00+07	0638.64.102	1	\N	\N	SA	\N	10111	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7438	\N	1	1020	\N	\N	2022-06-06 00:00:00+07	1020.64.102	1	\N	\N	SA	\N	10111	Amalia Mahardani	\N	\N	f	f	2022-06-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5153	\N	1	2359	\N	\N	2019-03-27 00:00:00+07	2359.62.101	2	2020-12-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5300	\N	1	2108	\N	\N	2019-05-16 00:00:00+07	2108.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5163	\N	1	918	\N	\N	2019-04-01 00:00:00+07	0918.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5165	\N	1	919	\N	\N	2019-04-02 00:00:00+07	0919.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5173	\N	1	757	\N	\N	2019-04-04 00:00:00+07	0757.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5174	\N	1	2059	\N	\N	2019-04-04 00:00:00+07	2059.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5297	\N	1	477	\N	\N	2019-05-16 00:00:00+07	0477.63.101	2	2020-03-01 00:00:00+07	\N	OA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5179	\N	1	2362	\N	\N	2019-04-04 00:00:00+07	2362.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5180	\N	1	2363	\N	\N	2019-04-04 00:00:00+07	2363.62.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5182	\N	1	2364	\N	\N	2019-04-04 00:00:00+07	2364.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5183	\N	1	2365	\N	\N	2019-04-04 00:00:00+07	2365.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5184	\N	1	2366	\N	\N	2019-04-04 00:00:00+07	2366.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5186	\N	1	922	\N	\N	2019-04-05 00:00:00+07	0922.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5187	\N	1	923	\N	\N	2019-04-05 00:00:00+07	0923.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5188	\N	1	2065	\N	\N	2019-04-05 00:00:00+07	2065.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5190	\N	1	2367	\N	\N	2019-04-07 00:00:00+07	2367.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5196	\N	1	2070	\N	\N	2019-04-08 00:00:00+07	2070.65.101	2	2021-10-04 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5197	\N	1	2368	\N	\N	2019-04-08 00:00:00+07	2368.62.101	2	2021-09-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5198	\N	1	2369	\N	\N	2019-04-08 00:00:00+07	2369.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5199	\N	1	759	\N	\N	2019-04-08 00:00:00+07	0759.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5218	\N	1	2081	\N	\N	2019-04-12 00:00:00+07	2081.65.101	2	2020-10-18 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5219	\N	1	761	\N	\N	2019-04-12 00:00:00+07	0761.64.101	2	2019-10-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5225	\N	1	2082	\N	\N	2019-04-16 00:00:00+07	2082.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5229	\N	1	2375	\N	\N	2019-04-18 00:00:00+07	2375.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5230	\N	1	930	\N	\N	2019-04-22 00:00:00+07	0930.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5231	\N	1	762	\N	\N	2019-04-22 00:00:00+07	0762.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5266	\N	1	2378	\N	\N	2019-05-07 00:00:00+07	2378.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5287	\N	1	2384	\N	\N	2019-05-13 00:00:00+07	2384.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5286	\N	1	765	\N	\N	2019-05-09 00:00:00+07	0765.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5285	\N	1	2099	\N	\N	2019-05-09 00:00:00+07	2099.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5284	\N	1	2383	\N	\N	2019-05-09 00:00:00+07	2383.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5242	\N	1	2377	\N	\N	2019-04-30 00:00:00+07	2377.62.101	1	\N	\N	OA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5273	\N	1	2096	\N	\N	2019-05-08 00:00:00+07	2096.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5275	\N	1	476	\N	\N	2019-05-08 00:00:00+07	0476.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5574	\N	1	2470	\N	\N	2019-08-05 00:00:00+07	2470.62.101	1	\N	2020-06-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3756	\N	1	346	\N	\N	2017-05-31 00:00:00+07	0346.63.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7486	\N	1	2911	\N	\N	2022-06-23 00:00:00+07	2911.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7485	\N	1	1030	\N	\N	2022-06-23 00:00:00+07	1030.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	10:00	21:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7472	\N	1	1175	\N	\N	2022-06-16 00:00:00+07	1175.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	16:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7465	\N	1	1174	\N	\N	2022-06-14 00:00:00+07	1174.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-14 00:00:00+07	\N	10:00	22:00	3	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7464	\N	1	2904	\N	\N	2022-06-14 00:00:00+07	2904.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-14 00:00:00+07	\N	06:30	15:30	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7461	\N	1	1173	\N	\N	2022-06-14 00:00:00+07	1173.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7460	\N	1	1172	\N	\N	2022-06-14 00:00:00+07	1172.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7459	\N	1	2901	\N	\N	2022-06-14 00:00:00+07	2901.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7451	\N	1	1170	\N	\N	2022-06-09 00:00:00+07	1170.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7450	\N	1	1025	\N	\N	2022-06-09 00:00:00+07	1025.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7457	\N	1	1171	\N	\N	2022-06-13 00:00:00+07	1171.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7443	\N	1	1022	\N	\N	2022-06-07 00:00:00+07	1022.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7442	\N	1	1021	\N	\N	2022-06-07 00:00:00+07	1021.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7428	\N	1	2897	\N	\N	2022-05-25 00:00:00+07	2897.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7426	\N	1	2895	\N	\N	2022-05-25 00:00:00+07	2895.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7435	\N	1	2899	\N	\N	2022-06-02 00:00:00+07	2899.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7425	\N	1	2894	\N	\N	2022-05-25 00:00:00+07	2894.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7433	\N	1	1169	\N	\N	2022-05-31 00:00:00+07	1169.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-31 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7421	\N	1	1018	\N	\N	2022-05-24 00:00:00+07	1018.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7411	\N	1	2716	\N	\N	2022-05-17 00:00:00+07	2716.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7412	\N	1	2717	\N	\N	2022-05-18 00:00:00+07	2717.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7410	\N	1	2887	\N	\N	2022-05-12 00:00:00+07	2887.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7404	\N	1	704	\N	\N	2022-04-27 00:00:00+07	0704.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7403	\N	1	2715	\N	\N	2022-04-27 00:00:00+07	2715.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7400	\N	1	2713	\N	\N	2022-04-21 00:00:00+07	2713.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-21 00:00:00+07	\N	09:00	21:00	90	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7398	\N	1	2884	\N	\N	2022-04-21 00:00:00+07	2884.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7397	\N	1	2883	\N	\N	2022-04-21 00:00:00+07	2883.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7395	\N	1	703	\N	\N	2022-04-20 00:00:00+07	0703.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7406	\N	1	706	\N	\N	2022-05-10 00:00:00+07	0706.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-05-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7386	\N	1	702	\N	\N	2022-04-18 00:00:00+07	0702.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7382	\N	1	2881	\N	\N	2022-04-07 00:00:00+07	2881.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7368	\N	1	700	\N	\N	2022-04-05 00:00:00+07	0700.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7363	\N	1	2699	\N	\N	2022-04-04 00:00:00+07	2699.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7362	\N	1	699	\N	\N	2022-04-04 00:00:00+07	0699.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7351	\N	1	698	\N	\N	2022-03-25 00:00:00+07	0698.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7349	\N	1	2872	\N	\N	2022-03-23 00:00:00+07	2872.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7346	\N	1	1011	\N	\N	2022-03-23 00:00:00+07	1011.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1752	\N	2	35	\N	\N	2011-02-07 00:00:00+07	0035.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7345	\N	1	2871	\N	\N	2022-03-23 00:00:00+07	2871.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-23 00:00:00+07	\N	11:00	22:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7342	\N	1	696	\N	\N	2022-03-22 00:00:00+07	0696.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7330	\N	1	1158	\N	\N	2022-03-16 00:00:00+07	1158.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-16 00:00:00+07	\N	11:30	20:30	40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7325	\N	1	2864	\N	\N	2022-03-16 00:00:00+07	2864.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-16 00:00:00+07	\N	10:00	21:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7323	\N	1	2862	\N	\N	2022-03-16 00:00:00+07	2862.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7295	\N	1	691	\N	\N	2022-03-08 00:00:00+07	0691.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-08 00:00:00+07	\N	10:00	23:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7293	\N	1	2681	\N	\N	2022-03-08 00:00:00+07	2681.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7292	\N	1	2856	\N	\N	2022-03-08 00:00:00+07	2856.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7291	\N	1	2855	\N	\N	2022-03-08 00:00:00+07	2855.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7296	\N	1	2682	\N	\N	2022-03-09 00:00:00+07	2682.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7284	\N	1	1001	\N	\N	2022-02-24 00:00:00+07	1001.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7283	\N	1	689	\N	\N	2022-02-24 00:00:00+07	0689.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7280	\N	1	1000	\N	\N	2022-02-23 00:00:00+07	1000.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-23 00:00:00+07	\N	10:00	21:00	30-40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7275	\N	1	2677	\N	\N	2022-02-23 00:00:00+07	2677.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7274	\N	1	688	\N	\N	2022-02-22 00:00:00+07	0688.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7268	\N	1	687	\N	\N	2022-02-21 00:00:00+07	0687.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7257	\N	1	2671	\N	\N	2022-02-10 00:00:00+07	2671.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7256	\N	1	2845	\N	\N	2022-02-10 00:00:00+07	2845.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7252	\N	1	686	\N	\N	2022-02-10 00:00:00+07	0686.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7246	\N	1	1143	\N	\N	2022-02-08 00:00:00+07	1143.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7245	\N	1	1142	\N	\N	2022-02-08 00:00:00+07	1142.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7244	\N	1	1141	\N	\N	2022-02-08 00:00:00+07	1141.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7249	\N	1	2667	\N	\N	2022-02-09 00:00:00+07	2667.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-09 00:00:00+07	\N	11:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7243	\N	1	995	\N	\N	2022-02-08 00:00:00+07	0995.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7222	\N	1	990	\N	\N	2022-02-04 00:00:00+07	0990.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7219	\N	1	2839	\N	\N	2022-02-04 00:00:00+07	2839.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7210	\N	1	685	\N	\N	2022-02-02 00:00:00+07	0685.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-02 00:00:00+07	\N	10:00	21:00	40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1769	\N	2	114	\N	\N	2011-02-07 00:00:00+07	0114.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7179	\N	1	2656	\N	\N	2022-01-27 00:00:00+07	2656.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7177	\N	1	2655	\N	\N	2022-01-27 00:00:00+07	2655.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7174	\N	1	682	\N	\N	2022-01-25 00:00:00+07	0682.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-25 00:00:00+07	\N	05:00	12:00	20-25 ORANG	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7170	\N	1	2822	\N	\N	2022-01-24 00:00:00+07	2822.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7146	\N	1	2814	\N	\N	2022-01-21 00:00:00+07	2814.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7144	\N	1	975	\N	\N	2022-01-20 00:00:00+07	0975.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7119	\N	1	2648	\N	\N	2022-01-14 00:00:00+07	2648.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7109	\N	1	678	\N	\N	2022-01-13 00:00:00+07	0678.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	07:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7085	\N	1	2643	\N	\N	2022-01-13 00:00:00+07	2643.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	00:00	23:59	80 - 100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7135	\N	1	2652	\N	\N	2022-01-18 00:00:00+07	2652.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7083	\N	1	675	\N	\N	2022-01-12 00:00:00+07	0675.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7084	\N	1	2642	\N	\N	2022-01-13 00:00:00+07	2642.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7070	\N	1	2639	\N	\N	2022-01-11 00:00:00+07	2639.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7020	\N	1	953	\N	\N	2022-01-03 00:00:00+07	0953.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7018	\N	1	2633	\N	\N	2021-12-28 00:00:00+07	2633.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7013	\N	1	2630	\N	\N	2021-12-22 00:00:00+07	2630.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7012	\N	1	2788	\N	\N	2021-12-22 00:00:00+07	2788.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7011	\N	1	2787	\N	\N	2021-12-22 00:00:00+07	2787.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7010	\N	1	950	\N	\N	2021-12-22 00:00:00+07	0950.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7009	\N	1	1109	\N	\N	2021-12-22 00:00:00+07	1109.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7008	\N	1	1108	\N	\N	2021-12-22 00:00:00+07	1108.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7004	\N	1	2628	\N	\N	2021-12-20 00:00:00+07	2628.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7003	\N	1	947	\N	\N	2021-12-20 00:00:00+07	0947.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7002	\N	1	946	\N	\N	2021-12-17 00:00:00+07	0946.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7001	\N	1	945	\N	\N	2021-12-17 00:00:00+07	0945.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6996	\N	1	668	\N	\N	2021-12-17 00:00:00+07	0668.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-17 00:00:00+07	\N	10:00	19:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6995	\N	1	2624	\N	\N	2021-12-16 00:00:00+07	2624.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-16 00:00:00+07	\N	11:00	21:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6993	\N	1	2785	\N	\N	2021-12-16 00:00:00+07	2785.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-16 00:00:00+07	\N	10:00	21:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6991	\N	1	667	\N	\N	2021-12-16 00:00:00+07	0667.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6977	\N	1	1104	\N	\N	2021-12-09 00:00:00+07	1104.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6971	\N	1	666	\N	\N	2021-12-08 00:00:00+07	0666.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-08 00:00:00+07	\N	09:00	19:00	4	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6970	\N	1	665	\N	\N	2021-12-08 00:00:00+07	0665.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-08 00:00:00+07	\N	10:00	20:41	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6951	\N	1	939	\N	\N	2021-12-02 00:00:00+07	0939.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6949	\N	1	659	\N	\N	2021-12-01 00:00:00+07	0659.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6965	\N	1	664	\N	\N	2021-12-06 00:00:00+07	0664.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6962	\N	1	2777	\N	\N	2021-12-03 00:00:00+07	2777.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1753	\N	2	192	\N	\N	2011-02-07 00:00:00+07	0192.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6961	\N	1	2614	\N	\N	2021-12-03 00:00:00+07	2614.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6960	\N	1	662	\N	\N	2021-12-03 00:00:00+07	0662.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6959	\N	1	661	\N	\N	2021-12-03 00:00:00+07	0661.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6958	\N	1	2613	\N	\N	2021-12-03 00:00:00+07	2613.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6957	\N	1	2612	\N	\N	2021-12-03 00:00:00+07	2612.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6956	\N	1	942	\N	\N	2021-12-03 00:00:00+07	0942.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6955	\N	1	941	\N	\N	2021-12-03 00:00:00+07	0941.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6954	\N	1	2776	\N	\N	2021-12-03 00:00:00+07	2776.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6937	\N	1	656	\N	\N	2021-11-25 00:00:00+07	0656.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6953	\N	1	1102	\N	\N	2021-12-03 00:00:00+07	1102.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6935	\N	1	2774	\N	\N	2021-11-24 00:00:00+07	2774.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-24 00:00:00+07	\N	10:00	20:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6928	\N	1	933	\N	\N	2021-11-23 00:00:00+07	0933.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	06:30	22:00	60	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6927	\N	1	2604	\N	\N	2021-11-23 00:00:00+07	2604.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	12:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6919	\N	1	2596	\N	\N	2021-11-22 00:00:00+07	2596.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-22 00:00:00+07	\N	06:00	12:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6918	\N	1	2773	\N	\N	2021-11-19 00:00:00+07	2773.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6917	\N	1	655	\N	\N	2021-11-18 00:00:00+07	0655.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-18 00:00:00+07	\N	09:00	18:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6916	\N	1	932	\N	\N	2021-11-18 00:00:00+07	0932.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6913	\N	1	1101	\N	\N	2021-11-17 00:00:00+07	1101.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6912	\N	1	1100	\N	\N	2021-11-17 00:00:00+07	1100.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6911	\N	1	2595	\N	\N	2021-11-17 00:00:00+07	2595.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	09:00	20:30	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6909	\N	1	653	\N	\N	2021-11-17 00:00:00+07	0653.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	06:00	21:00	40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6908	\N	1	652	\N	\N	2021-11-17 00:00:00+07	0652.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	06:00	21:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6689	\N	1	1069	\N	\N	2021-08-26 00:00:00+07	1069.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6907	\N	1	651	\N	\N	2021-11-17 00:00:00+07	0651.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	06:00	21:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6906	\N	1	650	\N	\N	2021-11-17 00:00:00+07	0650.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	06:00	21:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6905	\N	1	2593	\N	\N	2021-11-17 00:00:00+07	2593.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	12:00	21:00	60	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6904	\N	1	1099	\N	\N	2021-11-17 00:00:00+07	1099.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	08:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6893	\N	1	649	\N	\N	2021-11-10 00:00:00+07	0649.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6886	\N	1	648	\N	\N	2021-11-09 00:00:00+07	0648.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6885	\N	1	2585	\N	\N	2021-11-09 00:00:00+07	2585.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6882	\N	1	2584	\N	\N	2021-11-09 00:00:00+07	2584.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6879	\N	1	2768	\N	\N	2021-11-05 00:00:00+07	2768.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6878	\N	1	2767	\N	\N	2021-11-05 00:00:00+07	2767.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	10:00	21:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6877	\N	1	645	\N	\N	2021-11-05 00:00:00+07	0645.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	10:30	19:30	4	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6873	\N	1	2581	\N	\N	2021-11-04 00:00:00+07	2581.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6870	\N	1	2579	\N	\N	2021-11-03 00:00:00+07	2579.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-03 00:00:00+07	\N	10:00	21:00	12	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6869	\N	1	2766	\N	\N	2021-11-03 00:00:00+07	2766.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-03 00:00:00+07	\N	10:00	21:00	7	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6864	\N	1	2765	\N	\N	2021-11-01 00:00:00+07	2765.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-01 00:00:00+07	\N	12:00	21:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6862	\N	1	2576	\N	\N	2021-11-01 00:00:00+07	2576.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-01 00:00:00+07	\N	10:00	02:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6860	\N	1	925	\N	\N	2021-10-29 00:00:00+07	0925.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6856	\N	1	2764	\N	\N	2021-10-28 00:00:00+07	2764.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6855	\N	1	643	\N	\N	2021-10-27 00:00:00+07	0643.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-27 00:00:00+07	\N	08:00	20:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6854	\N	1	2573	\N	\N	2021-10-27 00:00:00+07	2573.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6853	\N	1	1091	\N	\N	2021-10-27 00:00:00+07	1091.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-27 00:00:00+07	\N	08:00	20:30	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6850	\N	1	922	\N	\N	2021-10-26 00:00:00+07	0922.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6849	\N	1	1090	\N	\N	2021-10-26 00:00:00+07	1090.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-26 00:00:00+07	\N	16:30	03:00	250	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6838	\N	1	2760	\N	\N	2021-10-21 00:00:00+07	2760.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-21 00:00:00+07	\N	09:00	21:30	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6837	\N	1	2759	\N	\N	2021-10-21 00:00:00+07	2759.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6834	\N	1	920	\N	\N	2021-10-19 00:00:00+07	0920.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-19 00:00:00+07	\N	13:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6823	\N	1	2757	\N	\N	2021-10-15 00:00:00+07	2757.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6820	\N	1	2754	\N	\N	2021-10-14 00:00:00+07	2754.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-14 00:00:00+07	\N	10:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6818	\N	1	2753	\N	\N	2021-10-13 00:00:00+07	2753.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6816	\N	1	641	\N	\N	2021-10-13 00:00:00+07	0641.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-13 00:00:00+07	\N	15:00	23:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6815	\N	1	2751	\N	\N	2021-10-13 00:00:00+07	2751.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6812	\N	1	2563	\N	\N	2021-10-08 00:00:00+07	2563.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6811	\N	1	640	\N	\N	2021-10-08 00:00:00+07	0640.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6808	\N	1	2562	\N	\N	2021-10-07 00:00:00+07	2562.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6805	\N	1	2560	\N	\N	2021-10-06 00:00:00+07	2560.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6802	\N	1	1083	\N	\N	2021-10-01 00:00:00+07	1083.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6801	\N	1	2558	\N	\N	2021-10-01 00:00:00+07	2558.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-01 00:00:00+07	\N	10:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6800	\N	1	2745	\N	\N	2021-09-30 00:00:00+07	2745.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-30 00:00:00+07	\N	10:00	21:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6791	\N	1	2743	\N	\N	2021-09-28 00:00:00+07	2743.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-28 00:00:00+07	\N	13:00	23:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6784	\N	1	2741	\N	\N	2021-09-27 00:00:00+07	2741.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-27 00:00:00+07	\N	10:00	19:30	400	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6783	\N	1	2740	\N	\N	2021-09-27 00:00:00+07	2740.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-27 00:00:00+07	\N	10:00	20:00	90	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6782	\N	1	2548	\N	\N	2021-09-27 00:00:00+07	2548.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6778	\N	1	1082	\N	\N	2021-09-23 00:00:00+07	1082.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6777	\N	1	636	\N	\N	2021-09-23 00:00:00+07	0636.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-23 00:00:00+07	\N	10:00	19:30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6776	\N	1	635	\N	\N	2021-09-23 00:00:00+07	0635.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6775	\N	1	1081	\N	\N	2021-09-23 00:00:00+07	1081.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6767	\N	1	2543	\N	\N	2021-09-20 00:00:00+07	2543.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	1500000	f	f	2021-09-20 00:00:00+07	\N	08:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6762	\N	1	1077	\N	\N	2021-09-17 00:00:00+07	1077.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6724	\N	1	2726	\N	\N	2021-09-10 00:00:00+07	2726.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-10 00:00:00+07	\N	00:00	00:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6723	\N	1	2538	\N	\N	2021-09-10 00:00:00+07	2538.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-10 00:00:00+07	\N	12:00	20:30	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6719	\N	1	632	\N	\N	2021-09-09 00:00:00+07	0632.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-09 00:00:00+07	\N	09:00	20:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6712	\N	1	631	\N	\N	2021-09-08 00:00:00+07	0631.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-08 00:00:00+07	\N	10:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6709	\N	1	2532	\N	\N	2021-09-07 00:00:00+07	2532.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-07 00:00:00+07	\N	11:00	20:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6708	\N	1	2724	\N	\N	2021-09-07 00:00:00+07	2724.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6704	\N	1	2717	\N	\N	2021-09-02 00:00:00+07	2717.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-02 00:00:00+07	\N	11:00	19:30	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6700	\N	1	2531	\N	\N	2021-09-01 00:00:00+07	2531.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-09-01 00:00:00+07	\N	07:00	17:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6699	\N	1	2721	\N	\N	2021-08-31 00:00:00+07	2721.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-31 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6696	\N	1	2529	\N	\N	2021-08-30 00:00:00+07	2529.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6686	\N	1	1066	\N	\N	2021-08-25 00:00:00+07	1066.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6685	\N	1	913	\N	\N	2021-08-25 00:00:00+07	0913.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6684	\N	1	912	\N	\N	2021-08-25 00:00:00+07	0912.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6681	\N	1	2720	\N	\N	2021-08-25 00:00:00+07	2720.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6678	\N	1	2525	\N	\N	2021-08-24 00:00:00+07	2525.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-24 00:00:00+07	\N	08:00	22:00	20 - 40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6669	\N	1	625	\N	\N	2021-08-19 00:00:00+07	0625.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-19 00:00:00+07	\N	11:00	21:00	7	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6668	\N	1	624	\N	\N	2021-08-18 00:00:00+07	0624.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-18 00:00:00+07	\N	10:00	17:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6562	\N	1	2716	\N	\N	2021-08-12 00:00:00+07	2716.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6561	\N	1	2715	\N	\N	2021-08-10 00:00:00+07	2715.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6664	\N	1	1062	\N	\N	2021-08-16 00:00:00+07	1062.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-16 00:00:00+07	\N	08:00	19:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6665	\N	1	621	\N	\N	2021-08-16 00:00:00+07	0621.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6539	\N	1	2712	\N	\N	2021-08-05 00:00:00+07	2712.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-05 00:00:00+07	\N	09:00	21:00	150	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6537	\N	1	908	\N	\N	2021-07-21 00:00:00+07	0908.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-07-21 00:00:00+07	\N	08:00	20:00	60	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6519	\N	1	2515	\N	\N	2021-06-29 00:00:00+07	2515.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-29 00:00:00+07	\N	10:00	21:00	45	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6515	\N	1	2514	\N	\N	2021-06-24 00:00:00+07	2514.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-24 00:00:00+07	\N	09:30	21:30	50-100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6506	\N	1	615	\N	\N	2021-06-21 00:00:00+07	0615.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-21 00:00:00+07	\N	10:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6501	\N	1	2509	\N	\N	2021-06-15 00:00:00+07	2509.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-15 00:00:00+07	\N	11:30	20:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6498	\N	1	2708	\N	\N	2021-06-15 00:00:00+07	2708.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-15 00:00:00+07	\N	10:00	16:00	150	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6491	\N	1	2508	\N	\N	2021-06-09 00:00:00+07	2508.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-09 00:00:00+07	\N	09:00	22:00	30-50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6490	\N	1	612	\N	\N	2021-06-09 00:00:00+07	0612.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-09 00:00:00+07	\N	11:00	20:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6485	\N	1	2705	\N	\N	2021-06-03 00:00:00+07	2705.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6488	\N	1	2706	\N	\N	2021-06-07 00:00:00+07	2706.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-07 00:00:00+07	\N	11:00	20:30	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6489	\N	1	2507	\N	\N	2021-06-08 00:00:00+07	2507.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-08 00:00:00+07	\N	07:30	19:30	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6466	\N	1	2697	\N	\N	2021-05-05 00:00:00+07	2697.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-05-05 00:00:00+07	\N	09:00	21:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6464	\N	1	1055	\N	\N	2021-05-04 00:00:00+07	1055.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-01 00:00:00+07	\N	04:00	15:00	35	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6461	\N	1	890	\N	\N	2021-04-30 00:00:00+07	0890.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-30 00:00:00+07	\N	09:00	22:00	24	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6456	\N	1	2500	\N	\N	2021-04-26 00:00:00+07	2500.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-26 00:00:00+07	\N	10:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6453	\N	1	2498	\N	\N	2021-04-21 00:00:00+07	2498.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-21 00:00:00+07	\N	12:00	21:00	40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3939	\N	1	27	\N	\N	2017-08-16 00:00:00+07	0027.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3942	\N	1	2013	\N	\N	2017-08-21 00:00:00+07	2013.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6454	\N	1	2692	\N	\N	2021-04-22 00:00:00+07	2692.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6422	\N	1	2486	\N	\N	2021-04-08 00:00:00+07	2486.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-08 00:00:00+07	\N	08:00	22:00	35	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6419	\N	1	1051	\N	\N	2021-04-05 00:00:00+07	1051.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-05 00:00:00+07	\N	10:00	21:00	30 - 50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6418	\N	1	2485	\N	\N	2021-04-01 00:00:00+07	2485.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-01 00:00:00+07	\N	10:00	20:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6417	\N	1	2685	\N	\N	2021-03-31 00:00:00+07	2685.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-31 00:00:00+07	\N	09:00	21:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6416	\N	1	2484	\N	\N	2021-03-31 00:00:00+07	2484.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-31 00:00:00+07	\N	11:00	20:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6415	\N	1	2684	\N	\N	2021-03-30 00:00:00+07	2684.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-30 00:00:00+07	\N	11:00	21:00	200	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6413	\N	1	2483	\N	\N	2021-03-25 00:00:00+07	2483.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-25 00:00:00+07	\N	09:00	19:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6411	\N	1	2481	\N	\N	2021-03-24 00:00:00+07	2481.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3943	\N	1	481	\N	\N	2017-08-21 00:00:00+07	0481.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6409	\N	1	2683	\N	\N	2021-03-24 00:00:00+07	2683.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-24 00:00:00+07	\N	05:00	15:00	75	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6408	\N	1	2682	\N	\N	2021-03-24 00:00:00+07	2682.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-24 00:00:00+07	\N	09:00	23:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6406	\N	1	2478	\N	\N	2021-03-19 00:00:00+07	2478.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6405	\N	1	2681	\N	\N	2021-03-18 00:00:00+07	2681.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-18 00:00:00+07	\N	10:00	19:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2944	\N	1	1584	\N	\N	2017-02-01 00:00:00+07	1584.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6401	\N	1	2678	\N	\N	2021-03-16 00:00:00+07	2678.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-16 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6395	\N	1	1049	\N	\N	2021-03-05 00:00:00+07	1049.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-05 00:00:00+07	\N	08:00	18:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6392	\N	1	880	\N	\N	2021-03-04 00:00:00+07	0880.64.101	2	2021-12-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-04 00:00:00+07	\N	07:00	14:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6390	\N	1	879	\N	\N	2021-03-04 00:00:00+07	0879.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6389	\N	1	877	\N	\N	2021-03-04 00:00:00+07	0877.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-04 00:00:00+07	\N	11:00	21:30	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6387	\N	1	2675	\N	\N	2021-02-19 00:00:00+07	2675.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6386	\N	1	1047	\N	\N	2021-02-19 00:00:00+07	1047.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-19 00:00:00+07	\N	10:00	22:00	45	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6382	\N	1	2672	\N	\N	2021-02-11 00:00:00+07	2672.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-11 00:00:00+07	\N	11:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6381	\N	1	2671	\N	\N	2021-02-10 00:00:00+07	2671.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-10 00:00:00+07	\N	10:00	21:00	20 - 30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6376	\N	1	2669	\N	\N	2021-02-04 00:00:00+07	2669.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-04 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6374	\N	1	2668	\N	\N	2021-02-04 00:00:00+07	2668.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-04 00:00:00+07	\N	09:00	20:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6373	\N	1	2667	\N	\N	2021-02-04 00:00:00+07	2667.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-04 00:00:00+07	\N	09:00	20:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6372	\N	1	2666	\N	\N	2021-02-03 00:00:00+07	2666.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-03 00:00:00+07	\N	09:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6370	\N	1	2474	\N	\N	2021-02-03 00:00:00+07	2474.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-03 00:00:00+07	\N	10:00	12:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6364	\N	1	2472	\N	\N	2021-01-28 00:00:00+07	2472.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-28 00:00:00+07	\N	10:00	20:00	4	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6360	\N	1	1045	\N	\N	2021-01-26 00:00:00+07	1045.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-26 00:00:00+07	\N	10:00	20:00	75	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6361	\N	1	2660	\N	\N	2021-01-27 00:00:00+07	2660.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-27 00:00:00+07	\N	12:00	20:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6350	\N	1	2656	\N	\N	2021-01-15 00:00:00+07	2656.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-15 00:00:00+07	\N	09:00	21:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2928	\N	1	2026	\N	\N	2016-12-01 00:00:00+07	2026.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6346	\N	1	1043	\N	\N	2021-01-13 00:00:00+07	1043.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6344	\N	1	2655	\N	\N	2021-01-13 00:00:00+07	2655.62.101	2	2021-01-17 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-13 00:00:00+07	\N	10:00	21:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6343	\N	1	2654	\N	\N	2021-01-13 00:00:00+07	2654.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-13 00:00:00+07	\N	10:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6338	\N	1	2467	\N	\N	2021-01-11 00:00:00+07	2467.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-11 00:00:00+07	\N	10:00	21:45	45	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6335	\N	1	1040	\N	\N	2021-01-05 00:00:00+07	1040.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6330	\N	1	2464	\N	\N	2020-12-23 00:00:00+07	2464.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-23 00:00:00+07	\N	10:00	21:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6328	\N	1	2650	\N	\N	2020-12-15 00:00:00+07	2650.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6314	\N	1	596	\N	\N	2020-12-07 00:00:00+07	0596.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-07 00:00:00+07	\N	\N	\N	40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6292	\N	1	869	\N	\N	2020-12-03 00:00:00+07	0869.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6291	\N	1	2455	\N	\N	2020-12-03 00:00:00+07	2455.65.101	2	2021-01-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-03 00:00:00+07	\N	07:00	20:30	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6290	\N	1	594	\N	\N	2020-12-03 00:00:00+07	0594.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6283	\N	1	2645	\N	\N	2020-11-27 00:00:00+07	2645.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-27 00:00:00+07	\N	13:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6281	\N	1	593	\N	\N	2020-11-26 00:00:00+07	0593.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-26 00:00:00+07	\N	10:00	21:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6280	\N	1	2451	\N	\N	2020-11-26 00:00:00+07	2451.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3945	\N	1	2184	\N	\N	2017-08-21 00:00:00+07	2184.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6277	\N	1	2644	\N	\N	2020-11-23 00:00:00+07	2644.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-23 00:00:00+07	\N	10:00	21:00	12	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6273	\N	1	591	\N	\N	2020-11-19 00:00:00+07	0591.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6263	\N	1	2637	\N	\N	2020-11-06 00:00:00+07	2637.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6270	\N	1	2643	\N	\N	2020-11-13 00:00:00+07	2643.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-13 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6269	\N	1	2642	\N	\N	2020-11-12 00:00:00+07	2642.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6258	\N	1	2636	\N	\N	2020-11-03 00:00:00+07	2636.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6255	\N	1	2634	\N	\N	2020-11-02 00:00:00+07	2634.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-02 00:00:00+07	\N	10:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6268	\N	1	2641	\N	\N	2020-11-12 00:00:00+07	2641.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6267	\N	1	2640	\N	\N	2020-11-12 00:00:00+07	2640.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2927	\N	1	1518	\N	\N	2016-12-01 00:00:00+07	1518.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6266	\N	1	2639	\N	\N	2020-11-11 00:00:00+07	2639.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-11 00:00:00+07	8	10:00	21:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1681	\N	2	925	\N	\N	2011-02-08 00:00:00+07	0925.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3947	\N	1	770	\N	\N	2016-02-25 00:00:00+07	0770.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1683	\N	2	1288	\N	\N	2011-02-08 00:00:00+07	1288.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1684	\N	2	1287	\N	\N	2011-02-08 00:00:00+07	1287.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1686	\N	2	769	\N	\N	2011-02-08 00:00:00+07	0769.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1687	\N	2	1261	\N	\N	2011-02-08 00:00:00+07	1261.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1700	\N	2	1369	\N	\N	2011-02-08 00:00:00+07	1369.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1703	\N	2	1364	\N	\N	2011-02-08 00:00:00+07	1364.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1705	\N	2	1306	\N	\N	2011-02-08 00:00:00+07	1306.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1707	\N	2	1278	\N	\N	2011-02-08 00:00:00+07	1278.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1708	\N	2	1348	\N	\N	2011-02-08 00:00:00+07	1348.62.101	2	2020-09-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3948	\N	1	1643	\N	\N	2017-08-24 00:00:00+07	1643.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1712	\N	2	1319	\N	\N	2011-02-08 00:00:00+07	1319.62.101	1	\N	2020-05-31 00:00:00+07	OA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1713	\N	2	1391	\N	\N	2011-02-08 00:00:00+07	1391.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1715	\N	2	772	\N	\N	2011-02-08 00:00:00+07	0772.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1719	\N	2	1284	\N	\N	2011-02-08 00:00:00+07	1284.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1720	\N	2	1032	\N	\N	2011-02-08 00:00:00+07	1032.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6253	\N	1	2445	\N	\N	2020-10-16 00:00:00+07	2445.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1754	\N	2	193	\N	\N	2011-02-07 00:00:00+07	0193.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1755	\N	2	196	\N	\N	2011-02-07 00:00:00+07	0196.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1756	\N	2	205	\N	\N	2011-02-07 00:00:00+07	0205.60.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1757	\N	2	206	\N	\N	2011-02-07 00:00:00+07	0206.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1758	\N	2	208	\N	\N	2011-02-07 00:00:00+07	0208.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1759	\N	2	209	\N	\N	2011-02-07 00:00:00+07	0209.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1760	\N	2	210	\N	\N	2011-02-07 00:00:00+07	0210.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1761	\N	2	235	\N	\N	2011-02-07 00:00:00+07	0235.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1762	\N	2	249	\N	\N	2011-02-07 00:00:00+07	0249.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1763	\N	2	93	\N	\N	2011-02-07 00:00:00+07	0093.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3937	\N	1	1873	\N	\N	2015-05-01 00:00:00+07	1873.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1765	\N	2	99	\N	\N	2011-02-07 00:00:00+07	0099.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1766	\N	2	111	\N	\N	2011-02-07 00:00:00+07	0111.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1767	\N	2	112	\N	\N	2011-02-07 00:00:00+07	0112.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1768	\N	2	113	\N	\N	2011-02-07 00:00:00+07	0113.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1770	\N	2	125	\N	\N	2011-02-07 00:00:00+07	0125.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1771	\N	2	126	\N	\N	2011-02-07 00:00:00+07	0126.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1772	\N	2	131	\N	\N	2011-02-07 00:00:00+07	0131.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1773	\N	2	134	\N	\N	2011-02-07 00:00:00+07	0134.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1774	\N	2	943	\N	\N	2011-02-07 00:00:00+07	0943.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1775	\N	2	136	\N	\N	2011-02-07 00:00:00+07	0136.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-09 00:00:00+07	48	06:30	17:30	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1776	\N	2	137	\N	\N	2011-02-07 00:00:00+07	0137.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1777	\N	2	140	\N	\N	2011-02-07 00:00:00+07	0140.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3950	\N	1	474	\N	\N	2017-08-11 00:00:00+07	0474.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1778	\N	2	799	\N	\N	2011-02-07 00:00:00+07	0799.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1779	\N	2	803	\N	\N	2011-02-07 00:00:00+07	0803.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1780	\N	2	1021	\N	\N	2011-02-07 00:00:00+07	1021.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1781	\N	2	1103	\N	\N	2011-02-07 00:00:00+07	1103.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1782	\N	2	19	\N	\N	2011-02-07 00:00:00+07	0019.61.101	2	2020-02-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1783	\N	2	22	\N	\N	2011-02-07 00:00:00+07	0022.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3951	\N	1	488	\N	\N	2017-07-03 00:00:00+07	0488.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1785	\N	2	89	\N	\N	2011-02-07 00:00:00+07	0089.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1786	\N	2	90	\N	\N	2011-02-07 00:00:00+07	0090.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1787	\N	2	95	\N	\N	2011-02-07 00:00:00+07	0095.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1788	\N	2	96	\N	\N	2011-02-07 00:00:00+07	0096.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1789	\N	2	101	\N	\N	2011-02-07 00:00:00+07	0101.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1790	\N	2	102	\N	\N	2011-02-07 00:00:00+07	0102.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1791	\N	2	104	\N	\N	2011-02-07 00:00:00+07	0104.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1792	\N	2	106	\N	\N	2011-02-07 00:00:00+07	0106.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1793	\N	2	107	\N	\N	2011-02-07 00:00:00+07	0107.62.101	1	2017-05-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1794	\N	2	119	\N	\N	2011-02-07 00:00:00+07	0119.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1795	\N	2	120	\N	\N	2011-02-07 00:00:00+07	0120.62.101	1	2017-05-16 00:00:00+07	2020-08-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1796	\N	2	805	\N	\N	2011-02-07 00:00:00+07	0805.62.101	2	2017-01-01 00:00:00+07	\N	SA	\N	10112	solusiti	\N	\N	f	f	0001-01-20 00:07:12+07:07:12	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1797	\N	2	924	\N	\N	2011-02-07 00:00:00+07	0924.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1798	\N	2	141	\N	\N	2011-02-07 00:00:00+07	0141.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1799	\N	2	142	\N	\N	2011-02-07 00:00:00+07	0142.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1800	\N	2	143	\N	\N	2011-02-07 00:00:00+07	0143.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1801	\N	2	144	\N	\N	2011-02-07 00:00:00+07	0144.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1802	\N	2	146	\N	\N	2011-02-07 00:00:00+07	0146.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1803	\N	2	127	\N	\N	2011-02-07 00:00:00+07	0127.62.101	1	2017-07-07 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1804	\N	2	175	\N	\N	2011-02-07 00:00:00+07	0175.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6252	\N	1	2444	\N	\N	2020-10-16 00:00:00+07	2444.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1806	\N	2	169	\N	\N	2011-02-07 00:00:00+07	0169.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1807	\N	2	957	\N	\N	2011-02-07 00:00:00+07	0957.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1808	\N	2	952	\N	\N	2011-02-07 00:00:00+07	0952.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1809	\N	2	327	\N	\N	2011-02-07 00:00:00+07	0327.61.101	2	2019-04-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1810	\N	2	10	\N	\N	2011-02-07 00:00:00+07	0010.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	06:00	13:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1811	\N	2	27	\N	\N	2011-02-07 00:00:00+07	0027.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1812	\N	2	316	\N	\N	2011-02-07 00:00:00+07	0316.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1813	\N	2	474	\N	\N	2011-02-07 00:00:00+07	0474.62.101	1	2019-10-14 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1814	\N	2	433	\N	\N	2011-02-07 00:00:00+07	0433.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1815	\N	2	186	\N	\N	2011-02-07 00:00:00+07	0186.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1816	\N	2	1209	\N	\N	2011-02-07 00:00:00+07	1209.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1817	\N	2	1282	\N	\N	2011-02-07 00:00:00+07	1282.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1818	\N	2	1302	\N	\N	2011-02-07 00:00:00+07	1302.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1819	\N	2	200	\N	\N	2011-02-07 00:00:00+07	0200.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1820	\N	2	468	\N	\N	2011-02-07 00:00:00+07	0468.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1821	\N	2	1315	\N	\N	2011-02-07 00:00:00+07	1315.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3766	\N	1	841	\N	\N	2017-06-05 00:00:00+07	0841.61.101	2	2017-07-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1823	\N	2	472	\N	\N	2011-02-07 00:00:00+07	0472.61.101	1	2017-05-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1824	\N	2	474	\N	\N	2011-02-07 00:00:00+07	0474.61.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1825	\N	2	1349	\N	\N	2011-02-07 00:00:00+07	1349.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1826	\N	2	475	\N	\N	2011-02-07 00:00:00+07	0475.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1827	\N	2	477	\N	\N	2011-02-07 00:00:00+07	0477.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1828	\N	2	478	\N	\N	2011-02-07 00:00:00+07	0478.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1829	\N	2	480	\N	\N	2011-02-07 00:00:00+07	0480.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1830	\N	2	11	\N	\N	2011-02-07 00:00:00+07	0011.61.101	2	2020-02-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1831	\N	2	18	\N	\N	2011-02-07 00:00:00+07	0018.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1832	\N	2	481	\N	\N	2011-02-07 00:00:00+07	0481.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1833	\N	2	483	\N	\N	2011-02-07 00:00:00+07	0483.61.101	1	2017-05-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1835	\N	2	1368	\N	\N	2011-02-07 00:00:00+07	1368.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1836	\N	2	1186	\N	\N	2011-02-07 00:00:00+07	1186.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6251	\N	1	2443	\N	\N	2020-10-16 00:00:00+07	2443.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1838	\N	2	1102	\N	\N	2011-02-07 00:00:00+07	1102.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1839	\N	2	956	\N	\N	2011-02-07 00:00:00+07	0956.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1840	\N	2	490	\N	\N	2011-02-07 00:00:00+07	0490.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1996	\N	2	159	\N	\N	2011-02-08 00:00:00+07	0159.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1841	\N	2	129	\N	\N	2011-02-07 00:00:00+07	0129.62.101	1	\N	2021-04-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1842	\N	2	130	\N	\N	2011-02-07 00:00:00+07	0130.62.101	2	2020-01-24 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1843	\N	2	475	\N	\N	2011-02-07 00:00:00+07	0475.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1844	\N	2	929	\N	\N	2011-02-07 00:00:00+07	0929.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1845	\N	2	310	\N	\N	2011-02-07 00:00:00+07	0310.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1846	\N	2	12	\N	\N	2011-02-07 00:00:00+07	0012.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1847	\N	2	5	\N	\N	2011-02-07 00:00:00+07	0005.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1848	\N	2	2	\N	\N	2011-02-07 00:00:00+07	0002.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1849	\N	2	31	\N	\N	2011-02-07 00:00:00+07	0031.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1851	\N	2	8	\N	\N	2011-02-07 00:00:00+07	0008.61.101	2	2019-03-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1852	\N	2	7	\N	\N	2011-02-07 00:00:00+07	0007.61.101	1	\N	2018-06-27 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1853	\N	2	146	\N	\N	2011-02-07 00:00:00+07	0146.65.101	2	2021-07-28 00:00:00+07	\N	OA	\N	10112	solusiti	\N	\N	f	f	2021-07-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1855	\N	2	353	\N	\N	2011-02-07 00:00:00+07	0353.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1856	\N	2	317	\N	\N	2011-02-07 00:00:00+07	0317.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1857	\N	2	318	\N	\N	2011-02-07 00:00:00+07	0318.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1858	\N	2	320	\N	\N	2011-02-07 00:00:00+07	0320.65.101	2	2020-03-01 00:00:00+07	\N	OA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-07-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1859	\N	2	319	\N	\N	2011-02-07 00:00:00+07	0319.65.101	2	2019-10-16 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1860	\N	2	133	\N	\N	2011-02-07 00:00:00+07	0133.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1861	\N	2	40	\N	\N	2011-02-07 00:00:00+07	0040.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1862	\N	2	1206	\N	\N	2011-02-07 00:00:00+07	1206.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1863	\N	2	490	\N	\N	2011-02-07 00:00:00+07	0490.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1864	\N	2	200	\N	\N	2011-02-07 00:00:00+07	0200.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1865	\N	2	1305	\N	\N	2011-02-07 00:00:00+07	1305.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1866	\N	2	1317	\N	\N	2011-02-07 00:00:00+07	1317.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1867	\N	2	494	\N	\N	2011-02-07 00:00:00+07	0494.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1868	\N	2	497	\N	\N	2011-02-07 00:00:00+07	0497.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3963	\N	1	366	\N	\N	2017-08-29 00:00:00+07	0366.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1870	\N	2	495	\N	\N	2011-02-07 00:00:00+07	0495.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1871	\N	2	245	\N	\N	2011-02-07 00:00:00+07	0245.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1872	\N	2	499	\N	\N	2011-02-07 00:00:00+07	0499.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1873	\N	2	219	\N	\N	2011-02-07 00:00:00+07	0219.64.101	2	2018-07-18 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1874	\N	2	504	\N	\N	2011-02-07 00:00:00+07	0504.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1875	\N	2	503	\N	\N	2011-02-07 00:00:00+07	0503.64.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1876	\N	2	233	\N	\N	2011-02-07 00:00:00+07	0233.64.101	2	2018-07-18 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1877	\N	2	1372	\N	\N	2011-02-07 00:00:00+07	1372.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1878	\N	2	1370	\N	\N	2011-02-07 00:00:00+07	1370.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1879	\N	2	221	\N	\N	2011-02-07 00:00:00+07	0221.64.101	2	2018-07-18 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1880	\N	2	10	\N	\N	2011-02-07 00:00:00+07	0010.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1881	\N	2	177	\N	\N	2011-02-07 00:00:00+07	0177.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1882	\N	2	2	\N	\N	2011-02-07 00:00:00+07	0002.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1883	\N	2	139	\N	\N	2011-02-07 00:00:00+07	0139.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1884	\N	2	1198	\N	\N	2011-02-07 00:00:00+07	1198.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1885	\N	2	804	\N	\N	2011-02-07 00:00:00+07	0804.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1886	\N	2	476	\N	\N	2011-02-07 00:00:00+07	0476.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1887	\N	2	1	\N	\N	2011-02-08 00:00:00+07	0001.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1888	\N	2	146	\N	\N	2011-02-08 00:00:00+07	0146.63.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1889	\N	2	126	\N	\N	2011-02-08 00:00:00+07	0126.63.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3964	\N	1	1646	\N	\N	2017-08-30 00:00:00+07	1646.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1891	\N	2	10	\N	\N	2011-02-08 00:00:00+07	0010.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1892	\N	2	9	\N	\N	2011-02-08 00:00:00+07	0009.63.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1893	\N	2	6	\N	\N	2011-02-08 00:00:00+07	0006.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1894	\N	2	5	\N	\N	2011-02-08 00:00:00+07	0005.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1895	\N	2	4	\N	\N	2011-02-08 00:00:00+07	0004.63.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1896	\N	2	138	\N	\N	2011-02-08 00:00:00+07	0138.63.101	2	2019-10-30 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1897	\N	2	136	\N	\N	2011-02-08 00:00:00+07	0136.63.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1898	\N	2	137	\N	\N	2011-02-08 00:00:00+07	0137.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1899	\N	2	135	\N	\N	2011-02-08 00:00:00+07	0135.63.101	1	2020-01-02 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1900	\N	2	134	\N	\N	2011-02-08 00:00:00+07	0134.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1901	\N	2	133	\N	\N	2011-02-08 00:00:00+07	0133.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1902	\N	2	128	\N	\N	2011-02-08 00:00:00+07	0128.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1903	\N	2	13	\N	\N	2011-02-08 00:00:00+07	0013.63.101	2	2018-12-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1904	\N	2	14	\N	\N	2011-02-08 00:00:00+07	0014.63.101	1	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1905	\N	2	1351	\N	\N	2011-02-08 00:00:00+07	1351.62.101	2	2018-05-02 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1906	\N	2	154	\N	\N	2011-02-08 00:00:00+07	0154.63.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1907	\N	2	156	\N	\N	2011-02-08 00:00:00+07	0156.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1908	\N	2	157	\N	\N	2011-02-08 00:00:00+07	0157.63.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1909	\N	2	178	\N	\N	2011-02-08 00:00:00+07	0178.62.101	1	2017-07-07 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1910	\N	2	132	\N	\N	2011-02-08 00:00:00+07	0132.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1911	\N	2	15	\N	\N	2011-02-08 00:00:00+07	0015.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1912	\N	2	415	\N	\N	2011-02-08 00:00:00+07	0415.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1913	\N	2	1197	\N	\N	2011-02-08 00:00:00+07	1197.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1914	\N	2	754	\N	\N	2011-02-08 00:00:00+07	0754.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1915	\N	2	1022	\N	\N	2011-02-08 00:00:00+07	1022.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1916	\N	2	758	\N	\N	2011-02-08 00:00:00+07	0758.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1917	\N	2	771	\N	\N	2011-02-08 00:00:00+07	0771.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1918	\N	2	761	\N	\N	2011-02-08 00:00:00+07	0761.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1919	\N	2	179	\N	\N	2011-02-08 00:00:00+07	0179.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1920	\N	2	189	\N	\N	2011-02-08 00:00:00+07	0189.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1921	\N	2	774	\N	\N	2011-02-08 00:00:00+07	0774.62.101	2	2018-01-23 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1922	\N	2	1030	\N	\N	2011-02-08 00:00:00+07	1030.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1923	\N	2	1031	\N	\N	2011-02-08 00:00:00+07	1031.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1924	\N	2	1032	\N	\N	2011-02-08 00:00:00+07	1032.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1925	\N	2	744	\N	\N	2011-02-08 00:00:00+07	0744.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1926	\N	2	789	\N	\N	2011-02-08 00:00:00+07	0789.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1927	\N	2	1023	\N	\N	2011-02-08 00:00:00+07	1023.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1928	\N	2	794	\N	\N	2011-02-08 00:00:00+07	0794.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1929	\N	2	734	\N	\N	2011-02-08 00:00:00+07	0734.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1930	\N	2	742	\N	\N	2011-02-08 00:00:00+07	0742.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1931	\N	2	743	\N	\N	2011-02-08 00:00:00+07	0743.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1932	\N	2	748	\N	\N	2011-02-08 00:00:00+07	0748.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1933	\N	2	749	\N	\N	2011-02-08 00:00:00+07	0749.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1934	\N	2	752	\N	\N	2011-02-08 00:00:00+07	0752.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1935	\N	2	753	\N	\N	2011-02-08 00:00:00+07	0753.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1936	\N	2	1019	\N	\N	2011-02-08 00:00:00+07	1019.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1937	\N	2	1036	\N	\N	2011-02-08 00:00:00+07	1036.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1938	\N	2	765	\N	\N	2011-02-08 00:00:00+07	0765.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1939	\N	2	781	\N	\N	2011-02-08 00:00:00+07	0781.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1940	\N	2	787	\N	\N	2011-02-08 00:00:00+07	0787.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3965	\N	1	1647	\N	\N	2017-08-30 00:00:00+07	1647.65.101	2	2019-06-04 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1942	\N	2	1033	\N	\N	2011-02-08 00:00:00+07	1033.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1943	\N	2	778	\N	\N	2011-02-08 00:00:00+07	0778.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1944	\N	2	1303	\N	\N	2011-02-08 00:00:00+07	1303.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1945	\N	2	1379	\N	\N	2011-02-08 00:00:00+07	1379.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1947	\N	2	1380	\N	\N	2011-02-08 00:00:00+07	1380.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1948	\N	2	1381	\N	\N	2011-02-08 00:00:00+07	1381.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1949	\N	2	194	\N	\N	2011-02-08 00:00:00+07	0194.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1950	\N	2	224	\N	\N	2011-02-08 00:00:00+07	0224.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1951	\N	2	253	\N	\N	2011-02-08 00:00:00+07	0253.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1952	\N	2	45	\N	\N	2011-02-08 00:00:00+07	0045.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1953	\N	2	191	\N	\N	2011-02-08 00:00:00+07	0191.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1954	\N	2	207	\N	\N	2011-02-08 00:00:00+07	0207.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1956	\N	2	243	\N	\N	2011-02-08 00:00:00+07	0243.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	06:30	15:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1957	\N	2	246	\N	\N	2011-02-08 00:00:00+07	0246.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1958	\N	2	256	\N	\N	2011-02-08 00:00:00+07	0256.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1959	\N	2	605	\N	\N	2011-02-08 00:00:00+07	0605.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1960	\N	2	606	\N	\N	2011-02-08 00:00:00+07	0606.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1961	\N	2	604	\N	\N	2011-02-08 00:00:00+07	0604.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1962	\N	2	948	\N	\N	2011-02-08 00:00:00+07	0948.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1963	\N	2	82	\N	\N	2011-02-08 00:00:00+07	0082.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1964	\N	2	7	\N	\N	2011-02-08 00:00:00+07	0007.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1965	\N	2	609	\N	\N	2011-02-08 00:00:00+07	0609.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1966	\N	2	91	\N	\N	2011-02-08 00:00:00+07	0091.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1967	\N	2	92	\N	\N	2011-02-08 00:00:00+07	0092.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1968	\N	2	97	\N	\N	2011-02-08 00:00:00+07	0097.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1969	\N	2	117	\N	\N	2011-02-08 00:00:00+07	0117.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1970	\N	2	945	\N	\N	2011-02-08 00:00:00+07	0945.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1971	\N	2	949	\N	\N	2011-02-08 00:00:00+07	0949.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1973	\N	2	1265	\N	\N	2011-02-08 00:00:00+07	1265.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1974	\N	2	171	\N	\N	2011-02-08 00:00:00+07	0171.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1975	\N	2	103	\N	\N	2011-02-08 00:00:00+07	0103.62.101	1	\N	2018-06-27 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1977	\N	2	174	\N	\N	2011-02-08 00:00:00+07	0174.62.101	2	2019-08-25 00:00:00+07	\N	OA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-05-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1978	\N	2	194	\N	\N	2011-02-08 00:00:00+07	0194.64.101	1	\N	2019-01-01 00:00:00+07	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1979	\N	2	1264	\N	\N	2011-02-08 00:00:00+07	1264.62.101	2	2019-10-16 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1980	\N	2	919	\N	\N	2011-02-08 00:00:00+07	0919.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1981	\N	2	1300	\N	\N	2011-02-08 00:00:00+07	1300.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1982	\N	2	199	\N	\N	2011-02-09 00:00:00+07	0199.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1983	\N	2	156	\N	\N	2011-02-08 00:00:00+07	0156.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1984	\N	2	172	\N	\N	2011-02-08 00:00:00+07	0172.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1985	\N	2	1363	\N	\N	2011-02-08 00:00:00+07	1363.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1986	\N	2	1359	\N	\N	2011-02-08 00:00:00+07	1359.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2076	\N	2	749	\N	\N	2011-02-09 00:00:00+07	0749.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1987	\N	2	230	\N	\N	2011-02-08 00:00:00+07	0230.64.101	1	2019-10-16 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1988	\N	2	1360	\N	\N	2011-02-08 00:00:00+07	1360.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1989	\N	2	1361	\N	\N	2011-02-08 00:00:00+07	1361.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1990	\N	2	220	\N	\N	2011-02-08 00:00:00+07	0220.64.101	2	2018-07-18 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1991	\N	2	1365	\N	\N	2011-02-08 00:00:00+07	1365.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1992	\N	2	1366	\N	\N	2011-02-08 00:00:00+07	1366.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1993	\N	2	1371	\N	\N	2011-02-08 00:00:00+07	1371.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1994	\N	2	232	\N	\N	2011-02-08 00:00:00+07	0232.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1995	\N	2	148	\N	\N	2011-02-08 00:00:00+07	0148.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1997	\N	2	160	\N	\N	2011-02-08 00:00:00+07	0160.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1998	\N	2	163	\N	\N	2011-02-08 00:00:00+07	0163.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1999	\N	2	165	\N	\N	2011-02-08 00:00:00+07	0165.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2000	\N	2	166	\N	\N	2011-02-08 00:00:00+07	0166.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2001	\N	2	167	\N	\N	2011-02-09 00:00:00+07	0167.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2002	\N	2	1029	\N	\N	2011-02-09 00:00:00+07	1029.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2003	\N	2	1378	\N	\N	2011-02-09 00:00:00+07	1378.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2004	\N	2	158	\N	\N	2011-02-09 00:00:00+07	0158.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2005	\N	2	164	\N	\N	2011-02-09 00:00:00+07	0164.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2006	\N	2	24	\N	\N	2011-02-09 00:00:00+07	0024.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2007	\N	2	419	\N	\N	2011-02-09 00:00:00+07	0419.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2008	\N	2	413	\N	\N	2011-02-09 00:00:00+07	0413..65	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6250	\N	1	2633	\N	\N	2020-10-16 00:00:00+07	2633.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-16 00:00:00+07	\N	10:00	17:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2010	\N	2	325	\N	\N	2011-02-09 00:00:00+07	0325.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2011	\N	2	9	\N	\N	2011-02-09 00:00:00+07	0009.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2012	\N	2	315	\N	\N	2011-02-09 00:00:00+07	0315.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2013	\N	2	421	\N	\N	2011-02-09 00:00:00+07	0421.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2014	\N	2	199	\N	\N	2011-02-09 00:00:00+07	0199.65.101	2	2019-01-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2015	\N	2	16	\N	\N	2011-02-09 00:00:00+07	0016.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2016	\N	2	34	\N	\N	2011-02-09 00:00:00+07	0034.65.101	2	2018-03-21 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2017	\N	2	88	\N	\N	2011-02-09 00:00:00+07	0088.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2019	\N	2	28	\N	\N	2011-02-09 00:00:00+07	0028.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2020	\N	2	108	\N	\N	2011-02-09 00:00:00+07	0108.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2021	\N	2	96	\N	\N	2011-02-09 00:00:00+07	0096.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2022	\N	2	44	\N	\N	2011-02-09 00:00:00+07	0044.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2023	\N	2	404	\N	\N	2011-02-09 00:00:00+07	0404.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2024	\N	2	190	\N	\N	2011-02-09 00:00:00+07	0190.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2025	\N	2	1194	\N	\N	2011-02-09 00:00:00+07	1194.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2026	\N	2	331	\N	\N	2011-02-09 00:00:00+07	0331.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2027	\N	2	1109	\N	\N	2011-02-09 00:00:00+07	1109.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2028	\N	2	47	\N	\N	2011-02-09 00:00:00+07	0047.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2029	\N	2	34	\N	\N	2011-02-09 00:00:00+07	0034.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2030	\N	2	4	\N	\N	2011-02-09 00:00:00+07	0004.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2031	\N	2	187	\N	\N	2011-02-09 00:00:00+07	0187.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2032	\N	2	39	\N	\N	2011-02-09 00:00:00+07	0039.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2033	\N	2	953	\N	\N	2011-02-09 00:00:00+07	0953.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2034	\N	2	45	\N	\N	2011-02-09 00:00:00+07	0045.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2035	\N	2	41	\N	\N	2011-02-09 00:00:00+07	0041.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2036	\N	2	15	\N	\N	2011-02-09 00:00:00+07	0015.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2037	\N	2	20	\N	\N	2011-02-09 00:00:00+07	0020.65.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2038	\N	2	10	\N	\N	2011-02-09 00:00:00+07	0010.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3966	\N	1	1648	\N	\N	2017-08-30 00:00:00+07	1648.65.101	2	2019-10-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2040	\N	2	23	\N	\N	2011-02-09 00:00:00+07	0023.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2041	\N	2	321	\N	\N	2011-02-09 00:00:00+07	0321.65.101	1	2019-03-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2042	\N	2	18	\N	\N	2011-02-09 00:00:00+07	0018.65.101	1	2017-07-07 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2043	\N	2	13	\N	\N	2011-02-09 00:00:00+07	0013.61.101	1	2017-07-07 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2044	\N	2	26	\N	\N	2011-02-09 00:00:00+07	0026.61.101	1	2017-07-07 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2045	\N	2	25	\N	\N	2011-02-09 00:00:00+07	0025.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2046	\N	2	180	\N	\N	2011-02-09 00:00:00+07	0180.62.101	1	\N	2019-06-07 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2047	\N	2	239	\N	\N	2011-02-09 00:00:00+07	0239.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2048	\N	2	237	\N	\N	2011-02-09 00:00:00+07	0237.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2049	\N	2	183	\N	\N	2011-02-09 00:00:00+07	0183.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2050	\N	2	54	\N	\N	2011-02-09 00:00:00+07	0054.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2051	\N	2	461	\N	\N	2011-02-09 00:00:00+07	0461.61.101	1	2017-07-07 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2052	\N	2	458	\N	\N	2011-02-09 00:00:00+07	0458.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2053	\N	2	457	\N	\N	2011-02-09 00:00:00+07	0457.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2054	\N	2	486	\N	\N	2011-02-09 00:00:00+07	0486.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2055	\N	2	508	\N	\N	2011-02-09 00:00:00+07	0508.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2056	\N	2	489	\N	\N	2011-02-09 00:00:00+07	0489.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2057	\N	2	488	\N	\N	2011-02-09 00:00:00+07	0488.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2058	\N	2	459	\N	\N	2011-02-09 00:00:00+07	0459.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2059	\N	2	30	\N	\N	2011-02-09 00:00:00+07	0030.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2060	\N	2	466	\N	\N	2011-02-09 00:00:00+07	0466.61.101	2	2019-12-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2061	\N	2	352	\N	\N	2011-02-09 00:00:00+07	0352.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2062	\N	2	498	\N	\N	2011-02-09 00:00:00+07	0498.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2063	\N	2	491	\N	\N	2011-02-09 00:00:00+07	0491.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2064	\N	2	496	\N	\N	2011-02-09 00:00:00+07	0496.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2065	\N	2	1356	\N	\N	2011-02-09 00:00:00+07	1356.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2066	\N	2	500	\N	\N	2011-02-09 00:00:00+07	0500.65.101	1	\N	2020-07-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2067	\N	2	412	\N	\N	2011-02-09 00:00:00+07	0412.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2068	\N	2	505	\N	\N	2011-02-09 00:00:00+07	0505.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2069	\N	2	507	\N	\N	2011-02-09 00:00:00+07	0507.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-02-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6249	\N	1	587	\N	\N	2020-10-15 00:00:00+07	0587.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2071	\N	2	502	\N	\N	2011-02-09 00:00:00+07	0502.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2072	\N	2	1308	\N	\N	2011-02-09 00:00:00+07	1308.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2074	\N	2	1301	\N	\N	2011-02-09 00:00:00+07	1301.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2075	\N	2	197	\N	\N	2011-02-09 00:00:00+07	0197.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2077	\N	2	1353	\N	\N	2011-02-09 00:00:00+07	1353.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2078	\N	2	1186	\N	\N	2011-02-09 00:00:00+07	1186.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2079	\N	2	1270	\N	\N	2011-02-09 00:00:00+07	1270.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6248	\N	1	2442	\N	\N	2020-10-15 00:00:00+07	2442.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-15 00:00:00+07	\N	11:00	20:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2081	\N	2	470	\N	\N	2011-02-09 00:00:00+07	0470.61.101	2	2017-06-07 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2083	\N	2	473	\N	\N	2011-02-09 00:00:00+07	0473.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2085	\N	2	1	\N	\N	2011-02-09 00:00:00+07	0001.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2086	\N	2	233	\N	\N	2011-02-09 00:00:00+07	0233.62.101	2	2018-08-31 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2087	\N	2	219	\N	\N	2011-02-09 00:00:00+07	0219.62.101	1	\N	2020-04-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2088	\N	2	1281	\N	\N	2011-02-09 00:00:00+07	1281.62.101	2	2019-01-30 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2089	\N	2	240	\N	\N	2011-02-09 00:00:00+07	0240.62.101	1	\N	2020-04-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2090	\N	2	323	\N	\N	2011-02-09 00:00:00+07	0323.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3764	\N	1	348	\N	\N	2017-06-05 00:00:00+07	0348.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2093	\N	2	228	\N	\N	2011-02-09 00:00:00+07	0228.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2094	\N	2	1192	\N	\N	2011-02-09 00:00:00+07	1192.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2095	\N	2	42	\N	\N	2011-02-09 00:00:00+07	0042.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2096	\N	2	241	\N	\N	2011-02-09 00:00:00+07	0241.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2097	\N	2	46	\N	\N	2011-02-09 00:00:00+07	0046.61.101	2	2020-04-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2098	\N	2	220	\N	\N	2011-02-09 00:00:00+07	0220.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2100	\N	2	44	\N	\N	2011-02-09 00:00:00+07	0044.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2101	\N	2	739	\N	\N	2011-02-09 00:00:00+07	0739.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2103	\N	2	262	\N	\N	2011-02-09 00:00:00+07	0262.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2105	\N	2	222	\N	\N	2011-02-09 00:00:00+07	0222.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2106	\N	2	12	\N	\N	2011-02-09 00:00:00+07	0012.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2107	\N	2	431	\N	\N	2011-02-09 00:00:00+07	0431.62.101	2	2017-04-03 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2108	\N	2	85	\N	\N	2011-02-09 00:00:00+07	0085.65.101	2	2017-04-03 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2111	\N	2	195	\N	\N	2011-02-09 00:00:00+07	0195.62.101	2	2018-02-25 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2112	\N	2	241	\N	\N	2011-02-09 00:00:00+07	0241.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2113	\N	2	48	\N	\N	2011-02-09 00:00:00+07	0048.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2114	\N	2	218	\N	\N	2011-02-09 00:00:00+07	0218.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2115	\N	2	11	\N	\N	2011-02-09 00:00:00+07	0011.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2116	\N	2	215	\N	\N	2011-02-09 00:00:00+07	0215.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2117	\N	2	1269	\N	\N	2011-02-09 00:00:00+07	1269.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2118	\N	2	731	\N	\N	2011-02-09 00:00:00+07	0731.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2119	\N	2	38	\N	\N	2011-02-09 00:00:00+07	0038.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2120	\N	2	211	\N	\N	2011-02-09 00:00:00+07	0211.62.101	1	2017-07-24 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2121	\N	2	255	\N	\N	2011-02-09 00:00:00+07	0255.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2122	\N	2	225	\N	\N	2011-02-09 00:00:00+07	0225.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2124	\N	2	201	\N	\N	2011-02-09 00:00:00+07	0201.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2125	\N	2	232	\N	\N	2011-02-09 00:00:00+07	0232.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2126	\N	2	750	\N	\N	2011-02-09 00:00:00+07	0750.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2127	\N	2	1268	\N	\N	2011-02-09 00:00:00+07	1268.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2128	\N	2	230	\N	\N	2011-02-09 00:00:00+07	0230.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2129	\N	2	231	\N	\N	2011-02-09 00:00:00+07	0231.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2130	\N	2	199	\N	\N	2011-02-09 00:00:00+07	0199.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2131	\N	2	1286	\N	\N	2011-02-09 00:00:00+07	1286.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2132	\N	2	416	\N	\N	2011-02-09 00:00:00+07	0416.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2133	\N	2	1307	\N	\N	2011-02-09 00:00:00+07	1307.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2134	\N	2	1207	\N	\N	2011-02-09 00:00:00+07	1207.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2135	\N	2	1354	\N	\N	2011-02-09 00:00:00+07	1354.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2136	\N	2	20	\N	\N	2011-02-09 00:00:00+07	0020.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2137	\N	2	100	\N	\N	2011-02-09 00:00:00+07	0100.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2138	\N	2	1383	\N	\N	2011-02-09 00:00:00+07	1383.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2139	\N	2	572	\N	\N	2011-08-01 00:00:00+07	0572.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2140	\N	2	573	\N	\N	2011-08-01 00:00:00+07	0573.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2141	\N	2	22	\N	\N	2011-02-10 00:00:00+07	0022.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2142	\N	2	236	\N	\N	2011-02-10 00:00:00+07	0236.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2143	\N	2	43	\N	\N	2011-02-10 00:00:00+07	0043.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2144	\N	2	2803	\N	\N	2011-02-10 00:00:00+07	2803.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2146	\N	2	212	\N	\N	2011-02-10 00:00:00+07	0212.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2148	\N	2	11	\N	\N	2011-02-10 00:00:00+07	0011.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2150	\N	2	1388	\N	\N	2011-02-11 00:00:00+07	1388.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2151	\N	2	1389	\N	\N	2011-02-11 00:00:00+07	1389.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2152	\N	2	158	\N	\N	2011-02-11 00:00:00+07	0158.63.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2153	\N	2	509	\N	\N	2011-02-11 00:00:00+07	0509.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2154	\N	2	510	\N	\N	2011-02-11 00:00:00+07	0510.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2155	\N	2	511	\N	\N	2011-02-11 00:00:00+07	0511.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2156	\N	2	512	\N	\N	2011-02-11 00:00:00+07	0512.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2157	\N	2	495	\N	\N	2011-02-11 00:00:00+07	0495.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2158	\N	2	1384	\N	\N	2011-02-11 00:00:00+07	1384.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2159	\N	2	234	\N	\N	2011-02-11 00:00:00+07	0234.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2160	\N	2	204	\N	\N	2011-02-11 00:00:00+07	0204.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2161	\N	2	935	\N	\N	2011-02-11 00:00:00+07	0935.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2163	\N	2	739	\N	\N	2011-02-11 00:00:00+07	0739.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2165	\N	2	1273	\N	\N	2011-02-11 00:00:00+07	1273.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2166	\N	2	1274	\N	\N	2011-02-11 00:00:00+07	1274.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2167	\N	2	251	\N	\N	2011-02-11 00:00:00+07	0251.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2169	\N	2	1393	\N	\N	2011-02-11 00:00:00+07	1393.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2170	\N	2	513	\N	\N	2011-02-11 00:00:00+07	0513.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2171	\N	2	514	\N	\N	2011-02-11 00:00:00+07	0514.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2172	\N	2	515	\N	\N	2011-02-11 00:00:00+07	0515.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2174	\N	2	517	\N	\N	2011-02-11 00:00:00+07	0517.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2175	\N	2	56	\N	\N	2011-02-11 00:00:00+07	0056.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2176	\N	2	235	\N	\N	2011-02-11 00:00:00+07	0235.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2177	\N	2	1395	\N	\N	2011-02-16 00:00:00+07	1395.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2178	\N	2	518	\N	\N	2011-02-16 00:00:00+07	0518.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2179	\N	2	227	\N	\N	2011-02-16 00:00:00+07	0227.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2180	\N	2	1396	\N	\N	2011-02-16 00:00:00+07	1396.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2181	\N	2	159	\N	\N	2011-02-16 00:00:00+07	0159.63.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2182	\N	2	1394	\N	\N	2011-02-16 00:00:00+07	1394.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2183	\N	2	221	\N	\N	2011-02-18 00:00:00+07	0221.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2185	\N	2	496	\N	\N	2011-02-21 00:00:00+07	0496.61.101	2	2021-04-01 00:00:00+07	\N	OA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-07-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2186	\N	2	410	\N	\N	2011-02-21 00:00:00+07	0410.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2187	\N	2	215	\N	\N	2011-02-21 00:00:00+07	0215.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2191	\N	2	1399	\N	\N	2011-02-23 00:00:00+07	1399.62.101	2	2019-10-17 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2192	\N	2	1528	\N	\N	2011-01-31 00:00:00+07	1528.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2193	\N	2	935	\N	\N	2011-02-24 00:00:00+07	0935.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2194	\N	2	792	\N	\N	2011-03-21 00:00:00+07	0792.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2197	\N	2	1355	\N	\N	2011-03-23 00:00:00+07	1355.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2198	\N	2	516	\N	\N	2011-03-23 00:00:00+07	0516.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2200	\N	2	179	\N	\N	2011-03-25 00:00:00+07	0179.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2201	\N	2	328	\N	\N	2011-03-28 00:00:00+07	0328.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2202	\N	2	346	\N	\N	2011-03-29 00:00:00+07	0346.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2205	\N	2	579	\N	\N	2011-08-01 00:00:00+07	0579.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2206	\N	2	1542	\N	\N	2011-02-17 00:00:00+07	1542.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2207	\N	2	1543	\N	\N	2011-02-21 00:00:00+07	1543.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2208	\N	2	482	\N	\N	2011-02-22 00:00:00+07	0482.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2210	\N	2	1541	\N	\N	2011-03-02 00:00:00+07	1541.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2211	\N	2	1546	\N	\N	2011-03-03 00:00:00+07	1546.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2212	\N	2	1556	\N	\N	2011-03-15 00:00:00+07	1556.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2213	\N	2	1549	\N	\N	2011-03-09 00:00:00+07	1549.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2214	\N	2	1553	\N	\N	2011-03-17 00:00:00+07	1553.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2215	\N	2	1555	\N	\N	2011-03-21 00:00:00+07	1555.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2216	\N	2	1557	\N	\N	2011-03-23 00:00:00+07	1557.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2218	\N	2	1558	\N	\N	2011-03-25 00:00:00+07	1558.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2219	\N	2	588	\N	\N	2011-03-28 00:00:00+07	0588.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2220	\N	2	1559	\N	\N	2011-03-28 00:00:00+07	1559.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2221	\N	2	1560	\N	\N	2011-03-28 00:00:00+07	1560.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2222	\N	2	1560	\N	\N	2011-03-28 00:00:00+07	1560.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2223	\N	2	567	\N	\N	2011-03-29 00:00:00+07	0567.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2224	\N	2	568	\N	\N	2011-03-29 00:00:00+07	0568.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2225	\N	2	214	\N	\N	2011-03-30 00:00:00+07	0214.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2226	\N	2	631	\N	\N	2011-03-23 00:00:00+07	0631.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2227	\N	2	1554	\N	\N	2011-04-01 00:00:00+07	1554.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2228	\N	2	561	\N	\N	2011-04-01 00:00:00+07	0561.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2229	\N	2	575	\N	\N	2011-04-14 00:00:00+07	0575.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2231	\N	2	1566	\N	\N	2011-04-15 00:00:00+07	1566.62.101	2	2018-10-19 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2232	\N	2	580	\N	\N	2011-04-20 00:00:00+07	0580.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2233	\N	2	1569	\N	\N	2011-04-25 00:00:00+07	1569.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2234	\N	2	1568	\N	\N	2011-04-21 00:00:00+07	1568.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2235	\N	2	635	\N	\N	2011-04-27 00:00:00+07	0635.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2236	\N	2	636	\N	\N	2011-04-27 00:00:00+07	0636.61.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2237	\N	2	634	\N	\N	2011-05-02 00:00:00+07	0634.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2238	\N	2	576	\N	\N	2011-05-02 00:00:00+07	0576.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2239	\N	2	577	\N	\N	2011-05-02 00:00:00+07	0577.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2240	\N	2	578	\N	\N	2011-05-02 00:00:00+07	0578.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2241	\N	2	637	\N	\N	2011-05-02 00:00:00+07	0637.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2242	\N	2	583	\N	\N	2011-05-09 00:00:00+07	0583.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2243	\N	2	581	\N	\N	2011-05-10 00:00:00+07	0581.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2244	\N	2	582	\N	\N	2011-05-10 00:00:00+07	0582.65.101	2	2019-12-26 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2245	\N	2	1266	\N	\N	2011-05-09 00:00:00+07	1266.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2246	\N	2	1567	\N	\N	2011-05-18 00:00:00+07	1567.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2247	\N	2	1573	\N	\N	2011-05-20 00:00:00+07	1573.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2248	\N	2	584	\N	\N	2011-05-20 00:00:00+07	0584.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2249	\N	2	349	\N	\N	2011-06-01 00:00:00+07	0349.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2250	\N	2	587	\N	\N	2011-06-01 00:00:00+07	0587.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2251	\N	2	1577	\N	\N	2011-06-01 00:00:00+07	1577.62.101	2	2017-12-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2252	\N	2	348	\N	\N	2011-06-01 00:00:00+07	0348.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2253	\N	2	1576	\N	\N	2011-06-01 00:00:00+07	1576.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2254	\N	2	1575	\N	\N	2011-06-01 00:00:00+07	1575.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2255	\N	2	585	\N	\N	2011-06-01 00:00:00+07	0585.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2256	\N	2	586	\N	\N	2011-06-01 00:00:00+07	0586.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2258	\N	2	1578	\N	\N	2011-06-15 00:00:00+07	1578.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2260	\N	2	491	\N	\N	2011-07-01 00:00:00+07	0491.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2261	\N	2	217	\N	\N	2011-07-01 00:00:00+07	0217.63.101	2	2019-04-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2262	\N	2	640	\N	\N	2011-07-01 00:00:00+07	0640.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2263	\N	2	433	\N	\N	2011-07-01 00:00:00+07	0433.63.101	2	2017-09-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2264	\N	2	1581	\N	\N	2011-07-01 00:00:00+07	1581.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2265	\N	2	1580	\N	\N	2011-07-01 00:00:00+07	1580.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2266	\N	2	1579	\N	\N	2011-07-01 00:00:00+07	1579.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2267	\N	2	591	\N	\N	2011-07-01 00:00:00+07	0591.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2268	\N	2	590	\N	\N	2011-07-01 00:00:00+07	0590.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2269	\N	2	589	\N	\N	2011-07-01 00:00:00+07	0589.65.101	1	2017-05-17 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2270	\N	2	639	\N	\N	2011-07-26 00:00:00+07	0639.61.101	2	2016-11-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2271	\N	2	1587	\N	\N	2011-07-27 00:00:00+07	1587.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2272	\N	2	595	\N	\N	2011-08-01 00:00:00+07	0595.65.101	2	2021-07-28 00:00:00+07	\N	OA	\N	10112	solusiti	\N	\N	f	f	2021-07-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2273	\N	2	220	\N	\N	2011-08-01 00:00:00+07	0220.63.101	2	2019-07-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2274	\N	2	594	\N	\N	2011-08-01 00:00:00+07	0594.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2275	\N	2	593	\N	\N	2011-08-01 00:00:00+07	0593.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2276	\N	2	1585	\N	\N	2011-08-01 00:00:00+07	1585.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2277	\N	2	641	\N	\N	2011-08-01 00:00:00+07	0641.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2388	\N	2	1675	\N	\N	2013-05-01 00:00:00+07	1675.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2278	\N	2	219	\N	\N	2011-08-01 00:00:00+07	0219.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2279	\N	2	596	\N	\N	2011-08-01 00:00:00+07	0596.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2280	\N	2	597	\N	\N	2011-08-01 00:00:00+07	0597.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2281	\N	2	1590	\N	\N	2011-08-01 00:00:00+07	1590.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2282	\N	2	1589	\N	\N	2011-08-01 00:00:00+07	1589.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2287	\N	2	223	\N	\N	2011-12-01 00:00:00+07	0223.63.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2288	\N	2	224	\N	\N	2011-12-01 00:00:00+07	0224.63.101	1	2017-07-07 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2289	\N	2	644	\N	\N	2011-12-01 00:00:00+07	0644.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2290	\N	2	601	\N	\N	2011-12-01 00:00:00+07	0601.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2291	\N	2	1597	\N	\N	2011-12-01 00:00:00+07	1597.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2292	\N	2	225	\N	\N	2011-12-01 00:00:00+07	0225.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2293	\N	2	226	\N	\N	2011-12-01 00:00:00+07	0226.63.101	1	2017-04-27 00:00:00+07	2019-06-23 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2294	\N	2	602	\N	\N	2011-12-01 00:00:00+07	0602.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2295	\N	2	645	\N	\N	2011-12-01 00:00:00+07	0645.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2296	\N	2	1598	\N	\N	2011-12-01 00:00:00+07	1598.62.101	2	2017-12-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2297	\N	2	616	\N	\N	2011-12-01 00:00:00+07	0616.61.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2298	\N	2	647	\N	\N	2011-12-01 00:00:00+07	0647.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2299	\N	2	648	\N	\N	2011-12-01 00:00:00+07	0648.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2300	\N	2	1599	\N	\N	2011-12-01 00:00:00+07	1599.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2301	\N	2	652	\N	\N	2011-12-01 00:00:00+07	0652.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2302	\N	2	1600	\N	\N	2011-12-01 00:00:00+07	1600.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2303	\N	2	603	\N	\N	2011-12-01 00:00:00+07	0603.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2304	\N	2	604	\N	\N	2011-12-01 00:00:00+07	0604.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2305	\N	2	607	\N	\N	2011-12-01 00:00:00+07	0607.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2306	\N	2	608	\N	\N	2011-12-14 00:00:00+07	0608.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2307	\N	2	1593	\N	\N	2011-10-03 00:00:00+07	1593.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2308	\N	2	610	\N	\N	2011-12-21 00:00:00+07	0610.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2310	\N	2	1609	\N	\N	2011-12-23 00:00:00+07	1609.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2312	\N	2	1613	\N	\N	2011-12-30 00:00:00+07	1613.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2313	\N	2	1619	\N	\N	2012-01-03 00:00:00+07	1619.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2314	\N	2	1622	\N	\N	2011-12-19 00:00:00+07	1622.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2316	\N	2	613	\N	\N	2012-01-19 00:00:00+07	0613.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2318	\N	2	128	\N	\N	2012-02-01 00:00:00+07	0128.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3577	\N	1	2113	\N	\N	2017-03-01 00:00:00+07	2113.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2320	\N	2	1623	\N	\N	2012-02-01 00:00:00+07	1623.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2323	\N	2	614	\N	\N	2012-02-17 00:00:00+07	0614.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2324	\N	2	1630	\N	\N	2012-02-22 00:00:00+07	1630.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2325	\N	2	616	\N	\N	2012-03-09 00:00:00+07	0616.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2326	\N	2	360	\N	\N	2012-03-14 00:00:00+07	0360.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2327	\N	2	1633	\N	\N	2012-03-15 00:00:00+07	1633.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2328	\N	2	232	\N	\N	2012-04-02 00:00:00+07	0232.63.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2329	\N	2	619	\N	\N	2012-04-05 00:00:00+07	0619.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2330	\N	2	1634	\N	\N	2012-04-10 00:00:00+07	1634.62.101	2	2018-03-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2331	\N	2	1635	\N	\N	2012-04-10 00:00:00+07	1635.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2332	\N	2	1642	\N	\N	2012-04-23 00:00:00+07	1642.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2333	\N	2	620	\N	\N	2012-04-23 00:00:00+07	0620.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2336	\N	2	623	\N	\N	2012-05-02 00:00:00+07	0623.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2337	\N	2	361	\N	\N	2012-05-02 00:00:00+07	0361.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2338	\N	2	663	\N	\N	2012-05-08 00:00:00+07	0663.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2339	\N	2	664	\N	\N	2012-05-09 00:00:00+07	0664.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2340	\N	2	1643	\N	\N	2012-05-15 00:00:00+07	1643.62.101	2	2018-07-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2341	\N	2	667	\N	\N	2012-05-25 00:00:00+07	0667.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2342	\N	2	362	\N	\N	2012-05-28 00:00:00+07	0362.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2343	\N	2	669	\N	\N	2012-05-30 00:00:00+07	0669.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2344	\N	2	624	\N	\N	2012-06-04 00:00:00+07	0624.65.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2345	\N	2	1645	\N	\N	2012-06-04 00:00:00+07	1645.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2346	\N	2	1644	\N	\N	2012-06-04 00:00:00+07	1644.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2347	\N	2	626	\N	\N	2012-06-05 00:00:00+07	0626.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2348	\N	2	619	\N	\N	2012-06-13 00:00:00+07	0619.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2349	\N	2	1646	\N	\N	2012-06-19 00:00:00+07	1646.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2350	\N	2	1647	\N	\N	2012-07-02 00:00:00+07	1647.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2352	\N	2	630	\N	\N	2012-07-04 00:00:00+07	0630.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2353	\N	2	629	\N	\N	2012-07-10 00:00:00+07	0629.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2355	\N	2	631	\N	\N	2012-08-10 00:00:00+07	0631.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2356	\N	2	234	\N	\N	2012-09-11 00:00:00+07	0234.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2357	\N	2	633	\N	\N	2012-10-08 00:00:00+07	0633.65.101	2	2017-08-25 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2358	\N	2	637	\N	\N	2012-10-23 00:00:00+07	0637.65.101	2	2018-01-15 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2359	\N	2	1653	\N	\N	2012-11-13 00:00:00+07	1653.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2361	\N	2	237	\N	\N	2012-12-03 00:00:00+07	0237.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2362	\N	2	638	\N	\N	2012-12-17 00:00:00+07	0638.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3465	\N	1	1	\N	\N	2017-03-01 00:00:00+07	0001.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2368	\N	2	1102	\N	\N	2013-02-04 00:00:00+07	1102.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2369	\N	2	1290	\N	\N	2013-02-04 00:00:00+07	1290.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2370	\N	2	1285	\N	\N	2013-02-04 00:00:00+07	1285.62.101	2	2017-09-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3983	\N	1	288	\N	\N	2015-08-18 00:00:00+07	0288.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2372	\N	2	1396	\N	\N	2013-02-04 00:00:00+07	1396.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2373	\N	2	641	\N	\N	2013-02-13 00:00:00+07	0641.65.101	1	\N	2020-06-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2377	\N	2	642	\N	\N	2013-02-28 00:00:00+07	0642.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3578	\N	1	2144	\N	\N	2017-03-01 00:00:00+07	2144.62.101	1	\N	2018-06-16 00:00:00+07	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2380	\N	2	670	\N	\N	2013-04-01 00:00:00+07	0670.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2381	\N	2	1672	\N	\N	2013-04-08 00:00:00+07	1672.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2385	\N	2	643	\N	\N	2013-04-30 00:00:00+07	0643.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2387	\N	2	1674	\N	\N	2013-05-01 00:00:00+07	1674.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2389	\N	2	645	\N	\N	2013-05-10 00:00:00+07	0645.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2390	\N	2	1683	\N	\N	2013-07-19 00:00:00+07	1683.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2391	\N	2	1684	\N	\N	2013-07-22 00:00:00+07	1684.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2393	\N	2	1539	\N	\N	2013-08-26 00:00:00+07	1539.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2397	\N	2	1692	\N	\N	2013-09-17 00:00:00+07	1692.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2400	\N	2	1698	\N	\N	2013-10-16 00:00:00+07	1698.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-13 00:00:00+07	\N	10:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2402	\N	2	1677	\N	\N	2013-10-22 00:00:00+07	1677.62.101	2	2017-09-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2403	\N	2	1701	\N	\N	2013-10-24 00:00:00+07	1701.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2405	\N	2	682	\N	\N	2013-10-31 00:00:00+07	0682.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2407	\N	2	679	\N	\N	2013-11-15 00:00:00+07	0679.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2408	\N	2	1702	\N	\N	2013-11-25 00:00:00+07	1702.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2410	\N	2	236	\N	\N	2013-12-06 00:00:00+07	0236.63.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2411	\N	2	1697	\N	\N	2013-12-09 00:00:00+07	1697.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2412	\N	2	237	\N	\N	2013-12-09 00:00:00+07	0237.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2413	\N	2	1705	\N	\N	2013-12-20 00:00:00+07	1705.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2414	\N	2	1707	\N	\N	2013-12-30 00:00:00+07	1707.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2418	\N	2	1699	\N	\N	2014-01-30 00:00:00+07	1699.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2419	\N	2	638	\N	\N	2014-02-05 00:00:00+07	0638.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2420	\N	2	688	\N	\N	2014-02-17 00:00:00+07	0688.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2422	\N	2	711	\N	\N	2014-02-27 00:00:00+07	0711.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2423	\N	2	1710	\N	\N	2014-02-27 00:00:00+07	1710.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2427	\N	2	376	\N	\N	2014-03-17 00:00:00+07	0376.64.101	1	\N	2019-06-05 00:00:00+07	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2428	\N	2	736	\N	\N	2014-03-27 00:00:00+07	0736.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2429	\N	2	737	\N	\N	2014-03-28 00:00:00+07	0737.65.101	2	2020-02-17 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2430	\N	2	1720	\N	\N	2014-04-08 00:00:00+07	1720.62.101	2	2019-04-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2431	\N	2	724	\N	\N	2014-04-14 00:00:00+07	0724.65.101	2	2017-03-27 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2432	\N	2	1752	\N	\N	2014-04-17 00:00:00+07	1752.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2433	\N	2	694	\N	\N	2014-04-17 00:00:00+07	0694.61.101	2	2021-01-31 00:00:00+07	\N	OA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-07-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2434	\N	2	695	\N	\N	2014-04-25 00:00:00+07	0695.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2435	\N	2	746	\N	\N	2014-04-25 00:00:00+07	0746.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2436	\N	2	747	\N	\N	2014-04-29 00:00:00+07	0747.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2437	\N	2	377	\N	\N	2014-04-30 00:00:00+07	0377.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2438	\N	2	690	\N	\N	2014-05-13 00:00:00+07	0690.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2440	\N	2	700	\N	\N	2014-05-28 00:00:00+07	0700.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2442	\N	2	767	\N	\N	2014-06-04 00:00:00+07	0767.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2443	\N	2	378	\N	\N	2014-06-04 00:00:00+07	0378.64.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2444	\N	2	702	\N	\N	2014-06-09 00:00:00+07	0702.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2445	\N	2	701	\N	\N	2014-06-09 00:00:00+07	0701.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2446	\N	2	703	\N	\N	2014-06-09 00:00:00+07	0703.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2447	\N	2	1756	\N	\N	2014-06-10 00:00:00+07	1756.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2448	\N	2	1758	\N	\N	2014-06-12 00:00:00+07	1758.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2450	\N	2	1764	\N	\N	2014-06-23 00:00:00+07	1764.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2451	\N	2	781	\N	\N	2014-07-07 00:00:00+07	0781.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2454	\N	2	1767	\N	\N	2014-08-15 00:00:00+07	1767.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2456	\N	2	945	\N	\N	2014-09-02 00:00:00+07	0945.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2457	\N	2	802	\N	\N	2014-09-05 00:00:00+07	0802.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6233	\N	1	865	\N	\N	2020-10-07 00:00:00+07	0865.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-07 00:00:00+07	\N	07:00	20:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2459	\N	2	383	\N	\N	2014-09-08 00:00:00+07	0383.64.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2460	\N	2	711	\N	\N	2014-09-11 00:00:00+07	0711.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2461	\N	2	244	\N	\N	2014-09-15 00:00:00+07	0244.63.101	2	2017-10-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2463	\N	2	1783	\N	\N	2014-09-29 00:00:00+07	1783.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2465	\N	2	805	\N	\N	2014-10-07 00:00:00+07	0805.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2466	\N	2	818	\N	\N	2014-10-07 00:00:00+07	0818.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2467	\N	2	859	\N	\N	2014-10-09 00:00:00+07	0859.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2468	\N	2	785	\N	\N	2014-10-10 00:00:00+07	0785.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2469	\N	2	807	\N	\N	2014-10-10 00:00:00+07	0807.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2470	\N	2	1788	\N	\N	2014-10-16 00:00:00+07	1788.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2471	\N	2	710	\N	\N	2014-10-16 00:00:00+07	0710.61.101	2	2019-07-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2473	\N	2	1779	\N	\N	2014-10-31 00:00:00+07	1779.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2474	\N	2	897	\N	\N	2014-11-03 00:00:00+07	0897.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2475	\N	2	1780	\N	\N	2014-11-11 00:00:00+07	1780.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2476	\N	2	1793	\N	\N	2014-11-14 00:00:00+07	1793.62.101	1	2017-05-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2477	\N	2	243	\N	\N	2014-11-17 00:00:00+07	0243.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2478	\N	2	714	\N	\N	2014-11-24 00:00:00+07	0714.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2479	\N	2	382	\N	\N	2014-11-28 00:00:00+07	0382.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2480	\N	2	954	\N	\N	2014-12-02 00:00:00+07	0954.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2481	\N	2	1794	\N	\N	2014-12-03 00:00:00+07	1794.62.101	2	2017-11-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2482	\N	2	1799	\N	\N	2014-12-03 00:00:00+07	1799.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2484	\N	2	1796	\N	\N	2014-12-11 00:00:00+07	1796.62.101	2	2017-04-03 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2485	\N	2	1719	\N	\N	2014-12-11 00:00:00+07	1719.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2486	\N	2	1803	\N	\N	2014-12-15 00:00:00+07	1803.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2488	\N	2	1797	\N	\N	2014-12-19 00:00:00+07	1797.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2490	\N	2	1795	\N	\N	2014-12-22 00:00:00+07	1795.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2491	\N	2	715	\N	\N	2014-12-22 00:00:00+07	0715.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2492	\N	2	718	\N	\N	2014-12-23 00:00:00+07	0718.61.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2493	\N	2	1801	\N	\N	2014-12-23 00:00:00+07	1801.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2494	\N	2	1032	\N	\N	2015-01-12 00:00:00+07	1032.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2495	\N	2	1812	\N	\N	2015-01-12 00:00:00+07	1812.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2498	\N	2	252	\N	\N	2015-01-27 00:00:00+07	0252.63.101	2	2019-01-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2499	\N	2	99	\N	\N	2015-01-28 00:00:00+07	0099.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2500	\N	2	1819	\N	\N	2015-02-02 00:00:00+07	1819.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2501	\N	2	420	\N	\N	2015-02-03 00:00:00+07	0420.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2502	\N	2	994	\N	\N	2015-02-04 00:00:00+07	0994.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2504	\N	2	1810	\N	\N	2015-02-09 00:00:00+07	1810.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2505	\N	2	723	\N	\N	2015-02-11 00:00:00+07	0723.61.101	1	\N	2019-01-02 00:00:00+07	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2507	\N	2	717	\N	\N	2015-02-18 00:00:00+07	0717.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2508	\N	2	404	\N	\N	2015-02-18 00:00:00+07	0404.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2509	\N	2	407	\N	\N	2015-02-18 00:00:00+07	0407.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6232	\N	1	2436	\N	\N	2020-10-07 00:00:00+07	2436.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-07 00:00:00+07	\N	07:00	\N	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2513	\N	2	1076	\N	\N	2015-02-20 00:00:00+07	1076.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2514	\N	2	1821	\N	\N	2015-02-23 00:00:00+07	1821.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2515	\N	2	960	\N	\N	2015-02-24 00:00:00+07	0960.65.101	2	2018-10-04 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2517	\N	2	1823	\N	\N	2015-03-09 00:00:00+07	1823.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2518	\N	2	1138	\N	\N	2015-03-10 00:00:00+07	1138.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2519	\N	2	1832	\N	\N	2015-03-10 00:00:00+07	1832.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2520	\N	2	1834	\N	\N	2015-03-12 00:00:00+07	1834.62.101	2	2018-04-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2521	\N	2	833	\N	\N	2015-03-12 00:00:00+07	0833.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2523	\N	2	1838	\N	\N	2015-03-16 00:00:00+07	1838.62.101	1	2017-06-20 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2524	\N	2	1839	\N	\N	2015-03-16 00:00:00+07	1839.62.101	1	2017-06-20 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2525	\N	2	725	\N	\N	2015-03-18 00:00:00+07	0725.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2526	\N	2	1826	\N	\N	2015-03-20 00:00:00+07	1826.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2528	\N	2	1159	\N	\N	2015-03-26 00:00:00+07	1159.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2529	\N	2	1843	\N	\N	2015-03-30 00:00:00+07	1843.62.101	2	2018-09-02 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2530	\N	2	1844	\N	\N	2015-03-30 00:00:00+07	1844.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2531	\N	2	1847	\N	\N	2015-04-06 00:00:00+07	1847.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2532	\N	2	1848	\N	\N	2015-04-07 00:00:00+07	1848.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2533	\N	2	413	\N	\N	2015-04-07 00:00:00+07	0413.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2534	\N	2	1161	\N	\N	2015-04-08 00:00:00+07	1161.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2535	\N	2	1816	\N	\N	2015-04-08 00:00:00+07	1816.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2536	\N	2	247	\N	\N	2015-04-09 00:00:00+07	0247.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2537	\N	2	926	\N	\N	2015-04-09 00:00:00+07	0926.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2538	\N	2	1825	\N	\N	2015-04-13 00:00:00+07	1825.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2539	\N	2	1167	\N	\N	2015-04-13 00:00:00+07	1167.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2540	\N	2	1811	\N	\N	2015-04-14 00:00:00+07	1811.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2541	\N	2	1849	\N	\N	2015-04-15 00:00:00+07	1849.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2542	\N	2	1850	\N	\N	2015-04-15 00:00:00+07	1850.62.101	2	2018-04-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2543	\N	2	1829	\N	\N	2015-04-15 00:00:00+07	1829.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2544	\N	2	415	\N	\N	2015-04-21 00:00:00+07	0415.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2545	\N	2	1045	\N	\N	2015-04-23 00:00:00+07	1045.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2547	\N	2	726	\N	\N	2015-04-28 00:00:00+07	0726.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2549	\N	2	1215	\N	\N	2015-04-28 00:00:00+07	1215.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2550	\N	2	254	\N	\N	2015-04-28 00:00:00+07	0254.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2551	\N	2	915	\N	\N	2015-04-29 00:00:00+07	0915.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2552	\N	2	1	\N	\N	2015-04-30 00:00:00+07	0001.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2553	\N	2	1217	\N	\N	2015-04-30 00:00:00+07	1217.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2554	\N	2	1218	\N	\N	2015-04-30 00:00:00+07	1218.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2555	\N	2	1225	\N	\N	2015-05-04 00:00:00+07	1225.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2556	\N	2	155	\N	\N	2015-05-04 00:00:00+07	0155.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2557	\N	2	256	\N	\N	2015-05-06 00:00:00+07	0256.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2558	\N	2	1231	\N	\N	2015-05-06 00:00:00+07	1231.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2559	\N	2	417	\N	\N	2015-05-06 00:00:00+07	0417.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2560	\N	2	257	\N	\N	2015-05-06 00:00:00+07	0257.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2561	\N	2	1234	\N	\N	2015-05-06 00:00:00+07	1234.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2563	\N	2	259	\N	\N	2015-05-06 00:00:00+07	0259.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6231	\N	1	2435	\N	\N	2020-10-05 00:00:00+07	2435.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-05 00:00:00+07	120	12:00	21:00	33	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2568	\N	2	1233	\N	\N	2015-05-07 00:00:00+07	1233.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2569	\N	2	1216	\N	\N	2015-05-07 00:00:00+07	1216.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2570	\N	2	1856	\N	\N	2015-05-07 00:00:00+07	1856.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2571	\N	2	1857	\N	\N	2015-05-08 00:00:00+07	1857.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3990	\N	1	2190	\N	\N	2017-09-12 00:00:00+07	2190.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2573	\N	2	261	\N	\N	2015-05-08 00:00:00+07	0261.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2574	\N	2	8	\N	\N	2015-05-08 00:00:00+07	0008.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2575	\N	2	263	\N	\N	2015-05-08 00:00:00+07	0263.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2576	\N	2	716	\N	\N	2015-05-08 00:00:00+07	0716.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2577	\N	2	264	\N	\N	2015-05-08 00:00:00+07	0264.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2578	\N	2	7	\N	\N	2015-05-08 00:00:00+07	0007.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2579	\N	2	730	\N	\N	2015-05-08 00:00:00+07	0730.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2580	\N	2	6	\N	\N	2015-05-08 00:00:00+07	0006.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2581	\N	2	5	\N	\N	2015-05-08 00:00:00+07	0005.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2582	\N	2	421	\N	\N	2015-05-08 00:00:00+07	0421.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2583	\N	2	732	\N	\N	2015-05-08 00:00:00+07	0732.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2584	\N	2	1238	\N	\N	2015-05-08 00:00:00+07	1238.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2588	\N	2	1858	\N	\N	2015-05-11 00:00:00+07	1858.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2589	\N	2	731	\N	\N	2015-05-11 00:00:00+07	0731.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2590	\N	2	267	\N	\N	2015-05-11 00:00:00+07	0267.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2593	\N	2	1241	\N	\N	2015-05-11 00:00:00+07	1241.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2594	\N	2	268	\N	\N	2015-05-12 00:00:00+07	0268.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2595	\N	2	269	\N	\N	2015-05-12 00:00:00+07	0269.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2596	\N	2	270	\N	\N	2015-05-12 00:00:00+07	0270.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2597	\N	2	1212	\N	\N	2015-05-12 00:00:00+07	1212.65.101	2	2019-01-05 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2598	\N	2	271	\N	\N	2015-05-12 00:00:00+07	0271.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2599	\N	2	423	\N	\N	2015-05-13 00:00:00+07	0423.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2600	\N	2	422	\N	\N	2015-05-13 00:00:00+07	0422.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2601	\N	2	1859	\N	\N	2015-05-13 00:00:00+07	1859.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2602	\N	2	1860	\N	\N	2015-05-15 00:00:00+07	1860.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2603	\N	2	425	\N	\N	2015-05-15 00:00:00+07	0425.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2604	\N	2	735	\N	\N	2015-05-15 00:00:00+07	0735.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2605	\N	2	424	\N	\N	2015-05-15 00:00:00+07	0424.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2606	\N	2	734	\N	\N	2015-05-15 00:00:00+07	0734.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2607	\N	2	11	\N	\N	2015-05-15 00:00:00+07	0011.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2608	\N	2	1861	\N	\N	2015-05-15 00:00:00+07	1861.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2609	\N	2	272	\N	\N	2015-05-15 00:00:00+07	0272.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2611	\N	2	258	\N	\N	2015-05-18 00:00:00+07	0258.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2612	\N	2	427	\N	\N	2015-05-18 00:00:00+07	0427.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2613	\N	2	1248	\N	\N	2015-05-18 00:00:00+07	1248.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2614	\N	2	1251	\N	\N	2015-05-18 00:00:00+07	1251.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2615	\N	2	428	\N	\N	2015-05-19 00:00:00+07	0428.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2616	\N	2	1891	\N	\N	2015-05-19 00:00:00+07	1891.62.101	2	2021-01-01 00:00:00+07	2020-04-30 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2617	\N	2	429	\N	\N	2015-05-19 00:00:00+07	0429.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2618	\N	2	280	\N	\N	2015-05-19 00:00:00+07	0280.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2619	\N	2	1252	\N	\N	2015-05-19 00:00:00+07	1252.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2620	\N	2	265	\N	\N	2015-05-20 00:00:00+07	0265.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2621	\N	2	1255	\N	\N	2015-05-21 00:00:00+07	1255.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2622	\N	2	737	\N	\N	2015-05-21 00:00:00+07	0737.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2623	\N	2	277	\N	\N	2015-05-21 00:00:00+07	0277.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2626	\N	2	1879	\N	\N	2015-05-27 00:00:00+07	1879.62.101	1	\N	2020-04-30 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2627	\N	2	279	\N	\N	2015-05-28 00:00:00+07	0279.63.101	2	2018-07-19 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2629	\N	2	1777	\N	\N	2015-06-03 00:00:00+07	1777.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2630	\N	2	1261	\N	\N	2015-06-04 00:00:00+07	1261.64.101	2	2018-02-06 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2631	\N	2	281	\N	\N	2015-06-08 00:00:00+07	0281.63.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2632	\N	2	741	\N	\N	2015-06-08 00:00:00+07	0741.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2633	\N	2	729	\N	\N	2015-06-08 00:00:00+07	0729.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2634	\N	2	1885	\N	\N	2015-06-09 00:00:00+07	1885.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2639	\N	2	1889	\N	\N	2015-06-12 00:00:00+07	1889.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2641	\N	2	1272	\N	\N	2015-06-15 00:00:00+07	1272.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2642	\N	2	1880	\N	\N	2015-06-17 00:00:00+07	1880.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2643	\N	2	275	\N	\N	2015-06-19 00:00:00+07	0275.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2644	\N	2	1876	\N	\N	2015-06-19 00:00:00+07	1876.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2646	\N	2	273	\N	\N	2015-06-19 00:00:00+07	0273.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2647	\N	2	1244	\N	\N	2015-06-19 00:00:00+07	1244.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2648	\N	2	12	\N	\N	2015-06-19 00:00:00+07	0012.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2649	\N	2	276	\N	\N	2015-06-19 00:00:00+07	0276.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2988	\N	1	2027	\N	\N	2016-01-01 00:00:00+07	2027.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2651	\N	2	1862	\N	\N	2015-06-19 00:00:00+07	1862.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2652	\N	2	736	\N	\N	2015-06-19 00:00:00+07	0736.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2653	\N	2	19	\N	\N	2015-06-19 00:00:00+07	0019.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2654	\N	2	13	\N	\N	2015-06-19 00:00:00+07	0013.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2657	\N	2	1865	\N	\N	2015-06-19 00:00:00+07	1865.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2660	\N	2	1863	\N	\N	2015-06-19 00:00:00+07	1863.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2661	\N	2	1247	\N	\N	2015-06-19 00:00:00+07	1247.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2664	\N	2	1892	\N	\N	2015-06-19 00:00:00+07	1892.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2665	\N	2	1896	\N	\N	2015-06-19 00:00:00+07	1896.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2666	\N	2	1893	\N	\N	2015-06-19 00:00:00+07	1893.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2667	\N	2	1895	\N	\N	2015-06-19 00:00:00+07	1895.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2668	\N	2	1864	\N	\N	2015-06-19 00:00:00+07	1864.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2669	\N	2	1894	\N	\N	2015-06-19 00:00:00+07	1894.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2670	\N	2	742	\N	\N	2015-06-22 00:00:00+07	0742.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2671	\N	2	726	\N	\N	2015-06-23 00:00:00+07	0726.65.101	2	2019-01-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2672	\N	2	1751	\N	\N	2015-06-24 00:00:00+07	1751.62.101	1	2016-11-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2673	\N	2	1750	\N	\N	2015-06-24 00:00:00+07	1750.62.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2674	\N	2	1262	\N	\N	2015-06-26 00:00:00+07	1262.65.101	2	2017-10-31 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2675	\N	2	1884	\N	\N	2015-06-26 00:00:00+07	1884.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2676	\N	2	740	\N	\N	2015-06-26 00:00:00+07	0740.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2677	\N	2	282	\N	\N	2015-06-26 00:00:00+07	0282.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2678	\N	2	283	\N	\N	2015-06-26 00:00:00+07	0283.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2681	\N	2	4	\N	\N	2015-07-08 00:00:00+07	0004.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2682	\N	2	1899	\N	\N	2015-07-10 00:00:00+07	1899.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6230	\N	1	2434	\N	\N	2020-10-05 00:00:00+07	2434.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-05 00:00:00+07	\N	09:00	20:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2684	\N	2	746	\N	\N	2015-07-14 00:00:00+07	0746.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2685	\N	2	745	\N	\N	2015-07-14 00:00:00+07	0745.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2687	\N	2	1882	\N	\N	2015-07-27 00:00:00+07	1882.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2689	\N	2	1897	\N	\N	2015-07-30 00:00:00+07	1897.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2691	\N	2	1291	\N	\N	2015-08-10 00:00:00+07	1291.65.101	2	2017-11-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2693	\N	2	751	\N	\N	2015-08-10 00:00:00+07	0751.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2694	\N	2	1286	\N	\N	2015-08-14 00:00:00+07	1286.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2695	\N	2	1285	\N	\N	2015-08-14 00:00:00+07	1285.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2698	\N	2	1846	\N	\N	2015-08-31 00:00:00+07	1846.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2699	\N	2	1910	\N	\N	2015-08-31 00:00:00+07	1910.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2701	\N	2	1906	\N	\N	2015-09-03 00:00:00+07	1906.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2702	\N	2	1907	\N	\N	2015-09-03 00:00:00+07	1907.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2703	\N	2	440	\N	\N	2015-09-03 00:00:00+07	0440.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2704	\N	2	1908	\N	\N	2015-09-03 00:00:00+07	1908.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2705	\N	2	441	\N	\N	2015-09-03 00:00:00+07	0441.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2707	\N	2	1288	\N	\N	2015-09-07 00:00:00+07	1288.65.101	2	2020-12-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2708	\N	2	21	\N	\N	2015-09-07 00:00:00+07	0021.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2709	\N	2	1688	\N	\N	2015-09-09 00:00:00+07	1688.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2710	\N	2	753	\N	\N	2015-09-09 00:00:00+07	0753.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2711	\N	2	738	\N	\N	2015-09-09 00:00:00+07	0738.61.101	2	2017-04-03 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2712	\N	2	739	\N	\N	2015-09-09 00:00:00+07	0739.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2713	\N	2	749	\N	\N	2015-09-09 00:00:00+07	0749.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2714	\N	2	1886	\N	\N	2015-09-10 00:00:00+07	1886.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2715	\N	2	1904	\N	\N	2015-09-11 00:00:00+07	1904.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2716	\N	2	1326	\N	\N	2015-09-14 00:00:00+07	1326.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3746	\N	1	1627	\N	\N	2017-05-29 00:00:00+07	1627.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6229	\N	1	583	\N	\N	2020-09-30 00:00:00+07	0583.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-30 00:00:00+07	\N	07:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2719	\N	2	1345	\N	\N	2015-10-06 00:00:00+07	1345.65.101	2	2017-06-13 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2725	\N	2	1327	\N	\N	2015-10-19 00:00:00+07	1327.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2726	\N	2	22	\N	\N	2015-10-26 00:00:00+07	0022.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2727	\N	2	1927	\N	\N	2015-11-02 00:00:00+07	1927.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2729	\N	2	289	\N	\N	2015-11-03 00:00:00+07	0289.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2730	\N	2	1928	\N	\N	2015-11-04 00:00:00+07	1928.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2732	\N	2	1367	\N	\N	2015-11-05 00:00:00+07	1367.65.101	2	2017-12-31 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2733	\N	2	1366	\N	\N	2015-11-06 00:00:00+07	1366.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2735	\N	2	290	\N	\N	2015-11-10 00:00:00+07	0290.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2736	\N	2	1332	\N	\N	2015-11-10 00:00:00+07	1332.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2737	\N	2	1338	\N	\N	2015-11-10 00:00:00+07	1338.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2740	\N	2	189	\N	\N	2015-11-11 00:00:00+07	0189.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2741	\N	2	1926	\N	\N	2015-11-12 00:00:00+07	1926.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2742	\N	2	759	\N	\N	2015-11-12 00:00:00+07	0759.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2744	\N	2	1933	\N	\N	2015-11-23 00:00:00+07	1933.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2745	\N	2	1925	\N	\N	2015-11-24 00:00:00+07	1925.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2746	\N	2	1375	\N	\N	2015-11-24 00:00:00+07	1375.65.101	1	\N	2020-04-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2747	\N	2	1934	\N	\N	2015-12-01 00:00:00+07	1934.62.101	2	2020-01-07 00:00:00+07	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2749	\N	2	1924	\N	\N	2015-12-08 00:00:00+07	1924.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2750	\N	2	452	\N	\N	2015-12-10 00:00:00+07	0452.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2752	\N	2	1920	\N	\N	2015-12-10 00:00:00+07	1920.62.101	1	\N	2020-04-30 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2753	\N	2	1868	\N	\N	2015-12-29 00:00:00+07	1868.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2754	\N	2	295	\N	\N	2016-01-07 00:00:00+07	0295.63.101	1	\N	2020-07-01 00:00:00+07	OA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2758	\N	2	457	\N	\N	2016-01-14 00:00:00+07	0457.64.101	1	2017-08-09 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2759	\N	2	1401	\N	\N	2016-01-15 00:00:00+07	1401.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2760	\N	2	927	\N	\N	2016-01-18 00:00:00+07	0927.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2761	\N	2	1919	\N	\N	2016-01-18 00:00:00+07	1919.62.101	1	\N	2020-04-30 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3472	\N	1	472	\N	\N	2017-03-01 00:00:00+07	0472.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2763	\N	2	754	\N	\N	2016-01-20 00:00:00+07	0754.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2764	\N	2	1389	\N	\N	2016-01-20 00:00:00+07	1389.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2765	\N	2	761	\N	\N	2016-01-20 00:00:00+07	0761.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2769	\N	2	1376	\N	\N	2016-02-02 00:00:00+07	1376.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2771	\N	2	1946	\N	\N	2016-02-09 00:00:00+07	1946.62.101	2	2017-07-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2772	\N	2	1299	\N	\N	2016-02-09 00:00:00+07	1299.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2773	\N	2	1947	\N	\N	2016-02-09 00:00:00+07	1947.62.101	2	2020-02-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2774	\N	2	1932	\N	\N	2016-02-09 00:00:00+07	1932.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3473	\N	1	473	\N	\N	2017-03-01 00:00:00+07	0473.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2782	\N	2	1399	\N	\N	2016-02-11 00:00:00+07	1399.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2783	\N	2	1368	\N	\N	2016-02-11 00:00:00+07	1368.65.101	1	2017-06-20 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2785	\N	2	1659	\N	\N	2016-02-17 00:00:00+07	1659.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2786	\N	2	408	\N	\N	2016-02-18 00:00:00+07	0408.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2787	\N	2	743	\N	\N	2016-02-19 00:00:00+07	0743.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2788	\N	2	1774	\N	\N	2016-02-29 00:00:00+07	1774.62.101	2	2018-02-26 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2791	\N	2	1209	\N	\N	2016-03-03 00:00:00+07	1209.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2792	\N	2	1261	\N	\N	2016-03-03 00:00:00+07	1261.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2794	\N	2	1429	\N	\N	2016-03-07 00:00:00+07	1429.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2795	\N	2	1430	\N	\N	2016-03-07 00:00:00+07	1430.65.101	2	2018-05-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2796	\N	2	469	\N	\N	2016-03-07 00:00:00+07	0469.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2797	\N	2	1431	\N	\N	2016-03-07 00:00:00+07	1431.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2798	\N	2	1961	\N	\N	2016-03-07 00:00:00+07	1961.62.101	2	2018-09-29 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2801	\N	2	1948	\N	\N	2016-03-10 00:00:00+07	1948.62.101	2	2022-01-01 00:00:00+07	2020-04-30 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2805	\N	2	1444	\N	\N	2016-04-04 00:00:00+07	1444.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2807	\N	2	1403	\N	\N	2016-04-05 00:00:00+07	1403.65.101	1	2017-06-02 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2810	\N	2	1975	\N	\N	2016-04-22 00:00:00+07	1975.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2811	\N	2	1976	\N	\N	2016-04-22 00:00:00+07	1976.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6227	\N	1	2433	\N	\N	2020-09-28 00:00:00+07	2433.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2813	\N	2	1973	\N	\N	2016-04-28 00:00:00+07	1973.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2814	\N	2	1974	\N	\N	2016-04-28 00:00:00+07	1974.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2815	\N	2	771	\N	\N	2016-04-28 00:00:00+07	0771.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2816	\N	2	1472	\N	\N	2016-05-02 00:00:00+07	1472.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2817	\N	2	773	\N	\N	2016-05-02 00:00:00+07	0773.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2818	\N	2	1451	\N	\N	2016-05-04 00:00:00+07	1451.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2819	\N	2	1982	\N	\N	2016-05-04 00:00:00+07	1982.62.101	2	2018-11-25 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2822	\N	2	1455	\N	\N	2016-05-09 00:00:00+07	1455.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2823	\N	2	1477	\N	\N	2016-05-09 00:00:00+07	1477.65.101	2	2019-06-28 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2824	\N	2	1479	\N	\N	2016-05-10 00:00:00+07	1479.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2825	\N	2	1978	\N	\N	2016-05-10 00:00:00+07	1978.62.101	2	2020-02-25 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2826	\N	2	286	\N	\N	2016-05-10 00:00:00+07	0286.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2827	\N	2	772	\N	\N	2016-05-16 00:00:00+07	0772.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2828	\N	2	1481	\N	\N	2016-05-17 00:00:00+07	1481.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3555	\N	1	31	\N	\N	2017-03-01 00:00:00+07	0031.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2833	\N	2	1983	\N	\N	2016-06-09 00:00:00+07	1983.62.101	2	2018-04-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2834	\N	2	1994	\N	\N	2016-06-21 00:00:00+07	1994.62.101	2	2021-04-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-07-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2836	\N	2	1480	\N	\N	2016-07-11 00:00:00+07	1480.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2952	\N	1	2001	\N	\N	2017-03-20 00:00:00+07	2001.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2839	\N	2	1992	\N	\N	2016-07-12 00:00:00+07	1992.62.101	2	2019-04-30 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2840	\N	2	1998	\N	\N	2016-07-25 00:00:00+07	1998.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2841	\N	2	1988	\N	\N	2016-08-04 00:00:00+07	1988.62.101	1	\N	2019-05-01 00:00:00+07	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2844	\N	2	1491	\N	\N	2016-08-08 00:00:00+07	1491.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2845	\N	2	781	\N	\N	2016-08-10 00:00:00+07	0781.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2846	\N	2	1492	\N	\N	2016-08-10 00:00:00+07	1492.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2847	\N	2	1822	\N	\N	2016-08-22 00:00:00+07	1822.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2849	\N	2	1502	\N	\N	2016-09-05 00:00:00+07	1502.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2850	\N	2	786	\N	\N	2016-09-06 00:00:00+07	0786.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2851	\N	2	1499	\N	\N	2016-09-07 00:00:00+07	1499.65.101	2	2018-03-10 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2852	\N	2	1500	\N	\N	2016-09-07 00:00:00+07	1500.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2854	\N	2	2015	\N	\N	2016-09-08 00:00:00+07	2015.62.101	2	2016-10-31 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2855	\N	2	2008	\N	\N	2016-09-08 00:00:00+07	2008.62.101	2	2018-12-31 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2856	\N	2	1480	\N	\N	2016-09-09 00:00:00+07	1480.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2857	\N	2	1981	\N	\N	2016-09-09 00:00:00+07	1981.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2860	\N	2	1503	\N	\N	2016-09-09 00:00:00+07	1503.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-05-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2861	\N	2	1501	\N	\N	2016-09-09 00:00:00+07	1501.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2862	\N	2	2000	\N	\N	2016-09-09 00:00:00+07	2000.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2864	\N	2	305	\N	\N	2016-09-09 00:00:00+07	0305.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2865	\N	2	2009	\N	\N	2016-09-09 00:00:00+07	2009.62.101	2	2019-02-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2866	\N	2	2005	\N	\N	2016-09-13 00:00:00+07	2005.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2868	\N	2	796	\N	\N	2016-09-16 00:00:00+07	0796.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3475	\N	1	2094	\N	\N	2017-04-21 00:00:00+07	2094.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2871	\N	2	1522	\N	\N	2016-10-07 00:00:00+07	1522.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2872	\N	2	1520	\N	\N	2016-10-10 00:00:00+07	1520.65.101	2	2016-11-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2873	\N	2	2017	\N	\N	2016-10-17 00:00:00+07	2017.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2874	\N	2	1527	\N	\N	2016-10-18 00:00:00+07	1527.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2895	\N	1	434	\N	\N	2017-03-09 00:00:00+07	0434.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2931	\N	1	2022	\N	\N	2016-12-01 00:00:00+07	2022.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2901	\N	1	797	\N	\N	2017-03-10 00:00:00+07	0797.61.101	2	2019-03-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2932	\N	1	1545	\N	\N	2016-12-01 00:00:00+07	1545.65.101	2	2019-03-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2913	\N	1	799	\N	\N	2016-12-01 00:00:00+07	0799.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2940	\N	1	2036	\N	\N	2016-12-01 00:00:00+07	2036.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2914	\N	1	2034	\N	\N	2017-03-13 00:00:00+07	2034.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2942	\N	1	1530	\N	\N	2017-02-01 00:00:00+07	1530.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6225	\N	1	2630	\N	\N	2020-09-23 00:00:00+07	2630.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-23 00:00:00+07	\N	09:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2957	\N	1	2014	\N	\N	2016-11-07 00:00:00+07	2014.62.101	2	2018-05-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6195	\N	1	2624	\N	\N	2020-09-10 00:00:00+07	2624.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-09 00:00:00+07	\N	08:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6193	\N	1	2430	\N	\N	2020-09-09 00:00:00+07	2430.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	2020-09-09 00:00:00+07	\N	10:00	20:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2925	\N	1	2043	\N	\N	2016-12-01 00:00:00+07	2043.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2945	\N	1	1538	\N	\N	2016-12-01 00:00:00+07	1538.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2993	\N	2	2084	\N	\N	2016-11-01 00:00:00+07	2084.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2996	\N	1	328	\N	\N	2016-11-01 00:00:00+07	0328.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6191	\N	1	2623	\N	\N	2020-09-09 00:00:00+07	2623.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-09 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6190	\N	1	2622	\N	\N	2020-09-09 00:00:00+07	2622.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-09 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2973	\N	2	1872	\N	\N	2017-01-02 00:00:00+07	1872.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6187	\N	1	2621	\N	\N	2020-09-09 00:00:00+07	2621.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	2020-09-09 00:00:00+07	\N	08:00	21:00	40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6186	\N	1	2428	\N	\N	2020-09-09 00:00:00+07	2428.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	2020-09-09 00:00:00+07	\N	08:00	18:00	20-30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2985	\N	1	1561	\N	\N	2016-12-19 00:00:00+07	1561.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6178	\N	1	2422	\N	\N	2020-08-31 00:00:00+07	2422.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2987	\N	2	2071	\N	\N	2016-12-05 00:00:00+07	2071.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6177	\N	1	1030	\N	\N	2020-08-31 00:00:00+07	1030.61.101	2	2021-06-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-07-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6168	\N	1	2419	\N	\N	2020-08-14 00:00:00+07	2419.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6167	\N	1	862	\N	\N	2020-08-14 00:00:00+07	0862.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6164	\N	1	859	\N	\N	2020-08-13 00:00:00+07	0859.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6163	\N	1	2418	\N	\N	2020-08-12 00:00:00+07	2418.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6162	\N	1	2417	\N	\N	2020-08-12 00:00:00+07	2417.65.101	2	2021-12-15 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6161	\N	1	2619	\N	\N	2020-08-12 00:00:00+07	2619.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6160	\N	1	2618	\N	\N	2020-08-10 00:00:00+07	2618.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6159	\N	1	2617	\N	\N	2020-08-07 00:00:00+07	2617.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6158	\N	1	1028	\N	\N	2020-08-05 00:00:00+07	1028.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6156	\N	1	2416	\N	\N	2020-08-03 00:00:00+07	2416.65.101	2	2021-10-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6154	\N	1	577	\N	\N	2020-07-29 00:00:00+07	0577.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6151	\N	1	2413	\N	\N	2020-07-28 00:00:00+07	2413.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6149	\N	1	2615	\N	\N	2020-07-22 00:00:00+07	2615.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6142	\N	1	2409	\N	\N	2020-07-15 00:00:00+07	2409.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6141	\N	1	2408	\N	\N	2020-07-15 00:00:00+07	2408.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6135	\N	1	2614	\N	\N	2020-07-10 00:00:00+07	2614.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6134	\N	1	573	\N	\N	2020-07-10 00:00:00+07	0573.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6114	\N	1	2612	\N	\N	2020-07-02 00:00:00+07	2612.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6110	\N	1	1023	\N	\N	2020-07-01 00:00:00+07	1023.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6107	\N	1	2611	\N	\N	2020-06-25 00:00:00+07	2611.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6106	\N	1	2610	\N	\N	2020-06-25 00:00:00+07	2610.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6102	\N	1	2398	\N	\N	2020-06-17 00:00:00+07	2398.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6100	\N	1	851	\N	\N	2020-06-17 00:00:00+07	0851.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6095	\N	1	2397	\N	\N	2020-06-10 00:00:00+07	2397.65.101	2	2021-01-09 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6090	\N	1	2395	\N	\N	2020-06-08 00:00:00+07	2395.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6064	\N	1	561	\N	\N	2020-03-24 00:00:00+07	0561.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6059	\N	1	2603	\N	\N	2020-03-12 00:00:00+07	2603.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6058	\N	1	559	\N	\N	2020-03-11 00:00:00+07	0559.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6057	\N	1	2602	\N	\N	2020-03-11 00:00:00+07	2602.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6055	\N	1	558	\N	\N	2020-03-11 00:00:00+07	0558.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6040	\N	1	839	\N	\N	2020-02-25 00:00:00+07	0839.64.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6034	\N	1	2597	\N	\N	2020-02-18 00:00:00+07	2597.62.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6033	\N	1	2379	\N	\N	2020-02-17 00:00:00+07	2379.65.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6031	\N	1	1011	\N	\N	2020-02-13 00:00:00+07	1011.61.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6030	\N	1	2595	\N	\N	2020-02-12 00:00:00+07	2595.62.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6013	\N	1	1009	\N	\N	2020-01-22 00:00:00+07	1009.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6009	\N	1	2587	\N	\N	2020-01-22 00:00:00+07	2587.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-31 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6000	\N	1	2585	\N	\N	2020-01-16 00:00:00+07	2585.62.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5999	\N	1	2369	\N	\N	2020-01-13 00:00:00+07	2369.65.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5996	\N	1	2584	\N	\N	2020-01-13 00:00:00+07	2584.62.101	1	\N	2020-05-08 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5995	\N	1	832	\N	\N	2020-01-13 00:00:00+07	0832.64.101	1	\N	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5978	\N	1	2579	\N	\N	2020-01-03 00:00:00+07	2579.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5974	\N	1	2360	\N	\N	2019-12-30 00:00:00+07	2360.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5969	\N	1	2577	\N	\N	2019-12-26 00:00:00+07	2577.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5959	\N	1	2352	\N	\N	2019-12-11 00:00:00+07	2352.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5936	\N	1	2572	\N	\N	2019-11-28 00:00:00+07	2572.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5925	\N	1	2565	\N	\N	2019-11-21 00:00:00+07	2565.62.101	2	2020-06-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5919	\N	1	2562	\N	\N	2019-11-18 00:00:00+07	2562.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5913	\N	1	2561	\N	\N	2019-11-14 00:00:00+07	2561.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5912	\N	1	2560	\N	\N	2019-11-14 00:00:00+07	2560.62.101	2	2021-01-27 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5908	\N	1	824	\N	\N	2019-11-13 00:00:00+07	0824.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5893	\N	1	2339	\N	\N	2019-11-07 00:00:00+07	2339.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5892	\N	1	992	\N	\N	2019-11-07 00:00:00+07	0992.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5890	\N	1	2556	\N	\N	2019-11-07 00:00:00+07	2556.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5883	\N	1	2554	\N	\N	2019-11-06 00:00:00+07	2554.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5878	\N	1	2552	\N	\N	2019-10-31 00:00:00+07	2552.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5875	\N	1	2551	\N	\N	2019-10-31 00:00:00+07	2551.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5874	\N	1	2550	\N	\N	2019-10-31 00:00:00+07	2550.62.101	2	2022-02-14 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5873	\N	1	2333	\N	\N	2019-10-31 00:00:00+07	2333.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5871	\N	1	2549	\N	\N	2019-10-30 00:00:00+07	2549.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5870	\N	1	991	\N	\N	2019-10-30 00:00:00+07	0991.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5869	\N	1	2548	\N	\N	2019-10-30 00:00:00+07	2548.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5867	\N	1	2546	\N	\N	2019-10-30 00:00:00+07	2546.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5866	\N	1	2545	\N	\N	2019-10-30 00:00:00+07	2545.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5865	\N	1	2544	\N	\N	2019-10-30 00:00:00+07	2544.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5863	\N	1	2543	\N	\N	2019-10-29 00:00:00+07	2543.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5856	\N	1	2123	\N	\N	2019-10-28 00:00:00+07	2123.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5852	\N	1	2540	\N	\N	2019-10-23 00:00:00+07	2540.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5851	\N	1	2539	\N	\N	2019-10-23 00:00:00+07	2539.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5849	\N	1	2537	\N	\N	2019-10-23 00:00:00+07	2537.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5847	\N	1	2536	\N	\N	2019-10-22 00:00:00+07	2536.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5846	\N	1	2535	\N	\N	2019-10-22 00:00:00+07	2535.62.101	2	2020-04-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5841	\N	1	2534	\N	\N	2019-10-17 00:00:00+07	2534.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5840	\N	1	2533	\N	\N	2019-10-17 00:00:00+07	2533.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5837	\N	1	2529	\N	\N	2019-10-17 00:00:00+07	2529.62.101	2	2020-04-08 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5831	\N	1	1997	\N	\N	2019-10-16 00:00:00+07	1997.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5834	\N	1	2532	\N	\N	2019-10-16 00:00:00+07	2532.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5833	\N	1	2531	\N	\N	2019-10-16 00:00:00+07	2531.62.101	2	2020-09-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5822	\N	1	2528	\N	\N	2019-10-16 00:00:00+07	2528.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5815	\N	1	2524	\N	\N	2019-10-14 00:00:00+07	2524.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5814	\N	1	2523	\N	\N	2019-10-14 00:00:00+07	2523.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5813	\N	1	2323	\N	\N	2019-10-14 00:00:00+07	2323.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5812	\N	1	2522	\N	\N	2019-10-14 00:00:00+07	2522.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5811	\N	1	2521	\N	\N	2019-10-14 00:00:00+07	2521.62.101	2	2020-03-30 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5810	\N	1	2520	\N	\N	2019-10-14 00:00:00+07	2520.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5809	\N	1	2519	\N	\N	2019-10-14 00:00:00+07	2519.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5808	\N	1	2518	\N	\N	2019-10-14 00:00:00+07	2518.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5807	\N	1	2517	\N	\N	2019-10-11 00:00:00+07	2517.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5805	\N	1	2515	\N	\N	2019-10-11 00:00:00+07	2515.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5804	\N	1	2514	\N	\N	2019-10-11 00:00:00+07	2514.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5803	\N	1	2513	\N	\N	2019-10-11 00:00:00+07	2513.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5802	\N	1	2511	\N	\N	2019-10-11 00:00:00+07	2511.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5801	\N	1	2510	\N	\N	2019-10-11 00:00:00+07	2510.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5800	\N	1	2512	\N	\N	2019-10-11 00:00:00+07	2512.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5799	\N	1	2509	\N	\N	2019-10-11 00:00:00+07	2509.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5798	\N	1	2508	\N	\N	2019-10-11 00:00:00+07	2508.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5797	\N	1	2507	\N	\N	2019-10-11 00:00:00+07	2507.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5795	\N	1	986	\N	\N	2019-10-10 00:00:00+07	0986.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5794	\N	1	2506	\N	\N	2019-10-10 00:00:00+07	2506.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5793	\N	1	2505	\N	\N	2019-10-10 00:00:00+07	2505.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5792	\N	1	2504	\N	\N	2019-10-10 00:00:00+07	2504.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5791	\N	1	2322	\N	\N	2019-10-10 00:00:00+07	2322.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5786	\N	1	2503	\N	\N	2019-10-09 00:00:00+07	2503.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5785	\N	1	2319	\N	\N	2019-10-09 00:00:00+07	2319.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5783	\N	1	530	\N	\N	2019-10-08 00:00:00+07	0530.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5782	\N	1	813	\N	\N	2019-10-08 00:00:00+07	0813.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5781	\N	1	812	\N	\N	2019-10-08 00:00:00+07	0812.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5763	\N	1	2307	\N	\N	2019-10-02 00:00:00+07	2307.65.101	2	2020-06-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5762	\N	1	808	\N	\N	2019-10-02 00:00:00+07	0808.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5760	\N	1	807	\N	\N	2019-10-02 00:00:00+07	0807.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5751	\N	1	2494	\N	\N	2019-09-26 00:00:00+07	2494.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5748	\N	1	2493	\N	\N	2019-09-26 00:00:00+07	2493.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-10-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3552	\N	1	1583	\N	\N	2017-04-01 00:00:00+07	1583.65.101	2	2018-05-07 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5731	\N	1	2291	\N	\N	2019-09-23 00:00:00+07	2291.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5729	\N	1	2289	\N	\N	2019-09-23 00:00:00+07	2289.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3553	\N	1	1556	\N	\N	2017-04-01 00:00:00+07	1556.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3340	\N	2	2029	\N	\N	2016-10-03 00:00:00+07	2029.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3556	\N	1	531	\N	\N	2017-03-01 00:00:00+07	0531.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3342	\N	1	487	\N	\N	2016-10-19 00:00:00+07	0487.65.101	2	2017-03-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3343	\N	1	332	\N	\N	2017-04-03 00:00:00+07	0332.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3344	\N	1	331	\N	\N	2017-04-03 00:00:00+07	0331.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3345	\N	1	259	\N	\N	2017-03-01 00:00:00+07	0259.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3346	\N	1	3072	\N	\N	2017-01-02 00:00:00+07	3072.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3348	\N	2	2018	\N	\N	2016-10-24 00:00:00+07	2018.62.101	2	2018-07-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3449	\N	1	29	\N	\N	2016-11-01 00:00:00+07	0029.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3352	\N	1	329	\N	\N	2016-02-01 00:00:00+07	0329.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3356	\N	1	436	\N	\N	2016-06-01 00:00:00+07	436.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3357	\N	1	314	\N	\N	2016-10-03 00:00:00+07	0314.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3360	\N	1	822	\N	\N	2016-05-02 00:00:00+07	0822.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3619	\N	1	2038	\N	\N	2017-03-01 00:00:00+07	2038.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3362	\N	1	2053	\N	\N	2017-01-02 00:00:00+07	2053.62.101	2	2018-10-31 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3370	\N	1	320	\N	\N	2016-08-02 00:00:00+07	0320.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3374	\N	1	2050	\N	\N	2016-12-05 00:00:00+07	2050.62.101	2	2017-05-31 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3375	\N	1	1553	\N	\N	2016-12-05 00:00:00+07	1553.65.101	2	2018-01-29 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3376	\N	1	1554	\N	\N	2017-01-02 00:00:00+07	1554.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3377	\N	1	323	\N	\N	2017-01-02 00:00:00+07	0323.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3378	\N	2	32	\N	\N	2016-11-02 00:00:00+07	0032.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3381	\N	2	1525	\N	\N	2017-01-02 00:00:00+07	1525.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3387	\N	1	475	\N	\N	2016-10-03 00:00:00+07	0475.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3388	\N	2	506	\N	\N	2016-12-05 00:00:00+07	0506.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3389	\N	1	317	\N	\N	2017-01-02 00:00:00+07	0317.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3390	\N	1	2019	\N	\N	2016-10-03 00:00:00+07	2019.62.101	2	2017-11-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3401	\N	1	2056	\N	\N	2017-01-02 00:00:00+07	2056.62.101	2	2018-09-09 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3404	\N	1	1582	\N	\N	2017-01-02 00:00:00+07	1582.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3406	\N	2	2011	\N	\N	2016-12-05 00:00:00+07	2011.62.101	2	2017-07-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3409	\N	1	2007	\N	\N	2017-01-02 00:00:00+07	2007.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3410	\N	1	2044	\N	\N	2017-01-03 00:00:00+07	2044.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3412	\N	1	499	\N	\N	2017-02-02 00:00:00+07	0499.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3414	\N	2	1572	\N	\N	2017-01-02 00:00:00+07	1572.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3415	\N	1	794	\N	\N	2017-01-02 00:00:00+07	0794.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3417	\N	1	2078	\N	\N	2017-01-02 00:00:00+07	2078.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3422	\N	2	1577	\N	\N	2017-03-01 00:00:00+07	1577.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5692	\N	1	2490	\N	\N	2019-09-18 00:00:00+07	2490.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5691	\N	1	2489	\N	\N	2019-09-18 00:00:00+07	2489.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3430	\N	1	828	\N	\N	2017-03-01 00:00:00+07	0828.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3431	\N	1	829	\N	\N	2017-03-01 00:00:00+07	0829.61.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3433	\N	1	2110	\N	\N	2017-03-01 00:00:00+07	2110.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3434	\N	1	2111	\N	\N	2017-03-01 00:00:00+07	2111.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3435	\N	1	2112	\N	\N	2017-03-01 00:00:00+07	2112.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5690	\N	1	2488	\N	\N	2019-09-18 00:00:00+07	2488.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3451	\N	1	1575	\N	\N	2017-04-03 00:00:00+07	1575.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3442	\N	1	1507	\N	\N	2017-01-03 00:00:00+07	1507.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3446	\N	1	2041	\N	\N	2017-01-04 00:00:00+07	2041.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3447	\N	1	2121	\N	\N	2017-03-01 00:00:00+07	2121.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3452	\N	1	1574	\N	\N	2017-04-03 00:00:00+07	1574.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3453	\N	1	1576	\N	\N	2017-04-03 00:00:00+07	1576.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3554	\N	1	2109	\N	\N	2017-04-03 00:00:00+07	2109.62.101	2	2018-10-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3492	\N	1	539	\N	\N	2017-03-01 00:00:00+07	0539.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3508	\N	1	14	\N	\N	2017-03-01 00:00:00+07	0014.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5685	\N	1	2487	\N	\N	2019-09-16 00:00:00+07	2487.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5683	\N	1	2262	\N	\N	2019-09-16 00:00:00+07	2262.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5682	\N	1	2261	\N	\N	2019-09-13 00:00:00+07	2261.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3513	\N	1	307	\N	\N	2017-03-01 00:00:00+07	0307.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3514	\N	1	24	\N	\N	2017-03-01 00:00:00+07	0024.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3515	\N	1	2049	\N	\N	2017-01-02 00:00:00+07	2049.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3519	\N	1	2	\N	\N	2017-04-27 00:00:00+07	0002.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3525	\N	1	1569	\N	\N	2017-03-01 00:00:00+07	1569.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3526	\N	1	2129	\N	\N	2017-03-01 00:00:00+07	2129.62.101	2	2019-08-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5676	\N	1	2485	\N	\N	2019-09-11 00:00:00+07	2485.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5675	\N	1	2257	\N	\N	2019-09-11 00:00:00+07	2257.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3534	\N	1	2115	\N	\N	2017-04-01 00:00:00+07	2115.62.101	2	2018-07-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3535	\N	1	2052	\N	\N	2017-05-03 00:00:00+07	2052.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5674	\N	1	2484	\N	\N	2019-09-11 00:00:00+07	2484.62.101	2	2020-08-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3537	\N	1	2137	\N	\N	2017-04-01 00:00:00+07	2137.62.101	2	2017-10-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3538	\N	1	542	\N	\N	2017-04-01 00:00:00+07	0542.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3557	\N	1	2105	\N	\N	2017-03-01 00:00:00+07	2105.62.101	2	2018-09-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3558	\N	1	2091	\N	\N	2017-03-01 00:00:00+07	2091.62.101	2	2017-09-14 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3540	\N	1	1617	\N	\N	2017-04-01 00:00:00+07	1617.65.101	2	2017-12-01 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3541	\N	1	66	\N	\N	2017-04-03 00:00:00+07	0066.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5670	\N	1	2255	\N	\N	2019-09-09 00:00:00+07	2255.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3546	\N	1	477	\N	\N	2017-04-01 00:00:00+07	0477.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5663	\N	1	518	\N	\N	2019-09-06 00:00:00+07	0518.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3550	\N	1	2090	\N	\N	2017-03-01 00:00:00+07	2090.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3551	\N	1	2098	\N	\N	2017-03-01 00:00:00+07	2098.62.101	1	2017-05-23 00:00:00+07	2018-07-02 00:00:00+07	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3560	\N	1	2097	\N	\N	2017-03-01 00:00:00+07	2097.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3566	\N	1	2082	\N	\N	2017-03-01 00:00:00+07	2082.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3568	\N	1	315	\N	\N	2017-01-02 00:00:00+07	0315.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3569	\N	1	15	\N	\N	2015-05-01 00:00:00+07	0015.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3575	\N	1	2047	\N	\N	2017-03-01 00:00:00+07	2047.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3576	\N	1	2048	\N	\N	2017-03-01 00:00:00+07	2048.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3579	\N	1	1379	\N	\N	2017-05-10 00:00:00+07	1379.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3583	\N	1	319	\N	\N	2017-03-01 00:00:00+07	0319.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5659	\N	1	2249	\N	\N	2019-09-04 00:00:00+07	2249.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5658	\N	1	2248	\N	\N	2019-09-04 00:00:00+07	2248.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3584	\N	1	2072	\N	\N	2017-03-01 00:00:00+07	2072.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3586	\N	1	835	\N	\N	2017-01-02 00:00:00+07	0835.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3644	\N	1	1560	\N	\N	2017-02-03 00:00:00+07	1560.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3587	\N	1	755	\N	\N	2017-01-02 00:00:00+07	0755.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3589	\N	1	762	\N	\N	2017-04-03 00:00:00+07	0762.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3645	\N	1	1567	\N	\N	2017-02-20 00:00:00+07	1567.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3591	\N	1	2127	\N	\N	2017-03-01 00:00:00+07	2127.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3646	\N	1	1570	\N	\N	2017-02-22 00:00:00+07	1570.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3592	\N	1	2080	\N	\N	2017-03-01 00:00:00+07	2080.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3597	\N	1	2039	\N	\N	2017-03-01 00:00:00+07	2039.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3598	\N	1	2079	\N	\N	2017-03-01 00:00:00+07	2079.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3599	\N	1	2116	\N	\N	2017-03-01 00:00:00+07	2116.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3648	\N	1	1580	\N	\N	2017-03-06 00:00:00+07	1580.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3600	\N	1	2117	\N	\N	2017-03-01 00:00:00+07	2117.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3601	\N	1	1610	\N	\N	2017-03-01 00:00:00+07	1610.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3650	\N	1	1586	\N	\N	2017-03-10 00:00:00+07	1586.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5657	\N	1	2247	\N	\N	2019-09-04 00:00:00+07	2247.65.101	2	2020-03-10 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3603	\N	1	1581	\N	\N	2017-03-01 00:00:00+07	1581.65.101	2	2017-11-26 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3604	\N	1	292	\N	\N	2017-03-01 00:00:00+07	0292.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3605	\N	1	798	\N	\N	2017-03-01 00:00:00+07	0798.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3606	\N	1	2104	\N	\N	2017-03-01 00:00:00+07	2104.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3653	\N	1	1599	\N	\N	2017-03-27 00:00:00+07	1599.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3607	\N	1	2042	\N	\N	2017-03-01 00:00:00+07	2042.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5656	\N	1	2246	\N	\N	2019-09-04 00:00:00+07	2246.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3662	\N	1	2051	\N	\N	2017-05-17 00:00:00+07	2051.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3664	\N	1	2067	\N	\N	2017-03-16 00:00:00+07	2067.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3665	\N	1	2073	\N	\N	2017-03-21 00:00:00+07	2073.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3666	\N	1	2077	\N	\N	2017-03-23 00:00:00+07	2077.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3669	\N	1	2081	\N	\N	2017-03-23 00:00:00+07	2081.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3670	\N	1	2085	\N	\N	2017-03-27 00:00:00+07	2085.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3671	\N	1	2089	\N	\N	2017-04-03 00:00:00+07	2089.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3672	\N	1	2092	\N	\N	2017-04-05 00:00:00+07	2092.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3673	\N	1	2093	\N	\N	2017-04-05 00:00:00+07	2093.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3674	\N	1	2096	\N	\N	2017-04-06 00:00:00+07	2096.62.101	2	2017-10-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5654	\N	1	799	\N	\N	2019-09-03 00:00:00+07	0799.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5653	\N	1	976	\N	\N	2019-09-02 00:00:00+07	0976.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3706	\N	1	2122	\N	\N	2017-05-18 00:00:00+07	2122.62.101	2	2018-07-11 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5650	\N	1	2245	\N	\N	2019-08-28 00:00:00+07	2245.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5649	\N	1	516	\N	\N	2019-08-28 00:00:00+07	0516.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3687	\N	1	495	\N	\N	2017-02-22 00:00:00+07	0495.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3688	\N	1	502	\N	\N	2017-03-17 00:00:00+07	0502.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5648	\N	1	2244	\N	\N	2019-08-28 00:00:00+07	2244.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5646	\N	1	798	\N	\N	2019-08-28 00:00:00+07	0798.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5645	\N	1	797	\N	\N	2019-08-28 00:00:00+07	0797.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3710	\N	1	2128	\N	\N	2017-05-18 00:00:00+07	2128.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3698	\N	1	20	\N	\N	2016-11-02 00:00:00+07	0020.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3699	\N	1	837	\N	\N	2016-11-02 00:00:00+07	0837.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5643	\N	1	2480	\N	\N	2019-08-22 00:00:00+07	2480.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3702	\N	1	2118	\N	\N	2017-04-13 00:00:00+07	2118.62.101	2	2017-06-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3712	\N	1	1	\N	\N	2017-05-02 00:00:00+07	001.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3714	\N	1	2133	\N	\N	2017-05-02 00:00:00+07	2133.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3716	\N	1	2131	\N	\N	2017-05-02 00:00:00+07	2131.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3717	\N	1	2132	\N	\N	2017-05-02 00:00:00+07	2132.62.101	2	2019-01-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3718	\N	1	838	\N	\N	2017-03-01 00:00:00+07	0838.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3719	\N	1	2138	\N	\N	2017-05-05 00:00:00+07	2138.62.101	1	2017-06-25 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3721	\N	1	2140	\N	\N	2017-05-10 00:00:00+07	2140.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3722	\N	2	2142	\N	\N	2017-05-18 00:00:00+07	2142.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3723	\N	2	2143	\N	\N	2017-05-10 00:00:00+07	2143.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3724	\N	1	2145	\N	\N	2017-05-17 00:00:00+07	2145.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3725	\N	1	2146	\N	\N	2017-05-17 00:00:00+07	2146.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3726	\N	1	2147	\N	\N	2017-05-17 00:00:00+07	2147.62.101	2	2018-06-04 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3729	\N	1	2055	\N	\N	2017-05-19 00:00:00+07	2055.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5638	\N	1	515	\N	\N	2019-08-21 00:00:00+07	0515.63.101	2	2020-01-28 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5636	\N	1	514	\N	\N	2019-08-20 00:00:00+07	0514.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3932	\N	1	2182	\N	\N	2017-08-08 00:00:00+07	2182.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3747	\N	1	548	\N	\N	2017-05-29 00:00:00+07	0548.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3749	\N	1	1628	\N	\N	2017-05-29 00:00:00+07	1628.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3750	\N	1	344	\N	\N	2017-05-30 00:00:00+07	0344.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3765	\N	1	347	\N	\N	2017-06-05 00:00:00+07	0347.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3751	\N	1	345	\N	\N	2017-05-31 00:00:00+07	0345.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5633	\N	1	974	\N	\N	2019-08-19 00:00:00+07	0974.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3734	\N	1	343	\N	\N	2017-05-23 00:00:00+07	0343.63.101	2	2018-06-14 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5630	\N	1	2237	\N	\N	2019-08-15 00:00:00+07	2237.65.101	1	\N	2020-05-30 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3752	\N	1	35	\N	\N	2017-05-01 00:00:00+07	0035.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3755	\N	1	1629	\N	\N	2017-05-31 00:00:00+07	1629.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3768	\N	1	2153	\N	\N	2017-06-05 00:00:00+07	2153.62.101	2	2021-03-24 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-05-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3769	\N	1	349	\N	\N	2017-06-05 00:00:00+07	0349.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5627	\N	1	795	\N	\N	2019-08-15 00:00:00+07	0795.64.101	1	\N	\N	OA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5622	\N	1	973	\N	\N	2019-08-14 00:00:00+07	0973.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5621	\N	1	2232	\N	\N	2019-08-14 00:00:00+07	2232.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5620	\N	1	972	\N	\N	2019-08-14 00:00:00+07	0972.61.101	2	2021-01-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-08-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5595	\N	1	967	\N	\N	2019-08-08 00:00:00+07	0967.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5594	\N	1	966	\N	\N	2019-08-08 00:00:00+07	0966.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3805	\N	1	291	\N	\N	2017-05-01 00:00:00+07	0291.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-16 00:00:00+07	\N	11:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3807	\N	1	284	\N	\N	2015-07-01 00:00:00+07	0284.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3811	\N	1	2158	\N	\N	2017-06-13 00:00:00+07	2158.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5590	\N	1	965	\N	\N	2019-08-08 00:00:00+07	0965.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3821	\N	1	83	\N	\N	2017-06-13 00:00:00+07	0083.00.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5587	\N	1	963	\N	\N	2019-08-08 00:00:00+07	0963.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3823	\N	2	84	\N	\N	2017-06-14 00:00:00+07	0084.00.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5582	\N	1	2472	\N	\N	2019-08-07 00:00:00+07	2472.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5581	\N	1	790	\N	\N	2019-08-07 00:00:00+07	0790.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3845	\N	1	564	\N	\N	2017-07-03 00:00:00+07	0564.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3846	\N	1	1635	\N	\N	2017-06-01 00:00:00+07	1635.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3847	\N	1	2163	\N	\N	2017-07-07 00:00:00+07	2163.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5563	\N	1	2215	\N	\N	2019-07-31 00:00:00+07	2215.65.101	1	\N	\N	OA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5562	\N	1	2214	\N	\N	2019-07-31 00:00:00+07	2214.65.101	1	\N	\N	OA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5561	\N	1	2213	\N	\N	2019-07-31 00:00:00+07	2213.65.101	1	\N	\N	OA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5558	\N	1	2463	\N	\N	2019-07-29 00:00:00+07	2463.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3849	\N	1	356	\N	\N	2017-07-07 00:00:00+07	0356.63.101	2	2018-02-21 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3850	\N	1	355	\N	\N	2017-07-07 00:00:00+07	0355.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3852	\N	1	2164	\N	\N	2017-07-10 00:00:00+07	2164.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3854	\N	1	358	\N	\N	2017-07-10 00:00:00+07	0358.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3857	\N	2	565	\N	\N	2017-07-13 00:00:00+07	0565.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3858	\N	1	2165	\N	\N	2017-07-13 00:00:00+07	2165.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5556	\N	1	2462	\N	\N	2019-07-25 00:00:00+07	2462.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5555	\N	1	2461	\N	\N	2019-07-25 00:00:00+07	2461.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5554	\N	1	2460	\N	\N	2019-07-25 00:00:00+07	2460.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3893	\N	1	364	\N	\N	2017-07-21 00:00:00+07	0364.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3903	\N	1	362	\N	\N	2017-07-25 00:00:00+07	0362.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3905	\N	1	361	\N	\N	2017-07-25 00:00:00+07	0361.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3906	\N	1	2171	\N	\N	2017-07-26 00:00:00+07	2171.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3907	\N	1	2172	\N	\N	2017-07-26 00:00:00+07	2172.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3919	\N	1	2174	\N	\N	2017-07-27 00:00:00+07	2174.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5546	\N	1	2454	\N	\N	2019-07-24 00:00:00+07	2454.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3921	\N	1	2175	\N	\N	2017-07-31 00:00:00+07	2175.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5545	\N	1	2453	\N	\N	2019-07-24 00:00:00+07	2453.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3923	\N	1	2176	\N	\N	2017-08-01 00:00:00+07	2176.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5544	\N	1	504	\N	\N	2019-07-24 00:00:00+07	0504.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3927	\N	1	365	\N	\N	2017-08-02 00:00:00+07	0365.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5543	\N	1	2452	\N	\N	2019-07-24 00:00:00+07	2452.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5579	\N	1	2471	\N	\N	2019-08-05 00:00:00+07	2471.62.101	2	2020-04-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3996	\N	1	2191	\N	\N	2017-09-12 00:00:00+07	2191.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3998	\N	1	2192	\N	\N	2017-09-12 00:00:00+07	2192.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3999	\N	1	2193	\N	\N	2017-09-12 00:00:00+07	2193.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5542	\N	1	2451	\N	\N	2019-07-24 00:00:00+07	2451.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4040	\N	1	369	\N	\N	2017-09-14 00:00:00+07	0369.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4137	\N	1	1193	\N	\N	2008-01-01 00:00:00+07	1193.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5538	\N	1	2450	\N	\N	2019-07-24 00:00:00+07	2450.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4042	\N	1	575	\N	\N	2017-09-18 00:00:00+07	0575.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5537	\N	1	2209	\N	\N	2019-07-23 00:00:00+07	2209.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4043	\N	1	576	\N	\N	2017-09-18 00:00:00+07	0576.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5536	\N	1	2208	\N	\N	2019-07-23 00:00:00+07	2208.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5535	\N	1	2207	\N	\N	2019-07-23 00:00:00+07	2207.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5534	\N	1	2449	\N	\N	2019-07-23 00:00:00+07	2449.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5533	\N	1	2448	\N	\N	2019-07-23 00:00:00+07	2448.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5532	\N	1	2447	\N	\N	2019-07-23 00:00:00+07	2447.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5531	\N	1	2446	\N	\N	2019-07-23 00:00:00+07	2446.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5530	\N	1	2445	\N	\N	2019-07-23 00:00:00+07	2445.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4048	\N	1	579	\N	\N	2017-09-18 00:00:00+07	0579.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5529	\N	1	2444	\N	\N	2019-07-22 00:00:00+07	2444.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4138	\N	1	1190	\N	\N	2007-08-01 00:00:00+07	1190.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4128	\N	1	1352	\N	\N	2009-05-01 00:00:00+07	1352.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4052	\N	1	370	\N	\N	2017-09-20 00:00:00+07	0370.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4130	\N	1	1301	\N	\N	2015-08-28 00:00:00+07	1301.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4131	\N	1	584	\N	\N	2017-09-27 00:00:00+07	0584.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5527	\N	1	2443	\N	\N	2019-07-22 00:00:00+07	2443.62.101	2	2020-03-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-03-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5526	\N	1	2442	\N	\N	2019-07-22 00:00:00+07	2442.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4073	\N	1	2196	\N	\N	2017-09-20 00:00:00+07	2196.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4074	\N	1	3333	\N	\N	2017-09-20 00:00:00+07	3333.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4142	\N	1	372	\N	\N	2017-10-05 00:00:00+07	0372.63.101	2	2021-01-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4136	\N	1	1188	\N	\N	2007-01-01 00:00:00+07	1188.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5522	\N	1	961	\N	\N	2019-07-19 00:00:00+07	0961.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5521	\N	1	2440	\N	\N	2019-07-18 00:00:00+07	2440.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5520	\N	1	2439	\N	\N	2019-07-18 00:00:00+07	2439.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5519	\N	1	2438	\N	\N	2019-07-18 00:00:00+07	2438.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4143	\N	1	2202	\N	\N	2017-10-05 00:00:00+07	2202.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4144	\N	1	2203	\N	\N	2017-10-05 00:00:00+07	2203.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4145	\N	1	2204	\N	\N	2017-10-05 00:00:00+07	2204.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4146	\N	1	373	\N	\N	2017-10-05 00:00:00+07	0373.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5518	\N	1	2437	\N	\N	2019-07-18 00:00:00+07	2437.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5515	\N	1	2436	\N	\N	2019-07-17 00:00:00+07	2436.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4151	\N	1	374	\N	\N	2017-10-06 00:00:00+07	0374.63.101	2	2017-11-27 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5510	\N	1	2434	\N	\N	2019-07-17 00:00:00+07	2434.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5509	\N	1	2433	\N	\N	2019-07-17 00:00:00+07	2433.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5503	\N	1	2431	\N	\N	2019-07-15 00:00:00+07	2431.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4163	\N	1	205	\N	\N	2012-04-01 00:00:00+07	0205.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4165	\N	1	926	\N	\N	2006-01-01 00:00:00+07	0926.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5502	\N	1	2430	\N	\N	2019-07-15 00:00:00+07	2430.62.101	2	2019-10-15 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5501	\N	1	2429	\N	\N	2019-07-15 00:00:00+07	2429.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4168	\N	1	590	\N	\N	2017-10-16 00:00:00+07	0590.64.101	2	2020-05-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4189	\N	1	591	\N	\N	2017-10-19 00:00:00+07	0591.64.101	2	\N	\N	SA	\N	10112	solusiti	\N	\N	f	f	2021-01-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4190	\N	1	376	\N	\N	2017-10-19 00:00:00+07	0376.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4191	\N	1	592	\N	\N	2017-10-19 00:00:00+07	0592.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4192	\N	1	593	\N	\N	2017-10-19 00:00:00+07	0593.64.101	2	2018-02-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4193	\N	1	594	\N	\N	2017-10-19 00:00:00+07	0594.64.101	2	2020-02-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5500	\N	1	785	\N	\N	2019-07-15 00:00:00+07	0785.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4205	\N	1	3	\N	\N	2006-12-01 00:00:00+07	0003.63.101	2	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4206	\N	1	141	\N	\N	2008-06-02 00:00:00+07	0141.63.101	2	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4207	\N	1	6	\N	\N	2009-08-01 00:00:00+07	0006.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4208	\N	1	198	\N	\N	2011-11-01 00:00:00+07	0198.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4209	\N	1	596	\N	\N	2017-10-20 00:00:00+07	0596.64.101	2	2019-12-01 00:00:00+07	\N	SA	\N	10112	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5499	\N	1	2428	\N	\N	2019-07-15 00:00:00+07	2428.62.101	2	2019-09-28 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4213	\N	2	860	\N	\N	2017-10-23 00:00:00+07	0860.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4215	\N	1	170	\N	\N	2005-01-01 00:00:00+07	0170.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4216	\N	1	764	\N	\N	2009-09-01 00:00:00+07	0764.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4219	\N	1	773	\N	\N	2009-01-01 00:00:00+07	0773.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4224	\N	1	1	\N	\N	2008-08-01 00:00:00+07	001	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5498	\N	1	2427	\N	\N	2019-07-12 00:00:00+07	2427.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4235	\N	1	696	\N	\N	2008-07-01 00:00:00+07	0696.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5497	\N	1	2426	\N	\N	2019-07-12 00:00:00+07	2426.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4246	\N	1	1280	\N	\N	2009-01-01 00:00:00+07	1280.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4247	\N	1	597	\N	\N	2017-10-31 00:00:00+07	0597.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5496	\N	1	2425	\N	\N	2019-07-11 00:00:00+07	2425.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4251	\N	1	1291	\N	\N	2009-02-01 00:00:00+07	1291.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4253	\N	1	1304	\N	\N	2009-01-01 00:00:00+07	1304.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5494	\N	1	2424	\N	\N	2019-07-11 00:00:00+07	2424.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4257	\N	1	1320	\N	\N	2009-05-01 00:00:00+07	1320.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4258	\N	1	1262	\N	\N	2009-06-01 00:00:00+07	1262.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4259	\N	1	1	\N	\N	2009-03-01 00:00:00+07	001	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4261	\N	1	201	\N	\N	2009-03-01 00:00:00+07	0201..62	2	2018-05-16 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4262	\N	1	963	\N	\N	2006-11-01 00:00:00+07	0963.62.101	2	2018-05-16 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4264	\N	1	3	\N	\N	2007-05-01 00:00:00+07	0003.64.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5493	\N	1	2423	\N	\N	2019-07-10 00:00:00+07	2423.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5492	\N	1	2422	\N	\N	2019-07-10 00:00:00+07	2422.62.101	1	\N	2020-06-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5491	\N	1	2197	\N	\N	2019-07-10 00:00:00+07	2197.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5489	\N	1	2196	\N	\N	2019-07-10 00:00:00+07	2196.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4274	\N	1	601	\N	\N	2017-11-14 00:00:00+07	0601.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4276	\N	1	28	\N	\N	2008-01-01 00:00:00+07	0028.61.101	2	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5487	\N	1	2195	\N	\N	2019-07-10 00:00:00+07	2195.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4280	\N	1	863	\N	\N	2017-11-24 00:00:00+07	0863.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5482	\N	1	959	\N	\N	2019-07-09 00:00:00+07	0959.61.101	1	\N	\N	OA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4283	\N	1	865	\N	\N	2017-11-29 00:00:00+07	0865.61.101	2	2019-10-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4289	\N	1	411	\N	\N	2008-05-01 00:00:00+07	0411.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4290	\N	1	37	\N	\N	2005-11-01 00:00:00+07	0037.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4297	\N	1	492	\N	\N	2009-07-01 00:00:00+07	0492.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4298	\N	1	178	\N	\N	2006-02-01 00:00:00+07	0178.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4302	\N	1	541	\N	\N	2011-02-01 00:00:00+07	0541.65.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5480	\N	1	782	\N	\N	2019-07-08 00:00:00+07	0782.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4305	\N	1	39	\N	\N	2009-12-11 00:00:00+07	0039.61.101	2	2018-05-21 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4307	\N	1	471	\N	\N	2010-03-01 00:00:00+07	0471.61.101	2	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4308	\N	1	36	\N	\N	2007-05-01 00:00:00+07	0036.61.101	2	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4311	\N	1	1677	\N	\N	2017-12-13 00:00:00+07	1677.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4313	\N	1	200	\N	\N	2010-02-01 00:00:00+07	0200.64.101	2	2018-07-18 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4314	\N	1	15	\N	\N	2007-04-02 00:00:00+07	0015.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4315	\N	1	14	\N	\N	2009-02-02 00:00:00+07	0014.64.101	2	2018-07-18 00:00:00+07	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5475	\N	1	781	\N	\N	2019-07-03 00:00:00+07	0781.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5474	\N	1	780	\N	\N	2019-07-03 00:00:00+07	0780.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5472	\N	1	779	\N	\N	2019-07-03 00:00:00+07	0779.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5471	\N	1	778	\N	\N	2019-07-03 00:00:00+07	0778.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5470	\N	1	777	\N	\N	2019-07-03 00:00:00+07	0777.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4329	\N	1	2215	\N	\N	2017-12-27 00:00:00+07	2215.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5466	\N	1	2189	\N	\N	2019-06-28 00:00:00+07	2189.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5465	\N	1	2419	\N	\N	2019-06-28 00:00:00+07	2419.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5464	\N	1	2418	\N	\N	2019-06-28 00:00:00+07	2418.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5463	\N	1	2417	\N	\N	2019-06-28 00:00:00+07	2417.62.101	2	2021-02-28 00:00:00+07	2020-05-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5462	\N	1	2416	\N	\N	2019-06-28 00:00:00+07	2416.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5461	\N	1	2415	\N	\N	2019-06-28 00:00:00+07	2415.62.101	2	2019-02-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4336	\N	1	1979	\N	\N	2018-01-05 00:00:00+07	1979.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4337	\N	1	2216	\N	\N	2018-01-08 00:00:00+07	2216.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4344	\N	1	2218	\N	\N	2018-01-11 00:00:00+07	2218.62.101	2	2020-01-30 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5456	\N	1	2413	\N	\N	2019-06-27 00:00:00+07	2413.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5455	\N	1	2412	\N	\N	2019-06-27 00:00:00+07	2412.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4351	\N	1	2221	\N	\N	2018-01-16 00:00:00+07	2221.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4357	\N	1	285	\N	\N	2018-01-23 00:00:00+07	0285.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4358	\N	1	2224	\N	\N	2018-01-24 00:00:00+07	2224.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5449	\N	1	955	\N	\N	2019-06-26 00:00:00+07	0955.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4369	\N	1	2226	\N	\N	2018-01-30 00:00:00+07	2226.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-10-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4370	\N	1	2227	\N	\N	2018-01-31 00:00:00+07	2227.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4378	\N	1	2228	\N	\N	2018-02-01 00:00:00+07	2228.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4393	\N	1	2230	\N	\N	2018-02-06 00:00:00+07	2230.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4395	\N	1	2232	\N	\N	2018-02-07 00:00:00+07	2232.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4400	\N	1	1690	\N	\N	2018-02-12 00:00:00+07	1690.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4401	\N	1	2233	\N	\N	2018-02-12 00:00:00+07	2233.62.101	2	2021-12-25 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-04-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4415	\N	1	1694	\N	\N	2018-02-28 00:00:00+07	1694.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4416	\N	1	404	\N	\N	2018-02-28 00:00:00+07	0404.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5435	\N	1	2169	\N	\N	2019-06-21 00:00:00+07	2169.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4428	\N	1	1698	\N	\N	2018-03-07 00:00:00+07	1698.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4430	\N	1	1700	\N	\N	2018-03-08 00:00:00+07	1700.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5434	\N	1	2411	\N	\N	2019-06-21 00:00:00+07	2411.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5433	\N	1	2410	\N	\N	2019-06-21 00:00:00+07	2410.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5432	\N	1	776	\N	\N	2019-06-21 00:00:00+07	0776.64.101	2	2019-06-27 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4435	\N	1	1701	\N	\N	2018-03-12 00:00:00+07	1701.65.101	2	2020-03-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4438	\N	1	333	\N	\N	2018-03-12 00:00:00+07	0333.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4444	\N	1	2240	\N	\N	2018-03-14 00:00:00+07	2240.62.101	1	2020-11-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4445	\N	1	2241	\N	\N	2018-03-14 00:00:00+07	2241.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4449	\N	1	1704	\N	\N	2018-03-15 00:00:00+07	1704.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5431	\N	1	775	\N	\N	2019-06-21 00:00:00+07	0775.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4452	\N	1	1706	\N	\N	2018-03-19 00:00:00+07	1706.65.101	2	2018-08-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4453	\N	1	405	\N	\N	2018-03-19 00:00:00+07	0405.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5430	\N	1	774	\N	\N	2019-06-21 00:00:00+07	0774.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4456	\N	1	1708	\N	\N	2018-03-21 00:00:00+07	1708.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4457	\N	1	2246	\N	\N	2018-03-21 00:00:00+07	2246.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4460	\N	1	642	\N	\N	2018-03-22 00:00:00+07	0642.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4461	\N	1	643	\N	\N	2018-03-22 00:00:00+07	0643.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4462	\N	1	644	\N	\N	2018-03-22 00:00:00+07	0644.64.101	2	2018-11-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4463	\N	1	645	\N	\N	2018-03-22 00:00:00+07	0645.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5429	\N	1	2168	\N	\N	2019-06-21 00:00:00+07	2168.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4464	\N	1	646	\N	\N	2018-03-22 00:00:00+07	0646.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5428	\N	1	498	\N	\N	2019-06-21 00:00:00+07	0498.63.101	1	\N	\N	OA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4466	\N	1	647	\N	\N	2018-03-22 00:00:00+07	0647.64.101	1	\N	2018-08-01 00:00:00+07	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4477	\N	1	406	\N	\N	2018-03-28 00:00:00+07	0406.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4479	\N	1	407	\N	\N	2018-03-28 00:00:00+07	0407.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4483	\N	1	654	\N	\N	2018-04-05 00:00:00+07	0654.64.101	2	2019-11-07 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4484	\N	1	2248	\N	\N	2018-04-05 00:00:00+07	2248.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4485	\N	1	655	\N	\N	2018-04-05 00:00:00+07	0655.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4489	\N	1	1714	\N	\N	2018-04-05 00:00:00+07	1714.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4503	\N	1	1719	\N	\N	2018-04-12 00:00:00+07	1719.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4505	\N	1	411	\N	\N	2018-04-13 00:00:00+07	0411.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5421	\N	1	2408	\N	\N	2019-06-19 00:00:00+07	2408.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5420	\N	1	2163	\N	\N	2019-06-19 00:00:00+07	2163.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5418	\N	1	2162	\N	\N	2019-06-19 00:00:00+07	2162.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4516	\N	1	1721	\N	\N	2018-04-17 00:00:00+07	1721.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4518	\N	1	2254	\N	\N	2018-04-17 00:00:00+07	2254.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4520	\N	1	1724	\N	\N	2018-04-18 00:00:00+07	1724.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5417	\N	1	2161	\N	\N	2019-06-18 00:00:00+07	2161.65.101	2	2020-09-01 00:00:00+07	2020-06-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-04-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4530	\N	1	1731	\N	\N	2018-04-23 00:00:00+07	1731.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4523	\N	1	1727	\N	\N	2018-04-18 00:00:00+07	1727.65.101	1	\N	2020-06-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4525	\N	1	1728	\N	\N	2018-04-19 00:00:00+07	1728.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5416	\N	1	2160	\N	\N	2019-06-18 00:00:00+07	2160.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5415	\N	1	2159	\N	\N	2019-06-18 00:00:00+07	2159.65.101	2	2020-03-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4527	\N	1	1729	\N	\N	2018-04-19 00:00:00+07	1729.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4528	\N	1	1730	\N	\N	2018-04-19 00:00:00+07	1730.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5414	\N	1	2158	\N	\N	2019-06-18 00:00:00+07	2158.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4537	\N	1	665	\N	\N	2018-04-30 00:00:00+07	0665.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4542	\N	1	1735	\N	\N	2018-05-03 00:00:00+07	1735.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4543	\N	1	1736	\N	\N	2018-05-04 00:00:00+07	1736.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4561	\N	1	1743	\N	\N	2018-05-15 00:00:00+07	1743.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4562	\N	1	2260	\N	\N	2018-05-15 00:00:00+07	2260.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4563	\N	1	1744	\N	\N	2018-05-15 00:00:00+07	1744.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5412	\N	1	951	\N	\N	2019-06-17 00:00:00+07	0951.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4566	\N	1	419	\N	\N	2018-05-16 00:00:00+07	0419.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4567	\N	1	668	\N	\N	2018-05-16 00:00:00+07	0668.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4568	\N	1	1745	\N	\N	2018-05-16 00:00:00+07	1745.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4569	\N	1	669	\N	\N	2018-05-16 00:00:00+07	0669.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4573	\N	1	670	\N	\N	2018-05-17 00:00:00+07	0670.64.101	2	2020-03-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5411	\N	1	2156	\N	\N	2019-06-14 00:00:00+07	2156.65.101	2	2020-12-10 00:00:00+07	2020-07-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-02-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5410	\N	1	2155	\N	\N	2019-06-14 00:00:00+07	2155.65.101	1	\N	2020-07-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5409	\N	1	2154	\N	\N	2019-06-14 00:00:00+07	2154.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4580	\N	1	672	\N	\N	2018-06-05 00:00:00+07	0672.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4583	\N	1	2262	\N	\N	2018-06-07 00:00:00+07	2262.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5408	\N	1	2153	\N	\N	2019-06-14 00:00:00+07	2153.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4585	\N	1	2263	\N	\N	2018-06-07 00:00:00+07	2263.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5407	\N	1	2152	\N	\N	2019-06-14 00:00:00+07	2152.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4593	\N	1	1753	\N	\N	2018-07-06 00:00:00+07	1753.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4594	\N	1	1754	\N	\N	2018-07-06 00:00:00+07	1754.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4601	\N	1	1761	\N	\N	2018-07-12 00:00:00+07	1761.65.101	2	2018-08-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4602	\N	1	1762	\N	\N	2018-07-12 00:00:00+07	1762.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4603	\N	1	675	\N	\N	2018-07-12 00:00:00+07	0675.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5406	\N	1	2407	\N	\N	2019-06-13 00:00:00+07	2407.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5405	\N	1	2406	\N	\N	2019-06-13 00:00:00+07	2406.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5404	\N	1	2405	\N	\N	2019-06-13 00:00:00+07	2405.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4609	\N	1	882	\N	\N	2018-07-13 00:00:00+07	0882.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4612	\N	1	2266	\N	\N	2018-07-16 00:00:00+07	2266.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4613	\N	1	1766	\N	\N	2018-07-18 00:00:00+07	1766.65.101	2	2020-05-01 00:00:00+07	2020-05-01 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5403	\N	1	2404	\N	\N	2019-06-13 00:00:00+07	2404.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5402	\N	1	2151	\N	\N	2019-06-13 00:00:00+07	2151.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4617	\N	1	1768	\N	\N	2018-07-23 00:00:00+07	1768.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4621	\N	1	2267	\N	\N	2018-07-25 00:00:00+07	2267.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4622	\N	1	2268	\N	\N	2018-07-25 00:00:00+07	2268.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4623	\N	1	883	\N	\N	2018-07-25 00:00:00+07	0883.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4625	\N	1	682	\N	\N	2018-07-26 00:00:00+07	0682.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4626	\N	1	1770	\N	\N	2018-07-26 00:00:00+07	1770.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5400	\N	1	2150	\N	\N	2019-06-13 00:00:00+07	2150.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4630	\N	1	1773	\N	\N	2018-07-31 00:00:00+07	1773.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4632	\N	1	2270	\N	\N	2018-08-01 00:00:00+07	2270.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4637	\N	1	426	\N	\N	2018-08-02 00:00:00+07	0426.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4638	\N	1	1778	\N	\N	2018-08-03 00:00:00+07	1778.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4639	\N	1	1779	\N	\N	2018-08-03 00:00:00+07	1779.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4640	\N	1	2271	\N	\N	2018-08-06 00:00:00+07	2271.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4641	\N	1	684	\N	\N	2018-08-06 00:00:00+07	0684.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4642	\N	1	884	\N	\N	2018-08-06 00:00:00+07	0884.61.101	2	2019-06-30 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5398	\N	1	2149	\N	\N	2019-06-13 00:00:00+07	2149.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5397	\N	1	2148	\N	\N	2019-06-13 00:00:00+07	2148.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4648	\N	1	1781	\N	\N	2018-08-08 00:00:00+07	1781.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4650	\N	1	1782	\N	\N	2018-08-09 00:00:00+07	1782.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4651	\N	1	1783	\N	\N	2018-08-09 00:00:00+07	1783.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4653	\N	1	1784	\N	\N	2018-08-09 00:00:00+07	1784.65.101	2	2020-05-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-09-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4654	\N	1	1785	\N	\N	2018-08-09 00:00:00+07	1785.65.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-06-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4657	\N	1	2275	\N	\N	2018-08-10 00:00:00+07	2275.62.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4660	\N	1	1787	\N	\N	2018-08-13 00:00:00+07	1787.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4666	\N	1	885	\N	\N	2018-08-21 00:00:00+07	0885.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4667	\N	1	2278	\N	\N	2018-08-23 00:00:00+07	2278.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4668	\N	1	691	\N	\N	2018-08-23 00:00:00+07	0691.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4669	\N	1	1788	\N	\N	2018-08-23 00:00:00+07	1788.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4671	\N	1	1789	\N	\N	2018-08-27 00:00:00+07	1789.65.101	2	2019-10-08 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4672	\N	1	1790	\N	\N	2018-08-27 00:00:00+07	1790.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4673	\N	1	2279	\N	\N	2018-08-27 00:00:00+07	2279.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4674	\N	1	692	\N	\N	2018-08-28 00:00:00+07	0692.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4686	\N	1	693	\N	\N	2018-08-30 00:00:00+07	0693.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5393	\N	1	949	\N	\N	2019-06-13 00:00:00+07	0949.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4693	\N	1	1805	\N	\N	2018-09-05 00:00:00+07	1805.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5392	\N	1	948	\N	\N	2019-06-12 00:00:00+07	0948.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4728	\N	1	2285	\N	\N	2018-09-13 00:00:00+07	2285.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4732	\N	1	2287	\N	\N	2018-09-19 00:00:00+07	2287.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4733	\N	1	2288	\N	\N	2018-09-19 00:00:00+07	2288.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4734	\N	1	1819	\N	\N	2018-09-19 00:00:00+07	1819.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4735	\N	1	2289	\N	\N	2018-09-19 00:00:00+07	2289.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4737	\N	1	1821	\N	\N	2018-09-19 00:00:00+07	1821.65.101	2	2019-04-03 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4738	\N	1	711	\N	\N	2018-09-20 00:00:00+07	0711.64.101	2	2019-08-27 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4780	\N	1	1857	\N	\N	2018-10-04 00:00:00+07	1857.65.101	2	2019-02-01 00:00:00+07	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4781	\N	1	1858	\N	\N	2018-10-04 00:00:00+07	1858.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4783	\N	1	437	\N	\N	2018-10-05 00:00:00+07	0437.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4788	\N	1	1860	\N	\N	2018-10-11 00:00:00+07	1860.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4790	\N	1	1861	\N	\N	2018-10-11 00:00:00+07	1861.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4791	\N	1	1862	\N	\N	2018-10-11 00:00:00+07	1862.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4828	\N	1	2294	\N	\N	2018-10-25 00:00:00+07	2294.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4833	\N	1	1895	\N	\N	2018-11-05 00:00:00+07	1895.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4834	\N	1	1896	\N	\N	2018-11-05 00:00:00+07	1896.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4835	\N	1	720	\N	\N	2018-11-05 00:00:00+07	0720.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4836	\N	1	2296	\N	\N	2018-11-05 00:00:00+07	2296.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4837	\N	1	721	\N	\N	2018-11-05 00:00:00+07	0721.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4841	\N	1	2298	\N	\N	2018-11-07 00:00:00+07	2298.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4842	\N	1	1897	\N	\N	2018-11-07 00:00:00+07	1897.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4844	\N	1	1899	\N	\N	2018-11-13 00:00:00+07	1899.65.101	2	2019-06-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4846	\N	1	722	\N	\N	2018-11-21 00:00:00+07	0722.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4847	\N	1	723	\N	\N	2018-11-21 00:00:00+07	0723.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4848	\N	1	724	\N	\N	2018-11-21 00:00:00+07	0724.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4849	\N	1	725	\N	\N	2018-11-21 00:00:00+07	0725.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4850	\N	1	726	\N	\N	2018-11-21 00:00:00+07	0726.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4852	\N	1	442	\N	\N	2018-11-26 00:00:00+07	0442.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4853	\N	1	443	\N	\N	2018-11-26 00:00:00+07	0443.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4854	\N	1	727	\N	\N	2018-11-26 00:00:00+07	0727.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4855	\N	1	2299	\N	\N	2018-11-26 00:00:00+07	2299.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4856	\N	1	444	\N	\N	2018-11-26 00:00:00+07	0444.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4857	\N	1	445	\N	\N	2018-11-26 00:00:00+07	0445.63.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4858	\N	1	446	\N	\N	2018-11-26 00:00:00+07	0446.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4859	\N	1	1900	\N	\N	2018-11-27 00:00:00+07	1900.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4860	\N	1	1901	\N	\N	2018-11-27 00:00:00+07	1901.65.101	2	2019-03-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4862	\N	1	2300	\N	\N	2018-11-27 00:00:00+07	2300.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4863	\N	1	2301	\N	\N	2018-11-27 00:00:00+07	2301.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4865	\N	1	1904	\N	\N	2018-11-29 00:00:00+07	1904.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4866	\N	1	2302	\N	\N	2018-11-29 00:00:00+07	2302.62.101	1	\N	\N	OA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4867	\N	1	2303	\N	\N	2018-11-29 00:00:00+07	2303.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4868	\N	1	728	\N	\N	2018-11-29 00:00:00+07	0728.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4869	\N	1	1905	\N	\N	2018-11-29 00:00:00+07	1905.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4870	\N	1	1906	\N	\N	2018-11-29 00:00:00+07	1906.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4871	\N	1	1907	\N	\N	2018-11-30 00:00:00+07	1907.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4872	\N	1	1908	\N	\N	2018-11-30 00:00:00+07	1908.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4873	\N	1	729	\N	\N	2018-11-30 00:00:00+07	0729.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-03-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4874	\N	1	730	\N	\N	2018-11-30 00:00:00+07	0730.64.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4875	\N	1	1909	\N	\N	2018-11-30 00:00:00+07	1909.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4878	\N	1	1910	\N	\N	2018-12-04 00:00:00+07	1910.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4879	\N	1	1911	\N	\N	2018-12-04 00:00:00+07	1911.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4880	\N	1	1912	\N	\N	2018-12-04 00:00:00+07	1912.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4881	\N	1	1913	\N	\N	2018-12-04 00:00:00+07	1913.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4882	\N	1	1914	\N	\N	2018-12-04 00:00:00+07	1914.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4883	\N	1	731	\N	\N	2018-12-05 00:00:00+07	0731.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4884	\N	1	448	\N	\N	2018-12-05 00:00:00+07	0448.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4885	\N	1	449	\N	\N	2018-12-05 00:00:00+07	0449.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4886	\N	1	450	\N	\N	2018-12-05 00:00:00+07	0450.63.101	1	\N	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-11-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4887	\N	1	451	\N	\N	2018-12-05 00:00:00+07	0451.63.101	2	2019-05-01 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4888	\N	1	2305	\N	\N	2018-12-06 00:00:00+07	2305.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4890	\N	1	2307	\N	\N	2018-12-06 00:00:00+07	2307.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4895	\N	1	2309	\N	\N	2018-12-06 00:00:00+07	2309.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4901	\N	1	1920	\N	\N	2018-12-06 00:00:00+07	1920.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4902	\N	1	1921	\N	\N	2018-12-06 00:00:00+07	1921.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4903	\N	1	452	\N	\N	2018-12-07 00:00:00+07	0452.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4904	\N	1	453	\N	\N	2018-12-07 00:00:00+07	0453.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4905	\N	1	2311	\N	\N	2018-12-12 00:00:00+07	2311.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4906	\N	1	1922	\N	\N	2018-12-13 00:00:00+07	1922.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4907	\N	1	734	\N	\N	2018-12-13 00:00:00+07	0734.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5384	\N	1	494	\N	\N	2019-06-11 00:00:00+07	0494.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5383	\N	1	493	\N	\N	2019-06-11 00:00:00+07	0493.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4908	\N	1	1923	\N	\N	2018-12-13 00:00:00+07	1923.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5381	\N	1	2138	\N	\N	2019-06-11 00:00:00+07	2138.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4911	\N	1	1924	\N	\N	2018-12-17 00:00:00+07	1924.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5380	\N	1	2137	\N	\N	2019-06-11 00:00:00+07	2137.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4921	\N	1	1927	\N	\N	2019-01-02 00:00:00+07	1927.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5379	\N	1	2136	\N	\N	2019-06-11 00:00:00+07	2136.65.101	2	2019-06-12 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4924	\N	1	1928	\N	\N	2019-01-09 00:00:00+07	1928.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5378	\N	1	2135	\N	\N	2019-06-11 00:00:00+07	2135.65.101	2	2019-06-12 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5377	\N	1	2134	\N	\N	2019-06-11 00:00:00+07	2134.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5376	\N	1	2133	\N	\N	2019-06-11 00:00:00+07	2133.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4930	\N	1	1929	\N	\N	2019-01-21 00:00:00+07	1929.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5375	\N	1	2132	\N	\N	2019-06-11 00:00:00+07	2132.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5373	\N	1	947	\N	\N	2019-06-11 00:00:00+07	0947.61.101	2	2020-10-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2020-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5369	\N	1	2130	\N	\N	2019-05-31 00:00:00+07	2130.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5368	\N	1	944	\N	\N	2019-05-29 00:00:00+07	0944.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5367	\N	1	943	\N	\N	2019-05-29 00:00:00+07	0943.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5366	\N	1	942	\N	\N	2019-05-29 00:00:00+07	0942.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4947	\N	1	1936	\N	\N	2019-02-04 00:00:00+07	1936.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4950	\N	1	457	\N	\N	2019-02-07 00:00:00+07	0457.63.101	1	\N	2020-05-20 00:00:00+07	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4959	\N	1	900	\N	\N	2019-02-11 00:00:00+07	0900.61.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5365	\N	1	771	\N	\N	2019-05-28 00:00:00+07	0771.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5364	\N	1	770	\N	\N	2019-05-28 00:00:00+07	0770.64.101	2	2020-06-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5362	\N	1	941	\N	\N	2019-05-28 00:00:00+07	0941.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5026	\N	1	1981	\N	\N	2019-02-26 00:00:00+07	1981.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5030	\N	1	749	\N	\N	2019-02-26 00:00:00+07	0749.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5354	\N	1	2125	\N	\N	2019-05-27 00:00:00+07	2125.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5349	\N	1	769	\N	\N	2019-05-24 00:00:00+07	0769.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5038	\N	1	750	\N	\N	2019-02-27 00:00:00+07	0750.64.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5346	\N	1	2398	\N	\N	2019-05-23 00:00:00+07	2398.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4982	\N	1	1951	\N	\N	2019-02-14 00:00:00+07	1951.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5341	\N	1	490	\N	\N	2019-05-22 00:00:00+07	0490.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5340	\N	1	2397	\N	\N	2019-05-22 00:00:00+07	2397.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5050	\N	1	1993	\N	\N	2019-02-28 00:00:00+07	1993.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5339	\N	1	2395	\N	\N	2019-05-22 00:00:00+07	2395.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5338	\N	1	2119	\N	\N	2019-05-22 00:00:00+07	2119.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5052	\N	1	463	\N	\N	2019-02-28 00:00:00+07	0463.63.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5053	\N	1	1995	\N	\N	2019-02-28 00:00:00+07	1995.65.101	2	2019-08-30 00:00:00+07	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5336	\N	1	488	\N	\N	2019-05-22 00:00:00+07	0488.63.101	1	2019-12-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5335	\N	1	487	\N	\N	2019-05-22 00:00:00+07	0487.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5094	\N	1	2353	\N	\N	2019-03-13 00:00:00+07	2353.62.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5079	\N	1	2014	\N	\N	2019-03-11 00:00:00+07	2014.65.101	1	\N	\N	SA	\N	10112	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5330	\N	1	938	\N	\N	2019-05-22 00:00:00+07	0938.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5328	\N	1	2394	\N	\N	2019-05-21 00:00:00+07	2394.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5327	\N	1	768	\N	\N	2019-05-21 00:00:00+07	0768.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5326	\N	1	2393	\N	\N	2019-05-21 00:00:00+07	2393.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5325	\N	1	486	\N	\N	2019-05-21 00:00:00+07	0486.63.101	2	2020-02-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5323	\N	1	2392	\N	\N	2019-05-21 00:00:00+07	2392.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5322	\N	1	2391	\N	\N	2019-05-21 00:00:00+07	2391.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5321	\N	1	937	\N	\N	2019-05-21 00:00:00+07	0937.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5320	\N	1	936	\N	\N	2019-05-21 00:00:00+07	0936.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5113	\N	1	2032	\N	\N	2019-03-18 00:00:00+07	2032.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5318	\N	1	935	\N	\N	2019-05-21 00:00:00+07	0935.61.101	1	2020-03-01 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	f	f	2021-12-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5314	\N	1	481	\N	\N	2019-05-20 00:00:00+07	0481.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5124	\N	1	2040	\N	\N	2019-03-19 00:00:00+07	2040.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5125	\N	1	753	\N	\N	2019-03-19 00:00:00+07	0753.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5129	\N	1	2041	\N	\N	2019-03-20 00:00:00+07	2041.65.101	2	2020-07-11 00:00:00+07	\N	SA	\N	10112	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5130	\N	1	755	\N	\N	2019-03-20 00:00:00+07	0755.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5131	\N	1	465	\N	\N	2019-03-20 00:00:00+07	0465.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5132	\N	1	2357	\N	\N	2019-03-20 00:00:00+07	2357.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5309	\N	1	2389	\N	\N	2019-05-16 00:00:00+07	2389.62.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5306	\N	1	767	\N	\N	2019-05-16 00:00:00+07	0767.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5304	\N	1	478	\N	\N	2019-05-16 00:00:00+07	0478.63.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5149	\N	1	2049	\N	\N	2019-03-27 00:00:00+07	2049.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5150	\N	1	2050	\N	\N	2019-03-27 00:00:00+07	2050.65.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5151	\N	1	915	\N	\N	2019-03-27 00:00:00+07	0915.61.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5152	\N	1	2051	\N	\N	2019-03-27 00:00:00+07	2051.65.101	1	\N	\N	OA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5154	\N	1	756	\N	\N	2019-03-27 00:00:00+07	0756.64.101	1	\N	\N	SA	\N	10112	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6288	\N	1	868	\N	\N	2020-12-03 00:00:00+07	0868.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-03 00:00:00+07	\N	14:00	22:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2546	\N	2	414	\N	\N	2015-04-24 00:00:00+07	0414.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5437	\N	1	2171	\N	\N	2019-06-24 00:00:00+07	2171.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2527	\N	2	1791	\N	\N	2015-03-23 00:00:00+07	1791.62.103	1	\N	2020-06-01 00:00:00+07	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2512	\N	2	1815	\N	\N	2015-02-20 00:00:00+07	1815.65.103	2	2018-05-16 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2506	\N	2	1040	\N	\N	2015-02-12 00:00:00+07	1040.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2503	\N	2	1066	\N	\N	2015-02-06 00:00:00+07	1066.65.103	2	2018-04-01 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7422	\N	1	2892	\N	\N	2022-05-24 00:00:00+07	2892.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4427	\N	1	1697	\N	\N	2018-03-07 00:00:00+07	1697.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6133	\N	1	1025	\N	\N	2020-07-09 00:00:00+07	1025.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5209	\N	1	760	\N	\N	2019-04-11 00:00:00+07	0760.64.103	1	\N	2019-08-13 00:00:00+07	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6139	\N	1	1026	\N	\N	2020-07-14 00:00:00+07	1026.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7260	\N	1	2673	\N	\N	2022-02-14 00:00:00+07	2673.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-14 00:00:00+07	\N	08:30	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6140	\N	1	1027	\N	\N	2020-07-14 00:00:00+07	1027.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7258	\N	1	2672	\N	\N	2022-02-10 00:00:00+07	2672.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6143	\N	1	2410	\N	\N	2020-07-15 00:00:00+07	2410.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7261	\N	1	997	\N	\N	2022-02-16 00:00:00+07	0997.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6858	\N	1	923	\N	\N	2021-10-28 00:00:00+07	0923.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-28 00:00:00+07	\N	08:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6150	\N	1	2412	\N	\N	2020-07-22 00:00:00+07	2412.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6152	\N	1	2414	\N	\N	2020-07-29 00:00:00+07	2414.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6155	\N	1	2415	\N	\N	2020-07-29 00:00:00+07	2415.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6157	\N	1	2616	\N	\N	2020-08-03 00:00:00+07	2616.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7264	\N	1	1148	\N	\N	2022-02-16 00:00:00+07	1148.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5302	\N	1	2388	\N	\N	2019-05-16 00:00:00+07	2388.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5481	\N	1	783	\N	\N	2019-07-09 00:00:00+07	0783.64.103	2	2020-02-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6165	\N	1	860	\N	\N	2020-08-14 00:00:00+07	0860.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6172	\N	1	2620	\N	\N	2020-08-25 00:00:00+07	2620.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7269	\N	1	2676	\N	\N	2022-02-22 00:00:00+07	2676.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7270	\N	1	999	\N	\N	2022-02-22 00:00:00+07	0999.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5272	\N	1	2379	\N	\N	2019-05-08 00:00:00+07	2379.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6173	\N	1	2420	\N	\N	2020-08-25 00:00:00+07	2420.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5057	\N	1	2348	\N	\N	2019-03-01 00:00:00+07	2348.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7282	\N	1	2853	\N	\N	2022-02-24 00:00:00+07	2853.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5245	\N	1	934	\N	\N	2019-04-30 00:00:00+07	0934.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6861	\N	1	2575	\N	\N	2021-11-01 00:00:00+07	2575.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-01 00:00:00+07	\N	10:00	22:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7287	\N	1	2679	\N	\N	2022-03-02 00:00:00+07	2679.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-02 00:00:00+07	\N	10:00	20:00	5-10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7297	\N	1	692	\N	\N	2022-03-09 00:00:00+07	0692.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6863	\N	1	2577	\N	\N	2021-11-01 00:00:00+07	2577.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-01 00:00:00+07	\N	10:00	22:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7301	\N	1	2684	\N	\N	2022-03-10 00:00:00+07	2684.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5063	\N	1	2005	\N	\N	2019-03-06 00:00:00+07	2005.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1688	\N	2	1203	\N	\N	2011-02-08 00:00:00+07	1203.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7329	\N	1	2867	\N	\N	2022-03-16 00:00:00+07	2867.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-16 00:00:00+07	\N	10:00	21:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6180	\N	1	2424	\N	\N	2020-09-03 00:00:00+07	2424.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7331	\N	1	2868	\N	\N	2022-03-17 00:00:00+07	2868.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-17 00:00:00+07	\N	08:00	21:00	60	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5201	\N	1	2370	\N	\N	2019-04-10 00:00:00+07	2370.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7334	\N	1	2691	\N	\N	2022-03-17 00:00:00+07	2691.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7360	\N	1	1015	\N	\N	2022-03-30 00:00:00+07	1015.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5479	\N	1	2192	\N	\N	2019-07-04 00:00:00+07	2192.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5478	\N	1	2420	\N	\N	2019-07-04 00:00:00+07	2420.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4318	\N	1	2213	\N	\N	2017-12-18 00:00:00+07	2213.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5477	\N	1	958	\N	\N	2019-07-03 00:00:00+07	0958.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4320	\N	1	609	\N	\N	2017-12-19 00:00:00+07	0609.64.103	2	2018-07-01 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5476	\N	1	957	\N	\N	2019-07-03 00:00:00+07	0957.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5070	\N	1	2351	\N	\N	2019-03-08 00:00:00+07	2351.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5071	\N	1	2008	\N	\N	2019-03-08 00:00:00+07	2008.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5072	\N	1	2352	\N	\N	2019-03-08 00:00:00+07	2352.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6866	\N	1	2578	\N	\N	2021-11-02 00:00:00+07	2578.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6868	\N	1	927	\N	\N	2021-11-03 00:00:00+07	0927.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5469	\N	1	2191	\N	\N	2019-07-02 00:00:00+07	2191.65.103	2	2020-01-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6188	\N	1	1031	\N	\N	2020-09-09 00:00:00+07	1031.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	2020-09-09 00:00:00+07	\N	08:00	22:00	10-20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5073	\N	1	2009	\N	\N	2019-03-11 00:00:00+07	2009.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7364	\N	1	2700	\N	\N	2022-04-04 00:00:00+07	2700.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5439	\N	1	2173	\N	\N	2019-06-25 00:00:00+07	2173.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-06 00:00:00+07	\N	\N	\N	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7365	\N	1	1162	\N	\N	2022-04-04 00:00:00+07	1162.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	07:00	22:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7371	\N	1	2877	\N	\N	2022-04-06 00:00:00+07	2877.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-04-06 00:00:00+07	\N	10:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7373	\N	1	2879	\N	\N	2022-04-06 00:00:00+07	2879.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-04-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5202	\N	1	2371	\N	\N	2019-04-10 00:00:00+07	2371.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5238	\N	1	931	\N	\N	2019-04-26 00:00:00+07	0931.61.103	2	2021-01-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-06-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7374	\N	1	2880	\N	\N	2022-04-06 00:00:00+07	2880.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-04-06 00:00:00+07	\N	09:00	21:00	6	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7383	\N	1	2704	\N	\N	2022-04-11 00:00:00+07	2704.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-04-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4338	\N	1	2217	\N	\N	2018-01-10 00:00:00+07	2217.62.103	2	2020-02-17 00:00:00+07	2018-06-18 00:00:00+07	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2961	\N	1	1986	\N	\N	2016-01-01 00:00:00+07	1986.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2960	\N	1	1532	\N	\N	2017-01-01 00:00:00+07	1532.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2949	\N	1	1552	\N	\N	2016-12-01 00:00:00+07	1552.65.103	2	2017-12-01 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6224	\N	1	2432	\N	\N	2020-09-23 00:00:00+07	2432.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-09-23 00:00:00+07	\N	15:00	22:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7384	\N	1	2882	\N	\N	2022-04-13 00:00:00+07	2882.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-04-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7396	\N	1	1167	\N	\N	2022-04-21 00:00:00+07	1167.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-04-21 00:00:00+07	\N	13:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6875	\N	1	929	\N	\N	2021-11-04 00:00:00+07	0929.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6226	\N	1	2631	\N	\N	2020-09-23 00:00:00+07	2631.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-09-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5453	\N	1	2184	\N	\N	2019-06-26 00:00:00+07	2184.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7415	\N	1	2889	\N	\N	2022-05-20 00:00:00+07	2889.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-05-20 00:00:00+07	\N	07:00	19:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5452	\N	1	2183	\N	\N	2019-06-26 00:00:00+07	2183.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2863	\N	2	304	\N	\N	2016-09-09 00:00:00+07	0304.63.103	2	2019-06-01 00:00:00+07	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2859	\N	2	2003	\N	\N	2016-09-09 00:00:00+07	2003.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5451	\N	1	2182	\N	\N	2019-06-26 00:00:00+07	2182.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4891	\N	1	2308	\N	\N	2018-12-06 00:00:00+07	2308.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4892	\N	1	1915	\N	\N	2018-12-06 00:00:00+07	1915.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2858	\N	2	2006	\N	\N	2016-09-09 00:00:00+07	2006.62.103	2	2018-09-16 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5450	\N	1	2181	\N	\N	2019-06-26 00:00:00+07	2181.65.103	2	2020-04-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-07-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4893	\N	1	1916	\N	\N	2018-12-06 00:00:00+07	1916.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5448	\N	1	2180	\N	\N	2019-06-26 00:00:00+07	2180.65.103	2	2021-11-02 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2853	\N	2	1496	\N	\N	2016-09-08 00:00:00+07	1496.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2848	\N	2	1999	\N	\N	2016-09-01 00:00:00+07	1999.61.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4894	\N	1	1917	\N	\N	2018-12-06 00:00:00+07	1917.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7419	\N	1	2721	\N	\N	2022-05-24 00:00:00+07	2721.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2835	\N	2	476	\N	\N	2016-06-23 00:00:00+07	0476.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2832	\N	2	1473	\N	\N	2016-06-09 00:00:00+07	1473.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2831	\N	2	1484	\N	\N	2016-06-06 00:00:00+07	1484.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2829	\N	2	1437	\N	\N	2016-05-29 00:00:00+07	1437.65.103	2	2018-02-28 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2821	\N	2	1433	\N	\N	2016-05-04 00:00:00+07	1433.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2820	\N	2	1968	\N	\N	2016-05-04 00:00:00+07	1968.62.103	2	2022-05-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2808	\N	2	769	\N	\N	2016-04-08 00:00:00+07	0769.61.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4896	\N	1	1918	\N	\N	2018-12-06 00:00:00+07	1918.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2804	\N	2	1583	\N	\N	2016-03-28 00:00:00+07	1583.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2802	\N	2	1963	\N	\N	2016-03-11 00:00:00+07	1963.62.103	2	2017-05-22 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6228	\N	1	864	\N	\N	2020-09-29 00:00:00+07	0864.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-09-29 00:00:00+07	\N	12:00	23:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2789	\N	2	1944	\N	\N	2016-02-29 00:00:00+07	1944.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2775	\N	2	467	\N	\N	2016-02-09 00:00:00+07	0467.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5447	\N	1	2179	\N	\N	2019-06-26 00:00:00+07	2179.65.103	2	2021-11-02 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2743	\N	2	1840	\N	\N	2015-11-13 00:00:00+07	1840.62.103	2	2016-12-01 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2731	\N	2	757	\N	\N	2015-11-05 00:00:00+07	0757.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2721	\N	2	1336	\N	\N	2015-10-09 00:00:00+07	1336.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2720	\N	2	1329	\N	\N	2015-10-07 00:00:00+07	1329.65.103	1	\N	2020-06-01 00:00:00+07	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2700	\N	2	1901	\N	\N	2015-09-02 00:00:00+07	1901.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4897	\N	1	1919	\N	\N	2018-12-06 00:00:00+07	1919.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4394	\N	1	2231	\N	\N	2018-02-07 00:00:00+07	2231.62.103	2	2020-04-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4898	\N	1	732	\N	\N	2018-12-06 00:00:00+07	0732.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2697	\N	2	287	\N	\N	2015-08-26 00:00:00+07	0287.63.103	2	2016-10-23 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2680	\N	2	1690	\N	\N	2015-07-08 00:00:00+07	1690.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4899	\N	1	733	\N	\N	2018-12-06 00:00:00+07	0733.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4900	\N	1	2310	\N	\N	2018-12-06 00:00:00+07	2310.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5446	\N	1	2178	\N	\N	2019-06-26 00:00:00+07	2178.65.103	2	2021-11-02 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2662	\N	2	471	\N	\N	2016-06-01 00:00:00+07	0471.64.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5445	\N	1	2177	\N	\N	2019-06-26 00:00:00+07	2177.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2638	\N	2	262	\N	\N	2015-06-10 00:00:00+07	0262.63.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2636	\N	2	1210	\N	\N	2015-06-09 00:00:00+07	1210.65.103	1	2017-01-30 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2625	\N	2	1877	\N	\N	2015-05-27 00:00:00+07	1877.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5444	\N	1	2176	\N	\N	2019-06-26 00:00:00+07	2176.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4406	\N	1	1693	\N	\N	2018-02-20 00:00:00+07	1693.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5443	\N	1	2175	\N	\N	2019-06-26 00:00:00+07	2175.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5442	\N	1	2174	\N	\N	2019-06-26 00:00:00+07	2174.65.103	2	2021-11-02 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4412	\N	1	402	\N	\N	2018-02-26 00:00:00+07	0402.63.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2587	\N	2	1209	\N	\N	2015-05-11 00:00:00+07	1209.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6247	\N	1	2441	\N	\N	2020-10-15 00:00:00+07	2441.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-10-15 00:00:00+07	\N	11:00	23:00	35 - 50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5135	\N	1	2044	\N	\N	2019-03-21 00:00:00+07	2044.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7420	\N	1	2891	\N	\N	2022-05-24 00:00:00+07	2891.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4417	\N	1	1695	\N	\N	2018-03-01 00:00:00+07	1695.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5310	\N	1	2390	\N	\N	2019-05-16 00:00:00+07	2390.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7432	\N	1	2724	\N	\N	2022-05-30 00:00:00+07	2724.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-05-30 00:00:00+07	\N	09:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5093	\N	1	2019	\N	\N	2019-03-13 00:00:00+07	2019.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7427	\N	1	2896	\N	\N	2022-05-25 00:00:00+07	2896.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-05-25 00:00:00+07	\N	10:00	21:30	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7436	\N	1	709	\N	\N	2022-06-02 00:00:00+07	0709.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4436	\N	1	1702	\N	\N	2018-03-12 00:00:00+07	1702.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2497	\N	2	1583	\N	\N	2015-01-22 00:00:00+07	1583.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5235	\N	1	2087	\N	\N	2019-04-25 00:00:00+07	2087.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2496	\N	2	1814	\N	\N	2015-01-14 00:00:00+07	1814.62.103	2	2018-05-16 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2483	\N	2	1803	\N	\N	2014-12-09 00:00:00+07	1803.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4441	\N	1	1703	\N	\N	2018-03-14 00:00:00+07	1703.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4442	\N	1	2239	\N	\N	2018-03-14 00:00:00+07	2239.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2472	\N	2	245	\N	\N	2014-10-29 00:00:00+07	0245.63.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5382	\N	1	2139	\N	\N	2019-06-11 00:00:00+07	2139.65.103	1	\N	2019-10-01 00:00:00+07	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4446	\N	1	2242	\N	\N	2018-03-14 00:00:00+07	2242.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2462	\N	2	1735	\N	\N	2014-09-25 00:00:00+07	1735.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4448	\N	1	2244	\N	\N	2018-03-14 00:00:00+07	2244.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5103	\N	1	2028	\N	\N	2019-03-14 00:00:00+07	2028.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6234	\N	1	2437	\N	\N	2020-10-07 00:00:00+07	2437.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-10-07 00:00:00+07	\N	09:00	22:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6880	\N	1	2583	\N	\N	2021-11-05 00:00:00+07	2583.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	15:00	23:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6881	\N	1	2769	\N	\N	2021-11-08 00:00:00+07	2769.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7452	\N	1	2729	\N	\N	2022-06-09 00:00:00+07	2729.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4458	\N	1	2247	\N	\N	2018-03-22 00:00:00+07	2247.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2439	\N	2	1760	\N	\N	2014-05-26 00:00:00+07	1760.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4459	\N	1	641	\N	\N	2018-03-22 00:00:00+07	0641.64.103	1	2018-04-30 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2424	\N	2	1718	\N	\N	2014-02-28 00:00:00+07	1718.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5137	\N	1	2045	\N	\N	2019-03-25 00:00:00+07	2045.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4922	\N	1	895	\N	\N	2019-01-04 00:00:00+07	0895.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2421	\N	2	1694	\N	\N	2014-02-17 00:00:00+07	1694.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7453	\N	1	1023	\N	\N	2022-06-09 00:00:00+07	1023.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5162	\N	1	2053	\N	\N	2019-03-28 00:00:00+07	2053.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6884	\N	1	647	\N	\N	2021-11-09 00:00:00+07	0647.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2415	\N	2	1709	\N	\N	2014-01-07 00:00:00+07	1709.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6235	\N	1	2438	\N	\N	2020-10-07 00:00:00+07	2438.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-10-07 00:00:00+07	\N	09:00	22:30	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7458	\N	1	2732	\N	\N	2022-06-13 00:00:00+07	2732.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4467	\N	1	648	\N	\N	2018-03-22 00:00:00+07	0648.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6236	\N	1	2439	\N	\N	2020-10-07 00:00:00+07	2439.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-10-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2384	\N	2	1676	\N	\N	2013-04-26 00:00:00+07	1676.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2383	\N	2	1673	\N	\N	2013-04-22 00:00:00+07	1673.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5324	\N	1	485	\N	\N	2019-05-21 00:00:00+07	0485.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5427	\N	1	2167	\N	\N	2019-06-20 00:00:00+07	2167.65.103	2	2020-12-26 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5288	\N	1	2100	\N	\N	2019-05-14 00:00:00+07	2100.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4480	\N	1	408	\N	\N	2018-03-29 00:00:00+07	0408.63.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4482	\N	1	873	\N	\N	2018-04-02 00:00:00+07	0873.61.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5205	\N	1	2072	\N	\N	2019-04-10 00:00:00+07	2072.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5307	\N	1	2111	\N	\N	2019-05-16 00:00:00+07	2111.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6887	\N	1	2586	\N	\N	2021-11-09 00:00:00+07	2586.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	13:00	23:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6890	\N	1	2587	\N	\N	2021-11-09 00:00:00+07	2587.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	08:00	00:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5426	\N	1	2409	\N	\N	2019-06-20 00:00:00+07	2409.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5425	\N	1	2166	\N	\N	2019-06-19 00:00:00+07	2166.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2375	\N	2	636	\N	\N	2013-02-22 00:00:00+07	0636.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2374	\N	2	1652	\N	\N	2013-02-18 00:00:00+07	1652.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2351	\N	2	364	\N	\N	2012-07-03 00:00:00+07	0364.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2334	\N	2	1640	\N	\N	2012-04-23 00:00:00+07	1640.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5424	\N	1	2165	\N	\N	2019-06-19 00:00:00+07	2165.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4499	\N	1	2249	\N	\N	2018-04-12 00:00:00+07	2249.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2317	\N	2	1270	\N	\N	2012-01-25 00:00:00+07	1270.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2315	\N	2	355	\N	\N	2011-12-19 00:00:00+07	0355.64.103	1	\N	2019-06-09 00:00:00+07	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2311	\N	2	605	\N	\N	2011-11-21 00:00:00+07	0605.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6891	\N	1	2588	\N	\N	2021-11-09 00:00:00+07	2588.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	13:00	00:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2284	\N	2	1592	\N	\N	2011-08-01 00:00:00+07	1592.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6892	\N	1	2589	\N	\N	2021-11-09 00:00:00+07	2589.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	11:00	23:00	55	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2283	\N	2	598	\N	\N	2011-08-12 00:00:00+07	0598.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2259	\N	2	1583	\N	\N	2011-07-01 00:00:00+07	1583.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2217	\N	2	334	\N	\N	2011-03-09 00:00:00+07	0334.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2209	\N	2	1544	\N	\N	2011-02-07 00:00:00+07	1544.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5422	\N	1	2164	\N	\N	2019-06-19 00:00:00+07	2164.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2204	\N	2	1540	\N	\N	2011-04-01 00:00:00+07	1540.62.103	2	2017-06-15 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4512	\N	1	2253	\N	\N	2018-04-16 00:00:00+07	2253.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6899	\N	1	2770	\N	\N	2021-11-10 00:00:00+07	2770.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4936	\N	1	2321	\N	\N	2019-01-28 00:00:00+07	2321.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7456	\N	1	1026	\N	\N	2022-06-10 00:00:00+07	1026.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5370	\N	1	772	\N	\N	2019-06-10 00:00:00+07	0772.64.103	2	2020-02-07 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4517	\N	1	1722	\N	\N	2018-04-17 00:00:00+07	1722.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5301	\N	1	766	\N	\N	2019-05-16 00:00:00+07	0766.64.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6240	\N	1	2440	\N	\N	2020-10-14 00:00:00+07	2440.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-10-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6900	\N	1	2592	\N	\N	2021-11-11 00:00:00+07	2592.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5305	\N	1	2110	\N	\N	2019-05-16 00:00:00+07	2110.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4529	\N	1	2256	\N	\N	2018-04-20 00:00:00+07	2256.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1946	\N	2	1125	\N	\N	2011-02-08 00:00:00+07	1125.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7476	\N	1	715	\N	\N	2022-06-16 00:00:00+07	0715.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4531	\N	1	1732	\N	\N	2018-04-24 00:00:00+07	1732.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4524	\N	1	663	\N	\N	2018-04-19 00:00:00+07	0663.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6901	\N	1	2771	\N	\N	2021-11-12 00:00:00+07	2771.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-12 00:00:00+07	\N	12:00	20:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6910	\N	1	2594	\N	\N	2021-11-17 00:00:00+07	2594.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	09:00	22:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1750	\N	2	351	\N	\N	2011-02-07 00:00:00+07	0351.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1749	\N	2	962	\N	\N	2011-02-07 00:00:00+07	0962.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6914	\N	1	2772	\N	\N	2021-11-18 00:00:00+07	2772.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6920	\N	1	2597	\N	\N	2021-11-22 00:00:00+07	2597.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-22 00:00:00+07	\N	10:00	23:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4995	\N	1	1956	\N	\N	2019-02-20 00:00:00+07	1956.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1748	\N	2	1267	\N	\N	2011-02-08 00:00:00+07	1267.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6921	\N	1	2598	\N	\N	2021-11-22 00:00:00+07	2598.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-22 00:00:00+07	\N	13:00	23:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6922	\N	1	2599	\N	\N	2021-11-22 00:00:00+07	2599.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-22 00:00:00+07	\N	10:00	00:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1717	\N	2	785	\N	\N	2011-02-08 00:00:00+07	0785.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1716	\N	2	13	\N	\N	2011-02-08 00:00:00+07	0013.64.103	2	2019-02-11 00:00:00+07	2018-07-18 00:00:00+07	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1714	\N	2	18	\N	\N	2011-02-08 00:00:00+07	0018.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1711	\N	2	942	\N	\N	2011-02-08 00:00:00+07	0942.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7478	\N	1	2907	\N	\N	2022-06-16 00:00:00+07	2907.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6923	\N	1	2600	\N	\N	2021-11-22 00:00:00+07	2600.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-22 00:00:00+07	\N	08:00	00:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1709	\N	2	311	\N	\N	2011-02-08 00:00:00+07	0311.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1706	\N	2	248	\N	\N	2011-02-08 00:00:00+07	0248.62.103	2	2017-09-01 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1704	\N	2	184	\N	\N	2011-02-08 00:00:00+07	0184.62.103	2	2018-05-16 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1702	\N	2	49	\N	\N	2011-02-08 00:00:00+07	0049.61.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1701	\N	2	83	\N	\N	2011-02-08 00:00:00+07	0083.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1696	\N	2	767	\N	\N	2011-02-08 00:00:00+07	0767.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1695	\N	2	479	\N	\N	2011-02-08 00:00:00+07	0479.61.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4554	\N	1	1739	\N	\N	2018-05-07 00:00:00+07	1739.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1693	\N	2	1275	\N	\N	2011-02-08 00:00:00+07	1275.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1692	\N	2	454	\N	\N	2011-02-08 00:00:00+07	0454.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1691	\N	2	784	\N	\N	2011-02-08 00:00:00+07	0784.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1690	\N	2	261	\N	\N	2011-02-08 00:00:00+07	0261.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1689	\N	2	452	\N	\N	2011-02-08 00:00:00+07	0452.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6924	\N	1	2601	\N	\N	2021-11-22 00:00:00+07	2601.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-22 00:00:00+07	\N	17:00	00:00	7	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6929	\N	1	934	\N	\N	2021-11-23 00:00:00+07	0934.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7479	\N	1	2908	\N	\N	2022-06-16 00:00:00+07	2908.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6930	\N	1	935	\N	\N	2021-11-23 00:00:00+07	0935.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4565	\N	1	2261	\N	\N	2018-05-15 00:00:00+07	2261.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6931	\N	1	936	\N	\N	2021-11-23 00:00:00+07	0936.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6934	\N	1	2605	\N	\N	2021-11-24 00:00:00+07	2605.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-24 00:00:00+07	\N	10:00	22:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6942	\N	1	2611	\N	\N	2021-11-25 00:00:00+07	2611.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5025	\N	1	2341	\N	\N	2019-02-26 00:00:00+07	2341.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7487	\N	1	2912	\N	\N	2022-06-23 00:00:00+07	2912.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-23 00:00:00+07	\N	09:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6256	\N	1	2447	\N	\N	2020-11-02 00:00:00+07	2447.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6265	\N	1	2638	\N	\N	2020-11-10 00:00:00+07	2638.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-10 00:00:00+07	\N	10:00	23:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4582	\N	1	1750	\N	\N	2018-06-05 00:00:00+07	1750.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5289	\N	1	2101	\N	\N	2019-05-14 00:00:00+07	2101.65.103	2	2021-01-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7482	\N	1	2737	\N	\N	2022-06-20 00:00:00+07	2737.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5276	\N	1	2381	\N	\N	2019-05-08 00:00:00+07	2381.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4586	\N	1	422	\N	\N	2018-06-22 00:00:00+07	0422.63.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5363	\N	1	2129	\N	\N	2019-05-28 00:00:00+07	2129.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6272	\N	1	2448	\N	\N	2020-11-19 00:00:00+07	2448.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6274	\N	1	2449	\N	\N	2020-11-20 00:00:00+07	2449.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-20 00:00:00+07	\N	10:00	21:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4590	\N	1	2264	\N	\N	2018-07-05 00:00:00+07	2264.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6275	\N	1	866	\N	\N	2020-11-20 00:00:00+07	0866.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-20 00:00:00+07	\N	10:00	21:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4592	\N	1	2265	\N	\N	2018-07-06 00:00:00+07	2265.62.103	2	2022-03-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5317	\N	1	484	\N	\N	2019-05-20 00:00:00+07	0484.63.103	1	\N	2020-06-01 00:00:00+07	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5361	\N	1	2128	\N	\N	2019-05-27 00:00:00+07	2128.65.103	2	2020-02-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6282	\N	1	2452	\N	\N	2020-11-26 00:00:00+07	2452.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-26 00:00:00+07	\N	10:00	20:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6284	\N	1	867	\N	\N	2020-11-27 00:00:00+07	0867.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-27 00:00:00+07	\N	10:00	20:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6285	\N	2	2453	\N	\N	2020-11-30 00:00:00+07	2453.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-30 00:00:00+07	\N	10:00	22:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6286	\N	1	2454	\N	\N	2020-12-02 00:00:00+07	2454.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-02 00:00:00+07	\N	15:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3444	\N	1	948	\N	\N	2017-03-01 00:00:00+07	0948.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6287	\N	1	2646	\N	\N	2020-12-03 00:00:00+07	2646.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-03 00:00:00+07	\N	10:30	20:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4600	\N	1	1760	\N	\N	2018-07-11 00:00:00+07	1760.65.103	2	2018-11-01 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5316	\N	1	483	\N	\N	2019-05-20 00:00:00+07	0483.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5360	\N	1	2403	\N	\N	2019-05-27 00:00:00+07	2403.62.103	2	2020-02-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3450	\N	1	2025	\N	\N	2016-08-01 00:00:00+07	2025.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6938	\N	1	2607	\N	\N	2021-11-25 00:00:00+07	2607.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6939	\N	1	2608	\N	\N	2021-11-25 00:00:00+07	2608.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-25 00:00:00+07	\N	10:00	12:00	40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5358	\N	1	2402	\N	\N	2019-05-27 00:00:00+07	2402.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6289	\N	2	2647	\N	\N	2020-12-03 00:00:00+07	2647.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-03 00:00:00+07	\N	10:00	22:00	300 - 500	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5357	\N	1	2401	\N	\N	2019-05-27 00:00:00+07	2401.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6940	\N	1	2609	\N	\N	2021-11-25 00:00:00+07	2609.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-25 00:00:00+07	\N	07:00	23:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6308	\N	1	870	\N	\N	2020-12-07 00:00:00+07	0870.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-07 00:00:00+07	\N	10:00	22:00	7	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4611	\N	1	1765	\N	\N	2018-07-16 00:00:00+07	1765.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3429	\N	1	2021	\N	\N	2017-01-02 00:00:00+07	2021.62.103	1	2018-04-03 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5695	\N	1	2491	\N	\N	2019-09-18 00:00:00+07	2491.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6963	\N	1	2778	\N	\N	2021-12-06 00:00:00+07	2778.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5315	\N	1	482	\N	\N	2019-05-20 00:00:00+07	0482.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6967	\N	1	2615	\N	\N	2021-12-07 00:00:00+07	2615.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-07 00:00:00+07	\N	10:00	21:00	45	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6968	\N	1	2616	\N	\N	2021-12-07 00:00:00+07	2616.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-07 00:00:00+07	\N	10:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5033	\N	1	2344	\N	\N	2019-02-27 00:00:00+07	2344.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3564	\N	1	2130	\N	\N	2017-03-01 00:00:00+07	2130.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6309	\N	1	595	\N	\N	2020-12-07 00:00:00+07	0595.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-07 00:00:00+07	\N	09:25	22:45	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6310	\N	1	2456	\N	\N	2020-12-07 00:00:00+07	2456.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-07 00:00:00+07	\N	09:00	22:45	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3402	\N	1	1534	\N	\N	2017-01-02 00:00:00+07	1534.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6312	\N	1	871	\N	\N	2020-12-07 00:00:00+07	0871.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-07 00:00:00+07	\N	09:00	22:45	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6950	\N	1	660	\N	\N	2021-12-01 00:00:00+07	0660.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5195	\N	1	2069	\N	\N	2019-04-08 00:00:00+07	2069.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6952	\N	1	940	\N	\N	2021-12-02 00:00:00+07	0940.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-02 00:00:00+07	\N	11:00	23:00	40	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4624	\N	1	681	\N	\N	2018-07-25 00:00:00+07	0681.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3394	\N	1	1601	\N	\N	2017-03-30 00:00:00+07	1601.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5700	\N	1	520	\N	\N	2019-09-19 00:00:00+07	0520.63.103	1	\N	2020-01-01 00:00:00+07	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3372	\N	1	515	\N	\N	2016-12-02 00:00:00+07	0515.64.103	2	2019-06-19 00:00:00+07	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5352	\N	1	2400	\N	\N	2019-05-24 00:00:00+07	2400.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5351	\N	1	2124	\N	\N	2019-05-24 00:00:00+07	2124.65.103	2	2020-02-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6313	\N	1	2457	\N	\N	2020-12-07 00:00:00+07	2457.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-07 00:00:00+07	\N	\N	\N	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6315	\N	1	2648	\N	\N	2020-12-08 00:00:00+07	2648.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-08 00:00:00+07	\N	07:00	22:00	80	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5705	\N	1	982	\N	\N	2019-09-20 00:00:00+07	0982.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6973	\N	1	2618	\N	\N	2021-12-08 00:00:00+07	2618.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-08 00:00:00+07	\N	15:00	23:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6980	\N	1	2619	\N	\N	2021-12-10 00:00:00+07	2619.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-10 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4631	\N	1	2269	\N	\N	2018-08-01 00:00:00+07	2269.62.103	2	2020-02-20 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6318	\N	1	2459	\N	\N	2020-12-10 00:00:00+07	2459.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	07:00	21:00	60	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6981	\N	1	2620	\N	\N	2021-12-10 00:00:00+07	2620.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-10 00:00:00+07	\N	09:00	00:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5709	\N	1	2271	\N	\N	2019-09-23 00:00:00+07	2271.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6325	\N	1	2462	\N	\N	2020-12-10 00:00:00+07	2462.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	10:00	23:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6326	\N	1	2463	\N	\N	2020-12-10 00:00:00+07	2463.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	12:00	23:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6329	\N	1	2651	\N	\N	2020-12-17 00:00:00+07	2651.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5176	\N	1	2061	\N	\N	2019-04-04 00:00:00+07	2061.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5161	\N	1	917	\N	\N	2019-03-28 00:00:00+07	0917.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6982	\N	1	2621	\N	\N	2021-12-13 00:00:00+07	2621.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6987	\N	1	2784	\N	\N	2021-12-14 00:00:00+07	2784.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-14 00:00:00+07	\N	11:00	20:00	150	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3349	\N	1	2037	\N	\N	2017-01-02 00:00:00+07	2037.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6331	\N	1	1039	\N	\N	2020-12-23 00:00:00+07	1039.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-12-23 00:00:00+07	\N	\N	\N	120	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6999	\N	1	2627	\N	\N	2021-12-17 00:00:00+07	2627.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-17 00:00:00+07	\N	15:00	23:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6992	\N	1	2623	\N	\N	2021-12-16 00:00:00+07	2623.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4647	\N	1	2274	\N	\N	2018-08-08 00:00:00+07	2274.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6997	\N	1	2625	\N	\N	2021-12-17 00:00:00+07	2625.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3652	\N	1	1592	\N	\N	2017-03-17 00:00:00+07	1592.65.103	2	2018-05-31 00:00:00+07	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4649	\N	1	685	\N	\N	2018-08-08 00:00:00+07	0685.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6998	\N	1	2626	\N	\N	2021-12-17 00:00:00+07	2626.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5347	\N	1	2399	\N	\N	2019-05-24 00:00:00+07	2399.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5710	\N	1	2273	\N	\N	2019-09-23 00:00:00+07	2273.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5711	\N	1	2272	\N	\N	2019-09-23 00:00:00+07	2272.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5766	\N	1	2310	\N	\N	2019-10-03 00:00:00+07	2310.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5712	\N	1	2274	\N	\N	2019-09-23 00:00:00+07	2274.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6337	\N	1	2652	\N	\N	2021-01-08 00:00:00+07	2652.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-08 00:00:00+07	\N	11:00	19:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5715	\N	1	2275	\N	\N	2019-09-23 00:00:00+07	2275.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5716	\N	1	2276	\N	\N	2019-09-23 00:00:00+07	2276.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5717	\N	1	2277	\N	\N	2019-09-23 00:00:00+07	2277.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3661	\N	1	2045	\N	\N	2017-02-22 00:00:00+07	2045.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4698	\N	1	2283	\N	\N	2018-09-06 00:00:00+07	2283.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5274	\N	1	2380	\N	\N	2019-05-08 00:00:00+07	2380.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3732	\N	1	2149	\N	\N	2017-05-23 00:00:00+07	2149.62.103	2	2018-01-20 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7000	\N	1	944	\N	\N	2021-12-17 00:00:00+07	0944.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6339	\N	1	1041	\N	\N	2021-01-11 00:00:00+07	1041.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-11 00:00:00+07	\N	07:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6340	\N	1	873	\N	\N	2021-01-11 00:00:00+07	0873.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-11 00:00:00+07	\N	07:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5655	\N	1	800	\N	\N	2019-09-03 00:00:00+07	0800.64.103	2	2020-01-30 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5043	\N	1	2347	\N	\N	2019-02-28 00:00:00+07	2347.62.103	2	2021-07-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6345	\N	1	2468	\N	\N	2021-01-13 00:00:00+07	2468.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-13 00:00:00+07	\N	09:00	23:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6347	\N	1	2469	\N	\N	2021-01-14 00:00:00+07	2469.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-14 00:00:00+07	\N	08:00	20:00	8	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7006	\N	1	2629	\N	\N	2021-12-21 00:00:00+07	2629.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-12-21 00:00:00+07	\N	08:00	21:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4661	\N	1	689	\N	\N	2018-08-13 00:00:00+07	0689.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6348	\N	1	599	\N	\N	2021-01-14 00:00:00+07	0599.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-14 00:00:00+07	\N	08:00	20:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3675	\N	1	2095	\N	\N	2017-04-06 00:00:00+07	2095.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3677	\N	1	2100	\N	\N	2017-04-06 00:00:00+07	2100.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6349	\N	1	874	\N	\N	2021-01-14 00:00:00+07	0874.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-14 00:00:00+07	\N	08:00	20:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6352	\N	1	2658	\N	\N	2021-01-20 00:00:00+07	2658.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-20 00:00:00+07	\N	10:00	21:30	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5718	\N	1	2278	\N	\N	2019-09-23 00:00:00+07	2278.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6353	\N	1	2470	\N	\N	2021-01-20 00:00:00+07	2470.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-20 00:00:00+07	\N	13:30	20:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5719	\N	1	2279	\N	\N	2019-09-23 00:00:00+07	2279.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5652	\N	1	2481	\N	\N	2019-08-29 00:00:00+07	2481.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5651	\N	1	517	\N	\N	2019-08-28 00:00:00+07	0517.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7021	\N	1	954	\N	\N	2022-01-03 00:00:00+07	0954.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7026	\N	1	2792	\N	\N	2022-01-03 00:00:00+07	2792.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5313	\N	1	480	\N	\N	2019-05-20 00:00:00+07	0480.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7069	\N	1	2637	\N	\N	2022-01-07 00:00:00+07	2637.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6355	\N	1	1044	\N	\N	2021-01-21 00:00:00+07	1044.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-21 00:00:00+07	\N	10:00	21:00	60	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5647	\N	1	2243	\N	\N	2019-08-28 00:00:00+07	2243.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5720	\N	1	2280	\N	\N	2019-09-23 00:00:00+07	2280.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5721	\N	1	2281	\N	\N	2019-09-23 00:00:00+07	2281.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5722	\N	1	2282	\N	\N	2019-09-23 00:00:00+07	2282.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7082	\N	1	2641	\N	\N	2022-01-12 00:00:00+07	2641.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7133	\N	1	2651	\N	\N	2022-01-18 00:00:00+07	2651.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7090	\N	1	2801	\N	\N	2022-01-13 00:00:00+07	2801.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7113	\N	1	969	\N	\N	2022-01-14 00:00:00+07	0969.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6356	\N	1	600	\N	\N	2021-01-21 00:00:00+07	0600.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-21 00:00:00+07	\N	10:00	21:00	60	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6357	\N	1	2471	\N	\N	2021-01-21 00:00:00+07	2471.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-21 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5723	\N	1	2283	\N	\N	2019-09-23 00:00:00+07	2283.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6358	\N	1	875	\N	\N	2021-01-21 00:00:00+07	0875.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-21 00:00:00+07	\N	10:00	22:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6362	\N	1	2661	\N	\N	2021-01-28 00:00:00+07	2661.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-01-28 00:00:00+07	\N	09:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6368	\N	1	2665	\N	\N	2021-02-01 00:00:00+07	2665.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-02-01 00:00:00+07	\N	09:00	20:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5724	\N	1	2284	\N	\N	2019-09-23 00:00:00+07	2284.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4680	\N	1	1796	\N	\N	2018-08-29 00:00:00+07	1796.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6380	\N	1	603	\N	\N	2021-02-10 00:00:00+07	0603.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-02-10 00:00:00+07	\N	11:00	21:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6388	\N	1	1048	\N	\N	2021-03-02 00:00:00+07	1048.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-02 00:00:00+07	\N	07:00	22:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6394	\N	1	2676	\N	\N	2021-03-05 00:00:00+07	2676.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-05 00:00:00+07	8	11:00	20:30	22	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3720	\N	1	2139	\N	\N	2017-05-10 00:00:00+07	2139.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6397	\N	1	2677	\N	\N	2021-03-08 00:00:00+07	2677.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6398	\N	1	882	\N	\N	2021-03-08 00:00:00+07	0882.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-08 00:00:00+07	\N	10:00	22:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6400	\N	1	884	\N	\N	2021-03-10 00:00:00+07	0884.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-10 00:00:00+07	\N	12:00	\N	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7131	\N	1	2650	\N	\N	2022-01-18 00:00:00+07	2650.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7143	\N	1	974	\N	\N	2022-01-20 00:00:00+07	0974.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6404	\N	1	2680	\N	\N	2021-03-18 00:00:00+07	2680.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-18 00:00:00+07	\N	07:00	22:00	85	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5725	\N	1	2285	\N	\N	2019-09-23 00:00:00+07	2285.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5639	\N	1	2478	\N	\N	2019-08-21 00:00:00+07	2478.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6410	\N	1	2480	\N	\N	2021-03-24 00:00:00+07	2480.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-24 00:00:00+07	\N	12:00	20:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6412	\N	1	2482	\N	\N	2021-03-25 00:00:00+07	2482.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-25 00:00:00+07	\N	09:00	23:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5726	\N	1	2286	\N	\N	2019-09-23 00:00:00+07	2286.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6414	\N	1	605	\N	\N	2021-03-29 00:00:00+07	0605.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-29 00:00:00+07	\N	09:00	\N	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7145	\N	1	2813	\N	\N	2022-01-20 00:00:00+07	2813.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-20 00:00:00+07	\N	10:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5635	\N	1	2477	\N	\N	2019-08-20 00:00:00+07	2477.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4694	\N	1	887	\N	\N	2018-09-05 00:00:00+07	0887.61.103	2	2020-04-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5727	\N	1	2287	\N	\N	2019-09-23 00:00:00+07	2287.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5634	\N	1	975	\N	\N	2019-08-19 00:00:00+07	0975.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4695	\N	1	1806	\N	\N	2018-09-05 00:00:00+07	1806.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4696	\N	1	1807	\N	\N	2018-09-05 00:00:00+07	1807.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4697	\N	1	2282	\N	\N	2018-09-06 00:00:00+07	2282.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6424	\N	1	2487	\N	\N	2021-04-08 00:00:00+07	2487.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-04-08 00:00:00+07	\N	12:00	22:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4699	\N	1	695	\N	\N	2018-09-10 00:00:00+07	0695.64.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5728	\N	1	2288	\N	\N	2019-09-23 00:00:00+07	2288.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5730	\N	1	2290	\N	\N	2019-09-23 00:00:00+07	2290.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6455	\N	1	2693	\N	\N	2021-04-22 00:00:00+07	2693.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-04-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6441	\N	1	888	\N	\N	2021-04-15 00:00:00+07	0888.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-04-15 00:00:00+07	\N	\N	\N	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6444	\N	1	2491	\N	\N	2021-04-15 00:00:00+07	2491.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-04-15 00:00:00+07	\N	10:00	21:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7148	\N	1	2654	\N	\N	2022-01-21 00:00:00+07	2654.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5732	\N	1	2292	\N	\N	2019-09-23 00:00:00+07	2292.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5735	\N	1	2270	\N	\N	2019-09-24 00:00:00+07	2270.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6452	\N	1	2497	\N	\N	2021-04-21 00:00:00+07	2497.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-04-21 00:00:00+07	\N	09:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5736	\N	1	2294	\N	\N	2019-09-24 00:00:00+07	2294.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5738	\N	1	2296	\N	\N	2019-09-24 00:00:00+07	2296.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5739	\N	1	2297	\N	\N	2019-09-24 00:00:00+07	2297.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6480	\N	1	2703	\N	\N	2021-05-25 00:00:00+07	2703.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-05-25 00:00:00+07	\N	09:00	22:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6483	\N	1	894	\N	\N	2021-06-02 00:00:00+07	0894.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-06-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5740	\N	1	2298	\N	\N	2019-09-24 00:00:00+07	2298.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5750	\N	1	525	\N	\N	2019-09-26 00:00:00+07	0525.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5755	\N	1	806	\N	\N	2019-09-30 00:00:00+07	0806.64.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5790	\N	1	2321	\N	\N	2019-10-10 00:00:00+07	2321.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5806	\N	1	2516	\N	\N	2019-10-11 00:00:00+07	2516.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5857	\N	1	2541	\N	\N	2019-10-28 00:00:00+07	2541.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5868	\N	1	2547	\N	\N	2019-10-30 00:00:00+07	2547.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5882	\N	1	2553	\N	\N	2019-11-05 00:00:00+07	2553.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5611	\N	1	971	\N	\N	2019-08-12 00:00:00+07	0971.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5610	\N	1	970	\N	\N	2019-08-12 00:00:00+07	0970.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5598	\N	1	2222	\N	\N	2019-08-09 00:00:00+07	2222.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5891	\N	1	541	\N	\N	2019-11-07 00:00:00+07	0541.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5895	\N	1	542	\N	\N	2019-11-08 00:00:00+07	0542.63.103	2	2021-04-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6495	\N	1	2707	\N	\N	2021-06-14 00:00:00+07	2707.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-06-14 00:00:00+07	\N	07:00	22:00	70	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4708	\N	1	1810	\N	\N	2018-09-12 00:00:00+07	1810.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5897	\N	1	819	\N	\N	2019-11-11 00:00:00+07	0819.64.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6499	\N	1	2709	\N	\N	2021-06-15 00:00:00+07	2709.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-06-15 00:00:00+07	\N	11:00	23:00	40-50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5591	\N	1	512	\N	\N	2019-08-08 00:00:00+07	0512.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5926	\N	1	2566	\N	\N	2019-11-22 00:00:00+07	2566.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5589	\N	1	964	\N	\N	2019-08-08 00:00:00+07	0964.61.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5937	\N	1	2573	\N	\N	2019-11-29 00:00:00+07	2573.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6533	\N	1	618	\N	\N	2021-07-06 00:00:00+07	0618.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-07-06 00:00:00+07	\N	11:00	23:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5570	\N	1	2467	\N	\N	2019-08-02 00:00:00+07	2467.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6534	\N	1	619	\N	\N	2021-07-06 00:00:00+07	0619.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-07-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5564	\N	1	2465	\N	\N	2019-07-31 00:00:00+07	2465.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6541	\N	1	2713	\N	\N	2021-08-10 00:00:00+07	2713.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-08-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6560	\N	1	3963	\N	\N	2021-08-10 00:00:00+07	2714.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-08-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6563	\N	1	2523	\N	\N	2021-08-12 00:00:00+07	2523.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-08-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5965	\N	1	2355	\N	\N	2019-12-18 00:00:00+07	2355.65.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4731	\N	1	2286	\N	\N	2018-09-18 00:00:00+07	2286.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7176	\N	1	2824	\N	\N	2022-01-26 00:00:00+07	2824.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-26 00:00:00+07	\N	08:00	22:00	22	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7209	\N	1	2661	\N	\N	2022-02-02 00:00:00+07	2661.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-02 00:00:00+07	\N	10:00	20:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4996	\N	1	1957	\N	\N	2019-02-20 00:00:00+07	1957.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7211	\N	1	1132	\N	\N	2022-02-02 00:00:00+07	1132.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-02 00:00:00+07	\N	11:00	23:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6671	\N	1	2718	\N	\N	2021-08-20 00:00:00+07	2718.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-08-20 00:00:00+07	\N	08:00	21:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5970	\N	1	830	\N	\N	2019-12-26 00:00:00+07	0830.64.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5972	\N	1	831	\N	\N	2019-12-30 00:00:00+07	0831.64.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5973	\N	1	2359	\N	\N	2019-12-30 00:00:00+07	2359.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5977	\N	1	2361	\N	\N	2020-01-02 00:00:00+07	2361.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4998	\N	1	1959	\N	\N	2019-02-20 00:00:00+07	1959.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-10-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5984	\N	1	552	\N	\N	2020-01-07 00:00:00+07	0552.63.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5985	\N	1	553	\N	\N	2020-01-07 00:00:00+07	0553.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4999	\N	1	1960	\N	\N	2019-02-20 00:00:00+07	1960.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5987	\N	1	2364	\N	\N	2020-01-07 00:00:00+07	2364.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5988	\N	1	2365	\N	\N	2020-01-07 00:00:00+07	2365.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5989	\N	1	2580	\N	\N	2020-01-07 00:00:00+07	2580.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6001	\N	1	2586	\N	\N	2020-01-20 00:00:00+07	2586.62.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6002	\N	1	2370	\N	\N	2020-01-20 00:00:00+07	2370.65.103	2	2020-12-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6003	\N	1	2371	\N	\N	2020-01-20 00:00:00+07	2371.65.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6004	\N	1	835	\N	\N	2020-01-20 00:00:00+07	0835.64.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3928	\N	2	2179	\N	\N	2017-08-02 00:00:00+07	2179.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6005	\N	1	836	\N	\N	2020-01-20 00:00:00+07	0836.64.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3936	\N	2	2183	\N	\N	2017-08-14 00:00:00+07	2183.62.103	1	\N	2018-06-17 00:00:00+07	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6006	\N	1	837	\N	\N	2020-01-20 00:00:00+07	0837.64.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6010	\N	1	2588	\N	\N	2020-01-22 00:00:00+07	2588.62.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6015	\N	1	2590	\N	\N	2020-01-22 00:00:00+07	2590.62.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6016	\N	1	2591	\N	\N	2020-01-22 00:00:00+07	2591.62.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6018	\N	1	2593	\N	\N	2020-01-27 00:00:00+07	2593.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6020	\N	1	2373	\N	\N	2020-01-28 00:00:00+07	2373.65.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6023	\N	1	2375	\N	\N	2020-02-04 00:00:00+07	2375.65.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6025	\N	1	2376	\N	\N	2020-02-05 00:00:00+07	2376.65.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6026	\N	1	2377	\N	\N	2020-02-05 00:00:00+07	2377.65.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4778	\N	1	1855	\N	\N	2018-10-03 00:00:00+07	1855.65.103	1	\N	2019-09-01 00:00:00+07	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6028	\N	1	2594	\N	\N	2020-02-05 00:00:00+07	2594.62.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6041	\N	1	2598	\N	\N	2020-02-26 00:00:00+07	2598.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2020-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6045	\N	1	2600	\N	\N	2020-02-28 00:00:00+07	2600.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5000	\N	1	1961	\N	\N	2019-02-20 00:00:00+07	1961.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5001	\N	1	1962	\N	\N	2019-02-20 00:00:00+07	1962.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6046	\N	1	1014	\N	\N	2020-03-02 00:00:00+07	1014.61.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3991	\N	1	176	\N	\N	2009-03-01 00:00:00+07	0176.65.303	2	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6690	\N	1	628	\N	\N	2021-08-26 00:00:00+07	0628.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-08-26 00:00:00+07	\N	09:00	20:00	6	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5525	\N	1	2205	\N	\N	2019-07-22 00:00:00+07	2205.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5002	\N	1	2334	\N	\N	2019-02-20 00:00:00+07	2334.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-06-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6049	\N	1	840	\N	\N	2020-03-04 00:00:00+07	0840.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6698	\N	1	2530	\N	\N	2021-08-31 00:00:00+07	2530.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-08-31 00:00:00+07	\N	10:00	23:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6706	\N	1	2722	\N	\N	2021-09-06 00:00:00+07	2722.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6054	\N	1	2601	\N	\N	2020-03-10 00:00:00+07	2601.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6063	\N	1	2386	\N	\N	2020-03-20 00:00:00+07	2386.65.103	2	2021-01-07 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-03-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6066	\N	1	2605	\N	\N	2020-03-30 00:00:00+07	2605.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6067	\N	1	2387	\N	\N	2020-03-30 00:00:00+07	2387.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5523	\N	1	2204	\N	\N	2019-07-19 00:00:00+07	2204.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-06 00:00:00+07	\N	\N	\N	26	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4787	\N	1	1859	\N	\N	2018-10-11 00:00:00+07	1859.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5003	\N	1	1963	\N	\N	2019-02-20 00:00:00+07	1963.65.103	2	2020-03-06 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6707	\N	1	2723	\N	\N	2021-09-06 00:00:00+07	2723.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5303	\N	1	2109	\N	\N	2019-05-16 00:00:00+07	2109.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5298	\N	1	2386	\N	\N	2019-05-16 00:00:00+07	2386.62.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6710	\N	1	2725	\N	\N	2021-09-07 00:00:00+07	2725.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-07 00:00:00+07	\N	08:00	21:00	15	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6749	\N	1	2728	\N	\N	2021-09-14 00:00:00+07	2728.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-14 00:00:00+07	\N	12:00	20:00	60	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6750	\N	1	918	\N	\N	2021-09-15 00:00:00+07	0918.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-15 00:00:00+07	\N	07:00	22:00	22 - 50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6751	\N	1	2729	\N	\N	2021-09-15 00:00:00+07	2729.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-15 00:00:00+07	\N	12:00	02:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6753	\N	1	2731	\N	\N	2021-09-15 00:00:00+07	2731.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-15 00:00:00+07	\N	10:00	18:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6069	\N	1	1018	\N	\N	2020-04-09 00:00:00+07	1018.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-03-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6754	\N	1	2539	\N	\N	2021-09-15 00:00:00+07	2539.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6070	\N	1	1019	\N	\N	2020-04-09 00:00:00+07	1019.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-01-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6760	\N	1	2733	\N	\N	2021-09-16 00:00:00+07	2733.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6071	\N	1	562	\N	\N	2020-04-09 00:00:00+07	0562.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6761	\N	1	2734	\N	\N	2021-09-16 00:00:00+07	2734.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-16 00:00:00+07	\N	09:00	20:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6766	\N	1	2542	\N	\N	2021-09-20 00:00:00+07	2542.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-20 00:00:00+07	\N	07:00	23:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4155	\N	1	2207	\N	\N	2017-10-09 00:00:00+07	2207.62.103	2	2017-12-01 00:00:00+07	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6072	\N	1	2388	\N	\N	2020-04-09 00:00:00+07	2388.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4158	\N	1	2208	\N	\N	2017-10-09 00:00:00+07	2208.62.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6093	\N	1	567	\N	\N	2020-06-10 00:00:00+07	0567.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6084	\N	1	2392	\N	\N	2020-06-03 00:00:00+07	2392.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6085	\N	1	564	\N	\N	2020-06-05 00:00:00+07	0564.63.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6768	\N	1	2544	\N	\N	2021-09-20 00:00:00+07	2544.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-20 00:00:00+07	\N	08:00	22:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6781	\N	1	2547	\N	\N	2021-09-27 00:00:00+07	2547.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6088	\N	1	2606	\N	\N	2020-06-08 00:00:00+07	2606.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6785	\N	1	2742	\N	\N	2021-09-27 00:00:00+07	2742.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-27 00:00:00+07	\N	09:00	21:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6786	\N	1	2549	\N	\N	2021-09-27 00:00:00+07	2549.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-27 00:00:00+07	\N	07:00	19:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6787	\N	1	2550	\N	\N	2021-09-27 00:00:00+07	2550.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-27 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6788	\N	1	2551	\N	\N	2021-09-28 00:00:00+07	2551.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-28 00:00:00+07	\N	09:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6089	\N	1	2607	\N	\N	2020-06-08 00:00:00+07	2607.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6094	\N	1	2396	\N	\N	2020-06-10 00:00:00+07	2396.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6789	\N	1	2552	\N	\N	2021-09-28 00:00:00+07	2552.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-28 00:00:00+07	\N	13:00	20:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6790	\N	1	2553	\N	\N	2021-09-28 00:00:00+07	2553.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-28 00:00:00+07	\N	13:00	20:00	18	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6792	\N	1	2554	\N	\N	2021-09-29 00:00:00+07	2554.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-29 00:00:00+07	\N	13:00	22:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6793	\N	1	2744	\N	\N	2021-09-29 00:00:00+07	2744.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-29 00:00:00+07	\N	08:00	22:00	12	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6794	\N	1	2555	\N	\N	2021-09-29 00:00:00+07	2555.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-29 00:00:00+07	\N	16:00	22:00	25	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6798	\N	1	2556	\N	\N	2021-09-30 00:00:00+07	2556.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-30 00:00:00+07	\N	12:00	22:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6799	\N	1	2557	\N	\N	2021-09-30 00:00:00+07	2557.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-09-30 00:00:00+07	\N	\N	\N	48	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6803	\N	1	2559	\N	\N	2021-10-05 00:00:00+07	2559.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-05 00:00:00+07	\N	13:00	22:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6804	\N	1	2746	\N	\N	2021-10-05 00:00:00+07	2746.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-05 00:00:00+07	\N	07:00	20:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6806	\N	1	2561	\N	\N	2021-10-06 00:00:00+07	2561.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6807	\N	1	2747	\N	\N	2021-10-06 00:00:00+07	2747.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-06 00:00:00+07	\N	10:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6096	\N	1	2608	\N	\N	2020-06-10 00:00:00+07	2608.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6097	\N	1	849	\N	\N	2020-06-10 00:00:00+07	0849.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6813	\N	1	2564	\N	\N	2021-10-11 00:00:00+07	2564.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-11 00:00:00+07	\N	08:00	20:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6814	\N	1	2750	\N	\N	2021-10-13 00:00:00+07	2750.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-13 00:00:00+07	\N	07:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6103	\N	1	853	\N	\N	2020-06-18 00:00:00+07	0853.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6817	\N	1	2752	\N	\N	2021-10-13 00:00:00+07	2752.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-13 00:00:00+07	\N	07:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6821	\N	1	2755	\N	\N	2021-10-14 00:00:00+07	2755.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-14 00:00:00+07	\N	09:00	21:00	10	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6104	\N	1	2399	\N	\N	2020-06-24 00:00:00+07	2399.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6822	\N	1	2756	\N	\N	2021-10-14 00:00:00+07	2756.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-14 00:00:00+07	\N	10:00	21:00	100	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6826	\N	1	2567	\N	\N	2021-10-18 00:00:00+07	2567.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6827	\N	1	1084	\N	\N	2021-10-18 00:00:00+07	1084.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	10:00	21:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6840	\N	1	2568	\N	\N	2021-10-21 00:00:00+07	2568.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-21 00:00:00+07	\N	08:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4826	\N	1	439	\N	\N	2018-10-24 00:00:00+07	0439.63.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6108	\N	1	854	\N	\N	2020-06-26 00:00:00+07	0854.64.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6118	\N	1	2402	\N	\N	2020-07-07 00:00:00+07	2402.65.103	2	2021-08-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6842	\N	1	2569	\N	\N	2021-10-22 00:00:00+07	2569.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7217	\N	1	1134	\N	\N	2022-02-03 00:00:00+07	1134.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4829	\N	1	1893	\N	\N	2018-10-25 00:00:00+07	1893.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6843	\N	1	2761	\N	\N	2021-10-22 00:00:00+07	2761.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6844	\N	1	2570	\N	\N	2021-10-25 00:00:00+07	2570.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5495	\N	1	2198	\N	\N	2019-07-11 00:00:00+07	2198.65.103	2	2020-01-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6845	\N	1	2571	\N	\N	2021-10-25 00:00:00+07	2571.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-25 00:00:00+07	\N	14:00	00:00	50	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7218	\N	1	1135	\N	\N	2022-02-03 00:00:00+07	1135.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-03 00:00:00+07	\N	08:00	23:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7250	\N	1	1145	\N	\N	2022-02-09 00:00:00+07	1145.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6119	\N	1	2403	\N	\N	2020-07-07 00:00:00+07	2403.65.103	2	2020-09-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7247	\N	1	1144	\N	\N	2022-02-08 00:00:00+07	1144.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6120	\N	1	2404	\N	\N	2020-07-07 00:00:00+07	2404.65.103	2	2021-08-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7251	\N	1	2668	\N	\N	2022-02-09 00:00:00+07	2668.65.103	2	2022-04-04 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6851	\N	1	2572	\N	\N	2021-10-26 00:00:00+07	2572.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4839	\N	1	441	\N	\N	2018-11-06 00:00:00+07	0441.63.103	1	\N	\N	SA	\N	10113	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6852	\N	1	2763	\N	\N	2021-10-26 00:00:00+07	2763.62.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-26 00:00:00+07	\N	13:00	21:00	30	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6121	\N	1	2405	\N	\N	2020-07-07 00:00:00+07	2405.65.103	2	2021-05-30 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7253	\N	1	1146	\N	\N	2022-02-10 00:00:00+07	1146.61.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7254	\N	1	2669	\N	\N	2022-02-10 00:00:00+07	2669.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4843	\N	1	1898	\N	\N	2018-11-08 00:00:00+07	1898.65.103	1	\N	\N	SA	\N	10113	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5200	\N	1	2071	\N	\N	2019-04-10 00:00:00+07	2071.65.103	1	\N	\N	SA	\N	10113	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6857	\N	1	2574	\N	\N	2021-10-28 00:00:00+07	2574.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-28 00:00:00+07	\N	08:00	21:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6122	\N	1	571	\N	\N	2020-07-07 00:00:00+07	0571.63.103	2	2021-07-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-10-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6123	\N	1	572	\N	\N	2020-07-07 00:00:00+07	0572.63.103	2	2020-08-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6124	\N	1	2613	\N	\N	2020-07-07 00:00:00+07	2613.62.103	2	2021-08-01 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6125	\N	1	856	\N	\N	2020-07-07 00:00:00+07	0856.64.103	2	2021-05-30 00:00:00+07	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2021-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7255	\N	1	2670	\N	\N	2022-02-10 00:00:00+07	2670.65.103	1	\N	\N	SA	\N	10113	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4434	\N	1	636	\N	\N	2018-03-09 00:00:00+07	0636.64.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
422	\N	2	432	\N	\N	2011-02-10 00:00:00+07	0432.62.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3808	\N	2	2159	\N	\N	2017-06-13 00:00:00+07	2159.62.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3647	\N	1	1573	\N	\N	2017-02-24 00:00:00+07	1573.65.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
421	\N	2	90	\N	\N	2011-02-10 00:00:00+07	0090.61.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
433	\N	2	419	\N	\N	2011-02-11 00:00:00+07	0419.62.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
429	\N	2	83	\N	\N	2011-02-10 00:00:00+07	0083.65.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
472	\N	2	1691	\N	\N	2013-09-12 00:00:00+07	1691.62.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4481	\N	1	872	\N	\N	2018-04-02 00:00:00+07	0872.61.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
428	\N	2	50	\N	\N	2011-02-10 00:00:00+07	0050.65.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
427	\N	2	53	\N	\N	2011-02-10 00:00:00+07	0053.64.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
426	\N	2	54	\N	\N	2011-02-10 00:00:00+07	0054.64.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5024	\N	1	2340	\N	\N	2019-02-26 00:00:00+07	2340.62.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
425	\N	2	29	\N	\N	2011-02-10 00:00:00+07	0029.63.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
432	\N	2	915	\N	\N	2011-02-10 00:00:00+07	0915.62.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2997	\N	1	1535	\N	\N	2016-12-01 00:00:00+07	1535.65.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
424	\N	2	429	\N	\N	2011-02-10 00:00:00+07	0429.62.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
430	\N	2	81	\N	\N	2011-02-10 00:00:00+07	0081.65.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
431	\N	2	77	\N	\N	2011-02-10 00:00:00+07	0077.65.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
423	\N	2	439	\N	\N	2011-02-10 00:00:00+07	0439.62.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
462	\N	2	51	\N	\N	2011-02-21 00:00:00+07	0051.64.302	1	\N	\N	SA	\N	10118	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4288	\N	1	1674	\N	\N	2017-12-04 00:00:00+07	1674.65.309	1	\N	\N	SA	\N	10123	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4419	\N	1	2234	\N	\N	2018-03-01 00:00:00+07	2234.62.309	1	\N	2019-06-09 00:00:00+07	SA	\N	10123	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
486	\N	2	384	\N	\N	2014-09-09 00:00:00+07	0384.64.309	1	\N	2019-06-08 00:00:00+07	SA	\N	10123	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4957	\N	1	740	\N	\N	2019-02-11 00:00:00+07	0740.64.309	1	\N	2019-06-08 00:00:00+07	SA	\N	10123	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6332	\N	1	2465	\N	\N	2020-12-28 00:00:00+07	2465.65.309	1	\N	\N	SA	\N	10123	Amalia Mahardani	\N	\N	f	f	2020-12-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4920	\N	1	2315	\N	\N	2018-12-28 00:00:00+07	2315.62.309	2	2019-03-01 00:00:00+07	\N	SA	\N	10123	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3339	\N	1	1703	\N	\N	2017-03-01 00:00:00+07	1703.61.307	1	\N	\N	SA	\N	10145	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4916	\N	1	2314	\N	\N	2018-12-20 00:00:00+07	2314.62.307	1	\N	\N	SA	\N	10145	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3995	\N	1	2003	\N	\N	2008-01-01 00:00:00+07	2003.65.303	2	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
441	\N	2	1392	\N	\N	2011-02-11 00:00:00+07	1392.62.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4037	\N	1	176	\N	\N	2009-03-01 00:00:00+07	0176.65.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3986	\N	1	45	\N	\N	2003-01-01 00:00:00+07	0045.64.303	2	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
491	\N	2	1413	\N	\N	2016-02-10 00:00:00+07	1413.65.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4036	\N	1	9864	\N	\N	2005-04-01 00:00:00+07	9864.65.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4010	\N	1	9865	\N	\N	2006-12-01 00:00:00+07	9865.65.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
434	\N	2	47	\N	\N	2011-02-11 00:00:00+07	0047.64.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
435	\N	2	89	\N	\N	2011-02-11 00:00:00+07	0089.61.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
436	\N	2	415	\N	\N	2011-02-11 00:00:00+07	0415.62.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
437	\N	2	413	\N	\N	2011-02-11 00:00:00+07	0413.62.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
439	\N	2	191	\N	\N	2011-02-11 00:00:00+07	0191.64.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
438	\N	2	70	\N	\N	2011-02-11 00:00:00+07	0070.65.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
492	\N	2	1943	\N	\N	2016-02-29 00:00:00+07	1943.62.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4039	\N	1	964	\N	\N	2007-08-01 00:00:00+07	0964.62.303	1	\N	\N	SA	\N	10148	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
420	\N	2	43	\N	\N	2011-02-10 00:00:00+07	0043.64.308	1	\N	\N	SA	\N	10150	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
442	\N	2	407	\N	\N	2011-02-11 00:00:00+07	0407.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4273	\N	1	600	\N	\N	2017-11-13 00:00:00+07	0600.64.305	2	2020-02-10 00:00:00+07	\N	SA	\N	10153	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
446	\N	2	1373	\N	\N	2011-02-11 00:00:00+07	1373.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
445	\N	2	411	\N	\N	2011-02-11 00:00:00+07	0411.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
444	\N	2	683	\N	\N	2011-02-11 00:00:00+07	0683.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
443	\N	2	409	\N	\N	2011-02-11 00:00:00+07	0409.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7440	\N	1	710	\N	\N	2022-06-06 00:00:00+07	0710.63.305	1	\N	\N	SA	\N	10153	Amalia Mahardani	\N	\N	f	f	2022-06-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
496	\N	2	1412	\N	\N	2016-03-15 00:00:00+07	1412.65.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7423	\N	1	2893	\N	\N	2022-05-24 00:00:00+07	2893.62.305	1	\N	\N	SA	\N	10153	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	10:00	22:00	64	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
495	\N	2	1607	\N	\N	2016-03-10 00:00:00+07	1607.62.305	2	2017-02-28 00:00:00+07	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4927	\N	2	2317	\N	\N	2019-01-15 00:00:00+07	2317.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
488	\N	2	402	\N	\N	2015-02-05 00:00:00+07	0402.64.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3407	\N	1	2106	\N	\N	2017-03-01 00:00:00+07	2106.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4615	\N	1	1767	\N	\N	2018-07-18 00:00:00+07	1767.65.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3561	\N	1	2106	\N	\N	2017-05-09 00:00:00+07	2106.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3904	\N	1	360	\N	\N	2017-07-25 00:00:00+07	0360.63.305	2	2020-07-01 00:00:00+07	\N	SA	\N	10153	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
467	\N	2	1649	\N	\N	2012-08-30 00:00:00+07	1649.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5009	\N	1	2335	\N	\N	2019-02-21 00:00:00+07	2335.62.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2969	\N	1	0	\N	\N	2017-03-23 00:00:00+07	0000.65.305	1	\N	\N	SA	\N	10153	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5618	\N	1	2475	\N	\N	2019-08-13 00:00:00+07	2475.62.305	1	\N	\N	SA	\N	10153	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6494	\N	1	1057	\N	\N	2021-06-11 00:00:00+07	1057.61.306	1	\N	\N	SA	\N	10154	Amalia Mahardani	\N	\N	f	f	2021-06-11 00:00:00+07	\N	08:00	22:00	5	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
449	\N	2	405	\N	\N	2011-02-11 00:00:00+07	0405.62.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5246	\N	1	475	\N	\N	2019-04-30 00:00:00+07	0475.63.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5957	\N	1	1004	\N	\N	2019-12-11 00:00:00+07	1004.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5952	\N	1	1001	\N	\N	2019-12-09 00:00:00+07	1001.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
470	\N	2	634	\N	\N	2013-05-10 00:00:00+07	0634.65.306	2	2017-08-25 00:00:00+07	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5951	\N	1	1000	\N	\N	2019-12-09 00:00:00+07	1000.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5966	\N	1	2356	\N	\N	2019-12-23 00:00:00+07	2356.65.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5950	\N	1	999	\N	\N	2019-12-09 00:00:00+07	0999.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5949	\N	1	998	\N	\N	2019-12-09 00:00:00+07	0998.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
490	\N	2	1818	\N	\N	2015-05-26 00:00:00+07	1818.62.306	1	\N	2019-06-11 00:00:00+07	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4083	\N	1	2199	\N	\N	2017-09-25 00:00:00+07	2199.62.306	2	2020-02-17 00:00:00+07	2018-06-18 00:00:00+07	SA	\N	10154	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5948	\N	1	997	\N	\N	2019-12-09 00:00:00+07	0997.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7347	\N	1	1010	\N	\N	2022-03-23 00:00:00+07	1010.64.306	1	\N	\N	SA	\N	10154	Amalia Mahardani	\N	\N	f	f	2022-03-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3582	\N	1	543	\N	\N	2017-05-02 00:00:00+07	0543.64.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5947	\N	1	996	\N	\N	2019-12-09 00:00:00+07	0996.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5946	\N	1	995	\N	\N	2019-12-09 00:00:00+07	0995.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5976	\N	1	1008	\N	\N	2019-12-31 00:00:00+07	1008.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5945	\N	1	994	\N	\N	2019-12-09 00:00:00+07	0994.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4157	\N	1	858	\N	\N	2017-10-09 00:00:00+07	0858.61.306	1	\N	\N	SA	\N	10154	Amalia Mahardani	\N	\N	f	f	2020-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6047	\N	1	1015	\N	\N	2020-03-02 00:00:00+07	1015.61.306	1	\N	\N	SA	\N	10154	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7470	\N	1	1029	\N	\N	2022-06-15 00:00:00+07	1029.64.306	1	\N	\N	SA	\N	10154	Amalia Mahardani	\N	\N	f	f	2022-06-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3418	\N	1	497	\N	\N	2017-03-01 00:00:00+07	0497.64.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7281	\N	1	2852	\N	\N	2022-02-24 00:00:00+07	2852.62.306	1	\N	\N	SA	\N	10154	Amalia Mahardani	\N	\N	f	f	2022-02-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
448	\N	2	88	\N	\N	2011-02-11 00:00:00+07	0088.61.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4953	\N	1	2323	\N	\N	2019-02-07 00:00:00+07	2323.62.306	1	\N	2020-06-01 00:00:00+07	SA	\N	10154	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5958	\N	1	1005	\N	\N	2019-12-11 00:00:00+07	1005.61.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
474	\N	2	1706	\N	\N	2013-12-24 00:00:00+07	1706.62.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5662	\N	1	801	\N	\N	2019-09-05 00:00:00+07	0801.64.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5933	\N	1	546	\N	\N	2019-11-25 00:00:00+07	0546.63.306	1	\N	2020-06-01 00:00:00+07	SA	\N	10154	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
447	\N	2	87	\N	\N	2011-02-11 00:00:00+07	0087.61.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2917	\N	1	1557	\N	\N	2016-12-01 00:00:00+07	1557.65.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
465	\N	2	282	\N	\N	2011-02-24 00:00:00+07	0282.62.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2916	\N	1	1976	\N	\N	2016-12-01 00:00:00+07	1976.65.306	2	2017-09-01 00:00:00+07	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5921	\N	1	2564	\N	\N	2019-11-19 00:00:00+07	2564.62.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5917	\N	1	2345	\N	\N	2019-11-18 00:00:00+07	2345.65.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
497	\N	2	1454	\N	\N	2016-05-09 00:00:00+07	1454.65.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5910	\N	1	2342	\N	\N	2019-11-14 00:00:00+07	2342.65.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5906	\N	1	823	\N	\N	2019-11-13 00:00:00+07	0823.64.306	1	\N	\N	SA	\N	10154	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6944	\N	1	658	\N	\N	2021-11-26 00:00:00+07	0658.63.306	1	\N	\N	SA	\N	10154	Amalia Mahardani	\N	\N	f	f	2021-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6989	\N	1	1106	\N	\N	2021-12-15 00:00:00+07	1106.61.306	1	\N	\N	SA	\N	10154	Amalia Mahardani	\N	\N	f	f	2021-12-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5040	\N	1	2346	\N	\N	2019-02-27 00:00:00+07	2346.62.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6448	\N	1	2493	\N	\N	2021-04-20 00:00:00+07	2493.65.306	1	\N	\N	SA	\N	10154	Amalia Mahardani	\N	\N	f	f	2021-04-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4910	\N	1	2312	\N	\N	2018-12-14 00:00:00+07	2312.62.306	2	2020-04-01 00:00:00+07	\N	SA	\N	10154	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4955	\N	1	2325	\N	\N	2019-02-08 00:00:00+07	2325.62.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
450	\N	2	339	\N	\N	2011-02-11 00:00:00+07	0339.62.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4494	\N	1	874	\N	\N	2018-04-10 00:00:00+07	0874.61.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3419	\N	1	2054	\N	\N	2017-03-01 00:00:00+07	2054.62.306	1	\N	\N	SA	\N	10154	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
454	\N	2	546	\N	\N	2011-02-11 00:00:00+07	0546.62.304	1	\N	\N	SA	\N	10206	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
452	\N	2	182	\N	\N	2011-02-11 00:00:00+07	0182.65.304	1	\N	\N	SA	\N	10206	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
455	\N	2	406	\N	\N	2011-02-11 00:00:00+07	0406.62.304	1	\N	\N	SA	\N	10206	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
494	\N	2	428	\N	\N	2016-03-03 00:00:00+07	0428.65.304	1	\N	\N	SA	\N	10206	Amalia Mahardani	\N	\N	f	f	2021-04-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
453	\N	2	42	\N	\N	2011-02-11 00:00:00+07	0042.64.304	1	\N	\N	SA	\N	10206	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
456	\N	2	68	\N	\N	2011-02-11 00:00:00+07	0068.65.304	1	\N	\N	SA	\N	10206	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6462	\N	1	2695	\N	\N	2021-05-03 00:00:00+07	2695.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-05-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1605	\N	2	119	\N	\N	2010-02-25 00:00:00+07	0119.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1606	\N	2	120	\N	\N	2010-02-25 00:00:00+07	0120.61.600	1	\N	\N	SA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1607	\N	2	121	\N	\N	2010-02-25 00:00:00+07	0121.61.600	2	2019-07-01 00:00:00+07	\N	SA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1608	\N	2	122	\N	\N	2010-02-25 00:00:00+07	0122.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1609	\N	2	123	\N	\N	2010-02-25 00:00:00+07	0123.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1610	\N	2	124	\N	\N	2010-02-25 00:00:00+07	0124.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1611	\N	2	125	\N	\N	2010-02-25 00:00:00+07	0125.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1612	\N	2	126	\N	\N	2010-02-25 00:00:00+07	0126.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1613	\N	2	127	\N	\N	2010-02-25 00:00:00+07	0127.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1614	\N	2	293	\N	\N	2010-02-25 00:00:00+07	0293.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1615	\N	2	78	\N	\N	2010-02-25 00:00:00+07	0078.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1616	\N	2	440	\N	\N	2010-02-25 00:00:00+07	0440.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1617	\N	2	441	\N	\N	2010-02-25 00:00:00+07	0441.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1618	\N	2	442	\N	\N	2010-02-25 00:00:00+07	0442.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1619	\N	2	443	\N	\N	2010-02-25 00:00:00+07	0443.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1620	\N	2	444	\N	\N	2010-02-25 00:00:00+07	0444.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1621	\N	2	445	\N	\N	2010-02-25 00:00:00+07	0445.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1622	\N	2	446	\N	\N	2010-02-25 00:00:00+07	0446.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1623	\N	2	447	\N	\N	2010-02-25 00:00:00+07	0447.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1624	\N	2	448	\N	\N	2010-02-25 00:00:00+07	0448.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1626	\N	2	450	\N	\N	2010-02-25 00:00:00+07	0450.62.600	1	\N	\N	SA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1627	\N	2	456	\N	\N	2010-02-25 00:00:00+07	0456.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1628	\N	2	457	\N	\N	2010-02-25 00:00:00+07	0457.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1629	\N	2	458	\N	\N	2010-02-25 00:00:00+07	0458.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1630	\N	2	459	\N	\N	2010-02-25 00:00:00+07	0459.62.600	1	\N	\N	SA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1631	\N	2	460	\N	\N	2010-02-25 00:00:00+07	0460.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1632	\N	2	461	\N	\N	2010-02-25 00:00:00+07	0461.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1633	\N	2	462	\N	\N	2010-02-25 00:00:00+07	0462.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1634	\N	2	463	\N	\N	2010-02-25 00:00:00+07	0463.62.600	1	\N	2020-05-01 00:00:00+07	OA	\N	10205	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1635	\N	2	464	\N	\N	2010-02-25 00:00:00+07	0464.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1636	\N	2	465	\N	\N	2010-02-25 00:00:00+07	0465.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1637	\N	2	466	\N	\N	2010-02-25 00:00:00+07	0466.62.600	1	\N	2020-06-01 00:00:00+07	OA	\N	10205	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1638	\N	2	467	\N	\N	2010-02-25 00:00:00+07	0467.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1640	\N	2	469	\N	\N	2010-02-25 00:00:00+07	0469.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1641	\N	2	470	\N	\N	2010-02-25 00:00:00+07	0470.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1642	\N	2	471	\N	\N	2010-02-25 00:00:00+07	0471.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1643	\N	2	473	\N	\N	2010-02-25 00:00:00+07	0473.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1644	\N	2	476	\N	\N	2010-02-25 00:00:00+07	0476.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1645	\N	2	541	\N	\N	2010-02-25 00:00:00+07	0541.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1646	\N	2	658	\N	\N	2010-02-25 00:00:00+07	0658.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1647	\N	2	659	\N	\N	2010-02-25 00:00:00+07	0659.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1648	\N	2	731	\N	\N	2010-02-25 00:00:00+07	0731.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1649	\N	2	921	\N	\N	2010-02-25 00:00:00+07	0921.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1650	\N	2	1026	\N	\N	2010-02-25 00:00:00+07	1026.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1651	\N	2	1027	\N	\N	2010-02-25 00:00:00+07	1027.62.600	1	\N	\N	SA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1652	\N	2	1347	\N	\N	2010-02-25 00:00:00+07	1347.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1653	\N	2	75	\N	\N	2010-02-25 00:00:00+07	0075.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1654	\N	2	76	\N	\N	2010-02-25 00:00:00+07	0076.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1655	\N	2	77	\N	\N	2010-02-25 00:00:00+07	0077.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1656	\N	2	79	\N	\N	2010-02-25 00:00:00+07	0079.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1657	\N	2	153	\N	\N	2010-02-25 00:00:00+07	0153.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1658	\N	2	112	\N	\N	2010-02-25 00:00:00+07	0112.65.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1659	\N	2	113	\N	\N	2010-02-25 00:00:00+07	0113.65.600	1	\N	\N	SA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1660	\N	2	114	\N	\N	2010-02-25 00:00:00+07	0114.65.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1661	\N	2	117	\N	\N	2010-02-25 00:00:00+07	0117.65.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1662	\N	2	120	\N	\N	2010-04-30 00:00:00+07	0120.61.600	1	\N	\N	SA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1663	\N	2	449	\N	\N	2010-04-30 00:00:00+07	0449.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1665	\N	2	471	\N	\N	2011-02-28 00:00:00+07	0471.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1666	\N	2	671	\N	\N	2013-02-28 00:00:00+07	0671.61.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1667	\N	2	672	\N	\N	2013-02-28 00:00:00+07	0672.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1668	\N	2	1671	\N	\N	2013-04-30 00:00:00+07	1671.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1669	\N	2	1670	\N	\N	2013-04-30 00:00:00+07	1670.62.600	1	\N	2020-06-01 00:00:00+07	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-07-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1670	\N	2	1669	\N	\N	2013-04-30 00:00:00+07	1669.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1671	\N	2	676	\N	\N	2013-05-31 00:00:00+07	0676.61.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1673	\N	2	468	\N	\N	2016-02-29 00:00:00+07	0468.62.600	1	\N	2020-06-01 00:00:00+07	OA	\N	10205	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1674	\N	2	1954	\N	\N	2016-02-29 00:00:00+07	1954.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1675	\N	2	1945	\N	\N	2016-02-29 00:00:00+07	1945.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1677	\N	2	299	\N	\N	2016-03-31 00:00:00+07	0299.63.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1678	\N	2	470	\N	\N	2016-03-31 00:00:00+07	0470.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1679	\N	2	459	\N	\N	2016-04-29 00:00:00+07	0459.62.600	1	\N	\N	SA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4256	\N	1	1671	\N	\N	2017-11-01 00:00:00+07	1671.65.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4343	\N	2	396	\N	\N	2018-01-11 00:00:00+07	0396.63.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4347	\N	2	611	\N	\N	2018-01-11 00:00:00+07	0611.64.600	2	2019-10-01 00:00:00+07	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-02-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4349	\N	1	2219	\N	\N	2018-01-11 00:00:00+07	2219.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4425	\N	1	1696	\N	\N	2018-03-05 00:00:00+07	1696.65.600	1	\N	2019-06-09 00:00:00+07	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4509	\N	1	661	\N	\N	2018-04-13 00:00:00+07	0661.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4840	\N	1	2297	\N	\N	2018-11-06 00:00:00+07	2297.62.600	1	\N	2019-06-09 00:00:00+07	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4941	\N	1	456	\N	\N	2019-01-30 00:00:00+07	0456.63.600	2	2020-02-01 00:00:00+07	\N	OA	\N	10205	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5023	\N	1	2339	\N	\N	2019-02-25 00:00:00+07	2339.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4954	\N	1	2324	\N	\N	2019-02-07 00:00:00+07	2324.62.600	1	\N	2020-06-01 00:00:00+07	OA	\N	10205	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5027	\N	1	2342	\N	\N	2019-02-26 00:00:00+07	2342.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5028	\N	1	2343	\N	\N	2019-02-26 00:00:00+07	2343.62.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4974	\N	1	1946	\N	\N	2019-02-13 00:00:00+07	1946.65.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5068	\N	1	751	\N	\N	2019-03-06 00:00:00+07	0751.64.600	1	\N	\N	OA	\N	10205	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5106	\N	1	2354	\N	\N	2019-03-15 00:00:00+07	2354.62.600	1	\N	2020-06-01 00:00:00+07	OA	\N	10205	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5116	\N	1	2035	\N	\N	2019-03-18 00:00:00+07	2035.65.600	1	\N	\N	SA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5126	\N	1	910	\N	\N	2019-03-20 00:00:00+07	0910.61.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5177	\N	1	2062	\N	\N	2019-04-04 00:00:00+07	2062.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5356	\N	1	2127	\N	\N	2019-05-27 00:00:00+07	2127.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5484	\N	1	2194	\N	\N	2019-07-09 00:00:00+07	2194.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5485	\N	1	2421	\N	\N	2019-07-09 00:00:00+07	2421.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5508	\N	1	2432	\N	\N	2019-07-16 00:00:00+07	2432.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5511	\N	1	2201	\N	\N	2019-07-17 00:00:00+07	2201.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5514	\N	1	960	\N	\N	2019-07-17 00:00:00+07	0960.61.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5577	\N	1	788	\N	\N	2019-08-05 00:00:00+07	0788.64.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5540	\N	1	962	\N	\N	2019-07-24 00:00:00+07	0962.61.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5541	\N	1	2210	\N	\N	2019-07-24 00:00:00+07	2210.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5550	\N	1	2458	\N	\N	2019-07-25 00:00:00+07	2458.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5551	\N	1	2459	\N	\N	2019-07-25 00:00:00+07	2459.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5557	\N	1	2211	\N	\N	2019-07-26 00:00:00+07	2211.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5599	\N	1	792	\N	\N	2019-08-09 00:00:00+07	0792.64.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5600	\N	1	2473	\N	\N	2019-08-09 00:00:00+07	2473.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5601	\N	1	969	\N	\N	2019-08-09 00:00:00+07	0969.61.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5602	\N	1	2223	\N	\N	2019-08-09 00:00:00+07	2223.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5604	\N	1	511	\N	\N	2019-08-09 00:00:00+07	0511.63.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5605	\N	1	2224	\N	\N	2019-08-09 00:00:00+07	2224.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5606	\N	1	793	\N	\N	2019-08-09 00:00:00+07	0793.64.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5607	\N	1	794	\N	\N	2019-08-09 00:00:00+07	0794.64.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5608	\N	1	2474	\N	\N	2019-08-09 00:00:00+07	2474.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5609	\N	1	513	\N	\N	2019-08-09 00:00:00+07	0513.63.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5632	\N	1	2239	\N	\N	2019-08-15 00:00:00+07	2239.65.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5756	\N	1	2495	\N	\N	2019-10-01 00:00:00+07	2495.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5757	\N	1	2496	\N	\N	2019-10-01 00:00:00+07	2496.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5758	\N	1	2497	\N	\N	2019-10-01 00:00:00+07	2497.62.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5796	\N	1	987	\N	\N	2019-10-10 00:00:00+07	0987.61.600	1	\N	\N	OA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5928	\N	1	2568	\N	\N	2019-11-22 00:00:00+07	2568.62.600	1	\N	\N	SA	\N	10205	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5991	\N	1	2581	\N	\N	2020-01-09 00:00:00+07	2581.62.600	1	\N	\N	OA	\N	10205	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6077	\N	1	846	\N	\N	2020-04-28 00:00:00+07	0846.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6319	\N	1	597	\N	\N	2020-12-10 00:00:00+07	0597.63.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	09:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6320	\N	1	1038	\N	\N	2020-12-10 00:00:00+07	1038.61.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	09:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6321	\N	1	2649	\N	\N	2020-12-10 00:00:00+07	2649.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	09:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6322	\N	1	878	\N	\N	2020-12-10 00:00:00+07	0878.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	09:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6323	\N	1	2460	\N	\N	2020-12-10 00:00:00+07	2460.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	09:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6324	\N	1	2461	\N	\N	2020-12-10 00:00:00+07	2461.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2020-12-10 00:00:00+07	\N	09:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6431	\N	1	2489	\N	\N	2021-04-14 00:00:00+07	2489.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-04-14 00:00:00+07	\N	00:00	23:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6437	\N	1	2688	\N	\N	2021-04-14 00:00:00+07	2688.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-04-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6440	\N	1	2490	\N	\N	2021-04-15 00:00:00+07	2490.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-04-15 00:00:00+07	\N	08:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6442	\N	1	607	\N	\N	2021-04-15 00:00:00+07	0607.63.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-04-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6443	\N	1	2690	\N	\N	2021-04-15 00:00:00+07	2690.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-04-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6450	\N	1	2495	\N	\N	2021-04-20 00:00:00+07	2495.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-04-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6469	\N	1	891	\N	\N	2021-05-10 00:00:00+07	0891.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-05-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6474	\N	1	2502	\N	\N	2021-05-19 00:00:00+07	2502.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-05-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6475	\N	1	2503	\N	\N	2021-05-19 00:00:00+07	2503.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-05-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6476	\N	1	2702	\N	\N	2021-05-19 00:00:00+07	2702.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-05-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6477	\N	1	893	\N	\N	2021-05-20 00:00:00+07	0893.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-05-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6478	\N	1	2504	\N	\N	2021-05-20 00:00:00+07	2504.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-05-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6505	\N	1	2510	\N	\N	2021-06-18 00:00:00+07	2510.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6507	\N	1	2511	\N	\N	2021-06-21 00:00:00+07	2511.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6513	\N	1	2711	\N	\N	2021-06-24 00:00:00+07	2711.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6514	\N	1	899	\N	\N	2021-06-24 00:00:00+07	0899.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6517	\N	1	900	\N	\N	2021-06-28 00:00:00+07	0900.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6520	\N	1	2516	\N	\N	2021-06-30 00:00:00+07	2516.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6524	\N	1	902	\N	\N	2021-06-30 00:00:00+07	0902.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6522	\N	1	2518	\N	\N	2021-06-30 00:00:00+07	2518.65.600	1	\N	\N	SA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6525	\N	1	903	\N	\N	2021-06-30 00:00:00+07	0903.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-06-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6527	\N	1	1059	\N	\N	2021-07-01 00:00:00+07	1059.61.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-07-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6535	\N	1	2520	\N	\N	2021-07-08 00:00:00+07	2520.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-07-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6670	\N	1	1063	\N	\N	2021-08-19 00:00:00+07	1063.61.600	2	2021-10-29 00:00:00+07	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-08-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6673	\N	1	2524	\N	\N	2021-08-23 00:00:00+07	2524.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-08-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6674	\N	1	1064	\N	\N	2021-08-23 00:00:00+07	1064.61.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-08-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6676	\N	1	910	\N	\N	2021-08-23 00:00:00+07	0910.64.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-08-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6687	\N	1	1067	\N	\N	2021-08-26 00:00:00+07	1067.61.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-08-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6692	\N	1	2528	\N	\N	2021-08-26 00:00:00+07	2528.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-08-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6702	\N	1	1071	\N	\N	2021-09-02 00:00:00+07	1071.61.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6711	\N	1	630	\N	\N	2021-09-08 00:00:00+07	0630.63.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6713	\N	1	1073	\N	\N	2021-09-09 00:00:00+07	1073.61.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6714	\N	1	2533	\N	\N	2021-09-09 00:00:00+07	2533.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6716	\N	1	2535	\N	\N	2021-09-09 00:00:00+07	2535.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6752	\N	1	2730	\N	\N	2021-09-15 00:00:00+07	2730.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6756	\N	1	1076	\N	\N	2021-09-16 00:00:00+07	1076.61.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6758	\N	1	2540	\N	\N	2021-09-16 00:00:00+07	2540.65.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6764	\N	1	2735	\N	\N	2021-09-17 00:00:00+07	2735.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6795	\N	1	637	\N	\N	2021-09-30 00:00:00+07	0637.63.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6797	\N	1	639	\N	\N	2021-09-30 00:00:00+07	0639.63.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-09-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6809	\N	1	2748	\N	\N	2021-10-07 00:00:00+07	2748.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-10-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6810	\N	1	2749	\N	\N	2021-10-07 00:00:00+07	2749.62.600	1	\N	\N	OA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-10-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6932	\N	1	1096	\N	\N	2021-11-23 00:00:00+07	1096.61.600	1	\N	\N	SA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6933	\N	1	1097	\N	\N	2021-11-23 00:00:00+07	1097.61.600	1	\N	\N	SA	\N	10205	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7407	\N	1	1168	\N	\N	2022-05-11 00:00:00+07	1168.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-05-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1537	\N	2	227	\N	\N	2012-01-05 00:00:00+07	0227.63.400	1	\N	2020-07-01 00:00:00+07	OA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1536	\N	2	1610	\N	\N	2012-01-05 00:00:00+07	1610.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1535	\N	2	1611	\N	\N	2012-01-05 00:00:00+07	1611.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-10-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6351	\N	1	2657	\N	\N	2021-01-19 00:00:00+07	2657.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-01-19 00:00:00+07	108	11:00	21:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1534	\N	2	1618	\N	\N	2012-01-03 00:00:00+07	1618.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1533	\N	2	657	\N	\N	2012-01-03 00:00:00+07	0657.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7285	\N	1	1002	\N	\N	2022-03-01 00:00:00+07	1002.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7367	\N	1	2875	\N	\N	2022-04-04 00:00:00+07	2875.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4420	\N	1	2235	\N	\N	2018-03-01 00:00:00+07	2235.62.400	1	\N	2019-06-09 00:00:00+07	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1532	\N	2	656	\N	\N	2012-01-03 00:00:00+07	0656.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1531	\N	2	1616	\N	\N	2012-01-03 00:00:00+07	1616.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1530	\N	2	1615	\N	\N	2012-01-03 00:00:00+07	1615.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1529	\N	2	1617	\N	\N	2012-01-03 00:00:00+07	1617.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7376	\N	1	2701	\N	\N	2022-04-07 00:00:00+07	2701.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1528	\N	2	1621	\N	\N	2012-01-03 00:00:00+07	1621.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1598	\N	2	748	\N	\N	2015-09-09 00:00:00+07	0748.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1527	\N	2	1620	\N	\N	2012-01-03 00:00:00+07	1620.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1526	\N	2	1614	\N	\N	2012-01-03 00:00:00+07	1614.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1525	\N	2	655	\N	\N	2012-01-03 00:00:00+07	0655.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7366	\N	1	2874	\N	\N	2022-04-04 00:00:00+07	2874.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7377	\N	1	2702	\N	\N	2022-04-07 00:00:00+07	2702.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1597	\N	2	960	\N	\N	2015-08-10 00:00:00+07	0960.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1604	\N	2	1562	\N	\N	2016-07-12 00:00:00+07	1562.62.400	2	2019-10-30 00:00:00+07	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1524	\N	2	611	\N	\N	2012-01-03 00:00:00+07	0611.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1523	\N	2	358	\N	\N	2012-01-03 00:00:00+07	0358.64.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1522	\N	2	357	\N	\N	2012-01-03 00:00:00+07	0357.64.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1596	\N	2	1287	\N	\N	2015-07-23 00:00:00+07	1287.65.400	2	2017-08-01 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1521	\N	2	653	\N	\N	2012-01-03 00:00:00+07	0653.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1520	\N	2	654	\N	\N	2012-01-03 00:00:00+07	0654.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1519	\N	2	609	\N	\N	2012-01-03 00:00:00+07	0609.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-11-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1518	\N	2	1612	\N	\N	2012-01-02 00:00:00+07	1612.62.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1517	\N	2	1548	\N	\N	2012-01-02 00:00:00+07	1548.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1594	\N	2	488	\N	\N	2015-07-01 00:00:00+07	0488.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4845	\N	1	891	\N	\N	2018-11-15 00:00:00+07	0891.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5665	\N	1	2253	\N	\N	2019-09-09 00:00:00+07	2253.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7429	\N	1	2722	\N	\N	2022-05-30 00:00:00+07	2722.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-05-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5666	\N	1	802	\N	\N	2019-09-09 00:00:00+07	0802.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1516	\N	2	1601	\N	\N	2011-11-21 00:00:00+07	1601.62.400	1	\N	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1515	\N	2	1594	\N	\N	2011-10-10 00:00:00+07	1594.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5667	\N	1	2254	\N	\N	2019-09-09 00:00:00+07	2254.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5668	\N	1	977	\N	\N	2019-09-09 00:00:00+07	0977.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5669	\N	1	2482	\N	\N	2019-09-09 00:00:00+07	2482.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4402	\N	1	1691	\N	\N	2018-02-14 00:00:00+07	1691.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1514	\N	2	643	\N	\N	2011-10-10 00:00:00+07	0643.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1593	\N	2	1806	\N	\N	2015-04-02 00:00:00+07	1806.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1592	\N	2	406	\N	\N	2015-02-17 00:00:00+07	0406.64.400	2	2017-06-17 00:00:00+07	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1603	\N	2	920	\N	\N	2016-02-11 00:00:00+07	0920.62.400	2	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1513	\N	2	642	\N	\N	2011-09-19 00:00:00+07	0642.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6532	\N	1	907	\N	\N	2021-07-02 00:00:00+07	0907.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-07-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5337	\N	1	489	\N	\N	2019-05-22 00:00:00+07	0489.63.400	2	2021-09-01 00:00:00+07	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-09-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1558	\N	2	1641	\N	\N	2012-04-23 00:00:00+07	1641.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1557	\N	2	618	\N	\N	2012-04-02 00:00:00+07	0618.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6260	\N	1	588	\N	\N	2020-11-04 00:00:00+07	0588.63.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2020-11-04 00:00:00+07	196	10:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1600	\N	2	489	\N	\N	2016-02-10 00:00:00+07	0489.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5204	\N	1	2372	\N	\N	2019-04-10 00:00:00+07	2372.62.400	1	\N	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6138	\N	1	857	\N	\N	2020-07-14 00:00:00+07	0857.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1556	\N	2	617	\N	\N	2012-04-02 00:00:00+07	0617.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1599	\N	2	337	\N	\N	2016-01-28 00:00:00+07	0337.62.400	1	\N	2020-06-01 00:00:00+07	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3929	\N	2	567	\N	\N	2017-08-02 00:00:00+07	0567.64.400	1	\N	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6946	\N	1	1098	\N	\N	2021-11-30 00:00:00+07	1098.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1555	\N	2	615	\N	\N	2012-03-06 00:00:00+07	0615.65.400	2	2017-05-24 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1554	\N	2	660	\N	\N	2012-03-02 00:00:00+07	0660.61.400	1	\N	\N	SA	\N	10189	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1553	\N	2	1313	\N	\N	2012-03-01 00:00:00+07	1313.62.400	2	2018-09-01 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7302	\N	1	1004	\N	\N	2022-03-10 00:00:00+07	1004.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1552	\N	2	632	\N	\N	2012-02-21 00:00:00+07	0632.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6679	\N	1	2526	\N	\N	2021-08-25 00:00:00+07	2526.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-08-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6680	\N	1	2719	\N	\N	2021-08-25 00:00:00+07	2719.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-08-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6974	\N	1	1103	\N	\N	2021-12-09 00:00:00+07	1103.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-12-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6975	\N	1	2781	\N	\N	2021-12-09 00:00:00+07	2781.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-12-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1551	\N	2	659	\N	\N	2012-02-20 00:00:00+07	0659.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7394	\N	1	2712	\N	\N	2022-04-18 00:00:00+07	2712.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3886	\N	2	1639	\N	\N	2017-07-14 00:00:00+07	1639.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1550	\N	2	359	\N	\N	2012-02-10 00:00:00+07	0359.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5644	\N	1	796	\N	\N	2019-08-26 00:00:00+07	0796.64.400	1	\N	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7393	\N	1	2711	\N	\N	2022-04-18 00:00:00+07	2711.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7392	\N	1	2710	\N	\N	2022-04-18 00:00:00+07	2710.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7290	\N	1	2680	\N	\N	2022-03-04 00:00:00+07	2680.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1548	\N	2	658	\N	\N	2012-01-17 00:00:00+07	0658.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1547	\N	2	612	\N	\N	2012-01-17 00:00:00+07	0612.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7391	\N	1	2709	\N	\N	2022-04-18 00:00:00+07	2709.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7390	\N	1	2708	\N	\N	2022-04-18 00:00:00+07	2708.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1544	\N	2	665	\N	\N	2012-01-17 00:00:00+07	0665.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6176	\N	1	863	\N	\N	2020-08-28 00:00:00+07	0863.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1538	\N	2	353	\N	\N	2012-01-10 00:00:00+07	0353.64.400	2	2019-08-13 00:00:00+07	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3536	\N	1	2135	\N	\N	2017-04-03 00:00:00+07	2135.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1512	\N	2	1591	\N	\N	2011-08-12 00:00:00+07	1591.62.400	1	\N	2020-06-01 00:00:00+07	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1511	\N	2	651	\N	\N	2011-07-28 00:00:00+07	0651.61.400	1	\N	2018-07-01 00:00:00+07	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1510	\N	2	1588	\N	\N	2011-07-27 00:00:00+07	1588.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1509	\N	2	1586	\N	\N	2011-07-26 00:00:00+07	1586.62.400	1	\N	2020-05-01 00:00:00+07	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1591	\N	2	783	\N	\N	2014-12-05 00:00:00+07	0783.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5678	\N	1	804	\N	\N	2019-09-11 00:00:00+07	0804.64.400	1	\N	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1508	\N	2	1582	\N	\N	2011-07-07 00:00:00+07	1582.62.400	1	\N	2020-06-01 00:00:00+07	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-07-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1590	\N	2	1685	\N	\N	2014-10-29 00:00:00+07	1685.62.400	2	2020-03-01 00:00:00+07	\N	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4670	\N	1	430	\N	\N	2018-08-24 00:00:00+07	0430.63.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1507	\N	2	218	\N	\N	2011-06-24 00:00:00+07	0218.63.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1506	\N	2	1574	\N	\N	2011-06-09 00:00:00+07	1574.62.400	2	2018-05-01 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7379	\N	1	1164	\N	\N	2022-04-07 00:00:00+07	1164.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1589	\N	2	1768	\N	\N	2014-08-20 00:00:00+07	1768.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7380	\N	1	1165	\N	\N	2022-04-07 00:00:00+07	1165.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1505	\N	2	1572	\N	\N	2011-06-06 00:00:00+07	1572.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1504	\N	2	633	\N	\N	2011-04-20 00:00:00+07	0633.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7430	\N	1	2723	\N	\N	2022-05-30 00:00:00+07	2723.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-05-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7381	\N	1	1166	\N	\N	2022-04-07 00:00:00+07	1166.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1503	\N	2	1565	\N	\N	2011-04-15 00:00:00+07	1565.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7171	\N	1	2823	\N	\N	2022-01-25 00:00:00+07	2823.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4346	\N	2	610	\N	\N	2018-01-11 00:00:00+07	0610.64.400	2	2019-10-01 00:00:00+07	\N	OA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-02-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1502	\N	2	1564	\N	\N	2011-04-14 00:00:00+07	1564.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4342	\N	2	395	\N	\N	2018-01-11 00:00:00+07	0395.63.400	1	\N	\N	OA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1587	\N	2	704	\N	\N	2014-06-13 00:00:00+07	0704.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5992	\N	1	2582	\N	\N	2020-01-09 00:00:00+07	2582.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-12-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1586	\N	2	96	\N	\N	2014-03-21 00:00:00+07	0096.63.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1501	\N	2	1563	\N	\N	2011-04-13 00:00:00+07	1563.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1499	\N	2	1571	\N	\N	2011-04-02 00:00:00+07	1571.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1498	\N	2	1570	\N	\N	2011-04-02 00:00:00+07	1570.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1585	\N	2	363	\N	\N	2013-11-25 00:00:00+07	0363.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7385	\N	1	701	\N	\N	2022-04-18 00:00:00+07	0701.63.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1496	\N	2	213	\N	\N	2011-04-01 00:00:00+07	0213.63.400	2	2018-02-26 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7387	\N	1	2705	\N	\N	2022-04-18 00:00:00+07	2705.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1495	\N	2	1547	\N	\N	2011-03-07 00:00:00+07	1547.62.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1493	\N	2	623	\N	\N	2011-03-07 00:00:00+07	0623.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1492	\N	2	1545	\N	\N	2011-03-03 00:00:00+07	1545.62.400	2	2018-01-01 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7388	\N	1	2706	\N	\N	2022-04-18 00:00:00+07	2706.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1491	\N	2	506	\N	\N	2011-02-17 00:00:00+07	0506.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1490	\N	2	456	\N	\N	2011-02-17 00:00:00+07	0456.65.400	1	\N	\N	OA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1489	\N	2	198	\N	\N	2011-02-17 00:00:00+07	0198.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1488	\N	2	147	\N	\N	2011-02-17 00:00:00+07	0147.65.400	1	2018-04-03 00:00:00+07	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-02-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1487	\N	2	67	\N	\N	2011-02-17 00:00:00+07	0067.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1486	\N	2	66	\N	\N	2011-02-17 00:00:00+07	0066.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-11-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1485	\N	2	65	\N	\N	2011-02-17 00:00:00+07	0065.65.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1484	\N	2	89	\N	\N	2011-02-17 00:00:00+07	0089.64.400	2	2018-04-01 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1584	\N	2	683	\N	\N	2013-10-31 00:00:00+07	0683.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1483	\N	2	87	\N	\N	2011-02-17 00:00:00+07	0087.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7389	\N	1	2707	\N	\N	2022-04-18 00:00:00+07	2707.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-04-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1482	\N	2	81	\N	\N	2011-02-17 00:00:00+07	0081.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1481	\N	2	40	\N	\N	2011-02-17 00:00:00+07	0040.64.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1583	\N	2	1682	\N	\N	2013-09-16 00:00:00+07	1682.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1480	\N	2	39	\N	\N	2011-02-17 00:00:00+07	0039.64.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6029	\N	1	2378	\N	\N	2020-02-10 00:00:00+07	2378.65.400	1	\N	\N	SA	\N	10189	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1582	\N	2	644	\N	\N	2013-05-01 00:00:00+07	0644.65.400	1	\N	\N	OA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1581	\N	2	371	\N	\N	2013-04-22 00:00:00+07	0371.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1580	\N	2	233	\N	\N	2013-04-08 00:00:00+07	0233.63.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1479	\N	2	38	\N	\N	2011-02-17 00:00:00+07	0038.64.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1478	\N	2	37	\N	\N	2011-02-17 00:00:00+07	0037.64.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1578	\N	2	1658	\N	\N	2013-02-21 00:00:00+07	1658.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1477	\N	2	153	\N	\N	2011-02-17 00:00:00+07	0153.63.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1475	\N	2	28	\N	\N	2011-02-17 00:00:00+07	0028.63.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1577	\N	2	1664	\N	\N	2013-02-13 00:00:00+07	1664.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1576	\N	2	673	\N	\N	2013-02-08 00:00:00+07	0673.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6835	\N	1	1086	\N	\N	2021-10-21 00:00:00+07	1086.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-10-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6836	\N	1	2758	\N	\N	2021-10-21 00:00:00+07	2758.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-10-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1575	\N	2	368	\N	\N	2012-12-10 00:00:00+07	0368.64.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1574	\N	2	632	\N	\N	2012-11-01 00:00:00+07	0632.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6032	\N	1	2596	\N	\N	2020-02-13 00:00:00+07	2596.62.400	1	\N	\N	OA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1573	\N	2	365	\N	\N	2012-10-09 00:00:00+07	0365.64.400	1	2017-12-01 00:00:00+07	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-08-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1473	\N	2	1309	\N	\N	2011-02-17 00:00:00+07	1309.62.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1472	\N	2	1298	\N	\N	2011-02-17 00:00:00+07	1298.62.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1571	\N	2	1632	\N	\N	2012-09-03 00:00:00+07	1632.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1470	\N	2	1017	\N	\N	2011-02-17 00:00:00+07	1017.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-06-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1467	\N	2	741	\N	\N	2011-02-17 00:00:00+07	0741.62.400	2	2018-04-03 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6846	\N	1	1088	\N	\N	2021-10-25 00:00:00+07	1088.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-10-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6449	\N	1	2494	\N	\N	2021-04-20 00:00:00+07	2494.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-04-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4471	\N	1	1711	\N	\N	2018-03-26 00:00:00+07	1711.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7335	\N	1	1159	\N	\N	2022-03-21 00:00:00+07	1159.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1466	\N	2	455	\N	\N	2011-02-17 00:00:00+07	0455.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1465	\N	2	453	\N	\N	2011-02-17 00:00:00+07	0453.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1464	\N	2	452	\N	\N	2011-02-17 00:00:00+07	0452.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1570	\N	2	625	\N	\N	2012-08-16 00:00:00+07	0625.65.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7328	\N	1	2866	\N	\N	2022-03-16 00:00:00+07	2866.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5881	\N	1	538	\N	\N	2019-11-05 00:00:00+07	0538.63.400	1	\N	\N	SA	\N	10189	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1463	\N	2	451	\N	\N	2011-02-17 00:00:00+07	0451.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1462	\N	2	399	\N	\N	2011-02-17 00:00:00+07	0399.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7327	\N	1	2865	\N	\N	2022-03-16 00:00:00+07	2865.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7326	\N	1	2690	\N	\N	2022-03-16 00:00:00+07	2690.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1461	\N	2	398	\N	\N	2011-02-17 00:00:00+07	0398.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1569	\N	2	628	\N	\N	2012-08-08 00:00:00+07	0628.65.400	1	\N	\N	OA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1460	\N	2	397	\N	\N	2011-02-17 00:00:00+07	0397.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6053	\N	1	1016	\N	\N	2020-03-10 00:00:00+07	1016.61.400	2	2020-03-11 00:00:00+07	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2020-09-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1459	\N	2	396	\N	\N	2011-02-17 00:00:00+07	0396.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1568	\N	2	627	\N	\N	2012-08-08 00:00:00+07	0627.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1567	\N	2	230	\N	\N	2012-07-04 00:00:00+07	0230.63.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6470	\N	1	892	\N	\N	2021-05-10 00:00:00+07	0892.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-05-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1566	\N	2	229	\N	\N	2012-07-02 00:00:00+07	0229.63.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1458	\N	2	395	\N	\N	2011-02-17 00:00:00+07	0395.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1602	\N	2	1272	\N	\N	2016-02-10 00:00:00+07	1272.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1457	\N	2	394	\N	\N	2011-02-17 00:00:00+07	0394.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7354	\N	1	2697	\N	\N	2022-03-28 00:00:00+07	2697.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7318	\N	1	2688	\N	\N	2022-03-11 00:00:00+07	2688.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1565	\N	2	221	\N	\N	2012-06-01 00:00:00+07	0221.63.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7317	\N	1	2687	\N	\N	2022-03-11 00:00:00+07	2687.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7316	\N	1	2686	\N	\N	2022-03-11 00:00:00+07	2686.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7315	\N	1	2685	\N	\N	2022-03-11 00:00:00+07	2685.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7353	\N	1	2696	\N	\N	2022-03-28 00:00:00+07	2696.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1456	\N	2	393	\N	\N	2011-02-17 00:00:00+07	0393.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1601	\N	2	1955	\N	\N	2016-02-10 00:00:00+07	1955.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1455	\N	2	392	\N	\N	2011-02-17 00:00:00+07	0392.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1454	\N	2	391	\N	\N	2011-02-17 00:00:00+07	0391.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1453	\N	2	390	\N	\N	2011-02-17 00:00:00+07	0390.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1452	\N	2	389	\N	\N	2011-02-17 00:00:00+07	0389.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1451	\N	2	388	\N	\N	2011-02-17 00:00:00+07	0388.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1450	\N	2	387	\N	\N	2011-02-17 00:00:00+07	0387.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1449	\N	2	489	\N	\N	2011-02-17 00:00:00+07	0489.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6076	\N	1	845	\N	\N	2020-04-28 00:00:00+07	0845.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7314	\N	1	1157	\N	\N	2022-03-11 00:00:00+07	1157.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6894	\N	1	2590	\N	\N	2021-11-10 00:00:00+07	2590.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7313	\N	1	1156	\N	\N	2022-03-11 00:00:00+07	1156.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1448	\N	2	487	\N	\N	2011-02-17 00:00:00+07	0487.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6092	\N	1	848	\N	\N	2020-06-09 00:00:00+07	0848.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1447	\N	2	467	\N	\N	2011-02-17 00:00:00+07	0467.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1446	\N	2	118	\N	\N	2011-02-17 00:00:00+07	0118.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1445	\N	2	86	\N	\N	2011-02-17 00:00:00+07	0086.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1444	\N	2	85	\N	\N	2011-02-17 00:00:00+07	0085.61.400	2	2020-05-01 00:00:00+07	\N	SA	\N	10189	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1443	\N	2	84	\N	\N	2011-02-17 00:00:00+07	0084.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-03-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7341	\N	1	695	\N	\N	2022-03-22 00:00:00+07	0695.63.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7340	\N	1	2694	\N	\N	2022-03-22 00:00:00+07	2694.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7312	\N	1	1155	\N	\N	2022-03-11 00:00:00+07	1155.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7311	\N	1	1154	\N	\N	2022-03-11 00:00:00+07	1154.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1442	\N	2	83	\N	\N	2011-02-17 00:00:00+07	0083.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7337	\N	1	1007	\N	\N	2022-03-21 00:00:00+07	1007.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1441	\N	2	386	\N	\N	2011-02-17 00:00:00+07	0386.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6510	\N	1	898	\N	\N	2021-06-23 00:00:00+07	0898.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-06-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7309	\N	1	1153	\N	\N	2022-03-11 00:00:00+07	1153.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7308	\N	1	1152	\N	\N	2022-03-11 00:00:00+07	1152.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7307	\N	1	1151	\N	\N	2022-03-11 00:00:00+07	1151.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4002	\N	1	2195	\N	\N	2017-09-13 00:00:00+07	2195.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7336	\N	1	2692	\N	\N	2022-03-21 00:00:00+07	2692.65.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7306	\N	1	1150	\N	\N	2022-03-11 00:00:00+07	1150.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7305	\N	1	1149	\N	\N	2022-03-11 00:00:00+07	1149.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7304	\N	1	1006	\N	\N	2022-03-10 00:00:00+07	1006.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1563	\N	2	1638	\N	\N	2012-05-01 00:00:00+07	1638.62.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1562	\N	2	632	\N	\N	2012-05-01 00:00:00+07	0632.62.400	2	2018-01-31 00:00:00+07	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6198	\N	2	2627	\N	\N	2020-09-11 00:00:00+07	2627.62.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2020-11-09 00:00:00+07	\N	00:00	23:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1561	\N	2	1637	\N	\N	2012-05-01 00:00:00+07	1637.62.400	1	\N	\N	OA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4433	\N	1	635	\N	\N	2018-03-09 00:00:00+07	0635.64.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7303	\N	1	1005	\N	\N	2022-03-10 00:00:00+07	1005.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-03-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6113	\N	1	855	\N	\N	2020-07-02 00:00:00+07	0855.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1560	\N	2	662	\N	\N	2012-05-01 00:00:00+07	0662.61.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6529	\N	1	905	\N	\N	2021-07-02 00:00:00+07	0905.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-07-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6530	\N	1	906	\N	\N	2021-07-02 00:00:00+07	0906.64.400	1	\N	\N	SA	\N	10189	Amalia Mahardani	\N	\N	f	f	2021-07-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1559	\N	2	666	\N	\N	2012-05-01 00:00:00+07	0666.61.400	1	\N	\N	SA	\N	10189	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2921	\N	1	435	\N	\N	2016-12-01 00:00:00+07	0435.63.307	1	\N	2019-06-07 00:00:00+07	SA	\N	10211	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
487	\N	2	822	\N	\N	2014-09-17 00:00:00+07	0822.65.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
477	\N	2	375	\N	\N	2014-02-06 00:00:00+07	0375.64.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
476	\N	2	687	\N	\N	2014-02-04 00:00:00+07	0687.61.307	1	\N	2019-12-01 00:00:00+07	SA	\N	10211	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
475	\N	2	374	\N	\N	2014-01-30 00:00:00+07	0374.64.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
473	\N	2	1695	\N	\N	2013-10-17 00:00:00+07	1695.62.307	1	\N	2020-06-01 00:00:00+07	SA	\N	10211	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
471	\N	2	1678	\N	\N	2013-06-19 00:00:00+07	1678.62.307	1	\N	2020-06-01 00:00:00+07	SA	\N	10211	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
469	\N	2	1662	\N	\N	2013-05-01 00:00:00+07	1662.62.307	2	2017-01-31 00:00:00+07	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
460	\N	2	192	\N	\N	2011-02-11 00:00:00+07	0192.64.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3403	\N	1	532	\N	\N	2017-03-01 00:00:00+07	0532.64.307	1	\N	2019-06-09 00:00:00+07	SA	\N	10211	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4421	\N	1	2236	\N	\N	2018-03-01 00:00:00+07	2236.62.307	1	\N	2019-06-09 00:00:00+07	SA	\N	10211	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
489	\N	2	1845	\N	\N	2015-03-30 00:00:00+07	1845.62.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5080	\N	1	2015	\N	\N	2019-03-11 00:00:00+07	2015.65.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2922	\N	1	1977	\N	\N	2016-12-01 00:00:00+07	1977.65.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
457	\N	2	1131	\N	\N	2011-02-11 00:00:00+07	1131.65.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
485	\N	2	242	\N	\N	2014-06-19 00:00:00+07	0242.63.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
458	\N	2	384	\N	\N	2011-02-11 00:00:00+07	0384.65.307	1	\N	2019-06-07 00:00:00+07	SA	\N	10211	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2920	\N	1	783	\N	\N	2016-12-01 00:00:00+07	0783.61.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
499	\N	2	571	\N	\N	2016-08-10 00:00:00+07	0571.65.307	1	\N	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
482	\N	2	723	\N	\N	2014-04-14 00:00:00+07	0723.65.307	2	2017-03-27 00:00:00+07	\N	SA	\N	10211	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
481	\N	2	648	\N	\N	2014-03-28 00:00:00+07	0648.65.307	2	2018-03-01 00:00:00+07	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5547	\N	1	2455	\N	\N	2019-07-24 00:00:00+07	2455.62.307	1	\N	\N	SA	\N	10211	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
493	\N	2	1426	\N	\N	2016-03-01 00:00:00+07	1426.65.307	1	2018-01-31 00:00:00+07	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
479	\N	2	1704	\N	\N	2014-02-25 00:00:00+07	1704.62.307	2	2017-01-31 00:00:00+07	\N	SA	\N	10211	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
478	\N	2	686	\N	\N	2014-02-14 00:00:00+07	0686.65.307	1	2017-06-26 00:00:00+07	2019-06-09 00:00:00+07	SA	\N	10211	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
459	\N	2	41	\N	\N	2011-02-11 00:00:00+07	0041.64.307	2	2020-02-01 00:00:00+07	\N	SA	\N	10210	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
498	\N	2	1595	\N	\N	2016-06-06 00:00:00+07	1595.62.307	2	2017-12-27 00:00:00+07	\N	SA	\N	10210	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
483	\N	2	1689	\N	\N	2014-04-21 00:00:00+07	1689.62.307	1	\N	\N	SA	\N	10210	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7319	\N	1	2860	\N	\N	2022-03-15 00:00:00+07	2860.62.307	1	\N	\N	SA	\N	10210	Amalia Mahardani	\N	\N	f	f	2022-03-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
480	\N	2	1703	\N	\N	2014-03-11 00:00:00+07	1703.62.307	1	\N	\N	SA	\N	10210	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
468	\N	2	823	\N	\N	2013-01-16 00:00:00+07	0823.65.307	2	2018-10-11 00:00:00+07	\N	SA	\N	10210	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5549	\N	1	2457	\N	\N	2019-07-24 00:00:00+07	2457.62.307	1	\N	\N	SA	\N	10210	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
461	\N	2	455	\N	\N	2011-02-11 00:00:00+07	0455.65.307	1	\N	\N	SA	\N	10210	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
500	\N	2	354	\N	\N	2016-09-08 00:00:00+07	0354.64.307	2	2019-09-01 00:00:00+07	\N	SA	\N	10210	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4166	\N	1	1665	\N	\N	2017-10-12 00:00:00+07	1665.65.307	1	\N	\N	SA	\N	10210	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4564	\N	1	880	\N	\N	2018-05-15 00:00:00+07	0880.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5012	\N	1	2336	\N	\N	2019-02-21 00:00:00+07	2336.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6115	\N	1	568	\N	\N	2020-07-02 00:00:00+07	0568.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6116	\N	1	569	\N	\N	2020-07-02 00:00:00+07	0569.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6117	\N	1	570	\N	\N	2020-07-03 00:00:00+07	0570.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4919	\N	1	894	\N	\N	2018-12-28 00:00:00+07	0894.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5014	\N	1	2337	\N	\N	2019-02-21 00:00:00+07	2337.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6943	\N	1	657	\N	\N	2021-11-25 00:00:00+07	0657.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-11-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4045	\N	1	1656	\N	\N	2017-09-18 00:00:00+07	1656.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5108	\N	1	1560	\N	\N	2019-03-15 00:00:00+07	1560.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5112	\N	1	2356	\N	\N	2019-03-18 00:00:00+07	2356.62.101	1	\N	2019-05-01 00:00:00+07	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5118	\N	1	2037	\N	\N	2019-03-18 00:00:00+07	2037.65.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5203	\N	1	925	\N	\N	2019-04-10 00:00:00+07	0925.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5164	\N	1	920	\N	\N	2019-04-01 00:00:00+07	0920.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6136	\N	1	2406	\N	\N	2020-07-14 00:00:00+07	2406.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6137	\N	1	2407	\N	\N	2020-07-14 00:00:00+07	2407.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4536	\N	1	664	\N	\N	2018-04-27 00:00:00+07	0664.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4534	\N	1	878	\N	\N	2018-04-26 00:00:00+07	0878.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4041	\N	1	574	\N	\N	2017-09-15 00:00:00+07	0574.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4914	\N	1	454	\N	\N	2018-12-18 00:00:00+07	0454.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6144	\N	1	2411	\N	\N	2020-07-16 00:00:00+07	2411.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6145	\N	1	574	\N	\N	2020-07-20 00:00:00+07	0574.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6964	\N	1	663	\N	\N	2021-12-06 00:00:00+07	0663.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-12-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6146	\N	1	575	\N	\N	2020-07-20 00:00:00+07	0575.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6148	\N	1	858	\N	\N	2020-07-20 00:00:00+07	0858.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3935	\N	2	568	\N	\N	2017-08-11 00:00:00+07	0568.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6947	\N	1	938	\N	\N	2021-11-30 00:00:00+07	0938.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3924	\N	1	18	\N	\N	2016-11-01 00:00:00+07	0018.00.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6969	\N	1	2617	\N	\N	2021-12-08 00:00:00+07	2617.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-12-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3922	\N	1	851	\N	\N	2017-07-31 00:00:00+07	0851.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6153	\N	1	576	\N	\N	2020-07-29 00:00:00+07	0576.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5291	\N	1	2103	\N	\N	2019-05-14 00:00:00+07	2103.65.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5239	\N	1	2088	\N	\N	2019-04-29 00:00:00+07	2088.65.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4909	\N	1	735	\N	\N	2018-12-14 00:00:00+07	0735.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4515	\N	1	876	\N	\N	2018-04-17 00:00:00+07	0876.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7408	\N	1	707	\N	\N	2022-05-11 00:00:00+07	0707.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-05-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4913	\N	1	893	\N	\N	2018-12-18 00:00:00+07	0893.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6978	\N	1	2783	\N	\N	2021-12-09 00:00:00+07	2783.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-12-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6979	\N	1	1105	\N	\N	2021-12-10 00:00:00+07	1105.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-12-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4514	\N	1	412	\N	\N	2018-04-17 00:00:00+07	0412.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7073	\N	1	965	\N	\N	2022-01-11 00:00:00+07	0965.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4513	\N	1	875	\N	\N	2018-04-17 00:00:00+07	0875.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6166	\N	1	861	\N	\N	2020-08-14 00:00:00+07	0861.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6990	\N	1	2622	\N	\N	2021-12-15 00:00:00+07	2622.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-12-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5560	\N	1	2464	\N	\N	2019-07-30 00:00:00+07	2464.62.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6170	\N	1	578	\N	\N	2020-08-24 00:00:00+07	0578.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6171	\N	1	579	\N	\N	2020-08-24 00:00:00+07	0579.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6174	\N	1	580	\N	\N	2020-08-26 00:00:00+07	0580.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6175	\N	1	2421	\N	\N	2020-08-27 00:00:00+07	2421.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7044	\N	1	1112	\N	\N	2022-01-05 00:00:00+07	1112.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6221	\N	1	582	\N	\N	2020-09-21 00:00:00+07	0582.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-09-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6222	\N	1	1032	\N	\N	2020-09-22 00:00:00+07	1032.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-09-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6223	\N	1	2431	\N	\N	2020-09-23 00:00:00+07	2431.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-09-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7128	\N	1	679	\N	\N	2022-01-18 00:00:00+07	0679.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7127	\N	1	1117	\N	\N	2022-01-17 00:00:00+07	1117.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3828	\N	2	86	\N	\N	2017-06-15 00:00:00+07	0086.00.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3825	\N	1	85	\N	\N	2017-06-14 00:00:00+07	0085.00.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7126	\N	1	2809	\N	\N	2022-01-17 00:00:00+07	2809.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3934	\N	1	1508	\N	\N	2016-08-22 00:00:00+07	1508.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2565	\N	2	260	\N	\N	2015-05-07 00:00:00+07	0260.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3824	\N	1	557	\N	\N	2017-06-14 00:00:00+07	0557.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3933	\N	1	852	\N	\N	2017-08-09 00:00:00+07	0852.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3822	\N	1	1633	\N	\N	2017-06-14 00:00:00+07	1633.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4912	\N	1	736	\N	\N	2018-12-18 00:00:00+07	0736.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6238	\N	1	585	\N	\N	2020-10-13 00:00:00+07	0585.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-10-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3976	\N	1	1971	\N	\N	2017-09-07 00:00:00+07	1971.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5855	\N	1	990	\N	\N	2019-10-24 00:00:00+07	0990.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3969	\N	1	572	\N	\N	2017-08-31 00:00:00+07	0572.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3967	\N	2	853	\N	\N	2017-08-30 00:00:00+07	0853.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7118	\N	1	2805	\N	\N	2022-01-14 00:00:00+07	2805.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3806	\N	1	849	\N	\N	2017-06-12 00:00:00+07	0849.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7111	\N	1	968	\N	\N	2022-01-14 00:00:00+07	0968.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7138	\N	1	2811	\N	\N	2022-01-19 00:00:00+07	2811.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7110	\N	1	2644	\N	\N	2022-01-14 00:00:00+07	2644.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7139	\N	1	2811	\N	\N	2022-01-19 00:00:00+07	2811.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3760	\N	2	74	\N	\N	2017-06-05 00:00:00+07	0074.00.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7136	\N	1	1118	\N	\N	2022-01-19 00:00:00+07	1118.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6259	\N	1	1033	\N	\N	2020-11-03 00:00:00+07	1033.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-11-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6261	\N	1	1034	\N	\N	2020-11-05 00:00:00+07	1034.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-11-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6262	\N	1	1035	\N	\N	2020-11-05 00:00:00+07	1035.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-11-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6264	\N	1	589	\N	\N	2020-11-10 00:00:00+07	0589.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6271	\N	1	590	\N	\N	2020-11-16 00:00:00+07	0590.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-11-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6276	\N	1	1036	\N	\N	2020-11-20 00:00:00+07	1036.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-11-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7132	\N	1	2810	\N	\N	2022-01-18 00:00:00+07	2810.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7130	\N	1	973	\N	\N	2022-01-18 00:00:00+07	0973.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7350	\N	1	697	\N	\N	2022-03-24 00:00:00+07	0697.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-03-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6327	\N	1	598	\N	\N	2020-12-11 00:00:00+07	0598.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-12-11 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6333	\N	1	1313	\N	\N	2020-12-30 00:00:00+07	1313.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-12-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6334	\N	1	2466	\N	\N	2020-12-30 00:00:00+07	2466.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-12-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6336	\N	1	872	\N	\N	2021-01-07 00:00:00+07	0872.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-01-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6341	\N	1	1042	\N	\N	2021-01-12 00:00:00+07	1042.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6354	\N	1	2659	\N	\N	2021-01-20 00:00:00+07	2659.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-01-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6359	\N	1	876	\N	\N	2021-01-26 00:00:00+07	0876.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-01-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6365	\N	1	2662	\N	\N	2021-01-29 00:00:00+07	2662.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-01-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6366	\N	1	2663	\N	\N	2021-01-29 00:00:00+07	2663.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-01-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6367	\N	1	2664	\N	\N	2021-02-01 00:00:00+07	2664.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-02-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6371	\N	1	2475	\N	\N	2021-02-03 00:00:00+07	2475.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-02-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6377	\N	1	2670	\N	\N	2021-02-05 00:00:00+07	2670.62.101	2	2021-12-23 00:00:00+07	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-02-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6384	\N	1	2673	\N	\N	2021-02-17 00:00:00+07	2673.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-02-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4939	\N	1	899	\N	\N	2019-01-29 00:00:00+07	0899.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6385	\N	1	2674	\N	\N	2021-02-18 00:00:00+07	2674.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-02-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6391	\N	1	2477	\N	\N	2021-03-04 00:00:00+07	2477.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-03-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6393	\N	1	604	\N	\N	2021-03-04 00:00:00+07	0604.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-03-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6396	\N	1	881	\N	\N	2021-03-08 00:00:00+07	0881.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-03-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6402	\N	1	1050	\N	\N	2021-03-16 00:00:00+07	1050.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-03-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6403	\N	1	2679	\N	\N	2021-03-17 00:00:00+07	2679.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-03-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6407	\N	1	2479	\N	\N	2021-03-23 00:00:00+07	2479.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-03-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6421	\N	1	885	\N	\N	2021-04-07 00:00:00+07	0885.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-04-07 00:00:00+07	\N	08:00	20:00	20	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6425	\N	1	1052	\N	\N	2021-04-09 00:00:00+07	1052.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-04-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6429	\N	1	886	\N	\N	2021-04-13 00:00:00+07	0886.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-04-13 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6430	\N	1	887	\N	\N	2021-04-14 00:00:00+07	0887.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-04-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4491	\N	1	410	\N	\N	2018-04-06 00:00:00+07	0410.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4490	\N	1	409	\N	\N	2018-04-06 00:00:00+07	0409.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7431	\N	1	2898	\N	\N	2022-05-30 00:00:00+07	2898.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-05-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4838	\N	1	890	\N	\N	2018-11-05 00:00:00+07	0890.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7081	\N	1	674	\N	\N	2022-01-12 00:00:00+07	0674.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7447	\N	1	712	\N	\N	2022-06-07 00:00:00+07	0712.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-06-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4473	\N	1	651	\N	\N	2018-03-26 00:00:00+07	0651.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6446	\N	1	889	\N	\N	2021-04-20 00:00:00+07	0889.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-04-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6457	\N	1	2501	\N	\N	2021-04-27 00:00:00+07	2501.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-04-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6465	\N	1	608	\N	\N	2021-05-05 00:00:00+07	0608.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-05-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6468	\N	1	2699	\N	\N	2021-05-06 00:00:00+07	2699.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-05-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5884	\N	1	539	\N	\N	2019-11-06 00:00:00+07	0539.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6471	\N	1	1056	\N	\N	2021-05-10 00:00:00+07	1056.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-05-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6473	\N	1	2701	\N	\N	2021-05-11 00:00:00+07	2701.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-05-11 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5885	\N	1	540	\N	\N	2019-11-06 00:00:00+07	0540.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6481	\N	1	2704	\N	\N	2021-05-25 00:00:00+07	2704.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-05-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6482	\N	1	609	\N	\N	2021-05-31 00:00:00+07	0609.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-05-31 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6484	\N	1	610	\N	\N	2021-06-02 00:00:00+07	0610.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-06-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6502	\N	1	896	\N	\N	2021-06-16 00:00:00+07	0896.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-06-16 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6503	\N	1	1058	\N	\N	2021-06-17 00:00:00+07	1058.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-06-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7310	\N	1	2859	\N	\N	2022-03-11 00:00:00+07	2859.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-03-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6516	\N	1	617	\N	\N	2021-06-25 00:00:00+07	0617.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-06-25 00:00:00+07	\N	00:00	00:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4832	\N	1	889	\N	\N	2018-11-01 00:00:00+07	0889.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4831	\N	1	2295	\N	\N	2018-10-30 00:00:00+07	2295.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3733	\N	1	546	\N	\N	2017-05-23 00:00:00+07	0546.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5905	\N	1	822	\N	\N	2019-11-13 00:00:00+07	0822.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4432	\N	1	870	\N	\N	2018-03-09 00:00:00+07	0870.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5909	\N	1	993	\N	\N	2019-11-13 00:00:00+07	0993.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4827	\N	1	440	\N	\N	2018-10-25 00:00:00+07	0440.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6536	\N	1	620	\N	\N	2021-07-08 00:00:00+07	0620.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-07-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6538	\N	1	2521	\N	\N	2021-07-29 00:00:00+07	2521.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-07-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6540	\N	1	2522	\N	\N	2021-08-09 00:00:00+07	2522.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-08-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4426	\N	1	634	\N	\N	2018-03-07 00:00:00+07	0634.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3728	\N	1	839	\N	\N	2017-05-19 00:00:00+07	0839.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3727	\N	1	840	\N	\N	2017-05-19 00:00:00+07	0840.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7474	\N	1	2905	\N	\N	2022-06-16 00:00:00+07	2905.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5914	\N	1	825	\N	\N	2019-11-14 00:00:00+07	0825.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5916	\N	1	826	\N	\N	2019-11-18 00:00:00+07	0826.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7299	\N	1	2858	\N	\N	2022-03-10 00:00:00+07	2858.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-03-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5918	\N	1	2346	\N	\N	2019-11-18 00:00:00+07	2346.65.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4332	\N	1	393	\N	\N	2017-12-28 00:00:00+07	0393.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7298	\N	1	2857	\N	\N	2022-03-10 00:00:00+07	2857.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-03-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7288	\N	1	2854	\N	\N	2022-03-04 00:00:00+07	2854.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-03-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5923	\N	1	827	\N	\N	2019-11-20 00:00:00+07	0827.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4423	\N	1	633	\N	\N	2018-03-05 00:00:00+07	0633.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3686	\N	1	493	\N	\N	2017-05-17 00:00:00+07	0493.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5927	\N	1	2567	\N	\N	2019-11-22 00:00:00+07	2567.62.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6697	\N	1	916	\N	\N	2021-08-30 00:00:00+07	0916.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-08-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3685	\N	1	492	\N	\N	2017-12-05 00:00:00+07	0492.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3684	\N	1	491	\N	\N	2017-12-05 00:00:00+07	0491.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3683	\N	1	490	\N	\N	2017-12-05 00:00:00+07	0490.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6701	\N	1	917	\N	\N	2021-09-02 00:00:00+07	0917.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7267	\N	1	998	\N	\N	2022-02-18 00:00:00+07	0998.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-02-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6705	\N	1	629	\N	\N	2021-09-02 00:00:00+07	0629.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5934	\N	1	547	\N	\N	2019-11-25 00:00:00+07	0547.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7263	\N	1	1147	\N	\N	2022-02-16 00:00:00+07	1147.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-02-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7241	\N	1	2666	\N	\N	2022-02-08 00:00:00+07	2666.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5938	\N	1	548	\N	\N	2019-12-02 00:00:00+07	0548.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7475	\N	1	2906	\N	\N	2022-06-16 00:00:00+07	2906.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7208	\N	1	987	\N	\N	2022-01-31 00:00:00+07	0987.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-31 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5940	\N	1	549	\N	\N	2019-12-02 00:00:00+07	0549.63.101	2	2021-03-01 00:00:00+07	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5941	\N	1	828	\N	\N	2019-12-02 00:00:00+07	0828.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4418	\N	1	312	\N	\N	2018-03-01 00:00:00+07	0312.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5944	\N	1	829	\N	\N	2019-12-09 00:00:00+07	0829.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7182	\N	1	2657	\N	\N	2022-01-28 00:00:00+07	2657.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6720	\N	1	2537	\N	\N	2021-09-09 00:00:00+07	2537.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7181	\N	1	684	\N	\N	2022-01-27 00:00:00+07	0684.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6722	\N	1	634	\N	\N	2021-09-10 00:00:00+07	0634.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4789	\N	1	716	\N	\N	2018-10-11 00:00:00+07	0716.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4414	\N	1	403	\N	\N	2018-02-28 00:00:00+07	0403.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6748	\N	1	1074	\N	\N	2021-09-13 00:00:00+07	1074.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7173	\N	1	681	\N	\N	2022-01-25 00:00:00+07	0681.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4937	\N	1	1931	\N	\N	2019-01-28 00:00:00+07	1931.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4940	\N	1	2322	\N	\N	2019-01-29 00:00:00+07	2322.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4931	\N	1	897	\N	\N	2019-01-24 00:00:00+07	0897.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4929	\N	1	738	\N	\N	2019-01-18 00:00:00+07	0738.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6755	\N	1	1075	\N	\N	2021-09-15 00:00:00+07	1075.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6757	\N	1	2732	\N	\N	2021-09-16 00:00:00+07	2732.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5956	\N	1	2575	\N	\N	2019-12-10 00:00:00+07	2575.62.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5961	\N	1	551	\N	\N	2019-12-13 00:00:00+07	0551.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5962	\N	1	2353	\N	\N	2019-12-15 00:00:00+07	2353.65.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4279	\N	1	381	\N	\N	2017-11-22 00:00:00+07	0381.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6765	\N	1	1079	\N	\N	2021-09-17 00:00:00+07	1079.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4386	\N	1	867	\N	\N	2018-02-05 00:00:00+07	0867.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5672	\N	1	803	\N	\N	2019-09-10 00:00:00+07	0803.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6770	\N	1	2736	\N	\N	2021-09-21 00:00:00+07	2736.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5967	\N	1	1007	\N	\N	2019-12-23 00:00:00+07	1007.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7172	\N	1	1126	\N	\N	2022-01-25 00:00:00+07	1126.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6774	\N	1	2738	\N	\N	2021-09-22 00:00:00+07	2738.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4362	\N	1	399	\N	\N	2018-01-25 00:00:00+07	0399.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4361	\N	1	460	\N	\N	2018-01-24 00:00:00+07	0460.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4360	\N	1	1404	\N	\N	2018-01-24 00:00:00+07	1404.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6780	\N	1	919	\N	\N	2021-09-27 00:00:00+07	0919.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-09-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5681	\N	1	2260	\N	\N	2019-09-13 00:00:00+07	2260.65.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5990	\N	1	2366	\N	\N	2020-01-08 00:00:00+07	2366.65.101	1	\N	\N	SA	\N	10209	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4352	\N	1	398	\N	\N	2018-01-19 00:00:00+07	0398.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4928	\N	1	2318	\N	\N	2019-01-16 00:00:00+07	2318.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3493	\N	1	823	\N	\N	2017-03-01 00:00:00+07	0823.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4646	\N	1	2273	\N	\N	2018-08-06 00:00:00+07	2273.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4645	\N	1	2272	\N	\N	2018-08-06 00:00:00+07	2272.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4341	\N	1	1951	\N	\N	2018-01-11 00:00:00+07	1951.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5993	\N	1	834	\N	\N	2020-01-09 00:00:00+07	0834.64.101	1	\N	\N	SA	\N	10209	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5693	\N	1	2268	\N	\N	2019-09-18 00:00:00+07	2268.65.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5694	\N	1	978	\N	\N	2019-09-18 00:00:00+07	0978.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4340	\N	1	1969	\N	\N	2018-01-11 00:00:00+07	1969.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4339	\N	1	866	\N	\N	2018-01-10 00:00:00+07	0866.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5698	\N	1	979	\N	\N	2019-09-19 00:00:00+07	0979.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4335	\N	1	765	\N	\N	2018-01-02 00:00:00+07	0765.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4334	\N	1	763	\N	\N	2017-12-29 00:00:00+07	0763.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4333	\N	1	394	\N	\N	2017-12-29 00:00:00+07	0394.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5703	\N	1	521	\N	\N	2019-09-20 00:00:00+07	0521.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6008	\N	1	838	\N	\N	2020-01-21 00:00:00+07	0838.64.101	1	\N	\N	SA	\N	10209	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5704	\N	1	522	\N	\N	2019-09-20 00:00:00+07	0522.63.101	2	2020-02-04 00:00:00+07	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3365	\N	1	3073	\N	\N	2017-04-05 00:00:00+07	3073.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7141	\N	1	2811	\N	\N	2022-01-19 00:00:00+07	2811.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6014	\N	1	2589	\N	\N	2020-01-22 00:00:00+07	2589.62.101	1	\N	\N	SA	\N	10209	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5707	\N	1	523	\N	\N	2019-09-23 00:00:00+07	0523.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5708	\N	1	805	\N	\N	2019-09-23 00:00:00+07	0805.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7140	\N	1	2811	\N	\N	2022-01-19 00:00:00+07	2811.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6019	\N	1	554	\N	\N	2020-01-27 00:00:00+07	0554.63.101	1	\N	\N	SA	\N	10209	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4331	\N	1	321	\N	\N	2017-12-28 00:00:00+07	0321.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4330	\N	1	392	\N	\N	2017-12-27 00:00:00+07	0392.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4328	\N	1	2214	\N	\N	2017-12-27 00:00:00+07	2214.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4644	\N	1	427	\N	\N	2018-08-06 00:00:00+07	0427.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4327	\N	1	391	\N	\N	2017-12-27 00:00:00+07	0391.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4326	\N	1	390	\N	\N	2017-12-27 00:00:00+07	0390.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6828	\N	1	1085	\N	\N	2021-10-18 00:00:00+07	1085.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6829	\N	1	1085	\N	\N	2021-10-18 00:00:00+07	1085.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6830	\N	1	1085	\N	\N	2021-10-18 00:00:00+07	1085.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6831	\N	1	1085	\N	\N	2021-10-18 00:00:00+07	1085.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6832	\N	1	1085	\N	\N	2021-10-18 00:00:00+07	1085.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6833	\N	1	1085	\N	\N	2021-10-18 00:00:00+07	1085.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4325	\N	1	389	\N	\N	2017-12-22 00:00:00+07	0389.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4324	\N	1	388	\N	\N	2017-12-22 00:00:00+07	0388.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5473	\N	1	956	\N	\N	2019-07-03 00:00:00+07	0956.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4323	\N	1	387	\N	\N	2017-12-20 00:00:00+07	0387.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6841	\N	1	921	\N	\N	2021-10-21 00:00:00+07	0921.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4322	\N	1	386	\N	\N	2017-12-20 00:00:00+07	0386.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6035	\N	1	1012	\N	\N	2020-02-19 00:00:00+07	1012.61.101	1	\N	\N	SA	\N	10209	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4321	\N	1	385	\N	\N	2017-12-20 00:00:00+07	0385.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6037	\N	1	1013	\N	\N	2020-02-24 00:00:00+07	1013.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6847	\N	1	1089	\N	\N	2021-10-25 00:00:00+07	1089.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4629	\N	1	1772	\N	\N	2018-07-30 00:00:00+07	1772.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4317	\N	1	383	\N	\N	2017-12-15 00:00:00+07	0383.63.101	2	2018-01-10 00:00:00+07	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4316	\N	1	382	\N	\N	2017-12-14 00:00:00+07	0382.63.101	2	2018-03-20 00:00:00+07	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5359	\N	1	492	\N	\N	2019-05-27 00:00:00+07	0492.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4291	\N	1	2212	\N	\N	2017-12-06 00:00:00+07	2212.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6043	\N	1	555	\N	\N	2020-02-27 00:00:00+07	0555.63.101	1	\N	\N	SA	\N	10209	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6044	\N	1	556	\N	\N	2020-02-28 00:00:00+07	0556.63.101	1	\N	\N	SA	\N	10209	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4282	\N	1	864	\N	\N	2017-11-28 00:00:00+07	0864.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4620	\N	1	680	\N	\N	2018-07-24 00:00:00+07	0680.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4616	\N	1	678	\N	\N	2018-07-18 00:00:00+07	0678.64.101	1	2018-07-18 00:00:00+07	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6048	\N	1	2383	\N	\N	2020-03-04 00:00:00+07	2383.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6859	\N	1	924	\N	\N	2021-10-28 00:00:00+07	0924.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-10-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4926	\N	1	896	\N	\N	2019-01-15 00:00:00+07	0896.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6051	\N	1	557	\N	\N	2020-03-06 00:00:00+07	0557.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6052	\N	1	841	\N	\N	2020-03-09 00:00:00+07	0841.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6865	\N	1	926	\N	\N	2021-11-02 00:00:00+07	0926.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4278	\N	1	380	\N	\N	2017-11-22 00:00:00+07	0380.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6867	\N	1	1092	\N	\N	2021-11-02 00:00:00+07	1092.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4272	\N	1	862	\N	\N	2017-11-10 00:00:00+07	0862.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4268	\N	1	378	\N	\N	2017-11-08 00:00:00+07	0378.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5490	\N	1	784	\N	\N	2019-07-10 00:00:00+07	0784.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4267	\N	1	796	\N	\N	2017-11-08 00:00:00+07	0796.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4266	\N	1	377	\N	\N	2017-11-08 00:00:00+07	0377.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4265	\N	1	861	\N	\N	2017-11-08 00:00:00+07	0861.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6062	\N	1	2385	\N	\N	2020-03-18 00:00:00+07	2385.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4608	\N	1	425	\N	\N	2018-07-13 00:00:00+07	0425.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4250	\N	1	313	\N	\N	2017-11-01 00:00:00+07	0313.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6065	\N	1	842	\N	\N	2020-03-26 00:00:00+07	0842.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4248	\N	1	598	\N	\N	2017-10-31 00:00:00+07	0598.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4245	\N	1	300	\N	\N	2017-10-30 00:00:00+07	0300.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6068	\N	1	1017	\N	\N	2020-04-08 00:00:00+07	1017.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4226	\N	1	1668	\N	\N	2017-10-25 00:00:00+07	1668.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4972	\N	1	746	\N	\N	2019-02-13 00:00:00+07	0746.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6883	\N	1	646	\N	\N	2021-11-09 00:00:00+07	0646.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5047	\N	1	1990	\N	\N	2019-02-28 00:00:00+07	1990.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4167	\N	1	859	\N	\N	2017-10-13 00:00:00+07	0859.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5746	\N	1	984	\N	\N	2019-09-25 00:00:00+07	0984.61.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6074	\N	1	563	\N	\N	2020-04-15 00:00:00+07	0563.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6075	\N	1	844	\N	\N	2020-04-22 00:00:00+07	0844.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4162	\N	1	589	\N	\N	2017-10-10 00:00:00+07	0589.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6079	\N	1	1020	\N	\N	2020-04-29 00:00:00+07	1020.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6081	\N	1	2389	\N	\N	2020-06-02 00:00:00+07	2389.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6082	\N	1	2390	\N	\N	2020-06-02 00:00:00+07	2390.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6083	\N	1	2391	\N	\N	2020-06-02 00:00:00+07	2391.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4161	\N	1	2209	\N	\N	2017-10-10 00:00:00+07	2209.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5505	\N	1	503	\N	\N	2019-07-16 00:00:00+07	0503.63.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4605	\N	2	881	\N	\N	2018-07-12 00:00:00+07	0881.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6086	\N	1	565	\N	\N	2020-06-05 00:00:00+07	0565.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6087	\N	1	2394	\N	\N	2020-06-05 00:00:00+07	2394.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4156	\N	1	1494	\N	\N	2017-10-09 00:00:00+07	1494.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4604	\N	1	424	\N	\N	2018-07-12 00:00:00+07	0424.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4154	\N	1	2206	\N	\N	2017-10-09 00:00:00+07	2206.62.101	2	2017-10-10 00:00:00+07	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4153	\N	1	587	\N	\N	2017-10-09 00:00:00+07	0587.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4587	\N	1	423	\N	\N	2018-06-28 00:00:00+07	0423.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4923	\N	1	737	\N	\N	2019-01-08 00:00:00+07	0737.64.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4150	\N	1	857	\N	\N	2017-10-05 00:00:00+07	0857.61.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5759	\N	1	2498	\N	\N	2019-10-01 00:00:00+07	2498.62.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4576	\N	1	421	\N	\N	2018-05-28 00:00:00+07	0421.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7080	\N	1	673	\N	\N	2022-01-12 00:00:00+07	0673.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4149	\N	1	2205	\N	\N	2017-10-05 00:00:00+07	2205.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4575	\N	1	671	\N	\N	2018-05-25 00:00:00+07	0671.64.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4148	\N	1	1989	\N	\N	2017-10-05 00:00:00+07	1989.62.101	2	2017-10-10 00:00:00+07	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4147	\N	1	1990	\N	\N	2017-10-05 00:00:00+07	1990.62.101	2	2017-10-10 00:00:00+07	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4082	\N	1	1661	\N	\N	2017-09-22 00:00:00+07	1661.65.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4081	\N	1	371	\N	\N	2017-09-22 00:00:00+07	0371.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4080	\N	1	2198	\N	\N	2017-09-22 00:00:00+07	2198.62.101	2	2020-07-20 00:00:00+07	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4079	\N	1	16	\N	\N	2015-10-14 00:00:00+07	0016.00.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4574	\N	1	420	\N	\N	2018-05-21 00:00:00+07	0420.63.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5524	\N	1	2441	\N	\N	2019-07-22 00:00:00+07	2441.62.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6109	\N	1	1022	\N	\N	2020-06-30 00:00:00+07	1022.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4141	\N	1	2201	\N	\N	2017-10-04 00:00:00+07	2201.62.101	1	\N	\N	SA	\N	10209	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6111	\N	1	1024	\N	\N	2020-07-01 00:00:00+07	1024.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7019	\N	1	2634	\N	\N	2021-12-28 00:00:00+07	2634.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-12-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6216	\N	1	581	\N	\N	2020-09-14 00:00:00+07	0581.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-09-14 00:00:00+07	\N	09:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7079	\N	1	672	\N	\N	2022-01-12 00:00:00+07	0672.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7078	\N	1	2640	\N	\N	2022-01-12 00:00:00+07	2640.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7077	\N	1	966	\N	\N	2022-01-12 00:00:00+07	0966.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7076	\N	1	2800	\N	\N	2022-01-12 00:00:00+07	2800.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7075	\N	1	2799	\N	\N	2022-01-12 00:00:00+07	2799.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7072	\N	1	671	\N	\N	2022-01-11 00:00:00+07	0671.63.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6217	\N	1	2628	\N	\N	2020-09-15 00:00:00+07	2628.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-09-15 00:00:00+07	\N	09:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7071	\N	1	2638	\N	\N	2022-01-11 00:00:00+07	2638.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7068	\N	1	964	\N	\N	2022-01-06 00:00:00+07	0964.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7065	\N	1	2636	\N	\N	2022-01-06 00:00:00+07	2636.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6218	\N	1	2629	\N	\N	2020-09-15 00:00:00+07	2629.62.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2020-09-15 00:00:00+07	\N	09:00	22:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7129	\N	1	972	\N	\N	2022-01-18 00:00:00+07	0972.64.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2022-01-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7005	\N	1	1107	\N	\N	2021-12-21 00:00:00+07	1107.61.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-12-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6184	\N	1	2426	\N	\N	2020-09-08 00:00:00+07	2426.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	2020-08-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6185	\N	1	2427	\N	\N	2020-09-09 00:00:00+07	2427.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	\N	\N	2020-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5580	\N	1	1377	\N	\N	2019-08-06 00:00:00+07	1377.65.101	1	\N	\N	SA	\N	10209	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7016	\N	1	2631	\N	\N	2021-12-24 00:00:00+07	2631.65.101	1	\N	\N	SA	\N	10209	Amalia Mahardani	\N	\N	f	f	2021-12-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3336	\N	1	1111	\N	\N	2017-03-01 00:00:00+07	1111.62.310	2	2017-03-01 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3371	\N	1	0	\N	\N	2017-03-01 00:00:00+07	0000.61.310	2	2017-03-07 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3396	\N	1	2222	\N	\N	2017-03-01 00:00:00+07	2222.62.310	2	\N	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3530	\N	1	1234	\N	\N	2017-05-01 00:00:00+07	1234.62.310	2	2017-05-03 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3543	\N	1	0	\N	\N	2017-05-04 00:00:00+07	000.65.310	2	2017-05-04 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3484	\N	1	3162	\N	\N	2017-03-01 00:00:00+07	3162.62.310	2	2017-03-08 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7198	\N	1	2829	\N	\N	2022-01-28 00:00:00+07	2829.62.310	1	\N	\N	OA	\N	10119	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3397	\N	1	3333	\N	\N	2017-03-01 00:00:00+07	3333.62.310	2	2017-03-07 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3462	\N	1	362	\N	\N	2017-03-01 00:00:00+07	0362.62.310	2	2017-03-01 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3528	\N	1	3261	\N	\N	2017-04-03 00:00:00+07	3261.65.310	2	2017-04-10 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3511	\N	1	3262	\N	\N	2017-03-01 00:00:00+07	3262.62.310	2	2013-03-08 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3610	\N	1	3862	\N	\N	2017-05-02 00:00:00+07	3862.62.310	2	2017-05-08 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2968	\N	1	0	\N	\N	2017-03-23 00:00:00+07	0000.62.310	2	2017-03-23 00:00:00+07	\N	SA	\N	10119	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
344	\N	2	1412	\N	\N	2012-01-02 00:00:00+07	1412.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
343	\N	2	1410	\N	\N	2012-01-02 00:00:00+07	1410.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
342	\N	2	1403	\N	\N	2012-01-02 00:00:00+07	1403.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
341	\N	2	1415	\N	\N	2012-01-02 00:00:00+07	1415.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
340	\N	2	1433	\N	\N	2012-01-02 00:00:00+07	1433.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6839	\N	1	1087	\N	\N	2021-10-21 00:00:00+07	1087.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-10-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
339	\N	2	1439	\N	\N	2012-01-02 00:00:00+07	1439.62.700	1	\N	2020-06-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
338	\N	2	1453	\N	\N	2012-01-02 00:00:00+07	1453.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
337	\N	2	1459	\N	\N	2012-01-02 00:00:00+07	1459.62.700	1	\N	2020-06-01 00:00:00+07	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
336	\N	2	1466	\N	\N	2012-01-02 00:00:00+07	1466.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
335	\N	2	1476	\N	\N	2012-01-02 00:00:00+07	1476.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
334	\N	2	1425	\N	\N	2012-01-02 00:00:00+07	1425.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
333	\N	2	1500	\N	\N	2012-01-02 00:00:00+07	1500.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
332	\N	2	1424	\N	\N	2012-01-02 00:00:00+07	1424.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
331	\N	2	1531	\N	\N	2012-01-02 00:00:00+07	1531.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
330	\N	2	1505	\N	\N	2012-01-02 00:00:00+07	1505.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
329	\N	2	1422	\N	\N	2012-01-02 00:00:00+07	1422.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-06-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
328	\N	2	1416	\N	\N	2012-01-02 00:00:00+07	1416.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
327	\N	2	1432	\N	\N	2012-01-02 00:00:00+07	1432.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
326	\N	2	1509	\N	\N	2012-01-02 00:00:00+07	1509.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
325	\N	2	519	\N	\N	2012-01-02 00:00:00+07	0519.65.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
324	\N	2	1417	\N	\N	2012-01-02 00:00:00+07	1417.62.700	1	\N	2020-05-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
323	\N	2	592	\N	\N	2012-01-02 00:00:00+07	0592.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
322	\N	2	569	\N	\N	2012-01-02 00:00:00+07	0569.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
321	\N	2	562	\N	\N	2012-01-02 00:00:00+07	0562.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
320	\N	2	561	\N	\N	2012-01-02 00:00:00+07	0561.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
319	\N	2	559	\N	\N	2012-01-02 00:00:00+07	0559.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
318	\N	2	558	\N	\N	2012-01-02 00:00:00+07	0558.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
317	\N	2	556	\N	\N	2012-01-02 00:00:00+07	0556.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
316	\N	2	555	\N	\N	2012-01-02 00:00:00+07	0555.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
315	\N	2	554	\N	\N	2012-01-02 00:00:00+07	0554.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6871	\N	1	644	\N	\N	2021-11-03 00:00:00+07	0644.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
314	\N	2	553	\N	\N	2012-01-02 00:00:00+07	0553.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
313	\N	2	552	\N	\N	2012-01-02 00:00:00+07	0552.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
312	\N	2	551	\N	\N	2012-01-02 00:00:00+07	0551.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
311	\N	2	550	\N	\N	2012-01-02 00:00:00+07	0550.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
310	\N	2	543	\N	\N	2012-01-02 00:00:00+07	0543.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
309	\N	2	546	\N	\N	2012-01-02 00:00:00+07	0546.65.700	2	2021-07-28 00:00:00+07	\N	OA	\N	10191	solusiti	\N	\N	f	f	2021-07-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
308	\N	2	545	\N	\N	2012-01-02 00:00:00+07	0545.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
307	\N	2	544	\N	\N	2012-01-02 00:00:00+07	0544.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
306	\N	2	540	\N	\N	2012-01-02 00:00:00+07	0540.65.700	2	2019-10-23 00:00:00+07	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
305	\N	2	539	\N	\N	2012-01-02 00:00:00+07	0539.65.700	2	2016-11-01 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
304	\N	2	538	\N	\N	2012-01-02 00:00:00+07	0538.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
303	\N	2	537	\N	\N	2012-01-02 00:00:00+07	0537.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
302	\N	2	536	\N	\N	2012-01-02 00:00:00+07	0536.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
301	\N	2	533	\N	\N	2012-01-02 00:00:00+07	0533.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
300	\N	2	532	\N	\N	2012-01-02 00:00:00+07	0532.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
299	\N	2	529	\N	\N	2012-01-02 00:00:00+07	0529.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
298	\N	2	528	\N	\N	2012-01-02 00:00:00+07	0528.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
297	\N	2	527	\N	\N	2012-01-02 00:00:00+07	0527.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
296	\N	2	526	\N	\N	2012-01-02 00:00:00+07	0526.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
295	\N	2	350	\N	\N	2012-01-02 00:00:00+07	0350.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6896	\N	1	1094	\N	\N	2021-11-10 00:00:00+07	1094.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
294	\N	2	347	\N	\N	2012-01-02 00:00:00+07	0347.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
293	\N	2	342	\N	\N	2012-01-02 00:00:00+07	0342.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
292	\N	2	343	\N	\N	2012-01-02 00:00:00+07	0343.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
291	\N	2	330	\N	\N	2012-01-02 00:00:00+07	0330.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
290	\N	2	341	\N	\N	2012-01-02 00:00:00+07	0341.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6902	\N	1	930	\N	\N	2021-11-15 00:00:00+07	0930.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
289	\N	2	340	\N	\N	2012-01-02 00:00:00+07	0340.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
288	\N	2	337	\N	\N	2012-01-02 00:00:00+07	0337.64.700	2	2018-06-26 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
287	\N	2	336	\N	\N	2012-01-02 00:00:00+07	0336.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
286	\N	2	335	\N	\N	2012-01-02 00:00:00+07	0335.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
285	\N	2	326	\N	\N	2012-01-02 00:00:00+07	0326.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
284	\N	2	321	\N	\N	2012-01-02 00:00:00+07	0321.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
283	\N	2	320	\N	\N	2012-01-02 00:00:00+07	0320.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
282	\N	2	319	\N	\N	2012-01-02 00:00:00+07	0319.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
281	\N	2	318	\N	\N	2012-01-02 00:00:00+07	0318.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
280	\N	2	317	\N	\N	2012-01-02 00:00:00+07	0317.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
279	\N	2	316	\N	\N	2012-01-02 00:00:00+07	0316.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
278	\N	2	315	\N	\N	2012-01-02 00:00:00+07	0315.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
277	\N	2	314	\N	\N	2012-01-02 00:00:00+07	0314.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
276	\N	2	313	\N	\N	2012-01-02 00:00:00+07	0313.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
275	\N	2	311	\N	\N	2012-01-02 00:00:00+07	0311.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
274	\N	2	309	\N	\N	2012-01-02 00:00:00+07	0309.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
273	\N	2	307	\N	\N	2012-01-02 00:00:00+07	0307.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
272	\N	2	306	\N	\N	2012-01-02 00:00:00+07	0306.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
271	\N	2	305	\N	\N	2012-01-02 00:00:00+07	0305.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
270	\N	2	304	\N	\N	2012-01-02 00:00:00+07	0304.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
269	\N	2	303	\N	\N	2012-01-02 00:00:00+07	0303.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
268	\N	2	302	\N	\N	2012-01-02 00:00:00+07	0302.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
267	\N	2	301	\N	\N	2012-01-02 00:00:00+07	0301.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
266	\N	2	300	\N	\N	2012-01-02 00:00:00+07	0300.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
265	\N	2	299	\N	\N	2012-01-02 00:00:00+07	0299.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
264	\N	2	298	\N	\N	2012-01-02 00:00:00+07	0298.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
263	\N	2	297	\N	\N	2012-01-02 00:00:00+07	0297.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2	\N	2	498	\N	\N	2011-02-01 00:00:00+07	0498.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2926	\N	1	220	\N	\N	2017-02-01 00:00:00+07	220.62.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
262	\N	2	296	\N	\N	2012-01-02 00:00:00+07	0296.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-12-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
261	\N	2	295	\N	\N	2012-01-02 00:00:00+07	0295.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
260	\N	2	294	\N	\N	2012-01-02 00:00:00+07	0294.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
259	\N	2	293	\N	\N	2012-01-02 00:00:00+07	0293.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
258	\N	2	292	\N	\N	2012-01-02 00:00:00+07	0292.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
257	\N	2	291	\N	\N	2012-01-02 00:00:00+07	0291.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
256	\N	2	288	\N	\N	2012-01-02 00:00:00+07	0288.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
255	\N	2	288	\N	\N	2012-01-02 00:00:00+07	0288.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
254	\N	2	287	\N	\N	2012-01-02 00:00:00+07	0287.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
253	\N	2	285	\N	\N	2012-01-02 00:00:00+07	0285.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
252	\N	2	284	\N	\N	2012-01-02 00:00:00+07	0284.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
251	\N	2	283	\N	\N	2012-01-02 00:00:00+07	0283.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
250	\N	2	282	\N	\N	2012-01-02 00:00:00+07	0282.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
249	\N	2	281	\N	\N	2012-01-02 00:00:00+07	0281.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
248	\N	2	279	\N	\N	2012-01-02 00:00:00+07	0279.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
247	\N	2	278	\N	\N	2012-01-02 00:00:00+07	0278.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
246	\N	2	277	\N	\N	2012-01-02 00:00:00+07	0277.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
245	\N	2	237	\N	\N	2012-01-02 00:00:00+07	0237.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
244	\N	2	276	\N	\N	2012-01-02 00:00:00+07	0276.64.700	2	2017-08-14 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
243	\N	2	275	\N	\N	2012-01-02 00:00:00+07	0275.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6945	\N	1	937	\N	\N	2021-11-30 00:00:00+07	0937.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
242	\N	2	274	\N	\N	2012-01-02 00:00:00+07	0274.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
241	\N	2	273	\N	\N	2012-01-02 00:00:00+07	0273.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
240	\N	2	272	\N	\N	2012-01-02 00:00:00+07	0272.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
239	\N	2	271	\N	\N	2012-01-02 00:00:00+07	0271.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
238	\N	2	269	\N	\N	2012-01-02 00:00:00+07	0269.64.700	2	2018-05-01 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
237	\N	2	268	\N	\N	2012-01-02 00:00:00+07	0268.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
236	\N	2	267	\N	\N	2012-01-02 00:00:00+07	0267.64.700	2	2020-12-01 00:00:00+07	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-02-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
235	\N	2	266	\N	\N	2012-01-02 00:00:00+07	0266.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
234	\N	2	265	\N	\N	2012-01-02 00:00:00+07	0265.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
233	\N	2	264	\N	\N	2012-01-02 00:00:00+07	0264.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
232	\N	2	263	\N	\N	2012-01-02 00:00:00+07	0263.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
231	\N	2	262	\N	\N	2012-01-02 00:00:00+07	0262.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
230	\N	2	260	\N	\N	2012-01-02 00:00:00+07	0260.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
229	\N	2	259	\N	\N	2012-01-02 00:00:00+07	0259.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6976	\N	1	2782	\N	\N	2021-12-09 00:00:00+07	2782.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-12-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
228	\N	2	258	\N	\N	2012-01-02 00:00:00+07	0258.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
227	\N	2	257	\N	\N	2012-01-02 00:00:00+07	0257.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
226	\N	2	256	\N	\N	2012-01-02 00:00:00+07	0256.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
225	\N	2	255	\N	\N	2012-01-02 00:00:00+07	0255.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
224	\N	2	254	\N	\N	2012-01-02 00:00:00+07	0254.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
223	\N	2	253	\N	\N	2012-01-02 00:00:00+07	0253.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
222	\N	2	252	\N	\N	2012-01-02 00:00:00+07	0252.64.700	0	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
221	\N	2	251	\N	\N	2012-01-02 00:00:00+07	0251.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
220	\N	2	250	\N	\N	2012-01-02 00:00:00+07	0250.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
219	\N	2	249	\N	\N	2012-01-02 00:00:00+07	0249.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
218	\N	2	248	\N	\N	2012-01-02 00:00:00+07	0248.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
217	\N	2	247	\N	\N	2012-01-02 00:00:00+07	0247.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
216	\N	2	246	\N	\N	2012-01-02 00:00:00+07	0246.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
215	\N	2	245	\N	\N	2012-01-02 00:00:00+07	0245.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
214	\N	2	244	\N	\N	2012-01-02 00:00:00+07	0244.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
213	\N	2	243	\N	\N	2012-01-02 00:00:00+07	0243.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
212	\N	2	242	\N	\N	2012-01-02 00:00:00+07	0242.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
211	\N	2	240	\N	\N	2012-01-02 00:00:00+07	0240.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
210	\N	2	239	\N	\N	2012-01-02 00:00:00+07	0239.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
209	\N	2	238	\N	\N	2012-01-02 00:00:00+07	0238.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
208	\N	2	215	\N	\N	2012-01-02 00:00:00+07	0215.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
207	\N	2	211	\N	\N	2012-01-02 00:00:00+07	0211.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7007	\N	1	949	\N	\N	2021-12-21 00:00:00+07	0949.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-12-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
206	\N	2	199	\N	\N	2012-01-02 00:00:00+07	0199.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
205	\N	2	198	\N	\N	2012-01-02 00:00:00+07	0198.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
204	\N	2	197	\N	\N	2012-01-02 00:00:00+07	0197.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
203	\N	2	196	\N	\N	2012-01-02 00:00:00+07	0196.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
202	\N	2	1538	\N	\N	2012-01-02 00:00:00+07	1538.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
201	\N	2	195	\N	\N	2012-01-02 00:00:00+07	0195.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7014	\N	1	951	\N	\N	2021-12-23 00:00:00+07	0951.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-12-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7015	\N	1	952	\N	\N	2021-12-23 00:00:00+07	0952.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-12-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7017	\N	1	2632	\N	\N	2021-12-27 00:00:00+07	2632.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-12-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
200	\N	2	194	\N	\N	2012-01-02 00:00:00+07	0194.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
199	\N	2	193	\N	\N	2012-01-02 00:00:00+07	0193.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
198	\N	2	192	\N	\N	2012-01-02 00:00:00+07	0192.63.700	2	2019-01-01 00:00:00+07	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
197	\N	2	189	\N	\N	2012-01-02 00:00:00+07	0189.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4962	\N	1	458	\N	\N	2019-02-12 00:00:00+07	0458.63.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4960	\N	1	741	\N	\N	2019-02-12 00:00:00+07	0741.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-05-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4938	\N	1	455	\N	\N	2019-01-28 00:00:00+07	0455.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4994	\N	1	1955	\N	\N	2019-02-19 00:00:00+07	1955.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5371	\N	1	945	\N	\N	2019-06-10 00:00:00+07	0945.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5372	\N	1	946	\N	\N	2019-06-10 00:00:00+07	0946.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4993	\N	1	462	\N	\N	2019-02-18 00:00:00+07	0462.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4992	\N	1	1954	\N	\N	2019-02-18 00:00:00+07	1954.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5575	\N	1	787	\N	\N	2019-08-05 00:00:00+07	0787.64.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4702	\N	1	1809	\N	\N	2018-09-12 00:00:00+07	1809.65.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4701	\N	1	1808	\N	\N	2018-09-12 00:00:00+07	1808.65.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4688	\N	1	694	\N	\N	2018-08-31 00:00:00+07	0694.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4687	\N	1	886	\N	\N	2018-08-31 00:00:00+07	0886.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4662	\N	1	428	\N	\N	2018-08-15 00:00:00+07	0428.63.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4614	\N	1	677	\N	\N	2018-07-18 00:00:00+07	0677.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4584	\N	1	673	\N	\N	2018-06-07 00:00:00+07	0673.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4538	\N	1	879	\N	\N	2018-04-30 00:00:00+07	0879.61.700	2	2018-05-01 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4526	\N	1	2255	\N	\N	2018-04-19 00:00:00+07	2255.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4521	\N	1	1725	\N	\N	2018-04-18 00:00:00+07	1725.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5419	\N	1	952	\N	\N	2019-06-19 00:00:00+07	0952.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4510	\N	1	662	\N	\N	2018-04-16 00:00:00+07	0662.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4507	\N	1	659	\N	\N	2018-04-13 00:00:00+07	0659.64.700	2	2018-04-13 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4498	\N	1	620	\N	\N	2018-04-11 00:00:00+07	0620.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4478	\N	1	871	\N	\N	2018-03-28 00:00:00+07	0871.61.700	1	2018-03-31 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4454	\N	1	2245	\N	\N	2018-03-19 00:00:00+07	2245.62.700	1	\N	2020-05-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4450	\N	1	640	\N	\N	2018-03-15 00:00:00+07	0640.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4431	\N	1	2237	\N	\N	2018-03-08 00:00:00+07	2237.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5440	\N	1	953	\N	\N	2019-06-25 00:00:00+07	0953.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4411	\N	1	630	\N	\N	2018-02-23 00:00:00+07	0630.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4410	\N	1	401	\N	\N	2018-02-23 00:00:00+07	0401.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4409	\N	1	400	\N	\N	2018-02-23 00:00:00+07	0400.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4408	\N	1	399	\N	\N	2018-02-23 00:00:00+07	0399.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4356	\N	1	2223	\N	\N	2018-01-22 00:00:00+07	2223.62.700	1	\N	2019-06-09 00:00:00+07	SA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4350	\N	1	2220	\N	\N	2018-01-11 00:00:00+07	2220.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4348	\N	2	612	\N	\N	2018-01-11 00:00:00+07	0612.64.700	2	2019-10-01 00:00:00+07	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-02-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4345	\N	2	397	\N	\N	2018-01-11 00:00:00+07	0397.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5468	\N	1	499	\N	\N	2019-07-02 00:00:00+07	0499.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4319	\N	2	384	\N	\N	2017-12-18 00:00:00+07	0384.63.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5483	\N	1	2193	\N	\N	2019-07-09 00:00:00+07	2193.65.700	2	2021-02-01 00:00:00+07	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-04-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
65	\N	2	530	\N	\N	2011-02-01 00:00:00+07	0530.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
64	\N	2	286	\N	\N	2011-02-01 00:00:00+07	0286.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4255	\N	1	1670	\N	\N	2017-11-01 00:00:00+07	1670.65.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4160	\N	1	588	\N	\N	2017-10-10 00:00:00+07	0588.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
63	\N	2	183	\N	\N	2011-02-01 00:00:00+07	0183.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
62	\N	2	1499	\N	\N	2011-02-01 00:00:00+07	1499.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5512	\N	1	2202	\N	\N	2019-07-17 00:00:00+07	2202.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
61	\N	2	1484	\N	\N	2011-02-01 00:00:00+07	1484.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5516	\N	1	786	\N	\N	2019-07-18 00:00:00+07	0786.64.700	1	\N	\N	SA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4078	\N	2	854	\N	\N	2017-09-22 00:00:00+07	0854.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4140	\N	2	856	\N	\N	2017-10-02 00:00:00+07	0856.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4139	\N	2	855	\N	\N	2017-10-02 00:00:00+07	0855.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
60	\N	2	1452	\N	\N	2011-02-01 00:00:00+07	1452.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
59	\N	2	600	\N	\N	2011-02-01 00:00:00+07	0600.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
58	\N	2	574	\N	\N	2011-02-01 00:00:00+07	0574.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7417	\N	1	2719	\N	\N	2022-05-23 00:00:00+07	2719.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-05-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
57	\N	2	572	\N	\N	2011-02-01 00:00:00+07	0572.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3887	\N	1	566	\N	\N	2017-07-17 00:00:00+07	0566.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
56	\N	2	571	\N	\N	2011-02-01 00:00:00+07	0571.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3817	\N	1	1632	\N	\N	2017-06-01 00:00:00+07	1632.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
55	\N	2	570	\N	\N	2011-02-01 00:00:00+07	0570.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
54	\N	2	569	\N	\N	2011-02-01 00:00:00+07	0569.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
53	\N	2	568	\N	\N	2011-02-01 00:00:00+07	0568.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
52	\N	2	566	\N	\N	2011-02-01 00:00:00+07	0566.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7424	\N	1	1019	\N	\N	2022-05-24 00:00:00+07	1019.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
51	\N	2	565	\N	\N	2011-02-01 00:00:00+07	0565.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
50	\N	2	561	\N	\N	2011-02-01 00:00:00+07	0561.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3736	\N	1	836	\N	\N	2017-05-23 00:00:00+07	0836.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2020-10-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7434	\N	1	708	\N	\N	2022-06-02 00:00:00+07	0708.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-06-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3748	\N	2	2150	\N	\N	2017-05-29 00:00:00+07	2150.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3742	\N	2	547	\N	\N	2017-05-24 00:00:00+07	0547.64.700	2	2018-05-01 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3741	\N	2	1626	\N	\N	2017-05-24 00:00:00+07	1626.65.700	2	2019-10-23 00:00:00+07	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3730	\N	1	1625	\N	\N	2017-05-02 00:00:00+07	1625.65.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3692	\N	2	810	\N	\N	2017-03-29 00:00:00+07	0810.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3691	\N	2	809	\N	\N	2017-05-17 00:00:00+07	0809.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3668	\N	1	545	\N	\N	2017-05-02 00:00:00+07	0545.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3651	\N	2	1588	\N	\N	2017-03-14 00:00:00+07	1588.65.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3581	\N	1	834	\N	\N	2017-05-02 00:00:00+07	0834.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3580	\N	1	833	\N	\N	2017-05-02 00:00:00+07	0833.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3549	\N	1	807	\N	\N	2017-04-03 00:00:00+07	0807.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3548	\N	1	2087	\N	\N	2017-04-03 00:00:00+07	2087.62.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3547	\N	1	2088	\N	\N	2017-04-03 00:00:00+07	2088.62.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3559	\N	1	2070	\N	\N	2017-04-03 00:00:00+07	2070.62.700	2	2018-10-22 00:00:00+07	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3544	\N	1	2059	\N	\N	2017-03-03 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3523	\N	1	325	\N	\N	2017-03-01 00:00:00+07	0325.63.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3522	\N	1	2057	\N	\N	2017-03-01 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3521	\N	1	2058	\N	\N	2017-03-01 00:00:00+07	2058.62.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3520	\N	1	2060	\N	\N	2017-03-01 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3527	\N	1	814	\N	\N	2017-04-03 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3509	\N	1	815	\N	\N	2017-03-01 00:00:00+07	0815.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3457	\N	1	820	\N	\N	2017-03-01 00:00:00+07	0820.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3439	\N	1	850	\N	\N	2017-03-01 00:00:00+07	0850.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3438	\N	1	813	\N	\N	2017-03-01 00:00:00+07	0813.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3428	\N	1	1607	\N	\N	2017-03-01 00:00:00+07	1607.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3427	\N	1	535	\N	\N	2017-03-01 00:00:00+07	0535.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3426	\N	1	534	\N	\N	2017-03-01 00:00:00+07	0534.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3425	\N	1	533	\N	\N	2017-03-01 00:00:00+07	0533.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3386	\N	2	328	\N	\N	2016-12-05 00:00:00+07	0328.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3383	\N	1	520	\N	\N	2017-03-01 00:00:00+07	0520.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3382	\N	1	519	\N	\N	2017-03-01 00:00:00+07	0519.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5702	\N	1	981	\N	\N	2019-09-19 00:00:00+07	0981.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3334	\N	1	3071	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3333	\N	1	3070	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3332	\N	1	3069	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3331	\N	1	3068	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3330	\N	1	3067	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3329	\N	1	3066	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3328	\N	1	3065	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3327	\N	1	3064	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3326	\N	1	3063	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3325	\N	1	3062	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3324	\N	1	3061	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3323	\N	1	3060	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	2020-06-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3322	\N	1	3059	\N	\N	2017-03-29 00:00:00+07	2069.62.700	2	2018-10-22 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3321	\N	1	3058	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	2020-06-01 00:00:00+07	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3320	\N	1	3057	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3319	\N	1	3056	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3318	\N	1	3055	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3317	\N	1	3054	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3316	\N	1	3053	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3315	\N	1	3052	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3314	\N	1	3051	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3313	\N	1	3050	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3458	\N	1	507	\N	\N	2017-03-01 00:00:00+07	0507.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3310	\N	1	3047	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3309	\N	1	3046	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3308	\N	1	3045	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3307	\N	1	3044	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3306	\N	1	3043	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3305	\N	1	3042	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3304	\N	1	3041	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3303	\N	1	3040	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	2020-05-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3302	\N	1	3039	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3301	\N	1	3038	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3300	\N	1	3037	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3298	\N	1	3035	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3296	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3295	\N	1	3034	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3294	\N	1	3033	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3293	\N	1	3032	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3292	\N	1	3031	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3291	\N	1	3030	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3290	\N	1	3029	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3289	\N	1	3028	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3288	\N	1	3027	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3287	\N	1	3026	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
49	\N	2	560	\N	\N	2011-02-01 00:00:00+07	0560.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
48	\N	2	559	\N	\N	2011-02-01 00:00:00+07	0559.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
47	\N	2	558	\N	\N	2011-02-01 00:00:00+07	0558.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3286	\N	1	3025	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3285	\N	1	3024	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3284	\N	1	3023	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3283	\N	1	3022	\N	\N	2017-03-29 00:00:00+07	2069.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3282	\N	1	3021	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3281	\N	1	3020	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3280	\N	1	3019	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3279	\N	1	3018	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5784	\N	1	2502	\N	\N	2019-10-08 00:00:00+07	2502.62.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3278	\N	1	3017	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3277	\N	1	3016	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3276	\N	1	3015	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5788	\N	1	2320	\N	\N	2019-10-09 00:00:00+07	2320.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3275	\N	1	3014	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3274	\N	1	3013	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3273	\N	1	3012	\N	\N	2017-03-29 00:00:00+07	1588.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3272	\N	1	3011	\N	\N	2017-03-29 00:00:00+07	0817.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3271	\N	1	3010	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3270	\N	1	3009	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3269	\N	1	3008	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
46	\N	2	557	\N	\N	2011-02-01 00:00:00+07	0557.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3268	\N	1	3007	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3267	\N	1	3006	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3266	\N	1	3005	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3265	\N	1	3004	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3264	\N	1	3003	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3263	\N	1	3002	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3262	\N	1	3001	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3261	\N	1	3000	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3260	\N	1	2999	\N	\N	2017-03-29 00:00:00+07	0325.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3259	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3258	\N	1	3034	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3257	\N	1	3033	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3256	\N	1	3032	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3255	\N	1	3031	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3254	\N	1	3030	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3253	\N	1	3029	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3252	\N	1	3028	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3251	\N	1	3027	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3250	\N	1	3026	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3249	\N	1	3025	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3248	\N	1	3024	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3247	\N	1	3023	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3246	\N	1	3022	\N	\N	2017-03-29 00:00:00+07	2069.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3245	\N	1	3021	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3244	\N	1	3020	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3243	\N	1	3019	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3242	\N	1	3018	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3241	\N	1	3017	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3240	\N	1	3016	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3239	\N	1	3015	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3238	\N	1	3014	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3237	\N	1	3013	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3236	\N	1	3012	\N	\N	2017-03-29 00:00:00+07	1588.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3235	\N	1	3011	\N	\N	2017-03-29 00:00:00+07	0817.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3234	\N	1	3010	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3233	\N	1	3009	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3232	\N	1	3008	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3231	\N	1	3007	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3230	\N	1	3006	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3229	\N	1	3005	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3228	\N	1	3004	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3227	\N	1	3003	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3226	\N	1	3002	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3225	\N	1	3001	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3224	\N	1	3000	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3223	\N	1	2999	\N	\N	2017-03-29 00:00:00+07	0325.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3222	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3221	\N	1	3034	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3220	\N	1	3033	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3219	\N	1	3032	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3218	\N	1	3031	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3217	\N	1	3030	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3216	\N	1	3029	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3215	\N	1	3028	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3214	\N	1	3027	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3213	\N	1	3026	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3212	\N	1	3025	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3211	\N	1	3024	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3210	\N	1	3023	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3209	\N	1	3022	\N	\N	2017-03-29 00:00:00+07	2069.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3208	\N	1	3021	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3207	\N	1	3020	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3206	\N	1	3019	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3205	\N	1	3018	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3204	\N	1	3017	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3203	\N	1	3016	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3202	\N	1	3015	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3201	\N	1	3014	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3200	\N	1	3013	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3199	\N	1	3012	\N	\N	2017-03-29 00:00:00+07	1588.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3198	\N	1	3011	\N	\N	2017-03-29 00:00:00+07	0817.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3197	\N	1	3010	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3196	\N	1	3009	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3195	\N	1	3008	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3194	\N	1	3007	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3193	\N	1	3006	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3192	\N	1	3005	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3191	\N	1	3004	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5894	\N	1	2336	\N	\N	2019-11-08 00:00:00+07	2336.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3190	\N	1	3003	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3189	\N	1	3002	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3188	\N	1	3001	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3187	\N	1	3000	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3186	\N	1	2999	\N	\N	2017-03-29 00:00:00+07	0325.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3185	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5901	\N	1	543	\N	\N	2019-11-12 00:00:00+07	0543.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5902	\N	1	544	\N	\N	2019-11-12 00:00:00+07	0544.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3184	\N	1	3034	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3183	\N	1	3033	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3182	\N	1	3032	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3181	\N	1	3031	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3180	\N	1	3030	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3179	\N	1	3029	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3178	\N	1	3028	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3177	\N	1	3027	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3176	\N	1	3026	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3175	\N	1	3025	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3174	\N	1	3024	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3173	\N	1	3023	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3172	\N	1	3022	\N	\N	2017-03-29 00:00:00+07	2069.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3171	\N	1	3021	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5922	\N	1	545	\N	\N	2019-11-19 00:00:00+07	0545.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3170	\N	1	3020	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3169	\N	1	3019	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3168	\N	1	3018	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
45	\N	2	554	\N	\N	2011-02-01 00:00:00+07	0554.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5929	\N	1	2569	\N	\N	2019-11-22 00:00:00+07	2569.62.700	1	\N	\N	SA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3167	\N	1	3017	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3166	\N	1	3016	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5932	\N	1	2571	\N	\N	2019-11-22 00:00:00+07	2571.62.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3165	\N	1	3015	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3164	\N	1	3014	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3163	\N	1	3013	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3162	\N	1	3012	\N	\N	2017-03-29 00:00:00+07	1588.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3161	\N	1	3011	\N	\N	2017-03-29 00:00:00+07	0817.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3160	\N	1	3010	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3159	\N	1	3009	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3158	\N	1	3008	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3157	\N	1	3007	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3156	\N	1	3006	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3155	\N	1	3005	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3154	\N	1	3004	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3153	\N	1	3003	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3152	\N	1	3002	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3151	\N	1	3001	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3150	\N	1	3000	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3149	\N	1	2999	\N	\N	2017-03-29 00:00:00+07	0325.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3148	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3145	\N	1	3034	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3144	\N	1	3033	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3143	\N	1	3032	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3142	\N	1	3031	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3141	\N	1	3030	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3140	\N	1	3029	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3139	\N	1	3028	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3138	\N	1	3027	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3137	\N	1	3026	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3136	\N	1	3025	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3135	\N	1	3024	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3134	\N	1	3023	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3133	\N	1	3022	\N	\N	2017-03-29 00:00:00+07	2069.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3132	\N	1	3021	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3131	\N	1	3020	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3130	\N	1	3019	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3129	\N	1	3018	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3128	\N	1	3017	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3127	\N	1	3016	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3126	\N	1	3015	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3125	\N	1	3014	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3124	\N	1	3013	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3123	\N	1	3012	\N	\N	2017-03-29 00:00:00+07	1588.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3122	\N	1	3011	\N	\N	2017-03-29 00:00:00+07	0817.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3121	\N	1	3010	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3120	\N	1	3009	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3119	\N	1	3008	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
44	\N	2	552	\N	\N	2011-02-01 00:00:00+07	0552.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3118	\N	1	3007	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3117	\N	1	3006	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3116	\N	1	3005	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3115	\N	1	3004	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5998	\N	1	2368	\N	\N	2020-01-13 00:00:00+07	2368.65.700	1	\N	\N	OA	\N	10191	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3114	\N	1	3003	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3113	\N	1	3002	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3112	\N	1	3001	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3111	\N	1	3000	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3110	\N	1	2999	\N	\N	2017-03-29 00:00:00+07	0325.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3109	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3108	\N	1	3034	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3107	\N	1	3033	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6007	\N	1	2372	\N	\N	2020-01-20 00:00:00+07	2372.65.700	1	\N	\N	SA	\N	10191	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3106	\N	1	3032	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3105	\N	1	3031	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3104	\N	1	3030	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3103	\N	1	3029	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3102	\N	1	3028	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6017	\N	1	2592	\N	\N	2020-01-23 00:00:00+07	2592.62.700	1	\N	\N	OA	\N	10191	amalia mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3101	\N	1	3027	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3100	\N	1	3026	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3099	\N	1	3025	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3098	\N	1	3024	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3097	\N	1	3023	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3096	\N	1	3022	\N	\N	2017-03-29 00:00:00+07	2069.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3095	\N	1	3021	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3094	\N	1	3020	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3093	\N	1	3019	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3092	\N	1	3018	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3091	\N	1	3017	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3090	\N	1	3016	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3089	\N	1	3015	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3088	\N	1	3014	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3087	\N	1	3013	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3086	\N	1	3012	\N	\N	2017-03-29 00:00:00+07	1588.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3085	\N	1	3011	\N	\N	2017-03-29 00:00:00+07	0817.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3084	\N	1	3010	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3083	\N	1	3009	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3082	\N	1	3008	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3081	\N	1	3007	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3080	\N	1	3006	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3079	\N	1	3005	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3078	\N	1	3004	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6056	\N	1	2384	\N	\N	2020-03-11 00:00:00+07	2384.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3077	\N	1	3003	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3076	\N	1	3002	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3075	\N	1	3001	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3074	\N	1	3000	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6061	\N	1	560	\N	\N	2020-03-13 00:00:00+07	0560.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3073	\N	1	2999	\N	\N	2017-03-29 00:00:00+07	0325.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3072	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3071	\N	1	3034	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3070	\N	1	3033	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3069	\N	1	3032	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3068	\N	1	3031	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3067	\N	1	3030	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3066	\N	1	3029	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6073	\N	1	843	\N	\N	2020-04-14 00:00:00+07	0843.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
43	\N	2	550	\N	\N	2011-02-01 00:00:00+07	0550.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6078	\N	1	847	\N	\N	2020-04-28 00:00:00+07	0847.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3065	\N	1	3028	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3064	\N	1	3027	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3063	\N	1	3026	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3062	\N	1	3025	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3061	\N	1	3024	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3060	\N	1	3023	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3059	\N	1	3022	\N	\N	2017-03-29 00:00:00+07	2069.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3058	\N	1	3021	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3057	\N	1	3020	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3056	\N	1	3019	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3055	\N	1	3018	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3054	\N	1	3017	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3053	\N	1	3016	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3052	\N	1	3015	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6101	\N	1	852	\N	\N	2020-06-17 00:00:00+07	0852.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3051	\N	1	3014	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3050	\N	1	3013	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3049	\N	1	3012	\N	\N	2017-03-29 00:00:00+07	1588.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3048	\N	1	3011	\N	\N	2017-03-29 00:00:00+07	0817.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3047	\N	1	3010	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3046	\N	1	3009	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3045	\N	1	3008	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3044	\N	1	3007	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3043	\N	1	3006	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3042	\N	1	3005	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3041	\N	1	3004	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3040	\N	1	3003	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3039	\N	1	3002	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3038	\N	1	3001	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3037	\N	1	3000	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3036	\N	1	2999	\N	\N	2017-03-29 00:00:00+07	0325.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3035	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3034	\N	1	3034	\N	\N	2017-03-29 00:00:00+07	0820.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3033	\N	1	3033	\N	\N	2017-03-29 00:00:00+07	0814.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3032	\N	1	3032	\N	\N	2017-03-29 00:00:00+07	2057.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3031	\N	1	3031	\N	\N	2017-03-29 00:00:00+07	0819.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3030	\N	1	3030	\N	\N	2017-03-29 00:00:00+07	0813.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3029	\N	1	3029	\N	\N	2017-03-29 00:00:00+07	0800.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3028	\N	1	3028	\N	\N	2017-03-29 00:00:00+07	2070.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3027	\N	1	3027	\N	\N	2017-03-29 00:00:00+07	2058.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3026	\N	1	3026	\N	\N	2017-03-29 00:00:00+07	0818.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3025	\N	1	3025	\N	\N	2017-03-29 00:00:00+07	0816.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3024	\N	1	3024	\N	\N	2017-03-29 00:00:00+07	0812.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3023	\N	1	3023	\N	\N	2017-03-29 00:00:00+07	0801.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3022	\N	1	3022	\N	\N	2017-03-29 00:00:00+07	2069.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3021	\N	1	3021	\N	\N	2017-03-29 00:00:00+07	2066.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3020	\N	1	3020	\N	\N	2017-03-29 00:00:00+07	2065.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3019	\N	1	3019	\N	\N	2017-03-29 00:00:00+07	2064.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3018	\N	1	3018	\N	\N	2017-03-29 00:00:00+07	2063.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3017	\N	1	3017	\N	\N	2017-03-29 00:00:00+07	2062.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3016	\N	1	3016	\N	\N	2017-03-29 00:00:00+07	2061.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3015	\N	1	3015	\N	\N	2017-03-29 00:00:00+07	2060.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3014	\N	1	3014	\N	\N	2017-03-29 00:00:00+07	2059.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3013	\N	1	3013	\N	\N	2017-03-29 00:00:00+07	1593.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3012	\N	1	3012	\N	\N	2017-03-29 00:00:00+07	1588.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3011	\N	1	3011	\N	\N	2017-03-29 00:00:00+07	0817.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3010	\N	1	3010	\N	\N	2017-03-29 00:00:00+07	0815.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3009	\N	1	3009	\N	\N	2017-03-29 00:00:00+07	0811.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3008	\N	1	3008	\N	\N	2017-03-29 00:00:00+07	0810.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3007	\N	1	3007	\N	\N	2017-03-29 00:00:00+07	0807.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3006	\N	1	3006	\N	\N	2017-03-29 00:00:00+07	0806.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3005	\N	1	3005	\N	\N	2017-03-29 00:00:00+07	0805.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3004	\N	1	3004	\N	\N	2017-03-29 00:00:00+07	0803.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3003	\N	1	3003	\N	\N	2017-03-29 00:00:00+07	0802.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2982	\N	1	804	\N	\N	2017-02-01 00:00:00+07	0804.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2976	\N	2	692	\N	\N	2017-03-23 00:00:00+07	0692.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3002	\N	1	3002	\N	\N	2017-03-29 00:00:00+07	0503.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3001	\N	1	3001	\N	\N	2017-03-29 00:00:00+07	0501.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3000	\N	1	3000	\N	\N	2017-03-29 00:00:00+07	0500.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6189	\N	1	2429	\N	\N	2020-09-09 00:00:00+07	2429.65.700	1	\N	\N	OA	\N	10191	solusiti	\N	\N	f	f	2020-09-09 00:00:00+07	\N	00:00	12:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2970	\N	1	458	\N	\N	2017-01-01 00:00:00+07	0458.64.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6196	\N	1	2625	\N	\N	2020-09-10 00:00:00+07	2625.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2020-10-09 00:00:00+07	\N	00:00	23:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6197	\N	1	2626	\N	\N	2020-09-10 00:00:00+07	2626.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2020-10-09 00:00:00+07	\N	00:00	23:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3505	\N	2	830	\N	\N	2016-12-01 00:00:00+07	0830.61.700	2	2017-12-12 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3456	\N	1	819	\N	\N	2017-03-01 00:00:00+07	0819.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3470	\N	1	800	\N	\N	2017-03-01 00:00:00+07	0800.61.700	1	\N	2020-06-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3987	\N	2	368	\N	\N	2017-09-12 00:00:00+07	0368.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3855	\N	2	2167	\N	\N	2017-07-12 00:00:00+07	2167.62.700	2	2018-01-01 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3469	\N	1	801	\N	\N	2017-03-01 00:00:00+07	0801.61.700	1	\N	2020-06-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3468	\N	1	802	\N	\N	2017-03-01 00:00:00+07	0802.61.700	1	\N	2020-05-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3455	\N	1	818	\N	\N	2017-02-01 00:00:00+07	0818.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3454	\N	1	817	\N	\N	2017-03-01 00:00:00+07	0817.61.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3962	\N	2	2185	\N	\N	2017-08-28 00:00:00+07	2185.62.700	1	\N	\N	SA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2999	\N	1	2999	\N	\N	2017-03-29 00:00:00+07	0325.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2998	\N	1	2998	\N	\N	2017-03-29 00:00:00+07	0324.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6257	\N	1	2635	\N	\N	2020-11-03 00:00:00+07	2635.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2020-11-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6316	\N	1	1037	\N	\N	2020-12-08 00:00:00+07	1037.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2020-12-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6317	\N	1	2458	\N	\N	2020-12-08 00:00:00+07	2458.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2020-12-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
42	\N	2	549	\N	\N	2011-02-01 00:00:00+07	0549.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
41	\N	2	548	\N	\N	2011-02-01 00:00:00+07	0548.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
40	\N	2	547	\N	\N	2011-02-01 00:00:00+07	0547.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
39	\N	2	546	\N	\N	2011-02-01 00:00:00+07	0546.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
38	\N	2	545	\N	\N	2011-02-01 00:00:00+07	0545.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
37	\N	2	544	\N	\N	2011-02-01 00:00:00+07	0544.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6378	\N	1	602	\N	\N	2021-02-09 00:00:00+07	0602.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-02-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6379	\N	1	1046	\N	\N	2021-02-09 00:00:00+07	1046.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-02-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6383	\N	1	2476	\N	\N	2021-02-15 00:00:00+07	2476.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-02-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6399	\N	1	883	\N	\N	2021-03-10 00:00:00+07	0883.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-03-10 00:00:00+07	\N	07:00	17:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
36	\N	2	543	\N	\N	2011-02-01 00:00:00+07	0543.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6434	\N	1	2689	\N	\N	2021-04-14 00:00:00+07	2689.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-04-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
35	\N	2	542	\N	\N	2011-02-01 00:00:00+07	0542.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
34	\N	2	541	\N	\N	2011-02-01 00:00:00+07	0541.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7448	\N	1	713	\N	\N	2022-06-08 00:00:00+07	0713.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-06-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7449	\N	1	1024	\N	\N	2022-06-08 00:00:00+07	1024.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-06-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6451	\N	1	2496	\N	\N	2021-04-20 00:00:00+07	2496.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-04-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6458	\N	1	1053	\N	\N	2021-04-28 00:00:00+07	1053.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-04-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6459	\N	1	1054	\N	\N	2021-04-29 00:00:00+07	1054.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-04-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1	\N	2	497	\N	\N	2011-02-01 00:00:00+07	0497.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6463	\N	1	2696	\N	\N	2021-05-03 00:00:00+07	2696.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-05-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
33	\N	2	540	\N	\N	2011-02-01 00:00:00+07	0540.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
32	\N	2	539	\N	\N	2011-02-01 00:00:00+07	0539.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
31	\N	2	538	\N	\N	2011-02-01 00:00:00+07	0538.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
30	\N	2	537	\N	\N	2011-02-01 00:00:00+07	0537.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
29	\N	2	536	\N	\N	2011-02-01 00:00:00+07	0536.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
28	\N	2	535	\N	\N	2011-02-01 00:00:00+07	0535.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6479	\N	1	2505	\N	\N	2021-05-20 00:00:00+07	2505.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-05-20 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6492	\N	1	613	\N	\N	2021-06-10 00:00:00+07	0613.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-06-10 00:00:00+07	\N	08:00	16:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6493	\N	1	895	\N	\N	2021-06-11 00:00:00+07	0895.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-06-11 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6504	\N	1	614	\N	\N	2021-06-17 00:00:00+07	0614.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-06-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
27	\N	2	534	\N	\N	2011-02-01 00:00:00+07	0534.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
26	\N	2	533	\N	\N	2011-02-01 00:00:00+07	0533.61.700	2	2018-06-07 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6512	\N	1	616	\N	\N	2021-06-24 00:00:00+07	0616.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-06-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
25	\N	2	532	\N	\N	2011-02-01 00:00:00+07	0532.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
24	\N	2	531	\N	\N	2011-02-01 00:00:00+07	0531.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
23	\N	2	529	\N	\N	2011-02-01 00:00:00+07	0529.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6518	\N	1	901	\N	\N	2021-06-28 00:00:00+07	0901.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-06-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
22	\N	2	528	\N	\N	2011-02-01 00:00:00+07	0528.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
21	\N	2	526	\N	\N	2011-02-01 00:00:00+07	0526.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
20	\N	2	524	\N	\N	2011-02-01 00:00:00+07	0524.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
19	\N	2	522	\N	\N	2011-02-01 00:00:00+07	0522.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6523	\N	1	2519	\N	\N	2021-06-30 00:00:00+07	2519.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-06-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6526	\N	1	904	\N	\N	2021-06-30 00:00:00+07	0904.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-06-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
18	\N	2	521	\N	\N	2011-02-01 00:00:00+07	0521.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6528	\N	1	1060	\N	\N	2021-07-01 00:00:00+07	1060.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-07-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6531	\N	1	1061	\N	\N	2021-07-02 00:00:00+07	1061.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-07-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
17	\N	2	520	\N	\N	2011-02-01 00:00:00+07	0520.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6667	\N	1	623	\N	\N	2021-08-18 00:00:00+07	0623.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
16	\N	2	519	\N	\N	2011-02-01 00:00:00+07	0519.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6672	\N	1	909	\N	\N	2021-08-23 00:00:00+07	0909.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
15	\N	2	518	\N	\N	2011-02-01 00:00:00+07	0518.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6675	\N	1	1065	\N	\N	2021-08-23 00:00:00+07	1065.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7467	\N	1	714	\N	\N	2022-06-15 00:00:00+07	0714.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-06-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6677	\N	1	911	\N	\N	2021-08-23 00:00:00+07	0911.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6683	\N	1	627	\N	\N	2021-08-25 00:00:00+07	0627.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
14	\N	2	517	\N	\N	2011-02-01 00:00:00+07	0517.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6688	\N	1	1068	\N	\N	2021-08-26 00:00:00+07	1068.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
419	\N	2	1996	\N	\N	2016-10-11 00:00:00+07	1996.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6691	\N	1	2527	\N	\N	2021-08-26 00:00:00+07	2527.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7469	\N	1	1028	\N	\N	2022-06-15 00:00:00+07	1028.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-06-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6693	\N	1	1070	\N	\N	2021-08-27 00:00:00+07	1070.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6694	\N	1	914	\N	\N	2021-08-30 00:00:00+07	0914.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6695	\N	1	915	\N	\N	2021-08-30 00:00:00+07	0915.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
418	\N	2	775	\N	\N	2016-05-24 00:00:00+07	0775.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
417	\N	2	1956	\N	\N	2016-02-10 00:00:00+07	1956.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
416	\N	2	1365	\N	\N	2015-11-03 00:00:00+07	1365.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
415	\N	2	874	\N	\N	2015-10-29 00:00:00+07	0874.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
13	\N	2	516	\N	\N	2011-02-01 00:00:00+07	0516.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6703	\N	1	1072	\N	\N	2021-09-02 00:00:00+07	1072.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-09-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
414	\N	2	451	\N	\N	2015-10-16 00:00:00+07	0451.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
413	\N	2	719	\N	\N	2015-02-05 00:00:00+07	0719.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
410	\N	2	685	\N	\N	2014-03-04 00:00:00+07	0685.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
408	\N	2	238	\N	\N	2013-10-02 00:00:00+07	0238.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
406	\N	2	677	\N	\N	2013-05-15 00:00:00+07	0677.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
12	\N	2	515	\N	\N	2011-02-01 00:00:00+07	0515.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
11	\N	2	513	\N	\N	2011-02-01 00:00:00+07	0513.61.700	2	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6715	\N	1	2534	\N	\N	2021-09-09 00:00:00+07	2534.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
10	\N	2	511	\N	\N	2011-02-01 00:00:00+07	0511.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
405	\N	2	1636	\N	\N	2012-04-12 00:00:00+07	1636.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6718	\N	1	2536	\N	\N	2021-09-09 00:00:00+07	2536.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-09-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
404	\N	2	1624	\N	\N	2012-02-08 00:00:00+07	1624.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
403	\N	2	563	\N	\N	2012-01-06 00:00:00+07	0563.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
402	\N	2	345	\N	\N	2012-01-05 00:00:00+07	0345.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
401	\N	2	327	\N	\N	2012-01-05 00:00:00+07	0327.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
400	\N	2	333	\N	\N	2012-01-05 00:00:00+07	0333.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
399	\N	2	332	\N	\N	2012-01-05 00:00:00+07	0332.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
9	\N	2	510	\N	\N	2011-02-01 00:00:00+07	0510.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
398	\N	2	331	\N	\N	2012-01-05 00:00:00+07	0331.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
396	\N	2	325	\N	\N	2012-01-05 00:00:00+07	0325.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
8	\N	2	507	\N	\N	2011-02-01 00:00:00+07	0507.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7	\N	2	504	\N	\N	2011-02-01 00:00:00+07	0504.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6759	\N	1	2541	\N	\N	2021-09-16 00:00:00+07	2541.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-09-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
395	\N	2	324	\N	\N	2012-01-05 00:00:00+07	0324.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
394	\N	2	323	\N	\N	2012-01-05 00:00:00+07	0323.64.700	2	2018-06-25 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
393	\N	2	322	\N	\N	2012-01-05 00:00:00+07	0322.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6763	\N	1	1078	\N	\N	2021-09-17 00:00:00+07	1078.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-09-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7480	\N	1	716	\N	\N	2022-06-16 00:00:00+07	0716.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-06-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
392	\N	2	206	\N	\N	2012-01-05 00:00:00+07	0206.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
391	\N	2	205	\N	\N	2012-01-05 00:00:00+07	0205.63.700	2	2017-07-03 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
390	\N	2	203	\N	\N	2012-01-05 00:00:00+07	0203.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
389	\N	2	201	\N	\N	2012-01-05 00:00:00+07	0201.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
388	\N	2	1526	\N	\N	2012-01-05 00:00:00+07	1526.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6772	\N	1	1080	\N	\N	2021-09-22 00:00:00+07	1080.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-09-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6773	\N	1	2546	\N	\N	2021-09-22 00:00:00+07	2546.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-09-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
387	\N	2	1525	\N	\N	2012-01-05 00:00:00+07	1525.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
386	\N	2	1523	\N	\N	2012-01-05 00:00:00+07	1523.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
385	\N	2	1522	\N	\N	2012-01-05 00:00:00+07	1522.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
384	\N	2	1521	\N	\N	2012-01-05 00:00:00+07	1521.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
383	\N	2	1517	\N	\N	2012-01-05 00:00:00+07	1517.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
382	\N	2	1518	\N	\N	2012-01-05 00:00:00+07	1518.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
381	\N	2	1516	\N	\N	2012-01-05 00:00:00+07	1516.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
380	\N	2	1515	\N	\N	2012-01-05 00:00:00+07	1515.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
379	\N	2	612	\N	\N	2012-01-05 00:00:00+07	0612.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
378	\N	2	611	\N	\N	2012-01-05 00:00:00+07	0611.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
377	\N	2	610	\N	\N	2012-01-05 00:00:00+07	0610.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
376	\N	2	609	\N	\N	2012-01-05 00:00:00+07	0609.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-16 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
375	\N	2	608	\N	\N	2012-01-05 00:00:00+07	0608.61.700	1	\N	\N	SA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
374	\N	2	606	\N	\N	2012-01-05 00:00:00+07	0606.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
373	\N	2	602	\N	\N	2012-01-05 00:00:00+07	0602.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
372	\N	2	1520	\N	\N	2012-01-05 00:00:00+07	1520.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
371	\N	2	1529	\N	\N	2012-01-04 00:00:00+07	1529.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
370	\N	2	542	\N	\N	2012-01-02 00:00:00+07	0542.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
369	\N	2	1536	\N	\N	2012-01-02 00:00:00+07	1536.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6	\N	2	503	\N	\N	2011-02-01 00:00:00+07	0503.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6796	\N	1	638	\N	\N	2021-09-30 00:00:00+07	0638.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-09-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5	\N	2	502	\N	\N	2011-02-01 00:00:00+07	0502.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
368	\N	2	172	\N	\N	2012-01-02 00:00:00+07	0172.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
367	\N	2	547	\N	\N	2012-01-02 00:00:00+07	0547.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
366	\N	2	535	\N	\N	2012-01-02 00:00:00+07	0535.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
365	\N	2	534	\N	\N	2012-01-02 00:00:00+07	0534.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
364	\N	2	1421	\N	\N	2012-01-02 00:00:00+07	1421.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
363	\N	2	1503	\N	\N	2012-01-02 00:00:00+07	1503.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
362	\N	2	593	\N	\N	2012-01-02 00:00:00+07	0593.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
361	\N	2	579	\N	\N	2012-01-02 00:00:00+07	0579.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
360	\N	2	1510	\N	\N	2012-01-02 00:00:00+07	1510.62.700	2	2018-05-01 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
359	\N	2	590	\N	\N	2012-01-02 00:00:00+07	0590.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
358	\N	2	1506	\N	\N	2012-01-02 00:00:00+07	1506.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4	\N	2	500	\N	\N	2011-02-01 00:00:00+07	0500.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3	\N	2	499	\N	\N	2011-02-01 00:00:00+07	0499.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
357	\N	2	565	\N	\N	2012-01-02 00:00:00+07	0565.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
356	\N	2	1535	\N	\N	2012-01-02 00:00:00+07	1535.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
355	\N	2	1533	\N	\N	2012-01-02 00:00:00+07	1533.62.700	1	\N	2020-06-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-07-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
354	\N	2	1532	\N	\N	2012-01-02 00:00:00+07	1532.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
353	\N	2	626	\N	\N	2012-01-02 00:00:00+07	0626.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
352	\N	2	560	\N	\N	2012-01-02 00:00:00+07	0560.65.700	1	\N	2020-06-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
351	\N	2	181	\N	\N	2012-01-02 00:00:00+07	0181.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
350	\N	2	506	\N	\N	2012-01-02 00:00:00+07	0506.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6819	\N	1	642	\N	\N	2021-10-14 00:00:00+07	0642.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-10-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
349	\N	2	514	\N	\N	2012-01-02 00:00:00+07	0514.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
348	\N	2	505	\N	\N	2012-01-02 00:00:00+07	0505.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
347	\N	2	523	\N	\N	2012-01-02 00:00:00+07	0523.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
346	\N	2	585	\N	\N	2012-01-02 00:00:00+07	0585.61.700	1	\N	2020-06-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
345	\N	2	1413	\N	\N	2012-01-02 00:00:00+07	1413.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6825	\N	1	2566	\N	\N	2021-10-18 00:00:00+07	2566.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-10-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7051	\N	1	963	\N	\N	2022-01-06 00:00:00+07	0963.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-01-06 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
196	\N	2	188	\N	\N	2012-01-02 00:00:00+07	0188.63.700	1	2017-10-12 00:00:00+07	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
195	\N	2	187	\N	\N	2012-01-02 00:00:00+07	0187.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
194	\N	2	186	\N	\N	2012-01-02 00:00:00+07	0186.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
193	\N	2	184	\N	\N	2012-01-02 00:00:00+07	0184.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
192	\N	2	182	\N	\N	2012-01-02 00:00:00+07	0182.63.700	2	2019-11-14 00:00:00+07	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
191	\N	2	180	\N	\N	2012-01-02 00:00:00+07	0180.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
190	\N	2	179	\N	\N	2012-01-02 00:00:00+07	0179.63.700	1	2018-06-28 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
189	\N	2	178	\N	\N	2012-01-02 00:00:00+07	0178.63.700	2	2018-06-28 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
188	\N	2	177	\N	\N	2012-01-02 00:00:00+07	0177.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
187	\N	2	176	\N	\N	2012-01-02 00:00:00+07	0176.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
186	\N	2	175	\N	\N	2012-01-02 00:00:00+07	0175.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
185	\N	2	174	\N	\N	2012-01-02 00:00:00+07	0174.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
184	\N	2	173	\N	\N	2012-01-02 00:00:00+07	0173.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
183	\N	2	171	\N	\N	2012-01-02 00:00:00+07	0171.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
182	\N	2	170	\N	\N	2012-01-02 00:00:00+07	0170.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
181	\N	2	169	\N	\N	2012-01-02 00:00:00+07	0169.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
180	\N	2	168	\N	\N	2012-01-02 00:00:00+07	0168.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
179	\N	2	167	\N	\N	2012-01-02 00:00:00+07	0167.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
178	\N	2	166	\N	\N	2012-01-02 00:00:00+07	0166.63.700	2	2020-05-18 00:00:00+07	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
177	\N	2	165	\N	\N	2012-01-02 00:00:00+07	0165.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
176	\N	2	163	\N	\N	2012-01-02 00:00:00+07	0163.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
175	\N	2	162	\N	\N	2012-01-02 00:00:00+07	0162.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
174	\N	2	161	\N	\N	2012-01-02 00:00:00+07	0161.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-05-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
173	\N	2	1528	\N	\N	2012-01-02 00:00:00+07	1528.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7175	\N	1	683	\N	\N	2022-01-26 00:00:00+07	0683.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-01-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
172	\N	2	1550	\N	\N	2012-01-02 00:00:00+07	1550.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
171	\N	2	1537	\N	\N	2012-01-02 00:00:00+07	1537.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
170	\N	2	1514	\N	\N	2012-01-02 00:00:00+07	1514.62.700	0	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7207	\N	1	2660	\N	\N	2022-01-31 00:00:00+07	2660.65.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-01-31 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
169	\N	2	1513	\N	\N	2012-01-02 00:00:00+07	1513.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
168	\N	2	1508	\N	\N	2012-01-02 00:00:00+07	1508.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
167	\N	2	1507	\N	\N	2012-01-02 00:00:00+07	1507.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
166	\N	2	1504	\N	\N	2012-01-02 00:00:00+07	1504.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
165	\N	2	1502	\N	\N	2012-01-02 00:00:00+07	1502.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
164	\N	2	1501	\N	\N	2012-01-02 00:00:00+07	1501.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-24 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
163	\N	2	1498	\N	\N	2012-01-02 00:00:00+07	1498.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
162	\N	2	1497	\N	\N	2012-01-02 00:00:00+07	1497.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
161	\N	2	1495	\N	\N	2012-01-02 00:00:00+07	1495.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7221	\N	1	989	\N	\N	2022-02-04 00:00:00+07	0989.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-02-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
160	\N	2	1494	\N	\N	2012-01-02 00:00:00+07	1494.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
159	\N	2	1493	\N	\N	2012-01-02 00:00:00+07	1493.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
158	\N	2	1491	\N	\N	2012-01-02 00:00:00+07	1491.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
157	\N	2	1490	\N	\N	2012-01-02 00:00:00+07	1490.62.700	0	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
156	\N	2	1489	\N	\N	2012-01-02 00:00:00+07	1489.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
155	\N	2	1488	\N	\N	2012-01-02 00:00:00+07	1488.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
154	\N	2	1487	\N	\N	2012-01-02 00:00:00+07	1487.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
153	\N	2	1482	\N	\N	2012-01-02 00:00:00+07	1482.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
152	\N	2	1481	\N	\N	2012-01-02 00:00:00+07	1481.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
151	\N	2	1480	\N	\N	2012-01-02 00:00:00+07	1480.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
150	\N	2	1479	\N	\N	2012-01-02 00:00:00+07	1479.62.700	1	2013-01-01 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
149	\N	2	1478	\N	\N	2012-01-02 00:00:00+07	1478.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
148	\N	2	1477	\N	\N	2012-01-02 00:00:00+07	1477.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
147	\N	2	1475	\N	\N	2012-01-02 00:00:00+07	1475.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
146	\N	2	1474	\N	\N	2012-01-02 00:00:00+07	1474.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
145	\N	2	1473	\N	\N	2012-01-02 00:00:00+07	1473.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
144	\N	2	1472	\N	\N	2012-01-02 00:00:00+07	1472.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
143	\N	2	1471	\N	\N	2012-01-02 00:00:00+07	1471.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
142	\N	2	1470	\N	\N	2012-01-02 00:00:00+07	1470.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
141	\N	2	1469	\N	\N	2012-01-02 00:00:00+07	1469.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
140	\N	2	1467	\N	\N	2012-01-02 00:00:00+07	1467.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
139	\N	2	1465	\N	\N	2012-01-02 00:00:00+07	1465.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
138	\N	2	1464	\N	\N	2012-01-02 00:00:00+07	1464.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
137	\N	2	1463	\N	\N	2012-01-02 00:00:00+07	1463.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
136	\N	2	1462	\N	\N	2012-01-02 00:00:00+07	1462.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
135	\N	2	1461	\N	\N	2012-01-02 00:00:00+07	1461.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
134	\N	2	1460	\N	\N	2012-01-02 00:00:00+07	1460.62.700	2	2017-06-16 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
133	\N	2	1458	\N	\N	2012-01-02 00:00:00+07	1458.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
132	\N	2	1457	\N	\N	2012-01-02 00:00:00+07	1457.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-05-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
131	\N	2	1456	\N	\N	2012-01-02 00:00:00+07	1456.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
130	\N	2	1455	\N	\N	2012-01-02 00:00:00+07	1455.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
129	\N	2	1451	\N	\N	2012-01-02 00:00:00+07	1451.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
128	\N	2	1450	\N	\N	2012-01-02 00:00:00+07	1450.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
127	\N	2	1449	\N	\N	2012-01-02 00:00:00+07	1449.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
126	\N	2	1448	\N	\N	2012-01-02 00:00:00+07	1448.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
125	\N	2	1447	\N	\N	2012-01-02 00:00:00+07	1447.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
124	\N	2	1446	\N	\N	2012-01-02 00:00:00+07	1446.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7289	\N	1	1003	\N	\N	2022-03-04 00:00:00+07	1003.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-03-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
123	\N	2	1445	\N	\N	2012-01-02 00:00:00+07	1445.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
122	\N	2	1444	\N	\N	2012-01-02 00:00:00+07	1444.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
121	\N	2	1443	\N	\N	2012-01-02 00:00:00+07	1443.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-25 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
120	\N	2	1442	\N	\N	2012-01-02 00:00:00+07	1442.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
119	\N	2	1441	\N	\N	2012-01-02 00:00:00+07	1441.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
118	\N	2	1440	\N	\N	2012-01-02 00:00:00+07	1440.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
117	\N	2	1435	\N	\N	2012-01-02 00:00:00+07	1435.62.700	0	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
116	\N	2	1434	\N	\N	2012-01-02 00:00:00+07	1434.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
115	\N	2	1431	\N	\N	2012-01-02 00:00:00+07	1431.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
114	\N	2	1430	\N	\N	2012-01-02 00:00:00+07	1430.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
113	\N	2	1429	\N	\N	2012-01-02 00:00:00+07	1429.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
112	\N	2	1426	\N	\N	2012-01-02 00:00:00+07	1426.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
111	\N	2	1420	\N	\N	2012-01-02 00:00:00+07	1420.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
110	\N	2	1414	\N	\N	2012-01-02 00:00:00+07	1414.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
109	\N	2	1411	\N	\N	2012-01-02 00:00:00+07	1411.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7333	\N	1	694	\N	\N	2022-03-17 00:00:00+07	0694.63.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-03-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
108	\N	2	1409	\N	\N	2012-01-02 00:00:00+07	1409.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
107	\N	2	1408	\N	\N	2012-01-02 00:00:00+07	1408.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
106	\N	2	1407	\N	\N	2012-01-02 00:00:00+07	1407.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
105	\N	2	1406	\N	\N	2012-01-02 00:00:00+07	1406.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
104	\N	2	1404	\N	\N	2012-01-02 00:00:00+07	1404.62.700	2	2018-10-22 00:00:00+07	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
103	\N	2	1402	\N	\N	2012-01-02 00:00:00+07	1402.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7343	\N	1	1008	\N	\N	2022-03-22 00:00:00+07	1008.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-03-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7344	\N	1	1009	\N	\N	2022-03-22 00:00:00+07	1009.64.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-03-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
102	\N	2	1401	\N	\N	2012-01-02 00:00:00+07	1401.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
101	\N	2	1400	\N	\N	2012-01-02 00:00:00+07	1400.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
100	\N	2	628	\N	\N	2012-01-02 00:00:00+07	0628.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
99	\N	2	584	\N	\N	2012-01-02 00:00:00+07	0584.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
98	\N	2	629	\N	\N	2012-01-02 00:00:00+07	0629.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
97	\N	2	622	\N	\N	2012-01-02 00:00:00+07	0622.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
96	\N	2	621	\N	\N	2012-01-02 00:00:00+07	0621.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
95	\N	2	619	\N	\N	2012-01-02 00:00:00+07	0619.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-17 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
94	\N	2	618	\N	\N	2012-01-02 00:00:00+07	0618.61.700	2	2021-05-01 00:00:00+07	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-08-02 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
93	\N	2	617	\N	\N	2012-01-02 00:00:00+07	0617.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
92	\N	2	614	\N	\N	2012-01-02 00:00:00+07	0614.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
91	\N	2	616	\N	\N	2012-01-02 00:00:00+07	0616.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
90	\N	2	613	\N	\N	2012-01-02 00:00:00+07	0613.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
89	\N	2	599	\N	\N	2012-01-02 00:00:00+07	0599.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
88	\N	2	598	\N	\N	2012-01-02 00:00:00+07	0598.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
87	\N	2	597	\N	\N	2012-01-02 00:00:00+07	0597.61.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2020-09-22 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7370	\N	1	1016	\N	\N	2022-04-05 00:00:00+07	1016.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-04-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
86	\N	2	596	\N	\N	2012-01-02 00:00:00+07	0596.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
85	\N	2	595	\N	\N	2012-01-02 00:00:00+07	0595.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
84	\N	2	594	\N	\N	2012-01-02 00:00:00+07	0594.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
83	\N	2	592	\N	\N	2012-01-02 00:00:00+07	0592.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
82	\N	2	591	\N	\N	2012-01-02 00:00:00+07	0591.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
81	\N	2	589	\N	\N	2012-01-02 00:00:00+07	0589.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
80	\N	2	582	\N	\N	2012-01-02 00:00:00+07	0582.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
79	\N	2	581	\N	\N	2012-01-02 00:00:00+07	0581.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
78	\N	2	580	\N	\N	2012-01-02 00:00:00+07	0580.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
77	\N	2	578	\N	\N	2012-01-02 00:00:00+07	0578.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
76	\N	2	577	\N	\N	2012-01-02 00:00:00+07	0577.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
75	\N	2	576	\N	\N	2012-01-02 00:00:00+07	0576.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
74	\N	2	575	\N	\N	2012-01-02 00:00:00+07	0575.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
73	\N	2	1524	\N	\N	2011-03-08 00:00:00+07	1524.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
72	\N	2	1486	\N	\N	2011-03-01 00:00:00+07	1486.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-30 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7399	\N	1	1017	\N	\N	2022-04-21 00:00:00+07	1017.64.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-04-21 00:00:00+07	\N	08:00	16:00	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
71	\N	2	212	\N	\N	2011-03-01 00:00:00+07	0212.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
70	\N	2	624	\N	\N	2011-03-01 00:00:00+07	0624.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
69	\N	2	564	\N	\N	2011-02-28 00:00:00+07	0564.65.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
68	\N	2	1534	\N	\N	2011-02-01 00:00:00+07	1534.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
67	\N	2	557	\N	\N	2011-02-01 00:00:00+07	0557.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5244	\N	1	474	\N	\N	2019-04-30 00:00:00+07	0474.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5241	\N	1	933	\N	\N	2019-04-30 00:00:00+07	0933.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5240	\N	1	932	\N	\N	2019-04-30 00:00:00+07	0932.61.700	2	2021-10-01 00:00:00+07	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-26 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5236	\N	1	2376	\N	\N	2019-04-26 00:00:00+07	2376.62.700	1	\N	\N	SA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5224	\N	1	929	\N	\N	2019-04-15 00:00:00+07	0929.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5223	\N	1	928	\N	\N	2019-04-15 00:00:00+07	0928.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5290	\N	1	2102	\N	\N	2019-05-14 00:00:00+07	2102.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5222	\N	1	927	\N	\N	2019-04-15 00:00:00+07	0927.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5293	\N	1	2104	\N	\N	2019-05-15 00:00:00+07	2104.65.700	1	\N	2020-07-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5221	\N	1	926	\N	\N	2019-04-15 00:00:00+07	0926.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5295	\N	1	2106	\N	\N	2019-05-15 00:00:00+07	2106.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5189	\N	1	2066	\N	\N	2019-04-05 00:00:00+07	2066.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5178	\N	1	2063	\N	\N	2019-04-04 00:00:00+07	2063.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5299	\N	1	2387	\N	\N	2019-05-16 00:00:00+07	2387.62.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5155	\N	1	916	\N	\N	2019-03-28 00:00:00+07	0916.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5147	\N	1	914	\N	\N	2019-03-26 00:00:00+07	0914.61.700	2	2020-02-01 00:00:00+07	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5146	\N	1	470	\N	\N	2019-03-26 00:00:00+07	0470.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5144	\N	1	913	\N	\N	2019-03-26 00:00:00+07	0913.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5143	\N	1	912	\N	\N	2019-03-26 00:00:00+07	0912.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5142	\N	1	469	\N	\N	2019-03-26 00:00:00+07	0469.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5141	\N	1	468	\N	\N	2019-03-26 00:00:00+07	0468.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5140	\N	1	467	\N	\N	2019-03-26 00:00:00+07	0467.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5308	\N	1	2112	\N	\N	2019-05-16 00:00:00+07	2112.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5138	\N	1	466	\N	\N	2019-03-25 00:00:00+07	0466.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5312	\N	1	479	\N	\N	2019-05-20 00:00:00+07	0479.63.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5127	\N	1	911	\N	\N	2019-03-20 00:00:00+07	0911.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5123	\N	1	909	\N	\N	2019-03-19 00:00:00+07	0909.61.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-01 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5194	\N	1	924	\N	\N	2019-04-08 00:00:00+07	0924.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5122	\N	1	908	\N	\N	2019-03-19 00:00:00+07	0908.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5121	\N	1	907	\N	\N	2019-03-19 00:00:00+07	0907.61.700	2	2019-03-22 00:00:00+07	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5117	\N	1	2036	\N	\N	2019-03-18 00:00:00+07	2036.65.700	1	\N	\N	SA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5193	\N	1	2068	\N	\N	2019-04-08 00:00:00+07	2068.65.700	1	\N	\N	SA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5109	\N	1	464	\N	\N	2019-03-15 00:00:00+07	0464.63.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5107	\N	1	2355	\N	\N	2019-03-15 00:00:00+07	2355.62.700	1	\N	2020-06-01 00:00:00+07	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5192	\N	1	758	\N	\N	2019-04-08 00:00:00+07	0758.64.700	1	\N	\N	SA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5105	\N	1	2030	\N	\N	2019-03-15 00:00:00+07	2030.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5101	\N	1	2026	\N	\N	2019-03-14 00:00:00+07	2026.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5091	\N	1	906	\N	\N	2019-03-12 00:00:00+07	0906.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5088	\N	1	905	\N	\N	2019-03-12 00:00:00+07	0905.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5087	\N	1	904	\N	\N	2019-03-12 00:00:00+07	0904.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5086	\N	1	903	\N	\N	2019-03-12 00:00:00+07	0903.61.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5069	\N	1	752	\N	\N	2019-03-06 00:00:00+07	0752.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5066	\N	1	2349	\N	\N	2019-03-06 00:00:00+07	2349.62.700	1	\N	\N	SA	\N	10191	Amalia Mahardani	\N	\N	f	f	2022-01-12 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5059	\N	1	2000	\N	\N	2019-03-05 00:00:00+07	2000.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5049	\N	1	1992	\N	\N	2019-02-28 00:00:00+07	1992.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4990	\N	1	461	\N	\N	2019-02-15 00:00:00+07	0461.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4984	\N	1	2331	\N	\N	2019-02-14 00:00:00+07	2331.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4981	\N	1	2330	\N	\N	2019-02-14 00:00:00+07	2330.62.700	2	2020-03-01 00:00:00+07	\N	OA	\N	10191	Amalia Mahardani	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5345	\N	1	2122	\N	\N	2019-05-23 00:00:00+07	2122.65.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4980	\N	1	2329	\N	\N	2019-02-14 00:00:00+07	2329.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4979	\N	1	2328	\N	\N	2019-02-14 00:00:00+07	2328.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4975	\N	1	1947	\N	\N	2019-02-13 00:00:00+07	1947.65.700	0	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4973	\N	1	747	\N	\N	2019-02-13 00:00:00+07	0747.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5037	\N	1	1986	\N	\N	2019-02-27 00:00:00+07	1986.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4969	\N	1	459	\N	\N	2019-02-13 00:00:00+07	0459.63.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5353	\N	1	940	\N	\N	2019-05-27 00:00:00+07	0940.61.700	1	\N	\N	OA	\N	10191	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5035	\N	1	2345	\N	\N	2019-02-27 00:00:00+07	2345.62.700	1	\N	\N	OA	\N	10191	Amalia Mahardani	\N	\N	f	f	2021-11-09 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5029	\N	1	748	\N	\N	2019-02-26 00:00:00+07	0748.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
66	\N	2	531	\N	\N	2011-02-01 00:00:00+07	0531.65.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4967	\N	1	745	\N	\N	2019-02-12 00:00:00+07	0745.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4966	\N	1	744	\N	\N	2019-02-12 00:00:00+07	0744.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4965	\N	1	2327	\N	\N	2019-02-12 00:00:00+07	2327.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4964	\N	1	2326	\N	\N	2019-02-12 00:00:00+07	2326.62.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4963	\N	1	743	\N	\N	2019-02-12 00:00:00+07	0743.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4961	\N	1	742	\N	\N	2019-02-12 00:00:00+07	0742.64.700	1	\N	\N	OA	\N	10191	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4212	\N	1	96	\N	\N	2017-10-23 00:00:00+07	0096.00.401	1	\N	\N	OA	\N	10126	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7037	\N	1	959	\N	\N	2022-01-04 00:00:00+07	0959.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7038	\N	1	2795	\N	\N	2022-01-04 00:00:00+07	2795.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7039	\N	1	960	\N	\N	2022-01-04 00:00:00+07	0960.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7040	\N	1	961	\N	\N	2022-01-04 00:00:00+07	0961.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7041	\N	1	2796	\N	\N	2022-01-05 00:00:00+07	2796.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7042	\N	1	66666	\N	\N	2022-01-05 00:00:00+07	66666..401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7043	\N	1	1111	\N	\N	2022-01-05 00:00:00+07	1111..401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7036	\N	1	2794	\N	\N	2022-01-04 00:00:00+07	2794.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7045	\N	1	962	\N	\N	2022-01-05 00:00:00+07	0962.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7046	\N	1	2797	\N	\N	2022-01-05 00:00:00+07	2797.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7047	\N	1	670	\N	\N	2022-01-05 00:00:00+07	0670.63.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7048	\N	1	1113	\N	\N	2022-01-05 00:00:00+07	1113.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7049	\N	1	2635	\N	\N	2022-01-05 00:00:00+07	2635.65.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7050	\N	1	1114	\N	\N	2022-01-05 00:00:00+07	1114.61.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-05 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7355	\N	1	1013	\N	\N	2022-03-29 00:00:00+07	1013.64.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-03-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7035	\N	1	958	\N	\N	2022-01-04 00:00:00+07	0958.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7034	\N	1	669	\N	\N	2022-01-04 00:00:00+07	0669.63.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7356	\N	1	2873	\N	\N	2022-03-29 00:00:00+07	2873.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-03-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3679	\N	1	2103	\N	\N	2017-04-07 00:00:00+07	2103.62.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7032	\N	1	2793	\N	\N	2022-01-04 00:00:00+07	2793.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7030	\N	1	957	\N	\N	2022-01-04 00:00:00+07	0957.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7357	\N	1	1161	\N	\N	2022-03-29 00:00:00+07	1161.61.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-03-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7025	\N	1	2790	\N	\N	2022-01-03 00:00:00+07	2790.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7024	\N	1	2791	\N	\N	2022-01-03 00:00:00+07	2791.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7023	\N	1	1110	\N	\N	2022-01-03 00:00:00+07	1110.61.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7029	\N	1	956	\N	\N	2022-01-04 00:00:00+07	0956.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-04 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7022	\N	1	2789	\N	\N	2022-01-03 00:00:00+07	2789.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-03 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7358	\N	1	1012	\N	\N	2022-03-29 00:00:00+07	1012.64.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-03-29 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3759	\N	2	68	\N	\N	2017-06-02 00:00:00+07	0068.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3696	\N	1	525	\N	\N	2017-04-07 00:00:00+07	0525.62.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3840	\N	2	93	\N	\N	2017-06-22 00:00:00+07	0093.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3841	\N	2	92	\N	\N	2017-06-22 00:00:00+07	0092.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3708	\N	1	2126	\N	\N	2017-04-27 00:00:00+07	2126.62.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7091	\N	1	1115	\N	\N	2022-01-13 00:00:00+07	1115.61.401	1	\N	\N	OA	\N	10125	solusiti	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3848	\N	1	94	\N	\N	2017-07-07 00:00:00+07	0094.00.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3844	\N	1	91	\N	\N	2017-06-22 00:00:00+07	0091.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7092	\N	1	676	\N	\N	2022-01-13 00:00:00+07	0676.63.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7093	\N	1	967	\N	\N	2022-01-13 00:00:00+07	0967.64.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7094	\N	1	2802	\N	\N	2022-01-13 00:00:00+07	2802.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-13 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3803	\N	1	79	\N	\N	2017-06-07 00:00:00+07	0079.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3793	\N	2	846	\N	\N	2017-06-07 00:00:00+07	0846.61.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3830	\N	1	2162	\N	\N	2017-06-21 00:00:00+07	2162.62.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3831	\N	1	560	\N	\N	2017-06-21 00:00:00+07	0560.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3795	\N	2	845	\N	\N	2017-06-07 00:00:00+07	0845.61.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3796	\N	2	71	\N	\N	2017-06-07 00:00:00+07	0071.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3832	\N	2	563	\N	\N	2017-06-21 00:00:00+07	0563.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3798	\N	2	76	\N	\N	2017-06-07 00:00:00+07	0076.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3833	\N	2	562	\N	\N	2017-06-21 00:00:00+07	0562.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3834	\N	1	561	\N	\N	2017-06-21 00:00:00+07	0561.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3799	\N	2	75	\N	\N	2017-06-07 00:00:00+07	0075.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3801	\N	2	77	\N	\N	2017-06-07 00:00:00+07	0077.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7259	\N	1	2846	\N	\N	2022-02-10 00:00:00+07	2846.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-10 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7266	\N	1	2847	\N	\N	2022-02-18 00:00:00+07	2847.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3835	\N	2	90	\N	\N	2017-06-21 00:00:00+07	0090.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3918	\N	1	1641	\N	\N	2017-07-27 00:00:00+07	1641.65.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3678	\N	1	2102	\N	\N	2017-04-07 00:00:00+07	2102.62.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3836	\N	1	89	\N	\N	2017-06-21 00:00:00+07	0089.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3837	\N	1	88	\N	\N	2017-06-22 00:00:00+07	0088.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3802	\N	2	78	\N	\N	2017-06-07 00:00:00+07	0078.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3838	\N	1	87	\N	\N	2017-06-22 00:00:00+07	0087.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7276	\N	1	2851	\N	\N	2022-02-23 00:00:00+07	2851.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7277	\N	1	2850	\N	\N	2022-02-23 00:00:00+07	2850.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7278	\N	1	2849	\N	\N	2022-02-23 00:00:00+07	2849.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7279	\N	1	2848	\N	\N	2022-02-23 00:00:00+07	2848.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-23 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3839	\N	1	354	\N	\N	2017-06-22 00:00:00+07	0354.63.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3818	\N	2	80	\N	\N	2017-06-13 00:00:00+07	0080.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3762	\N	2	69	\N	\N	2017-06-05 00:00:00+07	0069.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3829	\N	2	2161	\N	\N	2017-06-21 00:00:00+07	2161.62.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7112	\N	1	2804	\N	\N	2022-01-14 00:00:00+07	2804.62.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7114	\N	1	2645	\N	\N	2022-01-14 00:00:00+07	2645.65.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7115	\N	1	2646	\N	\N	2022-01-14 00:00:00+07	2646.65.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7116	\N	1	1116	\N	\N	2022-01-14 00:00:00+07	1116.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7117	\N	1	2647	\N	\N	2022-01-14 00:00:00+07	2647.65.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3763	\N	2	70	\N	\N	2017-06-05 00:00:00+07	0070.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3819	\N	2	81	\N	\N	2017-06-13 00:00:00+07	0081.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7120	\N	1	2649	\N	\N	2022-01-14 00:00:00+07	2649.65.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7121	\N	1	970	\N	\N	2022-01-14 00:00:00+07	0970.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7122	\N	1	971	\N	\N	2022-01-14 00:00:00+07	0971.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7123	\N	1	2806	\N	\N	2022-01-14 00:00:00+07	2806.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7124	\N	1	2807	\N	\N	2022-01-14 00:00:00+07	2807.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7125	\N	1	2808	\N	\N	2022-01-14 00:00:00+07	2808.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-14 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3797	\N	2	847	\N	\N	2017-06-07 00:00:00+07	0847.61.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3787	\N	1	2156	\N	\N	2017-06-07 00:00:00+07	2156.62.401	1	\N	\N	OA	\N	10125	solusiti	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5862	\N	1	2330	\N	\N	2019-10-29 00:00:00+07	2330.65.401	1	\N	\N	SA	\N	10125	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3700	\N	1	2107	\N	\N	2017-05-01 00:00:00+07	2107.62.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5861	\N	1	2329	\N	\N	2019-10-29 00:00:00+07	2329.65.401	1	\N	\N	SA	\N	10125	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5860	\N	1	2328	\N	\N	2019-10-29 00:00:00+07	2328.65.401	1	\N	\N	SA	\N	10125	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7137	\N	1	2653	\N	\N	2022-01-19 00:00:00+07	2653.65.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-19 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5859	\N	1	2327	\N	\N	2019-10-29 00:00:00+07	2327.65.401	1	\N	\N	SA	\N	10125	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5858	\N	1	2542	\N	\N	2019-10-29 00:00:00+07	2542.62.401	1	\N	\N	SA	\N	10125	NANANG SWEISTINURA	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3489	\N	1	522	\N	\N	2017-04-25 00:00:00+07	0522.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3701	\N	1	2108	\N	\N	2017-05-01 00:00:00+07	2108.62.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3820	\N	2	82	\N	\N	2017-06-13 00:00:00+07	0082.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3715	\N	2	2134	\N	\N	2017-05-02 00:00:00+07	2134.62.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3745	\N	2	67	\N	\N	2017-05-24 00:00:00+07	0067.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3767	\N	2	72	\N	\N	2017-06-05 00:00:00+07	0072.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3490	\N	1	522	\N	\N	2016-12-01 00:00:00+07	`0522.64.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7149	\N	1	976	\N	\N	2022-01-21 00:00:00+07	0976.64.401	1	\N	\N	SA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7150	\N	1	2815	\N	\N	2022-01-21 00:00:00+07	2815.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7151	\N	1	2816	\N	\N	2022-01-21 00:00:00+07	2816.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7152	\N	1	2817	\N	\N	2022-01-21 00:00:00+07	2817.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7153	\N	1	2818	\N	\N	2022-01-21 00:00:00+07	2818.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7154	\N	1	2819	\N	\N	2022-01-21 00:00:00+07	2819.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7155	\N	1	2820	\N	\N	2022-01-21 00:00:00+07	2820.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7156	\N	1	1119	\N	\N	2022-01-21 00:00:00+07	1119.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7157	\N	1	1120	\N	\N	2022-01-21 00:00:00+07	1120.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7158	\N	1	1121	\N	\N	2022-01-21 00:00:00+07	1121.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7159	\N	1	2821	\N	\N	2022-01-21 00:00:00+07	2821.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7160	\N	1	1122	\N	\N	2022-01-21 00:00:00+07	1122.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7161	\N	1	1123	\N	\N	2022-01-21 00:00:00+07	1123.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7162	\N	1	1124	\N	\N	2022-01-21 00:00:00+07	1124.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7163	\N	1	977	\N	\N	2022-01-21 00:00:00+07	0977.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7164	\N	1	978	\N	\N	2022-01-21 00:00:00+07	0978.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7165	\N	1	979	\N	\N	2022-01-21 00:00:00+07	0979.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7166	\N	1	980	\N	\N	2022-01-21 00:00:00+07	0980.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7167	\N	1	680	\N	\N	2022-01-21 00:00:00+07	0680.63.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7168	\N	1	1125	\N	\N	2022-01-21 00:00:00+07	1125.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-21 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3770	\N	2	73	\N	\N	2017-06-05 00:00:00+07	0073.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3772	\N	2	549	\N	\N	2017-06-06 00:00:00+07	0549.65.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3773	\N	1	550	\N	\N	2017-06-06 00:00:00+07	0550.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3920	\N	1	95	\N	\N	2017-07-27 00:00:00+07	0095.00.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3774	\N	2	552	\N	\N	2017-06-06 00:00:00+07	0552.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7183	\N	1	2658	\N	\N	2022-01-28 00:00:00+07	2658.65.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7184	\N	1	2659	\N	\N	2022-01-28 00:00:00+07	2659.65.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3775	\N	2	553	\N	\N	2017-06-06 00:00:00+07	0553.62.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3776	\N	2	551	\N	\N	2017-06-06 00:00:00+07	0551.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7180	\N	1	1127	\N	\N	2022-01-27 00:00:00+07	1127.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-27 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3641	\N	1	1603	\N	\N	2017-04-07 00:00:00+07	1603.65.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7185	\N	1	1128	\N	\N	2022-01-28 00:00:00+07	1128.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7186	\N	1	1129	\N	\N	2022-01-28 00:00:00+07	1129.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7187	\N	1	1130	\N	\N	2022-01-28 00:00:00+07	1130.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7188	\N	1	1131	\N	\N	2022-01-28 00:00:00+07	1131.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7189	\N	1	982	\N	\N	2022-01-28 00:00:00+07	0982.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7190	\N	1	983	\N	\N	2022-01-28 00:00:00+07	0983.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7191	\N	1	984	\N	\N	2022-01-28 00:00:00+07	0984.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7192	\N	1	985	\N	\N	2022-01-28 00:00:00+07	0985.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7193	\N	1	986	\N	\N	2022-01-28 00:00:00+07	0986.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7194	\N	1	987	\N	\N	2022-01-28 00:00:00+07	0987.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7195	\N	1	2826	\N	\N	2022-01-28 00:00:00+07	2826.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7196	\N	1	2827	\N	\N	2022-01-28 00:00:00+07	2827.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7197	\N	1	2828	\N	\N	2022-01-28 00:00:00+07	2828.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7199	\N	1	2830	\N	\N	2022-01-28 00:00:00+07	2830.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7200	\N	1	2831	\N	\N	2022-01-28 00:00:00+07	2831.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7201	\N	1	2832	\N	\N	2022-01-28 00:00:00+07	2832.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7202	\N	1	2833	\N	\N	2022-01-28 00:00:00+07	2833.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7203	\N	1	2834	\N	\N	2022-01-28 00:00:00+07	2834.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7204	\N	1	2835	\N	\N	2022-01-28 00:00:00+07	2835.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7205	\N	1	2836	\N	\N	2022-01-28 00:00:00+07	2836.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7206	\N	1	2837	\N	\N	2022-01-28 00:00:00+07	2837.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-01-28 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3780	\N	2	350	\N	\N	2017-06-06 00:00:00+07	0350.63.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3781	\N	2	351	\N	\N	2017-06-06 00:00:00+07	0351.63.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3782	\N	2	352	\N	\N	2017-06-06 00:00:00+07	0352.63.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3783	\N	1	353	\N	\N	2017-06-06 00:00:00+07	0353.63.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3784	\N	2	554	\N	\N	2017-06-07 00:00:00+07	0554.64.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3785	\N	2	2154	\N	\N	2017-06-07 00:00:00+07	2154.62.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3786	\N	1	2155	\N	\N	2017-06-07 00:00:00+07	2155.62.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3788	\N	1	2157	\N	\N	2017-06-07 00:00:00+07	2157.62.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3789	\N	2	842	\N	\N	2017-06-07 00:00:00+07	0842.61.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3898	\N	1	2168	\N	\N	2017-07-25 00:00:00+07	2168.62.401	2	2018-02-21 00:00:00+07	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7224	\N	1	2841	\N	\N	2022-02-07 00:00:00+07	2841.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4204	\N	1	595	\N	\N	2017-10-19 00:00:00+07	0595.64.401	1	\N	\N	SA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7226	\N	1	1136	\N	\N	2022-02-07 00:00:00+07	1136.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7227	\N	1	1137	\N	\N	2022-02-07 00:00:00+07	1137.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7228	\N	1	2664	\N	\N	2022-02-07 00:00:00+07	2664.65.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7229	\N	1	2665	\N	\N	2022-02-07 00:00:00+07	2665.65.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7230	\N	1	1138	\N	\N	2022-02-07 00:00:00+07	1138.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7231	\N	1	1138	\N	\N	2022-02-07 00:00:00+07	1138.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7232	\N	1	1138	\N	\N	2022-02-07 00:00:00+07	1138.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7233	\N	1	1139	\N	\N	2022-02-07 00:00:00+07	1139.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7234	\N	1	991	\N	\N	2022-02-07 00:00:00+07	0991.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7235	\N	1	992	\N	\N	2022-02-07 00:00:00+07	0992.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7236	\N	1	993	\N	\N	2022-02-07 00:00:00+07	0993.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7237	\N	1	994	\N	\N	2022-02-07 00:00:00+07	0994.64.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-07 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7238	\N	1	2842	\N	\N	2022-02-08 00:00:00+07	2842.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7239	\N	1	2843	\N	\N	2022-02-08 00:00:00+07	2843.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7240	\N	1	2844	\N	\N	2022-02-08 00:00:00+07	2844.62.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7242	\N	1	1140	\N	\N	2022-02-08 00:00:00+07	1140.61.401	1	\N	\N	OA	\N	10125	Amalia Mahardani	\N	\N	f	f	2022-02-08 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3790	\N	2	843	\N	\N	2017-06-07 00:00:00+07	0843.61.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3791	\N	1	844	\N	\N	2017-06-07 00:00:00+07	0844.61.401	1	\N	\N	OA	\N	10125	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3529	\N	1	3462	\N	\N	2017-04-03 00:00:00+07	3462.62.311	2	2017-06-08 00:00:00+07	\N	SA	\N	10157	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3448	\N	1	3030	\N	\N	2017-03-01 00:00:00+07	3030.62.311	2	2017-03-31 00:00:00+07	\N	SA	\N	10157	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3615	\N	1	3962	\N	\N	2017-05-02 00:00:00+07	3962.62.311	2	2017-05-31 00:00:00+07	\N	SA	\N	10157	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3602	\N	1	3662	\N	\N	2017-05-02 00:00:00+07	3662.62.311	2	2017-06-08 00:00:00+07	\N	SA	\N	10157	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3567	\N	1	3562	\N	\N	2017-05-10 00:00:00+07	3562.62.311	2	2017-05-31 00:00:00+07	\N	SA	\N	10157	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3609	\N	1	3762	\N	\N	2017-05-02 00:00:00+07	3762.62.311	2	2017-06-08 00:00:00+07	\N	SA	\N	10157	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4121	\N	1	75	\N	\N	2003-01-01 00:00:00+07	0075.65.302	1	\N	\N	SA	\N	10213	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4070	\N	1	428	\N	\N	2003-01-01 00:00:00+07	0428.61.302	1	\N	\N	SA	\N	10213	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4120	\N	1	36	\N	\N	2008-08-01 00:00:00+07	0036.63.302	1	\N	\N	SA	\N	10213	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4122	\N	1	84	\N	\N	2005-09-01 00:00:00+07	0084.65.302	1	\N	\N	SA	\N	10213	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7322	\N	1	2689	\N	\N	2022-03-15 00:00:00+07	2689.65.312	1	\N	\N	SA	\N	10156	Amalia Mahardani	\N	\N	f	f	2022-03-15 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6915	\N	1	654	\N	\N	2021-11-18 00:00:00+07	0654.63.312	1	\N	\N	SA	\N	10156	Amalia Mahardani	\N	\N	f	f	2021-11-18 00:00:00+07	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4123	\N	1	77	\N	\N	2009-02-01 00:00:00+07	0077..65.302	1	\N	\N	SA	\N	10213	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4089	\N	1	428	\N	\N	2003-01-01 00:00:00+07	0428..61.302	1	\N	\N	SA	\N	10213	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
\.


--
-- Name: Npwpd_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."Npwpd_Id_seq"', 7487, false);


--
-- PostgreSQL database dump complete
--

