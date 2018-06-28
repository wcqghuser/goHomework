package main

import "testing"

func TestStatistal(t *testing.T) {

	str := ""
	a, b, c, d := statistal(str)
	if a != 0 || b != 0 || c != 0 || d != 0 {
		t.Errorf("%s result should be 字母:%d,空格:%d,数字:%d,其它字符:%d, but is 字母:%d,空格:%d,数字:%d,其它字符:%d,", str, 0, 0, 0, 0, a, b, c, d)
	}

	str = "abcd"
	a, b, c, d = statistal(str)
	if a != 4 || b != 0 || c != 0 || d != 0 {
		t.Errorf("%s result should be 字母:%d,空格:%d,数字:%d,其它字符:%d, but is 字母:%d,空格:%d,数字:%d,其它字符:%d,", str, 0, 0, 0, 0, a, b, c, d)
	}

	str = "    "
	a, b, c, d = statistal(str)
	if a != 0 || b != 4 || c != 0 || d != 0 {
		t.Errorf("%s result should be 字母:%d,空格:%d,数字:%d,其它字符:%d, but is 字母:%d,空格:%d,数字:%d,其它字符:%d,", str, 0, 4, 0, 0, a, b, c, d)
	}

	str = "1234567890"
	a, b, c, d = statistal(str)
	if a != 0 || b != 0 || c != 10 || d != 0 {
		t.Errorf("%s result should be 字母:%d,空格:%d,数字:%d,其它字符:%d, but is 字母:%d,空格:%d,数字:%d,其它字符:%d,", str, 0, 0, 10, 0, a, b, c, d)
	}

	str = "你好世界"
	a, b, c, d = statistal(str)
	if a != 0 || b != 0 || c != 0 || d != 4 {
		t.Errorf("%s result should be 字母:%d,空格:%d,数字:%d,其它字符:%d, but is 字母:%d,空格:%d,数字:%d,其它字符:%d,", str, 0, 0, 0, 4, a, b, c, d)
	}

}
