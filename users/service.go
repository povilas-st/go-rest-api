package users

import "io/ioutil"

const usersFile = "/go/src/go-rest-api/users.json"

//todo: implement GetUsers and GetUser methods

func readUsersFile() []byte {
	data, err := ioutil.ReadFile(usersFile)
	if err != nil {
		panic(err)
	}
	return data
}

//func writeUsersFile(users []User) {
//	data, err := json.Marshal(users)
//	if err != nil {
//		panic(err)
//	}
//	err = ioutil.WriteFile(usersFile, data, 0644)
//	if err != nil {
//		panic(err)
//	}
//}
