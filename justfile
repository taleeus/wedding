tailwind_cmd := join("bin", os(), "tailwindcss")

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
    @go tool wgo -file=.go -file=.templ -xfile=_templ.go go tool templ generate :: go run main.go
