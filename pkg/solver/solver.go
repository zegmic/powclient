package solver

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

func Solve(challenge string) int {
	for i := 0; i < math.MaxInt; i++ {
		sol := fmt.Sprintf("%d.%s", i, challenge)
		dig := sha256.New()
		dig.Write([]byte(sol))
		sum := dig.Sum(nil)

		if solved(hex.EncodeToString(sum[:])) {
			return i
		}
	}

	return -1
}

func solved(sum string) bool {
	return strings.HasPrefix(sum, "00000")
}
