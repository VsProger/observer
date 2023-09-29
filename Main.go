package main

import "fmt"

type Subject interface {
	register(observer Observer)
	unregister(observer Observer)
	notify()
}

type Observer interface {
	update(string)
	getID() string
}

type Customer struct {
	id string
}

type Item struct {
	name      string
	observers []Observer
	available bool
}

func (c *Customer) update(itemName string) {
	fmt.Printf("meesage has been sent to %s for %s", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("%s is now available\n", i.name)
	i.available = true
	i.notify()
}

func (i *Item) notify() {
	for _, observer := range i.observers {
		observer.update(i.name)
	}
}

func (i *Item) register(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) unregister(o Observer) {
	i.observers = removeFromList(i.observers, o)
}

func removeFromList(observers []Observer, observerToRemove Observer) []Observer {
	observersLength := len(observers)
	for i, o := range observers {
		if observerToRemove.getID() == o.getID() {
			observers[observersLength-1], observers[i] = observers[i], observers[observersLength-1]
			return observers[:observersLength-1]
		}
	}
	return observers
}

func main() {

	phoneItem := newItem("Iphone 15")

	observer1 := &Customer{id: "bbvbatvv@abc"}
	observer2 := &Customer{id: "abc"}
	phoneItem.register(observer1)
	phoneItem.unregister(observer1)
	phoneItem.register(observer2)
	phoneItem.updateAvailability()

}
