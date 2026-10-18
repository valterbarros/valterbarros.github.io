Vamos escrever basicamente o que esse projeto go deve fazer

O projeto atual do blog é gerenciado usando jekyll e o gh pages, então a parte de templates, includes, variaveis de template está resolvido

Existe um workflow hoje que eu preciso fazer manualmente e ele seria:
- [x] criar um novo draft para um post, tudo que fica na pasta articles/drafts
- [x] criar um novo post, tudo que fica na pasta articles/_posts
- [x] converter um draft para um post ou vice versa
- [] "publicar" um novo post basicamente fazer o commit e push seguindo as regras para um novo post, passar layout title etc no metadata alem de colocar a data no nome do post
- [x] listar todos os drafts
- [x] listar todos os posts
- [x] Abrir um post ou draft diretamente no vscode sem estar com ele aberto

Seria interessante:
- Um liveview para ver o markdown do post enquanto digito poderia ser no navegador
- Hoje eu preciso mandar o post para o google docs quando ele passa a ficar grande e complexo, se tivesse uma maneira barata de fazer isso seria interessante usando o cli claro
- Talves uma maneira de ir coletando ideias e dai ir colocando para depois ver o que realmente vale a pena e não

Esses pontos poderiam ser resolvidos usando uma ferramente de linha de comando

Fora do escopo

- Será que da pra substituir o jekyll? no caso fazer na mão mesmo
- Funcionalidades atuais do jekyll
  - Fazer include de uma pagina dentro de outra
  - Lidar com variaveis que ficam no topo da pagina
  - Arquivo de config com algumas configs
  - Montar paginas dinamicamentes, com for e ifs, um bom exemplo seria a index de posts que exibe todos os posts
  - Lidar com imports de assets estaticos imagens etc
  - Lidar com layouts

- Paginas do blog e pagina inicial
  - Pagina inicial do site em ptbr
  - Pagina inicial do site em ingles
  - Pagina de listagem de artigos
  - Pagina do artigo
