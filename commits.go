package forward

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// -----------------------------------------------------------------------------

// Commit represents a commit.
type Commit struct {
	msg *CommitMessage
	md  *CommitMetadata
}

// -----------------------------------------------------------------------------

// CommitMetadata represents the metadata of a commit.
type CommitMetadata struct {
	taskID                      int64
	author, authorEmail         string
	assignee, assigneeEmail     string
	status                      string
	cDate, cModified, sModified time.Time
}

// NewMetadata parses `b` and returns a new CommitMetadata.
func NewMetadata(b []byte) (*CommitMetadata, error) { return nil, nil }

// -----------------------------------------------------------------------------

var (
	regexCommitCommandTitle *regexp.Regexp
	regexCommitAttributes   *regexp.Regexp
)

func init() {
	regexCommitCommandTitle = regexp.MustCompile(`^\[FastForward\] ((?:[a-z]+:?)+) > (.*)$`)
	regexCommitAttributes = regexp.MustCompile(`^([a-zA-Z0-9\-_]+):((?:[a-zA-Z0-9\-_]+,?)*)$`)
}

// Commitable exposes methods to retrieve command names and commit templates.
type Commitable interface {
	Command() string
	Template() []byte
}

// CommitMessage represents the message of a commit.
type CommitMessage struct {
	command            string
	title, description string
	attributes         map[string][]string
}

// Bytes returns the byte representation of a CommitMessage.
func (cm CommitMessage) Bytes() []byte {
	var msg string
	msg += fmt.Sprintf("[FastForward] %s > %s", cm.command, cm.title) +
		"\n\n" + cm.description + "\n\n"
	for a, vs := range cm.attributes {
		msg += a + ":" + strings.Join(vs, ",") + "\n"
	}
	return []byte(msg)
}

// NewMessage parses `b` and returns a new CommitMessage.
func NewMessage(b []byte) (*CommitMessage, error) {
	cm := &CommitMessage{}
	cm.attributes = make(map[string][]string)
	scanner := bufio.NewScanner(bytes.NewReader(b))
	i := 0
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())

		// comment
		if bytes.HasPrefix(line, []byte("#")) {
			continue
		}

		// empty line
		if len(line) <= 0 {
			i++
			continue
		}

		// i == 0 -> command + title
		if i == 0 {
			if cm.title != "" {
				return nil, fmt.Errorf("multi-lines title are forbidden")
			}
			if !regexCommitCommandTitle.Match(line) {
				return nil, fmt.Errorf("'%s': invalid command/title", line)
			}
			values := regexCommitCommandTitle.FindStringSubmatch(string(line))
			cm.command = values[1]
			cm.title = values[2]
			if len(cm.title) > 80 {
				cm.title = cm.title[:80]
			}
			continue
		}

		// i == 1 -> description
		if i == 1 {
			cm.description += string(line)
			continue
		}

		// i >= 2 -> attributes
		if i >= 2 {
			if !regexCommitAttributes.Match(line) {
				return nil, fmt.Errorf("'%s': invalid attributes", line)
			}
			attr := regexCommitAttributes.FindStringSubmatch(string(line))
			cm.attributes[attr[1]] = strings.Split(attr[2], ",")
		}
	}

	// empty commit message, abort
	if cm.title == "" {
		return nil, nil
	}

	return cm, scanner.Err()
}

// EditMessage starts the user's favorite editor, parses the output and returns
// a new CommitMessage.
func EditMessage(c Commitable) (*CommitMessage, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return nil, fmt.Errorf("$EDITOR isn't set")
	}
	path, err := exec.LookPath(editor)
	if err != nil {
		return nil, err
	}

	f, err := ioutil.TempFile(os.TempDir(), "fastforward")
	if err != nil {
		return nil, err
	}
	defer func() {
		// NOTE: file removal might fail, not much we can do though
		if err := os.Remove(f.Name()); err != nil {
			log.Println(err)
		}
	}()

	if _, err := f.Write(c.Template()); err != nil {
		return nil, err
	}

	cmd := exec.Command(path, f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}
	output, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	output = []byte(fmt.Sprintf("[FastForward] %s > ", c.Command()) + string(output))

	return NewMessage(output)
}
