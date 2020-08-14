package transport

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

type BaseData struct {
	Timestamp    int64  `json:"timestamp"`
	TransportKey string `json:"transportKey"`
	VerifyHash   string `json:"verifyHash"`
}

func (d *BaseData) IsVerified(productManagerKey string) bool {
	return d.VerifyHash == fmt.Sprintf("%x", md5.Sum([]byte(d.TransportKey+productManagerKey)))
}

func (d *BaseData) SetVerificationData(productManagerKey, transportKey string) {
	if transportKey == "" && d.TransportKey != "" {
		transportKey = d.TransportKey
	} else if transportKey == "" {
		transportKey = randomString(40)
		d.TransportKey = transportKey
	} else {
		d.TransportKey = transportKey
	}
	d.VerifyHash = fmt.Sprintf("%x", md5.Sum([]byte(transportKey+productManagerKey)))
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	llen := len(letter)
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(llen)]
	}
	return string(b)
}
