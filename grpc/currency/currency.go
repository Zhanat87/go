package currency

// protoc -I grpc/currency/ grpc/currency/currency.proto --go_out=plugins=grpc:grpc/currency
type Currency struct {
	Title       string  `json:"title" xml:"title"`
	PubDate     string  `json:"pubDate" xml:"pubDate"`
	Description float32 `json:"description" xml:"description"`
	Quant       int32   `json:"quant" xml:"quant"`
	Index       string  `json:"index" xml:"index"`
	Change      float32 `json:"change" xml:"change"`
}
