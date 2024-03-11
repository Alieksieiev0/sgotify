package api

import "encoding/json"

type Property func(b map[string]interface{})

// A JSON array containing the ID of the device on which playback should be started/transferred.
//
// For example:{device_ids:["74ASZWbe4lXaubB36ztrGX"]}
//
// Note: Although an array is accepted, only a single device_id is currently supported.
// Supplying more than one will return 400 Bad Request
func DeviceIds(ids []string) Property {
	return func(b map[string]interface{}) {
		b["device_ids"] = ids
	}
}

// true: ensure playback happens on new device.
//
// false or not provided: keep the current playback state.
func Play(state bool) Property {
	return func(b map[string]interface{}) {
		b["play"] = state
	}
}

// Spotify URI of the context to play. Valid contexts are albums, artists & playlists.
//
// {context_uri:"spotify:album:1Je1IMUlBXcx1Fz0WE7oPT"}
func ContextURI(uri string) Property {
	return func(b map[string]interface{}) {
		b["context_uri"] = uri
	}
}

// A JSON array of the Spotify URIs.
//
// For example: {"uris": ["spotify:track:4iV5W9uYEdYUVa79Axb7Rh", "spotify:track:1301WleyT98MSxVHPZCA6M"]}
func PropertyURIs(uris []string) Property {
	return func(b map[string]interface{}) {
		b["uris"] = uris
	}
}

// Indicates from where in the context playback should start.
// Only available when context_uri corresponds to an album or
// playlist object "position" is zero based and can’t be negative.
// Example: "offset": {"position": 5} "uri" is a string representing the uri of the item to start at.
// Example: "offset": {"uri": "spotify:track:1301WleyT98MSxVHPZCA6M"}
func PropertyOffset(v interface{}) Property {
	return func(b map[string]interface{}) {
		b["offset"] = v
	}
}

// Position of the Audio Recording in Ms.
func PositionMs(num int) Property {
	return func(b map[string]interface{}) {
		b["position_ms"] = num
	}
}

// The name for the playlist, for example "My New Playlist Title"
func Name(name string) Property {
	return func(b map[string]interface{}) {
		b["name"] = name
	}
}

// If true the playlist will be public, if false it will be private.
func Public(state bool) Property {
	return func(b map[string]interface{}) {
		b["public"] = state
	}
}

// If true, the playlist will become collaborative
// and other users will be able to modify the playlist in their Spotify client.
// Note: You can only set collaborative to true on non-public playlists.
func Collaborative(state bool) Property {
	return func(b map[string]interface{}) {
		b["collaborative"] = state
	}
}

// Value for playlist description as displayed in Spotify Clients and in the Web API.
func Description(description string) Property {
	return func(b map[string]interface{}) {
		b["description"] = description
	}
}

// The position of the first item to be reordered.
func RangeStart(num int) Property {
	return func(b map[string]interface{}) {
		b["range_start"] = num
	}
}

// The position where the items should be inserted.
//
// To reorder the items to the end of the playlist, simply set insert_before to the position after the last item.
//
// Examples:
//
// To reorder the first item to the last position in a playlist with 10 items,
// set range_start to 0, and insert_before to 10.
//
// To reorder the last item in a playlist with 10 items to the start of the playlist,
// set range_start to 9, and insert_before to 0.
func InsertBefore(num int) Property {
	return func(b map[string]interface{}) {
		b["insert_before"] = num
	}
}

// The amount of items to be reordered. Defaults to 1 if not set.
//
// The range of items to be reordered begins from the range_start position,
// and includes the range_length subsequent items.
//
// Example:
//
// To move the items at index 9-10 to the start of the playlist,
// range_start is set to 9, and range_length is set to 2.
func RangeLength(num int) Property {
	return func(b map[string]interface{}) {
		b["range_length"] = num
	}
}

// The playlist's snapshot ID against which you want to make the changes.
func SnapshotId(id string) Property {
	return func(b map[string]interface{}) {
		b["snapshot_id"] = id
	}
}

// The position to insert the items, a zero-based index.
// For example, to insert the items in the first position: position=0 ;
// to insert the items in the third position: position=2.
// If omitted, the items will be appended to the playlist.
// Items are added in the order they appear in the uris array.
// For example: {"uris": ["spotify:track:4iV5W9uYEdYUVa79Axb7Rh","spotify:track:1301WleyT98MSxVHPZCA6M"], "position": 3}
func PropertyPosition(num int) Property {
	return func(b map[string]interface{}) {
		b["position"] = num
	}
}

// An array of objects containing Spotify URIs of the tracks or episodes to remove.
// For example: { "tracks": [{ "uri": "spotify:track:4iV5W9uYEdYUVa79Axb7Rh" },{ "uri": "spotify:track:1301WleyT98MSxVHPZCA6M" }] }.
// A maximum of 100 objects can be sent at once.
func Tracks(v []interface{}) Property {
	return func(b map[string]interface{}) {
		b["tracks"] = v
	}
}

// createBodyFromProperties puts all properties into the map and marshals it into the json.
func createBodyFromProperties(properties []Property) ([]byte, error) {
	body := map[string]interface{}{}
	for _, p := range properties {
		p(body)
	}

	return json.Marshal(body)
}
