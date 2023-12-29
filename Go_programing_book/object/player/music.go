package player

import "errors"

type Music struct {
	Id       string
	Name     string
	Artist   string
	Location string
	FileType string
}

type Manager struct {
	musics []Music
}

func NewManager(list []Music) *Manager {
	return &Manager{musics: make([]Music, 0)}
}

func (this *Manager) Len() int {
	return len(this.musics)
}

func (this *Manager) Get(idx int) (res *Music, err error) {
	if idx > this.Len()-1 || idx < 0 {
		return nil, errors.New("idx out of range")
	}
	return &this.musics[idx], nil
}

func (this *Manager) Find(name string) *Music {
	for _, v := range this.musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (this *Manager) Add(music *Music) {
	this.musics = append(this.musics, *music)
}

func (this *Manager) Remove(idx int) (music *Music, err error) {
	if idx > this.Len()-1 || idx < 0 {
		return nil, errors.New("idx out of range")
	}
	music = &this.musics[idx]

	if idx == 0 {
		this.musics = this.musics[1:]
	} else if idx == this.Len()-1 {
		this.musics = this.musics[:idx-1]
	} else {
		this.musics = append(this.musics[:idx-1], this.musics[idx+1:]...)
	}
	return
}

func (this *Manager) RemoveByName(name string) (music *Music) {
	for i := 0; i < this.Len(); i++ {
		tmp, _ := this.Get(i)
		if tmp.Name == name {
			if i == 0 {
				this.musics = this.musics[1:]
			} else if i == this.Len()-1 {
				this.musics = this.musics[:i-1]
			} else {
				this.musics = append(this.musics[:i-1], this.musics[i+1:]...)
			}
			return tmp
		}
	}
	return nil
}
