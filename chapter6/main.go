package main

// #cgo LDFLAGS: -L. -lcgetcstring
// #include <stdlib.h>
// char* getCString(const char* str);
import "C"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	cStr := C.CString("Go String to C")

	cResult := C.getCString(cStr)

	goStr := C.GoString(cResult) // C 문자열을 Go 문자열로 변환
	fmt.Println("결과:", goStr)

	C.free(unsafe.Pointer(cStr))
	C.free(unsafe.Pointer(cResult))
}
