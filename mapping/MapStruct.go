package mapping

import (
	"fmt"
)

type PersonalDetails struct {
    Class 		string		`json:"class"`
    Height		int			`json:"height"`
    Weight		int			`json:"weight"`
}

func MapStruct() (map[Student]*PersonalDetails) {

	// declare and initialize studentList
	studentList := make(map[Student]*PersonalDetails, 0)

	// append Student to studentList
	studentList[Student{"John", "C3445"}] = &PersonalDetails{"6H", 174, 67}
	studentList[Student{"James", "C2634"}] = &PersonalDetails{"5B", 168, 59}
	studentList[Student{"Peter", "C5375"}] = &PersonalDetails{"6M", 167, 70}

	fmt.Println(*studentList[Student{"John", "C3445"}])

	if value, exist := studentList[Student{"John", "C3445"}]; exist {
		fmt.Printf("Student %v exists", value)
	}

	if value, exist := studentList[Student{"Jimmy", "C4889"}]; !exist {
		fmt.Printf("Student %v does not exist", value)
	}

	return studentList
}