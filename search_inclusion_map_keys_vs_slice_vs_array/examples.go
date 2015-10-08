package search

import s "strconv"

type (
	dMap   map[string]struct{}
	dSlice []string
)

func newDSlice(count int) (d dSlice) {
	d = make(dSlice, count)

	for i, _ := range d {
		d[i] = s.Itoa(i)
	}

	return
}

func newDMap(count int) (d dMap) {
	d = make(dMap, count)

	for i, _ := range make([]struct{}, count) {
		d[s.Itoa(i)] = struct{}{}
	}

	return
}

func (d dMap) includes(key string) (b bool) {
	if _, ok := d[key]; ok {
		b = true
	}

	return
}

func (d dSlice) includes(key string) (b bool) {
	for _, v := range d {
		if v == key {
			b = true
			break
		}
	}

	return
}
