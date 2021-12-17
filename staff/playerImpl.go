package staff

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
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
 */

func PlayerComparator(a , b interface{}) int {
	aAsserted := a.(Player).Level()
	bAsserted := b.(Player).Level()
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}


// Jack
func (jack *Jack) Level() int{
	return jack.LV
}

func (jack *Jack) Info() string{
	sa := make([]string,0)
	sa = append(sa,jack.Name)
	sa = append(sa,strconv.Itoa(jack.LV))
	return strings.Join(sa,"-")
}

func (jack *Jack) Introduce() {
	info := jack.Info()

	fmt.Printf("hello,everyone , this is %s \n",info)
	fmt.Printf("show name is %s \n",jack.ShowName)
}

func (jack *Jack) Play(){
	sum := 0
	for _,v := range jack.IntegerArray{
		sum += v
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("%v total is %d \n",jack.IntegerArray, sum)
}

func (jack *Jack) Thanks(){
	fmt.Println(jack.Peroration)
}

// Rose
func (rose *Rose) Level() int{
	return rose.LV
}

func (rose *Rose) Info() string{
	sa := make([]string,0)
	sa = append(sa,rose.Name)
	sa = append(sa,strconv.Itoa(rose.LV))
	return strings.Join(sa,"-")
}

func (rose *Rose) Introduce() {
	info := rose.Info()

	fmt.Printf("hello,everyone , this is %s \n",info)
	fmt.Printf("show name is %s \n",rose.ShowName)
}

func (rose *Rose) Play(){
	sum := ""
	for _,v := range rose.StringArray{
		sum += v
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("%v total is %s \n",rose.StringArray, sum)
}

func (rose *Rose) Thanks(){
	fmt.Println(rose.Peroration)
}

//PhonixLegend
func (pl *PhoenixLegend) Level() int{
	return pl.LV
}

func (pl *PhoenixLegend) Info() string{
	sa := make([]string,0)
	sa = append(sa,pl.Name)
	sa = append(sa,strconv.Itoa(pl.LV))
	return strings.Join(sa,"-")
}

func (pl *PhoenixLegend) Introduce() {
	info := pl.Info()

	fmt.Printf("hello,everyone , this is %s \n",info)
	fmt.Printf("show name is %s \n",pl.ShowName)
}

func (pl *PhoenixLegend) Play(){
	fmt.Println(pl.Song)
}

func (pl *PhoenixLegend) Thanks(){
	fmt.Println(pl.Peroration)
}


//FallOutBoy
func (fob *FallOutBoy) Level() int{
	return fob.LV
}

func (fob *FallOutBoy) Info() string{
	sa := make([]string,0)
	sa = append(sa,fob.Name)
	sa = append(sa,strconv.Itoa(fob.LV))
	return strings.Join(sa,"-")
}

func (fob *FallOutBoy) Introduce() {
	info := fob.Info()

	fmt.Printf("hello,everyone , this is %s \n",info)
	fmt.Printf("show name is %s \n",fob.ShowName)
}

func (fob *FallOutBoy) Play(){
	fmt.Println(fob.Song)
}

func (fob *FallOutBoy) Thanks(){
	fmt.Println(fob.Peroration)
}

