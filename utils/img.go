package utils

import (
	"fmt"
	"io"
	"os"
)

const path = "http://pilipali.cc"

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func SaveImg(url, imgPath, path1 string, str []string) {
	///upload/vod/20190712-1/37028a8a314e23ed79ef7e4c31dd14b4.jpg
	url = path + url
	bol := Exists(path1)

	if !bol {
		err1 := os.Mkdir(path1, os.ModePerm) //创建文件夹
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		resp, err := httpClient.Get(url)
		if err != nil {
			return
		}
		f, err := os.Create(imgPath)
		defer resp.Body.Close()

		buf := make([]byte, 4096)
		for {
			n, err1 := resp.Body.Read(buf)
			if n == 0 {
				break
			}
			if err1 != nil && err1 != io.EOF {
				err = err1
				fmt.Println("下载失败")
				return
			} else {
				fmt.Println("下载成功")
			}

			f.Write(buf[:n])
		}

	} else {
		resp, err := httpClient.Get(url)
		if err != nil {
			return
		}
		f, err := os.Create(imgPath)
		defer resp.Body.Close()

		buf := make([]byte, 4096)
		for {
			n, err1 := resp.Body.Read(buf)
			if n == 0 {
				break
			}
			if err1 != nil && err1 != io.EOF {
				err = err1
				fmt.Println("下载失败")
				return
			} else {
				fmt.Println("下载成功")
			}

			f.Write(buf[:n])
		}
	}
}
