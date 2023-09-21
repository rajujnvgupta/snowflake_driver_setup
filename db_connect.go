package main 
import (
	"database/sql"
	"fmt"
	_ "github.com/snowflakedb/gosnowflake"
)

func main() {

	// another syntax for making database connection 
//	db, err := sql.Open("snowflake", "user:password@my_organization-my_account/mydb")
	db, err := sql.Open("snowflake", "rajuugupta:Raju@0987@https://qg83307.ap-southeast-1.snowflakecomputing.com/demo")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println(db)
}
