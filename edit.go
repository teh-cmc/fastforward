package forward

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// -----------------------------------------------------------------------------

// Edit starts the user's favorite editor and returns the output.
func Edit(prefix string, templates ...string) ([]byte, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return nil, fmt.Errorf("$EDITOR isn't set")
	}
	path, err := exec.LookPath(editor)
	if err != nil {
		return nil, err
	}

	f, err := ioutil.TempFile(os.TempDir(), fmt.Sprintf("forward.%s-", prefix))
	if err != nil {
		return nil, err
	}
	defer func() {
		// NOTE: file removal might fail, not much we can do though.
		if err := os.Remove(f.Name()); err != nil {
			log.Println(err)
		}
	}()

	for _, t := range templates {
		if _, err := f.WriteString(t); err != nil {
			return nil, err
		}
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
	return ioutil.ReadAll(f)
}

// TitleAndDescription extracts the title and description from `b`.
//
// The title will be truncated if >80 characters.
//
// Example:
//    t, d, err := TitleAndDescription(Edit("init", TemplateInit))
func TitleAndDescription(b []byte, e error) (t string, d string, err error) {
	if e != nil {
		err = e
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(txt, "#") || len(txt) == 0 {
			continue
		}
		if t == "" {
			t = txt
			if len(t) > 80 {
				t = t[:80]
			}
			continue
		}
		d = txt
		break
	}
	err = scanner.Err()
	return
}
