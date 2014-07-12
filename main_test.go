package bencoding

import (
	"fmt"
	"testing"

	"github.com/izqui/helpers"
)

func TestString(t *testing.T) {

	i := helpers.RandomInt(1, 5000)
	stra := helpers.RandomString(i)
	s, err := Decode([]byte(str(stra)))

	if err != nil {

		t.Error(err)
	}

	if s.(string) != stra {

		t.Error("String parsing wrong")
	}
}

func TestInteger(t *testing.T) {

	d := helpers.RandomInt(1, 100000)
	i, err := Decode([]byte(integer(d)))

	if err != nil {

		t.Error(err)
	}

	if i.(int) != d {

		t.Error("Integer parsing wrong")
	}
}

func TestList(t *testing.T) {

	elements := helpers.RandomInt(35, 1500)

	stra := "l"
	for i := 0; i < elements; i++ {
		if i%2 == 0 {
			stra += integer(i)
		} else {

			stra += str(helpers.RandomString(i))
		}
	}
	stra += "e"

	d, err := Decode([]byte(stra))

	if err != nil {

		t.Error(err)
	}

	if d.([]interface{})[34].(int) != 34 {

		t.Error("List parsing error")
	}
}

func TestDict(t *testing.T) {

	elements := helpers.RandomInt(35, 1500)

	stra := "d"
	for i := 0; i < elements; i++ {

		stra += str(helpers.RandomString(i)) //key
		if i%2 == 0 {
			stra += integer(i)
		} else {

			stra += str(helpers.RandomString(i))
		}
	}
	stra += str("thekey") + "l" + str("hola") + str("adios") + integer(19) + "e"
	stra += "e"

	d, err := Decode([]byte(stra))

	if err != nil {

		t.Error(err)
	}

	if d.(map[string]interface{})["thekey"].([]interface{})[2] != 19 {

		t.Error("Dict parsing error")
	}
}

func integer(i int) string {

	return fmt.Sprintf("i%de", i)
}

func str(s string) string {

	return fmt.Sprintf("%d:%s", len(s), s)
}
