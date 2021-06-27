package main

import (
	"fmt"
	"os"
	"path"

	"github.com/akshaybabloo/go-cli-template/cmd"
	"github.com/akshaybabloo/go-cli-template/model"
	"github.com/akshaybabloo/go-cli-template/pkg/factory"
	"github.com/akshaybabloo/go-cli-template/pkg/io"
	"github.com/akshaybabloo/go-cli-template/pkg/update"
)

const Version = "0.0.0-DEV"
const BuildDate = "date"

func main() {
	f := factory.New()

	updateMessageChan := make(chan *update.ReleaseInfo)
	go func(v string, f *factory.Factory) {
		versionInfo, _ := update.CheckForNewRex(v, f)
		updateMessageChan <- versionInfo
	}(Version, f)

	// create `<user directory>/.config/rex/' path
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	rexFolder := path.Join(home, ".config", "rex")
	_, err = os.Stat(rexFolder)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(rexFolder, 0777); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	// creates `<user directory>/.config/rex/config.yaml'
	// and add default values
	globalConfig, err := f.Config().GetGlobalConfigPath()
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(globalConfig)
	if err != nil {
		if os.IsNotExist(err) {
			marshalNew, err := model.MarshalNew()
			if err != nil {
				panic(err)
			}
			err = io.WriteToFile(globalConfig, marshalNew)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	rootCmd := cmd.NewRootCmd(f, Version, BuildDate)
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cs := f.Colours()
	newRelease := <-updateMessageChan
	if newRelease != nil {
		fmt.Printf("\n\n%s %s → %s\n",
			cs.GreenString("A new release of rex is available:"),
			cs.GreenString(Version),
			cs.GreenString(newRelease.Version))
		fmt.Printf("%s\n\n", cs.GreenString(newRelease.URL))
	}
}
