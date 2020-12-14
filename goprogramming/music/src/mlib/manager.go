package mlib

import (
	"errors"
	"fmt"
)

type MusicEntry struct {
	Id     string //音乐ID
	Name   string //音乐名称
	Artist string //音乐家
	Source string //地址
	Type   string //类型
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

//Len 返回音乐单长度
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//Get 返回指定index音乐内容
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= m.Len() {
		return nil, errors.New("index 超出最大范围")
	}

	return &m.musics[index], nil
}

//Find 查找指定音乐名称内容
func (m *MusicManager) Find(name string) *MusicEntry {
	if m.Len() == 0 {
		fmt.Println("没有内容")
		return nil
	}

	for _, node := range m.musics {
		if node.Name == name {
			return &node
		}
	}
	return nil
}

//Add 增加音乐内容
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//Remove 删除音乐内容
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= m.Len() {
		fmt.Println("index 超出最大范围")
		return nil
	}

	removeMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)

	return removeMusic
}
