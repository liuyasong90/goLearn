package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type StateNum struct {
	State int
	Num int
}

func getNodeStateNum(stateNumList []StateNum)(int,int,int,int){

	norm,abnorm,unstart,stop:=0,0,0,0

	for _,stateNum:=range stateNumList{
		switch stateNum.State {
		case 0:
			unstart=stateNum.Num
		case 2:
			norm=stateNum.Num
		case 4:
			abnorm=stateNum.Num

		case 5:
			stop=stateNum.Num
		}
	}
	return norm,abnorm,unstart,stop
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3307)/fabric?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	var test []StateNum
	db.Table("peer").Select("count(*) as num,state").Group("state").Scan(&test)
	peernorm,peerabnorm,peerunstart,peerstop:=getNodeStateNum(test)

	peerNum,orderNum,caNum,kafakaNum,zookeeperNum:=0,0,0,0,0
	db.Table("peer").Count(&peerNum)
	fmt.Printf("peerNum:%d\n",peerNum)

	db.Table("orderer").Count(&orderNum)
	fmt.Printf("orderNum:%d\n",orderNum)

	db.Table("ca").Count(&caNum)
	fmt.Printf("caNum:%d\n",caNum)

	db.Table("zookeeper").Count(&zookeeperNum)
	fmt.Printf("zookeeperNum:%d\n",zookeeperNum)


	db.Table("kafka").Count(&kafakaNum)
	fmt.Printf("kafakaNum:%d\n",kafakaNum)



	//db.Table("orderer").Select("count(*) as num,state").Group("state").Scan(&test)
	ordernorm,orderabnorm,orderunstart,orderstop:=0,0,0,0


	println(peernorm+ordernorm,peerabnorm+orderabnorm,peerunstart+orderunstart,peerstop+orderstop)
	for _,j:=range test{
	fmt.Printf("State:%d,Num:%d\n",j.State,j.Num)
	}


	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}

	fmt.Println("连接成功！！！")
}

