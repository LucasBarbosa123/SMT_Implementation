package commandHandler

import (
	"smtpserver/stateManager"
	"strings"
)

func HandleHELO(stateManager *stateManager.StateManager, message string) string {

	msgParts := strings.Fields(message)
	if len(msgParts) > 2 {
		//because domains don't have spaces in it so the command should look like "HELO" or "HELO domain.net"
		return "500 Syntax error, command unrecognized\r\n"
	}

	if len(msgParts) == 2 {
		//goes after the domain that needs to be the first thing after the " "
		stateManager.Domain = msgParts[1]
	} else {
		stateManager.Domain = ""
	}

	stateManager.NextState()
	return "250 Hello\r\n"
}
