package vault

import (
	"context"
	"errors"

	"github.com/deep-rent/nexus/jose/jwk"
)

// ErrKeyNotFound is returned when a requested key cannot be found in the vault.
var ErrKeyNotFound = errors.New("key not found in vault")

// Vault represents a secure retrieval mechanism for cryptographic
// signing keys ([jwk.KeyPair]). It abstracts away the underlying implementation
// details of external sources like KMS, HSM, or HashiCorp Vault.
type Vault interface {
	jwk.Resolver

	// Next retrieves the currently active [jwk.KeyPair] intended for signing
	// new tokens. It returns [ErrKeyNotFound] if the vault is empty.
	Next() (jwk.KeyPair, error)
}

// Store represents an external backend capable of securely supplying and managing signing keys.
type Store interface {
	// Load retrieves all available signing keys from the store, decrypting them using the provided KEK.
	// The implementation must return the keys such that the currently active signing key is the first element.
	Load(ctx context.Context, kek []byte) ([]jwk.KeyPair, error)

	// Revoke invalidates a key by its Key ID, preventing it from being returned in future Load calls
	// or used for signing/verification.
	Revoke(ctx context.Context, kid string) error

	// Generate creates a new key pair, encrypts it using the provided KEK, and stores it in the backend.
	// This key typically becomes the new active key.
	Generate(ctx context.Context, kek []byte) (jwk.KeyPair, error)

	// Add imports an existing key pair from PEM format, encrypts it using the provided KEK,
	// and stores it in the backend.
	Add(ctx context.Context, kek []byte, pemData []byte) (jwk.KeyPair, error)
}
