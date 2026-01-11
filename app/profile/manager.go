package profile

import (
	"encoding/json"
	"fmt"
	"hidtool/app/util"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var List []Profile
var currentProfile Profile

func init() {
	folder := ""
	if util.IsDebug() {
		folder = "build"
	}

	pattern := util.GetPath(folder, "profile*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("Error globbing profiles: %v\n", err)
		return
	}

	var errorLogs []string

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			errorLogs = append(errorLogs, fmt.Sprintf("[%s] Error reading file: %v", filepath.Base(file), err))
			continue
		}

		var singleProfile *ProfileImpl
		if err := json.Unmarshal(data, &singleProfile); err != nil {
			errorLogs = append(errorLogs, fmt.Sprintf("[%s] Load Error: %v", filepath.Base(file), err))
			continue
		}

		if singleProfile.ID == "" {
			baseName := filepath.Base(file)
			singleProfile.ID = strings.TrimSuffix(baseName, filepath.Ext(baseName))
		}
		List = append(List, singleProfile)
	}

	if len(errorLogs) > 0 {
		timestamp := time.Now().Format("20060102_150405")
		logFileName := fmt.Sprintf("error_profile_%s.log", timestamp)
		logPath := util.GetPath(folder, logFileName)

		header := fmt.Sprintf("Profile Load Errors - %s\n=================================\n", time.Now().Format(time.RFC1123))
		content := header + strings.Join(errorLogs, "\n")

		if err := os.WriteFile(logPath, []byte(content), 0644); err != nil {
			fmt.Printf("Failed to write error log to %s: %v\n", logPath, err)
		} else {
			fmt.Printf("Profile errors logged to %s\n", logPath)
		}
	}

	if len(List) > 0 {
		currentProfile = List[0]
	}
}

func SetCurrentProfile(profile Profile) {
	currentProfile = profile
}

func GetCurrentProfile() Profile {
	return currentProfile
}
