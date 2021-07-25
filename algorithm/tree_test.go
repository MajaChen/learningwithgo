package algorithm

import (
	"strconv"
	"testing")


func TestPush(t *testing.T){
	stack := Stack{elems: make([]interface{}, 0)}
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)

	t.Error(stack.Pop())
	t.Error(stack.Pop()) 
	t.Error(stack.Pop())

	t.Error(stack.IsEmpty())
}

func TestInTraverse(t *testing.T) {
	left := TreeNode{Val: 2}
	right := TreeNode{Val: 3}
	root := TreeNode{Val: 1, Left: &left, Right: &right}
	res := InTraverse(&root)
	t.Errorf("res is %v", res)
}

func TestPostTraverse(t *testing.T) {
	left := TreeNode{Val: 2}
	right := TreeNode{Val: 3}
	root := TreeNode{Val: 1, Left: &left, Right: &right}
	res := PostTraverse(&root)
	t.Errorf("res is %v", res)
}

func TestPostTraverse2(t *testing.T) {
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	res := PostTraverse(&root)
	t.Errorf("res is %v", res)
}

func TestSlice(t *testing.T){
	s := []int{0}
	s = s[1:]
	t.Error(s)
}

func TestLayerTraverse(t *testing.T){
	left := TreeNode{Val: 2}
	right := TreeNode{Val: 3}
	root := TreeNode{Val: 1, Left: &left, Right: &right}
	res := layerTraverse(&root)
	t.Errorf("res is %v", res)
}

func TestGetNodesPerLayer(t *testing.T){
	right := TreeNode{Val: 3}
	left := TreeNode{Val: 2,Left: &right}
	root := TreeNode{Val: 1, Left: &left}
	res := getNodesPerLayer(&root)
	t.Errorf("double res is %v", res)
}

func TestBstFromPreorder(t *testing.T){
	preorder := []int{2,1,3}
	root := bstFromPreorder(preorder)
	t.Error(root)
}

func TestBstFromPreorder2(t *testing.T){
	preorder := []int{1,2,3}
	root := bstFromPreorder(preorder)
	t.Error(root)
}

func TestDistanceK(t *testing.T){

	seven := TreeNode{Val: 7}
	four := TreeNode{Val: 4}
	two := TreeNode{Val: 2,Left: &seven,Right: &four}
	zero := TreeNode{Val: 0}
	five := TreeNode{Val: 5,Left:&zero,Right: &two}
	one := TreeNode{Val: 1}
	three := TreeNode{Val: 3,Left: &five,Right: &one}
	ten := TreeNode{Val: 10}
	eleven := TreeNode{Val: 11,Left: &three,Right: &ten}
	
	res := distanceK(&eleven,&seven,3)
	t.Errorf("res is %v",res)
}

func TestIsValidBST(t *testing.T){
	right := TreeNode{Val: 3}
	left := TreeNode{Val: 1}
	root := TreeNode{Val: 2, Left: &left,Right: &right}
	t.Error(isValidBST(&root))
}

func TestFormatInt(t *testing.T){
	i := 7
	t.Error(strconv.FormatInt(int64(i),2))
}

func TestStringIte(t *testing.T){
	s := "01"
	t.Error(s[0])
}

func TestCountNodes(t *testing.T){
	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}
	two := TreeNode{Val: 2,Left: &four,Right: &five}
	six := TreeNode{Val: 6}
	seven := TreeNode{Val: 7}
	three := TreeNode{Val: 3,Left: &six,Right: &seven}
	one := TreeNode{Val: 1,Left: &two,Right: &three}
	
	c := countNodes(&one)
	t.Errorf("count is %v",c)
}

func TestDepth(t *testing.T){
	one := TreeNode{Val: 1}
	d := depth(&one)
	
	t.Errorf("count is %v",d)
}

func TestSerilize(t *testing.T){
	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}
	two := TreeNode{Val: 2,Left: &four,Right: &five}
	six := TreeNode{Val: 6}
	seven := TreeNode{Val: 7}
	three := TreeNode{Val: 3,Left: &six,Right: &seven}
	one := TreeNode{Val: 1,Left: &two,Right: &three}

	res := serilize(&one)
	t.Errorf("res is %v",res) 
}

func TestSumRootToLeaf(t *testing.T){
	zero := TreeNode{Val: 0}
	one := TreeNode{Val: 1}
	root := TreeNode{Val: 1,Left: &zero,Right: &one}
	t.Error(sumRootToLeaf(&root))
}

func TestGoodNodes(t *testing.T){

	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}
	two := TreeNode{Val: 2,Left: &four,Right: &five}
	six := TreeNode{Val: 6}
	seven := TreeNode{Val: 7}
	three := TreeNode{Val: 3,Left: &six,Right: &seven}
	one := TreeNode{Val: 1,Left: &two,Right: &three}

	t.Error(goodNodes(&one))
}

func TestFindTilt(t *testing.T){
	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}
	two := TreeNode{Val: 2,Left: &four,Right: &five}
	six := TreeNode{Val: 6}
	seven := TreeNode{Val: 7}
	three := TreeNode{Val: 3,Left: &six,Right: &seven}
	one := TreeNode{Val: 1,Left: &two,Right: &three}

	t.Error(findTilt(&one))
}

func TestPathSum_2(t *testing.T){
	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}
	two := TreeNode{Val: 2,Left: &four,Right: &five}
	six := TreeNode{Val: 6}
	seven := TreeNode{Val: 7}
	three := TreeNode{Val: 3,Left: &six,Right: &seven}
	one := TreeNode{Val: 1,Left: &two,Right: &three}

	t.Error(pathSum_v2(&one,7))
}

func TestPathSum(t *testing.T){
	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}
	two := TreeNode{Val: 2,Left: &four,Right: &five}
	six := TreeNode{Val: 6}
	seven := TreeNode{Val: 7}
	three := TreeNode{Val: 3,Left: &six,Right: &seven}
	one := TreeNode{Val: 1,Left: &two,Right: &three}

	t.Error(pathSum(&one,7))
}

func TestMaxPathSum(t *testing.T){
	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}
	two := TreeNode{Val: 2,Left: &four,Right: &five}
	six := TreeNode{Val: 6}
	seven := TreeNode{Val: 7}
	three := TreeNode{Val: 3,Left: &six,Right: &seven}
	one := TreeNode{Val: 1,Left: &two,Right: &three}

	t.Error(maxPathSum(&one))
}