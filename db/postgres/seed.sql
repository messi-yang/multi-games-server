INSERT INTO unit_types (name)
    VALUES
        ('static'),
        ('portal'),
        ('fence'),
        ('link')
    ON CONFLICT (name) DO NOTHING;



INSERT INTO items (id, name, traversable, created_at, updated_at, thumbnail_src, model_sources, compatible_unit_type)
    VALUES
        ('fb030c79-25a4-48c6-955a-b1188ed409f7', 'fence', FALSE, Now(), Now(), '/asset/item/fence/thumbnail.png', '{/asset/item/fence/model_1.gltf, /asset/item/fence/model_2.gltf, /asset/item/fence/model_3.gltf, /asset/item/fence/model_4.gltf}', 'fence'),
        ('c64f732c-9494-4693-b17b-43a2736aa67b', 'portal', TRUE, Now(), Now(), '/asset/item/portal/thumbnail.png', '{/asset/item/portal/model.gltf}', 'portal'),
        ('1b9ba8b1-c13e-4524-bddc-7cc6d981ee2c', 'trash bin', FALSE, Now(), Now(), '/asset/item/trash_bin/thumbnail.png', '{/asset/item/trash_bin/model.gltf}', 'static'),
        ('2b6ab30d-0a2a-4424-b245-99ec2c301844', 'chair A', FALSE, Now(), Now(), '/asset/item/chair/thumbnail.png', '{/asset/item/chair/model.gltf}', 'static'),
        ('34af14ab-42c5-4c55-a787-44f32012354e', 'torch', TRUE, Now(), Now(), '/asset/item/torch/thumbnail.png', '{/asset/item/torch/model.gltf}', 'static'),
        ('3c28537a-80c2-4ac1-917b-b1cd517c6b5e', 'stone', FALSE, Now(), Now(), '/asset/item/stone/thumbnail.png', '{/asset/item/stone/model.gltf}', 'static'),
        ('414b5703-91d1-42fc-a007-36dd8f25e329', 'tree', FALSE, Now(), Now(), '/asset/item/tree/thumbnail.png', '{/asset/item/tree/model.gltf}', 'static'),
        ('41de86e6-07a1-4a5d-ba6f-152d07f3ba1e', 'fan', FALSE, Now(), Now(), '/asset/item/fan/thumbnail.png', '{/asset/item/fan/model.gltf}', 'static'),
        ('52bdd7d3-799d-42dd-a2dc-cd438101cfca', 'chair B', FALSE, Now(), Now(), '/asset/item/chair_2/thumbnail.png', '{/asset/item/chair_2/model.gltf}', 'static'),
        ('6f127ae8-f1f8-4ff3-8148-fa8d2fef307a', 'ghost', FALSE, Now(), Now(), '/asset/item/ghost/thumbnail.png', '{/asset/item/ghost/model.gltf}', 'static'),
        ('c0a15d4a-24b7-4a81-8a39-9bbf4c7d6ccf', 'grass', TRUE, Now(), Now(), '/asset/item/grass/thumbnail.png', '{/asset/item/grass/model.gltf}', 'static'),
        ('d4d0850a-dbe0-451c-9e50-6ac280108d1c', 'cone', FALSE, Now(), Now(), '/asset/item/cone/thumbnail.png', '{/asset/item/cone/model.gltf}', 'static'),
        ('e495468b-e662-49cb-bc5b-96db204ad9d8', 'box', FALSE, Now(), Now(), '/asset/item/box/thumbnail.png', '{/asset/item/box/model.gltf}', 'static'),
        ('fb9d06f8-5d6d-4fa9-bdc5-ab760d55a442', 'potted plant', FALSE, Now(), Now(), '/asset/item/potted_plant/thumbnail.png', '{/asset/item/potted_plant/model.gltf}', 'static'), 
        ('ec3bf2ba-6e38-4b68-8bb2-15ef4e2a60a3', 'link', FALSE, Now(), Now(), '/asset/item/link/thumbnail.png', '{/asset/item/link/model.gltf}', 'link'),
        ('bb767e60-a5ae-43e9-ae0f-2aea00a1684f', 'tetris square i', FALSE, Now(), Now(), '/asset/item/tetris_square_i/thumbnail.png', '{/asset/item/tetris_square_i/model.gltf}', 'static'),
        ('b8155233-7850-49ab-a155-46fdc6687210', 'tetris square j', FALSE, Now(), Now(), '/asset/item/tetris_square_j/thumbnail.png', '{/asset/item/tetris_square_j/model.gltf}', 'static'),
        ('4fb0af0c-d27c-4ca9-bf5d-b8eed40ac869', 'tetris square l', FALSE, Now(), Now(), '/asset/item/tetris_square_l/thumbnail.png', '{/asset/item/tetris_square_l/model.gltf}', 'static'),
        ('ea819f56-b710-441b-89ea-9941b8bb75a0', 'tetris square o', FALSE, Now(), Now(), '/asset/item/tetris_square_o/thumbnail.png', '{/asset/item/tetris_square_o/model.gltf}', 'static'),
        ('f00d5865-eedc-4b3a-ab46-d973a9e02010', 'tetris square s', FALSE, Now(), Now(), '/asset/item/tetris_square_s/thumbnail.png', '{/asset/item/tetris_square_s/model.gltf}', 'static'),
        ('13aeff97-4f5f-4c9c-9a4c-2a92d4ae6cf5', 'tetris square z', FALSE, Now(), Now(), '/asset/item/tetris_square_z/thumbnail.png', '{/asset/item/tetris_square_z/model.gltf}', 'static'),
        ('2b694709-627a-4dfb-8f47-9435d46ef28f', 'tetris square t', FALSE, Now(), Now(), '/asset/item/tetris_square_t/thumbnail.png', '{/asset/item/tetris_square_t/model.gltf}', 'static')
    ON CONFLICT (id) DO UPDATE
        SET
            name = EXCLUDED.name,
            traversable = EXCLUDED.traversable,
            updated_at = EXCLUDED.updated_at,
            thumbnail_src = EXCLUDED.thumbnail_src,
            model_sources = EXCLUDED.model_sources,
            compatible_unit_type = EXCLUDED.compatible_unit_type;
