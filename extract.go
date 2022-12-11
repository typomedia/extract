package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	flag "github.com/spf13/pflag"
	"log"
	"os"
)

const version = "1.0"

const usage = `Usage: extract [OPTION]... -i file -o file
  -h, --help
  -i, --input file     json file
  -o, --output file    
  -b, --base64         bade64 encoded output string
  -r, --reverse        reverse the output string

Example: extract --base64 --reverse --input file.json --output file.keys
`

func main() {
	println("extract by Philipp Speck [Version " + version + "]")
	println("Copyright (C) 2022 Typomedia Foundation.")
	println()

	var input string
	var output string

	flag.StringVarP(&input,
		"input",
		"i",
		"",
		"Line separated file input")
	flag.StringVarP(&output,
		"output",
		"o",
		"",
		"Line separated file output")
	flag.BoolP("help", "h", false, "")
	flag.Usage = func() {
		//flag.PrintDefaults()
		fmt.Print(usage) // override default usage
	}
	reverse := flag.BoolP("reverse", "v", false, "")
	encode := flag.BoolP("base64", "b", false, "")
	flag.Parse()

	args := flag.Args()

	if len(input) == 0 && len(args) == 0 {
		flag.Usage()
	}

	if len(args) > 0 {
		for _, path := range args {
			fmt.Println(path)
		}
	}

	if input != "" {
		content, err := os.ReadFile(input)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		c := make(map[string]json.RawMessage)

		err = json.Unmarshal(content, &c)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		keys := make([]string, len(c))

		i := 0
		for s, _ := range c {
			keys[i] = s
			i++
		}

		//fmt.Printf("%#v\n", keys)

		data, _ := json.Marshal(keys)
		//fmt.Println(string(data))

		if *encode {
			data = []byte(base64.StdEncoding.EncodeToString(data))

			if *reverse {
				data = []byte(reverseStr(string(data)))
			}
		}

		err = os.WriteFile(output, data, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func reverseStr(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
