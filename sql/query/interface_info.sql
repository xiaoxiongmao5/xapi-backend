-- name: GetInterfaceInfo :one
SELECT * FROM xapi.`interface_info`
WHERE id = ? LIMIT 1;

-- name: ListInterfaces :many
SELECT * FROM xapi.`interface_info`
ORDER BY id;

-- name: CreateInterface :execresult
insert into xapi.`interface_info` (
    `name`, `description`, `url`, `requestParams`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`
    ) values (
        ?, ?, ?, ?, ?, ?, 0, ?, ?
    );

-- name: DeleteInterface :exec
UPDATE xapi.`interface_info` set `isDelete` = 1 
WHERE id = ?;