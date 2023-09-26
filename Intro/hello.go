// indica que esse é o pacote principal
// Go é uma linguagem compilada
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// import "reflect"

const monitoramento = 1
const delay = 1

func main() {

	// declarando variavel sintaxe
	// Em go nao e necessario declarar as variaveis int e string, ele infere qual o tipo da variavel automaticamente, as demais por boa pratica e bom declarar mas funcionaria tb
	// var nome = "Vitor"
	// var versao = 1.0
	// Ao declarar variaveis em go e não atribuir valor ele será 0 para números e espaço em branco para letras
	// atribuidor curto de variavel
	// idade := 20
	// Em go toda variavel declarada tem que ser usada se nao resultara em erro
	// fmt.Println("Olá, sr.", nome, "sua idade é:", idade)
	//

	// fmt.Println("O tipo da variável versao é:", reflect.TypeOf(versao))
	// fmt.Println("1- Iniciar Monitoramento")
	// fmt.Println("2- Exibir Logs")
	// fmt.Println("0- Sair do programa")

	// O problema é que se o valor passado for de um tipo diferente a variavel vai ficar em branco ou zerada ou com o valor atribuido inicialmente
	// var comando int
	// fmt.Scan(&comando)
	// fmt.Println("O comando escolhido foi ", comando)

	// Em go ao utilizar if e necessario colocar um valor que retorne um resultado booleano
	// Em go nao e necessario usar parenteses na funcao if

	// E possivel utilizar tambem o comando switch que se assemelha ao case em outras linguagens, ele tem a vantagem de nao precisar utilizar break
	// switch comando {
	// case 1:
	// 	fmt.Println("Monitorando...")
	// case 2:
	// 	fmt.Println("Exibindo Logs")
	// case 0:
	// 	fmt.Println("Saindo do programa")
	// default:
	// 	fmt.Println("Não conheço esse comando")
	// }

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		if comando == 1 {
			iniciarMonitoramento()
		} else if comando == 2 {
			fmt.Println("Exibindo Logs")
			imprimeLogs()
		} else if comando == 0 {
			fmt.Println("Saindo do programa")
			os.Exit(0)
		} else {
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	var nome = "Vitor"
	var versao = 1.0
	idade := 20
	fmt.Println("Olá, sr.", nome, "sua idade é:", idade)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi ", comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSites()
	// 	"https://www.alura.com.br",
	// 	"https://www.caelum.com.br"}

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "Esta com problemas.")
		registraLog(site, false)

	}

}

func leSites() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)

		fmt.Println(linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " -online: " + strconv.FormatBool(status) + "\n" )

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))

}