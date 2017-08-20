package model

import (
)

type Skill struct {
	Level int
	Exp int
}
func (s *Skill) GetExp() (int)  {
	return s.Exp
}
func (s *Skill) LevelUp() {
	s.Level ++
}

type ExpCon interface {
	GetName() string
	LevelUp()
	GetExp() int
	GetExpLimit() int
}
func ExpUp(con ExpCon, exp int, cb func())  {
	exp += con.GetExp()
	flag := false
	for exp > con.GetExpLimit() {
		exp -= con.GetExpLimit()
		con.LevelUp()
		flag = true
	}
	if flag {
		cb()
	}
}

type BodySkill struct {
	Skill
}
func (s *BodySkill) GetName() (string)  {
	return "Body"
}
func (s *BodySkill) GetExpLimit() (int)  {
	return s.Level * 200
}
type ActSkill struct {
	Skill
}
func (s *ActSkill) GetName() (string)  {
	return "Act"
}
func (s *ActSkill) GetExpLimit() (int)  {
	return s.Level * 50
}
