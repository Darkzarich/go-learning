/*
Assignment:

Imagine you're developing a multiplayer game. The game logic is processed on a server implemented in Go.

At some point, the designers request that players be able to cast mass spells.
This means the spell should affect MULTIPLE objects with different structures and types.

You realize that rewriting all the types and implementing the CastReceiver interface for each is too complex a task.

Implement spell casting using reflection.

Each spell satisfies the Spell interface - you can find out which object characteristic and what value it affects.
*/

package main

import (
	"fmt"
	"log"
	"reflect"
)

type Spell interface {
	// Spell name
	Name() string
	// Stat that is affected by spell
	Char() string
	// Value of spell
	Value() int
}

// CastReceiver — if a struct implements this interface, it should receive spells using ReceiveSpell
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {
	// My solution:

	if obj, ok := object.(CastReceiver); ok {
		obj.ReceiveSpell(spell)

		return
	}

	val := reflect.ValueOf(object)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	field := val.FieldByName(spell.Char())

	if !field.IsValid() || !field.CanSet() {
		return
	}

	// The field value must be an integer (any size)
	if field.Kind() != reflect.Int && field.Kind() != reflect.Int8 &&
		field.Kind() != reflect.Int16 && field.Kind() != reflect.Int32 &&
		field.Kind() != reflect.Int64 {
		return
	}

	field.SetInt(field.Int() + int64(spell.Value()))

	log.Printf("Casted spell %s to %#v", spell.Name(), object)
}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}

func main() {

	player := &Player{
		name:   "Player_1",
		health: 100,
	}

	enemies := []interface{}{
		&Zombie{Health: 1000},
		&Zombie{Health: 1000},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Daemon{Health: 1000},
		&Daemon{Health: 1000},
		&Wall{Durability: 100},
	}

	CastToAll(newSpell("fire", "Health", -50), append(enemies, player))
	CastToAll(newSpell("heal", "Health", 190), append(enemies, player))

	fmt.Println(player)
	for _, e := range enemies {
		fmt.Println(e)
	}
}
