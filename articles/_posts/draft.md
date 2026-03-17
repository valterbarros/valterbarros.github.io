Pontos:

Quero mesmo crisr o blog e pode escrever minhas criticas la
Mesmo que ninguém leia ou poucos a questão é mais conseguir por pra fora man
Seria mais como uma maneira de falar aquelas coisas que n posso falar com colegas de trabalho

Ideias:

- Falar sobre saido do javascript e como ta sendo a jornada
- [x] Falar sobre questão de devs com mente fechada, e acomodados.
  - [x] E como isso pode acabar atrapalhando a carreia do dev
- Falar sobre ia também e minha visão atual
- [x] Falar sobre sistemas legados, com vulnerabilidades
- [ ] Talvez poster sobre pr review e trabalho em equipa
---

Talvez outro post aqui:

### Back vs Front e liderança com vies

> Como meu background é muito mais forte no front vou falar sobre o front aqui, mas acredito que possa acontecer com outras linguagens também

Tenho visto também um certo padrão pelo menos nas empresas que passei, onde o front é ignorado, principalmente nesses pontos que citei, geralmente os tech leaders, ctos, heads de engenharia possuem um background de backend, com aqueles jargões de "aws", "subir uma lambda", "banco de dados", "throughput", etc. Esses muitas vezes nem sabem o que ta acontecendo e como o codigo frontend é feito.
isso dificulta bastante pois é certo que "de cima" não vira o direcionamento para melhorias. O ponto cego nisso é, mesmo que eu admita que o backend seja mais importante para a receita da empresa e saude
Não adianta criar um foguete no backend, com todas as bibliotecas atualizadas, boas praticas, baixa latencia, e todas as tecnicas avançadas possivel. Se na outra ponta o codigo é feito de forma
inferior.

Um exemplo seria trabalhar para melhorar a performance no lado de backend, porem no front as metricas de [core vitals](https://developers.google.com/search/docs/appearance/core-web-vitals) estão ruins, o endpoint vem em miliseconds, porem para abrir um dropdown o front trava por um tempo tendo metricas como [INP](https://web.dev/articles/inp) bem ruins. Ou mesmo para sistemas usados em mobile onde bundle de frontend é gigante e carrega scripts, css, imgs, htmls que talvez o usuário nem va precisar já no primeiro carregamento, torrando assim a quota de dados do usuario tornando a experiencia como um todo ruim. E se acha que isso não é importante tem aqui um [post do web.dev](https://web.dev/case-studies/vitals-business-impact) falando sobre o impacto dessas metricas para o negocio (ou seja dinheiro).

Muito se fala sobre: não deveriamos ter "linguagem favorita", e que o que devemos focar é na logica de programação, e entender o negocio, porem na pratica eu vejo uma dificuldade de um backend pelo menos
olhar para um codigo de javascript

---

Talvez outro post

### IA

Aqui entramos num ponto mais complicado atualmente com o avanço das IAs, as vezes sinto que a concorrencia agora não é mais os humanos e sim a IA em si, hoje trabalhos simples de programação e até intermediarios podem se resolvidos apenas com um prompt, então como um dev que nem atualiza suas dependencias nem busca melhorar seu codigo pode competir com isso? Por isso se faz ainda mais necessário a capacitação, o entendendimento do sistema, e não apenas resolver o problema sem se questionar como foi feito, pois isso a IA já consegue fazer bem.

---

### pr review 

Em muitos empregos o trabalho em equipe é importante e ajuda, como seres humanos podemos deixar detalhes passar que podem ser identificados por outros humanos ou hoje em dia com o uso da ia.

Para algumas pessoas pode ser mal visto em alguns momentos ter alguém dando opinião spbre seu trabalho, principalmente quando se leva para o pessoal e quando criamos um vínculo pessoal com o trabalho que foi feito

Como eu havia dito é importante ter uma segunda visão e opinião no nosso trabalho, pessoalmente não posso dizer que amo feebbakcs e comentários eu meus pull requestes, porem ai entramos num outro ponto: é trabalho e não tem como gostar de tudo que é relacionado a ele vamos sim ter coisas desconfortáveis porem precisamos fazer para conseguir evoluir como profissionais 

Acrdito que também ninguém é obrigado a aceitar uma sugestão, e em alguns momentos essas sugestões não fazem sentido mesmo. Geralmente coisas que n fazem sentido que vejo é perfeccionismo uma variável que ja descreve bem para que serve porem a pessoa ainda quer um nome mais descritivo

Alem disso otimização prematura, uma função usada em apenas um local e o commenter pede para separar a mesma num arquivo 

São cenarios que percebo n ser necessário executar 

Fala sobre prazo e que isso atrapalha, pq a galera quer mandar o pr nas carreiras
