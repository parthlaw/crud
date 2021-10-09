package db

import (
	"fmt"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	err := mgm.SetDefaultConfig(nil, "assignment", options.Client().ApplyURI("mongodb://mongo:0yzfuEzh36pmDLihJ9mA@containers-us-west-19.railway.app:6907"))
	fmt.Println("Connected")
	if err != nil {
		panic(err)
	}
}
