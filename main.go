package main

import (
	"log"
	"os"

	common "github.com/OlafRadicke/banwoldang/common"
)

func main() {
	startPasth := ""
	givenArguments := len(os.Args)
	minimumArguments := 2
	if givenArguments < minimumArguments {
		log.Fatal("To less arguments!")
	} else {
		programName := os.Args[0]
		startPasth = os.Args[1]
		log.Println("Arg numbers: ", givenArguments)
		log.Println("Starting the application: ", programName)
	}

	fileTree := common.FileTree{}
	fileTree.StartPath = startPasth
	fileTree.Collecting()
	log.Println("Findings: ", fileTree.Findings)
}

func checkInput(input string) {

}
