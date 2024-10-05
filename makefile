watch-assets:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

templ:
	# templ generate --watch
	@templ generate --watch --proxy="http://localhost:8091" --open-browser=false

server:
	air

dev:
	@make -j3 server templ watch-assets
