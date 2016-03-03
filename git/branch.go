package git

import (
	"bufio"
	"bytes"
	"log"
)

// -----------------------------------------------------------------------------

// BranchType represents the type of the `Branch` command.
type BranchType int

const (
	// BranchTypeNew creates a new branch
	BranchTypeNew BranchType = iota
	// BranchTypeSwitch switches to an existing branch
	BranchTypeSwitch BranchType = iota
	// BranchTypeCurrent gets the name of the current branch
	BranchTypeCurrent BranchType = iota
)

// Branch implements `git checkout` & `git branch` commands.
type Branch struct {
	typ    BranchType
	branch string
}

// NewBranch returns a new `Branch` command of type `t`.
func NewBranch(t BranchType, branch string) *Branch {
	return &Branch{typ: t, branch: branch}
}

// -----------------------------------------------------------------------------

// AllowAutoCheckout always returns false.
func (b Branch) AllowAutoCheckout() bool { return false }

// Input always returns `nil`.
func (b Branch) Input() []byte { return nil }

// Command returns a `git checkout` or `git branch` command.
func (b Branch) Command() []string {
	switch b.typ {
	case BranchTypeNew:
		return []string{"checkout", "-b", b.branch}
	case BranchTypeSwitch:
		return []string{"checkout", b.branch}
	case BranchTypeCurrent:
		return []string{"branch"}
	}
	return nil
}

// Transform parses the `output` to find the current branch.
func (b Branch) Transform(output []byte) []byte {
	if b.typ != BranchTypeCurrent {
		return output
	}

	return branchExtractCurrent(output)
}

func branchExtractCurrent(b []byte) []byte {
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		txt := scanner.Bytes()
		if bytes.HasPrefix(txt, []byte("* ")) {
			return txt[2:]
		}
	}
	log.Fatal("couldn't determine current branch")
	return nil
}
