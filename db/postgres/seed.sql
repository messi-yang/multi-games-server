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
        ('ec3bf2ba-6e38-4b68-8bb2-15ef4e2a60a3', 'link', FALSE, Now(), Now(), '/asset/item/link/thumbnail.png', '{/asset/item/link/model.gltf}', 'link')
    ON CONFLICT (id) DO UPDATE
        SET
            name = EXCLUDED.name,
            traversable = EXCLUDED.traversable,
            updated_at = EXCLUDED.updated_at,
            thumbnail_src = EXCLUDED.thumbnail_src,
            model_sources = EXCLUDED.model_sources,
            compatible_unit_type = EXCLUDED.compatible_unit_type;
