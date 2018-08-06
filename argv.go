package argv

// Parse takes in a slice of arguments (probably from os.Args) and
// parses the arguments into an easier to use map[string]string.
func Parse(rawArgs []string) map[string]string {
	flags := make(map[int]string)
	args := make(map[string]string)
	for i, arg := range rawArgs {
		if arg[0:2] == "--" {
			flags[i] = arg[2:]
		} else if arg[0:1] == "-" {
			flags[i] = arg[1:]
		}
	}

	for indexOfKey, key := range flags {
		if indexOfKey+1 < len(rawArgs) {
			args[key] = checkAndReturnValue(rawArgs, indexOfKey)
		} else {
			args[key] = ""
		}
	}

	return args
}

func checkAndReturnValue(rawArgs []string, keyIndex int) string {
	return getValidValue(rawArgs[keyIndex+1])
}

func getValidValue(val string) string {
	if val[0:2] == "--" || val[0:1] == "-" {
		return ""
	}
	return val
}
