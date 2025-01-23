#include <stdio.h>
#include <stdbool.h>

void main() {
	// Char / String
	char value1 = 't';
	char value2[] = "test"; // Dynamic

	int value3 = 1; // 4 bytes
	short value4 = 1; // 2 bytes
	long value5 = 1; // 4 / 8 bytes
	long long value6 = 30; // 8 bytes

	// unsigned -> Does not allow the variable to have negative numbers

	float value7 = 10.00; // 4 bytes
	double value8 = 10.00; // 8 bytes
	long double value9 = 10.00; // 8/12/16 bytes

	bool value10 = true;

	int value11[5] = {1, 2, 3, 4, 5}; // Fixed size array

	char *value12 = &value1; // Pointer

	printf("%c\n", value1);
	printf("%s\n", value2);
	printf("%d\n", value3);
	printf("%d\n", value4);
	printf("%ld\n", value5);
	printf("%lld\n", value6);
	printf("%f\n", value7);
	printf("%.2f\n", value7); // Limit float return
	printf("%f\n", value8);
	printf("%Lf\n", value9);
	printf("%d\n", value10);
	printf("%d\n", value11[0]);
	printf("%c\n", *value12);
}
