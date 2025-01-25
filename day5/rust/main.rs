fn test_action() {
	println!("test")
}

fn test_action1(value:String) {
	println!("{}", value)
}

fn test_action2(value:i32) -> i32 {
	value
}

fn main() {
	test_action();
	test_action1(String::from("test"));
	println!("{}", test_action2(1));
}

