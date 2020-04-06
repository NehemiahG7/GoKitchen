package database

import(
	"fmt"
	"database/sql"
	//Contains the driver for the GoKitchen database
	_"github.com/lib/pq"
)
//CONNSTR is the connections string to pass into sql.Open for the accounts database
const CONNSTR string = "user=postgres password=pass dbname=GoKitchen host=localhost port = 8081  sslmode=disable"
//DRIVER is the database driver beings used for this database
const DRIVER string = "postgres"

//OpenConn returns a connection to the GoKitchen database
func OpenConn() *sql.DB{
	db, err := sql.Open(DRIVER, CONNSTR)
	if err != nil{
		fmt.Printf("OpenConn error: %s\n", err)
	}
	return db
}
//UpdateDB changes the corrosponding cell in id/column
func UpdateDB(column string, change string, id int){
	db := OpenConn()
	defer db.Close()

	res, err := db.Exec(`UPDATE accounts SET items = $1 WHERE id = $2`, change, id)

//	res, err := db.Exec(`UPDATE accounts SET $1 = $2 WHERE id = $3`, column, change, id)
	if err != nil {
		fmt.Printf("Error UpdateDB: %s\n", err)
	}
	fmt.Printf("UpdateDB resault: %s\n", res)
}