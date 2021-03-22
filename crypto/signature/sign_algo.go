package signature

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// md5 签名:
func WithMD5(data string, privateKey string) string {
	// do sign:
	digest := md5.Sum([]byte(data + privateKey))
	return hex.EncodeToString(digest[:])
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// sha256 签名:
func WithSHA256(data string, privateKey string) string {
	input := []byte(data + privateKey)

	h := sha256.New()
	h.Write(input)

	// do sign:
	digest := h.Sum(nil)
	return hex.EncodeToString(digest[:])
}

// 等价写法:
func withSHA256v2(data string, privateKey string) string {
	input := []byte(data + privateKey)

	// do sign:
	digest := sha256.Sum256(input)
	return hex.EncodeToString(digest[:])
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// sha512 签名:
func WithSHA512(data string, privateKey string) string {
	input := []byte(data + privateKey)

	h := sha512.New()
	h.Write(input)

	// do sign:
	digest := h.Sum(nil)
	return hex.EncodeToString(digest[:])
}

// 等价写法:
func withSHA512v2(data string, privateKey string) string {
	input := []byte(data + privateKey)

	// do sign:
	digest := sha512.Sum512(input)
	return hex.EncodeToString(digest[:])
}
