package db

import (
	"fmt"
	"user-service/types"
)

func InsertUserIntoDB(user types.NewUser) error {

	_, err := Db.Exec("insert into users(name,password) values($1,$2)", user.Name, user.Password)

	return err
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
