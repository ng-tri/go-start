package main

import (
	"go_start/advance"
	"go_start/basic"
	"go_start/controllers"
	"go_start/intermediate"
	"go_start/practice"
)

func main() {
	advance.RunChannel()
	advance.RunGoroutine()

	basic.RunTest()
	basic.RunBasic()

	controllers.Http()

	intermediate.RunInterface()
	// intermediate.WriteFile()
	// intermediate.ReadFile()
	// intermediate.WriteCsv()
	// intermediate.LogGoRoutine()
	// intermediate.RunCustomer()
	// intermediate.RunProduct()
	// intermediate.RunOrder()
	// intermediate.InfoProduct()
	// intermediate.InfoOrder()
	// intermediate.InOrder()

	practice.RunPractice()
}
