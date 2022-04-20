package exstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// この関数は、引数で指定された文字列の中に、2つ以上連続する空白があった場合に、空白が1つになるまで変換を行う。
func TestTrimMultipleSpace(t *testing.T) {

	// 空白が１つだけの文字列のテスト
	case1Str := " t e s t "
	case1StrRes := TrimMultipleSpace(case1Str)
	assert.Equal(t, case1Str, case1StrRes)

	// 連続する空白が先頭に存在する文字列のテスト
	case2Str := "  t e s t "
	case2StrRes := TrimMultipleSpace(case2Str)
	assert.NotEqual(t, case2Str, case2StrRes)

	// 連続する空白が文字列中に存在する文字列のテスト
	case3Str := " t  e  s t "
	case3StrRes := TrimMultipleSpace(case3Str)
	assert.NotEqual(t, case3Str, case3StrRes)

	// 連続する空白が文末に存在する文字列のテスト
	case4Str := " t e s t  "
	case4StrRes := TrimMultipleSpace(case4Str)
	assert.NotEqual(t, case4Str, case4StrRes)

	// 連続する空白が3つ以上存在する文字列に対して関数を実行した場合に、「  」の文字列が含まれない事
	case5Str := " t e                                                                                   s t "
	case5StrRes := TrimMultipleSpace(case5Str)
	assert.NotContains(t, case5StrRes, "  ")
}
