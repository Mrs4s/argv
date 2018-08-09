package argv

import (
	"testing"
)

func TestParse(t *testing.T) {
	args := []string{"./mycommand", "-abc", "123", "--testing", "omg", "NOKEY1", "NOKEY2", "-haha", "--haha2"}
	myArgs := Parse(args[1:])

	if myArgs.Nokeys[0] != "NOKEY1" || myArgs.Nokeys[1] != "NOKEY2" {
		t.Error("Keyless arguments were not parsed and provided correctly.")
	}

	if myArgs.Keys["abc"] != "123" {
		t.Error("'abc' key was not parsed correctly.")
	}

	if myArgs.Keys["testing"] != "omg" {
		t.Error("'testing' key was not parsed correctly.")
	}

	if myArgs.Keys["haha"] != "" {
		t.Error("'haha' key was not parsed correctly.")
	}

	if myArgs.Keys["haha2"] != "" {
		t.Error("'haha2' key was not parsed correctly.")
	}
}
