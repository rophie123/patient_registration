package main

import (
	"fmt"
	"github.com/rophie123/patient_registration/model"
	"github.com/rophie123/patient_registration/router"
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
