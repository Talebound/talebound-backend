-- name: CreateWorld :one
INSERT INTO worlds (
    name,
    description
) VALUES (
     @name, @description
 ) RETURNING *;

-- name: UpdateWorld :one
UPDATE worlds
SET
    name = COALESCE(sqlc.narg(name), name),
    public = COALESCE(sqlc.narg(public), public),
    description = COALESCE(sqlc.narg(description), description)
WHERE
    id = sqlc.arg(world_id)
RETURNING *;

-- name: DeleteWorld :exec
DELETE FROM worlds WHERE id = @world_id;

-- name: GetWorldByID :one
SELECT * FROM view_worlds WHERE id = @world_id LIMIT 1;

-- name: GetWorlds :many
SELECT * FROM view_worlds
WHERE (@is_public::boolean IS NULL OR public = @is_public)
ORDER BY
    CASE
     WHEN @order_result::bool
         THEN @order_by::VARCHAR
     ELSE 'activity'
     END
DESC
LIMIT @page_limit
OFFSET @page_offset;

-- name: CreateWorldAdmin :one
INSERT INTO world_admins (
    world_id,
    user_id,
    is_main
) VALUES (@world_id, @user_id, @is_main) RETURNING *;


-- name: GetWorldsOfUser :many
SELECT
    vw.*
FROM
    view_worlds vw
    JOIN world_admins wa ON wa.world_id = vw.id
WHERE
    wa.user_id = @user_id
;

-- name: GetAdminsOfWorld :many
SELECT
    vu.*,
    wa.is_main as is_main
FROM
    view_users vu
    JOIN world_admins wa on wa.user_id = vu.id
WHERE
    wa.world_id = @world_id
;

-- name: IsWorldAdmin :one
SELECT * FROM world_admins WHERE user_id = @user_id AND world_id = @world_id;