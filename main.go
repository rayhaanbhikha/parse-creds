package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type creds struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func formatValue(s string) string {
	prefix := "AWS"
	r := regexp.MustCompile(`[A-Z][a-z]*`)

	strMatches := r.FindAllString(s, -1)
	n := len(strMatches)

	newString := ""
	for i, subString := range strMatches {
		newString += strings.ToUpper(subString)
		if len(subString) > 1 && i < n-1 {
			newString += "_"
		}
	}
	return fmt.Sprintf("%s_%s", prefix, newString)
}

func main() {

	var c creds

	data, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)

	err = json.Unmarshal(data, &c)
	checkErr(err)

	ct := reflect.ValueOf(&c).Elem()
	typeOfC := ct.Type()

	for i := 0; i < ct.NumField(); i++ {
		fmt.Printf("export %s=\"%v\"\n", formatValue(typeOfC.Field(i).Name), ct.Field(i))
	}

}
