package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

//List is a list returned by Find function
type List struct {
	start  int
	end    int
	number string
}

//Find shall be used to find a phone number inside the s (string)
func Find(s string) []List {
	var res []List
	for {
		index := strings.IndexAny(s, "+0")
		end := index
		phonenum := ""

		phoneNumberLike := s[index : index+15]
		phonenum, actualPhoneNum, actuallength := readPhoneNumber(phoneNumberLike)
		end += actuallength

		res = append(res, List{
			start:  index,
			end:    0,
			number: phonenum,
		})

		//remove phone number found
		s = strings.Replace(s, actualPhoneNum, "******", 1)

		break
	}
	return res
}

func readPhoneNumber(s string) (string, string, int) {
	res := ""
	realNumber := s
	idx := 0
	totalLen := len(s)
	for idx < len(s) {
		char, _ := utf8.DecodeRune([]byte{s[idx]})
		if unicode.IsDigit(char) || char == 43 {
			res += string(char)
		}

		if !unicode.IsDigit(char) && char != 32 && char != 43 && char != 45 {
			fmt.Println("invalid char found:", string(char))
			s = strings.Replace(realNumber, string(char), "", 1)
			totalLen--
		}

		idx++

	}
	return res, realNumber, totalLen
}
