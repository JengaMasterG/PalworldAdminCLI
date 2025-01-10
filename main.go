package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JengaMasterG/palwrldcmdsgo"
)

func main() {

	args := os.Args[1:]
	argsLen := len(args)
	IPAddress := ""
	password := ""
	command := ""
	arg1 := ""
	arg2 := ""
	response := ""

	switch argsLen {
	case 1:
		if args[0] == "man" || args[0] == "help" {
			command = args[0]
		}
	case 3:
		IPAddress = args[0]
		password = args[1]
		command = strings.ToLower(args[2])
		fmt.Println(command)
	case 4:
		IPAddress = args[0]
		password = args[1]
		command = strings.ToLower(args[2])
		arg1 = args[3]
	case 5:
		IPAddress = args[0]
		password = args[1]
		command = strings.ToLower(args[2])
		arg1 = args[3]
		arg2 = args[4]
	default:
	}

	switch command {
	case "player":
		switch arg1 {
		case "ban":
			response = palwrldcmdsgo.BanPlayer(IPAddress, password, arg1)
		case "kick":
			response = palwrldcmdsgo.KickPlayer(IPAddress, password, arg1)
		default:
			response = "Usage:\n player ban <steamID>\n player kick <steamID>\nNote: use the showplayers command to view active player's steamIDs."
		}
	case "broadcast":
		palwrldcmdsgo.Broadcast(IPAddress, password, arg1)
		response = "Message Broadcasted on server."
	case "info":
		response = palwrldcmdsgo.Info(IPAddress, password)
	case "save":
		response = palwrldcmdsgo.Save(IPAddress, password)
	case "showplayers":
		response = palwrldcmdsgo.ShowPlayers(IPAddress, password)
	case "shutdown":
		switch arg1 {
		case "now":
			response = `===STARTING FORCE SHUTDOWN OF SERVER===\n` + palwrldcmdsgo.DoExit(IPAddress, password) + `===FORCE SHUTDOWN COMPLETE=== Please stop the systemd service to stop a server restart.`
		default:
			response = palwrldcmdsgo.Shutdown(IPAddress, password, arg1, arg2) + `Note: Please stop the systemd srevice to stop a server restart.`
		}
	case "teleporttome":
		response = "Can only be done by an admin playing the game."
	case "teleporttoplayer":
		response = "Can only be done by an admin playing the game."
	case "man", "help":
		dat, err := os.ReadFile("man")
		if err != nil {
			panic(err)
		}
		fmt.Print(string(dat), "\n")
	default:
		fmt.Print("palworldcli <IPAddress:RCONPort> <AdminPassword> <command> <args>", "\nUse command man or help for more info.\n")
	}
	fmt.Printf(response)
}
