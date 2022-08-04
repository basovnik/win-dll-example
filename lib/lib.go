package main

import (
	"fmt"
	"unsafe"

	/*
		typedef void (*logHook) (int, wchar_t*);

		static inline void call_c_func(logHook ptr, int intVal, wchar_t *strVal) {
		    (ptr)(intVal, strVal);
		}
	*/
	"C"
	"golang.org/x/sys/windows"
)

func wcharPtrFromString(s string) *C.wchar_t {
	p, _ := windows.UTF16PtrFromString(s)
	return (*C.wchar_t)(unsafe.Pointer(p))
}

//export hello0
func hello0() {
	fmt.Println("Hello, world!")
}

//export hello
func hello(namePtr *uint16) int32 {
	name := windows.UTF16PtrToString(namePtr)
	fmt.Println("Hello", name)
	return 0
}

//export callback
func callback(log C.logHook) int32 {
	C.call_c_func(log, 1, wcharPtrFromString("one"))
	C.call_c_func(log, 2, wcharPtrFromString("two"))
	C.call_c_func(log, 3, wcharPtrFromString("three"))

	return 0
}

func main() {}
