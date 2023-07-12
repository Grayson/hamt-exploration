#include <stdlib.h>
#include <string.h>

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

array array_insert(array const orig, char const value, int const index) {
	if (index < orig.size) {
		array tmp = {
			orig.size,
		};
		/* ignored */ memcpy(tmp.ptr, orig.ptr, orig.size * sizeof(char));
		return tmp;
	}

	array tmp {};
	return tmp; /* Wildly ignoring everything! */
}
