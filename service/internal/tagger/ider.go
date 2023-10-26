package tagger

import (
	"crypto/md5"
	"fmt"
)

func IdHash(value string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(value)))

}
