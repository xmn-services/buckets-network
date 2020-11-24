package miners

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter     hash.Adapter
	blockRepository blocks.Repository
	linkRepository  links.Repository
}

func createApplication(
	hashAdapter hash.Adapter,
	blockRepository blocks.Repository,
	linkRepository links.Repository,
) Application {
	out := application{
		hashAdapter:     hashAdapter,
		blockRepository: blockRepository,
		linkRepository:  linkRepository,
	}

	return &out
}

// Test tests the miner
func (app *application) Test(difficulty uint) (string, error) {
	data := strconv.Itoa(time.Now().UTC().Nanosecond())
	hsh, err := app.hashAdapter.FromBytes([]byte(data))

	if err != nil {
		return "", err
	}

	return app.mine(defaultMiningValue, difficulty, *hsh)
}

// Block executes the block miner
func (app *application) Block(blockHashStr string) (string, error) {
	blockHash, err := app.hashAdapter.FromString(blockHashStr)
	if err != nil {
		return "", err
	}

	block, err := app.blockRepository.Retrieve(*blockHash)
	if err != nil {
		return "", err
	}

	// calculate the difficulty:
	gen := block.Genesis()
	diff := gen.Difficulty()
	blockDiff := diff.Block()
	amountBuckets := uint(len(block.Buckets())) + block.Additional()
	difficulty := app.blockDifficulty(
		blockDiff.Base(),
		blockDiff.IncreasePerBucket(),
		amountBuckets,
	)

	if difficulty > maxDifficulty {
		str := fmt.Sprintf("the block cannot be mined because the required difficulty (%d) to mine it is higher than the maximum difficulty (%d), try to reduce the amount of buckets (%d) in your block in order to reduce its difficulty", difficulty, maxDifficulty, amountBuckets)
		return "", errors.New(str)
	}

	// mine the block:
	miningValue := gen.MiningValue()
	return app.mine(miningValue, difficulty, block.Hash())
}

// Link executes the link miner
func (app *application) Link(linkHashStr string) (string, error) {
	linkHash, err := app.hashAdapter.FromString(linkHashStr)
	if err != nil {
		return "", err
	}

	link, err := app.linkRepository.Retrieve(*linkHash)
	if err != nil {
		return "", err
	}

	gen := link.Next().Block().Genesis()
	miningValue := gen.MiningValue()
	linkDifficulty := gen.Difficulty().Link()
	return app.mine(miningValue, linkDifficulty, link.Hash())
}

func (app *application) blockDifficulty(baseDifficulty uint, increasePerBucket float64, amountBuckets uint) uint {
	sum := float64(0)
	base := float64(baseDifficulty)
	for i := 0; i < int(amountBuckets); i++ {
		sum += increasePerBucket
	}

	return uint(sum + base)
}

func (app *application) mine(
	miningValue uint8,
	difficulty uint,
	hsh hash.Hash,
) (string, error) {
	// create the requested prefix:
	requestedPrefix, err := app.prefix(miningValue, difficulty)
	if err != nil {
		return "", err
	}

	// execute the mining:
	return app.mineRecursively(
		requestedPrefix,
		hsh,
		"",
	)
}

func (app *application) mineRecursively(
	requestedPrefix string,
	hsh hash.Hash,
	baseStr string,
) (string, error) {
	str := ""
	for i := uint(0); i <= maxMiningValue; i++ {
		str = fmt.Sprintf("%s%s", baseStr, []byte(strconv.Itoa(int(i))))
		res, err := app.hashAdapter.FromMultiBytes([][]byte{
			[]byte(str),
			hsh.Bytes(),
		})

		if err != nil {
			return "", err
		}

		if strings.HasPrefix(res.String(), requestedPrefix) {
			return str, nil
		}
	}

	for i := 0; i < maxMiningTries; i++ {
		results, err := app.mineRecursively(requestedPrefix, hsh, str)
		if err != nil {
			continue
		}

		return results, nil
	}

	return "", errors.New("the mining was impossible")
}

func (app *application) prefix(miningValue uint8, difficulty uint) (string, error) {
	output := ""
	for i := 0; i < int(difficulty); i++ {
		output = fmt.Sprintf("%s%d", output, miningValue)
	}

	return output, nil
}
