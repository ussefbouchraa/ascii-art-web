package asciiart


import(
	"bufio"
	"os"
	S "strings"
)



var (
	_map  = make(map[int][8]string)
	lines = [8]string{}
)


func InitMap(banner string) {
	var file *os.File
	var err error

	switch banner {
		
	case "standard", "standard.txt":
		file, err = os.Open("Banners/standard.txt")
	case "shadow", "shadow.txt":
		file, err = os.Open("Banners/shadow.txt")
	case "thinkertoy", "thinkertoy.txt":
		file, err = os.Open("Banners/thinkertoy.txt")
	default:
		os.Stderr.WriteString("Err: Invalid Argument [BANNER]: " + banner + "\n")
		os.Exit(0)
	}

	if err != nil {
		os.Stderr.WriteString("Err: " + err.Error() + "\n")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(file)
	// to skip first empty line
	if banner == "shadow" || banner == "thinkertoy" ||
		banner == "shadow.txt" || banner == "thinkertoy.txt" {
		scanner.Scan()
	}
	// insert data on the map
	for i := 32; i < 127; i++ {
		_map[i] = InsertValue(scanner)
	}

	defer file.Close()
}

func Storing(inp string) string{
	res := ""
	if inp == "\\n" {
		return res +"\n"
	}

	SplArgs := S.Split(inp, "\\n")

	// Fix Multiple "\n"
	if IsOnlyNewLine(SplArgs) {
		for i := 0; i < len(SplArgs)-1; i++ {
			 res += "\n"
		}
		return res
	}
	// applying "\n"
	for _, arg := range SplArgs {
		if arg == "" {
			 res += "\n"
			continue
		}
		// Storing Data
		for _, val := range arg {
			for i := 0; i < 8; i++ {
				lines[i] += _map[int(val)][i]
			}
		}
			// Printing Data
			for i := 0; i < 8; i++ {
				// F.Println(lines[i])
				res += lines[i] +"\n"
				lines[i] = ""
			}

	}
	return res
}
