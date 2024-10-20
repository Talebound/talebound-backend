-- name: CreateMap :one
INSERT INTO maps (user_id, title, type, description, width, height, thumbnail_image_id, is_private)
VALUES (sqlc.arg(user_id), sqlc.arg(title), sqlc.narg(type), sqlc.narg(description), sqlc.arg(width), sqlc.arg(height), sqlc.narg(thumbnail_image_id), sqlc.arg(is_private))
RETURNING *;

-- name: GetMaps :many
WITH cte AS (
    SELECT
        *
    FROM get_maps( sqlc.narg(tags)::int[], sqlc.narg(module_id), sqlc.narg(module_type), sqlc.narg(order_by), sqlc.narg(order_direction), 0, 0)
)
SELECT
    CAST((SELECT count(*) FROM cte) as integer) as total_count,
    cte.*
FROM cte
ORDER BY id DESC
LIMIT sqlc.arg(page_limit)
    OFFSET sqlc.arg(page_offset);


-- name: GetMapsByIDs :many
SELECT * FROM maps WHERE id = ANY(@map_ids::int[]);

-- name: GetMapById :one
SELECT * FROM maps WHERE id = sqlc.arg(id);

-- name: UpdateMap :one
UPDATE maps
SET
    title = COALESCE(sqlc.narg(title), title),
    type = COALESCE(sqlc.narg(type), type),
    description = COALESCE(sqlc.narg(description), description),
    width = COALESCE(sqlc.narg(width), width),
    height = COALESCE(sqlc.narg(height), height),
    thumbnail_image_id = COALESCE(sqlc.narg(thumbnail_image_id), thumbnail_image_id),
    is_private = COALESCE(sqlc.narg(is_private), is_private),
    last_updated_at = now(),
    last_updated_user_id =  COALESCE(sqlc.narg(last_updated_user_id), last_updated_user_id)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteMap :exec
CALL delete_map(sqlc.arg(id));

-- name: CreateMapLayer :one
INSERT INTO map_layers (name, map_id, image_id, enabled, position)
VALUES (sqlc.arg(name), sqlc.arg(map_id), sqlc.arg(image_id), sqlc.arg(enabled), sqlc.arg(position))
RETURNING *;

-- name: GetMaxMapLayerPosition :one
SELECT CAST(MAX(position) as integer) FROM map_layers WHERE map_id = sqlc.arg(map_id);

-- name: IncreaseMapLayerPositions :exec
UPDATE map_layers SET position = position + 1
WHERE map_id = sqlc.arg(map_id) AND position >= sqlc.arg(position);

-- name: DecreaseMapLayerPositions :exec
UPDATE map_layers SET position = position - 1
WHERE map_id = sqlc.arg(map_id) AND position >= sqlc.arg(position);

-- name: GetMapLayers :many
SELECT * FROM view_map_layers WHERE map_id = sqlc.arg(map_id);

-- name: GetMapLayerByID :one
SELECT * FROM view_map_layers WHERE id = sqlc.arg(map_layer_id);

-- name: UpdateMapLayer :one
UPDATE map_layers
SET
    name = COALESCE(sqlc.narg(name), name),
    image_id = COALESCE(sqlc.narg(image_id), image_id),
    enabled = COALESCE(sqlc.narg(enabled), enabled),
    position = COALESCE(sqlc.narg(position), position)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: MoveMapLayer :exec
CALL move_map_layer(sqlc.arg(id), sqlc.arg(position));

-- name: DeleteMapLayer :exec
DELETE FROM map_layers WHERE id = sqlc.arg(id);

-- name: DeleteMapLayersForMap :exec
DELETE FROM map_layers WHERE map_id = sqlc.arg(map_id);

--------------------------------------
-- name: CreateMapPinTypeGroup :one
INSERT INTO map_pin_type_group (name)
VALUES (sqlc.arg(name))
RETURNING *;

-- name: UpdateMapPinTypeGroup :one
UPDATE map_pin_type_group
SET
    name = COALESCE(sqlc.narg(name), name)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteMapPinTypeGroup :exec
DELETE FROM map_pin_type_group WHERE id = sqlc.arg(id);

-- name: CreateModuleMapPinTypeGroup :one
INSERT INTO module_map_pin_type_groups (module_id, map_pin_type_group_id)
VALUES (sqlc.arg(module_id), sqlc.arg(map_pin_type_group_id))
RETURNING *;

-- name: DeleteModuleMapPinTypeGroup :exec
DELETE FROM module_map_pin_type_groups WHERE module_id = sqlc.arg(module_id) AND map_pin_type_group_id = sqlc.arg(map_pin_type_group_id);

--------------------------------------

-- name: CreateMapPinType :one
INSERT INTO map_pin_types (shape, background_color, border_color, icon_color, icon, icon_size, width, map_pin_type_group_id, is_default)
VALUES (sqlc.arg(shape), sqlc.arg(background_color), sqlc.arg(border_color), sqlc.arg(icon_color), sqlc.arg(icon), sqlc.arg(icon_size), sqlc.arg(width), sqlc.arg(map_pin_type_group_id), sqlc.narg(is_default) )
RETURNING *;

-- name: GetMapPinTypesForModule :many
SELECT
    mpt.*
FROM
    map_pin_types mpt
    JOIN module_map_pin_type_groups mmptg ON mpt.map_pin_type_group_id = mmptg.map_pin_type_group_id
WHERE mmptg.module_id = sqlc.arg(module_id);

-- name: GetMapPinTypesForMap :many
SELECT
    mpt.*
FROM
    map_pin_types mpt
    JOIN module_map_pin_type_groups mmptg ON mpt.map_pin_type_group_id = mmptg.map_pin_type_group_id
    JOIN entities e ON e.module_id = mmptg.module_id
WHERE e.map_id = sqlc.arg(map_id);

-- name: GetDefaultMapPinTypeForModule :one
SELECT
    mpt.id
FROM
    map_pin_types mpt
    JOIN module_map_pin_type_groups mmptg ON mpt.map_pin_type_group_id = mmptg.map_pin_type_group_id
WHERE mmptg.module_id = sqlc.arg(module_id) AND mpt.is_default = true;

-- name: GetDefaultMapPinTypeForMap :one
SELECT
    mpt.id
FROM
    map_pin_types mpt
    JOIN module_map_pin_type_groups mmptg ON mpt.map_pin_type_group_id = mmptg.map_pin_type_group_id
    JOIN entities e ON e.module_id = mmptg.module_id
WHERE e.map_id = sqlc.arg(map_id) AND mpt.is_default = true;

-- name: GetMapPinTypeGroupsForModule :many
SELECT
    mptg.*
FROM
    map_pin_type_group mptg
    JOIN module_map_pin_type_groups mmptg ON mptg.id = mmptg.map_pin_type_group_id
WHERE mmptg.module_id = sqlc.arg(module_id);

-- name: UpdateMapPinType :one
UPDATE map_pin_types
SET
    shape = COALESCE(sqlc.narg(shape), shape),
    background_color = COALESCE(sqlc.narg(background_color), background_color),
    border_color = COALESCE(sqlc.narg(border_color), border_color),
    icon_color = COALESCE(sqlc.narg(icon_color), icon_color),
    icon = COALESCE(sqlc.narg(icon), icon),
    icon_size = COALESCE(sqlc.narg(icon_size), icon_size),
    width = COALESCE(sqlc.narg(width), width),
    is_default = COALESCE(sqlc.narg(is_default), is_default)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteMapPinType :exec
DELETE FROM map_pin_types WHERE id = sqlc.arg(id);

-- name: CreateMapPin :one
INSERT INTO map_pins (name, map_id, map_pin_type_id, location_id, map_layer_id, x, y)
VALUES (sqlc.arg(name), sqlc.arg(map_id), sqlc.narg(map_pin_type_id), sqlc.narg(location_id), sqlc.narg(map_layer_id), sqlc.arg(x), sqlc.arg(y))
RETURNING *;

-- name: GetMapPins :many
SELECT * FROM view_map_pins WHERE map_id = sqlc.arg(map_id);

-- name: GetMapPinByID :one
SELECT * FROM view_map_pins WHERE id = sqlc.arg(id);

-- name: UpdateMapPin :one
UPDATE map_pins
SET
    name = COALESCE(sqlc.narg(name), name),
    map_pin_type_id = COALESCE(sqlc.narg(map_pin_type_id), map_pin_type_id),
    location_id = COALESCE(sqlc.narg(location_id), location_id),
    map_layer_id = COALESCE(sqlc.narg(map_layer_id), map_layer_id),
    x = COALESCE(sqlc.narg(x), x),
    y = COALESCE(sqlc.narg(y), y)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteMapPin :exec
DELETE FROM map_pins WHERE id = sqlc.arg(id);

-- name: DeleteMapPinsForMapLayer :exec
DELETE FROM map_pins WHERE map_layer_id = sqlc.arg(map_layer_id);

-- name: DeleteMapPinsForMap :exec
DELETE FROM map_pins WHERE map_id = sqlc.arg(map_id);

-- name: GetMapAssignments :one
SELECT
    m.*
FROM
    entities e
    LEFT JOIN modules m ON e.module_id = m.id
WHERE e.map_id = sqlc.arg(map_id)
;
