-- name: GetInterfaceInfo :one
SELECT * FROM xapi.`interface_info`
WHERE id = ? AND isDelete = 0 LIMIT 1;

-- name: ListInterfaces :many
SELECT * FROM xapi.`interface_info`
WHERE isDelete = 0
ORDER BY id;

-- name: ListPageInterfaces :many
SELECT * FROM xapi.`interface_info`
WHERE isDelete = 0
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CreateInterface :execresult
insert into xapi.`interface_info` (
    `name`, `description`, `url`, `requestParams`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`
    ) values (
        ?, ?, ?, ?, ?, ?, 0, ?, ?
    );

-- name: UpdateInterface :exec
UPDATE xapi.`interface_info` set `name`=?, `description`=?, `url`=?, `requestParams`=?, `requestHeader`=?, `responseHeader`=?, `status`=?, `method`=?, `userId`=?
WHERE id = ?;

-- name: DeleteInterface :exec
UPDATE xapi.`interface_info` set `isDelete` = 1 
WHERE id = ?;

-- name: UpdateInterfaceStatus :exec
UPDATE xapi.`interface_info` set `status`=?
WHERE id = ?;