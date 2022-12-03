package define

import (
	"fmt"
	"strconv"
)

type Immidiate int

func StrtoImm(s string) (Immidiate, error) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, err
	}
	if i > 0x0f {
		return 0, fmt.Errorf("Error: Immidiate out of range. %v", s)
	}

	return Immidiate(i), nil
}
