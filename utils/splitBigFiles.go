package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

const chunkSize int64 = 4 << 20

func main() {
	fileInfo, err := os.Stat("E:/system/CentOS-7-x86_64-Everything-1708.iso")
	if err != nil {
		fmt.Println(err)
	}

	num := int(math.Ceil(float64(fileInfo.Size())/float64(chunkSize)))
	fi, err := os.OpenFile("E:/system/CentOS-7-x86_64-Everything-1708.iso", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte, chunkSize)
	var i int64 = 1
	for ;i <= int64(num);i++ {
		fi.Seek((i-1)*chunkSize,0)

		if len(b) > int((fileInfo.Size() - (i-1)*chunkSize)) {
			b = make([]byte, fileInfo.Size() - (i-1)*chunkSize)
		}

		fi.Read(b)

		f, err := os.OpenFile("./" + strconv.Itoa(int(i)) + ".tmp", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Write(b)
		f.Close()
	}
	fi.Close()

	fii, err := os.OpenFile("E:/system/CentOS-7-x86_64-Everything-1708.iso", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i:=1; i<=num; i++ {
		f, err := os.OpenFile("./" + strconv.Itoa(int(i)) + ".tmp", os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		fii.Write(b)
		f.Close()
	}
}
