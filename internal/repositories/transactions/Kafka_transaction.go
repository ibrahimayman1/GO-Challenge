package transactions

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

func Produce(transaction *Transaction) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic_transaction", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	obj, _ := json.Marshal(&transaction)
	conn.WriteMessages(kafka.Message{Value: []byte(obj)})
}

func Consume() {
	config := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "topic_transaction",
		MaxBytes: 30,
	}
	reader := kafka.NewReader(config)
	for {
		message, error := reader.ReadMessage(context.Background())
		if error != nil {

			continue
		}
		fmt.Println(time.Now().String() + "::message of transaction consumed:: " + string(message.Value))
		TransactionRepo.UpdateTransaction(string(message.Value))
	}
}
