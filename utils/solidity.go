package utils

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func SoliditySha3(input string) string {
	// 使用Keccak-256哈希函数
	hasher := sha3.NewLegacyKeccak256()
	// 添加Solidity函数名
	hasher.Write([]byte(input))
	// 获取哈希值
	hash := hasher.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hexHash := hex.EncodeToString(hash)
	// 只取前4个字节（8个十六进制字符），这是Solidity函数选择器的长度
	selector := hexHash[:8]
	return selector
}
