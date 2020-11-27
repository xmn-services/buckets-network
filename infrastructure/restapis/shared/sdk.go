package shared

import "net/url"

// URLValuesToAuthenticate converts a url values to Authenticate instance
func URLValuesToAuthenticate(urlValues url.Values) (*Authenticate, error) {
	return nil, nil
}

// Base64ToAuthenticate converts base64 string to Authenticate instance
func Base64ToAuthenticate(token string) (*Authenticate, error) {
	return nil, nil
}

// AuthenticateToBase64 converts an authenticate instance to base64
func AuthenticateToBase64(auth *Authenticate) string {
	return ""
}

// URLValuesToIdentity converts a url values to Identity instance
func URLValuesToIdentity(urlValues url.Values) (*Identity, error) {
	return nil, nil
}

// URLValuesToInitChain converts a url values to InitChain instance
func URLValuesToInitChain(urlValues url.Values) (*InitChain, error) {
	return nil, nil
}

// Authenticate represents an authenticate instance
type Authenticate struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Seed     string `json:"seed"`
}

// Identity represents an identity instance
type Identity struct {
	Authenticate *Authenticate `json:"authenticate"`
	Root         string        `json:"root"`
}

// InitChain represents an init chain instance
type InitChain struct {
	MiningValue           uint8   `json:"mining_value"`
	BaseDifficulty        uint    `json:"base_difficulty"`
	IncreasePerBucket     float64 `json:"increase_per_bucket"`
	LinkDifficulty        uint    `json:"link_difficulty"`
	RootAdditionalBuckets uint    `json:"root_additional_buckets"`
	HeadAdditionalBuckets uint    `json:"head_additional_buckets"`
}
