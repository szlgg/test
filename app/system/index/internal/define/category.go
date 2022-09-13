package define

type CategoryList struct {
	Id          uint   `form:"id"`
	Name        string `form:"name"`
	ContentType string `form:"content_type"`
}
