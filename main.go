package main

import (
	"fmt"
)

func main() {

	type Author struct {
		Name string             //名字
		VIP bool                //是否是高贵的带会员
		Icon string             //头像
		Signature []string        //签名
		Focus int               //关注人数
	}
     {
      var hd Author
	  hd.Name = "红豆稀饭中"
	  hd.VIP = true
	  hd.Icon = "无"
	  hd.Signature =[]string{"学习新东西ing"}
	  hd.Focus = 1340000
     fmt.Println(hd)
	}
}