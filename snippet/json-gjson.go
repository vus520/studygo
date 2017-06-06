package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

var TestJson = `{"took":200,"timed_out":false,"_shards":{"total":11,"successful":11,"failed":0},"hits":{"total":648875,"max_score":0.0,"hits":[]},"aggregations":{"day":{"doc_count_error_upper_bound":0,"sum_other_doc_count":0,"buckets":[{"key":"20170417","doc_count":648875,"channel":{"doc_count_error_upper_bound":0,"sum_other_doc_count":0,"buckets":[{"key":"kingpin_09","doc_count":648875,"acttype":{"doc_count_error_upper_bound":0,"sum_other_doc_count":0,"buckets":[{"key":"show","doc_count":406595,"ct":{"value":139409}},{"key":"click","doc_count":242280,"ct":{"value":111973}}]}}]}}]}}}`

func main() {
	result := gjson.Get(TestJson, "aggregations")
	iter(result)
}

func iter(value gjson.Result) []string {
	k := []string{}

	if value.Type.String() == "JSON" {
		value.ForEach(func(key, value gjson.Result) bool {

			a := gjson.Get(value.Raw, "value")
			if a.Exists() {
				fmt.Println(key, "'s value = ", a.String())
			}

			b := gjson.Get(value.Raw, "buckets")
			if b.Exists() {
				iterArray(b)
			}

			fmt.Println(key)
			return true
		})
	}

	return k
}

func iterArray(value gjson.Result) {
	value.ForEach(func(key, value gjson.Result) bool {
		iter(value)
		return true
	})
}
