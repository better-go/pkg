package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

/*
bcrypt 算法:
	- 单向 hash 加密
	- Bcrypt 是一个用于密码哈希的加密算法，它基于 Blowfish 加密算法
	- 安全：每个密码的盐值都是随机的，并且计算过程经过多次迭代
	- 使用简单：Bcrypt 把算法版本、计算次数和 salt（盐值）都放到 hash 值里面去了，所以不用再单独维护盐值了
	- https://zh.wikipedia.org/wiki/Bcrypt

*/

// 加密
func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 验证:
func ValidatePassword(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
