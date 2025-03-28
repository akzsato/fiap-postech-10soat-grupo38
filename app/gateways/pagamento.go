package gateways

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type pagamentoRepository struct {
	database    mongo.Database
	collection  string
	redisClient *redis.Client
}

func NewPagamentoRepository(db mongo.Database, collection string) repository.PagamentoRepository {
	// Initialize the Redis client
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddress, // Redis server address
		Password: "",           // No password
		DB:       0,            // Default DB
	})

	// Ping Redis to check the connection
	//ctx := context.Background()
	_, err := redisClient.Ping().Result()
	if err != nil {
		//panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
		fmt.Printf("Failed to connect to Redis: %v\n", err)
	}

	return &pagamentoRepository{
		database:    db,
		collection:  collection,
		redisClient: redisClient,
	}
}

func (pg *pagamentoRepository) ConfirmarPagamento(ctx context.Context, pagamento *entities.Pagamento) error {
	filter := bson.M{"identificacao": pagamento.IdPagamento}

	update := bson.M{"$set": bson.M{"statuspagamento": pagamento.Status}}

	result, err := pg.database.Collection("pedido").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update pedido status: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no pedido found with idPedido: %s", pagamento.IdPagamento)
	}

	fmt.Printf("Pedido %s status updated to %s\n", pagamento.IdPagamento, pagamento.Status)
	return nil
}

func (pg *pagamentoRepository) EnviarPagamento(c context.Context, pagamento *entities.Pagamento) error {

	paymentData := map[string]interface{}{
		"url":       "http://host.docker.internal:8080/pagamento/confirmar",
		"webhookId": uuid.New().String(),
		"data": map[string]interface{}{
			"idPagamento": pagamento.IdPagamento,
			"valor":       pagamento.Valor,
			"status":      pagamento.Status,
			"dataCriacao": time.Now().Format("02/01/2006, 15:04:05"),
		},
	}

	webhookPayloadJSON, err := json.Marshal(paymentData)
	if err != nil {
		return fmt.Errorf("failed to marshal payment data: %v", err)
	}

	err = pg.redisClient.Publish("payments", webhookPayloadJSON).Err()
	if err != nil {
		return fmt.Errorf("failed to publish payment to Redis: %v", err)
	}

	fmt.Println("Payment sent to Redis:", string(webhookPayloadJSON))
	return nil
}

// RegisterPaymentRoutes registers the HTTP routes for payments.
func RegisterPaymentRoutes(redisClient *redis.Client) {
	http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		paymentData := getPayment()
		webhookPayloadJSON, err := json.Marshal(paymentData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = redisClient.Publish("payments", webhookPayloadJSON).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(webhookPayloadJSON)
		println("Payment sent to Redis")
	})
}

func getPayment() map[string]interface{} {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdxyzpqr")
	paymentID := make([]rune, 5)
	for i := range paymentID {
		paymentID[i] = letters[rand.Intn(len(letters))]
	}

	return map[string]interface{}{
		"url":       os.Getenv("WEBHOOK_ADDRESS"),
		"webhookId": uuid.New().String(),
		"data": map[string]interface{}{
			"id":      "123456",
			"payment": fmt.Sprintf("PY-%s", string(paymentID)),
			"event":   []string{"accepted", "completed", "canceled"}[rand.Intn(3)],
			"created": time.Now().Format("02/01/2006, 15:04:05"),
		},
	}
}