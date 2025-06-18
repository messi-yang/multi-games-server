CREATE TABLE public.unit_types (
    name character varying(20) NOT NULL
);

ALTER TABLE ONLY public.unit_types
    ADD CONSTRAINT unit_types_pkey PRIMARY KEY (name);

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

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_name_key UNIQUE (name);

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_compatible_unit_type_unit_types_name FOREIGN KEY (compatible_unit_type) REFERENCES public.unit_types(name);

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

ALTER TABLE ONLY public.units
    ADD CONSTRAINT unique_world_id_pos_x_pos_z UNIQUE (world_id, pos_x, pos_z);

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_pkey PRIMARY KEY (id);

CREATE INDEX unit_world_id_pos_x_pos_z ON public.units USING btree (world_id, pos_x, pos_z);

ALTER TABLE ONLY public.units
    ADD CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES public.items(id);

ALTER TABLE ONLY public.units
    ADD CONSTRAINT fk_world FOREIGN KEY (world_id) REFERENCES public.worlds(id);

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_type_unit_types_name FOREIGN KEY (type) REFERENCES public.unit_types(name);

CREATE TABLE public.occupied_positions (
    world_id uuid NOT NULL,
    pos_x integer NOT NULL,
    pos_z integer NOT NULL,
    unit_id uuid NOT NULL
);

ALTER TABLE ONLY public.occupied_positions
    ADD CONSTRAINT occupied_positions_unique_world_id_pos_x_pos_z UNIQUE (world_id, pos_x, pos_z);

ALTER TABLE ONLY public.occupied_positions
    ADD CONSTRAINT occupied_positions_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES public.units(id);

ALTER TABLE ONLY public.occupied_positions
    ADD CONSTRAINT occupied_positions_world_id_fkey FOREIGN KEY (world_id) REFERENCES public.worlds(id);

INSERT INTO unit_types (name)
    VALUES
        ('static'),
        ('portal'),
        ('fence'),
        ('link'),
        ('embed'),
        ('color'),
        ('sign')
    ON CONFLICT (name) DO NOTHING;

CREATE TABLE public.portal_unit_infos (
    world_id uuid NOT NULL,
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    target_unit_id uuid
);

ALTER TABLE ONLY public.portal_unit_infos
    ADD CONSTRAINT portal_unit_infos_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.portal_unit_infos
    ADD CONSTRAINT fk_world FOREIGN KEY (world_id) REFERENCES public.worlds(id);

ALTER TABLE ONLY public.portal_unit_infos
    ADD CONSTRAINT portal_unit_info_fk_target_unit_id FOREIGN KEY (target_unit_id) REFERENCES public.units(id);

CREATE TABLE public.embed_unit_infos (
    id uuid NOT NULL,
    world_id uuid NOT NULL,
    embed_code character varying(2048)
);

ALTER TABLE ONLY public.embed_unit_infos
    ADD CONSTRAINT embed_unit_infos_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.embed_unit_infos
    ADD CONSTRAINT embed_unit_infos_world_id_fkey FOREIGN KEY (world_id) REFERENCES public.worlds(id);

CREATE TABLE public.link_unit_infos (
    world_id uuid NOT NULL,
    url character varying(2048) NOT NULL,
    id uuid NOT NULL
);

ALTER TABLE ONLY public.link_unit_infos
    ADD CONSTRAINT link_unit_infos_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.link_unit_infos
    ADD CONSTRAINT link_unit_infos_world_id_fkey FOREIGN KEY (world_id) REFERENCES public.worlds(id);
