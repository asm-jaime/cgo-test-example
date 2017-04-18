#include "lib.h"

#include <stdlib.h>
#include <stdio.h>


void print(int i) {
	printf("print(%d)\n", i);
}

void call_func(void (*f)(int)) {
	int i;
	for (i = 0; i < 3; i++) {
		f(i);
	}
}
