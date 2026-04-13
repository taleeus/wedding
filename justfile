tailwind_cmd := join("bin", os(), "tailwindcss")
templui_path := `go list -m -f \{\{.Dir\}\} github.com/templui/templui`

default: dev

[parallel]
dev: tailwind-watch templ-watch

tailwind:
    @{{ tailwind_cmd }} -i ./static/css/input.css -o ./static/css/style.css

tailwind-watch:
    @{{ tailwind_cmd }} -i ./static/css/input.css -o ./static/css/style.css --watch

templ:
    @go tool templ generate

templ-watch:
    @wgo -file=.go -file=.templ -xfile=_templ.go go tool templ generate :: go run main.go
