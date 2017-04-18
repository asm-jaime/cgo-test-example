package lib

/*
#include "lib.h"
#include <stdio.h>
#include <stdlib.h>
#cgo LDFLAGS: -L. -llib

void test_print(void) {
	printf("test print");
}
void call_func_print(void) {
	printf("start calling");
	call_func(print);
}
*/
import "C"

func CallFuncPrint() {
	C.test_print()
	C.call_func_print()
}
