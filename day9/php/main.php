<?php

class Test {
	//public $value1;
	//public $value2;

	private $value1;
	private $value2;

	function __construct($value1, $value2) {
		$this->value1 = $value1;
		$this->value2 = $value2;
	}

	public function setValues($value1, $value2): void {
		$this->value1 = $value1;
		$this->value2 = $value2;
	}

	public function getValues(): void {
		echo $this->value1 . " --- " . $this->value2 . PHP_EOL;
	}

}

echo "---------------------------------" . PHP_EOL;

$test = new Test('t1', 't2');

if (isset($test->value1)) { echo "È possível chamar um elemento privado" . PHP_EOL; }

$test->getValues();
$test->setValues('t2', 't3');
$test->getValues();

echo "---------------------------------" . PHP_EOL;
