package utillib

import (
	"strconv"
	"strings"
)

/*
	1. N（北纬） 2238.5260
		1.1 2238.5260÷100=22.385260（取整） =22
		1.2 385260÷60=6421
		得到以度形式的纬度坐标为 N 22.642100°
	2. E（东经） 11401.9686
		2.1 11401.9686÷100=114.019686（取整） =114
		2.2 019686÷60=0328.1
		得到以度形式的经度坐标为 E 114.032810°
*/
func TransformLngLat(tempData string)string  {
	//lng := "11809.496623"
	//lat := "2431.890961"
	index := strings.LastIndex(tempData,".")-2
	temp1 := tempData[0:index]
	tempData = strings.Replace(tempData,".","",-1)
	temp2 := tempData[index:]
	temp2Num,_ := strconv.Atoi(temp2)
	temp2Num = temp2Num/60
	res := temp1+"."+strconv.Itoa(temp2Num)
	return res
}