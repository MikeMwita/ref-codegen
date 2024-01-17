package output

// interface_adapters/output/response.go

//package output

// Response is a type that defines the output data structures for the interface adapters layer
type Response struct {
	Value  string // the value of the code or key
	Format string // the format of the code or key, such as Luhn, UUID, KSUID, etc.
	Valid  bool   // the validation result of the code or key
}
