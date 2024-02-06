# Simple CLI for Palworld Administrators
 This is a Command Line Interface that an admin of a Palworld server can use to send admin commands through a console/terminal. No need to log into the game!

- [Requirements](#requirements)
- [Setup](#setup)
    1. [Download the source](#1-download-the-source)
    2. [Set Variables](#2-set-variables)
    3. [Compiling](#3-compiling)
- [Usage](#usage)
    - [Command List](#command-list)
## Requirements
 - Go version 1.21 or newer
 - A Palworld server with RCON enabled
 - Know the IP address and RCON port

## Setup
#### 1. Download the source
`git clone https://github.com/JengaMasterG/PalworldAdminCLI.git`

Download the zip from GitHub, and extract it.

#### 2. Set Variables
Open main.go and replace `"Public_IP:Port", "AdminPassword"` with the IP address, port number, and Admin Password of the Palserver.

_If you do not want to compile the source into a binary, continue at the [Usage](#usage) section._

#### 3. Compiling
Open a console/terminal in the foler you saved the source, then run:

Linux:
```
go mod tidy
go build -o palworldcli -v main.go
```

Windows:
```
go mod tidy
go build -o palworldcli.exe -v main.go
```

## Usage
Open a console/terminal in the folder you saved the source, then:

Inside the golang environment, run:

`go run . <command> <args>`

If the package was built:

Linux: `./palworldcli <command> <args>`

Windows: `./palworldcli.exe <command> <args>`

### Command List
    broadcast <"Message_to_send">

Sends a message out as "SYSTEM" to the palword chat.

 _Note: Broadcast message currently broken on Palworld's end. Messages grouped with `""` or ` `` ` will not send the full message. It is recommended to use hyphens or underscores to link words together._

    info

Returns the server information of the Palserver. Information included: Palserver version and name of palserver.

    player ban|kick <steamID>

Bans/Kicks the player from the server. The steamID of the player is required to know. Run `showplayers` to get a list of current logged in players and their steamIDs.

    save

Saves the current state of the Palserver to the server's HDD.

    showplayers

Shows the current online players logged into the Palserver.

    shutdown <timer> <message>

Gracefully shuts down the Palserver. A timer (seconds) is started before the shutdown begins. The message (string) displays the shutdown notice defined by the command.

    shutdown now

Force stops the Palserver. 

_Note: If running the server as a systemd process on Linux, the server will restart. It's recommended to run `systemctl stop <palworld_service_name.service>` after a `shutdown` command was used._

    teleporttome <steamID>
Teleports a player to the admin. The steamID must be known. Run `showplayers` to get a list of current logged in players and their steamIDs. **This is not supported through Palworld's steamCMDs. The admin must be a logged in player to perform this command.**

    teleporttoplayer <steamID>
Teleports an admin to a player. The steamID must be known. Run `showplayers` to get a list of current logged in players and their steamIDs. **This is not supported through Palworld's steamCMDs. The admin must be a logged in player to perform this command.**