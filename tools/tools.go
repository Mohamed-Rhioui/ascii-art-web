package tools

import (
	"log"
	"os"
)

func Read_Input(infile_name string) string {
	buffer, err := os.ReadFile(infile_name)
	CheckError(err, "Error: failed to read infile: \""+infile_name+"\"!!")
	if len(buffer) == 0 {
		log.Fatalln("Error: infile is empty: \"" + infile_name + "\"!!")
	}
	return string(buffer)
}

func RemoveEmptyString(slice []string) []string {
	var result []string
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			result = append(result, slice[i])
		}
	}
	return result
}

func CHeckTemplate(banner string) string {
	var data string
	if banner == "standard" {
		data = Read_Input("Templates/standard.txt")
	} else if banner == "shadow" {
		data = Read_Input("Templates/shadow.txt")
	} else if banner == "thinkertoy" {
		data = Read_Input("Templates/thinkertoy.txt")
	} else {
		log.Fatalln("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
	return data
}



func IsAllNl(result string) bool {
	for _, char := range result {
		if char != '\n' {
			return false
		}
	}
	return true
}

func CheckError(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg+"\n")
	}
}
