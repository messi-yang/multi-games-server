INSERT INTO items (id, name, traversable, model_src, thumbnail_src, created_at, updated_at)
VALUES ('3c28537a-80c2-4ac1-917b-b1cd517c6b5e', 'stone', false, '/asset/item/stone/thumbnail.png', '/asset/item/stone/model.gltf', NOW(), NOW())
ON CONFLICT DO NOTHING;

INSERT INTO items (id, name, traversable, model_src, thumbnail_src, created_at, updated_at)
VALUES ('34af14ab-42c5-4c55-a787-44f32012354e', 'torch', true, '/asset/item/torch/thumbnail.png', '/asset/item/torch/model.gltf', NOW(), NOW())
ON CONFLICT DO NOTHING;

INSERT INTO items (id, name, traversable, model_src, thumbnail_src, created_at, updated_at)
VALUES ('414b5703-91d1-42fc-a007-36dd8f25e329', 'tree', true, '/asset/item/tree/thumbnail.png', '/asset/item/tree/model.gltf', NOW(), NOW())
ON CONFLICT DO NOTHING;

