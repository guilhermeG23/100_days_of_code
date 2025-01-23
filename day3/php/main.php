<?php
// Variables
// Variables in PHP are mixed, they can assume any state at any time;
$value = 'test'; // String
$value = 1; // Int
$value = 1.0; // Float
$value = true; // Bool
$value = ['test']; // Array -> Key/value

class TestFunction {
	public $input_value;

	function __construct($value = '') {
		$this->input_value = $value;
	}
}

$value = new TestFunction(); // Object
$value = null; // Null

foreach ([1,2,3] as $value) { // iterable
	echo $value;
}

function hello_world() {
	echo 'Hello World';
}

function request_function(callable $callback) {
	$callback();
}

request_function('hello_world'); // Callback

// Constant
// Value that cannot be changed
define(VALUE, 'test');
