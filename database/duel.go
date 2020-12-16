package database

import "github.com/youremine2493/Dalaran/utils"

type Duel struct {
	EnemyID    int
	Coordinate utils.Location
	Started    bool
}
