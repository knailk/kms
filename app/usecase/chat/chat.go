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
	repo *repository.PostgresRepository
}

func NewUseCase(repo *repository.PostgresRepository) IUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) CreateChat(ctx context.Context, req *CreateChatRequest) (*CreateChatResponse, error) {
	const op errs.Op = "useCase.chat.CreateChat"

	errKind := req.Validate()
	if errKind != errs.Other {
		return nil, errs.E(op, errKind, "validate request error")
	}

	chatSessionID := uuid.New()

	users, err := uc.repo.User.Where(uc.repo.User.Username.In(append([]string{req.Owner}, req.Participants...)...)).Find()
	if err != nil {
		logger.Error(op, " get user error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	participants := []entity.ChatParticipant{}
	fullName := []string{}
	for _, user := range users {
		participants = append(participants, entity.ChatParticipant{
			ID:            uuid.New(),
			ChatSessionID: chatSessionID,
			Username:      user.Username,
			IsOwner:       user.Username == req.Owner,
		})

		fullName = append(fullName, user.FullName)
	}

	err = uc.repo.ChatSession.Create(&entity.ChatSession{
		ID:               chatSessionID,
		Name:             generateChatName(fullName),
		ChatParticipants: participants,
	})
	if err != nil {
		logger.Error(op, err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &CreateChatResponse{}, nil
}

func (uc *useCase) AddMember(ctx context.Context, req *AddMemberRequest) (*AddMemberResponse, error) {
	const op errs.Op = "useCase.chat.AddMember"
	chatSession, err := uc.repo.ChatSession.
		Where(uc.repo.ChatSession.ID.Eq(req.ChatSessionID)).
		Preload(uc.repo.ChatSession.ChatParticipants).
		First()
	if err != nil {
		logger.Error(op, " get chat session error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	if len(chatSession.ChatParticipants) >= 3 {
		err = uc.repo.ChatParticipant.Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "username"}, {Name: "chat_session_id"}},
				DoNothing: true,
			},
		).Create(&entity.ChatParticipant{
			ID:            uuid.New(),
			ChatSessionID: req.ChatSessionID,
			Username:      req.Username,
			IsOwner:       false,
		})
		if err != nil {
			logger.Error(op, err)
			return nil, errs.E(op, errs.Database, err)
		}
	} else {
		participants := []string{req.Username}
		for _, p := range chatSession.ChatParticipants {
			if p.Username == req.Adder {
				continue
			}
			participants = append(participants, p.Username)
		}

		_, err = uc.CreateChat(ctx, &CreateChatRequest{
			Participants: participants,
			Owner:        req.Adder,
		})
		if err != nil {
			logger.Error(op, err)
			return nil, errs.E(op, errs.Database, err)
		}
	}

	return &AddMemberResponse{
		IsNewChat: len(chatSession.ChatParticipants) < 3,
	}, nil
}

func (uc *useCase) ListChats(ctx context.Context, req *ListChatsRequest) (*ListChatsResponse, error) {
	const op errs.Op = "useCase.chat.ListChats"

	chatSessions, err := uc.repo.ChatSession.
		LeftJoin(
			uc.repo.ChatParticipant,
			uc.repo.ChatSession.ID.EqCol(uc.repo.ChatParticipant.ChatSessionID)).
		Where(uc.repo.ChatParticipant.Username.Eq(req.UserRequester)).
		Preload(uc.repo.ChatSession.LatestMessage).
		Preload(uc.repo.ChatSession.ChatParticipants).
		Preload(uc.repo.ChatSession.ChatParticipants.User).
		Find()
	if err != nil {
		logger.Error(op, " get chat session error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &ListChatsResponse{
		ChatSessions: toListChatResponse(chatSessions, req.UserRequester),
	}, nil
}

func (uc *useCase) GetChat(ctx context.Context, req *GetChatRequest) (*GetChatResponse, error) {
	const op errs.Op = "useCase.chat.GetChat"
	chatSession, err := uc.repo.ChatSession.
		LeftJoin(
			uc.repo.ChatParticipant,
			uc.repo.ChatSession.ID.EqCol(uc.repo.ChatParticipant.ChatSessionID)).
		Where(
			uc.repo.ChatSession.ID.Eq(req.ChatSessionID),
			uc.repo.ChatParticipant.Username.Eq(req.UserRequester),
		).
		Preload(uc.repo.ChatSession.ChatParticipants).
		Preload(uc.repo.ChatSession.ChatParticipants.User).
		Preload(uc.repo.ChatSession.ChatMessages.Order(uc.repo.ChatMessage.CreatedAt).Limit(req.Limit)).
		First()
	if err != nil {
		logger.Error(op, " get chat session error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return toGetChatResponse(chatSession, req.UserRequester), nil
}

func (uc *useCase) UpdateChat(ctx context.Context, req *UpdateChatRequest) (*UpdateChatResponse, error) {
	const op errs.Op = "useCase.chat.UpdateChat"
	_, err := uc.repo.ChatSession.Where(uc.repo.ChatSession.ID.Eq(req.ChatSessionID)).Updates(map[string]interface{}{
		"name": req.Name,
	})
	if err != nil {
		logger.Error(op, " update chat session error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &UpdateChatResponse{}, nil
}

func (uc *useCase) DeleteChat(ctx context.Context, req *DeleteChatRequest) (*DeleteChatResponse, error) {
	const op errs.Op = "useCase.chat.DeleteChat"
	_, err := uc.repo.ChatSession.Where(uc.repo.ChatSession.ID.Eq(req.ChatSessionID)).Delete()
	if err != nil {
		logger.Error(op, " delete chat session error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &DeleteChatResponse{}, nil
}

func (uc *useCase) CreateMessage(ctx context.Context, req *CreateMessageRequest) (*CreateMessageResponse, error) {
	const op errs.Op = "useCase.chat.CreateMessage"
	errKind := req.Validate()
	if errKind != errs.Other {
		return nil, errs.E(op, errKind, "validate request error")
	}

	count, err := uc.repo.ChatParticipant.Where(uc.repo.ChatParticipant.ChatSessionID.Eq(req.ChatSessionID), uc.repo.ChatParticipant.Username.Eq(req.Sender)).Count()
	if err != nil {
		logger.Error(op, " get participant error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	if count == 0 {
		return nil, errs.E(op, errs.Invalid, "sender not in chat")
	}

	uc.repo.Transaction(func(tx *repository.Query) error {
		msgID := uuid.New()
		err = tx.ChatMessage.Create(&entity.ChatMessage{
			ID:            msgID,
			ChatSessionID: req.ChatSessionID,
			Sender:        req.Sender,
			Message:       req.Message,
			Type:          req.Type,
		})
		if err != nil {
			return errs.E(op, errs.Database, err)
		}

		_, err := tx.ChatSession.Where(tx.ChatSession.ID.Eq(req.ChatSessionID)).Updates(
			&entity.ChatSession{
				LatestMessageID: msgID,
			},
		)
		if err != nil {
			return errs.E(op, errs.Database, err)
		}

		return nil
	})
	if err != nil {
		logger.Error(op, " create message error :", err)
		return nil, err
	}

	return &CreateMessageResponse{}, nil
}
