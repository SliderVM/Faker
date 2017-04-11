package faker

import (
	
	
	"github.com/google/uuid"
	
	"math/rand"
	
	"strconv"
	"time"	
	"github.com/SliderVM/Faker/Data"
	
)

var surnames []string
var names []string
var patronamics []string
var phones []string

func init() {	
	d := data.D
	surnames = d.Surnames
	names = d.Names
	patronamics = d.Patronamics
	phones = d.Phones
	rand.Seed(time.Now().UnixNano())
	/*
	var err error
	surnames, err = loadDate("Surname.csv")
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(os.Getenv("GOPATH"))
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
	*/	
}

/*
func loadDate(fileName string) ([]string, error) {
	file, err := os.Open(os.Getenv("GOPATH") + "github.com/SliderVM/Faker/Data/" + fileName)
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
*/

// GenerateUUID - сгенерировать uuid
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateSnils - сгенерировать СНИЛС персоны
func GenerateSnils() string {	
	var snils string
	var sum int
	for i := 9; i > 0; i-- {
		n := strconv.FormatInt(rand.Int63n(9), 10)
		snils += n
		c, _ := strconv.ParseInt(string(n), 10, 32)
		sum = sum + (int(c) * i)
	}

	checkSum := getCheckSum4Snils(sum)
	
	snils += checkSum
	
	seed, _ := strconv.ParseInt(snils,10,64)	
	rand.Seed(seed)

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

	if len(res) < 2{
		res = "0" + res
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
