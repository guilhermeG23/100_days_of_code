#include <stdio.h>

void test_action() {
	printf("Hello World!\n");
}

int test_action1() {
	return 1;
}

int test_action2(int value) {
	return value;
}

void main() {
	test_action();
	int value = test_action1();
	printf("%d\n", value);
	value = test_action2(1);
	printf("%d\n", value);
}
