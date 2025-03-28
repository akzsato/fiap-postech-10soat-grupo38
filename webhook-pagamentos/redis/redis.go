package redis  

import (  
   "context"  
   "encoding/json"  
   "log"  

   "github.com/go-redis/redis/v8"  
)  

// WebhookPayload defines the structure of the data expected  
// to be received from Redis, including URL, Webhook ID, and relevant data.  
type WebhookPayload struct {  
   Url       string `json:"url"`  
   WebhookId string `json:"webhookId"`  
   Data      struct {  
      IdPagamento   string `json:"idPagamento"`  
      Valor 		string `json:"valor"`  
      Status   		string `json:"status"`  
      DataCriacao	string `json:"dataCriacao"`  
   } `json:"data"`  
}

func Subscribe(ctx context.Context, client *redis.Client, webhookQueue chan WebhookPayload) error {  
	// Subscribe to the "webhooks" channel in Redis  
	pubSub := client.Subscribe(ctx, "payments")  
 
	// Ensure that the PubSub connection is closed when the function exits  
	defer func(pubSub *redis.PubSub) {  
	   if err := pubSub.Close(); err != nil {  
		  log.Println("Error closing PubSub:", err)  
	   }  
	}(pubSub)  
 
	var payload WebhookPayload  
 
	// Infinite loop to continuously receive messages from the "webhooks" channel  
	for {  
	   // Receive a message from the channel  
	   msg, err := pubSub.ReceiveMessage(ctx)  
	   if err != nil {  
		  return err // Return the error if there's an issue receiving the message  
	   }  
 
	   // Unmarshal the JSON payload into the WebhookPayload structure  
	   err = json.Unmarshal([]byte(msg.Payload), &payload)  
	   if err != nil {  
		  println(msg.Payload)
		  log.Println("Error unmarshalling payload:", err)  
		  continue // Continue with the next message if there's an error unmarshalling  
	   }  
 
	   webhookQueue <- payload // Sending the payload to the channel  
	}  
 }