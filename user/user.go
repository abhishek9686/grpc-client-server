package user

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
)

// Users data
var (
	UserList = map[int64]*UserInfo{
		1: {
			Id:      1,
			Fname:   "bob",
			City:    "NewYork",
			Phone:   2345321125,
			Height:  5.10,
			Married: false,
		},
		2: {
			Id:      2,
			Fname:   "alice",
			City:    "Bangalore",
			Phone:   2345388812,
			Height:  5.10,
			Married: false,
		},
		3: {
			Id:      3,
			Fname:   "max",
			City:    "Tel Aviv",
			Phone:   474747728,
			Height:  6.2,
			Married: true,
		},
	}
)

// Server Implements UserDetailsServer
type Server struct {
	UnimplementedUserDetailsServer
}

func getUserInfo(id int64) (*UserInfo, error) {
	if val, ok := UserList[id]; ok {
		return val, nil
	}
	return &UserInfo{}, fmt.Errorf("User not found with ID: %d", id)
}
func getUserList(ids []int64) ([]*UserInfo, error) {
	userList := []*UserInfo{}
	var err error
	for _, id := range ids {
		user, err := getUserInfo(id)
		if err == nil {
			userList = append(userList, user)
		}
	}
	if len(userList) == 0 {
		err = fmt.Errorf("No Users Found For The Requested IDs: %v", ids)
	}
	return userList, err
}

// GetUserByID - get user details by providing user ID.
func (*Server) GetUserByID(ctx context.Context, in *UserRequest) (*UserRequestResponse, error) {
	resp := &UserRequestResponse{}
	if in == nil {
		return nil, errors.New("Request is nil")
	}
	if in.Id <= 0 {
		resp.Code = int32(codes.InvalidArgument)
		resp.Message = "ID should be a positive value"
		return resp, nil
	}
	userInfo, err := getUserInfo(in.Id)
	if err != nil {
		resp.Code = int32(codes.NotFound)
		resp.Message = err.Error()
		return resp, err
	}
	resp.User = userInfo
	resp.Code = int32(codes.OK)
	return resp, err
}

// ListUsersByID - list users by IDs
func (*Server) ListUsersByID(ctx context.Context, in *UserListRequest) (*UserListResponse, error) {
	resp := &UserListResponse{}
	if in == nil {
		return nil, errors.New("Request is nil")
	}

	if len(in.UserIDs) == 0 {
		resp.Code = int32(codes.InvalidArgument)
		resp.Message = "ID list is empty"
		return resp, nil
	}
	userList, err := getUserList(in.UserIDs)
	if err != nil {
		resp.Code = int32(codes.NotFound)
		resp.Message = err.Error()
		return resp, err
	}
	resp.Code = int32(codes.OK)
	resp.Users = userList
	return resp, nil
}
