package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	var B Root

	dat, err := os.ReadFile("./test/test.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(dat, &B)
	if err != nil {
		panic(err)
	}

	log.Printf("block one name is :%s", B.FirstBlock.Name)
}
