#include <stdlib.h>

typedef struct {
	int size;
	char * const ptr;
} array;

typedef struct {
	int value;
} array_size;

array array_allocate(int size) {
	char * const ptr = calloc(sizeof(char), size);
	array tmp = {
		size,
		ptr,
	};
	return tmp;
}
