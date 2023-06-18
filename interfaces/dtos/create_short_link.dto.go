package dtos

type CreateShortLinkDto struct {
	Link string `form:"link" binding:"required" json:"link"`
}
