all:
	@echo building the frontend 
	cd scss && node-sass --output-style compressed app.sass ../../backend/static/jass.css
	temple build templates template.go --package main
	GOOS=linux gopherjs build *.go *.inc.js -o ../backend/static/jass.js -m
	cp assets/*.html ../backend/static
	cp assets/*.webmanifest ../backend/static
	cp assets/favicon.ico ../backend/static
	cp -R assets/img ../backend/static/img
	cp -R assets/fonts ../backend/static/fonts
	cp -R assets/css ../backend/static/css
	@ls -ltraR ../backend/static
