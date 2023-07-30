package handler

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestConvertBytesToArray(t *testing.T) {
	var arr = []byte(`a,b,c,
	d,e,`)
	ret := convertBytesToArray(arr)

	expectRet := [][]string{
		{"a", "b", "c"},
		{"d", "e"},
	}

	if len(ret[0]) != len(expectRet[0]) {
		t.Errorf("Should array length equal: %d", len(expectRet[0]))
	}
	if len(ret[1]) != len(expectRet[1]) {
		t.Errorf("Should array length equal: %d", len(expectRet[1]))
	}
}

func TestGetSheet(t *testing.T) {
	godotenv.Load("../.env")
	getSheet()
}
