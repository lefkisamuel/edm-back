package main

import (
	"edm-back/bo"
)

func main() {
	objectProxy := bo.NewBusinessObjectProxy(
		&bo.Product{
			ProductClient: bo.ProductClient{
				FirstName: "Samuel",
				LastName: "Lefki",
			},
		})
	_ = objectProxy.SetFieldValue("Code", "AZERTYU")
	_ = objectProxy.SetFieldValue("Price", 12345)
	//if (err != nil){
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	objectProxy.Save()



}
