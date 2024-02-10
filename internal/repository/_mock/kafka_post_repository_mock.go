package _mock

import (
	"github.com/PNYwise/post-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockPostRepository is a mock implementation of IPostRepository
type MockKafkaPostRepository struct {
	mock.Mock
}

// Create mocks the Create method of IPostRepository
func (m *MockKafkaPostRepository) PublishMessage(post *domain.Post) error {
	args := m.Called(post)
	return args.Error(0)
}
