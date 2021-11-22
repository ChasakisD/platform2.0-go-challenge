package sql

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) {
	sqlFile := readSQLFile("sql/schema.sql")
	if len(sqlFile) <= 0 {
		return
	}

	execSchemaSQL(db, sqlFile)
}

func readSQLFile(path string) string {
	sqlFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(sqlFile)
}

func execSchemaSQL(db *gorm.DB, sqlFileString string) {

	if err := db.Exec(sqlFileString).Error; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("SQL Executed successfully!")
}
