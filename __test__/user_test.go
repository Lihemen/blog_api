package tests

import (
	"testing"

	"github.com/lihemen/gin_api/models"
)


var user = models.User{
	FirstName: "John",
	LastName:  "Doe",
	Email:     "johndoe@gmail.com",
	Password:  "123456",

}

func Test_Get_User (t *testing.T) {}

func Test_Create_User (t *testing.T) {}

func Test_Get_All_Users (t *testing.T) {}

func Test_Delete_User (t *testing.T) {}

func Test_Update_User (t *testing.T) {}

