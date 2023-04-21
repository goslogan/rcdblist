package rcdblist

import (
	"io"
	"strings"

	"github.com/gocarina/gocsv"
)

type DBInfo struct {
	Status           string  `csv:"Status"`
	DatabaseId       string  `csv:"Database ID"`
	DatabaseName     string  `csv:"Database name"`
	SubscriptionId   string  `csv:"Subscription ID"`
	SubscriptionName string  `csv:"Subscription name"`
	PublicEndpoint   string  `csv:"Public endpoint"`
	PrivateEndpoint  string  `csv:"Private endpoint"`
	MemoryLimit      float32 `csv:"Memory limit (MB)"`
	MemoryUsed       float32 `csv:"Memory used (MB)"`
	Throughput       string  `csv:"Throughput (Ops/sec)"` // Contains commas in values
	Modules          string  `csv:"Modules"`
	Options          string  `csv:"Options"`
	ShardType        string  `csv:"Shard type"`
	ShardCount       int     `csv:"Shard count"`
	ShardPrice       float32 `csv:"Shard price ($/hr)"`
	DataPriceHour    float32 `csv:"Database price ($/hr)"`
}

func ParseDBList(input io.Reader) ([]*DBInfo, error) {

	dbList := []*DBInfo{}

	err := gocsv.Unmarshal(input, &dbList)
	return dbList, err

}

func (d *DBInfo) GetModules() []string {
	if d.Modules == "N/A" || d.Modules == "" {
		return []string{}
	} else {
		return strings.Split(d.Modules, "; ")
	}
}

func (d *DBInfo) GetOptions() []string {
	if d.Options == "N/A" || d.Options == "" {
		return []string{}
	} else {
		return strings.Split(d.Options, "; ")
	}
}

func (d *DBInfo) Replication() bool {
	for _, opt := range d.GetOptions() {
		if opt == "In-Memory Replication" {
			return true
		}
	}

	return false
}

func (d *DBInfo) Persistence() bool {
	for _, opt := range d.GetOptions() {
		if opt == "Data Persistence" {
			return true
		}
	}

	return false
}
