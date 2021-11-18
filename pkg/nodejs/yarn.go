package nodejs

import (
	"regexp"
	"strings"
)

func YarnDeps(yarnLockContents string) map[string]string {
	deps := make(map[string]string)
	lines := strings.Split(yarnLockContents, "\n")
	startOfEntries := determineStartsOfEntries(lines)
	for _, startOfEntry := range startOfEntries {
		depName := parseName(lines[startOfEntry])
		depVersion := parseVersion(lines[startOfEntry + 1])
		deps[depName] = depVersion
	}
	return deps
}

func determineStartsOfEntries(yarnLockLines []string) []int {
	var startsOfEntries []int
	previousLineisUseful := false
	for lineNr := range yarnLockLines {
		if isUseful(yarnLockLines[lineNr]) && !previousLineisUseful {
			startsOfEntries = append(startsOfEntries, lineNr)
			previousLineisUseful = true
		} else if !isUseful(yarnLockLines[lineNr]){
			previousLineisUseful = false
		}
	}

	return startsOfEntries
}

func isUseful(line string) bool {
	return !(startsWithHash(line) || isEmpty(line))
}

func isEmpty(line string) bool {
	return strings.TrimSpace(line) == ""
}

func startsWithHash(line string) bool {
	return strings.HasPrefix(line, "#")
}

func parseName(line string) string {
	regex := regexp.MustCompile(`^"?(?P<pkgname>.*)@[^~]?.*$`)
	matches := regex.FindStringSubmatch(line)
	pkgnameIndex := regex.SubexpIndex("pkgname")
	return matches[pkgnameIndex]
}

func parseVersion(line string) string {
	regex := regexp.MustCompile(`.*"(?P<pkgversion>.*)"$`)
	matches := regex.FindStringSubmatch(line)
	pkgversionIndex := regex.SubexpIndex("pkgversion")
	return matches[pkgversionIndex]
}

