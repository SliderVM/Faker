package faker

import (
	"testing"
	"strconv"
)



func TestGenerateSnils(t *testing.T){
	snils := GenerateSnils()

	if len(snils) != 11{
		t.Error("Не верная длинна СНИЛС, ожидалось 11 получено", len(snils), snils)
	}

	var sum int
	for i := 9; i>0; i--{
		c, err := strconv.ParseInt(string(snils[9-i]), 10, 64)

		if err != nil{
			t.Error("Ошибка при приведение типа string к int",err)
			break
		}
		sum = sum + (int(c) * i)
	}

	_sum := getCheckSum4Snils(sum)
	sum1 := snils[9:11]

	if _sum != sum1 {
		t.Error("Неврная контрольная сумма СНИЛС = "+snils+", ожидалос " + strconv.FormatInt(int64(sum),10) + " получено ", sum1)
	}
}