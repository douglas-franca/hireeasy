CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  name VARCHAR(255) NOT NULL,

  -- Login
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,

  -- Email
  email_verification_token_hash TEXT,
  email_verification_expires_at TIMESTAMPTZ,
  email_verification_sent_at TIMESTAMPTZ,

  email_verified BOOLEAN NOT NULL DEFAULT FALSE,
  email_verified_at TIMESTAMPTZ,

  -- Password
  password_reset_token_hash TEXT,
  password_reset_expires_at TIMESTAMPTZ,
  password_reset_sent_at TIMESTAMPTZ,

  -- Timestamps
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
