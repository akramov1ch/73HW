package model

type Message struct {
    RoutingKey string `json:"routing_key"`
    Content    string `json:"content"`
}
