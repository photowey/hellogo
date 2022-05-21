package mysql

import (
	"database/sql"
	"fmt"

	"github.com/hellogo/internal/jsonz"
)

/**
 * database.mysql
 */

const (
	driverMysql = "mysql"
	dsn         = "root:root@tcp(192.168.1.11:3307)/notices?charset=utf8mb4&parseTime=True&loc=Local"
)

const (
	BaseTableInfoSql = `
SELECT 
TABLE_NAME AS tableName, 
TABLE_COMMENT AS tableComment 
FROM information_schema.TABLES 
WHERE TABLE_SCHEMA = ? `
	BaseColumnInfoSql = `
SELECT
TABLE_NAME AS tableName,
COLUMN_NAME AS columnName,
COLUMN_COMMENT AS columnComment,
IS_NULLABLE AS notNull,
DATA_TYPE AS dataType,
CHARACTER_MAXIMUM_LENGTH AS dataLength,
COLUMN_KEY AS primaryKey,
NUMERIC_PRECISION AS maxBit,
NUMERIC_SCALE AS minBit
FROM
information_schema.COLUMNS
WHERE
TABLE_SCHEMA = ?
AND TABLE_NAME = ?
ORDER BY TABLE_NAME,
ORDINAL_POSITION`
)

type Database struct {
	Name   string
	Tables []Table `json:"tables"`
}

type Table struct {
	Name    string         `json:"tableName"`
	Comment sql.NullString `json:"tableComment"`
	Columns []Column       `json:"columns"`
}

type Column struct {
	TableName     string         `json:"tableName"`
	ColumnName    string         `json:"columnName"`
	ColumnComment sql.NullString `json:"columnComment"`
	NotNull       sql.NullString `json:"notNull"`
	DataType      sql.NullString `json:"dataType"`
	DataLength    sql.NullString `json:"dataLength"`
	PrimaryKey    sql.NullString `json:"primaryKey"`
	MaxBit        sql.NullString `json:"maxBit"`
	MinBit        sql.NullString `json:"minBit"`
}

func Init() {
	// Asia/Shanghai
	db, err := sql.Open(driverMysql, dsn)
	if err != nil {
		fmt.Println("数据库连接错误" + err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping数据库错误" + err.Error())
		return
	}
	rows, err := db.Query(BaseTableInfoSql, "notices")

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	if err != nil {
		fmt.Println("查询数据库信息错误" + err.Error())
		return
	}

	database := &Database{
		Name:   "notices",
		Tables: make([]Table, 0),
	}

	for rows.Next() {
		table := Table{
			Columns: make([]Column, 0),
		}
		err := rows.Scan(&table.Name, &table.Comment)
		if err != nil {
			fmt.Printf("scan table failed, err:%v\n", err)
			return
		}
		rowColumns, err := db.Query(BaseColumnInfoSql, database.Name, table.Name)
		defer rowColumns.Close()
		if err != nil {
			fmt.Printf("query column failed, err:%v\n", err)
			return
		}

		for rowColumns.Next() {
			column := Column{}
			err := rowColumns.Scan(
				&column.TableName, &column.ColumnName, &column.ColumnComment,
				&column.NotNull, &column.DataType, &column.DataLength,
				&column.PrimaryKey, &column.MaxBit, &column.MinBit,
			)
			if err != nil {
				fmt.Printf("scan column failed, err:%v\n", err)
				return
			}
			table.Columns = append(table.Columns, column)
		}

		database.Tables = append(database.Tables, table)
	}

	databaseInfo, _ := jsonz.String(database)
	fmt.Printf("the database info is:\n%s\n", databaseInfo)
}
