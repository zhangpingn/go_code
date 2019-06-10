package model

import (
	"fmt"
)

type Heroes struct {
	Blood uint
	BloodChan chan uint
	HeroName string
	Power uint
}

func (this *Heroes) GetMsg() {
	fmt.Printf("您的血量是%d, 名字是%v, 攻击力是%d", 
		this.Blood, this.HeroName, this.Power)
}

//这是服务器调用的方法，根据HeroName找到对应的Hero，然后调用其的该方法改变Blood
func (this *Heroes) ChangeBloodChan(Hero *Heroes) {
	this.BloodChan <- Hero.Power
}

func (this *Heroes) ChangeBlood() {
	for {
		select {
			case v := <- this.BloodChan : 
				this.Blood = this.Blood + v
			default :
		}
	}
}