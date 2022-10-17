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
-- Data for Name: User; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."User" ("Id", "Name", "Salt", "Password", "Position", "Ref_Id", "Group_Id", "Email", "Notes", "SysAdmin", "ValidPeriod", "VeifyStatus", "RegMode", "Status", "ApprovedAt", "InactiveAt", "CreatedAt", "UpdatedAt", "DeletedAt") FROM stdin;
1	adminold	1374199120	1bd0ae15e69076eea82f5affa5226e0a88856cff	\N	\N	1	email@bapenda.com	\N	t	\N	\N	\N	1	\N	\N	\N	\N	\N
2	admin	\N	$2a$10$qp8CA3obI5LD12gbzdkWeuvGLHJUy316iWEuT8GlmjkO8hAb1e.si	1	\N	\N	admin@gmail.com	\N	t	0001-01-01 07:07:12+07:07:12	1	1	1	\N	\N	2022-10-17 10:39:14.673588+07	2022-10-17 10:39:14.673588+07	\N
18	penerimaan	1481784948	e1dfb8e7aa0b363f4e69dfb1dc4c25c356592555	\N	\N	106	email2@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
74	migrasi	1533792484	f48541d4348283384a98e42029c0c777edf1fd32	\N	\N	1	email3@bapenda.com	\N	t	\N	\N	\N	1	\N	\N	\N	\N	\N
20	ismu	1487490807	22981c67f53ede6d673d65384a998a4875b1e5e7	\N	\N	1	email4@bapenda.com	\N	t	\N	\N	\N	1	\N	\N	\N	\N	\N
36	zainal	1490154613	3d3b3777ea1ec869fe73db60d21399d4700d59a9	\N	\N	1	email5@bapenda.com	\N	t	\N	\N	\N	1	\N	\N	\N	\N	\N
37	Made	1490321059	37dca53ccff064d74a059142a3e08e7150b6c993	\N	\N	38	email6@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
13	admin2	1458805347	baed57252d54fb955f3e6d6a226dae5e70544b41	\N	\N	1	email7@bapenda.com	\N	t	\N	\N	\N	1	\N	\N	\N	\N	\N
14	herdian	1481776343	a81d578a6391f63b75993633bc62bab26c484d91	\N	\N	1	email8@bapenda.com	\N	t	\N	\N	\N	1	\N	\N	\N	\N	\N
75	pemeriksaan	1537249566	cd7c3a3c3c8c5cd1b684a19cf52dae365e2cd146	\N	\N	68	email9@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
38	ANANG	1490844311	92458204cffbcc1c2e80b608f0ca595588a3cda3	\N	\N	106	email10@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
39	RETNO	1490844390	1c82c8dd7c9e5b7d700029a1be8785c09bcf4413	\N	\N	68	email11@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
40	ANELIES	1490844458	d1caaa51367890008a3482a11390ff7dab8bcce5	\N	\N	68	email12@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
41	JOKO	1490844541	4640e0ed29d1fce2dc55ad450a3091738c9c7c42	\N	\N	43	email13@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
44	SBETTASARI	1493084349	db64c4441fae4bf9168f4a95c16327c751d931c2	\N	\N	106	email14@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
45	PUJI UTAMI	1495002069	f6e7e6b268dd8fe3ce599aefb11d0202c46cace1	\N	\N	106	email15@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
46	MARIYANTO	1496710979	50a181243e24d2032e758fbb384e13d0286d5432	\N	\N	43	email16@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
47	CHUSNUL ABDILAH	1497245412	89215e311649b21529c6eb4903e760c45449bae3	\N	\N	106	email17@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
43	eka	1492046561	a732ede3cd319977a5d3d16cbdc65354d99a8ff8	\N	\N	5	email18@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
27	wiwik	1488855150	b4ddbbf249ac2882f228417c68b1d8d4926f5351	\N	\N	196	email19@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
48	NILUH	1502171704	412e26e3769bf2f3c75f136341f5423f6dd24f67	\N	\N	46	email20@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
28	prek	1488855445	cd340d64633a4e13e9ae084221dd264fd7702e06	\N	\N	69	email21@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
49	SWEIST	1502246499	428dd5cc0f66fcf7d2314f204ce39e27fcb64081	\N	\N	3	email22@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
30	wiwit	1488855583	e0f87053f4dccaecf2b26601742756b491df6fa8	\N	\N	41	email23@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
31	tita	1488855706	932fe65b24d5f6ef4281dfe6d8f1728d2f8a28a0	\N	\N	204	email24@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
50	sugeng	1504165972	9a9044aa2793efd17536363d8449c67d793c44da	\N	\N	68	email25@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
32	nurul	1488855770	c14e421ecc4316b0d372d082d722f18df305a946	\N	\N	43	email26@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
51	jazilatul	1504166030	7573d983d1b10cdbde30b7f3287132795cb73765	\N	\N	68	email27@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
33	hariyono	1488856060	813385eba6988266f02e9593b1c22d3dd5011eef	\N	\N	175	email28@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
34	nengah	1488856124	1b408e0720702dfb497c58b6449b1ac8ebbd2657	\N	\N	175	email29@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
35	solikin	1489108929	69d74b74619644c2da5a2e13f053031b4a3f84c9	\N	\N	6	email30@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
52	hari	1504166077	9cb2d03ce791f25214a447466715f8a65e836546	\N	\N	68	email31@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
53	mariani	1504166172	853a4668fb3dd2cbe5af97c62eb37e5421ce930f	\N	\N	68	email32@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
54	Dull	1504166213	f94b0f0eab1c1c236f4cb4c8f5fe2febe5663482	\N	\N	68	email33@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
55	ida	1504166264	45bfd4523b3682e230e9711223d20e2c590c6194	\N	\N	68	email34@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
56	irma	1504166306	27a13abffb6050b31e75f4b0d09c83f8758ec3ee	\N	\N	106	email35@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
57	UPTKLOJEN	1505703919	3df84b8e89d420625119c13b74e6603018e8f2c2	\N	\N	69	email36@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
58	UPTKEDUNGKANDANG	1505713105	e0ce10f4fa24580d5ebb9817b241d2dc9eab6978	\N	\N	69	email37@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
59	UPTKD	1506390665	bc7eafce5c399486bf7734a6543bcd09940e6092	\N	\N	69	email38@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
60	UPTBLIMBING	1506390868	386b41991efb0e7a8ebdea20c5397ff6fc9a4a7c	\N	\N	69	email39@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
61	UPTLW	1506391257	1e89357811eb1cf36636fb90dda1a3b6e6af7e09	\N	\N	69	email40@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
62	UPTSUKUN	1506391336	c91785b6d94202cd0821d133b1b7f1c6210da9eb	\N	\N	69	email41@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
72	reklame	1524707513	972fa806146e748d8ae2277634da542434ce5935	\N	\N	69	email42@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
66	LITA	1517370594	3ade7ff6bfa53af6b218d12d0957bf1d560ae23e	\N	\N	4	email43@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
67	GILANG	1517388310	ecb15a097f688905d4f247653f51bfdd97c5c6e0	\N	\N	46	email44@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
68	Igha	1519620691	f37d37495d7128cf4306c1480d634292c567ef36	\N	\N	106	email45@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
70	iga	1519962637	c9a6499049815afdb83acb1c53e377283c8eab16	\N	\N	43	email46@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
71	IGA2	1520228324	003c50c62b28ca8500ce5873c359ecfdeac8dc51	\N	\N	43	email47@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
73	amalia	1530582220	24c59b3702b67b8bf5838aec764d4b1e47b64e3e	\N	\N	180	email48@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
76	tika	1541121305	8390fad973cc32629a4bc25548412c627b68a079	\N	\N	69	email49@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
77	rochma	1541121568	b5a61ac17bbff0ab27ddb31edbe535db0b06cd5c	\N	\N	69	email50@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
78	DIYAH	1541988146	e6f349db10f30b3c99f526d5dd2c8fdebd08edcd	\N	\N	6	email51@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
79	reklame123	1545902500	010c97c445b4de741af3816929337e7ea646cbfb	\N	\N	126	email52@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
80	genio	1551337900	ef04f350ef0a8d1c351d5320cd658785272f8e3f	\N	\N	69	email53@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
81	ARBI	1552632008	23b7fbff09c6d186b9a83e2a44694dd5eb3e24d2	\N	\N	69	email54@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
83	YUNIARSO	1592961090	825e58612d06dcf9a9ba1e272cf3aa1c86700143	\N	\N	43	email55@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
84	BENNI	1553827266	8c9a7497934f5fe32997379a40a237ce55fea447	\N	\N	69	email56@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
85	sameko	1555309029	1f79b2026bf0ea472fb113b3e69e881f23b19668	\N	\N	6	email57@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
86	helis	1559007725	4d96bc79dad93eba39e03ea1d363cd793c015940	\N	\N	69	email58@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
87	sanditia	1564537652	253e0771840fb48f745eeb0ca502d40650738781	\N	\N	43	email74@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
88	ilma	1564537861	99989127db23d36343cd19dd3cfc0229585a4229	\N	\N	40	email59@bapenda.com	\N	f	\N	\N	\N	0	\N	\N	\N	\N	\N
89	kpk	1568770517	9ae7dd518368a20bf234803ea83f16a89915377b	\N	\N	170	email60@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
90	diyah.	1578038521	b6215d52c1db0386c7330ef33c195a3b67f1dbf8	\N	\N	6	email61@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
91	anik	1578362704	f5220ca42151f5bbc3c208a835e05c2556c0b5a6	\N	\N	175	email62@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
92	dimas	1578534426	16bd7b38a136f84c7aa8f3cbbec477e9a371093f	\N	\N	3	email63@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
93	naufal	1578534504	23683dbbfe69b5b70a32e214a651c29e0fda9c1e	\N	\N	3	email64@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
94	nurulz	1578884124	3cb65b79467d0a5516d063f97882c6188650975f	\N	\N	6	email65@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
95	fauzan	1580453200	faf40093f1557f9c26c98a6f9902738dd2caa5f7	\N	\N	184	email66@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
96	diankuntari	1592963651	129a1a27cbe6c99009a38417a6cbc548fda271ab	\N	\N	189	email67@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
97	mike	1613364725	8f07cd7a3b91b4fbb83230b6cefa1f2cda99dd23	\N	\N	194	email68@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
98	agus	1616040811	3ec3a7ffdc49bbc0f6a10c628dfd212ffc9840b5	\N	\N	43	email69@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
99	alfiah	1616384976	ea6f38ade1cbf5830c8496723504c2e8a3180dc6	\N	\N	195	email70@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
100	penagihan1	1617076304	e987d0d7163cd8298abbbf4dd68fc1a38426dbde	\N	\N	68	email71@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
101	daeng	1624595049	df6dfb4239df6a354c4fe3dd076115112b79b985	\N	\N	200	email72@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
102	inspektorat	1654139552	fdc602132f36654a59e2a9513904319ab3b2ea7b	\N	\N	213	email73@bapenda.com	\N	f	\N	\N	\N	1	\N	\N	\N	\N	\N
\.


--
-- Name: User_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."User_Id_seq"', 1, false);


--
-- PostgreSQL database dump complete
--

