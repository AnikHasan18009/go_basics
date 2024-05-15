package db

import (
	"database/sql"
	"fmt"
	"user-service/types"
)

func InsertUserIntoDB(user types.NewUser) error {
	_, err := Db.Exec("insert into users(first_name,last_name,password) values($1,$2,$3)", user.FirstName, user.LastName, user.Password)
	return err
}

func RetriveUserByIdFromDB(id int) (types.User, error) {

	user := types.User{}

	err := Db.Get(&user, "select * from users where id =  $1", id)
	if user.FirstName == "" {
		err = fmt.Errorf("bad request")
	}
	return user, err
}

func CheckIfUserExistsInDB(id int) (bool, error) {
	var name string
	userExists := true
	var err = Db.Get(&name, "select first_name from users where id = $1", id)
	if name == "" || err != nil {
		userExists = false
	}
	return userExists, err

}

func RetriveUsersFromDB(users *[]types.User) ([]types.User, error) {

	if rows, err := Db.Query("SELECT id,first_name,last_name,password FROM users"); err != nil {

	}

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

func GetUsersFromResult(row *sql.Row, users *[]types.User) (*[]types.User, error) {

	// Exec
	if err != nil {
		return nil, err // Handle connection or query errors
	}
	defer rows.Close() // Close the rows object even on errors

	// Iterate through each row
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err // Handle scan errors
		}

		// Append the scanned user to the slice
		users = append(users, user)
	}

	// Check for errors after iterating
	if err := rows.Err(); err != nil {
		return nil, err // Handle any errors during iteration
	}

	// Handle empty result set (optional)
	if len(users) == 0 {
		// Log or return an empty slice based on your needs
		fmt.Println("No users found")
		// return users  // Return the empty slice
	}

	return users, nil
}
