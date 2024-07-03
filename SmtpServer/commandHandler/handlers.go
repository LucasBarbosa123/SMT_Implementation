package commandHandler

import (
	"smtpserver/stateManager"
	"strconv"
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
		// stateManager.Domain = ""
		return "500 Syntax error, missing an argument\r\n"
	}

	stateManager.NextState(msgParts[0])
	return "250 Hello.Fodase is my domain name\r\n"
}

func HandleEHLO(stateManager *stateManager.StateManager, message string, maxMsgSize int) string {
	msgParts := strings.Fields(message)
	response := HandleHELO(stateManager, message)

	resCode := strings.Fields(response)[0]
	if resCode != "250" {
		return response
	}

	response += "250 SIZE " + strconv.Itoa(maxMsgSize) + "\r\n"

	stateManager.NextState(msgParts[0])
	return response
}

func HandleMAILFROM(stateManager *stateManager.StateManager, message string, maxMsgSize int) string {
	msgParts := strings.Fields(message)
	if msgParts[1] != "FROM" {
		return "500 Syntax error, command unrecognized or unvalid\r\n"
	}

	return "250 Ok"
}
