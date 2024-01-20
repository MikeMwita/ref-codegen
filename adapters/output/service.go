package output

type CodeService interface {
	EncryptCode(code string, key []byte) (string, error)
	DecryptCode(code string, key []byte) (string, error)
}
