package main

import (
	"fmt"
	"os"
	"os/user"
	"compiler/evaluate"

)

func main(){
	
	user, err := user.Current()

	if err != nil{
		panic(err)
	}

	fmt.Printf("Hello %s!This is my IDK language\n ",user.Username)

	fmt.Printf("You can type something in command line,but first check ../grammar/toke.go\n")

	repl.Start(os.Stdin,os.Stdout)


}
