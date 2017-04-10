package faker

import (
	"strconv"
	"testing"
)

func TestGenerateSnils(t *testing.T) {
	snils := GenerateSnils()

	if len(snils) != 11 {
		t.Error("Не верная длинна СНИЛС, ожидалось 11 получено", len(snils), snils)
	}

	var sum int
	for i := 9; i > 0; i-- {
		c, err := strconv.ParseInt(string(snils[9-i]), 10, 64)

		if err != nil {
			t.Error("Ошибка при приведение типа string к int", err)
			break
		}
		sum = sum + (int(c) * i)
	}

	_sum := getCheckSum4Snils(sum)
	sum1 := snils[9:11]

	if _sum != sum1 {
		t.Error("Неврная контрольная сумма СНИЛС = "+snils+", ожидалос "+strconv.FormatInt(int64(sum), 10)+" получено ", sum1)
	}
}

func TestGetSurname(t *testing.T) {
	surname := GetSurname()

	if surname == "" {
		t.Error("Неудалось сгенерировать фамилию")
	}

	var ex bool
	ex = false
	for _, v := range surnames {
		if v == surname {
			ex = true
		}
	}
	if !ex {
		t.Error("Фамилия получена не верно, вернулось значение", surname)
	}
}

func TestGetName(t *testing.T) {
	name := GetName()

	if name == "" {
		t.Error("Неудалось сгенерировать имя")
	}

	var ex bool
	ex = false
	for _, v := range names {
		if v == name {
			ex = true
		}
	}
	if !ex {
		t.Error("Имя получено не верно, вернулось значение", name)
	}
}

func TestGetPatronamic(t *testing.T) {
	patronamic := GetPatronamic()

	if patronamic == "" {
		t.Error("Неудалось сгенерировать имя")
	}

	var ex bool
	ex = false
	for _, v := range patronamics {
		if v == patronamic {
			ex = true
		}
	}
	if !ex {
		t.Error("Отчестово получено не верно, вернулось значение", patronamic)
	}
}

func TestGetPhone(t *testing.T) {
	phone := GetPhone()

	if phone == "" {
		t.Error("Неудалось сгенерировать имя")
	}

	var ex bool
	ex = false
	for _, v := range phones {
		if v == phone {
			ex = true
		}
	}
	if !ex {
		t.Error("Телефон получен не верно, вернулось значение", phone)
	}
}
