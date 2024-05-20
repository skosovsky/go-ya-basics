package main

import (
	"log"
	"reflect"
)

type Spell interface {
	// Name название заклинания
	Name() string
	// Char характеристика, на которую воздействует
	Char() string
	// Value количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него.
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []any) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object any) {
	if receiver, ok := object.(CastReceiver); ok {
		receiver.ReceiveSpell(spell)

		return
	}

	val := reflect.ValueOf(object)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return
	}

	field := val.Elem().FieldByName(spell.Char())
	if !field.IsValid() {
		return
	}

	if !field.CanSet() {
		return
	}

	if field.Kind() != reflect.Int &&
		field.Kind() != reflect.Int8 &&
		field.Kind() != reflect.Int16 &&
		field.Kind() != reflect.Int32 &&
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

func newSpell(name string, char string, val int) Spell { //nolint:ireturn // example
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
		health: 100, //nolint:mnd // example
	}

	enemies := []any{
		&Zombie{Health: 1000},  //nolint:mnd // example
		&Zombie{Health: 1000},  //nolint:mnd // example
		&Orc{Health: 500},      //nolint:mnd // example
		&Orc{Health: 500},      //nolint:mnd // example
		&Orc{Health: 500},      //nolint:mnd // example
		&Daemon{Health: 1000},  //nolint:mnd // example
		&Daemon{Health: 1000},  //nolint:mnd // example
		&Wall{Durability: 100}, //nolint:mnd // example
	}

	CastToAll(newSpell("fire", "Health", -50), append(enemies, player))
	CastToAll(newSpell("heal", "Health", 190), append(enemies, player)) //nolint:mnd // example

	log.Println(player)
}
