-- name: GetInterfaceInfoById :one
SELECT * FROM xapi.`interface_info`
WHERE `id` = ? AND `isDelete` = 0 LIMIT 1;

-- name: GetInterfaceInfoByUniFullApi :one
SELECT * FROM xapi.interface_info
WHERE `host` = ? AND `url` = ? AND `method` = ? AND `isDelete` = 0 LIMIT 1;

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

-- name: GetInterfaceListCount :one
select COUNT(*) FROM xapi.`interface_info`
where isDelete = 0;

-- name: CreateInterface :execresult
insert into xapi.`interface_info` (
    `name`, `description`, `host`, `url`, `requestParams`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`
    ) values (
        ?, ?, ?, ?, ?, ?, ?, 0, ?, ?
    );

-- name: UpdateInterface :exec
UPDATE xapi.`interface_info` set `name`=?, `description`=?, `host`=?, `url`=?, `requestParams`=?, `requestHeader`=?, `responseHeader`=?, `method`=?, `userId`=?
WHERE id = ?;

-- name: DeleteInterface :exec
UPDATE xapi.`interface_info` set `isDelete` = 1 
WHERE id = ?;

-- name: UpdateInterfaceStatus :exec
UPDATE xapi.`interface_info` set `status`=?
WHERE id = ?;