package dbops

import (
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")

}

func TestMain(m *testing.M){
	//clearTables()
	m.Run()
	//clearTables()
}

func TestUserWorkFlow(t *testing.T){
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUserCredential)
}

func testAddUserCredential(t *testing.T) {

	err := AddUserCredential("zkk", "123")
	if err != nil {
		t.Errorf("testAddUserCredential error: %v", err)
	}
}

func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("zkk")
	if pwd != "123" || err != nil {
		t.Errorf("testGetUserCredential error: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("zkk", "123")
	if err != nil {
		t.Errorf("testDeleteUser error: %v", err)
	}
}

func testRegetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("zkk")
	if err != nil {
		t.Errorf("testRegetUserCredential error: %v", err)
	}
	if pwd != ""{
		t.Errorf("Deleting user test failed")
	}
}