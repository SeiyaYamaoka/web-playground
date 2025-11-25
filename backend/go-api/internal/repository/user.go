package repository

import "gorm.io/gorm"

// User はデータベースのユーザーモデル
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

// UserRepository は User データの永続化を扱うインターフェース
type UserRepository interface {
	Create(user *User) error
	FindByID(id uint) (*User, error)
	FindAll() ([]User, error)
	Update(user *User) error
	Delete(id uint) error
}

// userRepository は UserRepository の GORM 実装
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository は UserRepository の新しいインスタンスを作成
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create は新しいユーザーをDBに保存
func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

// FindByID はIDに基づいてユーザーを取得
func (r *userRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll は全てのユーザーを取得
func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	r.db.Find(&users)
	return users, nil
}

// Update はユーザー情報を更新
func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

// Delete はIDに基づいてユーザーを削除
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
