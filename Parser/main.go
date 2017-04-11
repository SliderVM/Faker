package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	file, err := os.Open("../data.csv")
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()
	
	/*
	fileSurname, err := os.Create("../Data/Surname.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer fileSurname.Close()

	fileName, err := os.Create("../Data/Name.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer fileName.Close()

	filePatronamic, err := os.Create("../Data/Patronamic.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer filePatronamic.Close()

	filePhone, err := os.Create("../Data/Phone.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer filePhone.Close()
	*/

	reader := csv.NewReader(file)
	reader.Comma = ';'	

	/*
	writerSurname := csv.NewWriter(fileSurname)
	writerSurname.Comma = ';'
	writerName := csv.NewWriter(fileName)
	writerName.Comma = ';'
	writerPatronamic := csv.NewWriter(filePatronamic)
	writerPatronamic.Comma = ';'
	writerPhone := csv.NewWriter(filePhone)
	writerPhone.Comma = ';'

	defer writerSurname.Flush()
	defer writerName.Flush()
	defer writerPatronamic.Flush()
	defer writerPhone.Flush()
	*/
	var surnames []string
	var names []string
	var patronamics []string
	var phones []string

	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		name := strings.Split(record[0], " ")
		surnames = append(surnames, name[0])
		names = append(names, name[1])
		patronamics = append(patronamics, name[2])
		phones = append(phones, record[1])
	}

	read, err := ioutil.ReadFile("../Data/data.go")
	if err != nil {
			panic(err)
	}	

	newContents := strings.Replace(string(read), "/*surname*/", getString(surnames), -1)
	newContents = strings.Replace(newContents, "/*name*/", getString(names), -1)
	newContents = strings.Replace(newContents, "/*patronamic*/", getString(patronamics), -1)
	newContents = strings.Replace(newContents, "/*phone*/", getString(phones), -1)
	err = ioutil.WriteFile("../Data/data.go", []byte(newContents), 0)
	if err != nil {
			panic(err)
	}

	/*
	err = writerSurname.Write(surnames)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = writerName.Write(names)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = writerPatronamic.Write(patronamics)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = writerPhone.Write(phones)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	*/
}

func getString(str []string) string{
	var r string
	for _, v := range str{
		r += `"` + v + `",`
	}
	return r
}