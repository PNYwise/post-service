package service

import (
	"testing"

	"github.com/PNYwise/post-service/internal/domain"
	"github.com/PNYwise/post-service/internal/repository/_mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	// Create a mock repository
	mockRepo := new(_mock.MockPostRepository)
	kafkaPostRepository := new(_mock.MockKafkaPostRepository)
	// Create a post service with the mock repository
	postService := NewPostService(mockRepo, kafkaPostRepository)

	fakeUserUUID := uuid.New().String()

	// Create a sample post request
	postRequest := &domain.PostRequest{
		UserUuid: fakeUserUUID,
		Caption:  "",
		ImageUrl: "http://example.com/image.jpg",
		Location: &domain.Location{
			Lat: 746.9327140312029,
			Lng: 400.7438706958651,
		},
	}

	post := &domain.Post{
		UserUuid: fakeUserUUID,
		Caption:  "",
		ImageUrl: "http://example.com/image.jpg",
		Location: &domain.Location{
			Lat: 746.9327140312029,
			Lng: 400.7438706958651,
		},
	}

	// Expect the Create method to be called with the correct argument
	mockRepo.On("Create", post).Return(nil)
	kafkaPostRepository.On("PublishMessage", post).Return(nil)

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
	mockRepo := new(_mock.MockPostRepository)
	kafkaPostRepo := new(_mock.MockKafkaPostRepository)
	// Create a post service with the mock repository
	postService := NewPostService(mockRepo, kafkaPostRepo)

	// Define a fake user UUID
	fakeUserUUID := uuid.New().String()

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

func TestExist(t *testing.T) {
	// Create a mock repository
	mockRepo := new(_mock.MockPostRepository)
	kafkaPostRepo := new(_mock.MockKafkaPostRepository)
	// Create a post service with the mock repository
	postService := NewPostService(mockRepo, kafkaPostRepo)

	// Define a fake user UUID
	fakeUUID := uuid.New().String()

	// Expect the Exist method to be called with the correct argument
	mockRepo.On("Exist", fakeUUID).Return(true, nil)

	// Call the Exist method of the post service
	exist, err := postService.Exist(fakeUUID)

	// Assert that the mock repository's Exist method was called with the correct argument
	mockRepo.AssertExpectations(t)

	// Assert that the returned bool and error match the expected values
	assert.NoError(t, err)
	assert.Equal(t, exist, true)
}

func TestDeletePost(t *testing.T) {
	// Create a mock repository
	mockRepo := new(_mock.MockPostRepository)
	kafkaPostRepo := new(_mock.MockKafkaPostRepository)

	// Create a post service with the mock repository
	postService := NewPostService(mockRepo, kafkaPostRepo)

	// Define a fake user UUID
	fakeUUID := uuid.New().String()

	// Expect the Delete method to be called with the correct argument
	mockRepo.On("Delete", fakeUUID).Return(nil)

	// Call the Delete method of the post service
	err := postService.Delete(fakeUUID)

	// Assert that the mock repository's Delete method was called with the correct argument
	mockRepo.AssertExpectations(t)

	// Assert that the returned posts and error match the expected values
	assert.NoError(t, err)
	assert.Equal(t, nil, err)
}
