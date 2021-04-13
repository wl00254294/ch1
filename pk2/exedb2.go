package dbmenu

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"runtime"

	message "eric.com/go/ch1/msg"
	a "github.com/ibmdb/go_ibm_db"
)

func ConnectDB2(host string, port string, dbname string, user string, pwd string) *a.DBP {

	con := "HOSTNAME=" + host + ";PORT=" + port + ";DATABASE=" + dbname + ";UID=" + user + ";PWD=" + pwd
	pool := a.Pconnect("PoolSize=1000")

	// SetConnMaxLifetime will take the value in SECONDS
	db := pool.Open(con, "SetConnMaxLifetime=30")
	fmt.Println("success build connection")
	return db
}

func CloseDB2(db *a.DBP) {
	db.Close()
	fmt.Println("success close connection")
}

func SelectDB2(db *a.DBP, sql string, ch chan string, parm []interface{}) {
	runtime.GOMAXPROCS(20)
	go exquery(db, sql, ch, parm)
}

func GetStringValue(col string, val string) string {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(val), &m)
	return reflect.ValueOf(m[col]).Interface().(string)

}

func GetIntValue(col string, val string) int {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(val), &m)
	return reflect.ValueOf(m[col]).Interface().(int)

}

func exquery(db *a.DBP, sql2 string, ch chan string, parm []interface{}) {

	st, err := db.Prepare(sql2)
	if err != nil {
		log.Printf("Error:%s, %s", message.Content["sqlqryerror"].Code, err)
		fmt.Println(err)
	}
	rows, err := st.Query(parm...)

	if err != nil {
		fmt.Println(err)
	}
	cols, _ := rows.Columns()
	vals := make([]interface{}, len(cols))
	defer rows.Close()
	for rows.Next() {

		for i := range cols {

			vals[i] = new(sql.RawBytes)
		}
		rows.Scan(vals...)
		entry := make(map[string]interface{})
		for i, value := range vals {

			content := reflect.ValueOf(value).Interface().(*sql.RawBytes)
			entry[cols[i]] = string(*content)
		}

		jsonrowData, _ := json.Marshal(entry)

		ch <- string(jsonrowData)
	}

	close(ch)
}

func InUpDeDB2(db *a.DBP, sql2 string, datas [][]interface{}, commitnum int) {

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("open db2 database fail", err)
	}

	for i, data := range datas {

		_, err := tx.Exec(sql2, data...)
		if err != nil {
			fmt.Println("execute failed,error： ", err)
			fmt.Println("error insert data : ", data)
			tx.Rollback()
			//return 失敗印出繼續往下處理

		}
		count := i + 1

		if count%commitnum == 0 {

			tx.Commit()
		}

	}

	tx.Commit()
}
