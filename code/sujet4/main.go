package main

import (
	"fmt"
	"sync"
	"time"
)

// Travail simulé par une goroutine
func worker(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	// Simuler un travail prenant du temps
	time.Sleep(time.Duration(id) * time.Second)
	result := fmt.Sprintf("Goroutine %d a terminé son travail", id)
	results <- result
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 3)

	// Lancer 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg, results)
	}

	// Attendre que toutes les goroutines se terminent
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recevoir et afficher les résultats
	for result := range results {
		fmt.Println(result)
	}
}
