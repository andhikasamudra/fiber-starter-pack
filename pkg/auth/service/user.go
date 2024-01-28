package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	"github.com/andhikasamudra/fiber-starter-pack/internal/env"
	"github.com/andhikasamudra/fiber-starter-pack/internal/middleware"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/dto"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/repo"
	"github.com/andhikasamudra/fiber-starter-pack/provider/mail"
	"github.com/andhikasamudra/fiber-starter-pack/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Dependency struct {
	AuthModel    repo.Interface
	RedisClient  *redis.Client
	MailProvider mail.Interface
	DbAdapter    postgres.Interface
	Logger       *zap.Logger
}

type AuthService struct {
	AuthModel    repo.Interface
	RedisClient  *redis.Client
	MailProvider mail.Interface
	DbAdapter    postgres.Interface
	Logger       *zap.Logger
}

func NewService(d Dependency) *AuthService {
	return &AuthService{
		AuthModel:    d.AuthModel,
		RedisClient:  d.RedisClient,
		MailProvider: d.MailProvider,
		Logger:       d.Logger,
		DbAdapter:    d.DbAdapter,
	}
}

func (s *AuthService) RegisterUser(ctx *fiber.Ctx, request dto.CreateUserRequest) error {
	user := repo.User{
		Email: request.Email,
	}

	err := s.DbAdapter.BeginTransaction()
	if err != nil {
		return err
	}

	newUser, err := s.AuthModel.CreateUser(ctx.Context(), s.DbAdapter, user)
	if err != nil {
		return err
	}

	_, err = s.AuthModel.CreateUserRole(ctx.Context(), s.DbAdapter, repo.UserRole{
		UserID:  newUser.ID,
		RoleKey: request.Role,
	})
	if err != nil {
		return err
	}

	profile, err := s.AuthModel.CreateUserProfile(ctx.Context(), s.DbAdapter, repo.UserProfile{
		UserID:      newUser.ID,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		CompanyName: request.CompanyName,
	})
	if err != nil {
		return err
	}

	err = s.DbAdapter.Commit()
	if err != nil {
		return err
	}

	go func() {
		go func(email string, name string) {
			token, _ := utils.GenerateToken(64)
			activatePath := fmt.Sprintf("activate?code=%s", token)
			err := s.RedisClient.Set(context.Background(), fmt.Sprintf("active_email_%s", token), email, 24*time.Hour).Err()
			if err != nil {
				s.Logger.Error(fmt.Sprintf("failed to store data to redis : %s", err))
			}
			link := fmt.Sprintf("%s/%s", env.BaseURL(), activatePath)
			emailData := map[string]interface{}{
				"Name": name,
				"Link": link,
			}
			parsedHTML, err := utils.ParsedHTMLTemplate("./template/activate_email.html", emailData)
			if err != nil {
				s.Logger.Error(fmt.Sprintf("failed to parsed email template: %s", err))
			}

			err = s.MailProvider.Send(mail.SendMailRequest{
				To:      []string{email},
				Message: parsedHTML,
				Subject: "Activate Email",
			})
			if err != nil {
				s.Logger.Error(fmt.Sprintf("failed to send email activate: %s", err))
			}
		}(newUser.Email, profile.Name)
	}()

	return nil
}

func (s *AuthService) ActivateUser(ctx *fiber.Ctx, request dto.ActivateUserRequest) error {
	redisKey := fmt.Sprintf("active_email_%s", request.Code)
	val, _ := s.RedisClient.Get(ctx.Context(), redisKey).Result()
	if val == "" {
		return fmt.Errorf("invalid activation code")
	}

	userData, err := s.AuthModel.ReadUser(ctx.Context(), s.DbAdapter, repo.GetUserRequest{
		Email: val,
	})
	if err != nil {
		return err
	}

	userData.IsActive = true
	err = s.AuthModel.UpdateUser(ctx.Context(), s.DbAdapter, *userData, []string{"is_active"})
	if err != nil {
		return err
	}

	go func() {
		err := s.RedisClient.Del(context.Background(), redisKey).Err()
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (s *AuthService) OTPRequest(ctx *fiber.Ctx, request dto.LoginRequest) error {
	userData, err := s.AuthModel.ReadUser(ctx.Context(), s.DbAdapter, repo.GetUserRequest{
		Email: request.Email,
	})
	if err != nil {
		return err
	}

	if !userData.IsActive {
		return fmt.Errorf("user need to be activated")
	}

	otpCode, err := utils.EncodeToString(6)
	if err != nil {
		return err
	}

	go func(otpCode string) {
		err := s.RedisClient.Set(context.Background(), fmt.Sprintf("otp_code_%s", request.Email), otpCode, 1*time.Minute).Err()
		if err != nil {
			s.Logger.Error(fmt.Sprintf("failed to store data to redis : %s", err))
		}

		emailData := map[string]interface{}{
			"name":     userData.Profile.Name,
			"otp_code": otpCode,
		}
		parsedHTML, err := utils.ParsedHTMLTemplate("./template/otp_email.html", emailData)
		if err != nil {
			s.Logger.Error(fmt.Sprintf("failed to parsed email template: %s", err))
		}

		err = s.MailProvider.Send(mail.SendMailRequest{
			To:      []string{userData.Email},
			Message: parsedHTML,
			Subject: "OTP Code",
		})
		if err != nil {
			s.Logger.Error(fmt.Sprintf("failed to send email : %s", err))
		}
	}(otpCode)

	return nil
}

func (s *AuthService) LoginWithOTPRequest(ctx *fiber.Ctx, request dto.LoginWithOTPRequest) (*dto.LoginResponse, error) {
	redisKey := fmt.Sprintf("otp_code_%s", request.Email)
	val, _ := s.RedisClient.Get(ctx.Context(), redisKey).Result()
	if val == "" {
		return nil, fmt.Errorf("invalid request")
	}

	if request.OTPCode != val {
		return nil, fmt.Errorf("invalid otp code")
	}

	userData, err := s.AuthModel.ReadUser(ctx.Context(), s.DbAdapter, repo.GetUserRequest{
		Email: request.Email,
	})
	if err != nil {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Email
	claims["guid"] = userData.GUID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	resultToken, err := token.SignedString([]byte(env.JWTSecret()))
	if err != nil {
		return nil, err
	}

	go func() {
		err := s.RedisClient.Del(context.Background(), redisKey).Err()
		if err != nil {
			log.Println(err)
		}
	}()

	return &dto.LoginResponse{
		AccessToken: resultToken,
	}, nil
}

func (s *AuthService) GetProfile(ctx *fiber.Ctx) (*dto.GetProfileResponse, error) {
	claims := middleware.GetClaims(ctx)
	userData, err := s.AuthModel.ReadUser(ctx.Context(), s.DbAdapter, repo.GetUserRequest{
		GUID: uuid.MustParse(claims["guid"].(string)),
	})
	if err != nil {
		return nil, err
	}

	var roles []string
	for _, i := range userData.Roles {
		roles = append(roles, i.RoleKey)
	}

	return &dto.GetProfileResponse{
		Email:       userData.Email,
		GUID:        userData.GUID,
		PICName:     userData.Profile.Name,
		CompanyName: userData.Profile.CompanyName,
		PhoneNumber: userData.Profile.PhoneNumber,
		Roles:       roles,
	}, nil
}
