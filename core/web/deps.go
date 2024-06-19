package web

import (
	"context"
	"time"
)

type (
	UserStorage interface {
		CreateUser(ctx context.Context, login, email string, passwordHash []byte) (int64, error)
		GetSignInDataByLogin(ctx context.Context, login string) (UserDataToSignIn, error)
		GetSignInDataById(ctx context.Context, id int64) (UserDataToSignIn, error)
		ActivateUser(ctx context.Context, userId int64) error
		UserIdByLoginAndEmail(ctx context.Context, login, email string) (int64, error)
		UpdatePassword(ctx context.Context, id int64, passwordHash []byte) error
	}

	RandomGenerator interface {
		GenerateString(ctx context.Context, set string, size int) (string, error)
	}

	ActivationCodeStorage interface {
		SaveActivationCode(ctx context.Context, code string, userId int64, ttl time.Duration) error
		VerifyActivationCode(ctx context.Context, code string) (int64, error)
	}

	AuthNotifier interface {
		SendActivationCode(ctx context.Context, email, code string) error
		SendRecoveryCode(ctx context.Context, email, code string) error
	}

	PasswordHasher interface {
		HashPassword(ctx context.Context, b []byte, cost int) ([]byte, error)
		CompareHashAndPassword(ctx context.Context, password string, hashedPassword []byte) error
	}

	RefreshTokenStorage interface {
		SaveRefreshToken(ctx context.Context, token string, userId int64, ttl time.Duration) error
		VerifyRefreshToken(ctx context.Context, token string) (int64, error)
	}

	TwoFactorCodeStorage interface {
		Save2FaCode(ctx context.Context, code string, userId int64, ttl time.Duration) error
		Verify2FaCode(ctx context.Context, code string) (int64, error)
	}

	TwoFactorCodeNotifier interface {
		Send2FaCode(ctx context.Context, telegramId int64, code string) error
	}

	RecoveryCodeStorage interface {
		SaveRecoveryCode(ctx context.Context, code string, userId int64, ttl time.Duration) error
		VerifyRecoveryCode(ctx context.Context, code string) (int64, error)
	}
)
