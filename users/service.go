package users

import (
	"encoding/json"
	"io/ioutil"
)

const usersFile = "/go/src/go-rest-api/users.json"

func GetUsers() []User {
	var users []User
	data := readUsersFile()
	json.Unmarshal(data, &users)

	return users
}

func GetUser(id int) User {
	var user User
	users := GetUsers()
	for _, value := range users {
		if value.Id == id {
			user = value
		}
	}

	return user
}

func CreateUser(user User) User {
	users := GetUsers()
	user.Id = 1
	for _, value := range users {
		if value.Id >= user.Id {
			user.Id = value.Id + 1
		}
	}
	users = append(users, user)
	writeUsersFile(users)

	return user
}

func UpdateUser(id int, user User) User {
	users := GetUsers()
	for index, value := range users {
		if value.Id == id {
			users[index].Email = user.Email
			users[index].Password = user.Password
			user = users[index]
			break
		}
	}
	writeUsersFile(users)

	return user
}

func DeleteUser(id int) {
	users := GetUsers()
	var mark int
	for index, value := range users {
		if value.Id == id {
			mark = index
			break
		}
	}
	users = append(users[:mark], users[mark+1:]...)
	writeUsersFile(users)
}

func writeUsersFile(users []User) {
	data, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(usersFile, data, 0644)
	if err != nil {
		panic(err)
	}
}

func readUsersFile() []byte {
	data, err := ioutil.ReadFile(usersFile)
	if err != nil {
		panic(err)
	}
	return data
}
