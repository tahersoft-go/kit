package faker

import (
	"time"

	"github.com/mhosseintaher/kit/dtp"
)

func SQLNow() dtp.NullTime {
	return dtp.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}
