package content

import (
	"fmt"
	"encoding/json"
	"io/ioutil"

	"github.com/itspaulyg/minecraft-mobs/model"
)

type Mobs struct {
	Mobs []Mob `json:"mobs"`
}

type Mob struct {
	Name		string		`json:"name"`
	Type		string		`json:"type"`
	HitPoints 	int			`json:"hitPoints"`
	Spawn		[]string	`json:"spawn"`
	Drops		[]string	`json:"drops"`
	Behavior []string `json:"behavior"`
}

var PassiveMobs = []string{}

var HostileMobs = []string{}

var mobs Mobs

func init() {
	data, err := ioutil.ReadFile("./data/mobs.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &mobs)
	if err != nil {
		fmt.Println(err)
	}

	for _, mob := range mobs.Mobs {
		if mob.Type == "passive" {
			PassiveMobs = append(PassiveMobs, mob.Name)
		} else if mob.Type == "hostile" {
			HostileMobs = append(HostileMobs, mob.Name)
		}
	}
}

func GetMobContent(mobName string) model.Content {
	var c model.Content
	for _, m := range mobs.Mobs {
		if m.Name == mobName {
			hitPoints := m.HitPoints
			spawn := make([]string, len(m.Spawn))
			copy(spawn, m.Spawn)
			drops := make([]string, len(m.Drops))
			copy(drops, m.Drops)
			behavior := make([]string, len(m.Behavior))
			copy(behavior, m.Behavior)

			c = model.Content{
				HitPoints:	hitPoints,
				Spawn:		spawn,
				Drops:		drops,
				Behavior: behavior,
			}
		}
	}

	return c
}
