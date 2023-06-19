package dtos

type GetShortLinkDto struct {
	ShortLink string `uri:"short_link" binding:"required" json:"short_Link"`
}
