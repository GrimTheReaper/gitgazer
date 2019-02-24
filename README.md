# GitGazer

A challenge entry for CenturyLink (1/3), written in Golang.

## Requirements:
- Golang >1.8

This uses no additional packages to what is provided in Golang.  

Please clone the project to:
`$GOPATH/src/github.com/grimthereaper/gitgazer`

Or if you are a Windows user:
`%GOPATH%/src/github.com/grimthereaper/gitgazer`

## Pre-Ignition:
Make sure to have a github oAuth token, or a github personal token, as rate
limiting cripples this application, as github limits non tokenized request to
60 per hour.

You can do that by following [this guide](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)!

## Running:

CommandLine Arguments
| Command | Type | Default | Description |
| ------- | ---- | ------- | ----------- |
| buffer | bool | `true` | Whether or not to buffer results from github |
| host | string | empty | What host to bind the application to |
| port | int | `8080` | What port to bind the application to |
| token | string | empty | What github personal access token to use (RECOMMENDED) |

To use any of the arguments, they are added by hyphen and the command name, such as...   
`go run main.go -token=ExampleToken`
