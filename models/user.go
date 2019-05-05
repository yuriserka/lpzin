package models

var db = make(map[int]string)

type User struct {
	ID int
	profilePic string
}

func init() {
	db[1] = "../static/img/yuri.jpg"
}

func RequestUserProfilePicture(usrId int) string {
	return db[usrId]
}