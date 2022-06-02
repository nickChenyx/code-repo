module example.com/import_panic_init

go 1.18

require panic_init v0.0.0

require (
	bou.ke/monkey v1.0.2
	util v0.0.0
)

replace panic_init => ../panic_init

replace util => ../util
