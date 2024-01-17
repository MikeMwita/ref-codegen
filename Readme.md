
# CodeGen-Ref: A Package for Generating Reference Codes and Keys

CodeGen is a package for generating reference codes and keys for various purposes, such as transactions, API keys, tokens, etc. It provides a simple and flexible way to create and manage codes and keys with different formats, lengths, characters, prefixes, suffixes, and delimiters. It also supports different algorithms and methods for generating codes and keys, such as Luhn, UUID, KSUID, etc. It also offers validation and verification functions to check the validity and integrity of the codes and keys, and encryption and decryption functions to secure and protect them.

## Installation

To install CodeGen, you need to have Go installed on your system. You can download it from [here](^1^). Then, you can use the following command to get the package:

```go
go get github.com/yourname/codegen
```

## Usage

To use CodeGen, you need to import it in your Go code:

```go
import "github.com/MikeMwita/codegen"
```

Then, you can use the functions and types provided by the package to generate and manage codes and keys. Here are some examples:

### Generate a 10-digit Luhn code with a "LUHN-" prefix

```go
code, err := codegen.Generate(codegen.Options{
    Length: 10,
    Format: codegen.Luhn,
    Prefix: "LUHN-",
})
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(code) // LUHN-1234567890
```

### Generate a UUID v4 code with a "-" delimiter

```go
code, err := codegen.Generate(codegen.Options{
    Format: codegen.UUID,
    Delimiter: "-",
})
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(code) // 123e4567-e89b-12d3-a456-426614174000
```

### Generate a KSUID code with a "KSUID" suffix

```go
code, err := codegen.Generate(codegen.Options{
    Format: codegen.KSUID,
    Suffix: "KSUID",
})
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(code) // 0ujsszwN8NRY24YaXiTIE2RlCKSUID
```

### Validate a Luhn code

```go
valid := codegen.Validate(codegen.Options{
    Format: codegen.Luhn,
}, "LUHN-1234567890")
fmt.Println(valid) // true
```

### Encrypt a code with AES

```go
key := []byte("your-secret-key")
encrypted, err := codegen.Encrypt(codegen.Options{
    Format: codegen.AES,
}, "your-code", key)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(encrypted) // some encrypted string
```

### Decrypt a code with AES

```go
key := []byte("your-secret-key")
decrypted, err := codegen.Decrypt(codegen.Options{
    Format: codegen.AES,
}, "some encrypted string", key)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(decrypted) // your-code
```
### Generate a private and public key pair in PEM format
    
```go
privateKey, publicKey, err := codegen.GenerateKeyPair(codegen.Options{
Format: codegen.RSA,
Length: 2048,
})
if err != nil {
fmt.Println(err)
return
}
fmt.Println(privateKey) // -----BEGIN RSA PRIVATE KEY----- ...
fmt.Println(publicKey) // -----BEGIN PUBLIC KEY
```

