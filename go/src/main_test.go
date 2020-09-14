package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error("Não foi possivel ler a saida")
	}
	os.Stdout = rescueStdout

	if string(out) == "Code.education Rocks!" {
		t.Errorf("Expected %s, got %s", "Code.education Rocks!: test", out)
	}
}
