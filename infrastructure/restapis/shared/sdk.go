package shared

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// PathKeyname represents the path keyname in the posted data
const PathKeyname = "path"

// TokenHeadKeyname represents the token keyname in the headers
const TokenHeadKeyname = "X-Session-Token"

// AuthenticateToURLValues converts an authenticate to url values
func AuthenticateToURLValues(auth Authenticate) url.Values {
	values := url.Values{}
	values.Set("name", auth.Name)
	values.Set("password", auth.Password)
	values.Set("seed", auth.Seed)
	return values
}

// URLValuesToAuthenticate converts a url values to Authenticate instance
func URLValuesToAuthenticate(urlValues url.Values) (*Authenticate, error) {
	name := urlValues.Get("name")
	if name == "" {
		return nil, errors.New("the name is mandatory in order to create an Authenticate instance")
	}

	password := urlValues.Get("password")
	if password == "" {
		return nil, errors.New("the password is mandatory in order to create an Authenticate instance")
	}

	seed := urlValues.Get("seed")
	if seed == "" {
		return nil, errors.New("the seed is mandatory in order to create an Authenticate instance")
	}

	return &Authenticate{
		Name:     name,
		Password: password,
		Seed:     seed,
	}, nil
}

// Base64ToAuthenticate converts base64 string to Authenticate instance
func Base64ToAuthenticate(token string) (*Authenticate, error) {
	js, err := b64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	ins := new(Authenticate)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// AuthenticateToBase64 converts an authenticate instance to base64
func AuthenticateToBase64(auth *Authenticate) (string, error) {
	js, err := json.Marshal(auth)
	if err != nil {
		return "", nil
	}

	return b64.StdEncoding.EncodeToString(js), nil
}

// URLValuesToIdentity converts a url values to Identity instance
func URLValuesToIdentity(urlValues url.Values) (*Identity, error) {
	auth, err := URLValuesToAuthenticate(urlValues)
	if err != nil {
		return nil, nil
	}

	root := urlValues.Get("root")
	if root == "" {
		return nil, errors.New("the root is mandatory in order to create an Identity instance")
	}

	return &Identity{
		Authenticate: auth,
		Root:         root,
	}, nil
}

// IdentityToURLValues converts an identity to url values
func IdentityToURLValues(identity Identity) url.Values {
	values := AuthenticateToURLValues(*identity.Authenticate)
	values.Set("root", identity.Root)
	return values
}

// URLValuesToInitChain converts a url values to InitChain instance
func URLValuesToInitChain(urlValues url.Values) (*InitChain, error) {
	miningValueStr := urlValues.Get("mining_value")
	if miningValueStr == "" {
		return nil, errors.New("the miningValue is mandatory in order to create an InitChain instance")
	}

	baseDifficultyStr := urlValues.Get("base_difficulty")
	if baseDifficultyStr == "" {
		return nil, errors.New("the baseDifficulty is mandatory in order to create an InitChain instance")
	}

	increasePerBucketStr := urlValues.Get("increase_per_bucket")
	if increasePerBucketStr == "" {
		return nil, errors.New("the increasePerBucket is mandatory in order to create an InitChain instance")
	}

	linkDifficultyStr := urlValues.Get("link_difficulty")
	if linkDifficultyStr == "" {
		return nil, errors.New("the linkDifficulty is mandatory in order to create an InitChain instance")
	}

	rootAdditionalBucketsStr := urlValues.Get("root_additional_buckets")
	if rootAdditionalBucketsStr == "" {
		return nil, errors.New("the rootAdditionalBuckets is mandatory in order to create an InitChain instance")
	}

	headAdditionalBucketsStr := urlValues.Get("head_additional_buckets")
	if headAdditionalBucketsStr == "" {
		return nil, errors.New("the headAdditionalBuckets is mandatory in order to create an InitChain instance")
	}

	miningValue, err := strconv.Atoi(miningValueStr)
	if err != nil {
		return nil, err
	}

	baseDifficulty, err := strconv.Atoi(baseDifficultyStr)
	if err != nil {
		return nil, err
	}

	increasePerBucket, err := strconv.ParseFloat(increasePerBucketStr, 64)
	if err != nil {
		return nil, err
	}

	linkDifficulty, err := strconv.Atoi(linkDifficultyStr)
	if err != nil {
		return nil, err
	}

	rootAdditionalBuckets, err := strconv.Atoi(rootAdditionalBucketsStr)
	if err != nil {
		return nil, err
	}

	headAdditionalBuckets, err := strconv.Atoi(headAdditionalBucketsStr)
	if err != nil {
		return nil, err
	}

	return &InitChain{
		MiningValue:           uint8(miningValue),
		BaseDifficulty:        uint(baseDifficulty),
		IncreasePerBucket:     increasePerBucket,
		LinkDifficulty:        uint(linkDifficulty),
		RootAdditionalBuckets: uint(rootAdditionalBuckets),
		HeadAdditionalBuckets: uint(headAdditionalBuckets),
	}, nil
}

// InitChainToURLValues converts an InitChain instance to url values
func InitChainToURLValues(initChain *InitChain) url.Values {
	incrPerBucketStr := strconv.FormatFloat(initChain.IncreasePerBucket, 'f', 10, 64)

	values := url.Values{}
	values.Set("mining_value", strconv.Itoa(int(initChain.MiningValue)))
	values.Set("base_difficulty", strconv.Itoa(int(initChain.BaseDifficulty)))
	values.Set("increase_per_bucket", incrPerBucketStr)
	values.Set("link_difficulty", strconv.Itoa(int(initChain.LinkDifficulty)))
	values.Set("root_additional_buckets", strconv.Itoa(int(initChain.RootAdditionalBuckets)))
	values.Set("head_additional_buckets", strconv.Itoa(int(initChain.HeadAdditionalBuckets)))
	return values
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
