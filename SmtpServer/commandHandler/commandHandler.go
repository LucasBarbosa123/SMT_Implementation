package commandHandler

import (
	"fmt"
	"smtpserver/stateManager"
	"smtpserver/utils"
	"strings"
)

// tem q ser usado quando recebermos o comando DATA para ver se excede este tamanho
const maxMsgSize = 36700160 // 35 MB

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
	case "EHLO":
		response = HandleEHLO(stateManager, message, maxMsgSize)
	default:
		response = "500 Syntax error, command unrecognized\r\n"
	}

	return response
}
