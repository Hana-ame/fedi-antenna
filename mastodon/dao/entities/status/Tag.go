package status

type Tag struct {
// Type: String
// Description: The value of the hashtag after the # sign.
Name string `json:"name"`
// Type: String (URL)
// Description: A link to the hashtag on the instance.
Url string `json:"url"`
}