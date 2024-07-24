package main

import (
	"73HW/consumer/service"
)

func main() {
	rabbitMQService := service.NewRabbitMQService()

	go rabbitMQService.ConsumeMessages("report.*", "report_queue_1")
	go rabbitMQService.ConsumeMessages("*.updated", "report_queue_2")
	go rabbitMQService.ConsumeMessages("report.#", "report_queue_3")

	select {}
}
