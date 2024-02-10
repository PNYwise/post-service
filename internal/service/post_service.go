package service

import (
	"errors"

	"github.com/PNYwise/post-service/internal/domain"
	"github.com/PNYwise/post-service/internal/util"
)

type postService struct {
	postRepository      domain.IPostRepository
	kafkaPostRepository domain.KafkaPostRepository
}

func NewPostService(postRepository domain.IPostRepository, kafkaPostRepository domain.KafkaPostRepository) domain.IPostService {
	return &postService{
		postRepository:      postRepository,
		kafkaPostRepository: kafkaPostRepository,
	}
}

func (p *postService) Create(request *domain.PostRequest) (*domain.Post, error) {
	if errs := util.Validate(request); len(errs) > 0 && errs[0].Error {
		return nil, util.ValidationErrMsg(errs)
	}
	post := &domain.Post{
		UserUuid: request.UserUuid,
		Caption:  request.Caption,
		ImageUrl: request.ImageUrl,
		Location: request.Location,
	}
	if err := p.postRepository.Create(post); err != nil {
		return nil, errors.New("Internal Server Error")
	}
	if err := p.kafkaPostRepository.PublishMessage(post); err != nil {
		return nil, errors.New("Internal Server Error")
	}
	return post, nil
}
func (p *postService) ReadAllByUserId(userUuid string) (*[]domain.Post, error) {
	if errs := util.Var(userUuid, "required,uuid4"); len(errs) > 0 && errs[0].Error {
		return nil, util.ValidationErrMsg(errs)
	}
	posts, err := p.postRepository.ReadAllByUserId(userUuid)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// Exist implements domain.IPostService.
func (p *postService) Exist(uuid string) (bool, error) {
	if errs := util.Var(uuid, "required,uuid4"); len(errs) > 0 && errs[0].Error {
		return false, util.ValidationErrMsg(errs)
	}
	exist, err := p.postRepository.Exist(uuid)
	if err != nil {
		return false, errors.New("internal server error")
	}
	return exist, nil
}

func (p *postService) Delete(uuid string) error {
	if errs := util.Var(uuid, "required,uuid4"); len(errs) > 0 && errs[0].Error {
		return util.ValidationErrMsg(errs)
	}
	if err := p.postRepository.Delete(uuid); err != nil {
		return errors.New(err.Error())
	}
	return nil
}
