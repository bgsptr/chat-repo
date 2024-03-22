package service

import (
	"errors"
	"fmt"
	"groupservice/model"
)

type GroupService struct {
	GroupRepository *repository.GroupRepository

}

func JoinGroup() {
	g.Repository.Create()
}

func (g *GroupService) CreateGroup(cl model.Client, roomID string) (*model.Message, error) {
	msg := &model.Message{
		From: cl,
		RoomID: roomID,
		Content: fmt.Sprintf("%s just create the room", cl.Username),
	}

	g.GroupRepository.Create()
	return msg, nil
}