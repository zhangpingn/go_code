package dao

import (
	"go_code/restudy/netstudy/netstudy03/server/db"
	"go_code/restudy/netstudy/netstudy03/common/model"
)

type HeroDao struct {

}

func (this *HeroDao) GetHeroes(heroName string) (hero *model.Heroes){
	hero = &model.Heroes{}
	db.DB.QueryRow("select blood, hero_name, power from heroes where hero_name = ?", heroName).Scan(&hero.Blood, &hero.HeroName, &hero.Power)
	return
}