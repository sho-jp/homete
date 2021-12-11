package text

import (
	"testing"
)

func TestRandomNumber(t *testing.T) {
	expect := []int{0, 1, 2, 3, 4, 5}
	actual := RandomNumber(5)

	if !contains(expect, actual) {
		t.Errorf("%d is not included in %d", actual, expect)
	}
}

func TestDefault(t *testing.T) {
	actual := Default().Ja
	if actual == "" {
		t.Errorf("%s is empty", actual)
	}

	actual = Default().En
	if actual == "" {
		t.Errorf("%s is empty", actual)
	}
}

func TestMotto(t *testing.T) {
	actual := Motto().Ja
	if actual == "" {
		t.Errorf("%s is empty", actual)
	}

	actual = Motto().En
	if actual == "" {
		t.Errorf("%s is empty", actual)
	}
}

func contains(ints []int, i int) bool {
	for _, v := range ints {
		if v == i {
			return true
		}
	}

	return false
}
