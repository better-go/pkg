package signature

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/better-go/pkg/convert"
	"github.com/better-go/pkg/random"
	"github.com/better-go/pkg/time"
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

		v := convert.StringsDictToDict(item)
		sign := signer.Sign(v, publicKey, nonce, ts, SignTypeMD5, false,
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
		v := convert.StringsDictToDict(item)
		ok := signer.Verify(v, SignTypeMD5, false,
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
			"sign":       []string{""},
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
			"sign":       []string{""},
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
			"sign":       []string{"e982271cde77017384482e7389c211f609d5b124ca311f874efb10aa7b6fa0df"},
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
			"sign":       []string{"c4f41a3cc87e9c21026f4b22f67d2a62c21befa2731d9ec6d9c63a1d69209b35"},
		},
	}

	for _, item := range in {
		v := convert.StringsDictToDict(item)
		sign := signer.Sign(v, publicKey, nonce, ts, SignTypeSHA256, false,
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
		v := convert.StringsDictToDict(item)
		ok := signer.Verify(v, SignTypeSHA256, false,
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
			"sign":       []string{""},
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
			"sign":       []string{""},
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
			"sign":       []string{"986e4653ac747d9d5f4daf32b9b5c947210c7a964f1c6d06a7e3c41e142d4908c4fb511c9427d2440bf73a9133f3e86af49c6fa7ec08318646c6689415eb015b"},
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
			"sign":       []string{"0d65094fff94490b4054d859dc82a4b28cbeabeb403a2d911552272cacc9b481c81cf247d80279f1d19807ce157807573506a916e1b13f6ab949cc98f8850f4a"},
		},
	}

	for _, item := range in {
		v := convert.StringsDictToDict(item)
		sign := signer.Sign(v, publicKey, nonce, ts, SignTypeSHA512, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
			func(data string, privateKey string) (digest string) {
				return WithSHA512(data, privateKey)
			},
		)

		sign2 := signer.SignSHA512(v, publicKey, nonce, ts, false, func(publicKey string) (privateKey string, err error) {
			return secretKey, nil
		})

		t.Logf("sign: %v, item =%+v", sign, item)
		t.Logf("sign2: %v, item =%+v", sign2, item)

	}

	// Sign:  will change item (del sign field)
	for _, item := range expect {
		v := convert.StringsDictToDict(item)
		ok := signer.Verify(v, SignTypeSHA512, false,
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

func TestSignSHA256(t *testing.T) {

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
			"sign":       []string{""},
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
			"sign":       []string{""},
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
			"sign":       []string{"e982271cde77017384482e7389c211f609d5b124ca311f874efb10aa7b6fa0df"},
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
			"sign":       []string{"c4f41a3cc87e9c21026f4b22f67d2a62c21befa2731d9ec6d9c63a1d69209b35"},
		},
	}

	for _, item := range in {
		v := convert.StringsDictToDict(item)
		sign := SignSHA256(v, publicKey, nonce, ts, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
		)
		t.Logf("sign: %v, item =%+v", sign, item)

	}

	// Sign:  will change item (del sign field)
	for _, item := range expect {
		v := convert.StringsDictToDict(item)
		ok := VerifySHA256(v, false,
			func(publicKey string) (privateKey string, err error) {
				return secretKey, nil
			},
		)
		t.Logf("verify: %v", ok)
	}
}
