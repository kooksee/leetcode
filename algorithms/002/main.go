package main

import (
	"fmt"
	"math"
	"time"
)

/*
https://leetcode-cn.com/problems/add-two-numbers/
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	_l1 := 0
	for _li1 := 0; l1 != nil; {
		_l1 = l1.Val*int(math.Pow10(_li1)) + _l1
		l1 = l1.Next
		_li1 += 1
	}

	_l2 := 0
	for _li2 := 0; l2 != nil; {
		_l2 = l2.Val*int(math.Pow10(_li2)) + _l2
		l2 = l2.Next
		_li2 += 1
	}

	_l3 := _l1 + _l2
	var l3 *ListNode
	_cur := l3
	for _l3 > 0 {
		_cur.Next = &ListNode{Val: _l3 % 10}
		_cur = _cur.Next
		_l3 /= 10
	}

	return l3
}

func main() {
	t := time.Now().UnixNano()
	fmt.Println("start: ", t)
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	fmt.Println(addTwoNumbers(l1, l2))

	fmt.Println("end: ", time.Now().UnixNano()-t)
}
