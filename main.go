package main

import (
	"fmt"
	"os"
	"groundhog"
)

func printHelp() {
    fmt.Println("SYNOPSIS")
    fmt.Println("\t./groundhog period")
    fmt.Println("")
    fmt.Println("DESCRIPTION")
    fmt.Println("\tperiodt\t\tthe number of days defining a period")
}

func main() {
    if groundhog.CheckHelp() {
        printHelp()
        os.Exit(0)
    }
    if !groundhog.CheckArgs() {
        printHelp()
        os.Exit(84)
    }
    groundhog.Groundhog()
}