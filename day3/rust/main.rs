struct TestValue {
	t1: String,
	t2: i8
}

#[allow(dead_code)] // Allows you to turn off warning notices for the section
enum Values {
	T1,
	T2,
	T3
}

fn main() {
	// String/char
	let value:String = String::from("test");
	println!("{}", value);
	let value: &str = "test";
	println!("{}", value);
	let value:char = 't';
	println!("{}", value);
	// Float
	let value:f32 = 1.0;
	println!("{}", value);
	let value:f64 = 1.0;
	println!("{}", value);
	// Int
	// i8, i16, i32, i64, i128: Signed integers (can be negative or positive);
	// u8, u16, u32, u64, u128: Unsigned integers (positive only)
	let value:i8 = 1;
	println!("{}", value);
	// Bool
	let value:bool = true;
	println!("{}", value);
	// Array
	// Arrays in Rust have to have a fixed size for performance reasons because it starts using the dynamic memory heap -> Too expensive
	let value: [i32; 3] = [1, 2, 3];
	println!("{}", value[0]);
	// Struct
	let value = TestValue {
		t1: String::from("t1"),
		t2: 1
	};
	println!("{} {}", value.t1, value.t2);
	// Enum
	let value = Values::T3;
	match value {
		Values::T1 => println!("t1"),
		Values::T2 => println!("t2"),
		Values::T3 => println!("t3"),
	}
	// Ok and Err
	let value = return_ok_err(2);
	match value {
		Ok(value_return) => println!("Result: {}", value_return),
		Err(error_return) => println!("Error: {}", error_return),
	}
}

fn return_ok_err(value: i8) -> Result<i8, String> {
	if value == 1 {
		Err("Error".to_string())
	} else {
		Ok(value)
	}
}
