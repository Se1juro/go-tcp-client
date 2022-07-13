package models

type User struct {
	id             string
	nickname       string
	currentChannel int
}

func (u *User) SetNickName(nickName string) {
	u.nickname = nickName
}

func (u *User) SetId(id string) {
	u.id = id
}

func (u *User) SetCurrentChannel(currentChannel int) {
	u.currentChannel = currentChannel
}
