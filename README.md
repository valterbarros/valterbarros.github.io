# My-own-website
builded in html/css bootstrap

# docker

```
docker build -t valterbarros-jekyll .
docker run --rm -it -p 4000:4000 -v "$PWD":/srv/jekyll -w /srv/jekyll valterbarros-jekyll
```
