package queue  

import (  
   "context"  
   "log"  
   "time"  
   "webhook-pagamentos/sender"  
   "webhook-pagamentos/payment"
   //"os"

   redisClient "webhook-pagamentos/redis"  
)

func ProcessWebhooks(ctx context.Context, webhookQueue chan redisClient.WebhookPayload) {  
	for payload := range webhookQueue {  
	   go func(p redisClient.WebhookPayload) {  
		  backoffTime := time.Second  // starting backoff time  
		  maxBackoffTime := time.Hour // maximum backoff time  
		  retries := 0  
		  maxRetries := 5 
		  p.Data.Status = payment.ProcessPayment(payment.PaymentRequest{IdPagamento: p.Data.IdPagamento, Valor: p.Data.Valor, Status: p.Data.Status, DataCriacao: p.Data.DataCriacao})
 
		  for {  
			 err := sender.SendWebhook(p.Data, p.Url, p.WebhookId)  
			 if err == nil {  
				break  
			 }  
			 log.Println("Error sending webhook:", err)  
 
			 retries++  
			 if retries >= maxRetries {  
				log.Println("Max retries reached. Giving up on webhook:", p.WebhookId)  
				break  
			 }  
 
			 time.Sleep(backoffTime)  
 
			 // Double the backoff time for the next iteration, capped at the max  
			 backoffTime *= 2  
			 log.Println(backoffTime)  
			 if backoffTime > maxBackoffTime {  
				backoffTime = maxBackoffTime  
			 }  
		  }  
	   }(payload)  
	}  
 }
 