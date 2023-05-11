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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: gamers; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.gamers (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.gamers OWNER TO main;

--
-- Name: items; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.items (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    traversable boolean NOT NULL,
    model_src character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    thumbnail_src character varying(255) NOT NULL
);


ALTER TABLE public.items OWNER TO main;

--
-- Name: players; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.players (
    id uuid NOT NULL,
    gamer_id uuid,
    world_id uuid NOT NULL,
    name character varying(50) NOT NULL,
    pos_x integer NOT NULL,
    pos_z integer NOT NULL,
    direction integer NOT NULL,
    held_item_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.players OWNER TO main;

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
    direction integer NOT NULL
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
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.users OWNER TO main;

--
-- Name: worlds; Type: TABLE; Schema: public; Owner: main
--

CREATE TABLE public.worlds (
    id uuid NOT NULL,
    gamer_id uuid NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.worlds OWNER TO main;

--
-- Name: gamers game_users_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.gamers
    ADD CONSTRAINT game_users_pkey PRIMARY KEY (id);


--
-- Name: gamers game_users_user_id_key; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.gamers
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
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (id);


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
-- Name: worlds worlds_pkey; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.worlds
    ADD CONSTRAINT worlds_pkey PRIMARY KEY (id);


--
-- Name: worlds worlds_user_id_key; Type: CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.worlds
    ADD CONSTRAINT worlds_user_id_key UNIQUE (gamer_id);


--
-- Name: players_world_id_pos_x_pos_z; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX players_world_id_pos_x_pos_z ON public.players USING btree (world_id, pos_x, pos_z);


--
-- Name: unit_world_id_pos_x_pos_z; Type: INDEX; Schema: public; Owner: main
--

CREATE INDEX unit_world_id_pos_x_pos_z ON public.units USING btree (world_id, pos_x, pos_z);


--
-- Name: worlds fk_game_user; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.worlds
    ADD CONSTRAINT fk_game_user FOREIGN KEY (gamer_id) REFERENCES public.gamers(id);


--
-- Name: players fk_gamer; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT fk_gamer FOREIGN KEY (gamer_id) REFERENCES public.gamers(id);


--
-- Name: units fk_item; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES public.items(id);


--
-- Name: players fk_item; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT fk_item FOREIGN KEY (held_item_id) REFERENCES public.items(id);


--
-- Name: gamers fk_user; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.gamers
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: units fk_world; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT fk_world FOREIGN KEY (world_id) REFERENCES public.worlds(id);


--
-- Name: players fk_world; Type: FK CONSTRAINT; Schema: public; Owner: main
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT fk_world FOREIGN KEY (world_id) REFERENCES public.worlds(id);


--
-- PostgreSQL database dump complete
--

