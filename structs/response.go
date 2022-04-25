package structs

type TagJson struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	NumberOfRefernce int    `json:"numberOfReference"`
}

type TagsJson struct {
	Tags []TagJson `json:"tags"`
}
