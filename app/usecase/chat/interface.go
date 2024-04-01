package chat

import "context"

type IUseCase interface {
	// chat
	CreateChat(ctx context.Context, req *CreateChatRequest) (*CreateChatResponse, error)
	AddMember(ctx context.Context, req *AddMemberRequest) (*AddMemberResponse, error)
	ListChats(ctx context.Context, req *ListChatsRequest) (*ListChatsResponse, error)
	GetChat(ctx context.Context, req *GetChatRequest) (*GetChatResponse, error)
	UpdateChat(ctx context.Context, req *UpdateChatRequest) (*UpdateChatResponse, error)
	DeleteChat(ctx context.Context, req *DeleteChatRequest) (*DeleteChatResponse, error)

	// message
	CreateMessage(ctx context.Context, req *CreateMessageRequest) (*CreateMessageResponse, error)
}
