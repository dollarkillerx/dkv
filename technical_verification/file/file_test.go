/**
 * @Author: DollarKillerX
 * @Description: file_test 存储的可行性验证
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:46 2019/12/18
 */
package file

import (
	"log"
	"os"
	"testing"
)

func TestStory(t *testing.T) {
	file, e := os.OpenFile("test1.dkv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 00666)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	//byt := make([]byte, 65535)
	_, e = file.Seek(65535, 0)
	if e != nil {
		panic(e)
	}

	//buf := bytes.NewBuffer(byt)
	//writer := bufio.NewWriter(buf)
	//writer.WriteString("adasdasdasdsadasd")
	//writer.Flush()

	//datas := "撒旦的点点滴滴的点点滴滴的点点滴滴\n"  // 52
	//datas := "\n"  // 1
	//datas := "\n"  // 1
	datas := "撒旦的点点滴滴的点点滴滴的点点滴滴\n" // 52

	n, e := file.Write([]byte(datas))
	if e != nil {
		panic(e)
	} else {
		log.Println(n)
	}
	// 上面插入的 数: 106
}

func TestReadDDE(t *testing.T) {
	// 读取指定位置
	// 读取第四个  52 + 1 + 1  （54～106）
	file, e := os.OpenFile("test1.dkv", os.O_RDONLY, 00666)
	if e != nil {
		panic(e)
	}

	ret, e := file.Seek(54, 0)
	if e != nil {
		panic(e)
	}
	log.Println(ret)
	data := make([]byte,52)
	_, e = file.Read(data)
	if e != nil {
		panic(e)
	}
	log.Println(string(data))
}
