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

	stmt, err := db.Prepare(`insert tbl_group (en_name,logo, name,region) values (?,?,?,?)`)
	if err != nil {
		panic(err)
	}

	execlFileName := "./edu_canada.xlsx"
	xlFile, err := xlsx.OpenFile(execlFileName)
	if err != nil {
		panic("error open file")
	}

	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {
			cells := row.Cells
			en_name := strings.TrimSpace(fmt.Sprintf("%s", cells[0]))
			logo := strings.TrimSpace(fmt.Sprintf("%s", cells[1]))
			name := strings.TrimSpace(fmt.Sprintf("%s", cells[2]))
			region := strings.TrimSpace(fmt.Sprintf("%s", cells[3]))

			res, err := stmt.Exec(name, logo, en_name, region)
			if err != nil {
				panic(err)
			}

			_, err = res.RowsAffected()
			if err != nil {
				panic(err)
			}

			fmt.Printf("index:%s, name:%s, logon:%s, en_name:%s, region:%s \n", index, name, logo, en_name, region)

		}
	}
}
