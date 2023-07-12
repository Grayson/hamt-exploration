#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>

#include "array.c"

int ascii_hash(char c) {
	return tolower(c) - 'a';
}

struct test
{
	char * const name;
	bool (*fn)(void);
};

bool testHash(void) {
	return ascii_hash('a') == 0 && ascii_hash('b') == 1 && ascii_hash('C') == 2;
}

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

struct test tests[] = {
	{ "testHash", testHash },
	{ "arrayAllocate", testArrayAllocate },
	{ "arrayResize", testArrayResize },
	{ "arrayInsertResize", testArrayInsertResize },
};

int main() {
	int const testCount = sizeof(tests) / sizeof(struct test);
	int numberOfFailingTests = 0;
	for (int testIndex = 0; testIndex < testCount; testIndex++) {
		struct test * const t = &tests[testIndex];
		printf("%d: %s ", testIndex, t->name);
		bool const result = t->fn();
		numberOfFailingTests += result == false;
		printf("-> %s\n", result ? "passed" : "failed");
	}
	if (numberOfFailingTests > 0) {
		printf("%d failing tests\n", numberOfFailingTests);
		return 1;
	}
	return 0;
}
