package fuck

import (
	logging "github.com/ipfs/go-log/v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

const FUCK_LOTUS_MINER = "FUCK_LOTUS_MINER"

var log = logging.Logger("main")
var fuckLock sync.Mutex

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CheckFuck() {
	ffp := os.Getenv(FUCK_LOTUS_MINER)
	log.Infof("%s=%s", FUCK_LOTUS_MINER, ffp)
}

func LoadFuck() int {
	fuckLock.Lock()
	defer fuckLock.Unlock()

	ffp := os.Getenv(FUCK_LOTUS_MINER)
	if ffp == "" {
		log.Errorf("cannot found env for LoadFuck:%s", FUCK_LOTUS_MINER)
		return 0
	}
	bs, err := ioutil.ReadFile(ffp)
	if err != nil {
		log.Errorf("LoadFuck error: %s", err)
		return 0
	}
	defer os.Remove(ffp)
	s := strings.ReplaceAll(string(bs),"\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, " ", "")
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Errorf("LoadFuck error: %s", err)
		return 0
	}
	return value
}