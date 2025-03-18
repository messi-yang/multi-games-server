--
-- PostgreSQL database dump
--

-- Dumped from database version 15.12
-- Dumped by pg_dump version 15.12

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
-- Name: embed_unit_infos; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.embed_unit_infos (
    id uuid NOT NULL,
    world_id uuid NOT NULL,
    embed_code character varying(2048)
);


ALTER TABLE public.embed_unit_infos OWNER TO main;

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
    model_sources character varying(150)[] DEFAULT '{}'::character varying[] NOT NULL,
    compatible_unit_type character varying(20) DEFAULT 'static'::character varying NOT NULL,
    dimension_width integer DEFAULT 1 NOT NULL,
    dimension_depth integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.items OWNER TO main;

--
-- Name: link_unit_infos; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.link_unit_infos (
    world_id uuid NOT NULL,
    url character varying(2048) NOT NULL,
    id uuid NOT NULL
);


ALTER TABLE public.link_unit_infos OWNER TO main;

--
-- Name: occupied_positions; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.occupied_positions (
    world_id uuid NOT NULL,
    pos_x integer NOT NULL,
    pos_z integer NOT NULL,
    unit_id uuid NOT NULL
);


ALTER TABLE public.occupied_positions OWNER TO main;

--
-- Name: portal_unit_infos; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.portal_unit_infos (
    world_id uuid NOT NULL,
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    target_unit_id uuid
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
-- Name: unit_types; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.unit_types (
    name character varying(20) NOT NULL
);


ALTER TABLE public.unit_types OWNER TO main;

--
-- Name: units; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.units (
    world_id uuid NOT NULL,
    pos_x integer NOT NULL,
    pos_z integer NOT NULL,
    item_id uuid NOT NULL,
    direction integer NOT NULL,
    id uuid NOT NULL,
    type character varying(20) DEFAULT 'static'::character varying NOT NULL,
    label character varying(20),
    dimension_width integer DEFAULT 1 NOT NULL,
    dimension_depth integer DEFAULT 1 NOT NULL,
    color character varying(7)
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
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.worlds OWNER TO main;

--
-- Name: embed_unit_infos embed_unit_infos_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.embed_unit_infos
    ADD CONSTRAINT embed_unit_infos_pkey PRIMARY KEY (id);


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
-- Name: link_unit_infos link_unit_infos_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.link_unit_infos
    ADD CONSTRAINT link_unit_infos_pkey PRIMARY KEY (id);


--
-- Name: occupied_positions occupied_positions_unique_world_id_pos_x_pos_z; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.occupied_positions
    ADD CONSTRAINT occupied_positions_unique_world_id_pos_x_pos_z UNIQUE (world_id, pos_x, pos_z);


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
-- Name: unit_types unit_types_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.unit_types
    ADD CONSTRAINT unit_types_pkey PRIMARY KEY (name);


--
-- Name: units units_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_pkey PRIMARY KEY (id);


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
-- Name: unit_world_id_pos_x_pos_z; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX unit_world_id_pos_x_pos_z ON public.units USING btree (world_id, pos_x, pos_z);


--
-- Name: world_roles_world_id_user_id; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX world_roles_world_id_user_id ON public.world_members USING btree (world_id, user_id);


--
-- Name: embed_unit_infos embed_unit_infos_world_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.embed_unit_infos
    ADD CONSTRAINT embed_unit_infos_world_id_fkey FOREIGN KEY (world_id) REFERENCES public.worlds(id);


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
-- Name: items items_compatible_unit_type_unit_types_name; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_compatible_unit_type_unit_types_name FOREIGN KEY (compatible_unit_type) REFERENCES public.unit_types(name);


--
-- Name: link_unit_infos link_unit_infos_world_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.link_unit_infos
    ADD CONSTRAINT link_unit_infos_world_id_fkey FOREIGN KEY (world_id) REFERENCES public.worlds(id);


--
-- Name: occupied_positions occupied_positions_unit_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.occupied_positions
    ADD CONSTRAINT occupied_positions_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES public.units(id);


--
-- Name: occupied_positions occupied_positions_world_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.occupied_positions
    ADD CONSTRAINT occupied_positions_world_id_fkey FOREIGN KEY (world_id) REFERENCES public.worlds(id);


--
-- Name: portal_unit_infos portal_unit_info_fk_target_unit_id; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.portal_unit_infos
    ADD CONSTRAINT portal_unit_info_fk_target_unit_id FOREIGN KEY (target_unit_id) REFERENCES public.units(id);


--
-- Name: units units_type_unit_types_name; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_type_unit_types_name FOREIGN KEY (type) REFERENCES public.unit_types(name);


--
-- PostgreSQL database dump complete
--

