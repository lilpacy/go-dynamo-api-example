package domain

// Item is struct for DynamoDB
type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

var ItemMock = Item{
	MyHashKey:  "MyHash",
	MyRangeKey: 1,
	MyText:     "My First Text",
}
