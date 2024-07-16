package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

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

// StringToBytes converts a string to a byte slice
func StringToBytes(str string) []byte {
	return []byte(str)
}

// BytesToHex converts a byte slice to a hexadecimal string
func BytesToHex(data []byte) string {
	return hex.EncodeToString(data)
}

// HexToBytes converts a hexadecimal string to a byte slice
func HexToBytes(hexStr string) ([]byte, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, errors.New("invalid hexadecimal string")
	}
	return bytes, nil
}

// Keccak256Hash computes the Keccak-256 hash of a given byte slice
func Keccak256Hash(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

// GenerateRandomBytes generates a random byte slice of the specified length
func GenerateRandomBytes(length int) ([]byte, error) {
	if length <= 0 {
		return nil, errors.New("length must be greater than 0")
	}
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
