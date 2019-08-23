package html

// Título obtem o título de uma página html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) { // Função Anônima
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Problema com método get ao pegar a url.")
			}
			html, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Erro ao ler o corpo do response ou html.")
			}
			// Pegando uma tag do html pegando o título
			r, err := regexp.Compile("<title>(.*?)<\\/title>")
			if err != nil {
				fmt.Println("Ocorreu um erro no Regex!")
			}
			// Convertendo em string o html
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) // Já ao mesmo tempo invocando a função logo após o término.
	}
	// Retornando o channel
	return c