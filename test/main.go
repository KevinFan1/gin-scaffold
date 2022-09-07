package main

import (
	"code/gin-scaffold/models"
	"encoding/json"
	"fmt"
)

func main() {
	var user models.User
	data := []byte(`{"id":1,"created_at":"2022-08-26 12:14:34","updated_at":"2022-08-26 12:14:34","deleted_at":null,"username":"user1","password":"","role_id":1,"role":{"id":1,"created_at":"2022-08-26 16:13:28","updated_at":"2022-08-26 16:13:30","deleted_at":null,"name":"管理员","e":"admin"}}`)
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}
