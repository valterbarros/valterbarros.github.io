Pontos:

Quero mesmo crisr o blog e pode escrever minhas criticas la
Mesmo que ninguém leia ou poucos a questão é mais conseguir por pra fora man
Seria mais como uma maneira de falar aquelas coisas que n posso falar com colegas de trabalho pois eles se doem

Ideias:

- Falar sobre saido do javascript e como ta sendo a jornada
- Falar sobre questão de devs com mente fechada, e acomodados.
  - E como isso pode acabar atrapalhando a carreia do dev
- Falar sobre ia também e minha visão atual
- Falar sobre sistemas legados, com vulnerabilidades

## O ambiente modificavel da area de tecnologia

Quando comecei minha carreira eu trabalhava com, jquery, js, ruby on rails, me lembro que na epoca eu precisava tocar numa serie de tecnologias e linguagens diferentes
como sql, css, scss, cofeescript. E sempre era dificil manter o nivel de qualidade em todas elas. Acredito que esse seja o cenario normal para qualquer dev. Desde então novas tecnologias surgiram e outras morreram, atualmente eu trabalho com vue.js, ruby on rails, e nossa a coisa mudou muito de cenario, praticamente tive que reaprender tudo

Com certeza se eu me mantivesse nessa mesma stack de anos atrás eu teria muita dificuldade em encontrar emprego e me manter em algum, sem me atualizar apenas me mantendo na mesma, existe emprego hoje para jquery? provavelmente não muitos.

Em alguns momentos essa mudança constante é muito complicada e desgastante pois, como seres humanos vamos sempre buscar algum conforto, o que é chamado hoje de "zona de conforto", para os coachs isso é algo que uma pessoa deve ter aversão e sempre procurar estar desconfortavel se desafiando e encontrando novos limites ainda não cruzados. Não concordo com essa ideia inclusive a psicologia moderna também não, e talvez essa ideia possa ser a causa de muito burnout que vemos hoje em dia. Acredito que devemos nos desafiar porem respeitando os limites

### Projeto Legado

Na area de tech isso pega um pouco mais pois, sempre vão ter novas tecnologias surgindo, e mesmo as que já existem continuam diariamente evoluindo, e isso nos faz ter que correr atrás para aprender. Isso me leva ao ponto principal que quero falar: trabalhar com projetos legados é um perigo, basicamente o projeto roda, os clientes estão acostumados, tem bugs mas são ignorados "se ninguem reclama deixa quieto ai", vulerabilidades também são ignoradas "se ninguem invadiu ate hoje ninguem vai mais". Bibliotecas a anos sem atualização muitas vezes por falta de confiança da equipe de mexer e quebrar algo que já está estavel rodando em produção. E dessa forma vão utilizando versões antigas e com sintaxe completamente deprecated e afundando cada vez mais esse projeto e criando uma dissonancia em relação ao mercado externo, e afasta o dev cada vez mais de se capacitar.

> Porem não sou o cara do hype, se uma determinada lib atende bem o sistema e consegue resolver os problemas, não vejo necessidade de fazer alguma troca, até pq existe um custo em trocar, correr o risco de quebrar em prod, ou de lá na frente essa lib deixar de ser mantida. Infelizmente essas mudanças de "agua para water" é algo que acontece muito no mundo tech onde ferramentas para fazer exatamente a mesma coisa são criadas a todo momento, com pretesto de serem mais rapidas ou com uma sintaxe melhor, porem na grande maioria é apenas mais do mesmo com "um roupa mais bonita". Muitas ideias hoje da tecnologia foram criadas a anos, e hoje são relançadas como algo novo.

Para a empresa isso quer dizer que novos usuários vão ter dificuldade para utilizar o sistema justamente por não estarem acostumados com os bugs considerados "normais", dessa forma a empresa não consegue agregar receita, alem de diminuir a confiança do usuário, essa empresa também vai ter dificuldade em fechar contratos com empresas grandes que geralmente tendem a se preocupar mais com segurança e ter baixa tolerancia a problemas de segurança como bibliotecas desatualizadas. Além de que essa empresa vai ter devs menos capacitados utilizando tecnologias antigas e muitas vezes a anos desaconselhada pela comunidade no geral.

Para o dev quer dizer que ele é suficiente para essa empresa(e talvez apenas isso mesmo) e em alguns casos ele acaba sendo indispensavel, mas será que ele consegue se realocar em outra empresa caso precise? Sera que ele consegue passar em uma entrevista apenas com esses conhecimentos "legados".

Para o usuário, quer dizer que ele vai pagar por um sistema, que tem baixas expectativas de qualidade. Lentidões para executar processos simples, carregamentos demorados, e insegurança com seus dados.

### IA

Aqui entramos num ponto mais complicado atualmente com o avanço das IAs, as vezes sinto que a concorrencia agora não é mais os humanos e sim a IA em si, hoje trabalhos simples de programação e até intermediarios podem se resolvidos apenas com um prompt, então como um dev que nem atualiza suas dependencias nem busca melhorar seu codigo pode competir com isso? Por isso se faz ainda mais necessário a capacitação, o entendendimento do sistema, e não apenas resolver o problema sem se questionar como foi feito, pois isso a IA já consegue fazer bem.

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

O ponto que quero levantar é, se vc acha que está seguro no seu emprego, será mesmo que está ou isso é uma ilusão?


