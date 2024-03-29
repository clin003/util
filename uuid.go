package util

import (
	"crypto/rand"
	"fmt"
	// "github.com/google/uuid"
)

// 生成随机UUID(RFC 4122)
// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func CreateUUID() (string, error) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		// log.Errorf(err, "Cannot generate UUID:%s")
		return "", err
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid, err
}

// func CreateUUIDStr() string {
// 	return uuid.New().String()
// }
// func GenShortId() (string, error) {
// 	// return shortid.Generate()
// 	u, err := uuid.NewRandom()
// 	if err != nil {
// 		return "", err
// 	}
// 	return u.String(), err
// }
