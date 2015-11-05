package main

import "fmt"

// 2,3,4,…というシーケンスをチャネル'ch'に送信
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // 'i'をチャネル'ch'に送信
	}
}

// チャネル'in'の値をチャネル'out'にコピー
// ただし'prime'で割り切れる値を除く
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // 'src'から受信した値でループ
		if i%prime != 0 {
			dst <- i // 'i'をチャネル'dst'に送信
		}
	}
}

// 素数のふるい：フィルターを数珠つなぎに組み合わせて処理する
func sieve() {
	ch := make(chan int) // 新しいチャネルを作成
	go generate(ch)      // generate()をサブプロセスとして開始
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
