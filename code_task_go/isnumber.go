package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Alter_Split(str string) (str2 string) {

	rege1 := regexp.MustCompile(`(?s:(alter table \w+ ))`)
	find1 := rege1.FindAllString(str, 1)
	fmt.Println("find1 = ", find1)
	//fmt.Printf("find1 is a %T", find1)
	//以“,”分割
	s := strings.SplitAfter(str, ",")
	//srt_n := len(s)
	//fmt.Println(s)
	result := strings.Join(s[0:], "  \n"+find1[0]) //进行组合
	//fmt.Println("result = ", result)
	str2 = result
	return

}
func main() {
	//测试str  有多个tab和多个空格之类的
	str := `
	alter	 table	 sdsdsdsdo	 add column xxxx, 
		
					
						
			modify   column   xxxxx , 
			
			add index xxxx dddd, 
			 
		drop column haha;
		`
	str = strings.Replace(str, "	", "", -1)  //去除tab
	str = strings.Replace(str, "\n", "", -1) //去除换行
	str = strings.Replace(str, "  ", "", -1) //去除俩个一起的空格
	fmt.Println("格式化后结果：\n", str)
	Alter_Split(str)
	fmt.Println("Split result : ")
	fmt.Print(Alter_Split(str))

}
