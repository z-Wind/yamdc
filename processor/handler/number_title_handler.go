package handler

import (
	"context"
	"strings"
	"yamdc/model"
	"yamdc/number"
)

type numberTitleHandler struct {
}

func (h *numberTitleHandler) Handle(ctx context.Context, fc *model.FileContext) error {
	title := number.GetCleanID(fc.Meta.Title)
	num := number.GetCleanID(fc.Number.GetNumberID())
	if strings.Contains(title, num) {
		return nil
	}
	fc.Meta.Title = fc.Number.GetNumberID() + " " + fc.Meta.Title
	if len(fc.Meta.TitleTranslated) > 0 {
		fc.Meta.TitleTranslated = fc.Number.GetNumberID() + " " + fc.Meta.TitleTranslated
	}

	return nil
}

func init() {
	Register(HNumberTitle, HandlerToCreator(&numberTitleHandler{}))
}
