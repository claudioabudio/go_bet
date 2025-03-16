package dog

import (
	"testing"
)

func TestYears(t *testing.T) {
	dogYears := Years(5)
	wanted := 35
	if dogYears != wanted {
		t.Error("got:", dogYears, "wanted", wanted)
	}

}

func TestYearsTwo(t *testing.T) {
	dogYears := YearsTwo(5)
	wanted := 35
	if dogYears != wanted {
		t.Error("got:", dogYears, "wanted", wanted)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
