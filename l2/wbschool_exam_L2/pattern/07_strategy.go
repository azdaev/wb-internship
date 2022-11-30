package main

import "fmt"

/*
	Стратегия.
	Вид: Поведенческий.
	Суть паттерна - позволяет определять семейство свхожих алгоритмов и помещать каждый их них
в свой отдельный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

	+: Замена алгоритмов на лету
	+: Изолирует код и данные алгоритмов от остальных классов
	+: Упрощает добавление новых продуктов в программу
	+: Реализуется принцип Open/Closed
	-: Усложняет программу за счёт дополнительных классов
	-: Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую
*/

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

type EvictionAlgo interface {
	evict(c *Cache)
}

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
