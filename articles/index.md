---
layout: minima
title: Meu Blog
---

# Artigos

Aqui está a lista de posts:

<ul>
  {% for post in site.posts %}
    <li>
      <span>{{ post.date | date: "%d/%m/%Y" }}</span> — 
      <a href="{{ post.url }}">{{ post.title }}</a>
    </li>
  {% endfor %}
</ul>
