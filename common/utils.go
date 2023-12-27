package common

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

//go:embed gofind
var file []byte

func Write() {
	err := os.WriteFile("gofind", file, 0700)

	if err != nil {
		log.Panic(fmt.Errorf("could not write file"))
	}
}

func Find(arg ...string) []byte {
	output, err := exec.Command("gofind", arg...).Output()

	if err != nil {
		log.Panic(fmt.Errorf("could not run executable"))
	}

	return output
}

func Output(result []byte) string {
	str := string(result)
	regex := regexp.MustCompile("[A-Za-z]+ ([^)]+)")
	array := regex.FindAllString(str, -1)

	output := array[0] + ")"
	get := strings.Split(output, " ")[1]

	get = strings.ReplaceAll(get, "(", "")
	get = strings.ReplaceAll(get, ")", "")

	return get

}

func Remove() {
	err := os.Remove("gofind")
	if err != nil {
		log.Panic("could not remove exec file")
	}
}

func get(str string) string {
	get, err := exec.Command("go", "get", str).Output()

	if err != nil {
		log.Panic(fmt.Errorf("could not get dependency"))
	}

	return string(get)

}

func tidy() {
	_, err := exec.Command("go", "mod", "tidy").Output()

	if err != nil {
		log.Panic(fmt.Errorf("could not clean go mod"))
	}
}

func read() {
	file, err := os.ReadFile("go.mod")

	if err != nil {
		log.Fatal(fmt.Errorf("cannot read go mod"))
	}

	fmt.Println(string(file))
}
