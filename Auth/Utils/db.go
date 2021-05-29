package Utils

import (
	config "Auth/Config"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

/*
	Description: Using gorm as ORM
*/
func GormConnectDB() bool {

	if dbConn == nil {
		var err error
		dsn := config.DB.Username + ":" + config.DB.Password + "@tcp(" + config.DB.Endpoint + ":3306)/" + config.DB.DBName
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println("db connection error response:  error:" + err.Error())
			return false
		}
	}

	return true
}

func RegisterUserInDB(data Register) int64 {
	fmt.Printf("data is %+v", data)
	if dbConn == nil {
		GormConnectDB()
	}
	result := dbConn.Exec("Insert into tblRegister (username,password,admin) values (?,?,?)", data.Username, data.Password, data.IsAdmin)
	fmt.Printf("result is %+v", result)
	return result.RowsAffected

}

func IsLoginExists(data Login) bool {
	var login Login
	fmt.Printf("data is %+v", data)
	if dbConn == nil {
		GormConnectDB()
	}
	dbConn.Table("tblRegister").Where("username = ?", data.Username).Find(&login)
	return login.Username != ""

}

func IsAdmin(username string) bool {
	var login Login
	if dbConn == nil {
		GormConnectDB()
	}
	dbConn.Table("tblRegister").Where("username = ?", username).Where("admin = true").Find(&login)
	return login.Username != ""

}

func IsUserRole(username string) bool {
	var login Login
	if dbConn == nil {
		GormConnectDB()
	}
	dbConn.Table("tblRegister").Where("username = ?", username).Where("admin = false").Find(&login)
	return login.Username != ""

}

func GetAllUsers() []Login {
	var logins []Login
	if dbConn == nil {
		GormConnectDB()
	}
	dbConn.Table("tblRegister").Find(&logins)
	return logins

}

func GetUserById(Id int) []Login {
	var logins []Login
	if dbConn == nil {
		GormConnectDB()
	}
	dbConn.Table("tblRegister").Where("Id = ?", Id).Find(&logins)
	return logins

}

func GetUserByName(name string) Login {
	var logins Login
	if dbConn == nil {
		GormConnectDB()
	}
	fmt.Println(name)
	dbConn.Table("tblRegister").Where("username = ?", name).Find(&logins)
	return logins

}

func UpdateUser(Id int) int64 {
	if dbConn == nil {
		GormConnectDB()
	}
	result := dbConn.Exec("Update tblRegister set admin = true where Id = ?", Id)
	fmt.Printf("result is %+v", result)
	return result.RowsAffected

}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"admin"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
