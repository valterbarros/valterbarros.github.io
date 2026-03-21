---
layout: default
title: Meu Blog
lang: pt-BR
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

<hr>

<style>
  .section-informations {
    position: absolute;
    translate: -50%;
    left: 50%;
    bottom: 0;
  }
</style>

{% include contact.html locale="pt" %}
