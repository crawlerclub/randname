package randomname

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func MakeRandomName() string {
	n1 := random(0, len(xing)-1)
	n2 := random(0, len(ming)-1)
	s1 := xing[n1]
	s2 := ming[n2]
	return s1 + s2
}
