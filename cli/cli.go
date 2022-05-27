package cli

import "os"

func GetArg(args string) string {

	if args == "cmd" {
		return getCommandLineArgs()[1]
	}

	if args == "option" {
		return getCommandLineArgs()[2]
	}

	return "none"
}

func getCommandLineArgs() []string {
	return os.Args
}
