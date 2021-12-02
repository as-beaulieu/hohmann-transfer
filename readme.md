# Hohmann Transfer

Inspired by the interview where Elon Musk openly stated that
"a bunch of people will probably die" on the trip to Mars, it remindind me of 
The Oregon Trail I used to play in school. 

So here I present, **The Hohmann Transfer**, named after the elliptical orbit maneuver
used to transfer from one circular orbit to another of different radii 
around a common body on the same plane. 

Launch missions that send crew and passengers from earth to mars, 
handling random events and mishaps which may occur along the way.

*Currently Working* - Colony mission. Send a colony to mars and hope that 
your crew can manage the provisions you have allocated correctly in order
to complete your mission.

# Known Bugs

- Windows executable loops on start screen no matter the input

# TODO

- Add maneuver events
- Add mission completion evaluation after mars landing
- Add passenger consumption of provisions
- Add loss of passengers if insufficient provisions
- Add Satellite mission (*Currently not implemented*)
- Add specific events depending on mission type

After event completion

- High score (File save and read)
- Difficulty (Options screen updating game state)

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