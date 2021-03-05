package utillib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

//RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}

func RandomNumber() int32 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return  rnd.Int31n(1000000)
}

func RandomNumber2(num int32) int32 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return  rnd.Int31n(num)
}
func RandomNumber3(num int) int {
	return rand.Intn(num)
}

// 函　数：生成随机数
// 概　要：
// 参　数：
//      min: 最小值
//      max: 最大值
// 返回值：
//      int64: 生成的随机数
func RandInt64(min, max int64) int64 {
	time.Sleep(100)
	rand.Seed(time.Now().Unix())
	randNum := rand.Int63n(max - min)
	randNum = randNum + min
	return randNum
	/*if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min*/
}

func RandNumber(max int)int  {
	time.Sleep(100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}



func Int2String(arg int) string {
	return strconv.Itoa(arg)
}

func Int32toString(arg int32) string {
	return strconv.FormatInt(int64(arg), 10)
}

// 生成32位MD5
func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
//6位随机验证码
func RandomCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.FormatInt(int64(rnd.Int31n(1000000)), 10)
}


var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func interface2String(inter interface{}) {

	switch inter.(type) {

	case string:
		fmt.Println("string", inter.(string))
		break
	case int:
		fmt.Println("int", inter.(int))
		break
	case float64:
		fmt.Println("float64", inter.(float64))
		break
	}

}

func SubString(s string,length int)(res string)  {
	return string([]rune(s)[:length])
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

/*
	判断字符串是否为utf8*/
func ValidUTF8(buf []byte) bool{
	nBytes := 0
	for i:= 0;i<len(buf);i++{
		if nBytes == 0{
			if (buf[i] & 0x80) != 0 {  //与操作之后不为0，说明首位为1
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1 //左移一位
					nBytes++ //记录字符共占几个字节
				}

				if nBytes < 2 || nBytes > 6 { //因为UTF8编码单字符最多不超过6个字节
					return false
				}

				nBytes-- //减掉首字节的一个计数
			}
		}else{ //处理多字节字符
			if buf[i] & 0xc0 != 0x80{ //判断多字节后面的字节是否是10开头
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}

