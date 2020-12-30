package main

import "fmt"

func main() {
	//data:=[25][]model.WafLog{}
	//tmp:=[]model.WafLog{}
	//err := model.Engine.Where("local_time REGEXP ?",time.Now().Format("2006-01-02*")).Find(&tmp)
	//if err != nil {
	//	panic(err)
	//}
	//for _,v:=range tmp{
	//	data[v.LocalTimeObj.Hour()]=append(data[v.LocalTimeObj.Hour()],v)
	//}
	//for x,y:=range data{
	//	println(x,len(y))
	//}
	//println(time.Now().Format("2006-01-02 15*"))
	//println(time.Now().Hour())
	//
	//
	//
	//tt:=time.Date(2020,12,28,1,23,33,2,time.Local)
	//println(tt.Hour())
	//println(tt.Format("2006-01-02 15*"))
	println(fmt.Sprintf("%.1f",float32(2)/16*100))
}
