package chat

import (
	"kms/app/domain/entity"
	"strings"
)

func toListChatResponse(chatSessions []*entity.ChatSession, userRequester string, isIgnoreMessage bool) []*GetChatResponse {
	result := make([]*GetChatResponse, 0)
	for _, chatSession := range chatSessions {
		chat := toGetChatResponse(chatSession, userRequester)
		if isIgnoreMessage {
			chat.ChatMessages = nil
		}
		result = append(result, chat)
	}

	return result
}

func toGetChatResponse(chatSession *entity.ChatSession, userRequester string) *GetChatResponse {
	chatName := ""
	if len(chatSession.ChatParticipants) > 2 {
		chatName = generateChatNameFromUser(chatSession.ChatParticipants)
	} else {
		if chatSession.ChatParticipants[0].Username == userRequester {
			chatName = chatSession.ChatParticipants[1].User.FullName
		} else {
			chatName = chatSession.ChatParticipants[0].User.FullName
		}
	}

	mapParticipants := make(map[string]entity.ChatParticipant)
	for _, p := range chatSession.ChatParticipants {
		mapParticipants[p.Username] = p
	}

	messages := make([]*MessageResponse, 0)
	for _, m := range chatSession.ChatMessages {
		messages = append(messages, &MessageResponse{
			ID:         m.ID,
			Sender:     m.Sender,
			Message:    m.Message,
			Type:       string(m.Type),
			CreatedAt:  m.CreatedAt,
			PictureURL: mapParticipants[m.Sender].User.PictureURL,
			SenderName: mapParticipants[m.Sender].User.FullName,
		})
	}

	var latestMessage *MessageResponse
	if len(messages) > 0 {
		latestMessage = messages[0]
	}

	return &GetChatResponse{
		ID:            chatSession.ID,
		Name:          chatName,
		ChatMessages:  messages,
		LatestMessage: latestMessage,
		CreatedAt:     chatSession.CreatedAt,
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
