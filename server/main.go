package main

import (
	"fmt"
	"github.com/rophie123/patient_registration/server/model"
	"github.com/rophie123/patient_registration/server/router"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	err = model.Init()
	if err != nil {
		return
	}
	engine := router.Init()
	err = engine.Run(":8080")
}
