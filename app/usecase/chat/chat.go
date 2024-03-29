package chat

import (
	"context"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/logger"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type useCase struct {
	repo repository.PostgresRepository
}

func NewUseCase(repo repository.PostgresRepository) IUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) CreateChat(ctx context.Context, req *CreateChatRequest) (*CreateChatResponse, error) {
	const op errs.Op = "useCase.chat.CreateChat"
	chatSessionID := uuid.New()

	participants := []entity.ChatParticipant{}
	for _, user := range req.Participants {
		participants = append(participants, entity.ChatParticipant{
			ID:            uuid.New(),
			ChatSessionID: chatSessionID,
			UserID:        user,
		})
	}

	participants = append(participants, entity.ChatParticipant{
		ID:            uuid.New(),
		ChatSessionID: chatSessionID,
		UserID:        req.Owner,
		IsOwner:       true,
	})

	message := entity.ChatMessage{
		ID:            uuid.New(),
		ChatSessionID: chatSessionID,
		Sender:        req.Owner,
		Message:       req.Message,
		Type:          req.MessageType,
	}

	err := uc.repo.ChatSession.Create(&entity.ChatSession{
		ID:               chatSessionID,
		Name:             req.Name,
		ChatParticipants: participants,
		ChatMessages:     []entity.ChatMessage{message},
	})
	if err != nil {
		logger.Error(op, err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &CreateChatResponse{}, nil
}

func (uc *useCase) AddMember(ctx context.Context, req *AddMemberRequest) (*AddMemberResponse, error) {
	const op errs.Op = "useCase.chat.AddMember"
	err := uc.repo.ChatParticipant.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "chat_session_id"}},
			DoNothing: true,
		},
	).Create(&entity.ChatParticipant{
		ID:            uuid.New(),
		ChatSessionID: req.ChatSessionID,
		UserID:        req.UserID,
		IsOwner:       false,
	})
	if err != nil {
		logger.Error(op, err)
		return nil, errs.E(op, errs.Database, err)
	}
	return &AddMemberResponse{}, nil
}

func (uc *useCase) ListChats(ctx context.Context, req *ListChatsRequest) (*ListChatsResponse, error) {
	return nil, nil
}
func (uc *useCase) GetChat(ctx context.Context, req *GetChatRequest) (*GetChatResponse, error) {
	return nil, nil
}
func (uc *useCase) UpdateChat(ctx context.Context, req *UpdateChatRequest) (*UpdateChatResponse, error) {
	return nil, nil
}
