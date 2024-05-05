package version

import (
	"fmt"
	"strings"
)

var (
	AppName    string
	AppVersion string
	GoVersion  string

	GitRepoUrl string
	GitTag     string
	GitHash    string

	BuildOS    string
	BuildArch  string
	TargetOS   string
	TargetArch string
	BuildTime  string
	BuildUser  string
)

type Version struct{}

func New() *Version {
	return &Version{}
}

func (v *Version) String() string {
	var bul strings.Builder
	bul.WriteString(fmt.Sprintf("\nAppName:        %s\n", AppName))
	bul.WriteString(fmt.Sprintf("AppVersion:     %s\n", AppVersion))
	bul.WriteString(fmt.Sprintf("GoVersion:      %s\n", GoVersion))
	bul.WriteString(fmt.Sprintf("BuildOS/Arch:   %s/%s\n", BuildOS, BuildArch))
	bul.WriteString(fmt.Sprintf("TargetOS/Arch:  %s/%s\n", TargetOS, TargetArch))
	bul.WriteString(fmt.Sprintf("BuildTime:      %s\n", BuildTime))
	bul.WriteString(fmt.Sprintf("BuildUser:      %s\n", BuildUser))
	bul.WriteString(fmt.Sprintf("GitTag/Branch:  %s\n", GitTag))
	bul.WriteString(fmt.Sprintf("GitHash:        %s\n", GitHash))
	bul.WriteString(fmt.Sprintf("GitRepoUrl:     %s\n", GitRepoUrl))
	return bul.String()
}
