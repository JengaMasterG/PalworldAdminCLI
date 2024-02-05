package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JengaMasterG/palwrldcmdsgo"
)

var IPAddress, password = "Public_IP:Port", "AdminPassword"

func main() {

	args := os.Args
	command := ""
	subcommand := ""
	response := ""

	if len(args) > 1 {
		command = strings.ToLower(os.Args[1])
	}

	switch command {
	case "player":
		if len(args) > 2 {
			subcommand = strings.ToLower(os.Args[2])
		}
		switch subcommand {
		case "ban":
			response = palwrldcmdsgo.BanPlayer(IPAddress, password, os.Args[3])
		case "kick":
			response = palwrldcmdsgo.KickPlayer(IPAddress, password, os.Args[3])
		default:
			response = "Usage:\n player ban <steamID>\n player kick <steamID>\nNote: use the showplayers command to view active player's steamIDs."
		}
	case "broadcast":
		palwrldcmdsgo.Broadcast(IPAddress, password, os.Args[2])
		response = "Message Broadcasted on server."
	case "info":
		response = palwrldcmdsgo.Info(IPAddress, password)
	case "save":
		response = palwrldcmdsgo.Save(IPAddress, password)
	case "showplayers":
		response = palwrldcmdsgo.ShowPlayers(IPAddress, password)
	case "shutdown":
		if len(args) > 2 {
			subcommand = strings.ToLower(os.Args[2])
		}
		switch subcommand {
		case "now":
			response = `===STARTING FORCE SHUTDOWN OF SERVER===\n` + palwrldcmdsgo.DoExit(IPAddress, password) + `===FORCE SHUTDOWN COMPLETE=== Please stop the systemd service to stop a server restart.`
		default:
			response = palwrldcmdsgo.Shutdown(IPAddress, password, os.Args[2], os.Args[3]) + `\nNote: Please stop the systemd srevice to stop a server restart.`
		}
	case "teleporttome":
		response = "Can only be done by an admin playing the game."
	case "teleporttoplayer":
		response = "Can only be done by an admin playing the game."
	default:
		dat, err := os.ReadFile("man")
		if err != nil {
			panic(err)
		}
		fmt.Print(string(dat))
	}
	fmt.Printf(response)
}
