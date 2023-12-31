package src

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Sh struct {
	CustomSt bool
	Stdin    bool
	Stdout   bool
	Stderr   bool
}

func (sh Sh) Getprefix() []string {
	prefix := make([]string, 2)
	if runtime.GOOS == "windows" {
		prefix[0] = "cmd"
		prefix[1] = "/C"
	}
	if runtime.GOOS == "linux" {
		prefix[0] = "sh"
		prefix[1] = "-c"
	}
	return prefix
}

func (sh Sh) Cmd(input string) error {
	prefix := sh.Getprefix()
	cmd := exec.Command(prefix[0], prefix[1], input)
	if sh.CustomSt {
		if sh.Stderr {
			cmd.Stderr = os.Stderr
		}
		if sh.Stdin {
			cmd.Stdin = os.Stdin
		}
		if sh.Stdout {
			cmd.Stdout = os.Stdout
		}
	} else {
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
	}
	err := cmd.Run()
	return err
}

func (sh Sh) Out(input string) (string, error) {
	prefix := sh.Getprefix()
	cmd := exec.Command(prefix[0], prefix[1], input)
	out, err := cmd.Output()
	return string(out), err
}

var sh = Sh{}

func CheckDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func CheckRclone() bool {
	var binName string = "rclone"
	if runtime.GOOS == "windows" {
		binName = binName + ".exe"
	}
	_, err := sh.Out(fmt.Sprintf("%v version", binName))
	if err != nil {
		return false
	} else {
		return true
	}
}

type yamlfile struct {
	VscodeFolder string `yaml:"vscode-folder"`
	BackupFolder string `yaml:"backup-folder"`
	Remotefolder string `yaml:"remote-folder"`
}

var ConfigData = getYamldata()

func getYamldata() yamlfile {
	yamldata := yamlfile{}
	if !CheckDir("config.yml") {
		NewJsonFile()
	}
	file, _ := os.ReadFile("config.yml")
	yaml.Unmarshal(file, &yamldata)
	return yamldata
}

func NewJsonFile() {
	jsonfile := yamlfile{}
	file, _ := os.Create("config.yml")
	defer file.Close()
	data, _ := yaml.Marshal(jsonfile)
	file.WriteString(string(data))
}

func Rclone() {
	if !CheckRclone() {
		fmt.Println("Rclone is not installed")
		return
	}
	if ConfigData.BackupFolder == "" || ConfigData.VscodeFolder == "" {
		fmt.Println("Rclone backupfolder or vscode folder is <null>")
		return
	}
	var win, command string
	if runtime.GOOS == "windows" {
		win = ".exe"
	}
	if os.Args[1] == "save" {
		command = fmt.Sprintf(
			"rclone%v sync %v %v -L -P",
			win,
			ConfigData.VscodeFolder,
			ConfigData.BackupFolder,
		)
	}
	if os.Args[1] == "restore" {
		command = fmt.Sprintf(
			"rclone%v sync %v %v -L -P",
			win,
			ConfigData.BackupFolder,
			ConfigData.VscodeFolder,
		)
	}
	if os.Args[1] == "download" {
		if ConfigData.Remotefolder == "" {
			fmt.Println("Remote dir is <null>")
		}
		command = fmt.Sprintf("rclone%v sync %v %v -L -P", win, ConfigData.Remotefolder, ConfigData.BackupFolder)
	}
	if os.Args[1] == "upload" {
		if ConfigData.Remotefolder == "" {
			fmt.Println("Remote dir is <null>")
		}
		command = fmt.Sprintf("rclone%v sync %v %v -L -P", win, ConfigData.BackupFolder, ConfigData.Remotefolder)
	}
	sh.Cmd(command)
}
