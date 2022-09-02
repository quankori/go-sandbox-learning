package advanced

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Khai báo một instance mutex
var mu sync.Mutex

//Khai báo biến count được truy cập bởi tất cả các routine
var count = 0

// Sao chép count vào temp, thực hiện một vài xử lý (tăng dần) và lưu lại vào count
// tạm dừng một khoảng ngẫu nhiên được thêm vào giữa lúc đọc và ghi count
func process(n int) {
	//Vòng lặp tăng count 10 lần
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		//Bắt đầu khoá ở đây
		mu.Lock()
		temp := count
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count = temp
		// Mở khoá
		mu.Unlock()
	}
	fmt.Println("Count after i="+strconv.Itoa(n)+" Count:", strconv.Itoa(count))
}

// Mutex func
func mutex() {
	//lặp gọi process() 3 lần
	for i := 1; i < 4; i++ {
		go process(i)
	}

	//Tạm dừng để đợi cho tất cả routine hoàn thành
	time.Sleep(25 * time.Second)
	fmt.Println("Final Count:", count)
}
