package globalconst

import "github.com/kukkar/common-golang/pkg/utils/queryparser"

var TigerHallQueryMap = map[string]queryparser.MapValue{

	"id": {"_id", "string", func(i interface{}) (bool, error) {
		//Implementation Logic provided by user
		return true, nil
	},
		[]string{"eq", "in"},
	},
	"name": {"name", "string", func(i interface{}) (bool, error) {
		//Implementation Logic provided by user
		return true, nil
	},
		[]string{"eq", "in", "lk"},
	},
	"lastcoordinates": {"lastSeenCoordinates", "string", func(i interface{}) (bool, error) {
		//Implementation Logic provided by user
		return true, nil
	},
		[]string{"eq"},
	},
}
