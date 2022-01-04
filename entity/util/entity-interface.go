package util

import "solar-faza/entity"

var Entities = func() []interface{} {
	ent := func(ent ...interface{}) []interface{} { return ent }
	return ent(
		&entity.User{},
		&entity.Participant{},
		&entity.Post{},
		&entity.PostComment{},
	)
}
