--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-2.pgdg20.04+2)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-2.pgdg20.04+2)

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
-- Data for Name: DbkbFasumDepMinMax; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."DbkbFasumDepMinMax" ("Id", "Provinsi_Kode", "Daerah_Kode", "Tahun", "Fasilitas_Kode", "KlsDepMin", "KlsDepMax", "Nilai", "CreatedAt", "UpdatedAt", "DeletedAt") FROM stdin;
1	35	73	2005	30	0	4	168572.00	\N	\N	\N
2	35	73	2005	30	5	9	351193.00	\N	\N	\N
3	35	73	2005	30	10	19	491671.00	\N	\N	\N
4	35	73	2005	30	20	99	632149.00	\N	\N	\N
5	35	73	2005	31	0	4	561911.00	\N	\N	\N
6	35	73	2005	31	5	9	842865.00	\N	\N	\N
7	35	73	2005	31	10	19	1194059.00	\N	\N	\N
8	35	73	2005	31	20	99	1404776.00	\N	\N	\N
9	35	73	2005	32	0	4	126429.00	\N	\N	\N
10	35	73	2005	32	5	9	210716.00	\N	\N	\N
11	35	73	2005	32	10	19	393337.00	\N	\N	\N
12	35	73	2005	32	20	99	561911.00	\N	\N	\N
13	35	73	2005	40	0	99	905.00	\N	\N	\N
14	35	73	2005	40	100	249	671.00	\N	\N	\N
15	35	73	2005	40	250	499	581.00	\N	\N	\N
16	35	73	2005	40	500	999999	517.00	\N	\N	\N
17	35	73	2007	12	0	50	934.00	\N	\N	\N
18	35	73	2007	12	51	100	847.00	\N	\N	\N
19	35	73	2007	12	101	200	620.00	\N	\N	\N
20	35	73	2007	12	201	400	590.00	\N	\N	\N
21	35	73	2007	12	401	999999	562.00	\N	\N	\N
22	35	73	2007	13	0	50	1018.00	\N	\N	\N
23	35	73	2007	13	51	100	930.00	\N	\N	\N
24	35	73	2007	13	101	200	705.00	\N	\N	\N
25	35	73	2007	13	201	400	674.00	\N	\N	\N
26	35	73	2007	13	401	999999	646.00	\N	\N	\N
27	35	73	2007	30	0	4	168572.00	\N	\N	\N
28	35	73	2014	12	0	50	934.00	\N	\N	\N
29	35	73	2014	12	51	100	847.00	\N	\N	\N
30	35	73	2014	12	101	200	620.00	\N	\N	\N
31	35	73	2014	12	201	400	590.00	\N	\N	\N
32	35	73	2014	12	401	999999	562.00	\N	\N	\N
33	35	73	2014	13	0	50	1018.00	\N	\N	\N
34	35	73	2014	13	51	100	930.00	\N	\N	\N
35	35	73	2014	13	101	200	705.00	\N	\N	\N
36	35	73	2014	13	201	400	674.00	\N	\N	\N
37	35	73	2014	13	401	999999	646.00	\N	\N	\N
38	35	73	2014	30	0	4	168572.00	\N	\N	\N
39	35	73	2014	30	5	9	351193.00	\N	\N	\N
40	35	73	2014	30	10	19	491671.00	\N	\N	\N
41	35	73	2014	30	20	99	632149.00	\N	\N	\N
42	35	73	2014	31	0	4	561911.00	\N	\N	\N
43	35	73	2014	31	5	9	842865.00	\N	\N	\N
44	35	73	2014	31	10	19	1194059.00	\N	\N	\N
45	35	73	2014	31	20	99	1404776.00	\N	\N	\N
46	35	73	2014	32	0	4	126429.00	\N	\N	\N
47	35	73	2014	32	5	9	210716.00	\N	\N	\N
48	35	73	2014	32	10	19	393337.00	\N	\N	\N
49	35	73	2014	32	20	99	561911.00	\N	\N	\N
50	35	73	2014	40	0	99	905.00	\N	\N	\N
51	35	73	2014	40	100	249	671.00	\N	\N	\N
52	35	73	2014	40	250	499	581.00	\N	\N	\N
53	35	73	2014	40	500	999999	517.00	\N	\N	\N
54	35	73	2015	12	0	50	934.00	\N	\N	\N
55	35	73	2015	12	51	100	847.00	\N	\N	\N
56	35	73	2015	12	101	200	620.00	\N	\N	\N
57	35	73	2015	12	201	400	590.00	\N	\N	\N
58	35	73	2015	12	401	999999	562.00	\N	\N	\N
59	35	73	2015	13	0	50	1018.00	\N	\N	\N
60	35	73	2015	13	51	100	930.00	\N	\N	\N
61	35	73	2015	13	101	200	705.00	\N	\N	\N
62	35	73	2015	13	201	400	674.00	\N	\N	\N
63	35	73	2015	13	401	999999	646.00	\N	\N	\N
64	35	73	2015	30	0	4	168572.00	\N	\N	\N
65	35	73	2015	30	5	9	351193.00	\N	\N	\N
66	35	73	2015	30	10	19	491671.00	\N	\N	\N
67	35	73	2015	30	20	99	632149.00	\N	\N	\N
68	35	73	2015	31	0	4	561911.00	\N	\N	\N
69	35	73	2015	31	5	9	842865.00	\N	\N	\N
70	35	73	2015	31	10	19	1194059.00	\N	\N	\N
71	35	73	2015	31	20	99	1404776.00	\N	\N	\N
72	35	73	2015	32	0	4	126429.00	\N	\N	\N
73	35	73	2015	32	5	9	210716.00	\N	\N	\N
74	35	73	2015	32	10	19	393337.00	\N	\N	\N
75	35	73	2015	32	20	99	561911.00	\N	\N	\N
76	35	73	2015	40	0	99	905.00	\N	\N	\N
77	35	73	2015	40	100	249	671.00	\N	\N	\N
78	35	73	2015	40	250	499	581.00	\N	\N	\N
79	35	73	2015	40	500	999999	517.00	\N	\N	\N
80	35	73	2016	12	0	50	934.00	\N	\N	\N
81	35	73	2016	12	51	100	847.00	\N	\N	\N
82	35	73	2016	12	101	200	620.00	\N	\N	\N
83	35	73	2016	12	201	400	590.00	\N	\N	\N
84	35	73	2016	12	401	999999	562.00	\N	\N	\N
85	35	73	2016	13	0	50	1018.00	\N	\N	\N
86	35	73	2016	13	51	100	930.00	\N	\N	\N
87	35	73	2016	13	101	200	705.00	\N	\N	\N
88	35	73	2016	13	201	400	674.00	\N	\N	\N
89	35	73	2016	13	401	999999	646.00	\N	\N	\N
90	35	73	2016	30	0	4	168572.00	\N	\N	\N
91	35	73	2016	30	5	9	351193.00	\N	\N	\N
92	35	73	2016	30	10	19	491671.00	\N	\N	\N
93	35	73	2016	30	20	99	632149.00	\N	\N	\N
94	35	73	2016	31	0	4	561911.00	\N	\N	\N
95	35	73	2016	31	5	9	842865.00	\N	\N	\N
96	35	73	2016	31	10	19	1194059.00	\N	\N	\N
97	35	73	2016	31	20	99	1404776.00	\N	\N	\N
98	35	73	2016	32	0	4	126429.00	\N	\N	\N
99	35	73	2016	32	5	9	210716.00	\N	\N	\N
100	35	73	2016	32	10	19	393337.00	\N	\N	\N
101	35	73	2016	32	20	99	561911.00	\N	\N	\N
102	35	73	2016	40	0	99	905.00	\N	\N	\N
103	35	73	2016	40	100	249	671.00	\N	\N	\N
104	35	73	2016	40	250	499	581.00	\N	\N	\N
105	35	73	2016	40	500	999999	517.00	\N	\N	\N
106	35	73	2017	12	0	50	934.00	\N	\N	\N
107	35	73	2017	12	51	100	847.00	\N	\N	\N
108	35	73	2017	12	101	200	620.00	\N	\N	\N
109	35	73	2017	12	201	400	590.00	\N	\N	\N
110	35	73	2017	12	401	999999	562.00	\N	\N	\N
111	35	73	2017	13	0	50	1018.00	\N	\N	\N
112	35	73	2017	13	51	100	930.00	\N	\N	\N
113	35	73	2017	13	101	200	705.00	\N	\N	\N
114	35	73	2017	13	201	400	674.00	\N	\N	\N
115	35	73	2017	13	401	999999	646.00	\N	\N	\N
116	35	73	2017	30	0	4	168572.00	\N	\N	\N
117	35	73	2017	30	5	9	351193.00	\N	\N	\N
118	35	73	2017	30	10	19	491671.00	\N	\N	\N
119	35	73	2017	30	20	99	632149.00	\N	\N	\N
120	35	73	2017	31	0	4	561911.00	\N	\N	\N
121	35	73	2017	31	5	9	842865.00	\N	\N	\N
122	35	73	2017	31	10	19	1194059.00	\N	\N	\N
123	35	73	2017	31	20	99	1404776.00	\N	\N	\N
124	35	73	2017	32	0	4	126429.00	\N	\N	\N
125	35	73	2017	32	5	9	210716.00	\N	\N	\N
126	35	73	2017	32	10	19	393337.00	\N	\N	\N
127	35	73	2017	32	20	99	561911.00	\N	\N	\N
128	35	73	2017	40	0	99	905.00	\N	\N	\N
129	35	73	2017	40	100	249	671.00	\N	\N	\N
130	35	73	2017	40	250	499	581.00	\N	\N	\N
131	35	73	2017	40	500	999999	517.00	\N	\N	\N
132	35	73	2007	30	5	9	351193.00	\N	\N	\N
133	35	73	2007	30	10	19	491671.00	\N	\N	\N
134	35	73	2007	30	20	99	632149.00	\N	\N	\N
135	35	73	2007	31	0	4	561911.00	\N	\N	\N
136	35	73	2007	31	5	9	842865.00	\N	\N	\N
137	35	73	2007	31	10	19	1194059.00	\N	\N	\N
138	35	73	2006	12	0	50	934.00	\N	\N	\N
139	35	73	2006	12	51	100	847.00	\N	\N	\N
140	35	73	2006	12	101	200	620.00	\N	\N	\N
141	35	73	2006	12	201	400	590.00	\N	\N	\N
142	35	73	2006	12	401	999999	562.00	\N	\N	\N
143	35	73	2006	13	0	50	1018.00	\N	\N	\N
144	35	73	2006	13	51	100	930.00	\N	\N	\N
145	35	73	2006	13	101	200	705.00	\N	\N	\N
146	35	73	2006	13	201	400	674.00	\N	\N	\N
147	35	73	2006	13	401	999999	646.00	\N	\N	\N
148	35	73	2006	30	0	4	168572.00	\N	\N	\N
149	35	73	2006	30	5	9	351193.00	\N	\N	\N
150	35	73	2006	30	10	19	491671.00	\N	\N	\N
151	35	73	2006	30	20	99	632149.00	\N	\N	\N
152	35	73	2006	31	0	4	561911.00	\N	\N	\N
153	35	73	2006	31	5	9	842865.00	\N	\N	\N
154	35	73	2006	31	10	19	1194059.00	\N	\N	\N
155	35	73	2006	31	20	99	1404776.00	\N	\N	\N
156	35	73	2006	32	0	4	126429.00	\N	\N	\N
157	35	73	2006	32	5	9	210716.00	\N	\N	\N
158	35	73	2006	32	10	19	393337.00	\N	\N	\N
159	35	73	2006	32	20	99	561911.00	\N	\N	\N
160	35	73	2006	40	0	99	905.00	\N	\N	\N
161	35	73	2006	40	100	249	671.00	\N	\N	\N
162	35	73	2006	40	250	499	581.00	\N	\N	\N
163	35	73	2006	40	500	999999	517.00	\N	\N	\N
164	35	73	2007	31	20	99	1404776.00	\N	\N	\N
165	35	73	2007	32	0	4	126429.00	\N	\N	\N
166	35	73	2007	32	5	9	210716.00	\N	\N	\N
167	35	73	2007	32	10	19	393337.00	\N	\N	\N
168	35	73	2007	32	20	99	561911.00	\N	\N	\N
169	35	73	2007	40	0	99	905.00	\N	\N	\N
170	35	73	2007	40	100	249	671.00	\N	\N	\N
171	35	73	2007	40	250	499	581.00	\N	\N	\N
172	35	73	2007	40	500	999999	517.00	\N	\N	\N
173	35	73	2008	12	0	50	934.00	\N	\N	\N
174	35	73	2008	12	51	100	847.00	\N	\N	\N
175	35	73	2008	12	101	200	620.00	\N	\N	\N
176	35	73	2008	12	201	400	590.00	\N	\N	\N
177	35	73	2008	12	401	999999	562.00	\N	\N	\N
178	35	73	2008	13	0	50	1018.00	\N	\N	\N
179	35	73	2008	13	51	100	930.00	\N	\N	\N
180	35	73	2008	13	101	200	705.00	\N	\N	\N
181	35	73	2008	13	201	400	674.00	\N	\N	\N
182	35	73	2008	13	401	999999	646.00	\N	\N	\N
183	35	73	2008	30	0	4	168572.00	\N	\N	\N
184	35	73	2008	30	5	9	351193.00	\N	\N	\N
185	35	73	2008	30	10	19	491671.00	\N	\N	\N
186	35	73	2008	30	20	99	632149.00	\N	\N	\N
187	35	73	2008	31	0	4	561911.00	\N	\N	\N
188	35	73	2008	31	5	9	842865.00	\N	\N	\N
189	35	73	2008	31	10	19	1194059.00	\N	\N	\N
190	35	73	2008	31	20	99	1404776.00	\N	\N	\N
191	35	73	2008	32	0	4	126429.00	\N	\N	\N
192	35	73	2008	32	5	9	210716.00	\N	\N	\N
193	35	73	2008	32	10	19	393337.00	\N	\N	\N
194	35	73	2008	32	20	99	561911.00	\N	\N	\N
195	35	73	2008	40	0	99	905.00	\N	\N	\N
196	35	73	2008	40	100	249	671.00	\N	\N	\N
197	35	73	2008	40	250	499	581.00	\N	\N	\N
198	35	73	2008	40	500	999999	517.00	\N	\N	\N
199	35	73	2000	12	0	50	934.00	\N	\N	\N
200	35	73	2000	12	51	100	847.00	\N	\N	\N
201	35	73	2000	12	101	200	620.00	\N	\N	\N
202	35	73	2000	12	201	400	590.00	\N	\N	\N
203	35	73	2000	12	401	999999	562.00	\N	\N	\N
204	35	73	2000	13	0	50	1018.00	\N	\N	\N
205	35	73	2000	13	51	100	930.00	\N	\N	\N
206	35	73	2000	13	101	200	705.00	\N	\N	\N
207	35	73	2000	13	201	400	674.00	\N	\N	\N
208	35	73	2000	13	401	999999	646.00	\N	\N	\N
209	35	73	2000	30	0	4	168572.00	\N	\N	\N
210	35	73	2000	30	5	9	351193.00	\N	\N	\N
211	35	73	2000	30	10	19	491671.00	\N	\N	\N
212	35	73	2000	30	20	99	632149.00	\N	\N	\N
213	35	73	2000	31	0	4	561911.00	\N	\N	\N
214	35	73	2000	31	5	9	842865.00	\N	\N	\N
215	35	73	2000	31	10	19	1194059.00	\N	\N	\N
216	35	73	2000	31	20	99	1404776.00	\N	\N	\N
217	35	73	2000	32	0	4	126429.00	\N	\N	\N
218	35	73	2000	32	5	9	210716.00	\N	\N	\N
219	35	73	2000	32	10	19	393337.00	\N	\N	\N
220	35	73	2000	32	20	99	561911.00	\N	\N	\N
221	35	73	2000	40	0	99	905.00	\N	\N	\N
222	35	73	2000	40	100	249	671.00	\N	\N	\N
223	35	73	2000	40	250	499	581.00	\N	\N	\N
224	35	73	2000	40	500	999999	517.00	\N	\N	\N
225	35	73	1999	12	0	50	869.00	\N	\N	\N
226	35	73	1999	12	51	100	788.00	\N	\N	\N
227	35	73	1999	12	101	200	577.00	\N	\N	\N
228	35	73	1999	12	201	400	549.00	\N	\N	\N
229	35	73	1999	12	401	999999	523.00	\N	\N	\N
230	35	73	1999	13	0	50	947.00	\N	\N	\N
231	35	73	1999	13	51	100	866.00	\N	\N	\N
232	35	73	1999	13	101	200	656.00	\N	\N	\N
233	35	73	1999	13	201	400	627.00	\N	\N	\N
234	35	73	1999	13	401	999999	601.00	\N	\N	\N
235	35	73	1999	30	0	4	156812.00	\N	\N	\N
236	35	73	1999	30	5	9	326692.00	\N	\N	\N
237	35	73	1999	30	10	19	457369.00	\N	\N	\N
238	35	73	1999	30	20	99	588046.00	\N	\N	\N
239	35	73	1999	31	0	4	522708.00	\N	\N	\N
240	35	73	2022	12	0	50	934.00	\N	\N	\N
241	35	73	2022	12	51	100	847.00	\N	\N	\N
242	35	73	2022	12	101	200	620.00	\N	\N	\N
243	35	73	2022	12	201	400	590.00	\N	\N	\N
244	35	73	2022	12	401	999999	562.00	\N	\N	\N
245	35	73	2022	13	0	50	1018.00	\N	\N	\N
246	35	73	2022	13	51	100	930.00	\N	\N	\N
247	35	73	2022	13	101	200	705.00	\N	\N	\N
248	35	73	2022	13	201	400	674.00	\N	\N	\N
249	35	73	2022	13	401	999999	646.00	\N	\N	\N
250	35	73	2022	30	0	4	168572.00	\N	\N	\N
251	35	73	2022	30	5	9	351193.00	\N	\N	\N
252	35	73	2022	30	10	19	491671.00	\N	\N	\N
253	35	73	2022	30	20	99	632149.00	\N	\N	\N
254	35	73	2022	31	0	4	561911.00	\N	\N	\N
255	35	73	2022	31	5	9	842865.00	\N	\N	\N
256	35	73	2022	31	10	19	1194059.00	\N	\N	\N
257	35	73	2022	31	20	99	1404776.00	\N	\N	\N
258	35	73	2022	32	0	4	126429.00	\N	\N	\N
259	35	73	2022	32	5	9	210716.00	\N	\N	\N
260	35	73	2022	32	10	19	393337.00	\N	\N	\N
261	35	73	2022	32	20	99	561911.00	\N	\N	\N
262	35	73	2022	40	0	99	905.00	\N	\N	\N
263	35	73	2022	40	100	249	671.00	\N	\N	\N
264	35	73	2022	40	250	499	581.00	\N	\N	\N
265	35	73	2022	40	500	999999	517.00	\N	\N	\N
266	35	73	1999	31	5	9	784061.00	\N	\N	\N
267	35	73	1999	31	10	19	1110753.00	\N	\N	\N
268	35	73	1999	31	20	99	1306769.00	\N	\N	\N
269	35	73	1999	32	0	4	117609.00	\N	\N	\N
270	35	73	1999	32	5	9	196015.00	\N	\N	\N
271	35	73	1999	32	10	19	365895.00	\N	\N	\N
272	35	73	1999	32	20	99	522708.00	\N	\N	\N
273	35	73	1999	40	0	99	842.00	\N	\N	\N
274	35	73	1999	40	100	249	625.00	\N	\N	\N
275	35	73	1999	40	250	499	541.00	\N	\N	\N
276	35	73	1999	40	500	999999	481.00	\N	\N	\N
277	35	73	1998	12	0	50	869.00	\N	\N	\N
278	35	73	1998	12	51	100	788.00	\N	\N	\N
279	35	73	1998	12	101	200	577.00	\N	\N	\N
280	35	73	1998	12	201	400	549.00	\N	\N	\N
281	35	73	1998	12	401	999999	523.00	\N	\N	\N
282	35	73	1998	13	0	50	947.00	\N	\N	\N
283	35	73	1998	13	51	100	866.00	\N	\N	\N
284	35	73	1998	13	101	200	656.00	\N	\N	\N
285	35	73	1998	13	201	400	627.00	\N	\N	\N
286	35	73	2013	12	0	50	934.00	\N	\N	\N
287	35	73	2013	12	51	100	847.00	\N	\N	\N
288	35	73	2013	12	101	200	620.00	\N	\N	\N
289	35	73	2013	12	201	400	590.00	\N	\N	\N
290	35	73	2013	12	401	999999	562.00	\N	\N	\N
291	35	73	2013	13	0	50	1018.00	\N	\N	\N
292	35	73	2013	13	51	100	930.00	\N	\N	\N
293	35	73	2013	13	101	200	705.00	\N	\N	\N
294	35	73	2013	13	201	400	674.00	\N	\N	\N
295	35	73	2013	13	401	999999	646.00	\N	\N	\N
296	35	73	2013	30	0	4	168572.00	\N	\N	\N
297	35	73	2013	30	5	9	351193.00	\N	\N	\N
298	35	73	2013	30	10	19	491671.00	\N	\N	\N
299	35	73	2013	30	20	99	632149.00	\N	\N	\N
300	35	73	2013	31	0	4	561911.00	\N	\N	\N
301	35	73	2013	31	5	9	842865.00	\N	\N	\N
302	35	73	2013	31	10	19	1194059.00	\N	\N	\N
303	35	73	2013	31	20	99	1404776.00	\N	\N	\N
304	35	73	2013	32	0	4	126429.00	\N	\N	\N
305	35	73	2013	32	5	9	210716.00	\N	\N	\N
306	35	73	2013	32	10	19	393337.00	\N	\N	\N
307	35	73	2013	32	20	99	561911.00	\N	\N	\N
308	35	73	2013	40	0	99	905.00	\N	\N	\N
309	35	73	2013	40	100	249	671.00	\N	\N	\N
310	35	73	2013	40	250	499	581.00	\N	\N	\N
311	35	73	2013	40	500	999999	517.00	\N	\N	\N
312	35	73	1998	13	401	999999	601.00	\N	\N	\N
313	35	73	1998	30	0	4	156812.00	\N	\N	\N
314	35	73	1998	30	5	9	326692.00	\N	\N	\N
315	35	73	1998	30	10	19	457369.00	\N	\N	\N
316	35	73	1998	30	20	99	588046.00	\N	\N	\N
317	35	73	1998	31	0	4	522708.00	\N	\N	\N
318	35	73	1998	31	5	9	784061.00	\N	\N	\N
319	35	73	1998	31	10	19	1110753.00	\N	\N	\N
320	35	73	1998	31	20	99	1306769.00	\N	\N	\N
321	35	73	1998	32	0	4	117609.00	\N	\N	\N
322	35	73	1998	32	5	9	196015.00	\N	\N	\N
323	35	73	1998	32	10	19	365895.00	\N	\N	\N
324	35	73	1998	32	20	99	522708.00	\N	\N	\N
325	35	73	1998	40	0	99	842.00	\N	\N	\N
326	35	73	1998	40	100	249	625.00	\N	\N	\N
327	35	73	1998	40	250	499	541.00	\N	\N	\N
328	35	73	1998	40	500	999999	481.00	\N	\N	\N
329	35	73	1997	12	0	50	665.00	\N	\N	\N
330	35	73	1997	12	51	100	603.00	\N	\N	\N
331	35	73	1997	12	101	200	442.00	\N	\N	\N
332	35	73	1997	12	201	400	420.00	\N	\N	\N
333	35	73	1997	12	401	999999	400.00	\N	\N	\N
334	35	73	1997	13	0	50	725.00	\N	\N	\N
335	35	73	1997	13	51	100	663.00	\N	\N	\N
336	35	73	1997	13	101	200	502.00	\N	\N	\N
337	35	73	1997	13	201	400	480.00	\N	\N	\N
338	35	73	1997	13	401	999999	460.00	\N	\N	\N
339	35	73	1997	30	0	4	120000.00	\N	\N	\N
340	35	73	1997	30	5	9	250000.00	\N	\N	\N
341	35	73	1997	30	10	19	350000.00	\N	\N	\N
342	35	73	1997	30	20	99	450000.00	\N	\N	\N
343	35	73	1997	31	0	4	400000.00	\N	\N	\N
344	35	73	1997	31	5	9	600000.00	\N	\N	\N
345	35	73	1997	31	10	19	850000.00	\N	\N	\N
346	35	73	1997	31	20	99	1000000.00	\N	\N	\N
347	35	73	1997	32	0	4	90000.00	\N	\N	\N
348	35	73	1997	32	5	9	150000.00	\N	\N	\N
349	35	73	1997	32	10	19	280000.00	\N	\N	\N
350	35	73	1997	32	20	99	400000.00	\N	\N	\N
351	35	73	1997	40	0	99	700.00	\N	\N	\N
352	35	73	1997	40	100	249	520.00	\N	\N	\N
353	35	73	1997	40	250	499	450.00	\N	\N	\N
354	35	73	1997	40	500	999999	400.00	\N	\N	\N
355	35	73	1996	12	0	50	665.00	\N	\N	\N
356	35	73	1996	12	51	100	603.00	\N	\N	\N
357	35	73	1996	12	101	200	442.00	\N	\N	\N
358	35	73	1996	12	201	400	420.00	\N	\N	\N
359	35	73	1996	12	401	999999	400.00	\N	\N	\N
360	35	73	1996	13	0	50	725.00	\N	\N	\N
361	35	73	1996	13	51	100	663.00	\N	\N	\N
362	35	73	1996	13	101	200	502.00	\N	\N	\N
363	35	73	1996	13	201	400	480.00	\N	\N	\N
364	35	73	1996	13	401	999999	460.00	\N	\N	\N
365	35	73	1996	30	0	4	120000.00	\N	\N	\N
366	35	73	1996	30	5	9	250000.00	\N	\N	\N
367	35	73	1996	30	10	19	350000.00	\N	\N	\N
368	35	73	1996	30	20	99	450000.00	\N	\N	\N
369	35	73	1996	31	0	4	400000.00	\N	\N	\N
370	35	73	1996	31	5	9	600000.00	\N	\N	\N
371	35	73	1996	31	10	19	850000.00	\N	\N	\N
372	35	73	1996	31	20	99	1000000.00	\N	\N	\N
373	35	73	1996	32	0	4	90000.00	\N	\N	\N
374	35	73	1996	32	5	9	150000.00	\N	\N	\N
375	35	73	1996	32	10	19	280000.00	\N	\N	\N
376	35	73	1996	32	20	99	400000.00	\N	\N	\N
377	35	73	1996	40	0	99	700.00	\N	\N	\N
378	35	73	1996	40	100	249	520.00	\N	\N	\N
379	35	73	1996	40	250	499	450.00	\N	\N	\N
380	35	73	1996	40	500	999999	400.00	\N	\N	\N
381	35	73	1995	12	0	50	665.00	\N	\N	\N
382	35	73	1995	12	51	100	603.00	\N	\N	\N
383	35	73	1995	12	101	200	442.00	\N	\N	\N
384	35	73	1995	12	201	400	420.00	\N	\N	\N
385	35	73	1995	12	401	999999	400.00	\N	\N	\N
386	35	73	1995	13	0	50	725.00	\N	\N	\N
387	35	73	1995	13	51	100	663.00	\N	\N	\N
388	35	73	1995	13	101	200	502.00	\N	\N	\N
389	35	73	1995	13	201	400	480.00	\N	\N	\N
390	35	73	1995	13	401	999999	460.00	\N	\N	\N
391	35	73	1995	30	0	4	120000.00	\N	\N	\N
392	35	73	1995	30	5	9	250000.00	\N	\N	\N
393	35	73	1995	30	10	19	350000.00	\N	\N	\N
394	35	73	1995	30	20	99	450000.00	\N	\N	\N
395	35	73	1995	31	0	4	400000.00	\N	\N	\N
396	35	73	1995	31	5	9	600000.00	\N	\N	\N
397	35	73	1995	31	10	19	850000.00	\N	\N	\N
398	35	73	1995	31	20	99	1000000.00	\N	\N	\N
399	35	73	1995	32	0	4	90000.00	\N	\N	\N
400	35	73	1995	32	5	9	150000.00	\N	\N	\N
401	35	73	1995	32	10	19	280000.00	\N	\N	\N
402	35	73	1995	32	20	99	400000.00	\N	\N	\N
403	35	73	1995	40	0	99	700.00	\N	\N	\N
404	35	73	1995	40	100	249	520.00	\N	\N	\N
405	35	73	1995	40	250	499	450.00	\N	\N	\N
406	35	73	1995	40	500	999999	400.00	\N	\N	\N
407	35	73	1994	12	0	50	665.00	\N	\N	\N
408	35	73	1994	12	51	100	603.00	\N	\N	\N
409	35	73	1994	12	101	200	442.00	\N	\N	\N
410	35	73	1994	12	201	400	420.00	\N	\N	\N
411	35	73	1994	12	401	999999	400.00	\N	\N	\N
412	35	73	1994	13	0	50	725.00	\N	\N	\N
413	35	73	1994	13	51	100	663.00	\N	\N	\N
414	35	73	1994	13	101	200	502.00	\N	\N	\N
415	35	73	1994	13	201	400	480.00	\N	\N	\N
416	35	73	1994	13	401	999999	460.00	\N	\N	\N
417	35	73	1994	30	0	4	120000.00	\N	\N	\N
418	35	73	1994	30	5	9	250000.00	\N	\N	\N
419	35	73	1994	30	10	19	350000.00	\N	\N	\N
420	35	73	1994	30	20	99	450000.00	\N	\N	\N
421	35	73	1994	31	0	4	400000.00	\N	\N	\N
422	35	73	1994	31	5	9	600000.00	\N	\N	\N
423	35	73	1994	31	10	19	850000.00	\N	\N	\N
424	35	73	1994	31	20	99	1000000.00	\N	\N	\N
425	35	73	1994	32	0	4	90000.00	\N	\N	\N
426	35	73	1994	32	5	9	150000.00	\N	\N	\N
427	35	73	1994	32	10	19	280000.00	\N	\N	\N
428	35	73	1994	32	20	99	400000.00	\N	\N	\N
429	35	73	1994	40	0	99	700.00	\N	\N	\N
430	35	73	1994	40	100	249	520.00	\N	\N	\N
431	35	73	1994	40	250	499	450.00	\N	\N	\N
432	35	73	1994	40	500	999999	400.00	\N	\N	\N
433	35	73	2001	12	0	50	934.00	\N	\N	\N
434	35	73	2001	12	51	100	847.00	\N	\N	\N
435	35	73	2001	12	101	200	620.00	\N	\N	\N
436	35	73	2001	12	201	400	590.00	\N	\N	\N
437	35	73	2001	12	401	999999	562.00	\N	\N	\N
438	35	73	2001	13	0	50	1018.00	\N	\N	\N
439	35	73	2001	13	51	100	930.00	\N	\N	\N
440	35	73	2001	13	101	200	705.00	\N	\N	\N
441	35	73	2001	13	201	400	674.00	\N	\N	\N
442	35	73	2001	13	401	999999	646.00	\N	\N	\N
443	35	73	2001	30	0	4	168572.00	\N	\N	\N
444	35	73	2001	30	5	9	351193.00	\N	\N	\N
445	35	73	2001	30	10	19	491671.00	\N	\N	\N
446	35	73	2001	30	20	99	632149.00	\N	\N	\N
447	35	73	2001	31	0	4	561911.00	\N	\N	\N
448	35	73	2001	31	5	9	842865.00	\N	\N	\N
449	35	73	2001	31	10	19	1194059.00	\N	\N	\N
450	35	73	2001	31	20	99	1404776.00	\N	\N	\N
451	35	73	2001	32	0	4	126429.00	\N	\N	\N
452	35	73	2001	32	5	9	210716.00	\N	\N	\N
453	35	73	2001	32	10	19	393337.00	\N	\N	\N
454	35	73	2001	32	20	99	561911.00	\N	\N	\N
455	35	73	2001	40	0	99	905.00	\N	\N	\N
456	35	73	2001	40	100	249	671.00	\N	\N	\N
457	35	73	2001	40	250	499	581.00	\N	\N	\N
458	35	73	2001	40	500	999999	517.00	\N	\N	\N
459	35	73	2002	12	0	50	934.00	\N	\N	\N
460	35	73	2002	12	51	100	847.00	\N	\N	\N
461	35	73	2002	12	101	200	620.00	\N	\N	\N
462	35	73	2002	12	201	400	590.00	\N	\N	\N
463	35	73	2002	12	401	999999	562.00	\N	\N	\N
464	35	73	2002	13	0	50	1018.00	\N	\N	\N
465	35	73	2002	13	51	100	930.00	\N	\N	\N
466	35	73	2002	13	101	200	705.00	\N	\N	\N
467	35	73	2002	13	201	400	674.00	\N	\N	\N
468	35	73	2002	13	401	999999	646.00	\N	\N	\N
469	35	73	2002	30	0	4	168572.00	\N	\N	\N
470	35	73	2002	30	5	9	351193.00	\N	\N	\N
471	35	73	2002	30	10	19	491671.00	\N	\N	\N
472	35	73	2002	30	20	99	632149.00	\N	\N	\N
473	35	73	2002	31	0	4	561911.00	\N	\N	\N
474	35	73	2002	31	5	9	842865.00	\N	\N	\N
475	35	73	2002	31	10	19	1194059.00	\N	\N	\N
476	35	73	2002	31	20	99	1404776.00	\N	\N	\N
477	35	73	2002	32	0	4	126429.00	\N	\N	\N
478	35	73	2002	32	5	9	210716.00	\N	\N	\N
479	35	73	2002	32	10	19	393337.00	\N	\N	\N
480	35	73	2002	32	20	99	561911.00	\N	\N	\N
481	35	73	2002	40	0	99	905.00	\N	\N	\N
482	35	73	2002	40	100	249	671.00	\N	\N	\N
483	35	73	2002	40	250	499	581.00	\N	\N	\N
484	35	73	2002	40	500	999999	517.00	\N	\N	\N
485	35	73	2012	12	0	50	934.00	\N	\N	\N
486	35	73	2012	12	51	100	847.00	\N	\N	\N
487	35	73	2012	12	101	200	620.00	\N	\N	\N
488	35	73	2012	12	201	400	590.00	\N	\N	\N
489	35	73	2012	12	401	999999	562.00	\N	\N	\N
490	35	73	2012	13	0	50	1018.00	\N	\N	\N
491	35	73	2012	13	51	100	930.00	\N	\N	\N
492	35	73	2012	13	101	200	705.00	\N	\N	\N
493	35	73	2012	13	201	400	674.00	\N	\N	\N
494	35	73	2012	13	401	999999	646.00	\N	\N	\N
495	35	73	2012	30	0	4	168572.00	\N	\N	\N
496	35	73	2012	30	5	9	351193.00	\N	\N	\N
497	35	73	2012	30	10	19	491671.00	\N	\N	\N
498	35	73	2012	30	20	99	632149.00	\N	\N	\N
499	35	73	2012	31	0	4	561911.00	\N	\N	\N
500	35	73	2012	31	5	9	842865.00	\N	\N	\N
501	35	73	2012	31	10	19	1194059.00	\N	\N	\N
502	35	73	2012	31	20	99	1404776.00	\N	\N	\N
503	35	73	2012	32	0	4	126429.00	\N	\N	\N
504	35	73	2012	32	5	9	210716.00	\N	\N	\N
505	35	73	2012	32	10	19	393337.00	\N	\N	\N
506	35	73	2012	32	20	99	561911.00	\N	\N	\N
507	35	73	2012	40	0	99	905.00	\N	\N	\N
508	35	73	2012	40	100	249	671.00	\N	\N	\N
509	35	73	2012	40	250	499	581.00	\N	\N	\N
510	35	73	2012	40	500	999999	517.00	\N	\N	\N
511	35	73	2009	12	0	50	934.00	\N	\N	\N
512	35	73	2009	12	51	100	847.00	\N	\N	\N
513	35	73	2009	12	101	200	620.00	\N	\N	\N
514	35	73	2009	12	201	400	590.00	\N	\N	\N
515	35	73	2009	12	401	999999	562.00	\N	\N	\N
516	35	73	2009	13	0	50	1018.00	\N	\N	\N
517	35	73	2009	13	51	100	930.00	\N	\N	\N
518	35	73	2009	13	101	200	705.00	\N	\N	\N
519	35	73	2009	13	201	400	674.00	\N	\N	\N
520	35	73	2009	13	401	999999	646.00	\N	\N	\N
521	35	73	2009	30	0	4	168572.00	\N	\N	\N
522	35	73	2009	30	5	9	351193.00	\N	\N	\N
523	35	73	2009	30	10	19	491671.00	\N	\N	\N
524	35	73	2009	30	20	99	632149.00	\N	\N	\N
525	35	73	2009	31	0	4	561911.00	\N	\N	\N
526	35	73	2009	31	5	9	842865.00	\N	\N	\N
527	35	73	2009	31	10	19	1194059.00	\N	\N	\N
528	35	73	2009	31	20	99	1404776.00	\N	\N	\N
529	35	73	2009	32	0	4	126429.00	\N	\N	\N
530	35	73	2009	32	5	9	210716.00	\N	\N	\N
531	35	73	2009	32	10	19	393337.00	\N	\N	\N
532	35	73	2009	32	20	99	561911.00	\N	\N	\N
533	35	73	2009	40	0	99	905.00	\N	\N	\N
534	35	73	2009	40	100	249	671.00	\N	\N	\N
535	35	73	2009	40	250	499	581.00	\N	\N	\N
536	35	73	2009	40	500	999999	517.00	\N	\N	\N
537	35	73	2010	12	0	50	934.00	\N	\N	\N
538	35	73	2010	12	51	100	847.00	\N	\N	\N
539	35	73	2010	12	101	200	620.00	\N	\N	\N
540	35	73	2010	12	201	400	590.00	\N	\N	\N
541	35	73	2010	12	401	999999	562.00	\N	\N	\N
542	35	73	2010	13	0	50	1018.00	\N	\N	\N
543	35	73	2010	13	51	100	930.00	\N	\N	\N
544	35	73	2010	13	101	200	705.00	\N	\N	\N
545	35	73	2010	13	201	400	674.00	\N	\N	\N
546	35	73	2010	13	401	999999	646.00	\N	\N	\N
547	35	73	2010	30	0	4	168572.00	\N	\N	\N
548	35	73	2010	30	5	9	351193.00	\N	\N	\N
549	35	73	2010	30	10	19	491671.00	\N	\N	\N
550	35	73	2010	30	20	99	632149.00	\N	\N	\N
551	35	73	2010	31	0	4	561911.00	\N	\N	\N
552	35	73	2010	31	5	9	842865.00	\N	\N	\N
553	35	73	2010	31	10	19	1194059.00	\N	\N	\N
554	35	73	2010	31	20	99	1404776.00	\N	\N	\N
555	35	73	2010	32	0	4	126429.00	\N	\N	\N
556	35	73	2010	32	5	9	210716.00	\N	\N	\N
557	35	73	2010	32	10	19	393337.00	\N	\N	\N
558	35	73	2010	32	20	99	561911.00	\N	\N	\N
559	35	73	2010	40	0	99	905.00	\N	\N	\N
560	35	73	2010	40	100	249	671.00	\N	\N	\N
561	35	73	2010	40	250	499	581.00	\N	\N	\N
562	35	73	2010	40	500	999999	517.00	\N	\N	\N
563	35	73	2011	12	0	50	934.00	\N	\N	\N
564	35	73	2011	12	51	100	847.00	\N	\N	\N
565	35	73	2011	12	101	200	620.00	\N	\N	\N
566	35	73	2011	12	201	400	590.00	\N	\N	\N
567	35	73	2011	12	401	999999	562.00	\N	\N	\N
568	35	73	2011	13	0	50	1018.00	\N	\N	\N
569	35	73	2011	13	51	100	930.00	\N	\N	\N
570	35	73	2011	13	101	200	705.00	\N	\N	\N
571	35	73	2011	13	201	400	674.00	\N	\N	\N
572	35	73	2011	13	401	999999	646.00	\N	\N	\N
573	35	73	2011	30	0	4	168572.00	\N	\N	\N
574	35	73	2011	30	5	9	351193.00	\N	\N	\N
575	35	73	2011	30	10	19	491671.00	\N	\N	\N
576	35	73	2011	30	20	99	632149.00	\N	\N	\N
577	35	73	2011	31	0	4	561911.00	\N	\N	\N
578	35	73	2011	31	5	9	842865.00	\N	\N	\N
579	35	73	2011	31	10	19	1194059.00	\N	\N	\N
580	35	73	2011	31	20	99	1404776.00	\N	\N	\N
581	35	73	2011	32	0	4	126429.00	\N	\N	\N
582	35	73	2011	32	5	9	210716.00	\N	\N	\N
583	35	73	2011	32	10	19	393337.00	\N	\N	\N
584	35	73	2011	32	20	99	561911.00	\N	\N	\N
585	35	73	2011	40	0	99	905.00	\N	\N	\N
586	35	73	2011	40	100	249	671.00	\N	\N	\N
587	35	73	2011	40	250	499	581.00	\N	\N	\N
588	35	73	2011	40	500	999999	517.00	\N	\N	\N
589	35	73	2003	12	0	50	934.00	\N	\N	\N
590	35	73	2003	12	51	100	847.00	\N	\N	\N
591	35	73	2003	12	101	200	620.00	\N	\N	\N
592	35	73	2003	12	201	400	590.00	\N	\N	\N
593	35	73	2003	12	401	999999	562.00	\N	\N	\N
594	35	73	2003	13	0	50	1018.00	\N	\N	\N
595	35	73	2003	13	51	100	930.00	\N	\N	\N
596	35	73	2003	13	101	200	705.00	\N	\N	\N
597	35	73	2003	13	201	400	674.00	\N	\N	\N
598	35	73	2003	13	401	999999	646.00	\N	\N	\N
599	35	73	2003	30	0	4	168572.00	\N	\N	\N
600	35	73	2003	30	5	9	351193.00	\N	\N	\N
601	35	73	2003	30	10	19	491671.00	\N	\N	\N
602	35	73	2003	30	20	99	632149.00	\N	\N	\N
603	35	73	2003	31	0	4	561911.00	\N	\N	\N
604	35	73	2003	31	5	9	842865.00	\N	\N	\N
605	35	73	2003	31	10	19	1194059.00	\N	\N	\N
606	35	73	2003	31	20	99	1404776.00	\N	\N	\N
607	35	73	2003	32	0	4	126429.00	\N	\N	\N
608	35	73	2003	32	5	9	210716.00	\N	\N	\N
609	35	73	2003	32	10	19	393337.00	\N	\N	\N
610	35	73	2003	32	20	99	561911.00	\N	\N	\N
611	35	73	2003	40	0	99	905.00	\N	\N	\N
612	35	73	2003	40	100	249	671.00	\N	\N	\N
613	35	73	2003	40	250	499	581.00	\N	\N	\N
614	35	73	2003	40	500	999999	517.00	\N	\N	\N
615	35	73	2004	12	0	50	934.00	\N	\N	\N
616	35	73	2004	12	51	100	847.00	\N	\N	\N
617	35	73	2004	12	101	200	620.00	\N	\N	\N
618	35	73	2004	12	201	400	590.00	\N	\N	\N
619	35	73	2004	12	401	999999	562.00	\N	\N	\N
620	35	73	2004	13	0	50	1018.00	\N	\N	\N
621	35	73	2004	13	51	100	930.00	\N	\N	\N
622	35	73	2004	13	101	200	705.00	\N	\N	\N
623	35	73	2004	13	201	400	674.00	\N	\N	\N
624	35	73	2004	13	401	999999	646.00	\N	\N	\N
625	35	73	2004	30	0	4	168572.00	\N	\N	\N
626	35	73	2004	30	5	9	351193.00	\N	\N	\N
627	35	73	2004	30	10	19	491671.00	\N	\N	\N
628	35	73	2004	30	20	99	632149.00	\N	\N	\N
629	35	73	2004	31	0	4	561911.00	\N	\N	\N
630	35	73	2004	31	5	9	842865.00	\N	\N	\N
631	35	73	2004	31	10	19	1194059.00	\N	\N	\N
632	35	73	2004	31	20	99	1404776.00	\N	\N	\N
633	35	73	2004	32	0	4	126429.00	\N	\N	\N
634	35	73	2004	32	5	9	210716.00	\N	\N	\N
635	35	73	2004	32	10	19	393337.00	\N	\N	\N
636	35	73	2004	32	20	99	561911.00	\N	\N	\N
637	35	73	2004	40	0	99	905.00	\N	\N	\N
638	35	73	2004	40	100	249	671.00	\N	\N	\N
639	35	73	2004	40	250	499	581.00	\N	\N	\N
640	35	73	2004	40	500	999999	517.00	\N	\N	\N
641	35	73	2005	12	0	50	934.00	\N	\N	\N
642	35	73	2005	12	51	100	847.00	\N	\N	\N
643	35	73	2005	12	101	200	620.00	\N	\N	\N
644	35	73	2005	12	201	400	590.00	\N	\N	\N
645	35	73	2005	12	401	999999	562.00	\N	\N	\N
646	35	73	2005	13	0	50	1018.00	\N	\N	\N
647	35	73	2020	12	0	50	934.00	\N	\N	\N
648	35	73	2020	12	51	100	847.00	\N	\N	\N
649	35	73	2020	12	101	200	620.00	\N	\N	\N
650	35	73	2020	12	201	400	590.00	\N	\N	\N
651	35	73	2020	12	401	999999	562.00	\N	\N	\N
652	35	73	2020	13	0	50	1018.00	\N	\N	\N
653	35	73	2020	13	51	100	930.00	\N	\N	\N
654	35	73	2020	13	101	200	705.00	\N	\N	\N
655	35	73	2020	13	201	400	674.00	\N	\N	\N
656	35	73	2020	13	401	999999	646.00	\N	\N	\N
657	35	73	2020	30	0	4	168572.00	\N	\N	\N
658	35	73	2020	30	5	9	351193.00	\N	\N	\N
659	35	73	2020	30	10	19	491671.00	\N	\N	\N
660	35	73	2020	30	20	99	632149.00	\N	\N	\N
661	35	73	2020	31	0	4	561911.00	\N	\N	\N
662	35	73	2020	31	5	9	842865.00	\N	\N	\N
663	35	73	2020	31	10	19	1194059.00	\N	\N	\N
664	35	73	2020	31	20	99	1404776.00	\N	\N	\N
665	35	73	2020	32	0	4	126429.00	\N	\N	\N
666	35	73	2020	32	5	9	210716.00	\N	\N	\N
667	35	73	2020	32	10	19	393337.00	\N	\N	\N
668	35	73	2020	32	20	99	561911.00	\N	\N	\N
669	35	73	2020	40	0	99	905.00	\N	\N	\N
670	35	73	2020	40	100	249	671.00	\N	\N	\N
671	35	73	2020	40	250	499	581.00	\N	\N	\N
672	35	73	2020	40	500	999999	517.00	\N	\N	\N
673	35	73	2018	12	0	50	934.00	\N	\N	\N
674	35	73	2018	12	51	100	847.00	\N	\N	\N
675	35	73	2018	12	101	200	620.00	\N	\N	\N
676	35	73	2018	12	201	400	590.00	\N	\N	\N
677	35	73	2018	12	401	999999	562.00	\N	\N	\N
678	35	73	2018	13	0	50	1018.00	\N	\N	\N
679	35	73	2018	13	51	100	930.00	\N	\N	\N
680	35	73	2018	13	101	200	705.00	\N	\N	\N
681	35	73	2018	13	201	400	674.00	\N	\N	\N
682	35	73	2018	13	401	999999	646.00	\N	\N	\N
683	35	73	2018	30	0	4	168572.00	\N	\N	\N
684	35	73	2018	30	5	9	351193.00	\N	\N	\N
685	35	73	2018	30	10	19	491671.00	\N	\N	\N
686	35	73	2018	30	20	99	632149.00	\N	\N	\N
687	35	73	2018	31	0	4	561911.00	\N	\N	\N
688	35	73	2018	31	5	9	842865.00	\N	\N	\N
689	35	73	2018	31	10	19	1194059.00	\N	\N	\N
690	35	73	2018	31	20	99	1404776.00	\N	\N	\N
691	35	73	2018	32	0	4	126429.00	\N	\N	\N
692	35	73	2018	32	5	9	210716.00	\N	\N	\N
693	35	73	2018	32	10	19	393337.00	\N	\N	\N
694	35	73	2018	32	20	99	561911.00	\N	\N	\N
695	35	73	2018	40	0	99	905.00	\N	\N	\N
696	35	73	2018	40	100	249	671.00	\N	\N	\N
697	35	73	2018	40	250	499	581.00	\N	\N	\N
698	35	73	2018	40	500	999999	517.00	\N	\N	\N
699	35	73	2021	12	0	50	934.00	\N	\N	\N
700	35	73	2021	12	51	100	847.00	\N	\N	\N
701	35	73	2021	12	101	200	620.00	\N	\N	\N
702	35	73	2021	12	201	400	590.00	\N	\N	\N
703	35	73	2021	12	401	999999	562.00	\N	\N	\N
704	35	73	2021	13	0	50	1018.00	\N	\N	\N
705	35	73	2021	13	51	100	930.00	\N	\N	\N
706	35	73	2021	13	101	200	705.00	\N	\N	\N
707	35	73	2021	13	201	400	674.00	\N	\N	\N
708	35	73	2021	13	401	999999	646.00	\N	\N	\N
709	35	73	2021	30	0	4	168572.00	\N	\N	\N
710	35	73	2021	30	5	9	351193.00	\N	\N	\N
711	35	73	2021	30	10	19	491671.00	\N	\N	\N
712	35	73	2021	30	20	99	632149.00	\N	\N	\N
713	35	73	2021	31	0	4	561911.00	\N	\N	\N
714	35	73	2021	31	5	9	842865.00	\N	\N	\N
715	35	73	2021	31	10	19	1194059.00	\N	\N	\N
716	35	73	2021	31	20	99	1404776.00	\N	\N	\N
717	35	73	2021	32	0	4	126429.00	\N	\N	\N
718	35	73	2021	32	5	9	210716.00	\N	\N	\N
719	35	73	2021	32	10	19	393337.00	\N	\N	\N
720	35	73	2021	32	20	99	561911.00	\N	\N	\N
721	35	73	2021	40	0	99	905.00	\N	\N	\N
722	35	73	2021	40	100	249	671.00	\N	\N	\N
723	35	73	2021	40	250	499	581.00	\N	\N	\N
724	35	73	2021	40	500	999999	517.00	\N	\N	\N
725	35	73	2019	12	0	50	934.00	\N	\N	\N
726	35	73	2019	12	51	100	847.00	\N	\N	\N
727	35	73	2019	12	101	200	620.00	\N	\N	\N
728	35	73	2019	12	201	400	590.00	\N	\N	\N
729	35	73	2019	12	401	999999	562.00	\N	\N	\N
730	35	73	2019	13	0	50	1018.00	\N	\N	\N
731	35	73	2019	13	51	100	930.00	\N	\N	\N
732	35	73	2019	13	101	200	705.00	\N	\N	\N
733	35	73	2019	13	201	400	674.00	\N	\N	\N
734	35	73	2019	13	401	999999	646.00	\N	\N	\N
735	35	73	2019	30	0	4	168572.00	\N	\N	\N
736	35	73	2019	30	5	9	351193.00	\N	\N	\N
737	35	73	2019	30	10	19	491671.00	\N	\N	\N
738	35	73	2019	30	20	99	632149.00	\N	\N	\N
739	35	73	2019	31	0	4	561911.00	\N	\N	\N
740	35	73	2019	31	5	9	842865.00	\N	\N	\N
741	35	73	2019	31	10	19	1194059.00	\N	\N	\N
742	35	73	2019	31	20	99	1404776.00	\N	\N	\N
743	35	73	2019	32	0	4	126429.00	\N	\N	\N
744	35	73	2019	32	5	9	210716.00	\N	\N	\N
745	35	73	2019	32	10	19	393337.00	\N	\N	\N
746	35	73	2019	32	20	99	561911.00	\N	\N	\N
747	35	73	2019	40	0	99	905.00	\N	\N	\N
748	35	73	2019	40	100	249	671.00	\N	\N	\N
749	35	73	2019	40	250	499	581.00	\N	\N	\N
750	35	73	2019	40	500	999999	517.00	\N	\N	\N
751	35	73	2005	13	51	100	930.00	\N	\N	\N
752	35	73	2005	13	101	200	705.00	\N	\N	\N
753	35	73	2005	13	201	400	674.00	\N	\N	\N
754	35	73	2005	13	401	999999	646.00	\N	\N	\N
\.


--
-- Name: DbkbFasumDepMinMax_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."DbkbFasumDepMinMax_Id_seq"', 754, true);


--
-- PostgreSQL database dump complete
--

