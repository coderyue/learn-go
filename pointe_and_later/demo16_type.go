package main

import "time"

func main() {

}

type myDuration = time.Duration

// 解决这个问题有两个方法， 1，在time包下 给duratio添加方法， 2，定义一个新的类型，而不是起别名，把等号去掉即可
func (m myDuration) SimpleSet() { //cannot define new methods on non-local type time.Duration
}
