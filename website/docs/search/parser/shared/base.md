### `tags`

<IndexProp name="tags%" type="array<string>" mapping="exact">

Tags associated to the document. They are sourced by the parser, for this specific result type.

- `+tags:"tagName"` will match all documents that have the given tag. Ensure you use double-quotes characters around the tag name if it contains colons.
</IndexProp>

### `displayName`

<IndexProp name="displayName%" type="string" mapping="fulltext">

Display name given by the parser, to be the title of the resulting document.

- `+displayName:MID` will match all documents that have a value of `MID` or a similar one in their display name, like `MIDI`.
</IndexProp>

### `pathAbsolute`

<IndexProp name="pathAbsolute%" type="string" mapping="exact">

The absolute path of the file the found document was sourced from. This can be used to filter for everything related
to a single file.

- `+pathAbsolute:/Users/adrian/Downloads/test.als` will match all documents with that exact path.
</IndexProp>

### `pathFolder`

<IndexProp name="pathFolder%" type="string" mapping="exact">

The path part of the `pathAbsolute` property. This can be used to filter for everything related to a single folder.

- `+pathFolder:/Users/adrian/Downloads` will match all documents that were sourced from the given folder.
</IndexProp>

### `filename`

<IndexProp name="filename%" type="string" mapping="exact">

The filename part of the `pathAbsolute` property. This can be used to filter for everything related to a single file.

- `+filename:test.als` will match all documents that were sourced from files with the exact given filename.
</IndexProp>
