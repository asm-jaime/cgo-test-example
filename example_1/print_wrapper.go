package main

/*
#include <stdlib.h>
#include "print.h"
*/
import "C"

func Print(str string) {
	cstr := C.CString(str)
	C.print(cstr)
}
