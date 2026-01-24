package data

import "authgo/models"

var Users = make(map[int]*models.User)
var NextUserID = 1
