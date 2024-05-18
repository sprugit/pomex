package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sprugit/pomex-go/internal/utils"
)

func main() {

	filename := flag.String("target-name", "dependencies.csv", "Name of the resulting filename.")
	estimate := flag.Uint("estimate", 10, "Rough estimate of how many pom files will be found.")
	directory := flag.String("target-dir", "", "Target directory to srape for maven projects.")

	fmt.Println(&filename)

	poms, err := utils.FetchPOMs(*directory, *estimate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(poms)

}
