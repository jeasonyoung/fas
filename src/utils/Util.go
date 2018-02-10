package utils

import (
	"os"
	"io/ioutil"
	"strings"
	"errors"
)

//判断文件或文件夹是否存在
func PathExists(path string) (bool, error){
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


//加载本地文件
func LocalPathData(fileName string) ([]byte, error) {
	if len(strings.TrimSpace(fileName)) == 0 {
		return nil, errors.New("文件名为空")
	}
	//获取当前运行目录
	dir, _ := os.Getwd()
	//检查路径是否存在
	ok, err := PathExists(fileName)
	if !ok {
		//检查是否已添加运行目录
		if strings.HasPrefix(fileName, dir) {
			return nil, err
		}
		//检查文件名路径是否有/
		if !strings.HasPrefix(fileName, "/") {
			return LocalPathData(dir + "/" + fileName)
		}
		return LocalPathData(dir + fileName)
	}
	//读取文件数据
	return ioutil.ReadFile(fileName)
}