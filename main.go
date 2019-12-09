package main

import (
	"encoding/xml"
	"fmt"
	"github.com/fiorix/wsdl2go/soap"
	"net/http"
	"wsdl_mine/open_api"
	"wsdl_mine/tranglo"
)

func Tranglo_Ping (trangloClient tranglo.API_Service1Soap){
	res, err := trangloClient.Ping(&tranglo.Ping{})
	fmt.Println(res, err)
}

func Tranglo_GetBalance (trangloClient tranglo.API_Service1Soap){
	var (
		UID = "ptartajasapembayaranelektronis_apigtsb"
		PWD = "b1X57QLs"
	)
	res, _ := trangloClient.Get_Balance(&tranglo.Get_Balance{
		UID: &UID,
		PWD: &PWD,
	})
	a := 7.0
	res.Get_BalanceResult.String.LastBal -= -a

	fmt.Println(res.Get_BalanceResult.String.LastBal)
	//asXML, err := xml.Marshal(&res.Get_BalanceResult)
	//fmt.Println(string(asXML), err)
}

func Tranglo_GetList (trangloClient tranglo.API_Service1Soap){
	var (
		UID = "ptartajasapembayaranelektronis_apigtsb"
		PWD = "b1X57QLs"
		ListName = "BANK"
	)
	res, err := trangloClient.Get_List(&tranglo.Get_List{
		UID:      &UID,
		PWD:      &PWD,
		ListName: &ListName,
	})
	fmt.Println(res.Get_ListResult.String.List, err)
}

func main(){
	clientSoapTranglo := soap.Client{
		URL:       "http://52.220.96.111:2012/API_Service.asmx",
		//URL:       "http://localhost:8088/mock-tranglox",
		Namespace: tranglo.Namespace,
		Pre: func(request *http.Request) {
			request.Header.Set("Content-type", "text/xml")
		},
	}
	trangloClient := tranglo.NewAPI_Service1Soap(&clientSoapTranglo)
	//Tranglo_Ping(trangloClient)
	Tranglo_GetBalance(trangloClient)
	//Tranglo_GetList(trangloClient)
}

type Envelope struct {
	XMLName    xml.Name `xml:"Envelope"`
	Val1       string   `xml:"xmlns:soapenv,attr"`
	Val2       string   `xml:"xmlns:web,attr"`
	CreateBody Body     `xml:"soapenv:Body"`
}

type Body struct {
	CreateText Text `xml:"ok.com web:parseText_XML"`
}

type Text struct {
	TextRow []byte `xml:"rawTextInput"`
}

func OPENAPI (){
	clientSoapOpenAPI := soap.Client{
		URL:       "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso",
		Namespace: open_api.Namespace,
		ContentType: "text/xml",
		Pre: func(request *http.Request) {
		request.Header.Set("Content-type", "text/xml")
		},
	}
	openAPIClient := open_api.NewCountryInfoServiceSoapType(&clientSoapOpenAPI)
	resOpenAPi, _ := openAPIClient.ListOfContinentsByCode(&open_api.ListOfContinentsByCode{})

	for _, row := range resOpenAPi.ListOfContinentsByCodeResult.TContinent {
		fmt.Println(*row.SName)
	}
}