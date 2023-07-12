#include <math.h>
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

array array_resize(array const orig, int const size) {
	if (size < orig.size) {
		array tmp = {
			orig.size,
		};
		/* ignored */ memcpy(tmp.ptr, orig.ptr, orig.size * sizeof(char));
		return tmp;
	}

	// find next size
	int increment = 1;
	int nextSize = pow(2, orig.size+increment);
	while (increment++, nextSize < size) {
		nextSize = pow(2, increment);
	}

	array tmp = {
		nextSize,
		calloc(nextSize, sizeof(char)),
	};
	/* ignored */ memcpy(tmp.ptr, orig.ptr, orig.size * sizeof(char));
	return tmp;
}

array array_insert(array const orig, char const value, int const index) {
	array tmp = array_resize(orig, index);
	tmp.ptr[index] = value;
	return tmp;
}
