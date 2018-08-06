package argv

import (
	"testing"
)

func TestParse(t *testing.T) {
	args := []string{"./mycommand", "-abc", "123", "--testing", "omg", "-haha", "--haha2"}
	myArgs := Parse(args[1:])

	if myArgs["abc"] != "123" {
		t.Error("'abc' key was not parsed correctly.")
	}

	if myArgs["testing"] != "omg" {
		t.Error("'testing' key was not parsed correctly.")
	}

	if myArgs["haha"] != "" {
		t.Error("'haha' key was not parsed correctly.")
	}

	if myArgs["haha2"] != "" {
		t.Error("'haha2' key was not parsed correctly.")
	}
}
