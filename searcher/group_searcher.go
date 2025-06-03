package searcher

import (
	"context"
	"fmt"
	"yamdc/model"
	"yamdc/number"

	"dario.cat/mergo"
	"github.com/xxxsen/common/logutil"
	"go.uber.org/zap"
)

type group struct {
	ss []ISearcher
}

func NewGroup(ss []ISearcher) ISearcher {
	return &group{ss: ss}
}
func (g *group) Name() string {
	return "group"
}

func (g *group) Search(ctx context.Context, number *number.Number) (*model.MovieMeta, bool, error) {
	return performGroupSearch(ctx, number, g.ss)
}

func (g *group) Check(ctx context.Context) error {
	return fmt.Errorf("unable to perform check on group searcher")
}

func performGroupSearch(ctx context.Context, number *number.Number, ss []ISearcher) (*model.MovieMeta, bool, error) {
	var lastErr error
	var meta_final model.MovieMeta
	for i, s := range ss {
		logutil.GetLogger(ctx).Debug("search number", zap.String("plugin", s.Name()))
		meta, found, err := s.Search(ctx, number)
		if err != nil {
			lastErr = err
			continue
		}
		if !found {
			continue
		}
		if i == 0 {
			meta_final = *meta
		} else if err := mergo.Merge(&meta_final, *meta); err != nil {
			lastErr = err
			continue
		}
		if meta_final.Number == "" || meta_final.Title == "" || meta_final.Plot == "" || len(meta_final.Actors) == 0 {
			continue
		}
		return meta, true, nil
	}
	if meta_final.Number != "" && meta_final.Title != "" {
		return &meta_final, true, nil
	}
	if lastErr != nil {
		return nil, false, lastErr
	}
	return nil, false, nil
}
