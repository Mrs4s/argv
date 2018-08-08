package argv

import (
	"testing"
)

func TestParse(t *testing.T) {
	args := []string{"./mycommand", "-abc", "123", "--testing", "omg", "NOKEY1", "NOKEY2", "-haha", "--haha2"}
	myArgs := Parse(args[1:])

	if myArgs.nokeys[0] != "NOKEY1" || myArgs.nokeys[1] != "NOKEY2" {
		t.Error("Keyless arguments were not parsed and provided correctly.")
	}

	if myArgs.keys["abc"] != "123" {
		t.Error("'abc' key was not parsed correctly.")
	}

	if myArgs.keys["testing"] != "omg" {
		t.Error("'testing' key was not parsed correctly.")
	}

	if myArgs.keys["haha"] != "" {
		t.Error("'haha' key was not parsed correctly.")
	}

	if myArgs.keys["haha2"] != "" {
		t.Error("'haha2' key was not parsed correctly.")
	}
}
