package main

type MyInit int

type MapKey interface {
	~int | float64
}

func Keys[K MapKey, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func main() {

}
