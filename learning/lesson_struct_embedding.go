package main

import (
	"fmt"
)

// 共通のフィールドを持つ構造体を定義
type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d", u.ID, u.Name, u.Age)
}

// 管理者構造体にUserを埋め込む
type Admin struct {
	User
	Permissions []string
}

func (a Admin) GetPermissions() string {
	return fmt.Sprintf("Admin %s has permissions: %v", a.Name, a.Permissions)
}

func main() {
	admin := Admin{
		User: User{
			ID:   1,
			Name: "Alice",
			Age:  30,
		},
		Permissions: []string{"read", "write", "delete"},
	}

	// Userのメソッドを呼び出す
	fmt.Println(admin.GetDetails())
	fmt.Println(admin.GetPermissions())
}
