package main

import (
	"time"
	"math/rand"
	"fmt"
	"log"
	"os"
)

var message string

func main () {
	colors := []string{ 
		"#1abc9c",
		"#2ecc71",
		"#3498db",
		"#9b59b6",
		"#34495e",
		"#f1c40f",
		"#e67e22",
		"#c0392b",
	}

	rand.Seed(time.Now().Unix())

	for {	
		f, err := os.Create("color.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(fmt.Sprint(colors[rand.Intn(len(colors))]))

		if err2 != nil {
			log.Fatal(err2)
		}

		time.Sleep(2 * time.Second)
	}	
}