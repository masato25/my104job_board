--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2 (Debian 11.2-1.pgdg90+1)
-- Dumped by pg_dump version 11.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: jobs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.jobs (
    id uuid NOT NULL,
    company_name character varying(255) NOT NULL,
    job_name character varying(255) NOT NULL,
    area character varying(255) NOT NULL,
    welfare text NOT NULL,
    j character varying(255) NOT NULL,
    c character varying(255) NOT NULL,
    job_cat character varying(255) NOT NULL,
    sal_low integer NOT NULL,
    sal_high integer NOT NULL,
    description text NOT NULL,
    other_des text NOT NULL,
    profile text NOT NULL,
    link character varying(255) NOT NULL,
    cat character varying(255) NOT NULL,
    manager boolean NOT NULL,
    need_on_bt boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.jobs OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: jobs jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jobs
    ADD CONSTRAINT jobs_pkey PRIMARY KEY (id);


--
-- Name: job_j_c_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX job_j_c_unique ON public.jobs USING btree (j, c);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

