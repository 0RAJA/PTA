package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
)

const (
	info = `1. 压缩文件
2. 解压文件
3. 退出`
	InputError = `输入有误,请重新输入`
	Over       = `GoodBye`
	Yes        = "y"
)

var (
	logger  *log.Logger
	ErrFile = errors.New("路径有误或文件已存在")
	cin     = bufio.NewReader(os.Stdin)
)

// HFMNode 哈夫曼节点
type HFMNode struct {
	weight                 int
	parent, lChild, rChild int
	c                      byte //对应字符
}

// HFMTree 哈夫曼树
type HFMTree struct {
	tree   []*HFMNode      //哈夫曼节点数组
	code   map[byte][]byte //字符对应的编码
	weight map[byte]int    //字符对应的权重
	rs     []byte          //字符数组(用于按ascii排序用)
}

//找最小的两个值的下标
func searchMin(list []*HFMNode, length int) (int, int) {
	min1, min2 := math.MaxInt32, math.MaxInt32
	var index1, index2 int
	for i := 1; i <= length; i++ {
		if list[i].parent != 0 {
			continue
		}
		if min2 > list[i].weight {
			t1, t2 := list[i].weight, i
			if min1 > list[i].weight {
				t1, t2 = min1, index1
				min1, index1 = list[i].weight, i
			}
			min2, index2 = t1, t2
		}
	}
	return index1, index2
}

//为字符寻找其对应编码
func (t *HFMTree) searchCode(c byte) (ret []byte) {
	index := 0
	for i := 1; i < len(t.tree); i++ {
		v := t.tree[i]
		if v.c == c {
			index = i
			break
		}
	}
	for t.tree[index].parent != 0 {
		p := t.tree[index].parent
		if t.tree[p].lChild == index {
			ret = append([]byte{0}, ret...)
		} else {
			ret = append([]byte{1}, ret...)
		}
		index = p
	}
	return
}

// NewHFMTree 通过一个字符串生成一棵哈夫曼树
func NewHFMTree(str []byte) (tree *HFMTree) {
	weight := make(map[byte]int)
	for _, v := range str {
		weight[v]++
	}
	return NewHFMTreeWithWright(weight)
}

// NewHFMTreeWithWright 通过权重生成哈夫曼树
func NewHFMTreeWithWright(weight map[byte]int) (tree *HFMTree) {
	rs := make([]byte, 0, len(weight))
	for k := range weight {
		rs = append(rs, k)
	}
	return newHFMTreeWithWright(weight, rs)
}

func newHFMTreeWithWright(weight map[byte]int, rs []byte) (tree *HFMTree) {
	tree = &HFMTree{code: map[byte][]byte{}, weight: weight, rs: rs}
	sort.Slice(tree.rs, func(i, j int) bool {
		return tree.rs[i] < tree.rs[j]
	})
	tree.tree = make([]*HFMNode, 2*len(tree.weight))
	m := 1
	for _, c := range tree.rs {
		tree.tree[m] = &HFMNode{weight: tree.weight[c], c: c}
		m++
	}
	for i := len(tree.weight) + 1; i < len(tree.tree); i++ {
		min1, min2 := searchMin(tree.tree, i-1)
		if tree.tree[min1].weight == tree.tree[min2].weight {
			if min2 < min1 {
				min2, min1 = min1, min2
			}
		}
		tree.tree[i] = &HFMNode{
			weight: tree.tree[min1].weight + tree.tree[min2].weight,
			lChild: min1,
			rChild: min2,
		}
		tree.tree[min1].parent = i
		tree.tree[min2].parent = i
	}
	for c := range tree.weight {
		code := tree.searchCode(c)
		tree.code[c] = code
	}
	return tree
}

// ToCode 给字符串编码
func (t *HFMTree) ToCode(str []byte) (ret []byte) {
	for _, v := range str {
		var x byte
		for i, n := range t.code[v] {
			if n == 1 {
				x = x | (1 << i)
			}
		}
		ret = append(ret, x)
	}
	return
}

// DeCode 解码
func (t *HFMTree) DeCode(str []byte) (ret string) {
	p := t.tree[len(t.tree)-1]
	for _, v := range str {
		for {
			m := 1 & v
			v >>= 1
			if m == 0 {
				p = t.tree[p.lChild]
			} else {
				p = t.tree[p.rChild]
			}
			if p.lChild == 0 && p.rChild == 0 {
				ret += string(p.c)
				p = t.tree[len(t.tree)-1]
				break
			}
		}
	}
	return
}

// Sum 求平均码长
func (t *HFMTree) Sum() float64 {
	var nums = 0
	for _, v := range t.weight {
		nums += v
	}
	var sum float64 = 0
	for k, v := range t.code {
		sum += (float64(t.weight[k])) / float64(nums) * float64(len(v))
	}
	return sum
}

func (t *HFMTree) Serialize() []byte {
	marshal, _ := json.Marshal(t.weight)
	return append(marshal)
}

func Deserialization(data []byte) (map[byte]int, error) {
	weight := make(map[byte]int)
	err := json.Unmarshal(data, &weight)
	if err != nil {
		return nil, err
	}
	return weight, nil
}

func OpenFile(path string) (*os.File, error) {
	return os.Open(path)
}

func CreateFile(path string) (*os.File, error) {
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		return nil, ErrFile
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func ReadFrom(reader *bufio.Reader) (string, error) {
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str, nil
}

func InputPath(str string) (string, bool) {
	for {
		fmt.Println(str)
		path, err := ReadFrom(cin)
		if err != nil {
			logger.Println(err)
			fmt.Println("输入有误,err:", err)
			fmt.Println("是否重新输入?(y/n)")
			s, err := ReadFrom(cin)
			if err != nil {
				logger.Println(err)
			}
			if s != Yes {
				return "", false
			}
			continue
		}
		path = strings.Replace(path, "\\", "/", -1)
		return path, true
	}
}

func OpenPath(str string) (*os.File, bool) {
	for {
		path, ok := InputPath(str)
		if !ok {
			return nil, false
		}
		f, err := OpenFile(path)
		if err != nil {
			logger.Println(err)
			fmt.Println("文件读取失败,err:", err)
			fmt.Println("是否重新输入?(y/n)")
			s, err := ReadFrom(cin)
			if err != nil {
				logger.Println(err)
			}
			if s != Yes {
				return nil, false
			}
			continue
		}
		return f, true
	}
}

func CreatePath(str string, b bool) (f1, f2 *os.File, ok bool) {
	var err error
	for {
		path, ok := InputPath(str)
		if !ok {
			return nil, nil, false
		}
		f1, err = CreateFile(path)
		if err != nil {
			logger.Println(err)
			fmt.Println("文件创建失败,err:", err)
			fmt.Println("是否重新输入?(y/n)")
			s, err := ReadFrom(cin)
			if err != nil {
				logger.Println(err)
			}
			if s != Yes {
				return nil, nil, false
			}
			continue
		}
		if b {
			f2, err = CreateFile(path + ".key")
			if err != nil {
				logger.Println(err)
				fmt.Println("文件创建失败,err:", err)
				fmt.Println("是否重新输入?(y/n)")
				s, err := ReadFrom(cin)
				if err != nil {
					logger.Println(err)
				}
				if s != Yes {
					return nil, nil, false
				}
				continue
			}
		}
		return f1, f2, true
	}
}

func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		logger.Println(err)
	}
}

func EnCode() {
	openFile, ok := OpenPath("输入需要压缩的文件的路径")
	defer CloseFile(openFile)
	if !ok {
		return
	}
	data, err := ioutil.ReadAll(openFile)
	if err != nil {
		logger.Println(err)
		fmt.Println("文件读取失败,err:", err)
		return
	}
	tree := NewHFMTree(data)
	result := tree.ToCode(data)
	fmt.Println("压缩完成")
	saveFile, saveFile2, ok := CreatePath("请输入保存路径:", true)
	if !ok {
		return
	}
	defer CloseFile(saveFile)
	defer CloseFile(saveFile2)
	res := tree.Serialize()
	wg := sync.WaitGroup{}
	wg.Add(2)
	isOk := true
	save := func(f *os.File, data []byte) {
		defer wg.Done()
		_, err = f.Write(data)
		if err != nil {
			logger.Println(err)
			fmt.Println("写入失败,err:", err)
			isOk = false
			return
		}
	}
	save(saveFile, result)
	save(saveFile2, res)
	if isOk {
		fmt.Println("保存成功")
	}
}

func DeCode() {
	openFile, ok := OpenPath("输入需要解压的文件的路径")
	if !ok {
		return
	}
	defer CloseFile(openFile)
	keyFile, ok := OpenPath("输入密钥文件的路径")
	if !ok {
		return
	}
	defer CloseFile(keyFile)
	bf := bufio.NewReader(openFile)
	str, err := ReadFrom(bf)
	bf2 := bufio.NewReader(keyFile)
	w, err := ReadFrom(bf2)
	if err != nil {
		logger.Println(err)
		fmt.Println("读取失败,err:")
		return
	}
	weight, err := Deserialization([]byte(w))
	if err != nil {
		logger.Println(err)
		fmt.Println("解析失败,err:")
		return
	}
	tree := NewHFMTreeWithWright(weight)
	result := tree.DeCode([]byte(str))
	fmt.Println("解压成功")
	saveFile, _, ok := CreatePath("请输入保存路径:", false)
	if !ok {
		return
	}
	defer CloseFile(saveFile)
	_, err = saveFile.Write([]byte(result))
	if err != nil {
		fmt.Println("写入失败,err:", err)
		return
	}
	fmt.Println("保存成功")
}

func InitLog() (*os.File, error) {
	if info, err := os.Stat("./log"); err != nil || !info.IsDir() {
		if errors.Is(err, os.ErrNotExist) || !info.IsDir() {
			if err := os.Mkdir("./log", 0777); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return os.OpenFile("./log/log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
}

func init() {
	logFile, err := InitLog()
	if err != nil {
		fmt.Println("初始化日志有误:", err)
		return
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	for {
		fmt.Println(info)
		catch, err := ReadFrom(cin)
		if err != nil {
			fmt.Println(InputError)
			continue
		}
		switch catch {
		case "1":
			EnCode()
		case "2":
			DeCode()
		case "3":
			fmt.Println(Over)
			return
		default:
			fmt.Println(InputError)
		}
	}
}
