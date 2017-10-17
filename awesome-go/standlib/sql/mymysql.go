package main

import (
	"fmt"
	"os"

	"github.com/ziutek/mymysql/mysql"

	// native engine
	_ "github.com/ziutek/mymysql/native"
	// thread safe engine
	// _ "github.com/ziutek/mymysql/thrsafe
)

func main() {
	user, pass, dbname := "root", "root", "test"
	db := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	rows, res, err := db.Query("select * from X where id > %d", 20)
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		for _, col := range row {
			if col == nil {
				//col has null value
			} else {
				// do something with text in col
			}
		}
		// get vlue for a row
		val1 := row[1].([]byte)
		// you can use it
		os.Stdout.Write(val1)

		// you can get converted value
		number, str, bignum := row.Int(0), row.Str(1), row.MustUint(2)

		// you may get values by columns name
		first := res.Map("FirstColumn")
		second := res.Map("SecondColumn")
		val1, val2 := row.Int(first), row.Str(second)
	}

	// prepaed statements
	stmt, err := db.Prepare("insert into x values (?, ?)")
	checkError(err)

	data = new(Data)
	for {
		err := getData(data)
		if err == endOfData {
			break
		}
		checkError(err)
		_, err = stmt.Run(data.Id, data.Tax)
		checkError(err)
	}

}

type Data struct {
	Id  int
	Tax *float32
}

func checkError(err error) {
	fmt.Printf("%s\n", err)
}
