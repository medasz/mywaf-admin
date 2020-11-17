package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"mywaf-admin/setting"
	"path"
)

//采用hash256加密
func Encryption(data string)string{
	tmp:=sha256.Sum256([]byte(data))
	return hex.EncodeToString(tmp[:])
}

//获取目录中文件列表
func FileList(place string)([]string,error){
	filesInfo,err:=ioutil.ReadDir(place)
	if err!=nil{
		return nil,err
	}
	filesList:=[]string{}
	for _,fileInfo:=range filesInfo{
		if !fileInfo.IsDir(){
			filePath:=path.Join(setting.RulePath,fileInfo.Name())
			filesList=append(filesList,filePath)
		}
	}
	return filesList,nil
}

