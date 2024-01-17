package output

// interface_adapters/output/service.go

//package output

// CodeService is an interface that defines the contract for the code service
type CodeService interface {
	Encrypt(code string, key []byte) (string, error) // encrypts a code with a key
	Decrypt(code string, key []byte) (string, error) // decrypts a code with a key
}
