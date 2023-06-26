package main

import "fmt"

type Fsubscriber struct {
	Subscriberid  int
	Imsi          int `json:"Imsi,string"`
	Plmn          string
	Skey          string
	Seq           int64
	Opc           string
	Amf           string
	Accountstatus int `json:"Accountstatus,string"`
	Defsst        int `json:"Defsst,string"`
	Defsd         string
	Uplink        int `json:"Uplink,string"`
	Downlink      int `json:"Downlink,string"`
	Dnnid1        string
	Dnn1Name      string
	Dnnid2        string
	Dnn2Name      string
	Dnnid3        string
	Dnn3Name      string
	Msisdn        string
	Createddate   string
	Createdby     int
	Updateddate   string
	Updatedby     int
	Isactive      int
	Rowversion    int
}

type Fsubscriber_list struct {
	TotalRecords int
	List         []Fsubscriber
}

func main() {
	var drag Fsubscriber_list
	fmt.Println(drag.List)

}
