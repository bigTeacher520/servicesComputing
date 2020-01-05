package main

import "fmt"

type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	var length int=len(nodes)
	var index int=length/2
	for index!=1{
		down(nodes[:],index,length)
		index--
	}
	down(nodes[:],index,length)
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	for i < n {
		var left=i*2
		var right=i*2+1
		if right>=n && left >=n{
			break
		}
		if right>=n && left<n{
			if nodes[left].Value<nodes[i].Value{
				var node Node=nodes[i]
				nodes[i]=nodes[left]
				nodes[left]=node
			}
			i=left
		}
		if right<n && left<n{
			var index int=left
			if nodes[left].Value>nodes[right].Value{
				index=right
			}
			if nodes[index].Value<nodes[i].Value{
				var node Node=nodes[i]
				nodes[i]=nodes[left]
				nodes[index]=node
			}
			i=index
		}
		
	}
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	for j>1{
		var index int =j/2
		if nodes[j].Value<nodes[index].Value{
			var node Node=nodes[j]
				nodes[j]=nodes[index]
				nodes[index]=node
		}
		j=index
	}
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	

	var index int=1
	for index<len(nodes){
		var left int=2*index
		var right int=2*index+1
		var mid int=left
		if left>=len(nodes) && right>=len(nodes){
			break
		}
		if(left<len(nodes) && right >=len(nodes)){
			mid=left
		}
		if(left<len(nodes) && right<len(nodes)){
			if nodes[left].Value>nodes[right].Value{
				mid=right
			}
		}
		var nod Node=nodes[index]
		nodes[index]=nodes[mid]
		nodes[mid]=nod
		index=mid
	}
	var nod Node=nodes[index]
		nodes[index]=nodes[len(nodes)-1]
		nodes[len(nodes)-1]=nod
	if index!=len(nodes)-1{
		up(nodes[:],index)
	}
	ha:=nodes[0:len(nodes)-1]
	return nodes[len(nodes)-1],ha
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	//var leng=len(nodes)+1
	var nodes2 [11]Node
	for i:=0;i<len(nodes);i++{
		nodes2[i]=nodes[i]
	}
	nodes2[len(nodes2)-1]=node
	up(nodes2[:],len(nodes2)-1)
	return nodes2[:]
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	var index int =0
	for i:=0;i<len(nodes);i++{
		if nodes[i].Value==node.Value{
			index=i
			break
		}
	}
	if index==0{
		return nodes[:]
	}
	fmt.Printf("%d \n",index)
	for index<len(nodes){
		var left int=2*index
		var right int=2*index+1
		var mid int=left
		if left>=len(nodes) && right>=len(nodes){
			break
		}
		if(left<len(nodes) && right >=len(nodes)){
			mid=left
		}
		if(left<len(nodes) && right<len(nodes)){
			if nodes[left].Value>nodes[right].Value{
				mid=right
			}
		}
		var nod Node=nodes[index]
		nodes[index]=nodes[mid]
		nodes[mid]=nod
		index=mid
	}
	var nod Node=nodes[index]
		nodes[index]=nodes[len(nodes)-1]
		nodes[len(nodes)-1]=nod
	if index!=len(nodes)-1{
		up(nodes[:],index)
	}
	ha:=nodes[0:len(nodes)-1]
	return ha
}
func construct(nodes []Node,array []int){
	for i:=0;i<len(array);i++{
		nodes[i].Value=array[i]
	}
}
func print(nodes []Node){
	for i:=0;i<len(nodes);i++{
		fmt.Printf("%d ",nodes[i].Value)
	}
	fmt.Println("\n");
}
func main() {
	var nodes [11]Node
	var array =[11]int{0,5, 1, 13, 3, 16, 7, 10, 14, 6, 9}
	construct(nodes[:],array[:])
	Init(nodes[:])
	print(nodes[:])
	var m,no=Pop(nodes[:])
	
	fmt.Printf("%d\n",m.Value)
	print(no)
	var nod Node
	nod.Value=2
	no=Push(nod,no[:])
	print(no[:])
	nod.Value=3
	
	no=Remove(no[:],nod)
	print(no[:])
}