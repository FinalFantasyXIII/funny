package staff

type Player interface {
	//表演优先级
	Level() int
	//获取表演者信息
	Info() string
	//自我介绍
	Introduce()
	//表演节目
	Play()
	//离场谢辞
	Thanks()
}


