package repository

import (
	"io"

	"github.com/fairytale5571/awesomeProject1/pkg/yt"
	ytl "github.com/kkdai/youtube/v2"
)

type Repository struct {
	yt *yt.Youtube
}

func NewRepository(yt *yt.Youtube) (*Repository, error) {
	return &Repository{
		yt: yt,
	}, nil
}

func (r *Repository) GetVideo(url string) (*ytl.Video, error) {
	return r.yt.GetVideoInfo(url)
}

func (r *Repository) DownloadVideo(video *ytl.Video, format *ytl.Format) (io.ReadCloser, int64, error) {
	return r.yt.DownloadVideo(video, format)
}
