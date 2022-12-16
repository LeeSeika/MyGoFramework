package framework

import (
	"errors"
	"github.com/godemo/coredemo/framework/gin"
	"strings"
)

type node struct {
	segment  string
	isLast   bool
	handlers []gin.HandlerFunc
	children []*node
	parent   *node
}

type Tree struct {
	root *node
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func newNode() *node {
	return &node{}
}

func newTree() *Tree {
	return &Tree{root: newNode()}
}

func (n *node) filterChildNode(segment string) []*node {
	if n.children == nil || len(n.children) == 0 {
		return nil
	}
	if isWildSegment(segment) {
		return n.children
	}
	nodes := make([]*node, 0, len(n.children))
	for _, child := range n.children {
		if isWildSegment(child.segment) {
			nodes = append(nodes, child)
		} else if segment == child.segment {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) matchNode(url string) *node {
	if len(url) == 0 {
		return nil
	}

	splitUrl := strings.SplitN(url, "/", 2)
	segment := splitUrl[0]

	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	filterNodes := n.filterChildNode(segment)
	if filterNodes == nil || len(filterNodes) == 0 {
		return nil
	}

	if len(splitUrl) == 1 {
		for _, node := range filterNodes {
			if node.isLast {
				return node
			}
		}
		return nil
	} else {
		for _, node := range filterNodes {
			resNode := node.matchNode(splitUrl[1])
			if resNode != nil {
				return resNode
			}
		}
	}

	return nil
}

func (tree *Tree) AddRouter(url string, handlers ...gin.HandlerFunc) error {
	if tree.root.matchNode(url) != nil {
		return errors.New("handler exist")
	}
	n := tree.root
	split := strings.Split(url, "/")
	for index, segment := range split {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		var objNode *node
		childNodes := n.filterChildNode(segment)

		isLast := index == len(split)-1

		if len(childNodes) != 0 {
			for _, targetNode := range childNodes {
				if targetNode.segment == segment {
					objNode = targetNode
					break
				}
			}
		}

		if childNodes == nil || len(childNodes) == 0 {
			newNode := newNode()
			newNode.segment = segment
			newNode.isLast = isLast
			newNode.parent = n
			if isLast {
				newNode.handlers = handlers
			}
			n.children = append(n.children, newNode)
			objNode = newNode
		}
		n = objNode
	}

	return nil
}

func (tree *Tree) FindHandler(url string) (*node, error) {
	matchNode := tree.root.matchNode(url)
	if matchNode == nil {
		return nil, errors.New("handler not found")
	}
	return matchNode, nil
}

func (n *node) parseParamFromEndNode(url string) map[string]string {
	if !n.isLast {
		return map[string]string{}
	}
	endNode := n.matchNode(url)
	segements := strings.Split(url, "/")
	len := len(segements)
	curr := endNode
	paramsMap := map[string]string{}
	for index := len - 1; index >= 0; index-- {
		segment := segements[index]
		if curr.segment == "" {
			break
		}
		if isWildSegment(curr.segment) {
			// 这里切片是为了去掉开头的"/"
			paramsMap[curr.segment[1:]] = segment
		}
		curr = curr.parent
	}
	return paramsMap
}
