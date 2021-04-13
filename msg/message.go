package message

import (
	"log"
)

type MsgEnum struct {
	Code    string `json:"Code"`
	Message string `json:"message"`
}

var Content = map[string]MsgEnum{}

func init() {
	log.Println("===Start Init Message Code===")
	Content["sqlprerror"] = MsgEnum{"E00001", "sql prepare error"}
	Content["jsonconverror"] = MsgEnum{"E00002", "json parser error"}
	Content["configreaderror"] = MsgEnum{"E00003", "config file read error"}
	Content["configdataerror"] = MsgEnum{"E00004", "service not found in config file"}
	Content["sqlqryerror"] = MsgEnum{"E00005", "sql query error"}
	log.Println("===End Init Message Code===")
}
