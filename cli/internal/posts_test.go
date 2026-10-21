package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
	"valterbarros/blog-tools/internal/utils"
)

func TestCreatePosts(t *testing.T) {
	CreateNewPost()

	path, _ := utils.GetFromRoot(postsFolder)
	_, err := os.ReadFile(path + "/new-post.md")

	if err != nil {
		t.Error("The file was not created")
	}
}

func TestConvertToDraft(t *testing.T) {
	name := "post-test.md"
	err := ConvertToDraft(name)

	if err != nil {
		fmt.Println("an error occurred during convert", err)
	}
}

func TestFormatPost(t *testing.T) {
	t.Run("should add metadata", func(t *testing.T) {
		_, dir := createFile(t, templateA)

		FormatPost("post-test.md", dir)

		newPath := getRenamed(t, dir)
		data := utils.GetFileData(newPath)

		reg := regexp.MustCompile(`(?m)^title:|layout: post`)

		if len(reg.FindAllStringSubmatch(data, 2)) < 2 {
			t.Error("Not found metadata")
		}
	})

	t.Run("should not duplicate metadata", func(t *testing.T) {
		_, dir := createFile(t, templateD)

		FormatPost("post-test.md", dir)

		newPath := getRenamed(t, dir)
		data := utils.GetFileData(newPath)

		reg := regexp.MustCompile(`(?m)^title:|layout: post`)

		if len(reg.FindAllStringSubmatch(data, 2)) != 2 {
			t.Error("Not found metadata")
		}
	})

	t.Run("should add justify tags", func(t *testing.T) {
		_, dir := createFile(t, templateA)

		FormatPost("post-test.md", dir)

		newPath := getRenamed(t, dir)
		data := utils.GetFileData(newPath)

		ct := strings.Count(data, `<p align="justify">`)

		if ct == 0 {
			t.Error("Not found justify tags:", ct)
		}

		if ct != 5 {
			t.Error("More tags than necessary:", ct)
		}
	})

	t.Run("should not add justify tags", func(t *testing.T) {
		_, dir := createFile(t, templateB)

		FormatPost("post-test.md", dir)

		newPath := getRenamed(t, dir)
		data := utils.GetFileData(newPath)

		ct := strings.Count(data, `<p align="justify">`)

		if ct != 0 {
			t.Error("Found justify tags:", ct)
		}
	})

	t.Run("should replace links and bold style for html tags", func(t *testing.T) {
		_, dir := createFile(t, templateC)

		FormatPost("post-test.md", dir)

		newPath := getRenamed(t, dir)
		data := utils.GetFileData(newPath)

		reg := regexp.MustCompile(`(?m)<b>|<a href`)

		total := len(reg.FindAllStringSubmatch(data, 2))

		if total != 2 {
			t.Error("Not found metadata", total)
		}
	})
}

func TestAddHtmlTags(t *testing.T) {
	t.Run("should replace links and bold style for html tags", func(t *testing.T) {
		result := addHtmlTags(templateC)
		fmt.Println("result", result)

		countB := strings.Count(result, "<b>")
		countB2 := strings.Count(result, "</b>")
		countA := strings.Count(result, "<a href")
		countA2 := strings.Count(result, "</a>")
		countI := strings.Count(result, "<i>")
		countI2 := strings.Count(result, "</i>")
		total := []int{countB, countB2, countA, countA2, countI, countI2}
		expected := slices.Compare(total, []int{3, 3, 2, 2, 2, 2})

		if expected != 0 {
			t.Errorf("Not found tags %v", total)
		}
	})
}

func getRenamed(t *testing.T, dir string) string {
	t.Helper()
	files, _ := utils.ListFiles(dir)
	index := slices.IndexFunc(files, func(name string) bool {
		return strings.Contains(name, "post-test.md")
	})
	newPath := filepath.Join(dir, files[index])
	return newPath
}

func createFile(t *testing.T, template string) (string, string) {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "post-test.md")
	utils.SaveFile(path, template)

	return path, dir
}

var templateA = `
titulo

## Introdução

Antes de qualquer coisa gostaria de deixar claro que utilizo IA todos os dias no meu trabalho, como por exemplo: claude, cursor, copilot, dentro de IDEs como cursor e vscode, não sou de forma alguma uma pessoa contrária à IA. O meu intuíto é tentar enxergar o todo de forma consciente sem aquele véu de animação/desânimo com tecnologia X ou Y que às vezes pode nos cegar. Inclusive o post atual foi feito com ajuda de IA. 😀

Neste texto quero compartilhar alguns conhecimentos que reuni a partir de diversos vídeos e conteúdos que venho acompanhando de perto nos últimos anos. A ideia aqui é focar menos na parte técnica de como a IA funcionar e olhar mais para como os programadores podem se preparar para as mudanças que vêm acontecendo, além de falar sobre como o marketing usado para promover a Inteligência Artificial foi construído e a forma como essas informações foram transmitidas acabou afetando o psicológico e a carreira de muitas pessoas pelo mundo.

Quero também trazer a importância do questionamento e o fato disso ser um dos pilares da área de programação e os perigos de silenciar isso.

Também vou trazer alguns pontos que são consenso entre os conteúdos que busquei em relação a como a IA está está mudando o papel do programador, que está ficando um pouco mais distante do código e atuando mais como um revisor, trabalhando no levantamento de requisitos, conversando com o cliente, validando requisitos, etc.

## O marketing do medo

É impossível abrir o linkedin ou youtube (sendo programador e os algoritmos sabendo disso) e não ver vários vídeos falando sobre o fim dos programadores e de outras profissões, bem negativo. Isso que estamos vendo, é o que podemos chamar de **marketing do medo** e pode ser explicado por: as empresas de IA dependem bastante de investimentos externos, é fácil ver isso pelos aportes bilionários feitos nos últimos anos, contudo a maneira encontrada para justificar todo esse investimento foi: substituição do quadro de funcionários das empresas por IA.`

var templateB = `Links anteriores sobre pessoas influentes falando sobre o fim dos programadores:

* [Sam Altman: AI Will Not Only Replace Experts But Create New Knowledge \- Fevereiro/2023](https://medium.datadriveninvestor.com/sam-altman-ai-will-not-only-replace-experts-but-create-new-knowledge-6d0726d8ef6e)
* [Stability AI CEO: There Will Be No (Human) Programmers in Five Years](https://finance.yahoo.com/news/stability-ai-ceo-no-human-193413248.html)  
* [In a leaked recording, Amazon cloud chief tells employees that most developers could stop coding soon as AI takes over \- Agosto/2024](https://www.businessinsider.com/aws-ceo-developers-stop-coding-ai-takes-over-2024-8)`

var templateC = `**Uma coisa que acho** que movimenta muito e reforça essa ideia e mantém esse medo no ar é o fato de CEOs de grandes empresas como; nvidia, amazon, openIA afirmam que em X meses/anos/dias a IA vai substituir todo mundo, apesar de alguns hoje estarem voltando atrás como é o caso do [Jenssen huang CEO da Nvidia acénto](https://tecnoblog.net/noticias/jensen-huang-critica-negatividade-em-torno-da-ia/?test=2) que falou que para os próximos anos vamos precisar muito de programadores.

*Neste texto quero* compartilhar alguns conhecimentos que reuni a partir de diversos **vídeos e conteúdos** que venho acompanhando de perto nos últimos anos. A ideia aqui é focar menos na parte técnica de como a IA funcionar e olhar mais para como [os programadores](https://testando.com) podem **se preparar para as** mudanças que vêm acontecendo, além de falar sobre como o marketing usado para promover a Inteligência Artificial foi construído e a forma como essas informações foram transmitidas acabou afetando o psicológico e a carreira de muitas *pessoas pelo mundo.*`

var templateD = `
---
layout: post
title: [title]
---

**Uma coisa que acho** que movimenta muito e reforça essa ideia e mantém esse medo no ar é o fato de CEOs de grandes empresas como; nvidia, amazon, openIA afirmam que em X meses/anos/dias a IA vai substituir todo mundo, apesar de alguns hoje estarem voltando atrás como é o caso do [Jenssen huang CEO da Nvidia acénto](https://tecnoblog.net/noticias/jensen-huang-critica-negatividade-em-torno-da-ia/?test=2) que falou que para os próximos anos vamos precisar muito de programadores.

*Neste texto quero* compartilhar alguns conhecimentos que reuni a partir de diversos **vídeos e conteúdos** que venho acompanhando de perto nos últimos anos. A ideia aqui é focar menos na parte técnica de como a IA funcionar e olhar mais para como [os programadores](https://testando.com) podem **se preparar para as** mudanças que vêm acontecendo, além de falar sobre como o marketing usado para promover a Inteligência Artificial foi construído e a forma como essas informações foram transmitidas acabou afetando o psicológico e a carreira de muitas *pessoas pelo mundo.*`
