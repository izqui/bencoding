package bencoding

import (
	"fmt"
	"testing"

	"github.com/izqui/helpers"
)

func TestString(t *testing.T) {

	i := helpers.RandomInt(1, 5000)
	str := helpers.RandomString(i)
	s, err := Decode([]byte(fmt.Sprintf("%d:%s", i, str)))

	if err != nil {

		t.Error(err)
	}

	if s.(string) != str {

		t.Error("String parsing wrong")
	}
}

func TestInteger(t *testing.T) {

	d := helpers.RandomInt(1, 100000)
	i, err := Decode([]byte(fmt.Sprintf("i%de", d)))

	if err != nil {

		t.Error(err)
	}

	if i.(int) != d {

		t.Error("Integer parsing wrong")
	}
}
