package example

import (
	"testing"
)

func Test_ADD_1(t *testing.T) {
	if 8 == Add(2, 6) {
		t.Log("通过")
	} else {
		t.Error("未通过1 数值不对")
	}
}

func Test_ADD_2(t *testing.T) {
	t.Log("没问题 过~")
}
