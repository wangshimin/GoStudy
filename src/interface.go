package main

import "fmt"

// 空接口
type empty interface {
}

/**** 接口中只有方法声明，没有数据字段 ********/
type USB interface {
	Name() string
	Connecter //	接口可以匿名嵌入其他接口，或嵌入到结构中
}

type Connecter interface {
	Connect() //	接口支持匿名字段方法
}

/* 在GO中，用户定义的类型实现了接口声明的所有方法
 * 那么这个类型就隐式地实现了这个接口
 * 而这个类型就是实现这个接口类型的实例
 */
//	因此，PhoneConnecter是接口  USB、Connecter的实例
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string { // PhoneConnecter实现了Name方法，方法接收者是值类型
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	// var a USB
	// a := PhoneConnecter{"OPPO"} // 字面值进行初始化
	// a.Connect()
	// Disconnect(a)
	// // DisconnectB(a)

	/******* 接口转换：可将拥有超集的接口转换为子集的接口 **********/
	b := PhoneConnecter{"huawei"}
	var c Connecter
	c = Connecter(b) // 类型强制转换
	c.Connect()

	/******* 对象赋值给接口时会发生拷贝，接口内部存储的是只想这个复制品的指针，即无法修改复制品的状态，也无法获取指针 *******/
	b.name = "macbook"
	fmt.Println(c) // huawei

	/********* 只有当接口存储的类型和对象都为nil时，接口才等于nil ********/
	var d interface{}
	fmt.Println(d == nil)
	var p *int = nil
	d = p
	fmt.Println(d == nil)
}

/****  使用type switch则可针对空接口进行比较全面的类型判断 ******/
func Disconnect(usb interface{}) { //	 空接口可以作为任何类型数据的容器
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected...", v.name)
	default:
		fmt.Println("Unknow device")
	}

}

/******  通过类型断言的ok pattern可以判断接口中的数据类型 ********/
func DisconnectB(usb USB) {
	if pc, status := usb.(PhoneConnecter); status {
		fmt.Println("Disconnectd:", pc.name)
		return
	}
	fmt.Println("Unknow device.")
}
