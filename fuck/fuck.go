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

func init() {
	CheckFuck()
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CheckFuck() {
	ffp := os.Getenv(FUCK_LOTUS_MINER)
	log.Infof("CheckFuck env: %s=%s", FUCK_LOTUS_MINER, ffp)
	bs, err := ioutil.ReadFile(ffp)
	if err != nil {
		log.Infof("CheckFuck file(%s) error: %s", ffp, err)
	} else {
		log.Infof("CheckFuck file(%s) value: %s", ffp, string(bs))
	}
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