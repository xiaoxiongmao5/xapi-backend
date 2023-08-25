package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"
	"xj/xapi-backend/config"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/** 生成登录验证token（jwt）
 */
func GenerateToken(userID, userRole string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"user_role": userRole,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey))
}

/** 加密密码（bcrypt 哈希算法）
 */
func HashPasswordByBcrypt(password string) (string, error) {
	// bcrypt.DefaultCost: 这是 bcrypt 哈希算法的工作因子（cost factor），表示计算哈希时使用的迭代次数。工作因子越高，计算哈希所需的时间和资源就越多，因此更难受到暴力破解。bcrypt.DefaultCost 是库中预定义的默认工作因子值。
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

/** 校验加密密码
 */
func CheckHashPasswordByBcrypt(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

/** 生成随机字符串
 */
func GenerateRandomKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	// 生成随机的字节序列
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// 转为base64编码的字符串
	return base64.StdEncoding.EncodeToString(randomBytes), nil
}

/** 生成带盐的哈希值（SHA-256 哈希算法）
 */
func HashBySHA256WithSalt(data, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data + salt))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

/** 签名工具
 */
func GetSign(body string, accessKey string) (sign string) {
	// 生成随机数 nonce

	// 请求体内容

	// 当前时间戳 timestamp

	// 生成签名
	return
}

func CheckSign(sign string, accessKey string) (res bool) {
	// 校验随机数

	// 校验时间戳与当前时间的差距，不能超过5分钟
	return
}
