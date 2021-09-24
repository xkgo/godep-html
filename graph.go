package main

import (
	"strings"
)

// Node 节点定义
type Node struct {
	Id       string  `json:"title"`
	Pid      string  `json:"pid"`
	Children []*Node `json:"children"`
	Expand   bool    `json:"expand"`
}

/*
Parse
解析 go mod graph 输出的结果, 返回 顶层 nodes
*/
func Parse(content string) (topNodes []*Node) {
	allNodes := make([]*Node, 0)
	nodeMap := make(map[string]*Node)
	childrens := make(map[string][]*Node)

	array := strings.Split(content, "\n")

	for _, line := range array {
		if len(line) == 0 {
			continue
		}
		subArr := strings.Split(line, " ")
		if len(subArr) != 2 {
			continue
		}

		id, pid := subArr[1], subArr[0]

		if _, ok := nodeMap[pid]; !ok {
			node := &Node{Id: pid}
			nodeMap[pid] = node
			allNodes = append(allNodes, node)
			topNodes = append(topNodes, node)
		}

		if _, ok := nodeMap[id]; !ok {
			node := &Node{Id: id, Pid: pid}
			nodeMap[id] = node
			allNodes = append(allNodes, node)
			pChildren := childrens[pid]
			if len(pChildren) == 0 {
				childrens[pid] = []*Node{node}
			} else {
				exist := false
				for _, item := range pChildren {
					if item.Id == id {
						exist = true
						break
					}
				}
				if !exist {
					childrens[subArr[0]] = append(pChildren, node)
				}
			}
		}
	}

	// 处理父子关系
	for _, node := range allNodes {
		node.Children = childrens[node.Id]
		node.Expand = true
	}
	return
}
