package biz

import (
	"fmt"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
)

func userModelsToReply(users []*User) []*v1.User {
	var res []*v1.User
	for _, user := range users {
		res = append(res, &v1.User{
			Id:    user.ID,
			Name:  user.Name,
			Color: fmt.Sprintf("%x", user.Color),
		})
	}
	return res
}

func msgModelsToReply(msgs []*ChatMessage) []*v1.NewMessage {
	var res []*v1.NewMessage
	for _, msg := range msgs {
		res = append(res, &v1.NewMessage{
			Id: msg.ID,
			User: &v1.User{
				Id:    msg.UserID,
				Name:  msg.Name,
				Color: fmt.Sprintf("%x", msg.Color),
			},
			Text: msg.Text,
		})
	}
	return res
}
