package chains

import (
	"net/url"
	"strconv"
)

type adapter struct {
	builder Builder
}

func createAdapter(builder Builder) Adapter {
	out := adapter{
		builder: builder,
	}

	return &out
}

// URLValuesToChain converts url values to Chain instance
func (app *adapter) URLValuesToChain(values url.Values) (Chain, error) {
	builder := app.builder.Create()
	miningValueStr := values.Get("mining_value")
	if miningValueStr != "" {
		miningValue, err := strconv.Atoi(miningValueStr)
		if err != nil {
			return nil, err
		}

		builder.WithMiningValue(uint8(miningValue))
	}

	baseDifficultyStr := values.Get("base_difficulty")
	if baseDifficultyStr != "" {
		baseDifficulty, err := strconv.Atoi(baseDifficultyStr)
		if err != nil {
			return nil, err
		}

		builder.WithBaseDifficulty(uint(baseDifficulty))
	}

	increasePerBucketStr := values.Get("increase_per_bucket")
	if increasePerBucketStr != "" {
		increasePerBucket, err := strconv.ParseFloat(increasePerBucketStr, 64)
		if err != nil {
			return nil, err
		}

		builder.WithIncreasePerBucket(increasePerBucket)
	}

	linkDifficultyStr := values.Get("link_difficulty")
	if linkDifficultyStr != "" {
		linkDifficulty, err := strconv.Atoi(linkDifficultyStr)
		if err != nil {
			return nil, err
		}

		builder.WithLinkDifficulty(uint(linkDifficulty))
	}

	rootAdditionalBucketsStr := values.Get("root_additional_buckets")
	if rootAdditionalBucketsStr != "" {
		rootAdditional, err := strconv.Atoi(rootAdditionalBucketsStr)
		if err != nil {
			return nil, err
		}

		builder.WithRootAdditionalBuckets(uint(rootAdditional))
	}

	headAdditionalBucketsStr := values.Get("head_additional_buckets")
	if headAdditionalBucketsStr != "" {
		headAdditional, err := strconv.Atoi(headAdditionalBucketsStr)
		if err != nil {
			return nil, err
		}

		builder.WithHeadAdditionalBuckets(uint(headAdditional))
	}

	return builder.Now()
}

// ChainToURLValues converts a Chain instance to url values
func (app *adapter) ChainToURLValues(chain Chain) url.Values {
	values := url.Values{}

	miningValue := chain.MiningValue()
	values.Add("mining_value", strconv.Itoa(int(miningValue)))

	baseDifficulty := chain.BaseDifficulty()
	values.Add("base_difficulty", strconv.Itoa(int(baseDifficulty)))

	increasePerBucket := chain.IncreasePerBucket()
	values.Add("increase_per_bucket", strconv.FormatFloat(increasePerBucket, 'f', 10, 64))

	linkDifficulty := chain.LinkDifficulty()
	values.Add("link_difficulty", strconv.Itoa(int(linkDifficulty)))

	rootAdditional := chain.RootAdditionalBuckets()
	values.Add("root_additional_buckets", strconv.Itoa(int(rootAdditional)))

	headAdditional := chain.HeadAdditionalBuckets()
	values.Add("head_additional_buckets", strconv.Itoa(int(headAdditional)))

	return values
}
