module example/hello

go 1.17

replace ravi-greeting.com/greetings => ../greetings

replace slice-greetings.com/slice_greetings => ../slice_greetings

replace multi-greetings.com/mul_greetings => ../multi-greetings

require (
	multi-greetings.com/multi_greetings v0.0.0-00010101000000-000000000000
	ravi-greeting.com/greetings v0.0.0-00010101000000-000000000000
	slice-greetings.com/slice_greetings v0.0.0-00010101000000-000000000000
)

replace multi-greetings.com/multi_greetings => ../multi-greetings
