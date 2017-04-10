package faker

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var surnames []string
var names []string
var patronamics []string
var phones []string

func init() {
	var err error
	surnames, err = loadDate("Surname.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	names, err = loadDate("Name.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	patronamics, err = loadDate("Patronamic.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	phones, err = loadDate("Phone.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func loadDate(fileName string) ([]string, error) {
	file, err := os.Open("Data/" + fileName)
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return nil, err
	}
	// automatically call Close() at the end of current method
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	var res []string
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		res = append(res, record...)
	}

	return res, err
}

// GenerateUUID - сгенерировать uuid
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateSnils - сгенерировать СНИЛС персоны
func GenerateSnils() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var snils string
	var sum int
	for i := 9; i > 0; i-- {
		n := strconv.FormatInt(r1.Int63n(9), 10)
		snils += n
		c, _ := strconv.ParseInt(string(n), 10, 32)
		sum = sum + (int(c) * i)
	}

	checkSum := getCheckSum4Snils(sum)
	snils += checkSum

	return snils
}

func getCheckSum4Snils(sum int) string {
	var res string
	if sum < 100 {
		res = strconv.FormatInt(int64(sum), 10)
	}
	if sum == 100 || sum == 101 {
		res = "00"
	}
	if sum > 101 {
		res = getCheckSum4Snils(sum % 101)
	}
	return res
}

// GetSurname - сгенерировать фамилию
func GetSurname() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return surnames[r1.Intn(len(surnames)-1)]
}

// GetName - сгенерировать фамилию
func GetName() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return names[r1.Intn(len(names)-1)]
}

// GetPatronamic - сгенерировать фамилию
func GetPatronamic() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return patronamics[r1.Intn(len(patronamics)-1)]
}

// GetPhone - сгенерировать фамилию
func GetPhone() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return phones[r1.Intn(len(phones)-1)]
}
