---
layout: default
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

<hr>

<style>
  .section-informations {
    position: fixed;

    bottom: 0;
    translate: -50%;
    left: 50%;
  }
</style>

{% include contact.html locale="pt" %}
