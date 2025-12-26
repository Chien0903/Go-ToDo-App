package services

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Chien0903/Go-ToDo-App/internal/config"
	"github.com/Chien0903/Go-ToDo-App/internal/models"
	"github.com/Chien0903/Go-ToDo-App/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUsernameExists     = errors.New("username already exists")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserService struct {
	repo *repository.UserRepository
	cfg  config.AppConfig
}

func NewUserService(repo *repository.UserRepository, cfg config.AppConfig) *UserService {
	return &UserService{repo: repo, cfg: cfg}
}

func (s *UserService) Register(username, password string) (*models.User, error) {
	username = strings.TrimSpace(strings.ToLower(username))
	password = strings.TrimSpace(password)

	// validate cơ bản
	if username == "" || len(password) < 6 {
		return nil, ErrInvalidInput
	}

	existed, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if existed != nil {
		return nil, ErrUsernameExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Username:     username,
		PasswordHash: string(hashed),
	}

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}

// Login xác thực user và trả về JWT token
func (s *UserService) Login(username, password string) (string, *models.User, error) {
	username = strings.TrimSpace(strings.ToLower(username))
	password = strings.TrimSpace(password)

	if username == "" || password == "" {
		return "", nil, ErrInvalidCredentials
	}

	// Tìm user theo username
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", nil, err
	}
	if user == nil {
		return "", nil, ErrInvalidCredentials
	}

	// So sánh password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, ErrInvalidCredentials
	}

	// Tạo JWT token
	expiresMinute, _ := strconv.Atoi(s.cfg.JWTExpiresMinute)
	if expiresMinute <= 0 {
		expiresMinute = 60 // default 60 phút
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Duration(expiresMinute) * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}
