package forward

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

// -----------------------------------------------------------------------------

// TemplateTaskNew is the template for the `fwd task new` command.
const TemplateTaskNew = `
# Please enter the name and description (optional) of your new task, separated
# by an empty line. Names longer than 80 characters will be truncated.
#
# Lines starting with '#' will be ignored, and an empty message aborts the
# creation of the Forward project.
`

// -----------------------------------------------------------------------------

// TaskNew implements the `fwd task new` command.
func TaskNew(c *cli.Context) {
	// get title and description
	t, d, err := TitleAndDescription(Edit("task-new", TemplateTaskNew))
	if err != nil {
		log.Fatal(err)
	}
	if t == "" { // NOTE: empty message -> abort
		return
	}

	// add task:new commit
	tid, err := LatestTaskID()
	if err != nil {
		log.Fatal(err)
	}
	t = fmt.Sprintf("[forward] task:new - (#%d) %s", tid+1, t)
	cmd := []string{"commit", "--allow-empty", "--file", "-"}
	if err := GitExec(cmd, []byte(t+"\n\n"+d)); err != nil {
		log.Fatal(err)
	}
}

// -----------------------------------------------------------------------------

type task struct {
	ID                          int64
	Author, Email               string
	Title, Description          string
	Status                      string
	CDate, CModified, SModified time.Time
}

// Parse extracts task information from a commit.
//
// `commit` must have format `--format=%aN;;;%aE;;;%at;;;%s;;;%b`.
func (t *task) Parse(commit string) error {
	strs := strings.Split(commit, ";;;")
	if len(strs) != 5 {
		return fmt.Errorf("'%v': wrong format", commit)
	}
	tid, err := ExtractTaskID([]byte(strs[3]))
	if err != nil {
		return err
	}
	t.ID = tid
	t.Author = strs[0]
	t.Email = strs[1]
	msg, err := ExtractMessage([]byte(strs[3]))
	if err != nil {
		return err
	}
	t.Title = msg
	t.Description = strs[4]
	ts, err := strconv.ParseInt(strs[2], 10, 64)
	if err != nil {
		return err
	}
	t.CDate = time.Unix(ts, 0)
	t.CModified = t.CDate
	t.SModified = t.CDate
	t.Status = "backlog"
	return nil
}

// String implements the `stringer` interface.
func (t task) String() string {
	template := "Task #%d - %s\n" +
		"  author: %s <%s>\n" +
		"  created: %v (last-modified: %v)\n" +
		"  status: %v (last-modified: %v)"
	return fmt.Sprintf(
		template,
		t.ID, t.Title,
		t.Author, t.Email,
		t.CDate.UTC().Format(time.RFC3339), t.CModified.UTC().Format(time.RFC3339),
		t.Status, t.SModified.UTC().Format(time.RFC3339),
	)
}

// TaskList implements the `fwd task list` command.
func TaskList(c *cli.Context) {
	cmd := []string{"log", "--grep", `\[forward\] task:`}
	cmd = append(cmd, "--format=%aN;;;%aE;;;%at;;;%s;;;")
	output, err := GitOutput(cmd)
	if err != nil {
		log.Fatal(err)
	}

	tasks := make([]*task, 0, 1024)
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		t := &task{}
		if err := t.Parse(scanner.Text()); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, t := range tasks {
		fmt.Println(t)
	}
}
