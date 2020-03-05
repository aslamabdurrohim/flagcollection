package flagcollection

import (
	"flag"
)

type FlagCollection struct {
	ConfigTest bool
}

func SetFlagCollection() FlagCollection {
	var flags FlagCollection
	flag.BoolVar(&flags.ConfigTest, "test", false, "read test flag")
	flag.Parse()

	return flags
}
