package commandHandler

import (
	"fmt"
	"smtpserver/stateManager"
	"smtpserver/utils"
	"strings"
)

func HandleCommand(stateManager *stateManager.StateManager, message string) string {
	response := ""
	message = strings.TrimSpace(message)
	fmt.Println("Message received:", message)

	msgParts := strings.Fields(message)
	msg := msgParts[0]

	if utils.ContainsString(stateManager.CurrentState.PossibleCommands, message) {
		return "500 Syntax error, command unrecognized or unvalid\r\n"
	}

	switch msg {
	case "HELO":
		response = HandleHELO(stateManager, message)
	default:
		response = "500 Syntax error, command unrecognized\r\n"
	}

	return response
}
