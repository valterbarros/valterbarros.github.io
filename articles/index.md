---
layout: minima
title: Meu Blog
---

## Posts Recentes

<html>
<head>
    <title>Meu Blog</title>
    <style>
        body { font-family: sans-serif; max-width: 800px; margin: 40px auto; padding: 0 20px; line-height: 1.6; }
        li { margin-bottom: 10px; }
    </style>
</head>
<body>
    <h1>Posts do Blog</h1>
    <ul>
      {% for post in site.posts %}
        <li>
          {{ post.date | date: "%d/%m/%Y" }} — 
          <a href="{{ post.url }}">{{ post.title }}</a>
        </li>
      {% endfor %}
    </ul>
    <hr>
    <a href="/">← Voltar para a Home Principal</a>
</body>
</html>
