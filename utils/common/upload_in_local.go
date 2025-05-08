package common

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
)

// 将文件保存至本地
func UploadFileInLocal(data []byte, path string) (err error) {

	//分割为目录和文件明
	basePath, _ := filepath.Split(path)
	//判断是否有这个文件目录
	_, err = os.ReadDir(basePath)
	if err != nil {
		//创建该文件目录
		err = os.MkdirAll(basePath, fs.ModePerm) //Mkdir 方法不能直接创建多级目录
		if err != nil {
			return err
		}
	}

	//文件下载至本地
	err = bufioWrite(data, path)
	if err != nil {
		return err
	}

	//	下载成功
	return nil
}

// bufio缓冲写入
func bufioWrite(data []byte, destName string) error {

	destFile, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 创建一个带缓冲的写入器
	writer := bufio.NewWriter(destFile)

	// 将数据写入缓冲区
	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	// 刷新缓冲区，将数据写入文件
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
