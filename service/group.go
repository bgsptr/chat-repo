package service

import (
	"context"
	"errors"
	"fmt"
	"groupservice/model"
	"groupservice/repository"
	"log"
	"time"
)

type GroupService struct {
	GroupRepository *repository.GroupRepository
}

func NewGroupService(GroupRepository *repository.GroupRepository) *GroupService {
	return &GroupService{
		GroupRepository: GroupRepository,
	}
}

var (
	deadline = time.Now().Add(10 * time.Second)
	errMessage = errors.New("already join")
)

func (g *GroupService) JoinGroup(username string, roomID string) (bool, error) {
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	findUser, err := g.GroupRepository.AlreadyJoin(ctx, username, roomID)
	if !findUser {
		log.Println(err)
		return false, errMessage
	}

	err = g.GroupRepository.Join(ctx, username, roomID)
	if err != nil {
		return false, errMessage
	}

	return true, nil
}

func (g *GroupService) CreateGroup(ctx context.Context, cl *model.Client, roomID string) (*model.Chat, error) {
    msg := &model.Chat{
        From:      cl,
        To:        make(map[string]*model.Client),
        Content:   fmt.Sprintf("%s just created room", cl.Username),
        CreatedAt: time.Now().Local().String(),
    }

    err := g.GroupRepository.Create(ctx, cl, roomID)
    if err != nil {
        return nil, errMessage
    }

    return msg, nil
}
