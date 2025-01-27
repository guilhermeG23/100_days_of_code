# Variables without scope
value='t1'

func1() {
	value='t2'
	echo $value
}

echo $value
func1
echo $value

echo '-----------------------------------'

# Scoped variables
value1='t1'

func2() {
	local value1='t2'
	echo $value1
}

echo $value1
func2
echo $value1
