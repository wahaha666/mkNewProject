// mkNewProject project main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	name := flag.String("name", "", "你要创建的项目名字")
	pack := flag.String("pack", "", "要测试的包的名字")
	flag.CommandLine.Parse(os.Args[1:])
	if _, err := os.Stat("./" + *name); err == nil {
		color.Red("already exists")
		return
	}
	if err := os.MkdirAll(*name+"/src/vendor/src/github.com/", 777); err != nil {
		color.Red("文件夹创建失败")
		return
	}
	if err := os.MkdirAll(*name+"/src/test/", 777); err != nil {
		color.Red("文件夹创建失败")
		return
	}
	if file, err := os.Create(*name + "/src/test/main.go"); err == nil {
		file.WriteString(`package main

import (
	"fmt"
	"github.com/` + *pack + `"
)

func main() {
	
}
		`)
	}

	//	a := strings.Split(*pack, "/")
	gitTarget := *name + "/src/vendor/src/github.com/" + *pack
	if err := os.MkdirAll(gitTarget, 777); err != nil {
		color.Red("文件夹创建失败")
		return
	}
	if ok := GitClone(*pack, gitTarget); !ok {
		color.Red(" 网络爆炸")
		return
	}
	color.HiYellow("项目创建完成")

}

//输入包名和路径，克隆包
func GitClone(pack string, path string) bool {
	c := "git clone https://github.com/" + pack + " " + path
	color.Green("正在克隆...")
	cmd := exec.Command("sh", "-c", c)
	if err := cmd.Run(); err != nil {
		color.HiMagenta("克隆出错")
		return false
	}
	return true
}

//func PathExists(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}
