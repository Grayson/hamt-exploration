#include <ctype.h>
#include <stdio.h>

int ascii_hash(char c) {
	return tolower(c) - 'a';
}

int main() {
	return !(0 != printf("HERE!\n"));
}