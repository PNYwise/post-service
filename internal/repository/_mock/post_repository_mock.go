package _mock

import (
	"github.com/PNYwise/post-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockPostRepository is a mock implementation of IPostRepository
type MockPostRepository struct {
	mock.Mock
}

// Create mocks the Create method of IPostRepository
func (m *MockPostRepository) Create(post *domain.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

// ReadAllByUserId mocks the ReadAllByUserId method of IPostRepository
func (m *MockPostRepository) ReadAllByUserId(userUuid string) (*[]domain.Post, error) {
	args := m.Called(userUuid)
	return args.Get(0).(*[]domain.Post), args.Error(1)
}

// Exist mocks the Exist method of IPostRepository
func (m *MockPostRepository) Exist(uuid string) (bool, error) {
	args := m.Called(uuid)
	return args.Get(0).(bool), args.Error(1)
}

// Delete mocks the Delete method of IPostRepository
func (m *MockPostRepository) Delete(uuid string) error {
	args := m.Called(uuid)
	return args.Error(0)
}
