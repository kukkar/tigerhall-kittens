package tigerhall

import "context"

func GetTigerHallKittens(c context.Context, config ConfigTigerHall) (Tigerhall, error) {
	return &tigherhall{
		stAdapter: &mongoAdapter{},
	}, nil
}
