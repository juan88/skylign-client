# skylign-client

Wrapper for generating Interactive logos for alignments and profile HMMs using
[Skylign.org](http://skylign.org/) server through its API.

## Instalation

Standalone client can be built just like any other Go app by building it in the standard way
`go build` inside the root directory.

`client` package for including the connector into other apps can be used in the same way
as the skylign-client app is using it so that you can use it elsewhere.

## Usage

```bash
skylign-client --file=filename_to_upload [--height=height_mode --processing=processing_mode --frag=fragmentation]
```

Parameters:

* `file`  the name of the file to be processed.
* `height`  Which height calculation algorithm should be used during logo generation. There are three different algorithms to choose from: *info_content_all*, *info_content_above*, *score*.
* `processing`  The type of post processing to be performed on an alignment: *observed*, *weighted*, *hmm*,*hmm_all*.

For more info checkout the official API [docs][20da26b6]

[20da26b6]: http://skylign.org/help/api/post "API docs"
