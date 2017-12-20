package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"strings"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/zd_big_wechat")
	if err != nil {
		panic("open database failed!")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic("ping database failed!")
	}

	stmt, err := db.Prepare(`insert tbl_group (name, en_name, short_name, region) values (?,?,?,?)`)
	if err != nil {
		panic(err)
	}

	execlFileName := "./edu.xlsx"
	xlFile, err := xlsx.OpenFile(execlFileName)
	if err != nil {
		panic("error open file")
	}

	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {
			cells := row.Cells
			name := strings.TrimSpace(fmt.Sprintf("%s", cells[2]))
			en_name := strings.TrimSpace(fmt.Sprintf("%s", cells[1]))
			short_name := strings.TrimSpace(fmt.Sprintf("%s", cells[4]))
			region := strings.TrimSpace(fmt.Sprintf("%s", cells[3]))

			res, err := stmt.Exec(name, en_name, short_name, region)
			if err != nil {
				panic(err)
			}

			_, err = res.RowsAffected()
			if err != nil {
				panic(err)
			}

			fmt.Printf("index:%s, name:%s, en_name:%s, short_name:%s, region:%s\n", index, name, en_name, short_name, region)

		}
	}
}
