package account
import(
	"fmt"
	"encoding/json"
	"github.com/NehemiahG7/GoKitchen/internal/database"
	//Contains the driver for the GoKitchen database
	_"github.com/lib/pq"
	"github.com/NehemiahG7/GoKitchen/internal/inventory"

)

//Login takes a username and password and checks the combination against the database. It returns the account id if present, and returns 0 if not present
func Login(account string, password string)(int){
	//Get connection to GoKitchen database
	db := database.OpenConn()
	defer db.Close()

	//Quarry database for username and password
	id := 0
	err := db.QueryRow(`SELECT id FROM accounts WHERE username = $1 AND password = $2`, account, password).Scan(&id)
	if err != nil{
		fmt.Printf("Err in QuerryRow: %s\n", err)
		fmt.Printf("id = %d\n", id)
		//return 0 if username is not found
		return id
	}
	fmt.Printf("id = %d\n", id)
	return id
}

//GetUsername takes the id of an user and returns the associated username
func GetUsername(id int) string{
	db := database.OpenConn()
	defer db.Close()

	var user string

	err := db.QueryRow(`SELECT username FROM accounts WHERE id = $1`, id).Scan(&user)
	if err != nil{
		fmt.Printf("GetUsername: Error retrieving name: %s\n", err)
		return "error"
	}

	return user
}

//GetInv querries the GoKitchen database for the items field associated with id
func GetInv(id int) inventory.Inventory{
	//Get connection to GoKitchen database
	db := database.OpenConn()
	defer db.Close()

	//inv to store items, and buff to get raw data from database
	inv := inventory.Inventory{}
	buff := make([]byte, 1024)

	//Querry data from database
	err := db.QueryRow(`SELECT items FROM accounts WHERE id = $1`, id).Scan(&buff)
	if err != nil{
		fmt.Printf("Err in GetInv: %s\n", err)
	}

	//Unmarshal information from database
	err = json.Unmarshal(buff, &inv)
	if err != nil{
		fmt.Printf("Err in GetInv Unmarshal: %s\n", err)
	}

	return inv
}

var defaultItems = `{"Inventory":{"canned":[],"dairy":[],"fruits":[],"grains":[],"meats":[],"other":[],"vegetables":[]}}`

//CreateAccount creates an entry in the GoKitchen database for username and password
func CreateAccount(username string, password string) int{
		//Get connection to GoKitchen database
		db := database.OpenConn()
		defer db.Close()

		res, err := db.Exec(`INSERT INTO accounts (username, password, items) VALUES ($1, $2, $3)`, username, password, defaultItems)
		if err != nil{
			fmt.Printf("CreateAccount error: %s\n", err)
		}
		fmt.Printf("Create resault: %s\n", res)
		return Login(username, password)
}
//AddItem adds the given item to the given cat in the given struct and passes it to the database
func AddItem(item string, cat string, accInv inventory.Inventory, id int){
	strs := []string{
		cat, item,
	}
	accInv.Add(strs)
	
	updateItems(accInv, id)
}
//AddCatagory adds the given catagory to the given struct and then passes it to the database
func AddCatagory(cat string, accInv inventory.Inventory, id int){
	strs := []string{cat,}
	accInv.Add(strs)
	
	updateItems(accInv, id)
}
func updateItems(accInv inventory.Inventory, id int){
	buff, err := json.Marshal(accInv)
	if err!=nil {
		fmt.Printf("updateItems: error Marshaling, %s\n", err)
	}
	database.UpdateDB("items", string(buff), id)
}