package util

import "strings"

var (
	base62 = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Format10To62(n uint64) string {
	var str strings.Builder
	var r uint64
	for n != 0 {
		r = n % uint64(len(base62))
		str.WriteByte(base62[r])
		n /= uint64(len(base62))
	}
	return str.String()
}
