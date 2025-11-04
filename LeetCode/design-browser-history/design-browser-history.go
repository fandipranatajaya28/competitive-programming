package main

import "fmt"

type Node struct {
	url        string
	prev, next *Node
}

type BrowserHistory struct {
	currPage *Node
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		currPage: &Node{
			url: homepage,
		},
	}
}

func (this *BrowserHistory) Visit(url string) {
	// Drop forward history
	this.currPage.next = nil
	// Append and move current page
	currPage := &Node{
		url:  url,
		prev: this.currPage,
	}
	this.currPage.next = currPage
	this.currPage = currPage
}

func (this *BrowserHistory) Back(steps int) string {
	if this == nil {
		return ""
	}
	for range steps {
		if this.currPage.prev == nil {
			break
		}
		this.currPage = this.currPage.prev
	}
	return this.currPage.url
}

func (this *BrowserHistory) Forward(steps int) string {
	if this == nil {
		return ""
	}
	for range steps {
		if this.currPage.next == nil {
			break
		}
		this.currPage = this.currPage.next
	}
	return this.currPage.url
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */

func main() {
	obj := Constructor("leetcode.com")
	obj.Visit("google.com")
	obj.Visit("facebook.com")
	obj.Visit("youtube.com")
	fmt.Println(obj.Back(1))
	fmt.Println(obj.Back(1))
	fmt.Println(obj.Forward(1))
	obj.Visit("linkedin.com")
	fmt.Println(obj.Forward(2))
	fmt.Println(obj.Back(2))
	fmt.Println(obj.Back(7))
}
