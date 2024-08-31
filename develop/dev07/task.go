package or

import (
	"sync"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 { // Проверка, что channels не пустой.
		return nil
	} else if len(channels) == 1 { // Если пришел всего один канал, то его и возвращаем.
		return channels[0]
	}

	done := make(chan interface{}) // Создаем общий канал.
	var wg sync.WaitGroup
	wg.Add(len(channels)) // Создаем счетчик, по количеству пришедших каналов.
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			for v := range ch {
				done <- v // Записываем данные в канал.
			}
			defer wg.Done() // Уменьшаем счетчик горутин.
		}(ch)
	}

	go func() {
		wg.Wait()   // Ожидаем выполнения всех горутин.
		close(done) // Закрываем канал.
	}()

	return done
}
