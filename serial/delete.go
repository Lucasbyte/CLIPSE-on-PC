package serial

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func testPort(p string, b int) bool {
	config := &serial.Config{
		Name:        p,
		Baud:        b,
		ReadTimeout: time.Second,
	}

	porta, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println("Erro ao abrir porta serial:", err)
		return false
	}
	defer porta.Close()

	fmt.Println(porta)
	return true
}

func listarItens(porta *serial.Port, codigo string) string {
	_, err := porta.Write([]byte(codigo))
	if err != nil {
		log.Fatalf("Erro ao escrever na porta serial: %v", err)
	}
	buf := make([]byte, 1)
	teste := ""
	c := 0
	for {
		c++
		n, err := porta.Read(buf)
		if err != nil {
			log.Fatalf("Erro ao ler da porta serial: %v", err)
		}
		if n == 0 || c == 3 {
			fmt.Print(c)
			break // Sai do loop se não há mais dados disponíveis
		}
		teste = fmt.Sprint(teste + string(buf[:n]))
	}

	return teste
}

func Delete(p int, speed int, plus []string) {

	porta := fmt.Sprintf("COM%d", p)
	velocidade := speed

	fmt.Printf("CONFIG: %s %d\n", porta, velocidade)

	if testPort(porta, velocidade) {
		fmt.Printf("Porta %s disponível.\n", porta)

		config := &serial.Config{
			Name:        porta,
			Baud:        velocidade,
			ReadTimeout: time.Second,
		}
		porta, err := serial.OpenPort(config)
		if err != nil {
			fmt.Printf("Erro ao abrir porta serial: %v s\n", err)
		}
		defer porta.Close()

		for _, value := range plus {
			resultado := listarItens(porta, fmt.Sprintf("\x04%s\x02", value))
			fmt.Printf("Resposta para %s: %s\n", value, resultado)
		}
	} else {
		fmt.Printf("Porta %s não disponível.\n", porta)
	}
}
