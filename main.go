package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	if len(os.Args) > 1 {
		cli()
	} else {
		ui()
	}
}

//命令行模式
func cli() {
	fmt.Println("Please input want remove file or directory:")
	input := bufio.NewScanner(os.Stdin) //初始化一个扫表对象
	for input.Scan() {                  //扫描输入内容
		file_path := input.Text() //把输入内容转换为字符串
		if _, err := os.Stat(file_path); err != nil {
			fmt.Println("input path not exists!!!")
			fmt.Println("Please input want remove file or directory:")
			continue
		}
		Remove(file_path)
		fmt.Println("Please input want remove file or directory:")
	}
}
//删除操作
func Remove(file_path string) {
	fmt.Println(file_path)
	s, err := os.Stat(file_path)
	if err != nil {
		fmt.Println("filepath not exists", err)
		return
	}
	if !s.IsDir() {
		err := os.Remove(file_path)
		if err != nil {
			fmt.Println("remove error:", err)
			return
		}
	} else {
		res := removeDir(file_path)
		if !res {
			fmt.Println("remove error")
			return
		}
	}
	fmt.Println("remove success")
}

//清除顽固文件夹
func removeDir(path string) bool {
	parent, _ := filepath.Split(path)
	temp := filepath.Join(parent, "t1")
	if _, err := os.Stat(temp); err != nil {
		err := os.Mkdir(temp, os.ModePerm)
		if err != nil {
			fmt.Println("make temp dir fail:", err)
			return false
		}
	}

	defer os.Remove(temp)

	c := exec.Command("robocopy", temp, path, "/MIR")
	if err := c.Run(); err != nil {
		fmt.Println("remove command run:", err)
	}
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//当前目录
func currentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return filepath.ToSlash(dir)
}


var mw = &walk.MainWindow{}
func ui() {
	var inTE *walk.TextEdit
	MainWindow{
		AssignTo: &mw,
		Title:   "深度删除",
		Size: Size{400, 200},
		MaxSize: Size{400, 200},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
				},
			},
			PushButton{
				Text: "选择文件",
				OnClicked: func() {
					file_path := selectFile()
					if file_path != "" {
						inTE.AppendText(file_path+"\r\n")
					}
				},
			},
			PushButton{
				Text: "选择文件夹",
				OnClicked: func() {
					ask()
					file_path := selectDir()
					if file_path != "" {
						inTE.AppendText(file_path+"\r\n")
					}
				},
			},
		},
	}.Run()
}

//选择文件操作
func selectFile() string {
	dlg := new(walk.FileDialog)
	dlg.Title = "请选择要删除的文件或目录"
	dlg.Filter = "所有文件 (*.*)|*.*"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		log.Println("Error : File Open")
		return ""
	} else if !ok {
		log.Println("Cancel")
		return ""
	}
	log.Println("Select :", dlg.FilePath)
	return dlg.FilePath
}
//选择目录操作
func selectDir() string {
	dlg := new(walk.FileDialog)
	dlg.Title = "请选择要删除的文件或目录"
	dlg.Filter = "所有文件 (*.*)|*.*"

	if ok, err := dlg.ShowBrowseFolder(mw); err != nil {
		log.Println("Error : File Open")
		return ""
	} else if !ok {
		log.Println("Cancel")
		return ""
	}
	log.Println("Select :", dlg.FilePath)
	return dlg.FilePath
}

func ask(){
	walk.MsgBox(mw, "About", "Walk Image Viewer Example", walk.MsgBoxIconInformation)
}
