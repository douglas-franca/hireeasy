package app

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
)

type VerificationTokenIssuer struct {
	repository domain.VerificationTokenStore
	sender     VerificationEmailSender
	ttl        time.Duration
}

func NewVerificationTokenIssuer(repository domain.VerificationTokenStore, sender VerificationEmailSender, ttl time.Duration) *VerificationTokenIssuer {
	return &VerificationTokenIssuer{repository: repository, sender: sender, ttl: ttl}
}

func (i *VerificationTokenIssuer) Issue(ctx context.Context, user *domain.User) error {
	token, err := newVerificationToken()
	if err != nil {
		return err
	}

	if err := i.repository.SetVerificationToken(ctx, user.ID, hashVerificationToken(token), time.Now().Add(i.ttl)); err != nil {
		return err
	}

	return i.sender.SendVerification(ctx, user.Email, token)
}

func newVerificationToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func hashVerificationToken(token string) string {
	sum := sha256.Sum256([]byte(token))

	return hex.EncodeToString(sum[:])
}
