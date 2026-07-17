FROM jekyll/jekyll:4.4.1

WORKDIR /srv/jekyll

COPY Gemfile ./
COPY Gemfile.lock ./

RUN bundle install

EXPOSE 4000

CMD ["bundle", "exec", "jekyll", "serve", "--watch", "--host", "0.0.0.0", "--port", "4000", "--livereload"]
