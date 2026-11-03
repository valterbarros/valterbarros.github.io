---
layout: post
title: Go na prática; automatizando o fluxo de posts do meu blog Jekyll
---

![imagem]({{"articles/imgs/2026-11-0311:55:49.3111304-03:00image.png" | relative_url }})

### Estudando Go

Para quem não me conhece, sou programador frontend, mas me interessei também pela área de backend e venho estudando Go. Estou gostando cada vez mais da linguagem, do sistema de tipos, da forma como ela estrutura os projetos e de como incentiva uma arquitetura simples. Tenho estudado bastante, lendo a documentação oficial e fazendo o Go Tour, mas todo esse estudo precisa ser colocado em prática. Por isso, parei para pensar em alguma demanda minha que pudesse ser automatizada.

Foi aí que me voltei para o [meu blog](https://valterbarros.github.io/articles/), que hoje é construído com Jekyll. Meu fluxo atual para criar e editar posts é feito direto no VSCode, mas ele tem pontos fracos: o preview de Markdown às vezes não sincroniza com o editor e o scroll fica estranho.

Além disso, como os posts do Jekyll são arquivos .md, preciso de uma forma prática de criar, deletar e renomear esses arquivos. Então decidi construir uma ferramenta em Go para resolver isso.

### O editor

Como eu tinha dito acima a edição de posts acaba sendo precária dentro do VSCode, e dai eu criei um editor usando Go, com preview em tempo real usando web-socket para a comunicação cliente/server, também adicionei nesse mesmo editor a possibilidade de colar imagens dentro dos textos, com isso consegui por em pratica:

- Servidor http
- Upgrade de http para web-socket e comunicação em tempo real
- Criação de arquivo e copia de bytes entre arquivos
- Serialize de JSON
- Uso de template HTML e text com variáveis dentro de Go
- Parse de markdown
- Testes unitários e integração 

E essa é a aparência final do editor + preview

![imagem]({{"articles/imgs/2026-11-0311:54:25.355995381-03:00image.png" | relative_url }})

### CLI

Um complemento ao editor+preview é a gerencia de arquivos .md dos posts então eu criei um CLI com alguns comandos utilitários para facilitar esse processo usando a lib [cobra](https://cobra.dev), que é muito eficiente para a criação de CLI tendo suporte a adição de flags nos comandos, sub comandos, documentação, help, etc.

Essa é basicamente a saída desses commandos:

```bash
./blogging

utils to create/list/delete posts.

Usage:
  blogging posts [command]

Available Commands:
  convert     Convert post to draft
  create      Create post
  delete      Delete a post
  edit        Edit posts
  format      Format a post
  list        List posts

Flags:
  -h, --help   help for posts

Use "blogging posts [command] --help" for more information about a command.
```

### Usando IA

Fiz uso de IA, mas optei por escrever as partes menos repetitivas por conta própria para tentar aprender mais, a ia também me ajudou a planejar a execução para antes mesmo de tocar no código ver os tradeoffs e quais as decisões que precisavam ser tomadas.

### Conclusão

Com esses commandos eu consigo fazer todo esse gerenciamento via terminal, foi uma experiencia bem interessante trabalhar nesse CLI consegui aplicar muita coisa do que eu estava aprendendo em algo que vou utilizar.
