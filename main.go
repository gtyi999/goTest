package main

import (
	"fmt"
	"flag"
	"Test_SW_Articleadd/config"
	"Test_SW_Articleadd/mylog"
	"Test_SW_Articleadd/handle"
)

func main(){
	fmt.Println("Hello main")
	flag.Parse()
	config.ConfigInit()
	mylog.SetupLogger("Test_SW_Articleadd")

	//db.SetupDBMysql()

	handle.PicurlTransfer()

}

