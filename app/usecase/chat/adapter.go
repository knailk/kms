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
		return result[i].UpdatedAt.After(result[j].UpdatedAt)
	})

	return result
}

func toGetChatResponse(chatSession *entity.ChatSession, userRequester string) *GetChatResponse {
	chatName := ""
	chatPicture := ""
	if len(chatSession.ChatParticipants) > 2 {
		chatName = generateChatNameFromUser(chatSession.ChatParticipants)
	} else {
		if chatSession.ChatParticipants[0].Username == userRequester {
			chatName = chatSession.ChatParticipants[1].User.FullName
		} else {
			chatName = chatSession.ChatParticipants[0].User.FullName
			chatPicture = chatSession.ChatParticipants[0].User.PictureURL
		}
	}

	mapParticipants := make(map[string]entity.ChatParticipant)
	for _, p := range chatSession.ChatParticipants {
		mapParticipants[p.Username] = p
	}

	messages := filterMessagesByDate(chatSession.ChatMessages, mapParticipants)

	var latestMsg *MessageResponse
	if chatSession.LatestMessage != nil {
		latestMsg = &MessageResponse{
			ID:      chatSession.LatestMessage.ID,
			Message: chatSession.LatestMessage.Message,
			Type:    chatSession.LatestMessage.Type.String(),

			CreatedAt: chatSession.LatestMessage.CreatedAt,
			UpdatedAt: chatSession.LatestMessage.UpdatedAt,
			SenderResponse: &SenderResponse{
				Username:   chatSession.LatestMessage.Sender,
				Avatar:     mapParticipants[chatSession.LatestMessage.Sender].User.PictureURL,
				SenderName: mapParticipants[chatSession.LatestMessage.Sender].User.FullName,
			},
		}
	}

	return &GetChatResponse{
		ID:            chatSession.ID,
		Name:          chatName,
		ChatMessages:  messages,
		ChatPicture:   chatPicture,
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

func generateChatNameFromUser(user []entity.ChatParticipant) string {
	listName := make([]string, 0)
	for _, p := range user {
		listName = append(listName, p.User.FullName)
	}

	return generateChatName(listName)
}

func filterMessagesByDate(messages []entity.ChatMessage, mapParticipants map[string]entity.ChatParticipant) []MessageByDate {
	messageMap := make(map[string][]*MessageResponse)

	for _, msg := range messages {
		date := msg.CreatedAt.Format("2006-01-02")
		messageMap[date] = append(messageMap[date], &MessageResponse{
			ID:        msg.ID,
			Message:   msg.Message,
			Type:      msg.Type.String(),
			CreatedAt: msg.CreatedAt,
			UpdatedAt: msg.UpdatedAt,
			SenderResponse: &SenderResponse{
				Username:   msg.Sender,
				Avatar:     mapParticipants[msg.Sender].User.PictureURL,
				SenderName: mapParticipants[msg.Sender].User.FullName,
			},
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
