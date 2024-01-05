# Heroicons Templ

[Heroicons](heroicons.com) icons compiled as [templ](templ.guide) components

## Usage

```bash
go get github.com/tvaintrob/heroicons-templ
```

To use in a `templ` file:

```templ
package main

// import the specific icon package (solid / outline) and size (16 / 20 / 24)
import "github.com/tvaintrob/heroicons-templ/solid/24"

templ ExamplePage() {
  ...
	@solid24.HeartIcon(templ.Attributes{})
  ...
}
```

For a working example refer to [./example/README.md]
