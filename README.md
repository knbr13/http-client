# HTTP Client

HTTP Client is a command-line tool for executing HTTP requests. It supports various HTTP methods such as GET, POST, PUT, DELETE, PATCH, OPTIONS, and HEAD.

## Screenshots:

![http-client](https://raw.githubusercontent.com/abdullah-alaadine/http-client/main/assets/screenshot.png)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/http-client.git
   ```
2. Build the project:
   ```shell
   cd http-client
   go build
   ```
## Usage

The general format of the command is:
  ```shell
  ./http-client [command] [arguments]
  ```
## Examples

1- Send an HTTP GET request:

```bash
  ./http-client httpget http://example.com
```

2- Send an HTTP POST request:
```bash
  ./http-client httppost http://example.com "body={key1: value, key2: value}"
```

3- Send an HTTP DELETE request:

```bash
  ./http-client httpdelete http://example.com
```
4- Run the tool without args to get a helpful message explaining how to use this tool.

## License

This project is licensed under the [MIT License](https://github.com/abdullah-alaadine/http-client/blob/main/LICENSE). See the [LICENSE](https://github.com/abdullah-alaadine/http-client/blob/main/LICENSE) file for details.

