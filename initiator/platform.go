package initiator

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"pharma-backend/platform"
	"pharma-backend/platform/logger"
	"pharma-backend/platform/token"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type PlatformLayer struct {
	Token platform.Token
}

func InitPlatformLayer(logger logger.Logger, privateKeyPath, publicKeyPath string) PlatformLayer {
	return PlatformLayer{
		Token: token.JWTInit(logger.Named("token-platform"), privateKey(privateKeyPath), publicKey(publicKeyPath)),
	}
}

func privateKey(privateKeyPath string) *rsa.PrivateKey {
	keyFile, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(context.Background(), "failed to read private key", zap.Error(err))
	}

	// privateKey, err := l.ParsePK(keyFile)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyFile)
	if err != nil {
		log.Fatal(context.Background(), "failed to parse private key", zap.Error(err))
	}
	return privateKey
}
func publicKey(publicKeyPath string) *rsa.PublicKey {
	certificate, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal(context.Background(), "Error reading own certificate : \n", zap.Error(err))
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(certificate)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(context.Background(), "Error parsing own certificate : \n", zap.Error(err))
	}
	return pubKey
}
