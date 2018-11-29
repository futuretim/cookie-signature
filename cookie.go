package cookie_signature

/**

	Golang HMAC w/64-bit encoding: https://www.jokecamp.com/blog/examples-of-creating-base64-hashes-using-hmac-sha256-in-different-languages/#go

 */

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func Signcookie(message string, secret string, replacer *strings.Replacer) (string) {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))

	encodedString := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	replaced := replacer.Replace(encodedString)

	return replaced
}

func Unsigncookie(cookie string, secret string, replacer *strings.Replacer) (bool) {
	valid := false

	idx := strings.Index(cookie, ".")

	testPart := cookie[:idx]

	signed := Signcookie(testPart, secret, replacer)

	signedPart := cookie[idx+1:]
	valid = signed == signedPart

	return valid
}
