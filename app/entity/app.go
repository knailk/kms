package entity

import (
	"encoding/hex"
	"fmt"
	"kms/app/errs"
	"kms/app/secure"
	"time"

	"github.com/google/uuid"
)

// APIKeyGenerator creates a random, 128 API key string
type APIKeyGenerator interface {
	RandomString(n int) (string, error)
}

// App is an application that interacts with the system
type App struct {
	ID               uuid.UUID
	ExternalID       secure.Identifier
	Org              *Org
	Name             string
	Description      string
	Provider         Provider
	ProviderClientID string
	APIKeys          []APIKey
}

// AddKey validates and adds an API key to the slice of App API keys
func (a *App) AddKey(key APIKey) error {
	const op errs.Op = "diygoapi/App.AddKey"

	err := key.validate()
	if err != nil {
		return errs.E(op, errs.Internal, err)
	}
	a.APIKeys = append(a.APIKeys, key)

	return nil
}

// ValidateKey determines if the app has a matching key for the input
// and if that key is valid
func (a *App) ValidateKey(realm, matchKey string) error {
	const op errs.Op = "diygoapi/App.ValidateKey"

	key, err := a.matchKey(realm, matchKey)
	if err != nil {
		return err
	}
	err = key.validate()
	if err != nil {
		return errs.E(op, errs.Unauthenticated, errs.Realm(realm), err)
	}
	return nil
}

// MatchKey returns the matching Key given the string, if exists.
// An error will be sent if no match is found.
func (a *App) matchKey(realm, matchKey string) (APIKey, error) {
	const op errs.Op = "diygoapi/App.matchKey"

	for _, apiKey := range a.APIKeys {
		if matchKey == apiKey.Key() {
			return apiKey, nil
		}
	}
	return APIKey{}, errs.E(op, errs.Unauthenticated, errs.Realm(realm), "Key does not match any keys for the App")
}


// APIKey is an API key for interacting with the system. The API key string
// is delivered to the client along with an App ID. The API Key acts as a
// password for the application.
type APIKey struct {
	// key: the unencrypted API key string
	key string
	// ciphertext: the encrypted API key as []byte
	ciphertextbytes []byte
	// deactivation: the date/time the API key is no longer usable
	deactivation time.Time
}

// NewAPIKey initializes an APIKey. It generates a random 128-bit (16 byte)
// base64 encoded string as an API key. The generated key is then encrypted
// using 256-bit AES-GCM and the encrypted bytes are added to the struct as
// well.
func NewAPIKey(g APIKeyGenerator, ek *[32]byte, deactivation time.Time) (APIKey, error) {
	const (
		n  int = 16
		op     = "diygoapi/NewAPIKey"
	)
	var (
		k   string
		err error
	)
	k, err = g.RandomString(n)
	if err != nil {
		return APIKey{}, errs.E(op, err)
	}

	var ctb []byte
	ctb, err = secure.Encrypt([]byte(k), ek)
	if err != nil {
		return APIKey{}, err
	}

	return APIKey{key: k, ciphertextbytes: ctb, deactivation: deactivation}, nil
}

// NewAPIKeyFromCipher initializes an APIKey given a ciphertext string.
func NewAPIKeyFromCipher(ciphertext string, ek *[32]byte) (APIKey, error) {
	const op errs.Op = "diygoapi/NewAPIKeyFromCipher"

	var (
		eak []byte
		err error
	)

	// encrypted api key is stored using hex in db. Decode to get ciphertext bytes.
	eak, err = hex.DecodeString(ciphertext)
	if err != nil {
		return APIKey{}, errs.E(op, errs.Internal, err)
	}

	var apiKey []byte
	apiKey, err = secure.Decrypt(eak, ek)
	if err != nil {
		return APIKey{}, errs.E(op, err)
	}

	return APIKey{key: string(apiKey), ciphertextbytes: eak}, nil
}

// Key returns the key for the API key
func (a *APIKey) Key() string {
	return a.key
}

// Ciphertext returns the hex encoded text of the encrypted cipher bytes for the API key
func (a *APIKey) Ciphertext() string {
	return hex.EncodeToString(a.ciphertextbytes)
}

// DeactivationDate returns the Deactivation Date for the API key
func (a *APIKey) DeactivationDate() time.Time {
	return a.deactivation
}

// SetDeactivationDate sets the deactivation date value to AppAPIkey
// TODO - try SetDeactivationDate as a candidate for generics with 1.18
func (a *APIKey) SetDeactivationDate(t time.Time) {
	a.deactivation = t
}

// SetStringAsDeactivationDate sets the deactivation date value to
// AppAPIkey given a string in RFC3339 format
func (a *APIKey) SetStringAsDeactivationDate(s string) error {
	const op errs.Op = "diygoapi/APIKey.SetStringAsDeactivationDate"

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return errs.E(op, errs.Validation, err)
	}
	a.deactivation = t

	return nil
}

func (a *APIKey) validate() error {
	const op errs.Op = "diygoapi/APIKey.validate"

	if a.ciphertextbytes == nil {
		return errs.E(op, "ciphertext must have a value")
	}

	now := time.Now()
	if a.deactivation.Before(now) {
		return errs.E(op, fmt.Sprintf("Key Deactivation %s is before current time %s", a.deactivation.String(), now.String()))
	}
	return nil
}
