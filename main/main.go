package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/quantum-craft/go-homework/main/myfolder01"
	"github.com/quantum-craft/go-homework/main/myfolder01/mypackage02"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(10))

	//sortalgorithm.Test2()
	myfolder01.TestPkg()
	mypackage02.TestAaaaa()
}
