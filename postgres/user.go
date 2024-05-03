package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/aliml92/anor"
	"github.com/aliml92/anor/postgres/store/user"
)

// Ensure service implements interface.
var _ anor.UserService = (*UserService)(nil)

type UserService struct {
	userStore user.Store
}

func NewUserService(us user.Store) *UserService {
	return &UserService{
		userStore: us,
	}
}

func (s UserService) CreateUser(ctx context.Context, u anor.User) error {
	err := s.userStore.ExecTx(ctx, func(tx pgx.Tx) error {
		err := s.userStore.WithTx(tx).CreateUser(ctx, u.Email, u.Password, u.FullName, &u.OTP, &u.OTPExpiry)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return anor.ErrUserExists
			}
		}

		return err
	}

	return nil
}

func (s UserService) GetUser(ctx context.Context, id int64) (anor.User, error) {
	ru, err := s.userStore.GetUser(ctx, id) // retrieved user
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return anor.User{}, anor.ErrNotFound
		}
		return anor.User{}, err
	}

	u := anor.User{
		ID:        ru.ID,
		Email:     ru.Email,
		Password:  ru.Password,
		FullName:  ru.FullName,
		Status:    anor.UserStatus(ru.Status),
		CreatedAt: ru.CreatedAt.Time,
		UpdatedAt: ru.UpdatedAt.Time,
	}

	if ru.PhoneNumber != nil {
		u.PhoneNumber = *ru.PhoneNumber
	}
	if ru.Otp != nil {
		u.OTP = *ru.Otp
	}
	if ru.Otp != nil {
		u.OTPExpiry = *ru.OtpExpiry
	}

	return u, nil
}

func (s UserService) GetUserByEmail(ctx context.Context, email string) (anor.User, error) {
	ru, err := s.userStore.GetUserByEmail(ctx, email) // retrieved user
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return anor.User{}, anor.ErrNotFound
		}
		return anor.User{}, err
	}

	u := anor.User{
		ID:        ru.ID,
		Email:     ru.Email,
		Password:  ru.Password,
		FullName:  ru.FullName,
		Status:    anor.UserStatus(ru.Status),
		CreatedAt: ru.CreatedAt.Time,
		UpdatedAt: ru.UpdatedAt.Time,
	}

	if ru.PhoneNumber != nil {
		u.PhoneNumber = *ru.PhoneNumber
	}
	if ru.Otp != nil {
		u.OTP = *ru.Otp
	}
	if ru.OtpExpiry != nil {
		u.OTPExpiry = *ru.OtpExpiry
	}

	return u, nil
}

func (s UserService) UpdateUserStatus(ctx context.Context, status anor.UserStatus, id int64) error {
	err := s.userStore.UpdateUserStatus(ctx, user.UserStatus(status), id)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) UpdateUserOTP(ctx context.Context, id int64, otp string, otpExpiry int64) error {
	err := s.userStore.UpdateUserOTP(ctx, id, &otp, &otpExpiry)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) UpdateUserPassword(ctx context.Context, id int64, password string) error {
	err := s.userStore.UpdateUserPassword(ctx, password, id)
	if err != nil {
		return err
	}
	return nil
}
