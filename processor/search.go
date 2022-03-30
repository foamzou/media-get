package processor

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	ccgo "github.com/ApesPlan/OpenCC-go"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/logger"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

func (p *Processor) Search() error {
	p.rebuildKeyword()
	return p.SearchSong()
}

func (p *Processor) rebuildKeyword() {
	if p.Opts.Search.Keyword == "" {
		p.Opts.Search.Keyword = strings.Join([]string{
			p.Opts.Search.SongName,
			p.Opts.Search.Artist,
		}, " ")

		logger.Debug("rebuild the keyword to " + p.Opts.Search.Keyword)
	}
}

func (p *Processor) SearchSong() error {
	var searchItems []*meta.SearchSongItem
	processors := p.getProcessors()
	var mutex sync.Mutex

	logger.Debug("Will search in ", p.Opts.Search.SourcesWillBeSearch)

	wg := sync.WaitGroup{}
	for _, processor := range processors {
		if !utils.InArray(p.Opts.Search.SourcesWillBeSearch, processor.GetSourceName()) {
			continue
		}
		wg.Add(1)
		processor := processor
		go func() {
			defer func() {
				wg.Done()
			}()
			items, err := processor.SearchSong()
			if err != nil {
				logger.Debug(err)
				return
			}
			const MaxPerSource = 3
			if len(items) > MaxPerSource {
				items = items[:MaxPerSource]
			}

			mutex.Lock()
			searchItems = append(searchItems, items...)
			mutex.Unlock()
		}()
	}
	wg.Wait()

	p.sortTheResult(searchItems)
	p.outputSearchResult(searchItems)
	return nil
}

func (p *Processor) sortTheResult(searchItems []*meta.SearchSongItem) {
	sort.SliceStable(searchItems, func(i, j int) bool {
		searchOption := p.Opts.Search
		searchItems[i].Score = p.calculateTheScore(searchOption, *searchItems[i])
		searchItems[j].Score = p.calculateTheScore(searchOption, *searchItems[j])

		return searchItems[i].Score > searchItems[j].Score
	})
}

func (p *Processor) calculateTheScore(searchOption args.Search, searchItem meta.SearchSongItem) float64 {
	// alignment text
	t2s, _ := ccgo.New("t2s")
	searchOption.Keyword = convertTC2SC(t2s, strings.ToLower(searchOption.Keyword))
	searchOption.SongName = convertTC2SC(t2s, strings.ToLower(searchOption.SongName))
	searchOption.Artist = convertTC2SC(t2s, strings.ToLower(searchOption.Artist))
	searchOption.Album = convertTC2SC(t2s, strings.ToLower(searchOption.Album))

	searchItem.Name = convertTC2SC(t2s, strings.ToLower(searchItem.Name))
	searchItem.Artist = convertTC2SC(t2s, strings.ToLower(searchItem.Artist))
	searchItem.Album = convertTC2SC(t2s, strings.ToLower(searchItem.Album))

	// cal the baseWeightCoefficient
	bwc := 1.0
	extraScoreConfig := []struct {
		artist               string
		source               string
		weightCoefficientBuf float64
	}{
		{artist: "周杰伦", source: consts.SourceNameMigu, weightCoefficientBuf: 0.25},
		{artist: searchItem.Artist, source: consts.SourceNameMigu, weightCoefficientBuf: 0.1}, // migu is a good platform
	}

	for _, config := range extraScoreConfig {
		if strings.Contains(searchItem.Artist, config.artist) && strings.Contains(searchOption.Keyword, config.artist) && config.source == searchItem.Source {
			bwc += config.weightCoefficientBuf
		}
	}

	// song name is very important
	if !strings.Contains(searchOption.Keyword, searchItem.Name) {
		bwc -= 1
	}

	provider := p.getProcessorBySourceName(searchItem.Source)
	if provider != nil {
		// for migu: remove the text within brackets since it's interfere with keyword
		if provider.GetSourceName() == consts.SourceNameMigu {
			searchItem.Name = utils.RemoveBracketsFromString(searchItem.Name)
		}

		// The music platform should have correct artist name
		if provider.IsMusicPlatform() {
			// the item should not be target if not match the user's input
			if searchOption.Artist != "" && !strings.Contains(searchItem.Artist, searchOption.Artist) {
				bwc -= 0.8
			}
		} else {
			// should balance the album and artist score for non-music platform since the platform may have target
			if searchOption.Artist != "" {
				searchItem.Artist = searchOption.Artist
			}

			if searchOption.Album != "" {
				searchItem.Album = searchOption.Album
			}

			// MV on the platform is close to the target
			if strings.Contains(searchItem.Name, "mv") {
				bwc += 1
			}
		}
	}

	if searchItem.CannotDownload {
		bwc -= 0.5
	}

	// cal the score
	keywordScore := utils.SimilarText(fmt.Sprintf("%s %s", searchItem.Name, searchItem.Artist), searchOption.Keyword) * (bwc + 0.1)
	nameScore := utils.SimilarText(searchItem.Name, searchOption.SongName) * (bwc + 0.2)
	albumScore := utils.SimilarText(searchItem.Album, searchOption.Album) * (bwc + 0.1)
	artistScore := utils.SimilarText(searchItem.Artist, searchOption.Artist) * bwc

	return keywordScore + nameScore + artistScore + albumScore
}

func (p *Processor) outputSearchResult(searchItems []*meta.SearchSongItem) {
	// json output
	if p.Opts.InfoFormat != consts.InfoFormatPlain {
		jsonByte, err := json.MarshalIndent(searchItems, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonByte))
		return
	}

	fmt.Println("search result:")
	for _, item := range searchItems {
		fmt.Printf("name: %s, album: %s, artist: %s, url: %s, duration: %d, source: %s, score: %f\n",
			item.Name, item.Album, item.Artist, item.Url, item.Duration, item.Source, item.Score)
	}
}

func convertTC2SC(occ *ccgo.OpenCC, in string) (out string) {
	out = in
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	out, _ = occ.Convert(in)
	if out == "" {
		return in
	}
	return out
}
