package tinytest

type Record struct {
	key   int
	value int
}

func MapValues() map[int]int {
	// this function should return a map where the key is the key of the record and the value is the sum of all values for that key
	// it should run 10 concurrent goroutines to process the records

	// those are records to process, it's a slice of []Record
	// example: [{key: 1, value: 1}, {key: 1, value: 2}, {key: 2, value: 2}, {key: 2, value: 3}]
	records := createRecords()
	_ = records

	return nil
}

func createRecords() []Record {
	var records []Record
	for i := 1; i < 100; i++ {
		record := Record{
			key:   i,
			value: i,
		}
		records = append(records, record)
		record.value = i + 1
		records = append(records, record)
	}
	return records
}
