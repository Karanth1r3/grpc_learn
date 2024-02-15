package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Karanth1r3/grpc_learn/internal/domain/models"
	"github.com/Karanth1r3/grpc_learn/internal/storage"
	"github.com/Karanth1r3/grpc_learn/internal/util/logger/slg"
	"golang.org/x/crypto/bcrypt"
)

// Authentification service layer
type (
	Auth struct {
		log          *slog.Logger
		userSaver    UserSaver
		userProvider UserProvider
		appProvider  AppProvider
		tokenTTL     time.Duration
	}
)

// "Cut" interface of storage declared at usage location
type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

// Same
type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

var (
	ErrInvalidCridentials = errors.New("invalid cridentials")
)

// New - CTOR, returns new instance of the Auth service
func New(
	log *slog.Logger, userSaver UserSaver, userProvider UserProvider, appProvider AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		userSaver:    userSaver,
		userProvider: userProvider,
		log:          log,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

// Login tries to login the user with provided account details
func (a *Auth) Login(ctx context.Context, email string, password string, appID int) (string, error) {
	const op = "auth.Login()"

	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email),
	)

	log.Info("attempting to login user")

	user, err := a.userProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", slg.Err(err))

			// Do not expose info about users in negative response
			return "", fmt.Errorf("%s : %s", op, ErrInvalidCridentials)
		}

		a.log.Error("failed to get user", slg.Err(err))

		return "", fmt.Errorf("%s : %w", op, err)
	}

	// Checking salted & hashed password through bcrypt
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Info("invalid cridentials", slg.Err(err))

		return "", fmt.Errorf("%s : %w", op, ErrInvalidCridentials)
	}

	// Getting app id from request && trying to create token for individual app
	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		return "", fmt.Errorf("%s : %w", op, err)
	}

	log.Info("user logged in successfully")
}

// RegisterNewUser registers new user in the system and returns user ID
func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (int64, error) {
	const op = "auth.RegisterNewUser()"

	// Dangerous to expose emails in logs
	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")

	// Additional protection layer => "salting" the password (with the hash)
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", slg.Err(err))

		return 0, fmt.Errorf("%s : %w", op, err)
	}

	// Trying to save user to the storage through interface
	id, err := a.userSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error("failed to save user", slg.Err(err))

		return 0, fmt.Errorf("%s : %w", op, err)
	}

	log.Info("user successfully registered")

	return id, nil
}

// IsAdmin checks if user is admin
func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	panic("not implemented")
}
