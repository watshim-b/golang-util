package exstring

import "strings"

// この関数は、引数で指定された文字列の中に、2つ以上連続する空白があった場合に、空白が1つになるまで変換を行う。
func TrimMultipleSpace(s string) string {
	for {
		// 2つ以上連続するスペースがなくなった場合は、ループを抜ける
		if !strings.Contains(s, "  ") {
			break
		}
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}
