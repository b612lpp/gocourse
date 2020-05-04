package main

import "fmt"

//Charakter contains list of stats
type Charakter struct {
	BaseDef    int
	Armor      int
	BaseDmg    int
	WeapoonDMG int
	LongRange  bool
}

//NonPlayerCharakter contains list of stats
type NonPlayerCharakter struct {
	BaseDef      int
	Armor        int
	BaseDmg      int
	PenaltyLevel int
	HasQuest     bool
}

//FullDmg calculates dbase and bonus damage
func (c Charakter) FullDmg() int {
	return c.BaseDmg + c.WeapoonDMG
}

//FullDmg calculates dbase and bonus damage
func (n NonPlayerCharakter) FullDmg() int {
	return n.BaseDmg - n.PenaltyLevel
}

//DefStats calculates base and bonus def
func (c Charakter) DefStats() int {
	return c.BaseDef + c.Armor
}

//DefStats calculates base and bonus def
func (n NonPlayerCharakter) DefStats() int {
	return n.BaseDef + n.Armor
}

//StatsCalculater composit interface
type StatsCalculater interface {
	DefStatsCalculater
	DmgCalculater
}

type DefStatsCalculater interface {
	DefStats() int
}

type DmgCalculater interface {
	FullDmg() int
}

func calcDmg(d StatsCalculater) int {

	metrick := d.FullDmg()
	fmt.Println(d)
	return metrick
}

func getDefStats(d StatsCalculater) int {
	metrick := d.DefStats()
	return metrick
}

func main() {

	var Marksman = Charakter{10, 20, 30, 10, true}
	var RaiDBoss = NonPlayerCharakter{40, 20, 20, 10, true}
	fmt.Println(calcDmg(Marksman))
	fmt.Println(calcDmg(RaiDBoss))

}
