package main

//#cgo CFLAGS: -I${SRCDIR}/callClib
//#cgo LDFLAGS: ${SRCDIR}/callC.a
//#include <stdlib.h>
//#include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function")
	C.cHello()
	myMessage := C.CString("This is Meeee!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfecly done!")
}
