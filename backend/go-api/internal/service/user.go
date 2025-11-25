package service

import (
	"errors"
	"go-api/internal/repository"
)

// UserService はユーザーに関するビジネスロジックを扱うインターフェース
type UserService interface {
	CreateUser(name string, email string) (*repository.User, error)
	GetUserByID(id uint) (*repository.User, error)
	GetAllUsers() ([]repository.User, error)
	UpdateUser(id uint, name string, email string) (*repository.User, error)
	DeleteUser(id uint) error
}

// userService は UserService の実装
type userService struct {
	repo repository.UserRepository
}

// NewUserService は UserService の新しいインスタンスを作成
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser はユーザーを作成するビジネスロジック
func (s *userService) CreateUser(name string, email string) (*repository.User, error) {
	// ここにバリデーションや複雑なビジネスルールを記述
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &repository.User{Name: name, Email: email}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID はユーザーを取得するビジネスロジック
func (s *userService) GetUserByID(id uint) (*repository.User, error) {
	return s.repo.FindByID(id)
}

// GetAllUsers は全ユーザーを取得するビジネスロジック
func (s *userService) GetAllUsers() ([]repository.User, error) {
	return s.repo.FindAll()
}

// UpdateUser はユーザーを更新するビジネスロジック
func (s *userService) UpdateUser(id uint, name string, email string) (*repository.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err // ユーザーが見つからない
	}

	user.Name = name
	user.Email = email

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser はユーザーを削除するビジネスロジック
func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
