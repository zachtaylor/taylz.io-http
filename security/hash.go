package security

import "crypto/md5"

// Hash is a func(string) string
type Hash = func(string) string

func NewMD5SaltHash(salt string) Hash {
	return func(password string) string {
		sum := md5.Sum([]byte(password + salt))
		return string(sum[:])
	}
}
