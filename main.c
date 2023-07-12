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
