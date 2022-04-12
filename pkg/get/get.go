package get

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// helper

type Students struct {
	XMLName     xml.Name  `xml:"students"`
	StudentList []Student `xml:"student"`
}

type Student struct {
	XMLName   xml.Name `xml:"student"`
	Firstname string   `xml:"type,attr"`
	Lastname  string   `xml:"type,attr"`
	City      string   `xml:"type,attr"`
}

func getStudents() {
	xmlFile, err := os.Open("db.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var students Students

	xml.Unmarshal(byteValue, &students)

	for i := 0; i < len(students.StudentList); i++ {
		fmt.Println("Firstname: " + students.StudentList[i].Firstname)
		fmt.Println("Lastname: " + students.StudentList[i].Lastname)
		fmt.Println("City: " + students.StudentList[i].City)
	}
}

func getCities(r *http.Request) {

	params := mux.Vars(r)

	xmlFile, err := os.Open("db.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var students Students

	xml.Unmarshal(byteValue, &students)

	for i := 0; i < len(students.StudentList); i++ {
		if students.StudentList[i].Firstname == params["firstname"] {
			fmt.Println("City: " + students.StudentList[i].City)
		}

	}
}
