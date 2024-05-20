package db

import (
	"fmt"
	"log"
	"user-service/types"
)

func InsertUserIntoDB(user types.NewUser) error {

	_, err := Db.Exec("insert into users(name,email,password) values($1,$2,$3)", user.Name, user.Email, user.Password)

	return err
}
func GetUserOrdersFromDB(userId int) ([]types.Order, error) {
	var orders []types.Order
	err := Db.Select(&orders, "select * from orders where user_id= $1", userId)
	fmt.Println(orders)
	return orders, err

}

func RetriveUserByIdFromDB(id int) (types.User, error) {

	user := types.User{}

	err := Db.Get(&user, "select * from users where id =  $1", id)
	if user.Name == "" {
		err = fmt.Errorf("bad request")
	}
	return user, err
}

func CheckIfUserExistsInDB(id int) (bool, error) {
	var name string
	userExists := true
	var err = Db.Get(&name, "select name from users where id = $1", id)
	if name == "" || err != nil {
		userExists = false
	}
	return userExists, err

}

func RetriveUsersFromDB(users []types.User) ([]types.User, error) {

	err := Db.Select(&users, "SELECT  * FROM users")

	if len(users) == 0 {
		err = fmt.Errorf("database error")
	}
	return users, err
}

func DeleteUserByIdFromDB(id int) error {

	_, err := Db.Exec("delete from users where id = $1", id)

	return err

}

func UpdateUserInDB(id int, updatedUser types.NewUser) error {

	_, err := Db.Exec("update users set name = $1 ,password = $2 where id = $3", updatedUser.Name, updatedUser.Password, id)

	return err
}

func CountOfRecordsBasedOnConditions(conditions map[string]string) int {
	var count int
	query := "select count(name) from users where "
	for k, v := range conditions {
		switch {
		case k == "id":
			query += k + " = " + v + " and "
		default:
			query += k + " = '" + v + "' and "
		}

	}
	query += "true"
	fmt.Println(query)
	err := Db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}
