package argv

// Args provides a way to store both a list of provided arguments
// without keys in addition to the key'ed values.
type Args struct {
	Nokeys []string
	Keys   map[string]string
}

// Parse takes in a slice of arguments (probably from os.Args) and
// parses the arguments into an easier to use map[string]string.
func Parse(rawArgs []string) Args {
	var allTheArgs Args
	var nokeyArgsIndices []int
	flags := make(map[int]string)
	allTheArgs.Keys = make(map[string]string)
	for i, arg := range rawArgs {
		if arg[0:2] == "--" {
			flags[i] = arg[2:]
		} else if arg[0:1] == "-" {
			flags[i] = arg[1:]
		} else {
			nokeyArgsIndices = append(nokeyArgsIndices, i)
		}
	}

	for indexOfKey, key := range flags {
		if indexOfKey+1 < len(rawArgs) {
			allTheArgs.Keys[key] = checkAndReturnValue(rawArgs, indexOfKey)
		} else {
			allTheArgs.Keys[key] = ""
		}
	}

	for _, ind := range nokeyArgsIndices {
		if ind > 0 {
			if getValidValue(rawArgs[ind-1]) != "" {
				allTheArgs.Nokeys = append(allTheArgs.Nokeys, rawArgs[ind])
			}
		}
	}

	return allTheArgs
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
