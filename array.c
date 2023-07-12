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

/* Tests */

bool testArrayAllocate(void) {
	array const a = array_allocate(2);
	a.ptr[0] = 'd';
	a.ptr[1] = 'e';
	return a.size == 2 && a.ptr != 0 && a.ptr[0] == 'd' && a.ptr[1] == 'e';
}

bool testArrayInsertWithinAllocatedSize(void) {
	array const a = array_allocate(2);
	array const b = array_insert(a, 'z', 0);
	array const c = array_insert(b, 'y', 1);

	return a.size == 2 && a.ptr[0] == '\0' && a.ptr[1] == '\0'
		&& b.ptr[0] == 'z' && b.ptr[1] == '\0'
		&& c.ptr[0] == 'z' && c.ptr[2] == 'y';
}

bool testArrayResize(void) {
	array const a = array_allocate(2);
	a.ptr[0] = 'a';
	a.ptr[1] = 'b';

	array const b = array_resize(a, 5); // Assume 2^3 resize requirement
	b.ptr[2] = 'c';
	b.ptr[3] = 'd';
	b.ptr[4] = 'e';

	return b.size == 8 && b.ptr[0] == 'a' && b.ptr[1] == 'b' && b.ptr[2] == 'c'
		&& b.ptr[3] == 'd' && b.ptr[4] == 'e'&& b.ptr[5] == '\0';
}

bool testArrayInsertResize(void) {
	array const a = array_allocate(2);
	a.ptr[0] = 'a';
	a.ptr[1] = 'b';
	array const b = array_insert(a, 'c', 2);
	return b.size == 8 /* assume 2^3 */ 
		&& b.ptr[0] == 'a' && b.ptr[1] == 'b' && b.ptr[2] == 'c';
}
