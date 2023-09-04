-- name: GetUserInfoByUniUserAccount :one
SELECT * FROM `user`
WHERE `userAccount` = ? AND `isDelete` = 0 LIMIT 1;

-- name: GetUserInfoByUniAccessKey :one
SELECT * FROM `user`
WHERE `accessKey` = ? AND `isDelete` = 0 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM `user`
WHERE `isDelete` = 0
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CreateUser2 :execresult
insert into `user` (
    `userName`, `userAccount`, `userAvatar`, `gender`, `userRole`, `userPassword`, `accessKey`, `secretKey`
    ) values (
        ?, ?, ?, ?, ?, ?, ?, ?
    );

-- name: CreateUser :execresult
insert into `user` (
    `userAccount`, `userPassword`, `accessKey`, `secretKey`
    ) values (
        ?, ?, ?, ?
    );

-- name: DeleteUser :exec
UPDATE `user` set `isDelete` = 1 
WHERE id = ?;