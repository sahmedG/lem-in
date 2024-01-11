package farm

/* used for ants */
type AntQueue struct {
	items []*Ant
}

func (q *AntQueue) Enqueue(item *Ant) {
	q.items = append(q.items, item)
}

func (q *AntQueue) Dequeue() *Ant {
	if len(q.items) == 0 {
		return nil // Queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
