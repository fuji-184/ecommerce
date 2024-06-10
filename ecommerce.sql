--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: detail_keranjang; Type: TABLE; Schema: public; Owner: u0_a1887
--

CREATE TABLE public.detail_keranjang (
    id integer NOT NULL,
    id_keranjang integer NOT NULL,
    id_item integer NOT NULL,
    jumlah integer NOT NULL,
    total_harga numeric(10,2) NOT NULL
);


ALTER TABLE public.detail_keranjang OWNER TO u0_a1887;

--
-- Name: detail_keranjang_id_seq; Type: SEQUENCE; Schema: public; Owner: u0_a1887
--

CREATE SEQUENCE public.detail_keranjang_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.detail_keranjang_id_seq OWNER TO u0_a1887;

--
-- Name: detail_keranjang_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: u0_a1887
--

ALTER SEQUENCE public.detail_keranjang_id_seq OWNED BY public.detail_keranjang.id;


--
-- Name: items; Type: TABLE; Schema: public; Owner: u0_a1887
--

CREATE TABLE public.items (
    id integer NOT NULL,
    nama character varying(50) NOT NULL,
    harga double precision NOT NULL,
    stok integer NOT NULL,
    gambar character varying(100)
);


ALTER TABLE public.items OWNER TO u0_a1887;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: u0_a1887
--

CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.items_id_seq OWNER TO u0_a1887;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: u0_a1887
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- Name: keranjang_belanja; Type: TABLE; Schema: public; Owner: u0_a1887
--

CREATE TABLE public.keranjang_belanja (
    id integer NOT NULL,
    id_pembeli integer NOT NULL,
    tanggal timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.keranjang_belanja OWNER TO u0_a1887;

--
-- Name: keranjang_belanja_id_seq; Type: SEQUENCE; Schema: public; Owner: u0_a1887
--

CREATE SEQUENCE public.keranjang_belanja_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.keranjang_belanja_id_seq OWNER TO u0_a1887;

--
-- Name: keranjang_belanja_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: u0_a1887
--

ALTER SEQUENCE public.keranjang_belanja_id_seq OWNED BY public.keranjang_belanja.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: u0_a1887
--

CREATE TABLE public.users (
    id integer NOT NULL,
    nama character varying(50) NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(50),
    is_admin boolean DEFAULT false,
    saldo double precision
);


ALTER TABLE public.users OWNER TO u0_a1887;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: u0_a1887
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: detail_keranjang id; Type: DEFAULT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.detail_keranjang ALTER COLUMN id SET DEFAULT nextval('public.detail_keranjang_id_seq'::regclass);


--
-- Name: items id; Type: DEFAULT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- Name: keranjang_belanja id; Type: DEFAULT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.keranjang_belanja ALTER COLUMN id SET DEFAULT nextval('public.keranjang_belanja_id_seq'::regclass);


--
-- Data for Name: detail_keranjang; Type: TABLE DATA; Schema: public; Owner: u0_a1887
--

COPY public.detail_keranjang (id, id_keranjang, id_item, jumlah, total_harga) FROM stdin;
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: u0_a1887
--

COPY public.items (id, nama, harga, stok, gambar) FROM stdin;
1	jdndnd	649494	6464644	/gambar.jpg
2	Produk 1	100	50	/gambar.jpg
3	Produk 2	200	40	/gambar.jpg
4	Produk 3	150	30	/gambar.jpg
5	Produk 4	120	20	/gambar.jpg
6	Produk 5	180	10	/gambar.jpg
7	Produk 6	220	5	/gambar.jpg
8	Produk 7	250	15	/gambar.jpg
9	Produk 8	300	25	/gambar.jpg
10	Produk 9	350	35	/gambar.jpg
11	Produk 10	400	45	/gambar.jpg
12	Produk 11	450	55	/gambar.jpg
13	Produk 12	500	60	/gambar.jpg
14	Produk 13	550	65	/gambar.jpg
15	Produk 14	600	70	/gambar.jpg
16	Produk 15	650	75	/gambar.jpg
17	Produk 16	700	80	/gambar.jpg
18	Produk 17	750	85	/gambar.jpg
19	Produk 18	800	90	/gambar.jpg
20	Produk 19	850	95	/gambar.jpg
21	Produk 20	900	100	/gambar.jpg
\.


--
-- Data for Name: keranjang_belanja; Type: TABLE DATA; Schema: public; Owner: u0_a1887
--

COPY public.keranjang_belanja (id, id_pembeli, tanggal) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: u0_a1887
--

COPY public.users (id, nama, username, password, is_admin, saldo) FROM stdin;
4	ss	nsns	nanana	f	\N
8	nnn	nsnssn	nbb	f	\N
12	nnnbwnnw	nsnssnnssbsn	nbb	f	\N
13	mfnsndm	krnendn	nsns	f	\N
14	jsnssn	nnsns	nsnssn	f	\N
15	jsnssnneen	nnsnsejjedn	nsnssn	f	\N
16	mxmxxnxn	mxndnx	nnn	f	\N
17	mdnsnn	mdsnsnsns	mssbsbs	f	\N
18		kkdns	mdbssnsns	f	\N
1	Fujiii	fuji	fuji	t	999999999
\.


--
-- Name: detail_keranjang_id_seq; Type: SEQUENCE SET; Schema: public; Owner: u0_a1887
--

SELECT pg_catalog.setval('public.detail_keranjang_id_seq', 1, false);


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: u0_a1887
--

SELECT pg_catalog.setval('public.items_id_seq', 21, true);


--
-- Name: keranjang_belanja_id_seq; Type: SEQUENCE SET; Schema: public; Owner: u0_a1887
--

SELECT pg_catalog.setval('public.keranjang_belanja_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: u0_a1887
--

SELECT pg_catalog.setval('public.users_id_seq', 20, true);


--
-- Name: detail_keranjang detail_keranjang_pkey; Type: CONSTRAINT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.detail_keranjang
    ADD CONSTRAINT detail_keranjang_pkey PRIMARY KEY (id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: keranjang_belanja keranjang_belanja_pkey; Type: CONSTRAINT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.keranjang_belanja
    ADD CONSTRAINT keranjang_belanja_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: detail_keranjang detail_keranjang_id_item_fkey; Type: FK CONSTRAINT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.detail_keranjang
    ADD CONSTRAINT detail_keranjang_id_item_fkey FOREIGN KEY (id_item) REFERENCES public.items(id);


--
-- Name: detail_keranjang detail_keranjang_id_keranjang_fkey; Type: FK CONSTRAINT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.detail_keranjang
    ADD CONSTRAINT detail_keranjang_id_keranjang_fkey FOREIGN KEY (id_keranjang) REFERENCES public.keranjang_belanja(id);


--
-- Name: keranjang_belanja keranjang_belanja_id_pembeli_fkey; Type: FK CONSTRAINT; Schema: public; Owner: u0_a1887
--

ALTER TABLE ONLY public.keranjang_belanja
    ADD CONSTRAINT keranjang_belanja_id_pembeli_fkey FOREIGN KEY (id_pembeli) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

