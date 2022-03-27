package youtube

type SearchSongResponse struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MainAppWebResponseContext struct {
			LoggedOut bool `json:"loggedOut"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			YtConfigData struct {
				VisitorData           string `json:"visitorData"`
				RootVisualElementType int    `json:"rootVisualElementType"`
			} `json:"ytConfigData"`
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	EstimatedResults string `json:"estimatedResults"`
	Contents         struct {
		TwoColumnSearchResultsRenderer struct {
			PrimaryContents struct {
				SectionListRenderer struct {
					Contents []struct {
						ItemSectionRenderer struct {
							Contents []struct {
								VideoRenderer struct {
									VideoId   string `json:"videoId"`
									Thumbnail struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnail"`
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
									} `json:"title"`
									LongBylineText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId         string `json:"browseId"`
													CanonicalBaseUrl string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"longBylineText"`
									PublishedTimeText struct {
										SimpleText string `json:"simpleText"`
									} `json:"publishedTimeText"`
									LengthText struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"lengthText"`
									ViewCountText struct {
										SimpleText string `json:"simpleText"`
									} `json:"viewCountText"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										WatchEndpoint struct {
											VideoId                            string `json:"videoId"`
											Params                             string `json:"params"`
											WatchEndpointSupportedOnesieConfig struct {
												Html5PlaybackOnesieConfig struct {
													CommonConfig struct {
														Url string `json:"url"`
													} `json:"commonConfig"`
												} `json:"html5PlaybackOnesieConfig"`
											} `json:"watchEndpointSupportedOnesieConfig"`
										} `json:"watchEndpoint"`
									} `json:"navigationEndpoint"`
									Badges []struct {
										MetadataBadgeRenderer struct {
											Style             string `json:"style"`
											Label             string `json:"label"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"metadataBadgeRenderer"`
									} `json:"badges,omitempty"`
									OwnerBadges []struct {
										MetadataBadgeRenderer struct {
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Style             string `json:"style"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"metadataBadgeRenderer"`
									} `json:"ownerBadges,omitempty"`
									OwnerText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId         string `json:"browseId"`
													CanonicalBaseUrl string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"ownerText"`
									ShortBylineText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId         string `json:"browseId"`
													CanonicalBaseUrl string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"shortBylineText"`
									TrackingParams     string `json:"trackingParams"`
									ShowActionMenu     bool   `json:"showActionMenu"`
									ShortViewCountText struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"shortViewCountText"`
									Menu struct {
										MenuRenderer struct {
											Items []struct {
												MenuServiceItemRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													ServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool `json:"sendPost"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														SignalServiceEndpoint struct {
															Signal  string `json:"signal"`
															Actions []struct {
																ClickTrackingParams  string `json:"clickTrackingParams"`
																AddToPlaylistCommand struct {
																	OpenMiniplayer      bool   `json:"openMiniplayer"`
																	VideoId             string `json:"videoId"`
																	ListType            string `json:"listType"`
																	OnCreateListCommand struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				ApiUrl   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		CreatePlaylistServiceEndpoint struct {
																			VideoIds []string `json:"videoIds"`
																			Params   string   `json:"params"`
																		} `json:"createPlaylistServiceEndpoint"`
																	} `json:"onCreateListCommand"`
																	VideoIds []string `json:"videoIds"`
																} `json:"addToPlaylistCommand"`
															} `json:"actions"`
														} `json:"signalServiceEndpoint"`
													} `json:"serviceEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"menuServiceItemRenderer"`
											} `json:"items"`
											TrackingParams string `json:"trackingParams"`
											Accessibility  struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
										} `json:"menuRenderer"`
									} `json:"menu"`
									ChannelThumbnailSupportedRenderers struct {
										ChannelThumbnailWithLinkRenderer struct {
											Thumbnail struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"thumbnail"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId         string `json:"browseId"`
													CanonicalBaseUrl string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
											Accessibility struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
										} `json:"channelThumbnailWithLinkRenderer"`
									} `json:"channelThumbnailSupportedRenderers"`
									ThumbnailOverlays []struct {
										ThumbnailOverlayTimeStatusRenderer struct {
											Text struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"text"`
											Style string `json:"style"`
										} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
										ThumbnailOverlayToggleButtonRenderer struct {
											IsToggled     bool `json:"isToggled,omitempty"`
											UntoggledIcon struct {
												IconType string `json:"iconType"`
											} `json:"untoggledIcon"`
											ToggledIcon struct {
												IconType string `json:"iconType"`
											} `json:"toggledIcon"`
											UntoggledTooltip         string `json:"untoggledTooltip"`
											ToggledTooltip           string `json:"toggledTooltip"`
											UntoggledServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														ApiUrl   string `json:"apiUrl,omitempty"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												PlaylistEditEndpoint struct {
													PlaylistId string `json:"playlistId"`
													Actions    []struct {
														AddedVideoId string `json:"addedVideoId"`
														Action       string `json:"action"`
													} `json:"actions"`
												} `json:"playlistEditEndpoint,omitempty"`
												SignalServiceEndpoint struct {
													Signal  string `json:"signal"`
													Actions []struct {
														ClickTrackingParams  string `json:"clickTrackingParams"`
														AddToPlaylistCommand struct {
															OpenMiniplayer      bool   `json:"openMiniplayer"`
															VideoId             string `json:"videoId"`
															ListType            string `json:"listType"`
															OnCreateListCommand struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		ApiUrl   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																CreatePlaylistServiceEndpoint struct {
																	VideoIds []string `json:"videoIds"`
																	Params   string   `json:"params"`
																} `json:"createPlaylistServiceEndpoint"`
															} `json:"onCreateListCommand"`
															VideoIds []string `json:"videoIds"`
														} `json:"addToPlaylistCommand"`
													} `json:"actions"`
												} `json:"signalServiceEndpoint,omitempty"`
											} `json:"untoggledServiceEndpoint"`
											ToggledServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														ApiUrl   string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												PlaylistEditEndpoint struct {
													PlaylistId string `json:"playlistId"`
													Actions    []struct {
														Action         string `json:"action"`
														RemovedVideoId string `json:"removedVideoId"`
													} `json:"actions"`
												} `json:"playlistEditEndpoint"`
											} `json:"toggledServiceEndpoint,omitempty"`
											UntoggledAccessibility struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"untoggledAccessibility"`
											ToggledAccessibility struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"toggledAccessibility"`
											TrackingParams string `json:"trackingParams"`
										} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
										ThumbnailOverlayNowPlayingRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
										} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
									} `json:"thumbnailOverlays"`
									RichThumbnail struct {
										MovingThumbnailRenderer struct {
											MovingThumbnailDetails struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
												LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
											} `json:"movingThumbnailDetails"`
											EnableHoveredLogging bool `json:"enableHoveredLogging"`
											EnableOverlay        bool `json:"enableOverlay"`
										} `json:"movingThumbnailRenderer"`
									} `json:"richThumbnail,omitempty"`
									DetailedMetadataSnippets []struct {
										SnippetText struct {
											Runs []struct {
												Text string `json:"text"`
												Bold bool   `json:"bold,omitempty"`
											} `json:"runs"`
										} `json:"snippetText"`
										SnippetHoverText struct {
											Runs []struct {
												Text string `json:"text"`
											} `json:"runs"`
										} `json:"snippetHoverText"`
										MaxOneLine bool `json:"maxOneLine"`
									} `json:"detailedMetadataSnippets,omitempty"`
								} `json:"videoRenderer,omitempty"`
								RadioRenderer struct {
									PlaylistId string `json:"playlistId"`
									Title      struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Thumbnail struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
										SampledThumbnailColor struct {
											Red   int `json:"red"`
											Green int `json:"green"`
											Blue  int `json:"blue"`
										} `json:"sampledThumbnailColor"`
									} `json:"thumbnail"`
									VideoCountText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"videoCountText"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										WatchEndpoint struct {
											VideoId          string `json:"videoId"`
											PlaylistId       string `json:"playlistId"`
											Params           string `json:"params"`
											ContinuePlayback bool   `json:"continuePlayback"`
											LoggingContext   struct {
												VssLoggingContext struct {
													SerializedContextData string `json:"serializedContextData"`
												} `json:"vssLoggingContext"`
											} `json:"loggingContext"`
											WatchEndpointSupportedOnesieConfig struct {
												Html5PlaybackOnesieConfig struct {
													CommonConfig struct {
														Url string `json:"url"`
													} `json:"commonConfig"`
												} `json:"html5PlaybackOnesieConfig"`
											} `json:"watchEndpointSupportedOnesieConfig"`
										} `json:"watchEndpoint"`
									} `json:"navigationEndpoint"`
									ShortBylineText struct {
										SimpleText string `json:"simpleText"`
									} `json:"shortBylineText"`
									TrackingParams string `json:"trackingParams"`
									Videos         []struct {
										ChildVideoRenderer struct {
											Title struct {
												SimpleText string `json:"simpleText"`
											} `json:"title"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoId        string `json:"videoId"`
													PlaylistId     string `json:"playlistId"`
													Params         string `json:"params"`
													LoggingContext struct {
														VssLoggingContext struct {
															SerializedContextData string `json:"serializedContextData"`
														} `json:"vssLoggingContext"`
													} `json:"loggingContext"`
													WatchEndpointSupportedOnesieConfig struct {
														Html5PlaybackOnesieConfig struct {
															CommonConfig struct {
																Url string `json:"url"`
															} `json:"commonConfig"`
														} `json:"html5PlaybackOnesieConfig"`
													} `json:"watchEndpointSupportedOnesieConfig"`
												} `json:"watchEndpoint"`
											} `json:"navigationEndpoint"`
											LengthText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"lengthText"`
											VideoId string `json:"videoId"`
										} `json:"childVideoRenderer"`
									} `json:"videos"`
									ThumbnailText struct {
										Runs []struct {
											Text string `json:"text"`
											Bold bool   `json:"bold,omitempty"`
										} `json:"runs"`
									} `json:"thumbnailText"`
									LongBylineText struct {
										SimpleText string `json:"simpleText"`
									} `json:"longBylineText"`
									ThumbnailOverlays []struct {
										ThumbnailOverlayBottomPanelRenderer struct {
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
										} `json:"thumbnailOverlayBottomPanelRenderer,omitempty"`
										ThumbnailOverlayHoverTextRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
										} `json:"thumbnailOverlayHoverTextRenderer,omitempty"`
										ThumbnailOverlayNowPlayingRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
										} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
									} `json:"thumbnailOverlays"`
									VideoCountShortText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"videoCountShortText"`
								} `json:"radioRenderer,omitempty"`
								ShelfRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Content struct {
										VerticalListRenderer struct {
											Items []struct {
												VideoRenderer struct {
													VideoId   string `json:"videoId"`
													Thumbnail struct {
														Thumbnails []struct {
															Url    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													Title struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
													} `json:"title"`
													LongBylineText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"longBylineText"`
													PublishedTimeText struct {
														SimpleText string `json:"simpleText"`
													} `json:"publishedTimeText"`
													LengthText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"lengthText"`
													ViewCountText struct {
														SimpleText string `json:"simpleText"`
													} `json:"viewCountText"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoId                            string `json:"videoId"`
															WatchEndpointSupportedOnesieConfig struct {
																Html5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		Url string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													Badges []struct {
														MetadataBadgeRenderer struct {
															Style             string `json:"style"`
															Label             string `json:"label"`
															TrackingParams    string `json:"trackingParams"`
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData,omitempty"`
														} `json:"metadataBadgeRenderer"`
													} `json:"badges,omitempty"`
													OwnerText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"ownerText"`
													ShortBylineText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"shortBylineText"`
													TrackingParams     string `json:"trackingParams"`
													ShowActionMenu     bool   `json:"showActionMenu"`
													ShortViewCountText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"shortViewCountText"`
													Menu struct {
														MenuRenderer struct {
															Items []struct {
																MenuServiceItemRenderer struct {
																	Text struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																	Icon struct {
																		IconType string `json:"iconType"`
																	} `json:"icon"`
																	ServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool `json:"sendPost"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		SignalServiceEndpoint struct {
																			Signal  string `json:"signal"`
																			Actions []struct {
																				ClickTrackingParams  string `json:"clickTrackingParams"`
																				AddToPlaylistCommand struct {
																					OpenMiniplayer      bool   `json:"openMiniplayer"`
																					VideoId             string `json:"videoId"`
																					ListType            string `json:"listType"`
																					OnCreateListCommand struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								SendPost bool   `json:"sendPost"`
																								ApiUrl   string `json:"apiUrl"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						CreatePlaylistServiceEndpoint struct {
																							VideoIds []string `json:"videoIds"`
																							Params   string   `json:"params"`
																						} `json:"createPlaylistServiceEndpoint"`
																					} `json:"onCreateListCommand"`
																					VideoIds []string `json:"videoIds"`
																				} `json:"addToPlaylistCommand"`
																			} `json:"actions"`
																		} `json:"signalServiceEndpoint"`
																	} `json:"serviceEndpoint"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"menuServiceItemRenderer"`
															} `json:"items"`
															TrackingParams string `json:"trackingParams"`
															Accessibility  struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"menuRenderer"`
													} `json:"menu"`
													ChannelThumbnailSupportedRenderers struct {
														ChannelThumbnailWithLinkRenderer struct {
															Thumbnail struct {
																Thumbnails []struct {
																	Url    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
															} `json:"thumbnail"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
															Accessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"channelThumbnailWithLinkRenderer"`
													} `json:"channelThumbnailSupportedRenderers"`
													ThumbnailOverlays []struct {
														ThumbnailOverlayTimeStatusRenderer struct {
															Text struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															Style string `json:"style"`
														} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
														ThumbnailOverlayToggleButtonRenderer struct {
															IsToggled     bool `json:"isToggled,omitempty"`
															UntoggledIcon struct {
																IconType string `json:"iconType"`
															} `json:"untoggledIcon"`
															ToggledIcon struct {
																IconType string `json:"iconType"`
															} `json:"toggledIcon"`
															UntoggledTooltip         string `json:"untoggledTooltip"`
															ToggledTooltip           string `json:"toggledTooltip"`
															UntoggledServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		ApiUrl   string `json:"apiUrl,omitempty"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																PlaylistEditEndpoint struct {
																	PlaylistId string `json:"playlistId"`
																	Actions    []struct {
																		AddedVideoId string `json:"addedVideoId"`
																		Action       string `json:"action"`
																	} `json:"actions"`
																} `json:"playlistEditEndpoint,omitempty"`
																SignalServiceEndpoint struct {
																	Signal  string `json:"signal"`
																	Actions []struct {
																		ClickTrackingParams  string `json:"clickTrackingParams"`
																		AddToPlaylistCommand struct {
																			OpenMiniplayer      bool   `json:"openMiniplayer"`
																			VideoId             string `json:"videoId"`
																			ListType            string `json:"listType"`
																			OnCreateListCommand struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						ApiUrl   string `json:"apiUrl"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				CreatePlaylistServiceEndpoint struct {
																					VideoIds []string `json:"videoIds"`
																					Params   string   `json:"params"`
																				} `json:"createPlaylistServiceEndpoint"`
																			} `json:"onCreateListCommand"`
																			VideoIds []string `json:"videoIds"`
																		} `json:"addToPlaylistCommand"`
																	} `json:"actions"`
																} `json:"signalServiceEndpoint,omitempty"`
															} `json:"untoggledServiceEndpoint"`
															ToggledServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		ApiUrl   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																PlaylistEditEndpoint struct {
																	PlaylistId string `json:"playlistId"`
																	Actions    []struct {
																		Action         string `json:"action"`
																		RemovedVideoId string `json:"removedVideoId"`
																	} `json:"actions"`
																} `json:"playlistEditEndpoint"`
															} `json:"toggledServiceEndpoint,omitempty"`
															UntoggledAccessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"untoggledAccessibility"`
															ToggledAccessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"toggledAccessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
														ThumbnailOverlayNowPlayingRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
														} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
													} `json:"thumbnailOverlays"`
													RichThumbnail struct {
														MovingThumbnailRenderer struct {
															MovingThumbnailDetails struct {
																Thumbnails []struct {
																	Url    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
																LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
															} `json:"movingThumbnailDetails"`
															EnableHoveredLogging bool `json:"enableHoveredLogging"`
															EnableOverlay        bool `json:"enableOverlay"`
														} `json:"movingThumbnailRenderer"`
													} `json:"richThumbnail"`
													DetailedMetadataSnippets []struct {
														SnippetText struct {
															Runs []struct {
																Text string `json:"text"`
																Bold bool   `json:"bold,omitempty"`
															} `json:"runs"`
														} `json:"snippetText"`
														SnippetHoverText struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"snippetHoverText"`
														MaxOneLine bool `json:"maxOneLine"`
													} `json:"detailedMetadataSnippets,omitempty"`
													OwnerBadges []struct {
														MetadataBadgeRenderer struct {
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															Style             string `json:"style"`
															Tooltip           string `json:"tooltip"`
															TrackingParams    string `json:"trackingParams"`
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"metadataBadgeRenderer"`
													} `json:"ownerBadges,omitempty"`
												} `json:"videoRenderer"`
											} `json:"items"`
											CollapsedItemCount       int `json:"collapsedItemCount"`
											CollapsedStateButtonText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
											} `json:"collapsedStateButtonText"`
											TrackingParams string `json:"trackingParams"`
										} `json:"verticalListRenderer"`
									} `json:"content"`
									TrackingParams string `json:"trackingParams"`
								} `json:"shelfRenderer,omitempty"`
								HorizontalCardListRenderer struct {
									Cards []struct {
										SearchRefinementCardRenderer struct {
											Thumbnail struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"thumbnail"`
											Query struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"query"`
											SearchEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SearchEndpoint struct {
													Query  string `json:"query"`
													Params string `json:"params"`
												} `json:"searchEndpoint"`
											} `json:"searchEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"searchRefinementCardRenderer"`
									} `json:"cards"`
									TrackingParams string `json:"trackingParams"`
									Header         struct {
										RichListHeaderRenderer struct {
											Title struct {
												SimpleText string `json:"simpleText"`
											} `json:"title"`
											TrackingParams string `json:"trackingParams"`
										} `json:"richListHeaderRenderer"`
									} `json:"header"`
									Style struct {
										Type string `json:"type"`
									} `json:"style"`
									PreviousButton struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Icon       struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"previousButton"`
									NextButton struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Icon       struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"nextButton"`
								} `json:"horizontalCardListRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
						} `json:"itemSectionRenderer,omitempty"`
						ContinuationItemRenderer struct {
							Trigger              string `json:"trigger"`
							ContinuationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										ApiUrl   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								ContinuationCommand struct {
									Token   string `json:"token"`
									Request string `json:"request"`
								} `json:"continuationCommand"`
							} `json:"continuationEndpoint"`
						} `json:"continuationItemRenderer,omitempty"`
					} `json:"contents"`
					TrackingParams string `json:"trackingParams"`
					SubMenu        struct {
						SearchSubMenuRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Groups []struct {
								SearchFilterGroupRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Filters []struct {
										SearchFilterRenderer struct {
											Label struct {
												SimpleText string `json:"simpleText"`
											} `json:"label"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SearchEndpoint struct {
													Query  string `json:"query"`
													Params string `json:"params"`
												} `json:"searchEndpoint"`
											} `json:"navigationEndpoint,omitempty"`
											Tooltip        string `json:"tooltip"`
											TrackingParams string `json:"trackingParams"`
											Status         string `json:"status,omitempty"`
										} `json:"searchFilterRenderer"`
									} `json:"filters"`
									TrackingParams string `json:"trackingParams"`
								} `json:"searchFilterGroupRenderer"`
							} `json:"groups"`
							TrackingParams string `json:"trackingParams"`
							Button         struct {
								ToggleButtonRenderer struct {
									Style struct {
										StyleType string `json:"styleType"`
									} `json:"style"`
									IsToggled   bool `json:"isToggled"`
									IsDisabled  bool `json:"isDisabled"`
									DefaultIcon struct {
										IconType string `json:"iconType"`
									} `json:"defaultIcon"`
									DefaultText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"defaultText"`
									Accessibility struct {
										Label string `json:"label"`
									} `json:"accessibility"`
									TrackingParams string `json:"trackingParams"`
									DefaultTooltip string `json:"defaultTooltip"`
									ToggledTooltip string `json:"toggledTooltip"`
									ToggledStyle   struct {
										StyleType string `json:"styleType"`
									} `json:"toggledStyle"`
									AccessibilityData struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibilityData"`
								} `json:"toggleButtonRenderer"`
							} `json:"button"`
						} `json:"searchSubMenuRenderer"`
					} `json:"subMenu"`
					HideBottomSeparator bool   `json:"hideBottomSeparator"`
					TargetId            string `json:"targetId"`
				} `json:"sectionListRenderer"`
			} `json:"primaryContents"`
		} `json:"twoColumnSearchResultsRenderer"`
	} `json:"contents"`
	TrackingParams string `json:"trackingParams"`
	Topbar         struct {
		DesktopTopbarRenderer struct {
			Logo struct {
				TopbarLogoRenderer struct {
					IconImage struct {
						IconType string `json:"iconType"`
					} `json:"iconImage"`
					TooltipText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"tooltipText"`
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								ApiUrl      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseId string `json:"browseId"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					TrackingParams    string `json:"trackingParams"`
					OverrideEntityKey string `json:"overrideEntityKey"`
				} `json:"topbarLogoRenderer"`
			} `json:"logo"`
			Searchbox struct {
				FusionSearchboxRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					PlaceholderText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"placeholderText"`
					Config struct {
						WebSearchboxConfig struct {
							RequestLanguage     string `json:"requestLanguage"`
							RequestDomain       string `json:"requestDomain"`
							HasOnscreenKeyboard bool   `json:"hasOnscreenKeyboard"`
							FocusSearchbox      bool   `json:"focusSearchbox"`
						} `json:"webSearchboxConfig"`
					} `json:"config"`
					TrackingParams string `json:"trackingParams"`
					SearchEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SearchEndpoint struct {
							Query string `json:"query"`
						} `json:"searchEndpoint"`
					} `json:"searchEndpoint"`
					ClearButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Icon       struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"clearButton"`
				} `json:"fusionSearchboxRenderer"`
			} `json:"searchbox"`
			TrackingParams string `json:"trackingParams"`
			CountryCode    string `json:"countryCode"`
			TopbarButtons  []struct {
				TopbarMenuButtonRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					MenuRenderer struct {
						MultiPageMenuRenderer struct {
							Sections []struct {
								MultiPageMenuSectionRenderer struct {
									Items []struct {
										CompactLinkRenderer struct {
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"title"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												UrlEndpoint struct {
													Url    string `json:"url"`
													Target string `json:"target"`
												} `json:"urlEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"compactLinkRenderer"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
								} `json:"multiPageMenuSectionRenderer"`
							} `json:"sections"`
							TrackingParams string `json:"trackingParams"`
							Style          string `json:"style"`
						} `json:"multiPageMenuRenderer"`
					} `json:"menuRenderer,omitempty"`
					TrackingParams string `json:"trackingParams"`
					Accessibility  struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					Tooltip     string `json:"tooltip"`
					Style       string `json:"style"`
					TargetId    string `json:"targetId,omitempty"`
					MenuRequest struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										MultiPageMenuRenderer struct {
											TrackingParams     string `json:"trackingParams"`
											Style              string `json:"style"`
											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
										} `json:"multiPageMenuRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
									BeReused  bool   `json:"beReused"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"menuRequest,omitempty"`
				} `json:"topbarMenuButtonRenderer,omitempty"`
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignInEndpoint struct {
							IdamTag string `json:"idamTag"`
						} `json:"signInEndpoint"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					TargetId       string `json:"targetId"`
				} `json:"buttonRenderer,omitempty"`
			} `json:"topbarButtons"`
			HotkeyDialog struct {
				HotkeyDialogRenderer struct {
					Title struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"title"`
					Sections []struct {
						HotkeyDialogSectionRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Options []struct {
								HotkeyDialogSectionOptionRenderer struct {
									Label struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"label"`
									Hotkey                   string `json:"hotkey"`
									HotkeyAccessibilityLabel struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"hotkeyAccessibilityLabel,omitempty"`
								} `json:"hotkeyDialogSectionOptionRenderer"`
							} `json:"options"`
						} `json:"hotkeyDialogSectionRenderer"`
					} `json:"sections"`
					DismissButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"dismissButton"`
					TrackingParams string `json:"trackingParams"`
				} `json:"hotkeyDialogRenderer"`
			} `json:"hotkeyDialog"`
			BackButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"backButton"`
			ForwardButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"forwardButton"`
			A11YSkipNavigationButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					Size       string `json:"size"`
					IsDisabled bool   `json:"isDisabled"`
					Text       struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"a11ySkipNavigationButton"`
			VoiceSearchButton struct {
				ButtonRenderer struct {
					Style           string `json:"style"`
					Size            string `json:"size"`
					IsDisabled      bool   `json:"isDisabled"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										VoiceSearchDialogRenderer struct {
											PlaceholderHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"placeholderHeader"`
											PromptHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"promptHeader"`
											ExampleQuery1 struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"exampleQuery1"`
											ExampleQuery2 struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"exampleQuery2"`
											PromptMicrophoneLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"promptMicrophoneLabel"`
											LoadingHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"loadingHeader"`
											ConnectionErrorHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"connectionErrorHeader"`
											ConnectionErrorMicrophoneLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"connectionErrorMicrophoneLabel"`
											PermissionsHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"permissionsHeader"`
											PermissionsSubtext struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"permissionsSubtext"`
											DisabledHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"disabledHeader"`
											DisabledSubtext struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"disabledSubtext"`
											MicrophoneButtonAriaLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"microphoneButtonAriaLabel"`
											ExitButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Icon       struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													TrackingParams    string `json:"trackingParams"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
												} `json:"buttonRenderer"`
											} `json:"exitButton"`
											TrackingParams            string `json:"trackingParams"`
											MicrophoneOffPromptHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"microphoneOffPromptHeader"`
										} `json:"voiceSearchDialogRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"serviceEndpoint"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					Tooltip           string `json:"tooltip"`
					TrackingParams    string `json:"trackingParams"`
					AccessibilityData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityData"`
				} `json:"buttonRenderer"`
			} `json:"voiceSearchButton"`
		} `json:"desktopTopbarRenderer"`
	} `json:"topbar"`
	Refinements []string `json:"refinements"`
}
