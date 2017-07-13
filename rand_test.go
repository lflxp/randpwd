package randpwd

import (
	"fmt"
	"testing"
)

func TestGetEasyPassword(t *testing.T) {
	tmp := GetEasyPassword(10)
	if tmp != "" {
		t.Log(fmt.Sprintf("ok : %s", tmp))
	} else {
		t.Fatal(fmt.Sprintf("fail : None"))
	}
}

func TestGetHardPassword(t *testing.T) {
	tmp := GetHardPassword(10)
	if tmp != "" {
		t.Log(fmt.Sprintf("ok : %s", tmp))
	} else {
		t.Fatal(fmt.Sprintf("fail : None"))
	}
}

func BenchGetEasyPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := GetEasyPassword(10)
		if tmp != "" {
			b.Log(tmp)
		} else {
			b.Fatal("fail")
		}
	}
}

func BenchGetHardPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := GetHardPassword(10)
		if tmp != "" {
			b.Log(tmp)
		} else {
			b.Fatal("fail")
		}
	}
}
