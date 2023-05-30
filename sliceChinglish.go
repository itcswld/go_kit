package main

import (
	"unicode"
)

func SliceChinglish(s string, cn int, cn_en int) (ss []string) {

	for i, v := range s {
		n := 0.0
		if IsChinese(v) {
			n += 1
		} else {
			n += 0.5
		}

		if i > 17 {
			if n == 17 {
				ss = SliceStr(s, cn)
			} else {
				ss = SliceStr(s, cn_en)
			}
		}
	}
	return ss
}

func IsChinese(char rune) (b bool) {
	//chinese
	if unicode.Is(unicode.Han, char) {
		b = true
	} else {
		//not chinese
		b = false
	}
	return b
}

func SliceStr(s string, slice int) (arr_str []string) {
	txtRune := []rune(s)
	for i := 1; i < len(txtRune); i++ {
		if i%slice == 0 {
			//slicing
			arr_str = append(arr_str, string(txtRune[:i])) //[0][1]
			txtRune = txtRune[i:]                          //[i]~[len(s)]
			i = 1
		}
	}
	arr_str = append(arr_str, string(txtRune[:]))
	return arr_str
}
