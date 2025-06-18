DROP TABLE link_unit_infos;
DROP TABLE embed_unit_infos;
DROP TABLE portal_unit_infos;

DELETE FROM items WHERE id != '414b5703-91d1-42fc-a007-36dd8f25e329';
DELETE FROM unit_types WHERE name != 'static';

DROP TABLE occupied_positions;
DROP TABLE units;
DROP TABLE items;
DROP TABLE unit_types;
