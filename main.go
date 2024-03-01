package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	Code      string
	Title     string
	Author    string
	Publisher string
	PageCount int
	Year      int
}

var library []Book

func main() {
	for {
		fmt.Println("\n1. Tambah Buku")
		fmt.Println("2. Tampilkan Daftar Buku")
		fmt.Println("3. Hapus Buku")
		fmt.Println("4. Ubah Buku")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBook()
		case 2:
			displayBooks()
		case 3:
			deleteBook()
		case 4:
			editBook()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func addBook() {
	var book Book
	fmt.Print("Kode Buku: ")
	fmt.Scanln(&book.Code)
	fmt.Print("Judul Buku: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		book.Title = scanner.Text()
	}
	fmt.Print("Pengarang: ")
	if scanner.Scan() {
		book.Author = scanner.Text()
	}
	fmt.Print("Penerbit: ")
	if scanner.Scan() {
		book.Publisher = scanner.Text()
	}
	fmt.Print("Jumlah Halaman: ")
	fmt.Scanln(&book.PageCount)
	fmt.Print("Tahun Terbit: ")
	fmt.Scanln(&book.Year)

	library = append(library, book)
	fmt.Println("Buku berhasil ditambahkan")
}

func displayBooks() {
	fmt.Println("\nDaftar Buku:")
	for _, book := range library {
		fmt.Printf("Kode: %s, Judul: %s, Pengarang: %s, Penerbit: %s, Halaman: %d, Tahun: %d\n",
			book.Code, book.Title, book.Author, book.Publisher, book.PageCount, book.Year)
	}
}

func deleteBook() {
	var code string
	fmt.Print("Masukkan Kode Buku yang akan dihapus: ")
	fmt.Scanln(&code)

	for i, book := range library {
		if book.Code == code {
			library = append(library[:i], library[i+1:]...)
			fmt.Println("Buku berhasil dihapus")
			return
		}
	}
	fmt.Println("Buku tidak ditemukan")
}

func editBook() {
	var code string
	fmt.Print("Masukkan Kode Buku yang akan diubah: ")
	fmt.Scanln(&code)

	for i, book := range library {
		if book.Code == code {
			fmt.Print("Judul Buku Baru (Kosongkan jika tidak ingin mengubah): ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() && scanner.Text() != "" {
				library[i].Title = scanner.Text()
			}
			fmt.Print("Pengarang Baru (Kosongkan jika tidak ingin mengubah): ")
			if scanner.Scan() && scanner.Text() != "" {
				library[i].Author = scanner.Text()
			}
			fmt.Print("Penerbit Baru (Kosongkan jika tidak ingin mengubah): ")
			if scanner.Scan() && scanner.Text() != "" {
				library[i].Publisher = scanner.Text()
			}
			fmt.Print("Jumlah Halaman Baru (Kosongkan jika tidak ingin mengubah): ")
			var pageCountStr string
			if scanner.Scan() && scanner.Text() != "" {
				pageCountStr = scanner.Text()
				pageCount, err := strconv.Atoi(pageCountStr)
				if err == nil {
					library[i].PageCount = pageCount
				} else {
					fmt.Println("Input tidak valid")
					return
				}
			}
			fmt.Print("Tahun Terbit Baru (Kosongkan jika tidak ingin mengubah): ")
			var yearStr string
			if scanner.Scan() && scanner.Text() != "" {
				yearStr = scanner.Text()
				year, err := strconv.Atoi(yearStr)
				if err == nil {
					library[i].Year = year
				} else {
					fmt.Println("Input tidak valid")
					return
				}
			}
			fmt.Println("Buku berhasil diubah")
			return
		}
	}
	fmt.Println("Buku tidak ditemukan")
}
