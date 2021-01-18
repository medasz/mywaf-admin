package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
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
			filePath:=path.Join(place,fileInfo.Name())//setting.RulePath,fileInfo.Name())
			filesList=append(filesList,filePath)
		}
	}
	return filesList,nil
}

//添加文件内容
func OverFile(content []byte,filepath string) error {
	file,err:=os.Create(filepath)
	if err!=nil{
		return err
	}
	_,err=file.Write(content)
	if err!=nil{
		return err
	}
	file.Close()
	return nil
}

//重载nginx
//func ReloadRule() error {
//	output,err:=exec.Command(setting.NginxBin,"-t").Output()
//	if err!=nil{
//		log.Println(output,err)
//		return err
//	}
//	output,err=exec.Command(setting.NginxBin,"-s","reload").Output()
//	if err!=nil{
//		log.Println(output,err)
//		return err
//	}
//	return nil
//}