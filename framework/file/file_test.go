package file

import (
	"testing"
)

// 以下是功能测试

func Test_GetCurrentDirectory_1(t *testing.T) {
	thisDir := GetCurrentDirectory()
	t.Log(thisDir)
}

func Test_GetFileExt_1(t *testing.T) {
	fn := "/download/vagrant_1.9.2.dmg"
	fileExt := GetFileExt(fn)
	if len(fileExt) < 1 {
		t.Error("fileExt null!")
	} else {
		t.Log(fileExt)
	}
}

func Test_GetFileExt_2(t *testing.T) {
	fn := "/download/vagrant_1"
	fileExt := GetFileExt(fn)
	if len(fileExt) < 1 {
		t.Error("fileExt null!")
	} else {
		t.Log(fileExt)
	}
}

func Test_Exist_1(t *testing.T) {}

// 以下是性能测试