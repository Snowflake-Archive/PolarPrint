package types

type QueuePrintRequestBody struct {
	FilePath string `json:"file"`
	Quantity int    `json:"quantity"`
}

type QueueItem struct {
	Id       int    `json:"id"`
	File     string `json:"file"`
	Quantity int    `json:"quantity"`
}

type Cluster struct {
	Id       int      `json:"id"`
	Printers []string `json:"printers"`
	Key      string   `json:"key"`
}
