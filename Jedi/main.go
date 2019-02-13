package main

import (
	"github.com/astaxie/beego"
	"project_wm/Jedi/models"
	_ "project_wm/Jedi/routers"
)

func main() {
  	models.Finit1()
	beego.Run()
}


/**/