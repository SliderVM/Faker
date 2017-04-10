package faker

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

// GenerateUUID - сгенерировать uuid
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateSnils - сгенерировать СНИЛС персоны
func GenerateSnils() string{	
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

	return snils
}

func getCheckSum4Snils(sum int) string {	
	var res string
	if sum < 100{
		res = strconv.FormatInt(int64(sum),10)
	}
	if sum == 100 || sum == 101{
		res = "00"
	}
	if sum > 101{
		res = getCheckSum4Snils(sum%101)
	}
	return res
}