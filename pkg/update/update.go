package update

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go-version"
	"gopkg.in/yaml.v3"

	"github.com/akshaybabloo/go-cli-template/pkg/factory"
)

type ReleaseInfo struct {
	Version     string    `json:"tag_name"`
	URL         string    `json:"html_url"`
	PublishedAt time.Time `json:"published_at" yaml:"published_at"`
}

type StateEntry struct {
	CheckedForUpdateAt time.Time   `yaml:"checked_for_update_at"`
	LatestRelease      ReleaseInfo `yaml:"latest_release"`
}

// CheckForNewRex checks for new version and returns its information if one exists.
func CheckForNewRex(v string, f *factory.Factory) (*ReleaseInfo, error) {

	path, err := f.Config().StatePath()
	if err != nil {
		return nil, err
	}

	stateEntry, _ := getStateEntry(path)
	if stateEntry != nil && time.Since(stateEntry.CheckedForUpdateAt).Hours() < 24 {
		return nil, nil
	}

	newRelease, err := latestRelease()
	if err != nil {
		return nil, err
	}

	err = setStateEntry(path, time.Now(), *newRelease)
	if err != nil {
		return nil, err
	}

	currentVersion, err := version.NewVersion(v)
	if err != nil {
		return nil, err
	}

	newVersion, err := version.NewVersion(newRelease.Version)
	if err != nil {
		return nil, err
	}

	if currentVersion.LessThan(newVersion) {
		return newRelease, nil
	}

	return nil, nil
}

func latestRelease() (*ReleaseInfo, error) {
	var releaseInfo ReleaseInfo

	get, err := http.Get("https://api.github.com/repos/akshaybabloo/rex/releases/latest")
	if err != nil {
		return nil, err
	}
	defer get.Body.Close()

	all, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(all, &releaseInfo)
	if err != nil {
		return nil, err
	}

	return &releaseInfo, nil
}

func getStateEntry(stateFilePath string) (*StateEntry, error) {
	content, err := ioutil.ReadFile(stateFilePath)
	if err != nil {
		return nil, err
	}

	var stateEntry StateEntry
	err = yaml.Unmarshal(content, &stateEntry)
	if err != nil {
		return nil, err
	}

	return &stateEntry, nil
}

func setStateEntry(stateFilePath string, t time.Time, r ReleaseInfo) error {
	data := StateEntry{CheckedForUpdateAt: t, LatestRelease: r}
	content, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	_ = ioutil.WriteFile(stateFilePath, content, 0600)

	return nil
}
