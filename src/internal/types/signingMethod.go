package types

type SigningMethod string

const (
	ES256   = "es256"
	ED25519 = "ed25519"
)

func (s SigningMethod) String() string {
	switch s {
	case ES256:
		return "es256"
	case ED25519:
		return "ed25519"
	}
	return "unknown"
}
