package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	finalMsg := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return finalMsg
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	finalMsg := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.01f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
	return finalMsg
}
