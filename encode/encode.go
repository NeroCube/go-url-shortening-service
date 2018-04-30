package encode

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func TinyURL(random_length int) string {

	var charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var b bytes.Buffer

	for i := 0; i <= random_length; i++ {
		rand.Seed(time.Now().UnixNano())
		// the length of charSet is 62
		b.WriteString(fmt.Sprintf("%v", string(charSet[rand.Intn(62)])))
	}

	return b.String()
}
