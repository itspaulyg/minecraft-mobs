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
	Height		float32		`json:"height"`
	Width		float32		`json:"width"`
	Spawn		[]string	`json:"spawn"`
	Drops		[]string	`json:"drops"`
	Behavior 	[]string 	`json:"behavior"`
}

var AllMobs = []string{}

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

	for _, m := range mobs.Mobs {
		AllMobs = append(AllMobs, m.Name)
	}
}

func GetMobsByFilter(mobType string) []string {
	if mobType == "all" {
		return AllMobs
	}

	var filteredMobs = []string{}
	for _, m := range mobs.Mobs {
		if m.Type == mobType {
			filteredMobs = append(filteredMobs, m.Name)
		}
	}

	return filteredMobs
}

func GetMobContent(mobName string) model.Content {
	var c model.Content
	for _, m := range mobs.Mobs {
		if m.Name == mobName {
			hitPoints := m.HitPoints
			height := m.Height
			width := m.Width
			spawn := make([]string, len(m.Spawn))
			copy(spawn, m.Spawn)
			drops := make([]string, len(m.Drops))
			copy(drops, m.Drops)
			behavior := make([]string, len(m.Behavior))
			copy(behavior, m.Behavior)

			c = model.Content{
				HitPoints:	hitPoints,
				Height:		height,
				Width:		width,
				Spawn:		spawn,
				Drops:		drops,
				Behavior: 	behavior,
			}
		}
	}

	return c
}
