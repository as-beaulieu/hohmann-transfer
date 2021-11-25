# Hohmann Transfer

# Known Bugs

- Windows executable loops on start screen no matter the input

# Building an executable

## Windows
`GOOS=windows GOARCH=386 go build -o hohmann.exe main.go`

## MacOS
`GOOS=darwin GOARCH=amd64 go build -o hohmann`

# Architecture

## Intended Flow Path

                    /------\    /----------\
                      Main       Game State
                    \------/    \----------/
                        |           |
                        |           |
                    /-------\       |
                     Phases         |   Game State
                    \-------/       |   Dependency
                        |           |   Injected
                        |           |
                    /-------\       |
                      Events        X   | Display |
                    \-------/           | Data    |
                        |                   |
                        |                   |
                    /-------\               |
                     Graphics               X
                    \-------/