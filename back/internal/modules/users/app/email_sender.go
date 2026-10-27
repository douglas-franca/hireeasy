package app

import "context"

type VerificationEmailSender interface {
	SendVerification(context.Context, string, string) error
}
