# GoRemote

GoRemote is a remote mouse project that allows you to control mouse actions from your mobile phone (preferably android) (left-click, right-click) and send mouse coordinates over a WebSocket connection. 
This README provides an overview of the project and instructions for setting up and running the server and client components.

## Features

- Control mouse actions remotely: left-click and right-click.
- Send mouse coordinates to the server for remote control.

## Server

The server component of GoRemote is written in Go and handles WebSocket connections from clients. It receives mouse coordinate input, left-click, and right-click commands from the connected clients. 
Designed to run on windows machines.

### Setup and Usage

1. Clone the GoRemote repository to your local machine:

   ```bash
   git clone https://github.com/OguzhanKazak/goremote.git
   ```

2. Navigate to the server directory:

   ```bash
   cd goremote/server
   ```

3. Install the required dependencies:

   ```bash
   go mod download
   ```

4. Start the server:

   ```bash
   go run .
   ```

   The server will start running on the default port `8080`. You can modify the port in the `main.go` file if needed.

OR

1. Clone the GoRemote repository to your local machine:

   ```bash
   execute /build/goremote.exe directly
   ```


## Client

The client component of GoRemote is a web page (`client.htm`) that contains JavaScript code for sending mouse actions and coordinates over the WebSocket connection. Client should change the local ip address to the host's local ip and port before open the file.

### Usage

1. Open the `client.htm` file in a web browser.

2. Connect to the server by clicking the "Connect" button on the web page.

3. Once connected, you can perform the following actions:

   - Left-click: Default touch action on android.
   - Right-click: Keep touching same point to activate rightClick event.
   - Mouse coordinates: Move your touch point, and the current touch coordinates will be sent to the server.

## Contributing

Contributions to GoRemote are welcome! If you find any bugs, have feature requests, or want to contribute enhancements, please submit an issue or create a pull request on the GitHub repository.

## Contact

For any questions or inquiries, please contact [oguzhankazak@outlook.com](mailto:oguzhankazak@outlook.com).

---