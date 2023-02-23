package battle

import (
	"fmt"
	"math/rand"
)

// Character represents the properties of character
type Character struct {
	Name   string  // name of the character
	Stats          // stats of the character
	Skills []Skill // skills associated with the character
}

// IsLucky checks if the character is lucky
func (c *Character) IsLucky() bool {
	luck := rand.Intn(101)
	if luck <= c.Luck {
		return true
	} else {
		return false
	}
}

// String used to log the health of the character
func (c *Character) String() string {
	if c.Health <= 0 {
		return fmt.Sprint(c.Name, " Health: ", 0)
	}
	return fmt.Sprint(c.Name, " Health: ", c.Health)
}
