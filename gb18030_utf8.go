package util

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GB2312
// 1980 年，中国发布了第一个汉字编码标准，也即 GB2312 ，全称 《信息交换用汉字编码字符集·基本集》，通常简称 GB （“国标”汉语拼音首字母）， 共收录了 6763 个常用的汉字和字符，此标准于次年5月实施，它满足了日常 99% 汉字的使用需求

// GBK
// 由于有些汉字是在 GB2312 标准发布之后才简化的，还有一些人名、繁体字、日语和朝鲜语中的汉字也没有包括在内，所以，在 GB2312 的基础上添加了这部分字符，就形成了 GBK ，全称 《汉字内码扩展规范》，共收录了两万多个汉字和字符，它完全兼容 GB2312
// GBK 于 1995 年发布，不过它只是 "技术规范指导性文件"，并不属于国家标准

// GB18030
// GB18030 全称《信息技术 中文编码字符集》 ，共收录七万多个汉字和字符， 它在 GBK 的基础上增加了中日韩语中的汉字 和 少数名族的文字及字符，完全兼容 GB2312，基本兼容 GBK
// GB18030 发布过两个版本，第一版于 2000 年发布，称为 GB18030-2000，第二版于 2005 年发布，称为 GB18030-2005

// https://zhuanlan.zhihu.com/p/453675608

func Gb18030ToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
func Gb18030ToUtf8String(s string) string {
	u, err := Gb18030ToUtf8([]byte(s))
	if err != nil {
		return ""
	}
	return string(u)
}

func Utf8ToGb18030(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGb18030String(s string) string {
	gbk, err := Utf8ToGb18030([]byte(s))
	if err != nil {
		return ""
	}
	return string(gbk)
}

// func main() {

// 	s := "GBK 与 UTF-8 编码转换测试"
// 	gbk, err := Utf8ToGbk([]byte(s))
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(string(gbk))
// 	}

// 	utf8, err := GbkToUtf8(gbk)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(string(utf8))
// 	}
// }
