-- name: GetUserInfo :one
SELECT * FROM `user`
WHERE `userAccount` = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM `user`
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CreateUser :execresult
insert into `user` (
    `userName`, `userAccount`, `userAvatar`, `gender`, `userRole`, `userPassword`, `accessKey`, `secretKey`
    ) values (
        ?, ?, ?, ?, ?, ?, ?, ?
    );

-- name: DeleteUser :exec
UPDATE `user` set `isDelete` = 1 
WHERE id = ?;