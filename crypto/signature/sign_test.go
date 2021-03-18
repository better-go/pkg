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
		sign := signer.SignMD5(item, publicKey, nonce, ts, false, func(publicKey string) (privateKey string, err error) {
			return secretKey, nil
		})
		t.Logf("sign: %v, item =%+v", sign, item)

	}

	// SignMD5:  will change item (del sign field)
	for _, item := range expect {
		ok := signer.VerifyMD5(item, false, func(publicKey string) (privateKey string, err error) {
			return secretKey, nil
		})
		t.Logf("verify: %v", ok)
	}

}
