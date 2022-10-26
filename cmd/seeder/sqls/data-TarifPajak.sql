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
-- Data for Name: TarifPajak; Type: TABLE DATA; Schema: public; Owner: dexwip
--

COPY public."TarifPajak" ("Id", "Rekening_Id", "Tahun", "TarifRp", "TarifPersen", "OmsetAwal", "OmsetAkhir") FROM stdin;
88	10189	2020	1.00	25.00	0.00	0.00
2	10101	2016	1.00	10.00	0.00	0.00
3	10102	2016	1.00	10.00	0.00	0.00
4	10103	2016	1.00	10.00	0.00	0.00
5	10104	2016	1.00	10.00	0.00	0.00
6	10105	2016	1.00	10.00	0.00	0.00
7	10106	2016	1.00	10.00	0.00	0.00
8	10107	2016	1.00	10.00	0.00	0.00
9	10108	2016	1.00	10.00	0.00	0.00
10	10109	2016	1.00	10.00	0.00	0.00
11	10111	2016	1.00	10.00	0.00	0.00
12	10112	2016	1.00	10.00	0.00	0.00
13	10113	2016	1.00	10.00	0.00	0.00
14	10114	2016	1.00	10.00	0.00	0.00
15	10115	2016	1.00	10.00	0.00	0.00
16	10116	2016	1.00	10.00	0.00	0.00
17	10118	2016	1.00	10.00	0.00	0.00
18	10119	2017	1.00	15.00	0.00	0.00
19	10120	2017	1.00	15.00	0.00	0.00
20	10121	2017	1.00	15.00	0.00	0.00
21	10122	2017	1.00	15.00	0.00	0.00
22	10123	2017	1.00	50.00	0.00	0.00
23	10125	2016	100000.00	10.00	0.00	0.00
24	10126	2017	1.00	2.50	0.00	0.00
25	10127	2016	3000.00	0.00	0.00	0.00
26	10132	2016	1.00	10.00	0.00	0.00
27	10130	2016	93750.00	0.00	0.00	0.00
28	10129	2016	0.00	0.00	0.00	0.00
29	10128	2016	3250.00	0.00	0.00	0.00
89	10191	2017	1.00	20.00	0.00	0.00
90	10189	2020	1.00	25.00	0.00	0.00
32	10137	2016	1.00	10.00	0.00	0.00
33	10138	2016	1.00	10.00	0.00	0.00
34	10139	2017	1.00	10.00	0.00	0.00
35	10140	2016	1.00	10.00	0.00	0.00
36	10141	2016	1.00	10.00	0.00	0.00
41	10145	2017	1.00	25.00	0.00	0.00
39	10143	2016	1.00	10.00	0.00	0.00
40	10144	2016	1.00	10.00	0.00	0.00
42	10146	2017	1.00	50.00	0.00	0.00
43	10147	2017	1.00	15.00	0.00	0.00
44	10148	2017	1.00	15.00	0.00	0.00
45	10149	2016	1.00	0.00	0.00	0.00
46	10150	2017	1.00	15.00	0.00	0.00
47	10151	2017	1.00	15.00	0.00	0.00
48	10152	2017	1.00	15.00	0.00	0.00
49	10153	2017	1.00	15.00	0.00	0.00
50	10154	2017	1.00	25.00	0.00	0.00
51	10155	2017	1.00	25.00	0.00	0.00
52	10156	2017	1.00	25.00	0.00	0.00
53	10157	2017	1.00	15.00	0.00	0.00
54	10158	2016	50000.00	0.00	0.00	0.00
55	10159	2016	20000.00	0.00	0.00	0.00
56	10160	2016	132500.00	0.00	0.00	0.00
57	10161	2016	25000.00	0.00	0.00	0.00
91	10101	2017	1.00	10.00	0.00	0.00
92	10102	2017	1.00	10.00	0.00	0.00
93	10103	2017	1.00	10.00	0.00	0.00
94	10104	2017	1.00	10.00	0.00	0.00
95	10105	2017	1.00	10.00	0.00	0.00
96	10106	2017	1.00	10.00	0.00	0.00
97	10107	2017	1.00	10.00	0.00	0.00
98	10108	2017	1.00	10.00	0.00	0.00
99	10109	2017	1.00	10.00	0.00	0.00
100	10111	2017	1.00	10.00	0.00	0.00
101	10112	2017	1.00	10.00	0.00	0.00
102	10113	2017	1.00	10.00	0.00	0.00
103	10114	2017	1.00	10.00	0.00	0.00
104	10115	2017	1.00	10.00	0.00	0.00
105	10116	2017	1.00	10.00	0.00	0.00
106	10118	2017	1.00	10.00	0.00	0.00
107	10119	2017	1.00	15.00	0.00	0.00
108	10120	2017	1.00	15.00	0.00	0.00
109	10121	2017	1.00	15.00	0.00	0.00
110	10122	2017	1.00	15.00	0.00	0.00
111	10123	2017	1.00	50.00	0.00	0.00
112	10125	2017	100000.00	10.00	0.00	0.00
113	10126	2017	1.00	2.50	0.00	0.00
114	10127	2017	3000.00	0.00	0.00	0.00
115	10132	2017	1.00	10.00	0.00	0.00
116	10130	2017	93750.00	0.00	0.00	0.00
117	10129	2017	0.00	0.00	0.00	0.00
118	10128	2017	3250.00	0.00	0.00	0.00
119	10191	2017	1.00	20.00	0.00	0.00
120	10137	2017	1.00	10.00	0.00	0.00
121	10138	2017	1.00	10.00	0.00	0.00
122	10139	2017	1.00	10.00	0.00	0.00
123	10140	2017	1.00	10.00	0.00	0.00
124	10141	2017	1.00	10.00	0.00	0.00
125	10145	2017	1.00	25.00	0.00	0.00
126	10143	2017	1.00	10.00	0.00	0.00
127	10144	2017	1.00	10.00	0.00	0.00
128	10146	2017	1.00	50.00	0.01	0.00
129	10147	2017	1.00	15.00	0.00	0.00
130	10148	2017	1.00	15.00	0.00	0.00
131	10149	2017	1.00	0.00	0.00	0.00
132	10150	2017	1.00	15.00	0.00	0.00
133	10151	2017	1.00	15.00	0.00	0.00
134	10152	2017	1.00	15.00	0.00	0.00
135	10153	2017	1.00	15.00	0.00	0.00
136	10154	2017	1.00	25.00	0.00	0.00
137	10155	2017	1.00	25.00	0.00	0.00
138	10156	2017	1.00	25.00	0.00	0.00
139	10157	2017	1.00	15.00	0.00	0.00
140	10158	2017	50000.00	0.00	0.00	0.00
141	10159	2017	20000.00	0.00	0.00	0.00
142	10160	2017	132500.00	0.00	0.00	0.00
143	10161	2017	25000.00	0.00	0.00	0.00
161	10205	2017	1.00	1.50	0.00	0.00
162	10206	2017	1.00	15.00	0.00	0.00
170	10213	2017	1.00	10.00	0.00	0.00
171	10119	2018	1.00	15.00	0.00	0.00
172	10120	2018	1.00	15.00	0.00	0.00
173	10121	2018	1.00	15.00	0.00	0.00
174	10122	2018	1.00	15.00	0.00	0.00
175	10123	2018	1.00	50.00	0.00	0.00
176	10126	2018	1.00	2.50	0.00	0.00
177	10191	2018	1.00	20.00	0.00	0.00
178	10189	2020	1.00	25.00	0.00	0.00
179	10139	2018	1.00	10.00	0.00	0.00
180	10145	2018	1.00	25.00	0.00	0.00
181	10146	2018	1.00	50.00	0.00	0.00
182	10147	2018	1.00	15.00	0.00	0.00
163	10208	2017	1.00	5.00	0.00	0.00
164	10207	2017	1.00	10.00	0.00	0.00
183	10148	2018	1.00	15.00	0.00	0.00
184	10150	2018	1.00	15.00	0.00	0.00
185	10151	2018	1.00	15.00	0.00	0.00
186	10152	2018	1.00	15.00	0.00	0.00
187	10153	2018	1.00	15.00	0.00	0.00
188	10154	2018	1.00	25.00	0.00	0.00
166	10209	2017	1.00	10.00	0.00	0.00
167	10210	2017	1.00	25.00	0.00	0.00
189	10155	2018	1.00	25.00	0.00	0.00
168	10211	2022	1.00	35.00	0.00	0.00
190	10156	2018	1.00	25.00	0.00	0.00
191	10157	2018	1.00	15.00	0.00	0.00
169	10212	2017	1.00	35.00	0.00	0.00
192	10101	2018	1.00	10.00	0.00	0.00
193	10102	2018	1.00	10.00	0.00	0.00
194	10103	2018	1.00	10.00	0.00	0.00
195	10104	2018	1.00	10.00	0.00	0.00
196	10105	2018	1.00	10.00	0.00	0.00
197	10106	2018	1.00	10.00	0.00	0.00
198	10107	2018	1.00	10.00	0.00	0.00
199	10108	2018	1.00	10.00	0.00	0.00
200	10109	2018	1.00	10.00	0.00	0.00
201	10111	2018	1.00	10.00	0.00	0.00
202	10112	2018	1.00	10.00	0.00	0.00
203	10113	2018	1.00	10.00	0.00	0.00
204	10114	2018	1.00	10.00	0.00	0.00
205	10115	2018	1.00	10.00	0.00	0.00
206	10116	2018	1.00	10.00	0.00	0.00
207	10118	2018	1.00	10.00	0.00	0.00
208	10119	2018	1.00	15.00	0.00	0.00
209	10120	2018	1.00	15.00	0.00	0.00
210	10121	2018	1.00	15.00	0.00	0.00
211	10122	2018	1.00	15.00	0.00	0.00
212	10123	2018	1.00	50.00	0.00	0.00
213	10125	2018	100000.00	10.00	0.00	0.00
214	10126	2018	1.00	2.50	0.00	0.00
215	10127	2018	3000.00	0.00	0.00	0.00
216	10132	2018	1.00	10.00	0.00	0.00
217	10130	2018	93750.00	0.00	0.00	0.00
218	10129	2018	0.00	0.00	0.00	0.00
219	10128	2018	3250.00	0.00	0.00	0.00
220	10191	2018	1.00	20.00	0.00	0.00
221	10137	2018	1.00	10.00	0.00	0.00
222	10138	2018	1.00	10.00	0.00	0.00
223	10139	2018	1.00	10.00	0.00	0.00
224	10140	2018	1.00	10.00	0.00	0.00
225	10141	2018	1.00	10.00	0.00	0.00
226	10145	2018	1.00	25.00	0.00	0.00
227	10143	2018	1.00	10.00	0.00	0.00
228	10144	2018	1.00	10.00	0.00	0.00
229	10146	2018	1.00	50.00	0.01	0.00
230	10147	2018	1.00	15.00	0.00	0.00
231	10148	2018	1.00	15.00	0.00	0.00
232	10149	2018	1.00	0.00	0.00	0.00
233	10150	2018	1.00	15.00	0.00	0.00
234	10151	2018	1.00	15.00	0.00	0.00
235	10152	2018	1.00	15.00	0.00	0.00
236	10153	2018	1.00	15.00	0.00	0.00
237	10154	2018	1.00	25.00	0.00	0.00
238	10155	2018	1.00	25.00	0.00	0.00
239	10156	2018	1.00	25.00	0.00	0.00
240	10157	2018	1.00	15.00	0.00	0.00
241	10158	2018	50000.00	0.00	0.00	0.00
242	10159	2018	20000.00	0.00	0.00	0.00
243	10160	2018	132500.00	0.00	0.00	0.00
244	10161	2018	25000.00	0.00	0.00	0.00
262	10205	2018	1.00	1.50	0.00	0.00
263	10206	2018	1.00	15.00	0.00	0.00
264	10213	2018	1.00	10.00	0.00	0.00
265	10208	2018	1.00	5.00	0.00	0.00
266	10207	2018	1.00	10.00	0.00	0.00
267	10209	2018	1.00	10.00	0.00	0.00
268	10210	2018	1.00	25.00	0.00	0.00
269	10211	2018	1.00	35.00	0.00	0.00
270	10212	2018	1.00	35.00	0.00	0.00
271	10214	2018	27000.00	0.00	0.00	0.00
272	10119	2019	1.00	15.00	0.00	0.00
273	10120	2019	1.00	15.00	0.00	0.00
274	10121	2019	1.00	15.00	0.00	0.00
275	10122	2019	1.00	15.00	0.00	0.00
276	10123	2019	1.00	50.00	0.00	0.00
277	10126	2019	1.00	2.50	0.00	0.00
278	10191	2019	1.00	20.00	0.00	0.00
279	10189	2020	1.00	25.00	0.00	0.00
280	10139	2019	1.00	10.00	0.00	0.00
281	10145	2019	1.00	25.00	0.00	0.00
282	10146	2019	1.00	50.00	0.00	0.00
283	10147	2019	1.00	15.00	0.00	0.00
284	10148	2019	1.00	15.00	0.00	0.00
285	10150	2019	1.00	15.00	0.00	0.00
286	10151	2019	1.00	15.00	0.00	0.00
287	10152	2019	1.00	15.00	0.00	0.00
288	10153	2019	1.00	15.00	0.00	0.00
289	10154	2019	1.00	25.00	0.00	0.00
290	10155	2019	1.00	25.00	0.00	0.00
291	10156	2019	1.00	25.00	0.00	0.00
292	10157	2019	1.00	15.00	0.00	0.00
293	10101	2019	1.00	10.00	0.00	0.00
294	10102	2019	1.00	10.00	0.00	0.00
295	10103	2019	1.00	10.00	0.00	0.00
296	10104	2019	1.00	10.00	0.00	0.00
297	10105	2019	1.00	10.00	0.00	0.00
298	10106	2019	1.00	10.00	0.00	0.00
299	10107	2019	1.00	10.00	0.00	0.00
300	10108	2019	1.00	10.00	0.00	0.00
301	10109	2019	1.00	10.00	0.00	0.00
302	10111	2019	1.00	10.00	0.00	0.00
303	10112	2019	1.00	10.00	0.00	0.00
304	10113	2019	1.00	10.00	0.00	0.00
305	10114	2019	1.00	10.00	0.00	0.00
306	10115	2019	1.00	10.00	0.00	0.00
307	10116	2019	1.00	10.00	0.00	0.00
308	10118	2019	1.00	10.00	0.00	0.00
309	10119	2019	1.00	15.00	0.00	0.00
310	10120	2019	1.00	15.00	0.00	0.00
311	10121	2019	1.00	15.00	0.00	0.00
312	10122	2019	1.00	15.00	0.00	0.00
313	10123	2019	1.00	50.00	0.00	0.00
314	10125	2019	100000.00	10.00	0.00	0.00
315	10126	2019	1.00	2.50	0.00	0.00
316	10127	2019	3000.00	0.00	0.00	0.00
317	10132	2019	1.00	10.00	0.00	0.00
318	10130	2019	93750.00	0.00	0.00	0.00
319	10129	2019	0.00	0.00	0.00	0.00
320	10128	2019	3250.00	0.00	0.00	0.00
321	10191	2019	1.00	20.00	0.00	0.00
322	10137	2019	1.00	10.00	0.00	0.00
323	10138	2019	1.00	10.00	0.00	0.00
324	10139	2019	1.00	10.00	0.00	0.00
325	10140	2019	1.00	10.00	0.00	0.00
326	10141	2019	1.00	10.00	0.00	0.00
327	10145	2019	1.00	25.00	0.00	0.00
328	10143	2019	1.00	10.00	0.00	0.00
329	10144	2019	1.00	10.00	0.00	0.00
330	10146	2019	1.00	50.00	0.01	0.00
331	10147	2019	1.00	15.00	0.00	0.00
332	10148	2019	1.00	15.00	0.00	0.00
333	10149	2019	1.00	0.00	0.00	0.00
334	10150	2019	1.00	15.00	0.00	0.00
335	10151	2019	1.00	15.00	0.00	0.00
336	10152	2019	1.00	15.00	0.00	0.00
337	10153	2019	1.00	15.00	0.00	0.00
338	10154	2019	1.00	25.00	0.00	0.00
339	10155	2019	1.00	25.00	0.00	0.00
340	10156	2019	1.00	25.00	0.00	0.00
341	10157	2019	1.00	15.00	0.00	0.00
342	10158	2019	50000.00	0.00	0.00	0.00
343	10159	2019	20000.00	0.00	0.00	0.00
344	10160	2019	132500.00	0.00	0.00	0.00
345	10161	2019	25000.00	0.00	0.00	0.00
363	10205	2019	1.00	1.50	0.00	0.00
364	10206	2019	1.00	15.00	0.00	0.00
365	10213	2019	1.00	10.00	0.00	0.00
366	10208	2019	1.00	5.00	0.00	0.00
367	10207	2019	1.00	10.00	0.00	0.00
368	10209	2019	1.00	10.00	0.00	0.00
369	10210	2019	1.00	25.00	0.00	0.00
370	10211	2019	1.00	35.00	0.00	0.00
371	10212	2019	1.00	35.00	0.00	0.00
372	10214	2019	27000.00	0.00	0.00	0.00
373	10119	2020	1.00	15.00	0.00	0.00
374	10120	2020	1.00	15.00	0.00	0.00
375	10121	2020	1.00	15.00	0.00	0.00
376	10122	2020	1.00	15.00	0.00	0.00
377	10123	2020	1.00	50.00	0.00	0.00
378	10126	2020	1.00	2.50	0.00	0.00
379	10191	2020	1.00	20.00	0.00	0.00
380	10189	2020	1.00	25.00	0.00	0.00
381	10139	2020	1.00	10.00	0.00	0.00
382	10145	2020	1.00	25.00	0.00	0.00
383	10146	2020	1.00	50.00	0.00	0.00
384	10147	2020	1.00	15.00	0.00	0.00
385	10148	2020	1.00	15.00	0.00	0.00
386	10150	2020	1.00	15.00	0.00	0.00
387	10151	2020	1.00	15.00	0.00	0.00
388	10152	2020	1.00	15.00	0.00	0.00
389	10153	2020	1.00	15.00	0.00	0.00
390	10154	2020	1.00	25.00	0.00	0.00
391	10155	2020	1.00	25.00	0.00	0.00
392	10156	2020	1.00	25.00	0.00	0.00
393	10157	2020	1.00	15.00	0.00	0.00
394	10101	2020	1.00	10.00	0.00	0.00
395	10102	2020	1.00	10.00	0.00	0.00
396	10103	2020	1.00	10.00	0.00	0.00
397	10104	2020	1.00	10.00	0.00	0.00
398	10105	2020	1.00	10.00	0.00	0.00
399	10106	2020	1.00	10.00	0.00	0.00
400	10107	2020	1.00	10.00	0.00	0.00
401	10108	2020	1.00	10.00	0.00	0.00
402	10109	2020	1.00	10.00	0.00	0.00
403	10111	2020	1.00	10.00	0.00	0.00
404	10112	2020	1.00	10.00	0.00	0.00
405	10113	2020	1.00	10.00	0.00	0.00
406	10114	2020	1.00	10.00	0.00	0.00
407	10115	2020	1.00	10.00	0.00	0.00
408	10116	2020	1.00	10.00	0.00	0.00
409	10118	2020	1.00	10.00	0.00	0.00
410	10119	2020	1.00	15.00	0.00	0.00
411	10120	2020	1.00	15.00	0.00	0.00
412	10121	2020	1.00	15.00	0.00	0.00
413	10122	2020	1.00	15.00	0.00	0.00
414	10123	2020	1.00	50.00	0.00	0.00
415	10125	2020	100000.00	10.00	0.00	0.00
416	10126	2020	1.00	2.50	0.00	0.00
417	10127	2020	3000.00	0.00	0.00	0.00
418	10132	2020	1.00	10.00	0.00	0.00
419	10130	2020	93750.00	0.00	0.00	0.00
420	10129	2020	0.00	0.00	0.00	0.00
421	10128	2020	3250.00	0.00	0.00	0.00
422	10191	2020	1.00	20.00	0.00	0.00
423	10137	2020	1.00	10.00	0.00	0.00
424	10138	2020	1.00	10.00	0.00	0.00
425	10139	2020	1.00	10.00	0.00	0.00
426	10140	2020	1.00	10.00	0.00	0.00
427	10141	2020	1.00	10.00	0.00	0.00
428	10145	2020	1.00	25.00	0.00	0.00
429	10143	2020	1.00	10.00	0.00	0.00
430	10144	2020	1.00	10.00	0.00	0.00
431	10146	2020	1.00	50.00	0.01	0.00
432	10147	2020	1.00	15.00	0.00	0.00
433	10148	2020	1.00	15.00	0.00	0.00
434	10149	2020	1.00	0.00	0.00	0.00
435	10150	2020	1.00	15.00	0.00	0.00
436	10151	2020	1.00	15.00	0.00	0.00
437	10152	2020	1.00	15.00	0.00	0.00
438	10153	2020	1.00	15.00	0.00	0.00
439	10154	2020	1.00	25.00	0.00	0.00
440	10155	2020	1.00	25.00	0.00	0.00
441	10156	2020	1.00	25.00	0.00	0.00
442	10157	2020	1.00	15.00	0.00	0.00
443	10158	2020	50000.00	0.00	0.00	0.00
444	10159	2020	20000.00	0.00	0.00	0.00
445	10160	2020	132500.00	0.00	0.00	0.00
446	10161	2020	25000.00	0.00	0.00	0.00
464	10205	2020	1.00	1.50	0.00	0.00
465	10206	2020	1.00	15.00	0.00	0.00
466	10213	2020	1.00	10.00	0.00	0.00
467	10208	2020	1.00	5.00	0.00	0.00
468	10207	2020	1.00	10.00	0.00	0.00
469	10209	2020	1.00	10.00	0.00	0.00
470	10210	2020	1.00	25.00	0.00	0.00
471	10211	2020	1.00	35.00	0.00	0.00
472	10212	2020	1.00	35.00	0.00	0.00
473	10214	2020	27000.00	0.00	0.00	0.00
474	10214	2021	27000.00	0.00	0.00	0.00
475	10212	2021	1.00	35.00	0.00	0.00
476	10211	2021	1.00	35.00	0.00	0.00
477	10210	2021	1.00	25.00	0.00	0.00
478	10209	2021	1.00	10.00	0.00	0.00
479	10207	2021	1.00	10.00	0.00	0.00
480	10208	2021	1.00	5.00	0.00	0.00
481	10213	2021	1.00	10.00	0.00	0.00
482	10206	2021	1.00	15.00	0.00	0.00
483	10205	2021	1.00	1.50	0.00	0.00
501	10161	2021	25000.00	0.00	0.00	0.00
502	10160	2021	132500.00	0.00	0.00	0.00
503	10159	2021	20000.00	0.00	0.00	0.00
504	10158	2021	50000.00	0.00	0.00	0.00
505	10157	2021	1.00	15.00	0.00	0.00
506	10156	2021	1.00	25.00	0.00	0.00
507	10155	2021	1.00	25.00	0.00	0.00
508	10154	2021	1.00	25.00	0.00	0.00
509	10153	2021	1.00	15.00	0.00	0.00
510	10152	2021	1.00	15.00	0.00	0.00
511	10151	2021	1.00	15.00	0.00	0.00
512	10150	2021	1.00	15.00	0.00	0.00
513	10149	2021	1.00	0.00	0.00	0.00
514	10148	2021	1.00	15.00	0.00	0.00
515	10147	2021	1.00	15.00	0.00	0.00
516	10146	2021	1.00	50.00	0.01	0.00
517	10144	2021	1.00	10.00	0.00	0.00
518	10143	2021	1.00	10.00	0.00	0.00
519	10145	2021	1.00	25.00	0.00	0.00
520	10141	2021	1.00	10.00	0.00	0.00
521	10140	2021	1.00	10.00	0.00	0.00
522	10139	2021	1.00	10.00	0.00	0.00
523	10138	2021	1.00	10.00	0.00	0.00
524	10137	2021	1.00	10.00	0.00	0.00
525	10191	2021	1.00	20.00	0.00	0.00
526	10128	2021	3250.00	0.00	0.00	0.00
527	10129	2021	0.00	0.00	0.00	0.00
528	10130	2021	93750.00	0.00	0.00	0.00
529	10132	2021	1.00	10.00	0.00	0.00
530	10127	2021	3000.00	0.00	0.00	0.00
531	10126	2021	1.00	2.50	0.00	0.00
532	10125	2021	100000.00	10.00	0.00	0.00
533	10123	2021	1.00	50.00	0.00	0.00
534	10122	2021	1.00	15.00	0.00	0.00
535	10121	2021	1.00	15.00	0.00	0.00
536	10120	2021	1.00	15.00	0.00	0.00
537	10119	2021	1.00	15.00	0.00	0.00
538	10118	2021	1.00	10.00	0.00	0.00
539	10116	2021	1.00	10.00	0.00	0.00
540	10115	2021	1.00	10.00	0.00	0.00
541	10114	2021	1.00	10.00	0.00	0.00
542	10113	2021	1.00	10.00	0.00	0.00
543	10112	2021	1.00	10.00	0.00	0.00
544	10111	2021	1.00	10.00	0.00	0.00
545	10109	2021	1.00	10.00	0.00	0.00
546	10108	2021	1.00	10.00	0.00	0.00
547	10107	2021	1.00	10.00	0.00	0.00
548	10106	2021	1.00	10.00	0.00	0.00
549	10105	2021	1.00	10.00	0.00	0.00
550	10104	2021	1.00	10.00	0.00	0.00
551	10103	2021	1.00	10.00	0.00	0.00
552	10102	2021	1.00	10.00	0.00	0.00
553	10101	2021	1.00	10.00	0.00	0.00
554	10157	2021	1.00	15.00	0.00	0.00
555	10156	2021	1.00	25.00	0.00	0.00
556	10155	2021	1.00	25.00	0.00	0.00
557	10154	2021	1.00	25.00	0.00	0.00
558	10153	2021	1.00	15.00	0.00	0.00
559	10152	2021	1.00	15.00	0.00	0.00
560	10151	2021	1.00	15.00	0.00	0.00
570	10150	2021	1.00	15.00	0.00	0.00
571	10148	2021	1.00	15.00	0.00	0.00
572	10147	2021	1.00	15.00	0.00	0.00
573	10146	2021	1.00	50.00	0.00	0.00
574	10145	2021	1.00	25.00	0.00	0.00
575	10139	2021	1.00	10.00	0.00	0.00
576	10189	2021	1.00	25.00	0.00	0.00
577	10191	2021	1.00	20.00	0.00	0.00
578	10126	2021	1.00	2.50	0.00	0.00
579	10123	2021	1.00	50.00	0.00	0.00
600	10122	2021	1.00	15.00	0.00	0.00
601	10121	2021	1.00	15.00	0.00	0.00
602	10120	2021	1.00	15.00	0.00	0.00
603	10119	2021	1.00	15.00	0.00	0.00
604	10189	2021	1.00	25.00	0.00	0.00
605	10189	2021	1.00	25.00	0.00	0.00
606	10189	2021	1.00	25.00	0.00	0.00
607	10189	2021	1.00	25.00	0.00	0.00
3266	10111	0	1.00	10.00	0.00	0.00
608	10214	2022	27000.00	0.00	0.00	0.00
609	10212	2022	1.00	35.00	0.00	0.00
610	10211	2022	1.00	35.00	0.00	0.00
611	10210	2022	1.00	25.00	0.00	0.00
612	10209	2022	1.00	10.00	0.00	0.00
613	10207	2022	1.00	10.00	0.00	0.00
614	10208	2022	1.00	5.00	0.00	0.00
615	10213	2022	1.00	10.00	0.00	0.00
616	10206	2022	1.00	15.00	0.00	0.00
617	10205	2022	1.00	1.50	0.00	0.00
635	10161	2022	25000.00	0.00	0.00	0.00
636	10160	2022	132500.00	0.00	0.00	0.00
637	10159	2022	20000.00	0.00	0.00	0.00
638	10158	2022	50000.00	0.00	0.00	0.00
639	10157	2022	1.00	15.00	0.00	0.00
640	10156	2022	1.00	25.00	0.00	0.00
641	10155	2022	1.00	25.00	0.00	0.00
642	10154	2022	1.00	25.00	0.00	0.00
643	10153	2022	1.00	15.00	0.00	0.00
644	10152	2022	1.00	15.00	0.00	0.00
645	10151	2022	1.00	15.00	0.00	0.00
646	10150	2022	1.00	15.00	0.00	0.00
647	10149	2022	1.00	0.00	0.00	0.00
648	10148	2022	1.00	15.00	0.00	0.00
649	10147	2022	1.00	15.00	0.00	0.00
650	10146	2022	1.00	50.00	0.01	0.00
651	10144	2022	1.00	10.00	0.00	0.00
652	10143	2022	1.00	10.00	0.00	0.00
653	10145	2022	1.00	25.00	0.00	0.00
654	10141	2022	1.00	10.00	0.00	0.00
655	10140	2022	1.00	10.00	0.00	0.00
656	10139	2022	1.00	10.00	0.00	0.00
657	10138	2022	1.00	10.00	0.00	0.00
658	10137	2022	1.00	10.00	0.00	0.00
659	10191	2022	1.00	20.00	0.00	0.00
660	10128	2022	3250.00	0.00	0.00	0.00
661	10129	2022	0.00	0.00	0.00	0.00
662	10130	2022	93750.00	0.00	0.00	0.00
663	10132	2022	1.00	10.00	0.00	0.00
664	10127	2022	3000.00	0.00	0.00	0.00
665	10126	2022	1.00	2.50	0.00	0.00
666	10125	2022	100000.00	10.00	0.00	0.00
667	10123	2022	1.00	50.00	0.00	0.00
668	10122	2022	1.00	15.00	0.00	0.00
669	10121	2022	1.00	15.00	0.00	0.00
670	10120	2022	1.00	15.00	0.00	0.00
671	10119	2022	1.00	15.00	0.00	0.00
672	10118	2022	1.00	10.00	0.00	0.00
673	10116	2022	1.00	10.00	0.00	0.00
674	10115	2022	1.00	10.00	0.00	0.00
675	10114	2022	1.00	10.00	0.00	0.00
676	10113	2022	1.00	10.00	0.00	0.00
677	10112	2022	1.00	10.00	0.00	0.00
678	10111	2022	1.00	10.00	0.00	0.00
679	10107	2022	1.00	10.00	0.00	0.00
680	10106	2022	1.00	10.00	0.00	0.00
681	10105	2022	1.00	10.00	0.00	0.00
682	10104	2022	1.00	10.00	0.00	0.00
683	10103	2022	1.00	10.00	0.00	0.00
684	10102	2022	1.00	10.00	0.00	0.00
685	10101	2022	1.00	10.00	0.00	0.00
686	10157	2022	1.00	15.00	0.00	0.00
687	10156	2022	1.00	25.00	0.00	0.00
688	10155	2022	1.00	25.00	0.00	0.00
689	10154	2022	1.00	25.00	0.00	0.00
690	10153	2022	1.00	15.00	0.00	0.00
691	10152	2022	1.00	15.00	0.00	0.00
692	10151	2022	1.00	15.00	0.00	0.00
693	10150	2022	1.00	15.00	0.00	0.00
694	10148	2022	1.00	15.00	0.00	0.00
697	10147	2022	1.00	15.00	0.00	0.00
698	10146	2022	1.00	50.00	0.00	0.00
699	10145	2022	1.00	25.00	0.00	0.00
700	10139	2022	1.00	10.00	0.00	0.00
701	10189	2022	1.00	25.00	0.00	0.00
702	10191	2022	1.00	20.00	0.00	0.00
703	10126	2022	1.00	2.50	0.00	0.00
704	10123	2022	1.00	50.00	0.00	0.00
705	10122	2022	1.00	15.00	0.00	0.00
706	10121	2022	1.00	15.00	0.00	0.00
707	10120	2022	1.00	15.00	0.00	0.00
708	10119	2022	1.00	15.00	0.00	0.00
709	10189	2022	1.00	25.00	0.00	0.00
710	10189	2022	1.00	25.00	0.00	0.00
711	10189	2022	1.00	25.00	0.00	0.00
712	10189	2022	1.00	25.00	0.00	0.00
\.


--
-- Name: TarifPajak_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: dexwip
--

SELECT pg_catalog.setval('public."TarifPajak_Id_seq"', 3266, false);


--
-- PostgreSQL database dump complete
--

