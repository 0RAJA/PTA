package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filePath = `D:\桌面\1.pdf`
const outPutPath = `D:\桌面\2`
const depressPath = `D:\桌面\result.pdf`
const contentBuffer = 5000000

type TreeNode struct {
	Val   int
	Times int
	Left  *TreeNode
	Right *TreeNode
}

type treeHeap []*TreeNode

func (p treeHeap) Less(i, j int) bool {
	return p[i].Times <= p[j].Times
}

func (p treeHeap) Len() int {
	return len(p)
}

func (p treeHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *treeHeap) Push(node interface{}) {
	*p = append(*p, node.(*TreeNode))
}

func (p *treeHeap) Pop() interface{} {
	n := len(*p)
	t := (*p)[n-1]
	*p = (*p)[:n-1]
	return t
}

// 测试压缩解压缩的正确性
func main() {
	HuffmanEncoding(filePath, outPutPath)
	depress(outPutPath, depressPath)
	originMD5, _ := MD5File(filePath)
	recoverMD5, _ := MD5File(depressPath)
	fmt.Println(originMD5 == recoverMD5)
}

func HuffmanEncoding(filePath, outPath string) {
	// 思路： 1. 读取文本内容，存放到内存中，或者以流的形式读取文本内容，构建二叉树即可。
	// 统计每个字出现的频次
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()
	// 我们不需要关心总量是多少，因为分母是固定的，只需要知道频率按次数排序即可。
	imap := getFrequencyMap(file)
	plist := make(treeHeap, 0)
	// 遍历map ,将键值对存入pair，然后按频率排序
	for k, v := range imap {
		plist = append(plist, &TreeNode{Val: k, Times: v})
	}
	sort.Sort(plist)
	//如果文件是空的，还构造个屁
	if len(plist) == 0 {
		return
	}
	hTree := initHuffmanTree(plist)
	/*遍历哈弗曼树，生成哈夫曼编码表(正表，用于编码),key(ASSCII),value(路径痕迹)*/
	encodeMap := make(map[int]string)
	createEncodingTable(hTree, encodeMap)

	// 将输入文件的字符通过码表编码，输出到另一个文件 , 压缩模块完成
	encoding(filePath, outPath, encodeMap)

}

func writeTable(path string, codeMap map[int]string, left int) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	// 第一行，写入文件头的长度
	var buff bytes.Buffer
	buff.WriteString(strconv.Itoa(len(codeMap)+1) + "\n")
	for k, v := range codeMap {
		buff.WriteString(strconv.Itoa(k) + ":" + v + "\n")
	}
	buff.WriteString(strconv.Itoa(left) + "\n")
	file.WriteString(buff.String())
	file.Close()
}

/* 一次性读入，存到string或者buffer.string中 */
func encoding(inPath string, outPath string, encodeMap map[int]string) {
	/*1.先尝试一次性读入*/
	inFile, err := os.Open(inPath)
	defer inFile.Close()
	if err != nil {
		return
	}
	reader := bufio.NewReader(inFile)
	fileContent := make([]byte, contentBuffer)
	count, _ := reader.Read(fileContent)
	var buff bytes.Buffer
	//string编码
	for i := 0; i < count; i++ {
		v := fileContent[i]
		if code, ok := encodeMap[int(v)]; len(code) != 0 && ok {
			buff.WriteString(code)
		}
	}
	res := make([]byte, 0)
	var buf byte = 0
	//bit编码
	//TODO 记录bit剩余位，很简单只要对buff.bytes取长度对8取余即可
	for idx, bit := range buff.Bytes() {
		//每八个位使用一个byte读取，结果存入res数组即可
		pos := idx % 8
		if pos == 0 && idx > 0 {
			res = append(res, buf)
			buf = 0
		}
		if bit == '1' {
			buf |= 1 << pos
		}
	}
	//TODO 这个left是剩余待处理的位数
	left := buff.Len() % 8
	res = append(res, buf)
	// 将编码数组写入文件 , TODO 先将码表和left数写入文件，解码时在开头读取
	writeTable(outPath, encodeMap, left)
	outFile, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	wcount, err := outFile.Write(res)
	if err != nil {
		fmt.Println(wcount)
		return
	}
}

//码表，考虑到性能必须要生成 map, key(int对应ASSCII🐎，string对应bit编码，后续转成bit)
func createEncodingTable(node *TreeNode, encodeMap map[int]string) {
	/*思路：回溯遍历二叉树，用byte记录0 1🐎，遇到叶子节点就转换成string存入码表
	  遍历顺序：根左右
	*/
	tmp := make([]byte, 0)
	var depth func(treeNode *TreeNode)
	depth = func(root *TreeNode) {
		//如果已经遍历到空，返回
		if root == nil {
			return
		}
		//如果遍历到的是叶子节点 , byte转换成string，入表
		if root.Left == nil && root.Right == nil {
			encodeMap[root.Val] = string(tmp)
		}
		//如果是普通节点，左右递归回溯即可 #规则： 左0右1
		tmp = append(tmp, '0')
		depth(root.Left)
		tmp[len(tmp)-1] = '1'
		depth(root.Right)
		tmp = tmp[:len(tmp)-1]
	}
	depth(node)
}

func initHuffmanTree(plist treeHeap) *TreeNode {
	//使用优先队列构造最小路径权值哈夫曼树
	heap.Init(&plist)
	for plist.Len() > 1 {
		t1 := heap.Pop(&plist).(*TreeNode)
		t2 := heap.Pop(&plist).(*TreeNode)
		root := &TreeNode{Times: t1.Times + t2.Times}
		if t1.Times > t2.Times {
			root.Right, root.Left = t1, t2
		} else {
			root.Right, root.Left = t2, t1
		}
		heap.Push(&plist, root)
	}
	return plist[0]
}

func getFrequencyMap(file *os.File) map[int]int {
	imap := make(map[int]int)
	// 读入文件数据，readline 记入map中，统计频次
	// 注意：Create不区分文件名大小写
	reader := bufio.NewReader(file)
	buffer := make([]byte, contentBuffer)
	readCount, _ := reader.Read(buffer)
	for i := 0; i < readCount; i++ {
		imap[int(buffer[i])]++
	}
	return imap
}

func depress(inPath, depressPath string) {
	// originPath 原文件(或者可以传入码表)， inPath 读入被压缩的文件 , depressPath 还原后的输出路径
	encodeMap := make(map[int]string)
	decodeMap := make(map[string]int)
	//2.读入压缩文件
	compressFile, _ := os.Open(inPath)
	// br 读取文件头 ,返回偏移量
	br := bufio.NewReader(compressFile)
	left, offset := readTable(*br, encodeMap)
	for idx, v := range encodeMap {
		decodeMap[v] = idx
	}
	// 解码string暂存区
	var buff bytes.Buffer
	// 编码bytes暂存区
	codeBuff := make([]byte, contentBuffer)
	codeLen, _ := compressFile.ReadAt(codeBuff, int64(offset))
	//遍历解码 , 读取比特
	for i := 0; i < codeLen; i++ {
		//对每个byte单独进行位运算转string
		perByte := codeBuff[i]
		for j := 0; j < 8; j++ {
			//与运算
			buff.WriteString(strconv.Itoa(int((perByte >> j) & 1)))
		}
	}
	// 对照码表，解码string , 对8取余目的是解决正好读满8个bit的情况发生
	contentStr := buff.String()[:buff.Len()-(8-left)%8]
	bytes := make([]byte, 0)
	//用切片读contenStr即可
	for star, end := 0, 1; end <= len(contentStr); {
		charValue, ok := decodeMap[contentStr[star:end]]
		if ok {
			bytes = append(bytes, byte(charValue))
			star = end
		}
		end++
	}

	depressFile, _ := os.Create(depressPath)
	depressFile.Write(bytes)
	depressFile.Close()
}

func readTable(br bufio.Reader, encodeMap map[int]string) (int, int) {
	lineStr, _, _ := br.ReadLine()
	lines, _ := strconv.Atoi(string(lineStr))
	for i := 0; i < lines-1; i++ {
		lineContent, _, _ := br.ReadLine()
		kvArr := strings.Split(string(lineContent), ":")
		k, v := kvArr[0], kvArr[1]
		kNum, _ := strconv.Atoi(k)
		encodeMap[kNum] = v
	}
	leftStr, _, _ := br.ReadLine()
	left, _ := strconv.Atoi(string(leftStr))
	return left, br.Size() - br.Buffered()
}

func MD5Bytes(s []byte) string {
	ret := md5.Sum(s)
	return hex.EncodeToString(ret[:])
}

func MD5File(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return MD5Bytes(data), nil
}
