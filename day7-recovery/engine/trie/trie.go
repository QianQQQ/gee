package trie

import "strings"

type Node struct {
	Pattern  string  // 待匹配路由
	Part     string  // 路由中的一部分
	Children []*Node // 一节一节的匹配的
	IsWild   bool    // 含有:, *时为模糊匹配
}

// 找到一个能匹配上的节点, 用于插入
func (n *Node) MatchChild(part string) *Node {
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			return child
		}
	}
	return nil
}

// 找到所有能匹配上的节点, 用于查找
func (n *Node) MatchChildren(part string) []*Node {
	var nodes []*Node
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *Node) Insert(pattern string, parts []string, height int) {
	// 递归终止条件
	if len(parts) == height {
		n.Pattern = pattern
		return
	}
	part := parts[height]
	child := n.MatchChild(part)
	if child == nil {
		child = &Node{Part: part, IsWild: part[0] == '*' || part[0] == ':'}
		// 竟然忘了加上去了, 体现为handlers加上去, 但是trie只有root
		n.Children = append(n.Children, child)
	}
	child.Insert(pattern, parts, height+1)
}

func (n *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.MatchChildren(part)
	for _, child := range children {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
