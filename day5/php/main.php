<?php
function test_action() {
	echo 'Hello World';
}

function test_action1() {
	return 1;
}

function test_action2(): int {
	return 1;
}

function test_action3(int $value): int {
	return $value;
}


function test_action4(int $value = 0): int {
	return $value;
}

test_action();
echo test_action1() . PHP_EOL;
echo test_action2() . PHP_EOL;
echo test_action3(1) . PHP_EOL;
echo test_action4(1) . PHP_EOL;
