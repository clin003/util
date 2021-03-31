package util

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"encoding/hex"

// 	//"fmt"
// 	"strings"
// )

// //unicode转换为字符串string
// func UscToString(form string) (to string, err error) {
// 	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
// 	if err != nil {
// 		return
// 	}
// 	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
// 		binary.Read(br, binary.BigEndian, &r)
// 		to += string(r)
// 	}
// 	return
// }
