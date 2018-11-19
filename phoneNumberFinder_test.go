package phonenumber

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {

	Find(`
	Nomor saya 085774568550
	silahkan cari saya di
	+6285774568550
	`)

}

func TestReadPhoneNumber(t *testing.T) {
	cases := []struct {
		s      string
		res    string
		resLen int
	}{
		{
			s:      "+62 8577456 8550 asd",
			res:    "+6285774568550",
			resLen: 16,
		},
	}
	for _, val := range cases {
		res, _, len := readPhoneNumber(val.s)
		if res != val.res {
			fmt.Println("res:", res)
			t.Fail()
		}

		if len != val.resLen {
			fmt.Println("len :", len)
			t.Fail()
		}
	}
}
