DROP TABLE embed_unit_infos;

ALTER TABLE link_unit_infos ALTER COLUMN url DROP NOT NULL;
ALTER TABLE link_unit_infos DROP CONSTRAINT link_unit_infos_pkey;
ALTER TABLE link_unit_infos DROP CONSTRAINT link_unit_infos_world_id_fkey;