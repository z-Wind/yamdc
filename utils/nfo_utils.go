package utils

import (
	"time"
	"yamdc/model"
	"yamdc/nfo"
)

func ConvertMetaToMovieNFO(m *model.MovieMeta) (*nfo.Movie, error) {
	mv := &nfo.Movie{
		ID:            m.Number,
		Plot:          m.Plot,
		Dateadded:     FormatTimeToDate(time.Now().UnixMilli()),
		Title:         m.Title,
		OriginalTitle: m.Title,
		SortTitle:     m.Title,
		Set:           m.Series,
		Rating:        0,
		Release:       FormatTimeToDate(m.ReleaseDate),
		ReleaseDate:   FormatTimeToDate(m.ReleaseDate),
		Premiered:     FormatTimeToDate(m.ReleaseDate),
		Runtime:       uint64(m.Duration) / 60, //分钟数
		Year:          time.UnixMilli(m.ReleaseDate).Year(),
		Tags:          m.Genres,
		Genres:        m.Genres,
		Studio:        m.Studio,
		Maker:         m.Studio,
		Art:           nfo.Art{},
		Mpaa:          "JP-18+",
		Director:      "",
		Label:         m.Label,
		Thumb:         "",
		ScrapeInfo: nfo.ScrapeInfo{
			Source: m.ExtInfo.ScrapeInfo.Source,
			Date:   time.UnixMilli(m.ExtInfo.ScrapeInfo.DateTs).Format(time.DateOnly),
		},
	}
	if len(m.TitleTranslated) > 0 {
		mv.Title = m.TitleTranslated
	}
	if len(m.PlotTranslated) > 0 {
		mv.Plot = m.PlotTranslated + " [原文:" + mv.Plot + "]"
	}
	if m.Poster != nil {
		mv.Art.Poster = m.Poster.Name
		mv.Poster = m.Poster.Name
		//
		mv.Art.Fanart = append(mv.Art.Fanart, m.Poster.Name)
	}
	if m.Cover != nil {
		mv.Cover = m.Cover.Name
		mv.Fanart = m.Cover.Name
		//
		mv.Art.Fanart = append(mv.Art.Fanart, m.Cover.Name)
	}
	for _, act := range m.Actors {
		mv.Actors = append(mv.Actors, nfo.Actor{
			Name: act,
		})
	}
	for _, image := range m.SampleImages {
		mv.Art.Fanart = append(mv.Art.Fanart, image.Name)
	}
	return mv, nil
}
