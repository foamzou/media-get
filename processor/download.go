package processor

import (
	"fmt"
	"path"

	"github.com/foamzou/audio-get/ffmpeg"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

func (p *Processor) download(mediaMeta *meta.MediaMeta, audios []*meta.Audio) error {
	for i, audio := range audios {
		fmt.Printf("-----------------\n start fetch %d, url: %s\n", i, audio.Resource.Url)
		tempOutputPath, targetOutPath := p.getOutputPaths(mediaMeta.Title)

		err := utils.Wget(audio.Resource.Url, tempOutputPath, audio.Resource.Headers)
		if err != nil {
			continue
		}

		fmt.Printf("start convert %d, download to : %s\n", i, tempOutputPath)
		err = ffmpeg.ConvertToMp3(tempOutputPath, targetOutPath, adjustFileMeta(tempOutputPath, audio))
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (p *Processor) getOutputPaths(fileTitle string) (tempFilePath, targetOutPath string) {
	if p.Opts.Out != "" {
		tempFilePath = p.Opts.Out
	} else {
		// Default: current dir
		dir, err := utils.GetCurrentDir()
		if err != nil {
			panic(err)
		}
		p.Opts.IsDir = true
		tempFilePath = dir
	}

	if p.Opts.IsDir {
		tempFilePath = path.Join(tempFilePath, fileTitle)
		targetOutPath = utils.ModifyFileExt(tempFilePath, "mp3")
	} else {
		targetOutPath = tempFilePath
		tempFilePath = path.Join(path.Dir(tempFilePath), fileTitle)
	}

	return
}

func adjustFileMeta(filePath string, audio *meta.Audio) *ffmpeg.MetaTag {
	metaTag := &ffmpeg.MetaTag{
		Album:  audio.Album,
		Title:  audio.Title,
		Artist: audio.Artist,
	}
	mediaFormat, err := ffmpeg.GetMediaFormat(filePath)

	if err != nil {
		fmt.Println(err)
		return metaTag
	}

	// Priority to use origin meta
	if mediaFormat.Format.Tags.Album != "" {
		metaTag.Album = mediaFormat.Format.Tags.Album
	}
	if mediaFormat.Format.Tags.Artist != "" {
		metaTag.Artist = mediaFormat.Format.Tags.Artist
	}
	if mediaFormat.Format.Tags.Title != "" {
		metaTag.Title = mediaFormat.Format.Tags.Title
	}

	return metaTag
}
