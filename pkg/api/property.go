package api

import "encoding/json"

type Property func(b map[string]interface{})

func DeviceIds(ids []string) Property {
	return func(b map[string]interface{}) {
		b["device_ids"] = ids
	}
}

func Play(state bool) Property {
	return func(b map[string]interface{}) {
		b["play"] = state
	}
}

func ContextURI(uri string) Property {
	return func(b map[string]interface{}) {
		b["context_uri"] = uri
	}
}

func PropertyURIs(uris []string) Property {
	return func(b map[string]interface{}) {
		b["uris"] = uris
	}
}

func PropertyOffset(v interface{}) Property {
	return func(b map[string]interface{}) {
		b["offset"] = v
	}
}

func PositionMs(num int) Property {
	return func(b map[string]interface{}) {
		b["position_ms"] = num
	}
}

func Name(name string) Property {
	return func(b map[string]interface{}) {
		b["name"] = name
	}
}

func Collaborative(state bool) Property {
	return func(b map[string]interface{}) {
		b["collaborative"] = state
	}
}

func Description(description string) Property {
	return func(b map[string]interface{}) {
		b["description"] = description
	}
}

func RangeStart(num int) Property {
	return func(b map[string]interface{}) {
		b["range_start"] = num
	}
}

func InsertBefore(num int) Property {
	return func(b map[string]interface{}) {
		b["insert_before"] = num
	}
}

func RangeLength(num int) Property {
	return func(b map[string]interface{}) {
		b["range_length"] = num
	}
}

func SnapshotId(id string) Property {
	return func(b map[string]interface{}) {
		b["snapshot_id"] = id
	}
}

func PropertyPosition(num int) Property {
	return func(b map[string]interface{}) {
		b["position"] = num
	}
}

func Tracks(v []interface{}) Property {
	return func(b map[string]interface{}) {
		b["tracks"] = v
	}
}

func createBodyFromProperties(properties []Property) ([]byte, error) {
	body := map[string]interface{}{}
	for _, p := range properties {
		p(body)
	}

	return json.Marshal(body)
}
