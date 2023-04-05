package TestDependency

import (
	"fmt"
	"github.com/Abhishek-Mali-Simform/TestDependency/Sample"
)

func DoSomething() {
	Sample.ExampleWrite()
	fmt.Println("Doing Some Work")
}
