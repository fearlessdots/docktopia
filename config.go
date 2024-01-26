package main

//
//// IMPORTS
//

import (
	// Modules in GOROOT
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	// External modules
	//color "github.com/gookit/color"
)

//
//// PROGRAM CONFIGURATION
//

type Program struct {
	name                    string
	nameAscii               string
	version                 string
	exec                    string
	shortDescription        string
	longDescription         string
	defaultShell            string
	userDataDir             string
	userCratesDir           string
	userTemplatesDir        string
	userTargetsTemplatesDir string
	userCratesTemplatesDir  string
	indentLevel             int
}

func getDefaultShellAbsolutePath(shellName string) string {
	// Get shell absolute path using `which`
	cmd := exec.Command("which", shellName)

	output, err := cmd.CombinedOutput()
	outputString := string(output)
	outputString = strings.TrimLeft(outputString, "\n")
	outputString = strings.TrimRight(outputString, "\n")

	if err != nil {
		showError(fmt.Sprintf("Failed to get absolute path to default shell '%v' -> %v", shellName, outputString), 0)
		finishProgram(1)
	}

	return outputString
}

func initializeDefaultProgram(customUserDataDir string) Program {
	// PROGRAM NAME
	programName := "docktopia"

	// PROGRAM NAME ASCII
	programNameAscii := "     _            _    _              _\n" +
		"  __| | ___   ___| | _| |_ ___  _ __ (_) __ _\n" +
		" / _  |/ _ \\ / __| |/ / __/ _ \\| '_ \\| |/ _` | \n" +
		"| (_| | (_) | (__|   <| || (_) | |_) | | (_| | \n" +
		" \\__,_|\\___/ \\___|_|\\_\\\\__\\___/| .__/|_|\\__,_| \n" +
		"                               |_|\n"

	// PROGRAM VERSION
	programVersion := "0.1.0"

	// PROGRAM EXEC
	programExec := os.Args[0] // Path for program executable

	// DESCRIPTIONS (SHORT AND LONG)
	programShortDescription := "A practical toolkit crafted specifically for effectively managing Docker-driven endeavors. Leveraging template functionalities, docktopia facilitates quick generation of crates and targets, leading to enhanced productivity."
	programLongDescription := fmt.Sprintf("%v is a practical toolkit crafted specifically for effectively administering Docker-driven endeavors. Leveraging template functionalities, docktopia facilitates quick generation of crates and targets, leading to enhanced productivity.", programName)

	// DEFAULT SHELL
	programDefaultShellName := "sh" // Should work on all Unix systems (Linux, Android, ...)
	programDefaultShellPath := getDefaultShellAbsolutePath(programDefaultShellName)

	// USER DIRECTORIES
	var userDataDir string
	if customUserDataDir != "" {
		userDataDir = customUserDataDir
	} else {
		userDataDir = getCurrentUserHomeDir(Program{indentLevel: 0}) + "/" + programName
	}

	userCratesDir := userDataDir + "/crates"
	userTemplatesDir := userDataDir + "/templates"
	userTargetsTemplatesDir := userTemplatesDir + "/targets"
	userCratesTemplatesDir := userTemplatesDir + "/crates"

	// INDENT LEVEL
	indentLevel := 0

	return Program{
		name:                    programName,
		nameAscii:               programNameAscii,
		version:                 programVersion,
		exec:                    programExec,
		shortDescription:        programShortDescription,
		longDescription:         programLongDescription,
		defaultShell:            programDefaultShellPath,
		userDataDir:             userDataDir,
		userCratesDir:           userCratesDir,
		userTemplatesDir:        userTemplatesDir,
		userTargetsTemplatesDir: userTargetsTemplatesDir,
		userCratesTemplatesDir:  userCratesTemplatesDir,
		indentLevel:             indentLevel,
	}
}

func getRootDirectory() string {
	// Check if the "PREFIX" environment variable is set
	prefix := os.Getenv("PREFIX")
	if prefix != "" {
		// Verify if the given prefix string ends with the directory suffix "/usr".
		// If it does, the program will proceed to strip it from the prefix.
		if strings.HasSuffix(prefix, "/usr") {
			return strings.TrimSuffix(prefix, "/usr")
		}
		return prefix
	} else {
		return "/"
	}
}

func displayProgramInfo(program Program) {
	showText(program.nameAscii, program.indentLevel)

	showText("Version: "+green.Sprintf(program.version), program.indentLevel)

	space()

	showText("Running on "+lightCopper.Sprintf(runtime.GOOS+"/"+runtime.GOARCH)+". Built with "+runtime.Version()+" using "+runtime.Compiler+" as compiler.", program.indentLevel)

	space()

	showText("This program is licensed under the GNU General Public License v3.0 (GPL-3.0).\nPlease refer to the LICENSE file for more information.", program.indentLevel)
}
