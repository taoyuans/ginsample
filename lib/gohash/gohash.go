package gohash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash256(str string) string {
	bytes := sha256.Sum256([]byte(str)) //计算哈希值，返回一个长度为32的数组
	return hex.EncodeToString(bytes[:]) //将数组转换成切片，转换成16进制，返回字符串
}
