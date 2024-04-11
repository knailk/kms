package chat

import (
	"kms/app/domain/entity"
	"sort"
	"strings"
)

func toListChatResponse(chatSessions []*entity.ChatSession, userRequester string) []*GetChatResponse {
	result := make([]*GetChatResponse, 0)
	for _, chatSession := range chatSessions {
		result = append(result, toGetChatResponse(chatSession, userRequester))
	}

	sort.SliceStable(result, func(i, j int) bool {
		result[i].Members = nil
		result[j].Members = nil
		return result[i].UpdatedAt.After(result[j].UpdatedAt)
	})

	return result
}

func toGetChatResponse(chatSession *entity.ChatSession, userRequester string) *GetChatResponse {
	chatName := ""
	chatPicture := ""
	if len(chatSession.ChatParticipants) > 2 {
		chatName = chatSession.Name
	} else {
		if chatSession.ChatParticipants[0].Username == userRequester {
			chatName = chatSession.ChatParticipants[1].User.FullName
			chatPicture = chatSession.ChatParticipants[1].User.PictureURL
		} else {
			chatName = chatSession.ChatParticipants[0].User.FullName
			chatPicture = chatSession.ChatParticipants[0].User.PictureURL
		}
	}

	mapParticipants := make(map[string]*MemberResponse)
	for _, p := range chatSession.ChatParticipants {
		mapParticipants[p.Username] = &MemberResponse{
			Username:   p.Username,
			Avatar:     p.User.PictureURL,
			SenderName: p.User.FullName,
		}
	}

	messages := filterMessagesByDate(chatSession.ChatMessages, mapParticipants)

	var latestMsg *MessageResponse
	if chatSession.LatestMessage != nil {
		latestMsg = &MessageResponse{
			ID:      chatSession.LatestMessage.ID,
			Message: chatSession.LatestMessage.Message,
			Type:    chatSession.LatestMessage.Type.String(),

			CreatedAt:      chatSession.LatestMessage.CreatedAt,
			UpdatedAt:      chatSession.LatestMessage.UpdatedAt,
			SenderResponse: mapParticipants[chatSession.LatestMessage.Sender],
		}
	}

	return &GetChatResponse{
		ID:            chatSession.ID,
		Name:          chatName,
		ChatPicture:   chatPicture,
		ChatMessages:  messages,
		Members:       toMemberSlice(mapParticipants),
		LatestMessage: latestMsg,
		CreatedAt:     chatSession.CreatedAt,
		UpdatedAt:     chatSession.UpdatedAt,
	}
}

func generateChatName(participants []string) string {
	// Join the usernames with commas
	joinedNames := strings.Join(participants, ", ")

	// Construct the group chat name
	return joinedNames
}

func filterMessagesByDate(messages []entity.ChatMessage, mapParticipants map[string]*MemberResponse) []MessageByDate {
	messageMap := make(map[string][]*MessageResponse)

	for _, msg := range messages {
		date := msg.CreatedAt.Format("2006-01-02")
		messageMap[date] = append(messageMap[date], &MessageResponse{
			ID:             msg.ID,
			Message:        msg.Message,
			Type:           msg.Type.String(),
			CreatedAt:      msg.CreatedAt,
			UpdatedAt:      msg.UpdatedAt,
			SenderResponse: mapParticipants[msg.Sender],
		})
	}

	var dates []string
	for date := range messageMap {
		dates = append(dates, date)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(dates)))

	var result []MessageByDate
	for _, date := range dates {
		result = append(result, MessageByDate{
			Date:     date,
			Messages: messageMap[date],
		})
	}

	return result
}

func toMemberSlice(req map[string]*MemberResponse) []*MemberResponse {
	result := make([]*MemberResponse, 0)
	for _, v := range req {
		result = append(result, v)
	}
	return result
}
