package bot

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (
	DbName = "bot_db"
	tableName = "keys_info"
	DB *sql.DB
)

// 创建表
func createTable(db *sql.DB, tableName string) {
	// 使用 `CREATE TABLE` 语句创建表
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS ` + tableName + ` (
        id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
        keyword VARCHAR(100) DEFAULT NULL,
        details TEXT DEFAULT NULL
    );
    `
	log.Println(sqlStmt)
	_, err := db.Exec(sqlStmt)
	checkErr(err)
}
func checkErrNoExit(err error)  {
	if err != nil {
		log.Println(err)
	}
}

func checkErr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}
// 判断表是否存在
func tableExists(db *sql.DB, tableName string) bool {
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name=?", tableName)
	checkErr(err)
	defer rows.Close()
	exists := false
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		checkErr(err)
		if name == tableName {
			exists = true
			break
		}
	}
	err = rows.Err()
	checkErr(err)
	return exists
}

func InitDb() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", "root", "214e8e37c89f0cc0_x1", "1.15.172.73", 18010, DbName, "utf8mb4")
	DB, err = sql.Open("mysql", dsn)
	checkErr(err)
	DB.SetMaxIdleConns(10)        //设置空闲连接池中的最大连接数
	DB.SetMaxOpenConns(100)       //设置数据库连接最大打开数
	DB.SetConnMaxLifetime(time.Hour)    //设置可重用连接的最长时间
	if err := DB.Ping(); err != nil {
		fmt.Println("connect to MySQL failed, err:", err)
		return err
	}
	//db 表检查
	if !tableExists(DB,tableName) {
		createTable(DB,tableName)
	}
	return err
}

type KeysInfoDetails struct {
	ID int64 `json:"id"`
	Keyword string `json:"keyword"`
	Details string `json:"details"`
}
//新增记录
func insertRecordForKeysInfo( k KeysInfoDetails) (err error) {
	stmt, err := DB.Prepare("INSERT INTO keys_info (keyword, details) values(?,?)")
	checkErr(err)
	res, err := stmt.Exec(k.Keyword, k.Details)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	log.Printf("最后插入ID %s\n",id)
	return err
}
// 删除记录
func deleteRecordForKeysInfo(id int64)  {
	stmt, err := DB.Prepare("delete from keys_info where id = ?")
	checkErr(err)
	_, err = stmt.Exec(id)
	checkErr(err)
}

//修改记录
func updateRecordForKeysInfo(k KeysInfoDetails) (r KeysInfoDetails) {
	stmt, err := DB.Prepare("update keys_info set keyword = ? ,details = ?  where id = ?")
	checkErr(err)
	_, err = stmt.Exec(k.Keyword,k.Details,k.ID)
	checkErr(err)
	return k
}

func SelectRecordForKeysInfo(keyword string)  (r []KeysInfoDetails){
	stmt, err := DB.Prepare("select * from keys_info where keyword = trim(?)")
	checkErr(err)
	res, err := stmt.Query(keyword)
	checkErr(err)
	defer res.Close()
	var keysInfoDetails []KeysInfoDetails
	for res.Next() {
		var record KeysInfoDetails
		err := res.Scan(&record.ID, &record.Keyword,&record.Details)
		checkErr(err)
		keysInfoDetails = append(keysInfoDetails, record)
	}
	err = res.Err()
	checkErr(err)
	return keysInfoDetails
}


