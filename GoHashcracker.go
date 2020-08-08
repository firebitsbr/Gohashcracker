package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {

	c := color.New(color.FgCyan).Add(color.Underline)
	d := color.New(color.FgGreen).Add(color.Underline)
	e := color.New(color.FgRed).Add(color.Underline)

	fmt.Println("Welcome to the hashcracker - made by r1me")

	fmt.Print("Do you want to generate or brute force hash (g / b) : ")
	var choice string
	fmt.Scanln(&choice)

	if choice == "g" {
		fmt.Println()
		fmt.Println("1) Generate MD5-hash")
		fmt.Println("2) Generate SHA1-hash")
		fmt.Println("3) Generate SHA256-hash")
		fmt.Println()
		fmt.Print("Make a choice : ")
		var keuze int
		fmt.Scanln(&keuze)

		switch keuze {
		case 1:
			fmt.Print("Enter the password : ")
			var password string
			fmt.Scanln(&password)
			data := []byte(password)
			d.Printf("MD5 : %x", md5.Sum(data))

		case 2:
			fmt.Print("Enter the password : ")
			var password string
			fmt.Scanln(&password)

			data := []byte(password)
			d.Printf("SHA1 : %x", sha1.Sum(data))

		case 3:
			fmt.Print("Enter the password : ")
			var password string
			fmt.Scanln(&password)

			hash := sha256.New()
			hash.Write([]byte(password))
			d.Printf("SHA256 : %x", hash.Sum(nil))

		}

	} else if choice == "b" {
		fmt.Println()
		fmt.Println("1) Brute force MD5-hash")
		fmt.Println("2) Brute force SHA1-hash")
		fmt.Println("3) Brute force SHA256-hash")
		fmt.Println()
		fmt.Print("Make a choice : ")
		var keuze int
		fmt.Scanln(&keuze)

		switch keuze {
		case 1:
			fmt.Print("Enter MD5 hash : ")
			var password string
			fmt.Scanln(&password)

			fmt.Print("File name : ")
			var wordlist string
			fmt.Scanln(&wordlist)

			lines, err := scanLines(wordlist)
			if err != nil {
				panic(err)
			}

			for _, line := range lines {
				encWord := []byte(line)
				lol := md5.Sum((encWord))
				pass := hex.EncodeToString(lol[:])
				if pass == password {
					fmt.Println("Password found")
					c.Print("Password is --> ", line)
					break
				}
			}

		case 2:
			fmt.Print("Enter SHA1 hash : ")
			var password string
			fmt.Scanln(&password)

			fmt.Print("File name : ")
			var wordlist string
			fmt.Scanln(&wordlist)

			lines, err := scanLines(wordlist)
			if err != nil {
				panic(err)
			}

			for _, line := range lines {
				encWord := []byte(line)
				lol := sha1.Sum((encWord))
				pass := hex.EncodeToString(lol[:])
				if pass == password {
					fmt.Println("Password found")
					c.Print("Password is --> ", line)
					break
				}
			}

		case 3:
			fmt.Print("Enter SHA256 hash : ")
			var password string
			fmt.Scanln(&password)

			fmt.Print("File name : ")
			var wordlist string
			fmt.Scanln(&wordlist)

			lines, err := scanLines(wordlist)
			if err != nil {
				panic(err)
			}

			for _, line := range lines {
				hash := sha256.New()
				hash.Write([]byte(line))
				lol := hash.Sum((nil))
				pass := hex.EncodeToString(lol[:])
				if pass == password {

					fmt.Println("Password found")
					c.Print("Password is --> ", line)
					break
				}
			}

		}

	} else {
		e.Println("Faulse input, quitting...")
	}

}
func scanLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
