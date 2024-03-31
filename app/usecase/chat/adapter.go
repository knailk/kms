package chat

import (
	"kms/app/domain/entity"
	"strings"
)

func toListChatResponse(chatSessions []*entity.ChatSession, userRequester string) []*GetChatResponse {
	result := make([]*GetChatResponse, 0)
	for _, chatSession := range chatSessions {
		result = append(result, toGetChatResponse(chatSession, userRequester))
	}

	return result
}

func toGetChatResponse(chatSession *entity.ChatSession, userRequester string) *GetChatResponse {
	chatName := ""
	if len(chatSession.ChatParticipants) > 2 {
		chatName = chatSession.Name
	} else {
		if chatSession.ChatParticipants[0].UserID == userRequester {
			chatName = chatSession.ChatParticipants[1].User.FullName
		} else {
			chatName = chatSession.ChatParticipants[0].User.FullName
		}
	}

	mapParticipants := make(map[string]entity.ChatParticipant)
	for _, p := range chatSession.ChatParticipants {
		mapParticipants[p.UserID] = p
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

	return &GetChatResponse{
		ID:           chatSession.ID,
		Name:         chatName,
		ChatMessages: messages,
		CreatedAt:    chatSession.CreatedAt,
	}
}

func generateChatName(participants []string) string {
	// Join the usernames with commas
	joinedNames := strings.Join(participants, ", ")

	// Truncate the joined names if it's too long
	const maxCharacters = 30
	if len(joinedNames) > maxCharacters {
		joinedNames = joinedNames[:maxCharacters] + "..."
	}

	// Construct the group chat name
	return joinedNames
}
