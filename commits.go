package forward

import (
	"log"
	"regexp"
	"strconv"
)

// -----------------------------------------------------------------------------

var regexExtractID *regexp.Regexp

func init() {
	regexExtractID = regexp.MustCompile(`\(#(\d+)\)`)
}

// -----------------------------------------------------------------------------

// LatestTaskID returns the ID of the last task created.
func LatestTaskID() (int64, error) {
	cmd := []string{
		"log", "--grep", `\[forward\] task:new`, "--oneline", "-n", "1",
	}
	output, err := GitOutput(cmd)
	if err != nil {
		log.Fatal(err)
	}
	return ExtractTaskID(output)
}

// ExtractTaskID extracts the task ID from a commit message.
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
