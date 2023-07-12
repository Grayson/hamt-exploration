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

struct test tests[] = {
	{ "testHash", testHash },
	{ "arrayAllocate", testArrayAllocate },
};

int main() {
	int const testCount = sizeof(tests) / sizeof(struct test);
	int numberOfFailingTests = 0;
	for (int testIndex = 0; testIndex < testCount; testIndex++) {
		struct test * const t = &tests[testIndex];
		bool const result = t->fn();
		numberOfFailingTests += result == false;
		printf("%d: %s -> %s\n", testIndex, t->name, result ? "passed" : "failed");
	}
	if (numberOfFailingTests > 0) {
		printf("%d failing tests\n", numberOfFailingTests);
		return 1;
	}
	return 0;
}
