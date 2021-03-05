package utillib

import (
	"fmt"
	"os"
)

func WriteContent(info string, filePath string) {

	//tempStr := IPAddress+"|"+Port+"|"+Account+"|"+Password
	//logger.Info("\n[======WriteContent "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"======")
	//var d1 = []byte(info)
	//err := ioutil.WriteFile("./info.txt", d1, 0666) //写入文件(字节数组)

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("os OpenFile error: ", err)
		return
	}
	defer f.Close()

	_, _ = f.WriteString(info)

}
