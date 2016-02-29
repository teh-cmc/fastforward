package forward

import (
	"regexp"
	"strconv"
)

// -----------------------------------------------------------------------------

var (
	regexExtractID      *regexp.Regexp
	regexExtractMessage *regexp.Regexp
)

func init() {
	regexExtractID = regexp.MustCompile(`^\[forward\] (?:.+:.+) -(?: \(#(\d+)\)).*`)
	regexExtractMessage = regexp.MustCompile(`^\[forward\] (?:.+:.+) -(?: \(#\d+\))? (.*)`)
}

// -----------------------------------------------------------------------------

// LatestTaskID returns the ID of the last task created.
func LatestTaskID() (int64, error) {
	cmd := []string{
		"log", "--grep", `\[forward\] task:new`, "--format=%s", "-n", "1",
	}
	output, err := GitOutput(cmd)
	if err != nil {
		return 0, err
	}
	return ExtractTaskID(output)
}

// ExtractTaskID extracts the task ID from a commit message.
//
// `commit` must have format `--format=%s`.
func ExtractTaskID(commit []byte) (int64, error) {
	sms := regexExtractID.FindSubmatch(commit)
	if len(sms) < 2 {
		return 0, nil
	}
	i, err := strconv.ParseInt(string(sms[1]), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ExtractMessage extracts the actual message from a commit message.
//
// `commit` must have format `--format=%s`.
func ExtractMessage(commit []byte) (string, error) {
	sms := regexExtractMessage.FindSubmatch(commit)
	if len(sms) < 2 {
		return "", nil
	}
	return string(sms[1]), nil
}
