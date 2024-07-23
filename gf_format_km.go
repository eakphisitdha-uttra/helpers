package helpers

import (
	"fmt"
	"strconv"
)

func FormatKM(n int64) string {
	if n < 0 {
		return "-" + FormatKM(-n)
	} else if n < 1000 {
		in := []byte(strconv.FormatInt(n, 10))
		if len(in) == 1 {
			return fmt.Sprintf("0+00%d", n)
		} else if len(in) == 2 {
			return fmt.Sprintf("0+0%d", n)
		} else {
			return fmt.Sprintf("0+%d", n)
		}
	} else {
		in := []byte(strconv.FormatInt(n, 10))
		var out []byte
		if i := len(in) % 3; i != 0 {
			if out, in = append(out, in[:i]...), in[i:]; len(in) > 0 {
				out = append(out, '+')
			}
		}
		for len(in) > 0 {
			if out, in = append(out, in[:3]...), in[3:]; len(in) > 0 {
				out = append(out, '+')
			}
		}
		return string(out)
	}
}
