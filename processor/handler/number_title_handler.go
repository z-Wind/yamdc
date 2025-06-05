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
	num := number.GetCleanID(fc.Number.GetNumberID())

	title := number.GetCleanID(fc.Meta.Title)
	title = strings.ReplaceAll(title, num, "")
	fc.Meta.Title = fc.Number.GetNumberID() + " " + title

	if len(fc.Meta.TitleTranslated) > 0 {
		titleTranslated := number.GetCleanID(fc.Meta.TitleTranslated)
		titleTranslated = strings.ReplaceAll(titleTranslated, num, "")
		fc.Meta.TitleTranslated = fc.Number.GetNumberID() + " " + titleTranslated
	}

	return nil
}

func init() {
	Register(HNumberTitle, HandlerToCreator(&numberTitleHandler{}))
}
