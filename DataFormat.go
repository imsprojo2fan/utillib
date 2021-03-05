package utillib

import (
	"strconv"
	"strings"
)

import (
	"fmt"
	"os"
	"time"
)

/*字节数组转16进制可以直接使用 用fmt就能转换*/
func ByteToHex(b []byte) (H string) {
	H = fmt.Sprintf("%x", b)
	return
}

/*16进制转字节数组*/
func HexToByte(str string) []byte {
	Glen := len(str)
	bHex := make([]byte, len(str)/2)
	ii := 0
	for i := 0; i < len(str); i = i + 2 {
		if Glen != 1 {
			ss := string(str[i]) + string(str[i+1])
			bt, _ := strconv.ParseInt(ss, 16, 32)
			bHex[ii] = byte(bt)
			ii = ii + 1
			Glen = Glen - 2
		}
	}
	return bHex
}

/*16进制转10进制*/
func HexToI(str string) int64 {
	//v := "000A"
	var r int64
	if s, err := strconv.ParseInt(str, 16, 32); err == nil {
		//fmt.Println(str)
		//fmt.Printf("%T|%d|\n", str, s)
		r = s
		//fmt.Println(reflect.TypeOf(s))
	}

	return r
}

// 10进制转任意进制
var tenToAny map[int64]string = map[int64]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

func DecimalToAny(num, n int64) string {
	newNumStr := ""
	var remainder int64
	var remainderString string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			remainderString = tenToAny[remainder]
		} else {
			remainderString = strconv.FormatInt(remainder, 10)
		}
		newNumStr = remainderString + newNumStr
		num = num / n
	}
	return newNumStr
}

//二进制转十六进制
func ByteToHex2(b string) string {
	base, _ := strconv.ParseInt(b, 2, 10)
	return strconv.FormatInt(base, 16)
}

/*字符串切割*/
func Substr(str string, start int, end int) string {

	//fmt.Println("start-->",start)
	//fmt.Println("end-->",end)
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		fmt.Println("start is wrong")
	}

	if end < 0 || end > length {
		fmt.Println("end is wrong")
	}

	return string(rs[start:end])
}

/*时间转时间戳*/
func T2TimeStamp(toBeCharge string) int64 {
	//fmt.Println(toBeCharge)
	//获取本地location
	//toBeCharge := "2015-01-01 00:00:00"                           //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)
	return sr
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

func MySplit(str string) []byte {
	var strArr []string
	length := len(str)
	var bArr []byte
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			temp := string([]rune(str)[i : i+2])
			strArr = append(strArr, temp)
		}

	}
	for _, item := range strArr {
		tempArr := HexToByte(item)
		for _, item02 := range tempArr {
			bArr = append(bArr, item02)
		}
	}
	return bArr
}

var MbTable = []uint16{
	0x0000, 0xC0C1, 0xC181, 0x0140, 0xC301, 0x03C0, 0x0280, 0xC241,
	0xC601, 0x06C0, 0x0780, 0xC741, 0x0500, 0xC5C1, 0xC481, 0x0440,
	0xCC01, 0x0CC0, 0x0D80, 0xCD41, 0x0F00, 0xCFC1, 0xCE81, 0x0E40,
	0x0A00, 0xCAC1, 0xCB81, 0x0B40, 0xC901, 0x09C0, 0x0880, 0xC841,
	0xD801, 0x18C0, 0x1980, 0xD941, 0x1B00, 0xDBC1, 0xDA81, 0x1A40,
	0x1E00, 0xDEC1, 0xDF81, 0x1F40, 0xDD01, 0x1DC0, 0x1C80, 0xDC41,
	0x1400, 0xD4C1, 0xD581, 0x1540, 0xD701, 0x17C0, 0x1680, 0xD641,
	0xD201, 0x12C0, 0x1380, 0xD341, 0x1100, 0xD1C1, 0xD081, 0x1040,
	0xF001, 0x30C0, 0x3180, 0xF141, 0x3300, 0xF3C1, 0xF281, 0x3240,
	0x3600, 0xF6C1, 0xF781, 0x3740, 0xF501, 0x35C0, 0x3480, 0xF441,
	0x3C00, 0xFCC1, 0xFD81, 0x3D40, 0xFF01, 0x3FC0, 0x3E80, 0xFE41,
	0xFA01, 0x3AC0, 0x3B80, 0xFB41, 0x3900, 0xF9C1, 0xF881, 0x3840,
	0x2800, 0xE8C1, 0xE981, 0x2940, 0xEB01, 0x2BC0, 0x2A80, 0xEA41,
	0xEE01, 0x2EC0, 0x2F80, 0xEF41, 0x2D00, 0xEDC1, 0xEC81, 0x2C40,
	0xE401, 0x24C0, 0x2580, 0xE541, 0x2700, 0xE7C1, 0xE681, 0x2640,
	0x2200, 0xE2C1, 0xE381, 0x2340, 0xE101, 0x21C0, 0x2080, 0xE041,
	0xA001, 0x60C0, 0x6180, 0xA141, 0x6300, 0xA3C1, 0xA281, 0x6240,
	0x6600, 0xA6C1, 0xA781, 0x6740, 0xA501, 0x65C0, 0x6480, 0xA441,
	0x6C00, 0xACC1, 0xAD81, 0x6D40, 0xAF01, 0x6FC0, 0x6E80, 0xAE41,
	0xAA01, 0x6AC0, 0x6B80, 0xAB41, 0x6900, 0xA9C1, 0xA881, 0x6840,
	0x7800, 0xB8C1, 0xB981, 0x7940, 0xBB01, 0x7BC0, 0x7A80, 0xBA41,
	0xBE01, 0x7EC0, 0x7F80, 0xBF41, 0x7D00, 0xBDC1, 0xBC81, 0x7C40,
	0xB401, 0x74C0, 0x7580, 0xB541, 0x7700, 0xB7C1, 0xB681, 0x7640,
	0x7200, 0xB2C1, 0xB381, 0x7340, 0xB101, 0x71C0, 0x7080, 0xB041,
	0x5000, 0x90C1, 0x9181, 0x5140, 0x9301, 0x53C0, 0x5280, 0x9241,
	0x9601, 0x56C0, 0x5780, 0x9741, 0x5500, 0x95C1, 0x9481, 0x5440,
	0x9C01, 0x5CC0, 0x5D80, 0x9D41, 0x5F00, 0x9FC1, 0x9E81, 0x5E40,
	0x5A00, 0x9AC1, 0x9B81, 0x5B40, 0x9901, 0x59C0, 0x5880, 0x9841,
	0x8801, 0x48C0, 0x4980, 0x8941, 0x4B00, 0x8BC1, 0x8A81, 0x4A40,
	0x4E00, 0x8EC1, 0x8F81, 0x4F40, 0x8D01, 0x4DC0, 0x4C80, 0x8C41,
	0x4400, 0x84C1, 0x8581, 0x4540, 0x8701, 0x47C0, 0x4680, 0x8641,
	0x8201, 0x42C0, 0x4380, 0x8341, 0x4100, 0x81C1, 0x8081, 0x4040}

func CheckSum(data []byte) uint16 {
	var crc16 uint16
	crc16 = 0xffff
	for _, v := range data {
		n := uint8(uint16(v) ^ crc16)
		crc16 >>= 8
		crc16 ^= MbTable[n]
	}
	return crc16
}

//通用modbus CRC校验算法
func ModBusCRC(dataString string) string {
	crc := 0xFFFF
	length := len(dataString)
	for i := 0; i < length; i++ {
		//通用modbus取寄存器的低8位参与异或运算
		crc = ((crc << 8) >> 8) ^ int(dataString[i])
		for j := 0; j < 8; j++ {
			flag := crc & 0x0001
			crc >>= 1
			if flag == 1 {
				crc ^= 0xA001
			}
		}
	}

	//得到的十六进制校验码是按照高字节在前低字节在后的字符串
	//要翻转，按照低字节在前高字节在后
	//校验码必须是4个字符，不足4位的需要在开头补0
	hex := strconv.FormatInt(int64(crc), 16) //格式化为16进制字符串
	tmp := hex[2:] + hex[:2]
	if len(tmp) == 3 {
		tmp = "0" + tmp
	}
	return tmp
}

//HJ212 CRC校验算法
func Hjt212CRC(dataString string) string {
	crc := 0xFFFF
	length := len(dataString)
	for i := 0; i < length; i++ {
		//hj212取寄存器的高8位参与异或运算
		crc = (crc >> 8) ^ int(dataString[i])
		for j := 0; j < 8; j++ {
			flag := crc & 0x0001
			crc >>= 1
			if flag == 1 {
				crc ^= 0xA001
			}
		}
	}
	//因为是基于右移位运算的结果，得到的本身就是低字节在前高字节在后的结果
	//不足4位的需要在开头补0
	hex := strconv.FormatInt(int64(crc), 16)
	if len(hex) == 3 {
		hex = "0" + hex
	}
	return hex
}

func Xor(data string) string {
	var tempArr []string
	for i := 0; i < len(data)/2; i++ {
		tempStr := data[i*2 : (i+1)*2]
		tempArr = append(tempArr, tempStr)
	}
	resArr := make([]int64, len(tempArr))
	for i := 0; i < len(tempArr); i++ {
		dec, err := strconv.ParseInt("0x"+tempArr[i], 0, 64)
		if err != nil {
			panic(err)
		}
		if i == 0 {
			resArr[i] = dec ^ 0
		} else {
			resArr[i] = dec ^ resArr[i-1]
		}
	}
	/*for i := 0; i < len(resArr); i ++ {
		fmt.Printf("%x\n", resArr[i])
	}*/
	resInt64 := resArr[len(tempArr)-1]
	return strings.ToUpper(strconv.FormatInt(resInt64, 16))
}
