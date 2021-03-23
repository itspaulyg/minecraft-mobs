module github.com/itspaulyg/minecraft-mobs

go 1.16

replace github.com/itspaulyg/minecraft-mobs/content => ./src/content

replace github.com/itspaulyg/minecraft-mobs/model => ./src/model

require (
	github.com/itspaulyg/minecraft-mobs/content v0.0.0-00010101000000-000000000000
	github.com/itspaulyg/minecraft-mobs/model v0.0.0-00010101000000-000000000000
)
