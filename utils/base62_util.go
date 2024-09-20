package utils

func Base62Encode(n int) string {
	var base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result string
	for n > 0 {
		result = string(base62[n%62]) + result
		n = n / 62
	}
	return result
}
