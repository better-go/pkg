package signature

import (
	"fmt"
	"github.com/better-go/pkg/random"
	"github.com/better-go/pkg/time"
	"net/url"
	"testing"
)

func TestSigner_SignMD5(t *testing.T) {
	signer := New("public_key", "nonce", "ts")

	publicKey := "this-is-public-key"
	secretKey := "use-this-do-sign"

	nonce := fmt.Sprintf("%v", random.SnowFlakeID())
	ts := time.Gen10BitTimestamp()
	// set:
	nonce = "705935809601077250"
	ts = "1616053176"

	in := []url.Values{
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"jim"},
			"age":        []string{"22"},
			"sex":        []string{"boy"},
			"address":    []string{"beijing"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"b06103ff6cdcb5660f065eecd508398f"},
		},
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"kate"},
			"age":        []string{"22"},
			"sex":        []string{"girl"},
			"address":    []string{"shanghai"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"95d483f5eae66707965dfecf312c6cdf"},
		},
	}

	expect := []url.Values{
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"jim"},
			"age":        []string{"22"},
			"sex":        []string{"boy"},
			"address":    []string{"beijing"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"b06103ff6cdcb5660f065eecd508398f"},
		},
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"kate"},
			"age":        []string{"22"},
			"sex":        []string{"girl"},
			"address":    []string{"shanghai"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"95d483f5eae66707965dfecf312c6cdf"},
		},
	}

	for _, item := range in {
		sign := signer.Sign(item, publicKey, nonce, ts, SignTypeMD5, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
			func(data string, privateKey string) (digest string) {
				return WithMD5(data, privateKey)
			},
		)
		t.Logf("sign: %v, item =%+v", sign, item)

	}

	// Sign:  will change item (del sign field)
	for _, item := range expect {
		ok := signer.Verify(item, SignTypeMD5, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
			func(data string, privateKey string) (digest string) {
				return WithMD5(data, privateKey)
			},
		)
		t.Logf("verify: %v", ok)
	}
}

func TestSigner_SignSHA256(t *testing.T) {
	signer := New("public_key", "nonce", "ts")

	publicKey := "this-is-public-key"
	secretKey := "use-this-do-sign"

	nonce := fmt.Sprintf("%v", random.SnowFlakeID())
	ts := time.Gen10BitTimestamp()
	// set:
	nonce = "705935809601077250"
	ts = "1616053176"

	in := []url.Values{
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"jim"},
			"age":        []string{"22"},
			"sex":        []string{"boy"},
			"address":    []string{"beijing"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"a2f3b946f9cc4a885275687c0fa4bdfe01f8a9e9c56e31969ae547b60af3f19f"},
		},
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"kate"},
			"age":        []string{"22"},
			"sex":        []string{"girl"},
			"address":    []string{"shanghai"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"64255506677740d621bab1a1f1e36d409d51923cf1d6881ae706d839ce9081f7"},
		},
	}

	expect := []url.Values{
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"jim"},
			"age":        []string{"22"},
			"sex":        []string{"boy"},
			"address":    []string{"beijing"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"a2f3b946f9cc4a885275687c0fa4bdfe01f8a9e9c56e31969ae547b60af3f19f"},
		},
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"kate"},
			"age":        []string{"22"},
			"sex":        []string{"girl"},
			"address":    []string{"shanghai"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"64255506677740d621bab1a1f1e36d409d51923cf1d6881ae706d839ce9081f7"},
		},
	}

	for _, item := range in {
		sign := signer.Sign(item, publicKey, nonce, ts, SignTypeSHA256, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
			func(data string, privateKey string) (digest string) {
				return WithSHA256(data, privateKey)
			},
		)
		t.Logf("sign: %v, item =%+v", sign, item)

	}

	// Sign:  will change item (del sign field)
	for _, item := range expect {
		ok := signer.Verify(item, SignTypeSHA256, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
			func(data string, privateKey string) (digest string) {
				return withSHA256v2(data, privateKey)
			},
		)
		t.Logf("verify: %v", ok)
	}
}

func TestSigner_SignSHA512(t *testing.T) {
	signer := New("public_key", "nonce", "ts")

	publicKey := "this-is-public-key"
	secretKey := "use-this-do-sign"

	nonce := fmt.Sprintf("%v", random.SnowFlakeID())
	ts := time.Gen10BitTimestamp()
	// set:
	nonce = "705935809601077250"
	ts = "1616053176"

	in := []url.Values{
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"jim"},
			"age":        []string{"22"},
			"sex":        []string{"boy"},
			"address":    []string{"beijing"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"91dacd0f5bad8f6c17c9dc3e632c0dcede777018b2b34a6f56f109c6ca5e8b299a05b571ffd4da3ca65c8bf7b4de624fd0fcbe27ba6bea06dcf8b97bc0bc6ebc"},
		},
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"kate"},
			"age":        []string{"22"},
			"sex":        []string{"girl"},
			"address":    []string{"shanghai"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"6e484eb7924ab926eb5197b02339f7db50bc74eadc5618fb627c5127f7e1eeea3380de50d0a8f3ad878e8ce08b967bf044be4398a10d86961335404022bab959"},
		},
	}

	expect := []url.Values{
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"jim"},
			"age":        []string{"22"},
			"sex":        []string{"boy"},
			"address":    []string{"beijing"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"91dacd0f5bad8f6c17c9dc3e632c0dcede777018b2b34a6f56f109c6ca5e8b299a05b571ffd4da3ca65c8bf7b4de624fd0fcbe27ba6bea06dcf8b97bc0bc6ebc"},
		},
		{
			"key1":       []string{"a1"},
			"key2":       []string{"b1"},
			"name":       []string{"kate"},
			"age":        []string{"22"},
			"sex":        []string{"girl"},
			"address":    []string{"shanghai"},
			"public_key": []string{publicKey},
			"nonce":      []string{nonce},
			"ts":         []string{ts},
			"sign":       []string{"6e484eb7924ab926eb5197b02339f7db50bc74eadc5618fb627c5127f7e1eeea3380de50d0a8f3ad878e8ce08b967bf044be4398a10d86961335404022bab959"},
		},
	}

	for _, item := range in {
		sign := signer.Sign(item, publicKey, nonce, ts, SignTypeSHA512, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
			func(data string, privateKey string) (digest string) {
				return WithSHA512(data, privateKey)
			},
		)

		sign2 := signer.SignSHA512(item, publicKey, nonce, ts, false, func(publicKey string) (privateKey string, err error) {
			return secretKey, nil
		})

		t.Logf("sign: %v, item =%+v", sign, item)
		t.Logf("sign2: %v, item =%+v", sign2, item)

	}

	// Sign:  will change item (del sign field)
	for _, item := range expect {
		ok := signer.Verify(item, SignTypeSHA512, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
			func(data string, privateKey string) (digest string) {
				return withSHA512v2(data, privateKey)
			},
		)

		//ok2 := signer.VerifySHA512(item, false, func(publicKey string) (privateKey string, err error) {
		//	return secretKey, nil
		//})

		t.Logf("verify: %v, %v", ok, ok)
	}
}
