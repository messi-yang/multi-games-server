--
-- PostgreSQL database dump
--

-- Dumped from database version 15.0
-- Dumped by pg_dump version 15.0

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
-- Name: unit_type; Type: TYPE; Schema: public; Owner: main
--

CREATE TYPE public.unit_type AS ENUM (
    'static',
    'portal'
);


ALTER TYPE public.unit_type OWNER TO main;

--
-- Name: world_role; Type: TYPE; Schema: public; Owner: main
--

CREATE TYPE public.world_role AS ENUM (
    'owner',
    'admin',
    'editor',
    'viewer'
);


ALTER TYPE public.world_role OWNER TO main;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: items; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.items (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    traversable boolean NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    thumbnail_src character varying(255) NOT NULL,
    compatible_unit_type public.unit_type DEFAULT 'static'::public.unit_type NOT NULL,
    model_sources character varying(150)[] DEFAULT '{}'::character varying[] NOT NULL
);


ALTER TABLE public.items OWNER TO main;

--
-- Name: portal_unit_infos; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.portal_unit_infos (
    world_id uuid NOT NULL,
    target_pos_x integer,
    target_pos_z integer,
    id uuid DEFAULT gen_random_uuid() NOT NULL
);


ALTER TABLE public.portal_unit_infos OWNER TO main;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO main;

--
-- Name: units; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.units (
    world_id uuid NOT NULL,
    pos_x integer NOT NULL,
    pos_z integer NOT NULL,
    item_id uuid NOT NULL,
    direction integer NOT NULL,
    type public.unit_type DEFAULT 'static'::public.unit_type NOT NULL,
    info_id uuid,
    info_snapshot jsonb NOT NULL
);


ALTER TABLE public.units OWNER TO main;

--
-- Name: users; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email_address character varying(255) NOT NULL,
    username character varying(50) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    friendly_name character varying(20) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.users OWNER TO main;

--
-- Name: world_accounts; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.world_accounts (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    worlds_count integer DEFAULT 0 NOT NULL,
    worlds_count_limit integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.world_accounts OWNER TO main;

--
-- Name: world_members; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.world_members (
    id uuid NOT NULL,
    world_id uuid NOT NULL,
    user_id uuid NOT NULL,
    role public.world_role NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.world_members OWNER TO main;

--
-- Name: worlds; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.worlds (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    bound_from_x integer DEFAULT '-50'::integer,
    bound_from_z integer DEFAULT '-50'::integer,
    bound_to_x integer DEFAULT 50,
    bound_to_z integer DEFAULT 50
);


ALTER TABLE public.worlds OWNER TO main;

--
-- Name: world_accounts game_users_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.world_accounts
    ADD CONSTRAINT game_users_pkey PRIMARY KEY (id);


--
-- Name: world_accounts game_users_user_id_key; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.world_accounts
    ADD CONSTRAINT game_users_user_id_key UNIQUE (user_id);


--
-- Name: items items_name_key; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_name_key UNIQUE (name);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: portal_unit_infos portal_unit_infos_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.portal_unit_infos
    ADD CONSTRAINT portal_unit_infos_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: units unique_world_id_pos_x_pos_z; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT unique_world_id_pos_x_pos_z UNIQUE (world_id, pos_x, pos_z);


--
-- Name: users users_email_address_key; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_address_key UNIQUE (email_address);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: world_members world_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.world_members
    ADD CONSTRAINT world_roles_pkey PRIMARY KEY (id);


--
-- Name: world_members world_roles_world_id_user_id_key; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.world_members
    ADD CONSTRAINT world_roles_world_id_user_id_key UNIQUE (world_id, user_id);


--
-- Name: worlds worlds_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.worlds
    ADD CONSTRAINT worlds_pkey PRIMARY KEY (id);


--
-- Name: item_compatible_unit_type; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX item_compatible_unit_type ON public.items USING btree (compatible_unit_type);


--
-- Name: portal_unit_infos_world_id_target_pos_x_target_pos_z; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX portal_unit_infos_world_id_target_pos_x_target_pos_z ON public.portal_unit_infos USING btree (world_id, target_pos_x, target_pos_z);


--
-- Name: unit_type; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX unit_type ON public.units USING btree (type);


--
-- Name: unit_world_id_pos_x_pos_z; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX unit_world_id_pos_x_pos_z ON public.units USING btree (world_id, pos_x, pos_z);


--
-- Name: world_roles_world_id_user_id; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX world_roles_world_id_user_id ON public.world_members USING btree (world_id, user_id);


--
-- Name: units fk_item; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES public.items(id);


--
-- Name: world_accounts fk_user; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.world_accounts
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: worlds fk_user; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.worlds
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: world_members fk_user; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.world_members
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: units fk_world; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT fk_world FOREIGN KEY (world_id) REFERENCES public.worlds(id);


--
-- Name: world_members fk_world; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.world_members
    ADD CONSTRAINT fk_world FOREIGN KEY (world_id) REFERENCES public.worlds(id);


--
-- Name: portal_unit_infos fk_world; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.portal_unit_infos
    ADD CONSTRAINT fk_world FOREIGN KEY (world_id) REFERENCES public.worlds(id);


--
-- PostgreSQL database dump complete
--

