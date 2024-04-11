package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

type SHFILEOPSTRUCT struct {
	hwnd   uintptr
	wFunc  uint32
	pFrom  *uint16
	pTo    *uint16
	fFlags uint32
}

const (
	FO_DELETE          = 3
	FOF_ALLOWUNDO      = 0x0040
	FOF_NOCONFIRMATION = 0x0010
)

var shell32 = syscall.NewLazyDLL("shell32.dll")
var procSHFileOperation = shell32.NewProc("SHFileOperationW")

func SHFileOperation(op *SHFILEOPSTRUCT) int {
	rc, _, _ := procSHFileOperation.Call(uintptr(unsafe.Pointer(op)))
	return int(rc)
}
func DeleteToBin(s string) bool {
	uu, _ := syscall.UTF16PtrFromString(s)
	op := &SHFILEOPSTRUCT{
		hwnd:   0,
		wFunc:  FO_DELETE,
		pFrom:  uu,
		pTo:    nil,
		fFlags: FOF_ALLOWUNDO | FOF_NOCONFIRMATION,
	}
	ret := SHFileOperation(op)
	if ret != 0 {
		err := syscall.Errno(ret)
		fmt.Println("error deleting file:", err)
	} else {
		fmt.Println("file deleted successfully")
		return true
	}
	return false
}
func main() {
	usr, _ := os.LookupEnv("USRPROFILE")

	fmt.Println(usr)
	u := os.Getenv("USERPROFILE")
	h := os.Getenv("HOMEPATH")
	fmt.Println(u)
	fmt.Println(h)
	a := "./a.txt"
	ab, _ := filepath.Abs(a)
	b := DeleteToBin(ab)
	if b {
		fmt.Println("成功！")
	} else {
		fmt.Println("失败！")
	}

}
