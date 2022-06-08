package controllers

import (
	"math/rand"
)

func GenerateColor() string {

	colorOptions := []string{
		"#2ecc71",
		"#3498db",
		"#9b59b6",
		"#34495e",
		"#f1c40f",
		"#e67e22",
		"#e74c3c",
		"#95a5a6",
	}

	return colorOptions[rand.Intn(len(colorOptions))]
}
