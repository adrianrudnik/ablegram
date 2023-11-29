package abletonsrc

/* Ableton MajorVersion="5" MinorVersion="11.0_11300"
<SampleRef>
	<FileRef>
		<RelativePathType Value="0" />
		<RelativePath Value="" />
		<Path Value="/Users/adrian/pCloud Drive/Projekte/Ablegram/Sample Project/CantinaBand3.wav" />
		<Type Value="2" />
		<LivePackName Value="" />
		<LivePackId Value="" />
		<OriginalFileSize Value="132344" />
		<OriginalCrc Value="56970" />
	</FileRef>
	<LastModDate Value="1699726595" />
	<SourceContext />
	<SampleUsageHint Value="0" />
	<DefaultDuration Value="66150" />
	<DefaultSampleRate Value="22050" />
</SampleRef>
*/

type XmlSampleRef11Node struct {
	SampleReference XmlSampleRef11 `xml:"SampleRef"`
}

type XmlSampleRef11 struct {
	XmlFileRef11Node
	DefaultSampleRate XmlIntValue `xml:"DefaultSampleRate"`
	DefaultDuration   XmlIntValue `xml:"DefaultDuration"`
}

type XmlFileRef11Node struct {
	FileReference XmlFileRef11 `xml:"FileRef"`
}

type XmlFileRef11 struct {
	RelativePathType XmlIntValue    `xml:"RelativePathType"`
	RelativePath     XmlStringValue `xml:"RelativePath"`
	Path             XmlStringValue `xml:"Path"`
	Name             XmlStringValue `xml:"Name"`
	OriginalFileSize XmlIntValue    `xml:"OriginalFileSize"`
}

/* Ableton MajorVersion="5" MinorVersion="10.0_370"
<SampleRef>
	<FileRef>
		<HasRelativePath Value="true" />
		<RelativePathType Value="5" />
		<RelativePath>
			<RelativePathElement Id="9" Dir="Samples" />
			<RelativePathElement Id="10" Dir="Sounds Of The 70s" />
			<RelativePathElement Id="11" Dir="115-Children Of The Sun" />
		</RelativePath>
		<Name Value="COTS-Drums.aif" />
		<Type Value="2" />
		<Data>...</Data>
		<RefersToFolder Value="false" />
		<SearchHint>
			<PathHint>
				<RelativePathElement Id="5" Dir="trunk" />
				<RelativePathElement Id="6" Dir="Vinyl Classics" />
				<RelativePathElement Id="7" Dir="Samples" />
				<RelativePathElement Id="8" Dir="Sounds Of The 70s" />
				<RelativePathElement Id="9" Dir="115-Children Of The Sun" />
			</PathHint>
			<FileSize Value="1104854" />
			<Crc Value="5728" />
			<MaxCrcSize Value="16384" />
			<HasExtendedInfo Value="true" />
		</SearchHint>
		<LivePackName Value="Vinyl Classics" />
		<LivePackId Value="www.ableton.com/33" />
	</FileRef>
	<LastModDate Value="1533741595" />
	<SourceContext>
		<SourceContext Id="0">
			<OriginalFileRef>
				<FileRef Id="2">
					<HasRelativePath Value="true" />
					<RelativePathType Value="5" />
					<RelativePath>
						<RelativePathElement Id="12" Dir="Samples" />
						<RelativePathElement Id="13" Dir="Sounds Of The 70s" />
						<RelativePathElement Id="14" Dir="115-Children Of The Sun" />
					</RelativePath>
					<Name Value="COTS-Drums.aif" />
					<Type Value="2" />
					<Data>...</Data>
					<RefersToFolder Value="false" />
					<SearchHint>
						<PathHint>
							<RelativePathElement Id="5" Dir="trunk" />
							<RelativePathElement Id="6" Dir="Vinyl Classics" />
							<RelativePathElement Id="7" Dir="Samples" />
							<RelativePathElement Id="8" Dir="Sounds Of The 70s" />
							<RelativePathElement Id="9" Dir="115-Children Of The Sun" />
						</PathHint>
						<FileSize Value="1104854" />
						<Crc Value="5728" />
						<MaxCrcSize Value="16384" />
						<HasExtendedInfo Value="true" />
					</SearchHint>
					<LivePackName Value="Vinyl Classics" />
					<LivePackId Value="www.ableton.com/33" />
				</FileRef>
			</OriginalFileRef>
			<BrowserContentPath Value="" />
		</SourceContext>
	</SourceContext>
	<SampleUsageHint Value="0" />
	<DefaultDuration Value="184070" />
	<DefaultSampleRate Value="44100" />
</SampleRef>
*/

type XmlSampleRef9Node struct {
	SampleReference XmlSampleRef9 `xml:"SampleRef"`
}

type XmlSampleRef9 struct {
	XmlFileRef9Node
	DefaultSampleRate XmlIntValue `xml:"DefaultSampleRate"`
	DefaultDuration   XmlIntValue `xml:"DefaultDuration"`
}

type XmlFileRef9Node struct {
	FileReference XmlFileRef9 `xml:"FileRef"`
}

type XmlFileRef9 struct {
	SearchHint       XmlSearchHint9            `xml:"SearchHint"`
	HasRelativePath  XmlBooleanValue           `xml:"HasRelativePath"`
	RelativePathType XmlIntValue               `xml:"RelativePathType"`
	RelativePath     []XmlFileRef9RelativePath `xml:"RelativePath>RelativePathElement"`
	Name             XmlStringValue            `xml:"Name"`
}

func (x *XmlFileRef9) RelativePathFolders() []string {
	var parts []string

	for _, e := range x.RelativePath {
		parts = append(parts, e.Folder)
	}

	return parts
}

type XmlSearchHint9 struct {
	PathHint []XmlFileRef9RelativePath `xml:"PathHint>RelativePathElement"`
	FileSize XmlIntValue               `xml:"FileSize"`
}

func (x *XmlSearchHint9) PathHintFolders() []string {
	var parts []string

	for _, e := range x.PathHint {
		parts = append(parts, e.Folder)
	}

	return parts
}

type XmlFileRef9RelativePath struct {
	Id     int64  `xml:"Id,attr"`
	Folder string `xml:"Dir,attr"`
}
