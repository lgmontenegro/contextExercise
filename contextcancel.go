package main

import (
	"context"
  "fmt"
  "time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	anyChannel := make(chan int)
	for i := 0; i <= 10; i++ {
		go connectAndRetrieve(ctx, i, anyChannel)
	}
  
  value := <- anyChannel
  cancel()
  
  fmt.Printf("usando o valor %d recebido.\n",value)
  
  time.Sleep(1*time.Second)
}

func connectAndRetrieve(ctx context.Context, id int, anyChannel chan int) {
  fmt.Printf("%d começou.\n", id)
  for i:=0;i<1000000;i++ {
    select {
    case <-ctx.Done():
      fmt.Printf("%d cancelada.\n", id)
      return
    default:
      continue
    }
  }
  fmt.Printf("%d terminou.\n", id)
  anyChannel <- id
}
