/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abhishek9686/grpc-client-server/user"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr string
	userId     int64
	idList     []int64
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func establishConnection() (*grpc.ClientConn, user.UserDetailsClient) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := user.NewUserDetailsClient(conn)
	return conn, c
}

var getUserCmd = &cobra.Command{
	Use: "getUser",
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		conn, userClient := establishConnection()
		if conn == nil {
			log.Fatal("Connection is lost quitting.")
		}
		defer conn.Close()
		resp, err := userClient.GetUserByID(ctx, &user.UserRequest{Id: userId})
		if err != nil {
			fmt.Printf("Failed to get UserInfo for userID: %d, Err: %v", userId, err)
		}
		fmt.Printf("Server Resp:: %+v\n", resp)
	},
}

var getUserListCmd = &cobra.Command{
	Use: "getUserList",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		conn, userClient := establishConnection()
		if conn == nil {
			log.Fatal("Connection is lost quitting.")
		}
		defer conn.Close()
		resp, err := userClient.ListUsersByID(ctx, &user.UserListRequest{UserIDs: idList})
		if err != nil {
			fmt.Printf("Failed to get UserInfoList for userIDs: %v, Err: %v", idList, err)
		}
		fmt.Printf("Server Resp: %+v\n", resp)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&serverAddr, "svrAddr", "s", "localhost:8008", "gRpc Server Address")
	rootCmd.AddCommand(getUserCmd)
	getUserCmd.Flags().Int64VarP(&userId, "userId", "i", 0, "user id")
	rootCmd.AddCommand(getUserListCmd)
	getUserListCmd.Flags().Int64SliceVarP(&idList, "idList", "l", []int64{}, "list of user ids")
}
