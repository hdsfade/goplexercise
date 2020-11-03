//@author: hdsfade
//@date: 2020-11-02-18:34
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("can't get %s: %d", url, resp.Status)
	}
	return resp, nil
}

func GetIssue(owner, repo, number string) (*Issue, error) {
	url := strings.Join([]string{APIURL,"repos",owner,repo,"issues",number},"/")
	resp, err := get(url)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue);err != nil{
		return nil,err
	}
	return &issue, nil
}

func GetIssues(owner, repo string) ([]Issue, error) {
	url := strings.Join([]string{APIURL,"repos",owner,"issues",repo},"/")
	resp, err := get(url)
	if err != nil{
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("can't get %s: %s", url, resp.Status)
	}
	defer resp.Body.Close()
	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(issues);err !=nil {
		return nil, err
	}
	return issues, nil
}

func EditIssue(owner, repo , number string, fields map[string]string) (*Issue, error) {
	buf := &bytes.Buffer{}     //*bytes.Buffer{}才是io.ReadWriter
	err := json.NewEncoder(buf).Encode(fields)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}  //创建客户端
	url := strings.Join([]string{APIURL,"repos",owner,repo,"issues",number},"/")
	req,err := http.NewRequest("PATCH",url,buf)   //创建请求
	req.SetBasicAuth(os.Getenv("GITHUB_USER"),os.Getenv("GITHUB_PASS"))  //从环境变量中获得github账号密码
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)  //发送请求
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't edit issue: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil{
		return nil, err
	}
	//返回修改后的issue
	return &issue, nil
}
