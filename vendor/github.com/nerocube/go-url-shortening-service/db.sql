-- docker-compose exec --user postgres app_postgres pg_dump "tiny_service" > db.sql
-- PostgreSQL database dump
--

-- Dumped from database version 10.3 (Debian 10.3-1.pgdg90+1)
-- Dumped by pg_dump version 10.3 (Debian 10.3-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: url_map; Type: TABLE; Schema: public; Owner: nero
--

CREATE TABLE public.url_map (
    original_url text NOT NULL,
    tiny_url text NOT NULL,
    created_at time with time zone NOT NULL
);


ALTER TABLE public.url_map OWNER TO nero;

--
-- Data for Name: url_map; Type: TABLE DATA; Schema: public; Owner: nero
--

COPY public.url_map (original_url, tiny_url, created_at) FROM stdin;
https://github.com/	QU5XzV	17:20:24.239648+00
\.


--
-- Name: url_map url_map_pkey; Type: CONSTRAINT; Schema: public; Owner: nero
--

ALTER TABLE ONLY public.url_map
    ADD CONSTRAINT url_map_pkey PRIMARY KEY (original_url);


--
-- PostgreSQL database dump complete
--

