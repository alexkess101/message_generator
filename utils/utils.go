package utils

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

func GetArguments(version string) []string {
	args := os.Args[1:]
	fmt.Printf("Running version: %s\n", version)

	if len(args) == 0 || len(args)%2 != 0 {
		fmt.Println("Error: Only accepts even number of arguments")
		os.Exit(1)
	}

	return args
}

func ApplyValueToMessage(fieldName, fieldValue string, messageValue reflect.Value) {
	fieldValueInt, err := strconv.ParseInt(fieldValue, 10, 64)
	if err != nil {
		fieldValueInt = 0
	}

	field := messageValue.FieldByName(strings.Title(fieldName))
	if !field.IsValid() {
		fmt.Printf("Error: field '%s' does not exist in Order struct\n", fieldName)
		os.Exit(1)
	}

	switch field.Kind() {
	case reflect.Int64:
		field.SetInt(fieldValueInt)
	case reflect.Uint64:
		field.SetUint(uint64(fieldValueInt))
	case reflect.Uint16:
		field.SetUint(uint64(fieldValueInt))
	case reflect.String:
		field.SetString(fieldValue)
	}
}

func WriteClipboard(str string) error {
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(str)); err != nil {
		return err
	}

	in.Close()
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func PrintSuccessMessage(context string) {
	fmt.Print("Success!! ")
	fmt.Print("\033[4m")
	fmt.Print("\033[1m")
	fmt.Printf("%s", context)
	fmt.Print("\033[0m")
	fmt.Print("\033[0m")
	fmt.Print(" test copied to clipboard")
}
