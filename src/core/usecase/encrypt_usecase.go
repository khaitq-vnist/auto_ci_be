package usecase

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/core/properties"
	"io"
)

type IEncryptUseCase interface {
	EncryptToken(ctx context.Context, token string) (string, error)
	DecryptToken(ctx context.Context, token string) (string, error)
}
type EncryptUseCase struct {
	props *properties.EncryptProperties
}

func (e *EncryptUseCase) DecryptToken(ctx context.Context, token string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		log.Error(ctx, "Decrypt token error: %v", err)
		return "", err
	}
	block, err := aes.NewCipher([]byte(e.props.Key))
	if err != nil {
		log.Error(ctx, "Decrypt token error: %v", err)
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		log.Error(ctx, "Decrypt token error: cipherText too short")
		return "", err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func (e *EncryptUseCase) EncryptToken(ctx context.Context, token string) (string, error) {
	block, err := aes.NewCipher([]byte(e.props.Key))
	if err != nil {
		log.Error(ctx, "Encrypt token error: %v", err)
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(token))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		log.Error(ctx, "Encrypt token error: %v", err)
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(token))
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func NewEncryptUseCase(props *properties.EncryptProperties) IEncryptUseCase {
	return &EncryptUseCase{
		props: props,
	}
}
