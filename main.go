package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type info struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

var i *info = &info{}

func init() {
	resp, err := http.Get("https://ipinfo.io")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
}

func main() {

	// Format string to output which includes available fields as %field_name%.
	var format string

	// Accept input flag for format string.
	flag.StringVar(&format, "f", "%city% %country%", "format of output")

	// Parse the input flags.
	flag.Parse()

	// Copy the format string to a new variable.
	output := format

	// Replace strings in format string.
	output = strings.Replace(output, "%ip%", i.IP, -1)
	output = strings.Replace(output, "%country%", i.Country, -1)
	output = strings.Replace(output, "%city%", i.City, -1)
	output = strings.Replace(output, "%loc%", i.Loc, -1)
	output = strings.Replace(output, "%org%", i.Org, -1)
	output = strings.Replace(output, "%postal%", i.Postal, -1)
	output = strings.Replace(output, "%readme%", i.Readme, -1)
	output = strings.Replace(output, "%region%", i.Region, -1)
	output = strings.Replace(output, "%timezone%", i.Timezone, -1)

	// Print the output.
	fmt.Print(output)

	// Return.
	return

}
