package utils

import (
	"strings"
	"unicode/utf8"
)

func BuildHasNext(limit, lengthData int64) bool {
	if lengthData < limit {
		return false
	}

	return true
}

func BuildPage(page int64) int64 {
	if page > 0 {
		return page - 1
	}

	return page
}

var tableConvertCases = map[rune]rune{
	'ấ':  'a',
	'ầ':  'a',
	'ẩ':  'a',
	'ẫ':  'a',
	'ậ':  'a',
	'Ấ':  'a',
	'Ầ':  'a',
	'Ẩ':  'a',
	'Ẫ':  'a',
	'Ậ':  'a',
	'ắ':  'a',
	'ằ':  'a',
	'ẳ':  'a',
	'ẵ':  'a',
	'ặ':  'a',
	'Ắ':  'a',
	'Ằ':  'a',
	'Ẳ':  'a',
	'Ẵ':  'a',
	'Ặ':  'a',
	'á':  'a',
	'à':  'a',
	'ả':  'a',
	'ã':  'a',
	'ạ':  'a',
	'â':  'a',
	'ă':  'a',
	'Á':  'a',
	'À':  'a',
	'Ả':  'a',
	'Ã':  'a',
	'Ạ':  'a',
	'Â':  'a',
	'Ă':  'a',
	'ế':  'e',
	'ề':  'e',
	'ể':  'e',
	'ễ':  'e',
	'ệ':  'e',
	'Ế':  'e',
	'Ề':  'e',
	'Ể':  'e',
	'Ễ':  'e',
	'Ệ':  'e',
	'é':  'e',
	'è':  'e',
	'ẻ':  'e',
	'ẽ':  'e',
	'ẹ':  'e',
	'ê':  'e',
	'É':  'e',
	'È':  'e',
	'Ẻ':  'e',
	'Ẽ':  'e',
	'Ẹ':  'e',
	'Ê':  'e',
	'í':  'i',
	'ì':  'i',
	'ỉ':  'i',
	'ĩ':  'i',
	'ị':  'i',
	'Í':  'i',
	'Ì':  'i',
	'Ỉ':  'i',
	'Ĩ':  'i',
	'Ị':  'i',
	'ố':  'o',
	'ồ':  'o',
	'ổ':  'o',
	'ỗ':  'o',
	'ộ':  'o',
	'Ố':  'o',
	'Ồ':  'o',
	'Ổ':  'o',
	'Ộ':  'o',
	'ớ':  'o',
	'ờ':  'o',
	'ở':  'o',
	'ỡ':  'o',
	'ợ':  'o',
	'Ớ':  'o',
	'Ờ':  'o',
	'Ở':  'o',
	'Ỡ':  'o',
	'Ợ':  'o',
	'ó':  'o',
	'ò':  'o',
	'ỏ':  'o',
	'õ':  'o',
	'ọ':  'o',
	'ô':  'o',
	'ơ':  'o',
	'Ó':  'o',
	'Ò':  'o',
	'Ỏ':  'o',
	'Õ':  'o',
	'Ọ':  'o',
	'Ô':  'o',
	'Ơ':  'o',
	'ứ':  'u',
	'ừ':  'u',
	'ử':  'u',
	'ữ':  'u',
	'ự':  'u',
	'Ứ':  'u',
	'Ừ':  'u',
	'Ử':  'u',
	'Ữ':  'u',
	'Ự':  'u',
	'ú':  'u',
	'ù':  'u',
	'ủ':  'u',
	'ũ':  'u',
	'ụ':  'u',
	'ư':  'u',
	'Ú':  'u',
	'Ù':  'u',
	'Ủ':  'u',
	'Ũ':  'u',
	'Ụ':  'u',
	'Ư':  'u',
	'ý':  'y',
	'ỳ':  'y',
	'ỷ':  'y',
	'ỹ':  'y',
	'ỵ':  'y',
	'Ý':  'y',
	'Ỳ':  'y',
	'Ỷ':  'y',
	'Ỹ':  'y',
	'Ỵ':  'y',
	'đ':  'd',
	'Đ':  'd',
	'&':  ' ',
	'<':  ' ',
	'>':  ' ',
	'-':  ' ',
	'ç':  'c',
	'~':  ' ',
	'¥':  'y',
	'©':  'c',
	'®':  'r',
	'\t': ' ',
	'_':  ' ',
	'+':  ' ',
	',':  ' ',
	'*':  ' ',
	'\\': ' ',
	'^':  ' ',
	'$':  ' ',
	'|':  ' ',
	'?':  ' ',
	'[':  ' ',
	']':  ' ',
	'{':  ' ',
	'}':  ' ',
	'(':  ' ',
	')':  ' ',
}

func RemoveAccents(s string) string {
	s1 := ""
	s = strings.ToLower(s)
	for _, c := range s {
		if v, ok := tableConvertCases[c]; ok {
			s1 += string(v)
		} else {
			s1 += string(c)
		}
	}
	return strings.TrimSpace(RemoveDuplicatedWhiteSpace(s1))
}

func RemoveDuplicatedWhiteSpace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	before := utf8.MaxRune
	for _, c := range str {
		if before == ' ' && c == ' ' {
			continue
		}
		b.WriteRune(c)
		before = c
	}
	return b.String()
}
