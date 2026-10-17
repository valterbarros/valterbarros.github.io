Vamos escrever basicamente o que esse projeto go deve fazer

O projeto atual do blog é gerenciado usando jekyll e o gh pages, então a parte de templates, includes, variaveis de template está resolvido

Existe um workflow hoje que eu preciso fazer manualmente e ele seria:
- criar um novo draft para um post, tudo que fica na pasta articles/drafts
- criar um novo post, tudo que fica na pasta articles/_posts
- converter um draft para um post ou vice versa
- "publicar" um novo post basicamente fazer o commit e push seguindo as regras para um novo post, passar layout title etc no metadata alem de colocar a data no nome do post
- listar todos os drafts e poder editar
- listar todos os posts e poder editar

Seria interessante:
- Um liveview para ver o markdown do post enquanto digito poderia ser no navegador
- Hoje eu preciso mandar o post para o google docs quando ele passa a ficar grande e complexo, se tivesse uma maneira barata de fazer isso seria interessante usando o cli claro
- Talves uma maneira de ir coletando ideias e dai ir colocando para depois ver o que realmente vale a pena e não

Esses pontos poderiam ser resolvidos usando uma ferramente de linha de comando
