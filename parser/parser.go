package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Item struct {
	Lsid       LsidType               `json:"lsid"`
	Ns         string                 `json:"ns"`
	O          map[string]interface{} `json:"o"`
	O2         map[string]interface{} `json:"o2"`
	Op         string                 `json:"op"`
	PrevOpTime TimeType               `json:"prevOpTime"`
	StmtId     int                    `json:"stmtId"`
	T          int                    `json:"t"`
	Ts         Timestamp              `json:"ts"`
	TxnNumber  int                    `json:"txnNumber"`
	Ui         UiType                 `json:"ui"`
	V          int                    `json:"v"`
	Wall       string                 `json:"wall"`
}

type LsidType struct {
	Id  IdType  `json:"id"`
	Uid UidType `json:"uid"`
}

type IdType struct {
	Subtype int    `json:"Subtype"`
	Data    string `json:"Data"`
}

type UidType struct {
	Subtype int    `json:"Subtype"`
	Data    string `json:"Data"`
}

type TimeType struct {
	T  int `json:"t"`
	Ts struct {
		T int `json:"T"`
		I int `json:"I"`
	} `json:"ts"`
}

type UiType struct {
	Subtype int    `json:"Subtype"`
	Data    string `json:"Data"`
}

type Timestamp struct {
	T int `json:"T"`
	I int `json:"I"`
}

func ConvertToSql() {
	data, err := ioutil.ReadFile("oplog.json")
	if err != nil {
		log.Fatal(err)
	}

	var items []Item
	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Fatal(err)
	}

	sql := ""

	for _, item := range items {
		operation := item.Op
		tableName := strings.Split(item.Ns, ".")[1]
		data := item.O
		o := item.O2
		values := make(map[string]interface{})
		for k, v := range data {
			values[k] = v

		}
		oval := make(map[string]interface{})
		for k, v := range o {
			oval[k] = v

		}

		sql += "\n"
		sql += CreateSql(tableName, operation, values, oval)
	}
	fmt.Println(sql)
}
