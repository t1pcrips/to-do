package user

import (
	"context"
	"todo/internal/converter"
	"todo/internal/handlers"
	"todo/internal/service"
	apiUsers "todo/internal/web/users"
)

type UserHanlderImpl struct {
	ServiceUser service.UserService
}

func NewUserHandler(uServ service.UserService) handlers.UserHandler {
	return &UserHanlderImpl{
		ServiceUser: uServ,
	}
}

func (h *UserHanlderImpl) PostUsers(ctx context.Context, req apiUsers.PostUsersRequestObject) (apiUsers.PostUsersResponseObject, error) {
	info := converter.ToModelFromApiCreateUser(&req)

	userId, err := h.ServiceUser.CreateUser(ctx, info)
	if err != nil {
		return nil, err
	}

	return apiUsers.PostUsers201JSONResponse{
		Id: &userId,
	}, nil
}

func (h *UserHanlderImpl) PatchUsers(ctx context.Context, req apiUsers.PatchUsersRequestObject) (apiUsers.PatchUsersResponseObject, error) {
	info := converter.ToModelFromApiUpdateUser(&req)

	err := h.ServiceUser.UpdateUser(ctx, info)
	if err != nil {
		return nil, err
	}

	return apiUsers.PatchUsers200JSONResponse{}, nil
}

func (h *UserHanlderImpl) DeleteUsers(ctx context.Context, req apiUsers.DeleteUsersRequestObject) (apiUsers.DeleteUsersResponseObject, error) {
	err := h.ServiceUser.DeleteUser(ctx, req.Params.Id)
	if err != nil {
		return nil, err
	}

	return apiUsers.DeleteUsers204Response{}, nil
}
