package rcutils

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/jszwec/csvutil"
)

// DBInfo represents a single line from the database export from Redis Cloud, representing
// information about a single database.
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

// Databases parses and reads the databases export from Redis Cloud.
func Databases(input io.Reader) ([]*DBInfo, error) {

	var dbList []*DBInfo
	reader := csv.NewReader(input)
	if decoder, err := csvutil.NewDecoder(reader); err != nil {
		return nil, err
	} else {
		decoder.Map = func(field, col string, v any) string {
			if field == "N/A" {
				if _, ok := v.(float32); ok {
					return "0.0"
				} else if _, ok := v.(int); ok {
					return "0"
				} else {
					return field
				}
			} else if field == "Fixed price" {
				return "NaN"
			}
			return field
		}

		for {
			var db DBInfo
			if err := decoder.Decode(&db); err == io.EOF {
				break
			} else if err != nil {
				return nil, err
			} else {
				dbList = append(dbList, &db)
			}
		}
	}

	return dbList, nil

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

func (d *DBInfo) Search() bool {
	for _, opt := range d.GetOptions() {
		if strings.Contains(opt, "Search") {
			return true
		}
	}

	return false
}
