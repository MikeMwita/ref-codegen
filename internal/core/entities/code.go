package entities

import (
	"crypto/rand"
	"fmt"
)

const (
	Luhn  = "Luhn"
	UUID  = "UUID"
	KSUID = "KSUID"
)

type Code struct {
	Value  string
	Format string
}

func NewCode(value, format string) (*Code, error) {
	// validate the value and format
	if value == "" {
		return nil, fmt.Errorf("value cannot be empty")
	}
	if format == "" {
		return nil, fmt.Errorf("format cannot be empty")
	}

	code := &Code{
		Value:  value,
		Format: format,
	}
	return code, nil
}

// Generate generates a new random value for the Code entity according to the format
func (c *Code) Generate(options Options) error {
	var value string
	var err error
	switch c.Format {
	case Luhn:
		value, err = generateLuhn(options)
	case UUID:
		value, err = generateUUID(options)
	case KSUID:
		value, err = generateKSUID(options)
	default:
		return fmt.Errorf("invalid format: %s", c.Format)
	}
	if err != nil {
		return err
	}
	c.Value = value
	return nil
}

// Validate validates the value of the Code entity according to the format

func (c *Code) Validate() (bool, error) {
	var valid bool
	var err error
	switch c.Format {
	case Luhn:
		valid, err = validateLuhn(c.Value)
	case UUID:
		valid, err = validateUUID(c.Value)
	case KSUID:
		valid, err = validateKSUID(c.Value)
	default:
		return false, fmt.Errorf("invalid format: %s", c.Format)
	}
	if err != nil {
		return false, err
	}
	return valid, nil
}

type Options struct {
	Length    int
	Suffix    string
	Delimiter string
	Prefix    string
}

// generateLuhn generates a Luhn code with the given options
func generateLuhn(options Options) (string, error) {
	// TODO: implement the logic for generating a Luhn code
	return "", nil
}

// validateLuhn validates a Luhn code
func validateLuhn(code string) (bool, error) {
	// TODO: implement the logic for validating a Luhn code
	return false, nil
}

// generateUUID generates a UUID code with the given options
func generateUUID(options Options) (string, error) {
	// TODO: implement the logic for generating a UUID code
	return "", nil
}

// validateUUID validates a UUID code
func validateUUID(code string) (bool, error) {
	// TODO: implement the logic for validating a UUID code
	return false, nil
}

// generateKSUID generates a KSUID code with the given options
func generateKSUID(options Options) (string, error) {
	// TODO: implement the logic for generating a KSUID code
	return "", nil
}

// validateKSUID validates a KSUID code
func validateKSUID(code string) (bool, error) {
	// TODO: implement the logic for validating a KSUID code
	return false, nil
}

// generateRandomBytes generates a slice of random bytes with the given length
func generateRandomBytes(length int) ([]byte, error) {
	// create a byte slice with the given length
	b := make([]byte, length)
	// fill the byte slice with random data
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
