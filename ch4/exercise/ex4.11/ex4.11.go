//@author: hdsfade
//@date: 2020-11-02-17:05
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//打开issue
func open(owner,repo,number string) {
	_, err := EditIssue(owner,repo,number,map[string]string{"state":"open"})
	if err != nil{
		log.Fatal(err)
	}
}
//关闭issue
func close(owner, repo,number string) {
	_, err := EditIssue(owner, repo,number,map[string]string{"state":"closed"})
	if err != nil{
		log.Fatal(err)
	}
}
//读取issue
func read(owner, repo, number string) {
	issue, err := GetIssue(owner, repo, number)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("repo: %s/%s\nissuenumber: %s\ncreatedat: %s\nstate: %s\nuser: %s\n",
		owner, repo, number, issue.CreatedAt,issue.State,issue.User.Login)
	fmt.Printf("title: %s\n", issue.Title)
	fmt.Printf("\n%s\n",issue.Body)
}
//打开编辑器，编辑issue
func edit(owner, repo, number string) {
	editor := os.Getenv("EDITOR") //从环境变量中获取默认编辑器
	//如果不存在编辑器环境变量，则使用notepad作为编辑器
	if editor == "" {
		editor = "notepad"
	}

	editorPath,err :=exec.LookPath(editor)  //寻找可执行文件路径
	if err != nil{
		log.Fatal(err)
	}

	tmpfile, err := ioutil.TempFile("", "issue")  //指定路径和文件名后缀创建临时文件
	if err != nil {
		log.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	issue, err := GetIssue(owner,repo,number)       //读取issue
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(tmpfile)
	//将issue内容以json形式写入tmpfile
	err = encoder.Encode(map[string]string{
		"title": issue.Title,
		"state":issue.State,
		"body":issue.Body,
	})
	if err != nil {
		log.Fatal(err)
	}

	//创建命令
	cmd := &exec.Cmd{
		Path:editorPath,
		Args: []string{editor,tmpfile.Name()},
		Stdin: os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	//运行命令
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	tmpfile.Seek(0,0) //读写指针返回文件开头
	fields := make(map[string]string)
	if err := json.NewDecoder(tmpfile).Decode(fields); err != nil{
		log.Fatal(err)
	}

	_,err = EditIssue(owner,repo,number,fields)   //修改内容写入github上的issue
	if err != nil{
		log.Fatal(err)
	}
}

var usage string = `usage:
[read|edit|close|open] OWNER REPO ISSUE_NUMBER
`


func main() {
	if len(os.Args) < 2 ||len(os.Args)>5 {
		fmt.Fprintln(os.Stderr,usage)
		os.Exit(1)
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	owner, repo, number := args[0], args[1], args[2]
	switch cmd {
	case "read":
		read(owner, repo, number)
	case "edit":
		edit(owner, repo, number)
	case "close":
		close(owner, repo, number)
	case "open":
		open(owner, repo, number)
	}
}

