package ItemChan

import "log"

func ItemProcedure() chan interface{} {
	out := make(chan interface{})
	var count int
	go func() {
		for {
			r := <-out
			//log.Printf("get Value: %v",r)
			log.Printf("第%d个结果为%v", count, r)
			count++
		}
	}()

	return out
}
