package userSession

import "crypto/rand"

func genToken(size int) string {
	token := make([]byte, size)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}
	return string(token)
}
