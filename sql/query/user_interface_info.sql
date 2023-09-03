-- name: GetUserInterfaceInfoById :one
SELECT * FROM xapi.`user_interface_info`
WHERE `id` = ? AND isDelete = 0 LIMIT 1;

-- name: GetUserInterfaceInfoByUserIdAndInterfaceId :one
SELECT * FROM xapi.`user_interface_info`
WHERE `userId` = ? AND `interfaceInfoId` = ? AND isDelete = 0 LIMIT 1;

-- name: ListUserInterfaceInfoByUserId :many
SELECT * FROM xapi.`user_interface_info`
WHERE `userId` = ? AND isDelete = 0
ORDER BY id;

-- name: ListUserInterfaceInfoByInterfaceinfoId :many
SELECT * FROM xapi.`user_interface_info`
WHERE `interfaceInfoId` = ? AND isDelete = 0
ORDER BY id;

-- name: ListPageUserInterfaceInfo :many
SELECT * FROM xapi.`user_interface_info`
WHERE isDelete = 0
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: GetUserInterfaceInfoListCount :one
select COUNT(*) FROM xapi.`user_interface_info`
where isDelete = 0;

-- name: CreateUserInterfaceInfo :execresult
insert into xapi.`user_interface_info` (
    `userId`, `interfaceInfoId`
    ) values (
        ?, ?
    );

-- name: InvokeUserInterfaceInfo :execresult
UPDATE xapi.`user_interface_info` set `totalNum`=`totalNum`+1, `leftNum`=`leftNum`-1
WHERE `userId`=? AND `interfaceInfoId`=? AND `isDelete`=0 AND `leftNum` > 0;

-- name: UpdateUserInterfaceInfoLeftNum :exec
UPDATE xapi.`user_interface_info` set `leftNum`=?
WHERE `userId`=? AND `interfaceInfoId`=? AND `isDelete`=0;

-- name: UpdateUserInterfaceInfoStatus :exec
UPDATE xapi.`user_interface_info` set `status`=?
WHERE `userId`=? AND `interfaceInfoId`=? AND `isDelete`=0;

-- name: DeleteUserInterfaceInfo :exec
UPDATE xapi.`user_interface_info` set `isDelete` = 1 
WHERE `userId`=? AND `interfaceInfoId`=?;