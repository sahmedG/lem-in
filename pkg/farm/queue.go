package farm

type Queue struct {
	items []*Room
}

func (q *Queue) Enqueue(item *Room) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() *Room {
	if len(q.items) == 0 {
		return nil // Queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
