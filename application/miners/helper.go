package miners

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/xmn-services/buckets-network/libs/hash"
)

// blockDifficulty calculates the block's difficulty
func blockDifficulty(baseDifficulty uint, increasePerBucket float64, amountBuckets uint) uint {
	sum := float64(0)
	base := float64(baseDifficulty)
	for i := 0; i < int(amountBuckets); i++ {
		sum += increasePerBucket
	}

	return uint(sum + base)
}

// prefix returns the prefix based on the difficulty
func prefix(miningValue uint8, difficulty uint) (string, error) {
	output := ""
	for i := 0; i < int(difficulty); i++ {
		output = fmt.Sprintf("%s%d", output, miningValue)
	}

	return output, nil
}

func mine(
	hashAdapter hash.Adapter,
	miningValue uint8,
	difficulty uint,
	hsh hash.Hash,
) (string, error) {
	// create the requested prefix:
	requestedPrefix, err := prefix(miningValue, difficulty)
	if err != nil {
		return "", err
	}

	// execute the mining:
	return mineRecursively(
		hashAdapter,
		requestedPrefix,
		hsh,
		"",
	)
}

func mineRecursively(
	hashAdapter hash.Adapter,
	requestedPrefix string,
	hsh hash.Hash,
	baseStr string,
) (string, error) {
	str := ""
	for i := uint(0); i <= maxMiningValue; i++ {
		str = fmt.Sprintf("%s%s", baseStr, []byte(strconv.Itoa(int(i))))
		res, err := hashAdapter.FromBytes([]byte(str))
		if err != nil {
			return "", err
		}

		if strings.HasPrefix(res.String(), requestedPrefix) {
			return str, nil
		}
	}

	for i := 0; i < maxMiningTries; i++ {
		results, err := mineRecursively(hashAdapter, requestedPrefix, hsh, str)
		if err != nil {
			continue
		}

		return results, nil
	}

	return "", errors.New("the mining was impossible")
}
