package main

//
//import (
//	"bytes"
//	"encoding/json"
//	"net/http"
//)
//
//func SendMail(email, address, productName string) {
//	mailData := map[string]string{
//		"email":        email,
//		"address":      address,
//		"product_name": productName,
//	}
//	marshalMailData, err := json.Marshal(mailData)
//	if err != nil {
//		panic(err)
//	}
//
//	jsonstr := marshalMailData
//
//	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
//	if err != nil {
//		panic(err)
//	}
//
//	client := &http.Client{}
//	response, err := client.Do(request)
//	if err != nil {
//		panic(err)
//	}
//}
