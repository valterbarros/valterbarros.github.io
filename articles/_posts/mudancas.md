---
layout: post
title: O perigo invisível dos projetos legados
---

### Introdução

No começo da minha carreira eu trabalhava como fullstack, me lembro que na época eu precisava lidar numa série de tecnologias e linguagens diferentes como sql, css, scss, coffeescript, jquery, js, ruby on rails, etc. Acredito que esse seja um cenário "comum" para muitos devs. Desde então novas tecnologias surgiram e outras morreram, muita coisa que eu fazia nessa época mudou e algumas até radicalmente.

Essa mudança constante pode ser difícil e desgastante pois, como seres humanos é normal buscar conforto, o que é chamado hoje de "zona de conforto", para os coachs isso é algo que uma pessoa deve ter aversão e sempre procurar estar desconfortável. Não concordo com essa ideia, inclusive a psicologia moderna também não, e talvez essa ideia possa ser a causa de muito burnout que vemos hoje em dia. Acredito que devemos nos desafiar porém respeitando os limites.

### Projeto Legado

Na área de tecnologia isso pega um pouco mais pois, novas tecnologias surgem muito mais rápido do que podemos nos atualizar, e mesmo as que já existem continuam diariamente evoluindo, isso nos força a uma corrida constante. Mas o perigo real surge quando paramos: os projetos legados. Muitas vezes o projeto roda, os clientes estão acostumados, tem bugs mas são ignorados "se ninguém reclama deixa quieto ai", vulnerabilidades também são ignoradas "se ninguém invadiu até hoje ninguém vai mais", bibliotecas a anos sem atualização muitas vezes por falta de confiança da equipe de mexer e quebrar algo que já está estável rodando em produção. E dessa forma vão utilizando versões antigas e com sintaxe completamente deprecated e afundando cada vez mais esse projeto e criando uma dissonância em relação ao que é praticado no mercado fora daquela empresa.

> Porém não sou o cara do hype, se uma determinada biblioteca atende bem o sistema e consegue resolver os problemas, não vejo necessidade de trocar, até porque existe toda mudança possui um custo, como correr o risco de quebrar algo em produção.

Temos 3 sujeitos que são afetados de formas diferentes pelos projetos legados:

**1. A Empresa:**

Para a organização isso quer dizer que novos usuários vão ter dificuldade para utilizar o sistema justamente por não estarem acostumados com os bugs considerados "normais", dessa forma a empresa não consegue agregar receita, além de diminuir a confiança do usuário, essa empresa também vai ter dificuldade em fechar contratos com empresas grandes que geralmente tendem a se preocupar mais com segurança e ter baixa tolerância a problemas de segurança como bibliotecas desatualizadas.

**2. Para o dev:**

Quer dizer que ele é suficiente para essa empresa(e talvez apenas isso mesmo) e em alguns casos ele acaba sendo indispensável, mas será que ele consegue se realocar em outra empresa caso precise? Será que ele consegue passar em uma entrevista apenas com esses conhecimentos "legados".

**3.Para o usuário**

Para o usuário quer dizer que ele vai pagar por um sistema, que tem baixas expectativas de qualidade. Lentidão para executar processos simples, carregamentos demorados, e insegurança com seus dados.

### O que fazer?

Mas aí vem a pergunta de um milhão: como se manter atualizado sem virar refém de cada nova biblioteca que aparece?

O segredo está em filtrar o que é puro hype do que realmente resolve problemas. Antes de sair trocando tudo, vale parar e fazer um 'check-up' de realidade:

- Fôlego a longo prazo: Essas melhorias vão se sustentar daqui a dois ou três anos, ou são só a moda do momento?
- Custo de oportunidade: Realmente vale o esforço de mexer em algo que já entrega valor hoje? O ganho justifica o risco?
- Saúde do ecossistema: Tem gente de verdade usando e mantendo isso? Se a comunidade não for engajada, você corre o risco de herdar um novo legado (e dessa vez, um que ninguém mais conhece) em pouco tempo."

Como eu havia citado anteriormente, não precisamos ficar a todo momento mudando a forma como as coisas são feitas apenas por ego ou para deixar o código mais "pop". Porém, por outro lado, também não podemos largar o projeto às traças e deixar de fazer o básico, como: atualizar dependências, fazer refatorações e levantamentos de débito técnico.

Mudar é difícil. Novas iniciativas dentro da empresa parecem desconfortáveis e às vezes parece não valer a pena mudar a forma como o código é feito, além do medo que falei sobre quebrar algo que está estável para os clientes de longo prazo.

Mas, como indivíduos, precisamos pensar em nossas carreiras. Afinal, o que seria melhor: modernizar a forma como faço meu trabalho ou ser um dev apenas de legado e me manter desatualizado?


