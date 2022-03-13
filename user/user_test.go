package user_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/abhishek9686/grpc-client-server/user"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	user.RegisterUserDetailsServer(server, &user.Server{})
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestGetUserByID(t *testing.T) {
	testCases := []struct {
		name         string
		req          *user.UserRequest
		expectedCode int32
		expectedMsg  string
		expectedErr  bool
	}{
		{
			name: "req ok",
			req: &user.UserRequest{
				Id: 1,
			},
			expectedCode: int32(codes.OK),
			expectedMsg:  "",
		},
		{
			name: "req not found",
			req: &user.UserRequest{
				Id: 5,
			},
			expectedCode: int32(codes.NotFound),
			expectedMsg:  fmt.Sprintf("User not found with ID: %d", 5),
			expectedErr:  true,
		},
		{
			name: "req with negative id number",
			req: &user.UserRequest{
				Id: -5,
			},
			expectedCode: int32(codes.InvalidArgument),
			expectedMsg:  "ID should be a positive value",
			expectedErr:  false,
		},
		{
			name:        "req is nil",
			req:         nil,
			expectedMsg: "Request is nil",
			expectedErr: true,
		},
	}
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(dialer()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := user.NewUserDetailsClient(conn)
	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			//  call GetUserByID
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			resp, err := c.GetUserByID(ctx, testCase.req)
			if resp != nil {
				if testCase.expectedCode == resp.Code && testCase.expectedMsg == resp.Message {
					t.Logf("Test case: %s Passed", testCase.name)
				} else {
					t.Errorf("Test case: %s Failed", testCase.name)
				}
			}
			if err != nil && !testCase.expectedErr {

				t.Errorf("Test case: %s Failed. Due to '%v'", testCase.name, err)
			}

		})
	}
}

func TestListUsersByID(t *testing.T) {
	testCases := []struct {
		name         string
		req          *user.UserListRequest
		expectedCode int32
		expectedMsg  string
		expectedErr  bool
	}{
		{
			name: "req ok",
			req: &user.UserListRequest{
				UserIDs: []int64{1, 2},
			},
			expectedCode: int32(codes.OK),
			expectedMsg:  "",
		},
		{
			name: "req not found",
			req: &user.UserListRequest{
				UserIDs: []int64{5, 6},
			},
			expectedCode: int32(codes.NotFound),
			expectedMsg:  fmt.Sprintf("No Users Found For The Requested IDs: %v", []int64{5, 6}),
			expectedErr:  true,
		},
		{
			name: "req id list is empty",
			req: &user.UserListRequest{
				UserIDs: []int64{},
			},
			expectedCode: int32(codes.InvalidArgument),
			expectedMsg:  "ID list is empty",
			expectedErr:  false,
		},
		{
			name:        "req is nil",
			req:         nil,
			expectedMsg: "Request is nil",
			expectedErr: true,
		},
	}
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(dialer()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := user.NewUserDetailsClient(conn)
	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			//  call GetUserByID
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			resp, err := c.ListUsersByID(ctx, testCase.req)
			if resp != nil {
				if testCase.expectedCode == resp.Code && testCase.expectedMsg == resp.Message {
					t.Logf("Test case: %s Passed", testCase.name)
				} else {
					t.Errorf("Test case: %s Failed", testCase.name)
				}
			}
			if err != nil && !testCase.expectedErr {

				t.Errorf("Test case: %s Failed. Due to '%v'", testCase.name, err)
			}

		})
	}
}
