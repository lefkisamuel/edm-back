package main

import (
	"proto2/bo"
)

func main() {
	objectProxy := bo.NewBusinessObjectProxy(&bo.Product{})
	//err := objectProxy.SetFieldValue("Code", "QWE123456")
	//if (err != nil){
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//_ := objectProxy.SetFieldValue("Price", 12345)

	objectProxy.Save()


}
