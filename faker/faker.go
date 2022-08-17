package faker

import (
	"gitag.ir/thepot/kit/dtp"
	"time"
)

func SQLNow() dtp.NullTime {
	return dtp.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}
