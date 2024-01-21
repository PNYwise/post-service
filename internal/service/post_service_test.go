package service

import (
	"testing"

	"github.com/PNYwise/post-service/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

// Delete mocks the Delete method of IPostRepository
func (m *MockPostRepository) Delete(uuid string) error {
	args := m.Called(uuid)
	return args.Error(0)
}

func TestCreatePost(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockPostRepository)

	// Create a post service with the mock repository
	postService := NewPostService(mockRepo)

	fakeUUID := uuid.New().String()

	// Create a sample post request
	postRequest := &domain.PostRequest{
		UserUuid: fakeUUID,
		Caption:  "Test caption",
		ImageUrl: "http://example.com/image.jpg",
		Location: &domain.Location{
			Lat: 746.9327140312029,
			Lng: 400.7438706958651,
		},
	}

	// Expect the Create method to be called with the correct argument
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Call the Create method of the post service
	createdPost, err := postService.Create(postRequest)

	// Assert that the mock repository's Create method was called with the correct argument
	mockRepo.AssertExpectations(t)

	// Assert that the returned post and error match the expected values
	assert.NoError(t, err)
	assert.Equal(t, postRequest.UserUuid, createdPost.UserUuid)
	assert.Equal(t, postRequest.Caption, createdPost.Caption)
	assert.Equal(t, postRequest.ImageUrl, createdPost.ImageUrl)
	assert.Equal(t, postRequest.Location, createdPost.Location)
}

func TestReadAllByUserId(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockPostRepository)

	// Create a post service with the mock repository
	postService := NewPostService(mockRepo)

	// Define a fake user UUID
	fakeUserUUID := "fakeUser123"

	// Create an example list of posts for the user
	fakePosts := []domain.Post{
		{
			UserUuid: fakeUserUUID, Caption: "Post 1", ImageUrl: "http://example.com/image.jpg",
			Location: &domain.Location{
				Lat: 746.9327140312029,
				Lng: 400.7438706958651,
			},
		},
		{
			UserUuid: fakeUserUUID, Caption: "Post 2", ImageUrl: "http://example.com/image.jpg",
			Location: &domain.Location{
				Lat: 746.9327140312029,
				Lng: 400.7438706958651,
			},
		},
	}

	// Expect the ReadAllByUserId method to be called with the correct argument
	mockRepo.On("ReadAllByUserId", fakeUserUUID).Return(&fakePosts, nil)

	// Call the ReadAllByUserId method of the post service
	resultPosts, err := postService.ReadAllByUserId(fakeUserUUID)

	// Assert that the mock repository's ReadAllByUserId method was called with the correct argument
	mockRepo.AssertExpectations(t)

	// Assert that the returned posts and error match the expected values
	assert.NoError(t, err)
	assert.Equal(t, fakePosts, *resultPosts)
}
