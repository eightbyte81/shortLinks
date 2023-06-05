package shorturl

import (
	"crypto/sha256"
	"encoding/base64"
	"math/big"
)

type IShortUrl interface {
	Sha256Of(input string) []byte
	Base64Encoded(bytes []byte) string
	Generate(link string) string
}

type ShortUrl struct {
	IShortUrl
}

func NewShortUrl() *ShortUrl {
	return &ShortUrl{}
}

func (s *ShortUrl) Sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))

	return algorithm.Sum(nil)
}

func (s *ShortUrl) Base64Encoded(bytes []byte) string {
	encoding := base64.URLEncoding
	return encoding.EncodeToString(bytes)
}

func (s *ShortUrl) Generate(link string) string {
	urlHashBytes := s.Sha256Of(link)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Bytes()
	finalString := s.Base64Encoded(generatedNumber)

	return "localhost:8080/" + finalString[:10]
}
