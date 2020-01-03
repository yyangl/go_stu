package dht

import (
	"fmt"
	"github.com/zeebo/bencode"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//好友节点，带你进入DHT网络
var BOOTSTRAP = []string{
	"67.215.246.10:6881",
	"91.121.59.153:6881",
	"82.221.103.244:6881",
	"212.129.33.50:6881",
}

const (
	try                = 3
	bucketExpiredAfter = 180 * time.Second
	nodeExpiredAfter   = 60 * time.Second
)

type DHT struct {
	// candidates are udp, udp4, udp6
	Network string
	// K桶的数量
	KBucketSize int
	// 监听的地址
	Address string
	// K桶的过期时间
	KBucketExpiredAfter time.Duration
	// node节点的过期时间
	NodeExpiredAfter time.Duration
	// 重试次数
	Try   int
	conn  *net.UDPConn
	route *Route
	nid   Id
}

func New(addr string) *DHT {
	return &DHT{
		Network:             "udp4",
		Try:                 try,
		KBucketExpiredAfter: bucketExpiredAfter,
		NodeExpiredAfter:    nodeExpiredAfter,
		Address:             addr,
		route:               NewRoute(),
		nid:                 GenerateId(),
	}
}

func (d *DHT) init() {
	listener, err := net.ListenPacket(d.Network, d.Address)
	if err != nil {
		panic(err)
	}
	// 断言监听的地址是udp连接
	d.conn = listener.(*net.UDPConn)
	// 监听udp传来的数据
	go d.ListenerAndServer()
}

func (d *DHT) Run() {
	// 初始化DHT
	d.init()
	// 启动
	go d.bootStrap()

	// 监听table里面的数据

	// 监听返回的数据

	// 系统关闭信号
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(twg *sync.WaitGroup) {
		sig := make(chan os.Signal, 2)
		signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
		<-sig

		// 停止各部分
		twg.Done()
	}(wg)

	wg.Wait()
}

// 启动DHT
func (d *DHT) bootStrap() {
	for _, host := range BOOTSTRAP {
		addr, err := net.ResolveUDPAddr("udp", host)
		if err != nil {
			log.Fatalf("节点错误%s", host)
			return
		}
		node := new(Node)
		node.addr = addr
		d.findNode(node)
	}
}

func (d *DHT) findNode(node *Node) {
	var id Id
	if node.id != nil {
		id = node.id.Neighbor(d.nid)
	} else {
		id = d.nid
	}
	v := make(map[string]interface{})
	v["t"] = fmt.Sprintf("%d", rand.Intn(100))
	v["y"] = "q"
	v["q"] = "find_node"
	args := make(map[string]string)
	args["id"] = string(id)
	args["target"] = string(GenerateId()) //查找自己，找到离自己较近的节点
	v["a"] = args
	data, err := bencode.EncodeBytes(v)
	if err != nil {
		log.Printf("编码失败%s", v)
		return
	}
	//log.Printf("发送数据%s",data)
	_ = d.send([]byte(data), node.addr)
}

func (d *DHT) send(data []byte, addr *net.UDPAddr) error {
	_, err := d.conn.WriteToUDP(data, addr)
	if err != nil {
		log.Fatalf("发送udp数据失败，地址%v", addr)
	} else {
		log.Printf("发送udp数据成功，地址%v", addr)
	}
	return err
}

func (d *DHT) ListenerAndServer() {
	b := make([]byte, 1000)
	for {
		n, addr, err := d.conn.ReadFromUDP(b)
		if err != nil {
			continue
		}
		//log.Printf("收到消息,addr%v",addr)
		go d.Decode(b[:n], addr)
	}
}

func (d *DHT) Decode(data []byte, addr *net.UDPAddr) {
	message := make(map[string]interface{})
	err := bencode.DecodeBytes(data, &message)
	if err != nil {
		log.Fatalf("解码回复消息失败%v", err)
	}
	if _, ok := message["t"]; !ok {
		return
	}
	if _, ok := message["y"]; !ok {
		return
	}

	switch message["y"] {
	case "r":
		if r, ok := message["r"].(map[string]interface{}); ok {
			// 向路由表中插入节点
			d.putNodes(r)
		}
		break
	case "q":
		break
	default:
		log.Print("未知的类型")
	}
}

func (d *DHT) putNodes(r map[string]interface{}) {
	if nodeStr, ok := r["nodes"].(string); ok {
		nodes := ParseByteStream([]byte(nodeStr))
		//fmt.Printf("%v",nodes)
		for _, node := range nodes {
			fmt.Printf("%v\n", node.id)
		}
	}
}

func (d *DHT) appendNode(node *Node) {
	if node.id.Int() == d.nid.Int() {
		return
	}
	d.route.appendNode(node)
}
