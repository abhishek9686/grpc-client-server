# gRpc User server/client

## To build the application

### server side: (All the below commands show be run the root of the project directory)

    1. To Compile and Dockerize application:
        `make build-dockerize-server`
    2. To Compile the server side code:
        `make build-server`
        - To execute the binary
            `./server/users --port=60000`

### client side: (client to interact with the server)

    1. To Compile the application:
        `make build-client`
    2. To execute the binary: (Client is a Cli tool)
    ``Usage:

    client [command]

    Available Commands:
    completion Generate the autocompletion script for the specified shell
    getUser
    getUserList
    help Help about any command

    Flags:
    -h, --help help for client
    -s, --svrAddr string gRpc Server Address (default "localhost:60000")
    -t, --toggle Help message for toggle

    Use "client [command] --help" for more information about a command.``

    3. To Get User By ID:
        ``./client/usersClient getUser --userId 1` (for more info user [-h] flag)``
        ``Usage:

        client getUser [flags]

        Flags:
        -h, --help help for getUser
        -i, --userId int user id``

    4. To Get List of Users By Their Ids:
        ``./client/usersClient getUserList --idList 1,2``
        ``Usage:

        client getUserList [flags]

        Flags:
        -h, --help help for getUserList
        -l, --idList int64Slice list of user ids (default [])``
