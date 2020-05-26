package fuck

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFuck4LocalEnv(t *testing.T) {
	os.Unsetenv(FUCK_LOTUS_MINER)
	noEnvValue := LoadFuck()
	if noEnvValue != 0 {
		t.Fatal("got invalid value:", noEnvValue)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("got user home error:", err)
	}
	ffp := home + "/fuck_lotus_miner"
	if checkFileIsExist(ffp) {
		os.Remove(ffp)
	}
	os.Setenv(FUCK_LOTUS_MINER, ffp)
	noFileValue := LoadFuck()
	if noFileValue != 0 {
		t.Fatal("got invalid value:", noFileValue)
	}

	f, err := os.Create(ffp)
	if err != nil {
		t.Fatal("create file error:", err)
	}
	t.Log("created file " + f.Name())
	if err := ioutil.WriteFile(ffp, []byte("1"), 0644); err != nil {
		t.Fatal("write file error:", err)
	}
	fuckFileValue := LoadFuck()
	if fuckFileValue != 1 {
		t.Fatal("got invalid value:", fuckFileValue)
	}
}

func TestFuck4SysEnv(t *testing.T) {
	CheckFuck()
	if value := LoadFuck(); value != 999 {
		t.Fatal("got invalid value:", value)
	}
}