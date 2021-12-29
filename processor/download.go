package processor

import (
	"fmt"
	"os"
	"path"

	"github.com/foamzou/audio-get/ffmpeg"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

func (p *Processor) download(mediaMeta *meta.MediaMeta) error {
	fmt.Printf("-----------------\n start fetch url: %s\n", mediaMeta.Audio.Url)
	tempOutputPath, targetOutPath := p.getOutputPaths(mediaMeta.Title)

	err := utils.Wget(mediaMeta.Audio.Url, tempOutputPath, mediaMeta.Headers)
	if err != nil {
		return err
	}

	fmt.Printf("start convert, download to : %s\n", tempOutputPath)
	err = ffmpeg.ConvertToMp3(tempOutputPath, targetOutPath, adjustFileMeta(tempOutputPath, mediaMeta))
	_ = os.Remove(tempOutputPath)
	if err != nil {
		fmt.Println(err)
		return err
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

func adjustFileMeta(filePath string, mediaMeta *meta.MediaMeta) *ffmpeg.MetaTag {
	metaTag := &ffmpeg.MetaTag{
		Album:  mediaMeta.Album,
		Title:  mediaMeta.Title,
		Artist: mediaMeta.Artist,
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
