package main

import (
	service "eric.com/go/ch1/pk1"
	//"fmt"
	//"time"
	//dbmenu "eric.com/go/ch1/pk2"
	//	jwtctrl "eric.com/go/ch1/pk3"
	//"encoding/json"
)

func main() {
	/*
		fmt.Println(time.Now())
		token := jwtctrl.GenJWToken(36000, 1, "eric", "10.8.0.116")
		fmt.Println("token==>", token)

		ret, err := jwtctrl.ParseToken(token, "10.8.0.111")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("userinfo: %v\n", ret.RequestIp)
		}

		fmt.Println(time.Now())
	*/
	service.ServerOn()

}

/*
func insertexample() {
	var db = dbmenu.ConnectDB2("134.251.80.228", "55000", "CR", "CRAP1", "1qaz2wsx")
	data := [][]interface{}{{"R1"}, {"R2"}, {"R3"}, {"R4"}}
	dbmenu.InUpDeDB2(db, "INSERT INTO ECSCRDB.eric_test (A) VALUES (?)", data, 1000)
	dbmenu.CloseDB2(db)
}

func selectexample() {
	var db = dbmenu.ConnectDB2("134.251.80.228", "55000", "CR", "CRAP1", "1qaz2wsx")
	ch := make(chan string, 100000)
	log.Println("==Start select")
	testary := []interface{}{"5466010000021191", 144} //parameter
	dbmenu.SelectDB2(db, "SELECT CARD_NO,P_SEQNO FROM ECSCRDB.CRD_CARD where CARD_NO = ? and RECNO = ?", ch, testary)
	log.Println("==after select")
	dbmenu.CloseDB2(db)

	for i := range ch {
		data := dbmenu.GetStringValue("CARD_NO", i)
		fmt.Println(data)
	}
}
*/
