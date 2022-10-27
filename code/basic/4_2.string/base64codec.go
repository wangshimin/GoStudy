// 示例：Base64编码 -- 电子邮件的基础编码格式
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//	需要处理的字符串
	message := "Away from keyboard.https://golang.org/"

	// 编码消息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message)) // 传入的字符串需要转换为字节数组才能供这个函数使用

	// 输出编码完成的消息
	fmt.Println(encodedMessage)

	// 解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMessage)

	// 出错处理
	if err != nil {
		fmt.Println(err)
	} else {
		// 打印解码完成的数据
		fmt.Println(string(data))
	}
}
