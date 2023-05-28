package main

import (
	"challenge-godb/entity"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "lelegit1109"
	dbname = "eniglaundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	showMainMenu(tx)
	
	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Data committed!")
	}
}


func showMainMenu(tx *sql.Tx) {
	for {
		fmt.Println("=== ENIGMA LAUNDRY ===")
		fmt.Println("1. Menu Master")
		fmt.Println("2. Menu Transaksi")
		fmt.Println("0. Keluar")

		var choice int
		fmt.Print("Pilihan: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			showMasterMenu(tx)
		case 2:
			showTransactionMenu(tx)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid!")
		}

		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}

}

func showMasterMenu(tx *sql.Tx) {
	for {
		fmt.Println("=== MENU MASTER ===")
		fmt.Println("1. Lihat Daftar Pelanggan")
		fmt.Println("2. Tambah Pelanggan Baru")
		fmt.Println("3. Hapus Pelanggan")
		fmt.Println("4. Update Pelanggan")
		fmt.Println("5. Daftar Pelayanan")
		fmt.Println("0. Kembali ke Menu Utama")

		var choice int
		fmt.Print("Pilihan: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1: 
			listCustomer(tx)
		case 2:
			addCustomer(tx)
		case 3:
			deleteCustomer(tx)
		case 4:
			updateCustomer(tx)
		case 5:
			listService(tx)
		case 0:
			showMainMenu(tx)
		default:
			fmt.Println("Pilihan tidak valid!")
		}
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}
}

func showTransactionMenu(tx *sql.Tx) {
	for {
		fmt.Println("=== MENU TRANSAKSI ===")
		fmt.Println("1. Lihat Daftar Transaksi")
		fmt.Println("2. Tambah Transaksi Baru")
		fmt.Println("0. Kembali ke Menu Utama")

		var choice int
		fmt.Print("Pilihan: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			listTransaction(tx)
		case 2:
			addTransaction(tx)
		case 0:
			showMainMenu(tx)
		default:
			fmt.Println("Pilihan tidak valid!")
		}
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}
}

func listService(tx *sql.Tx) {
	fmt.Println("=== DAFTAR PELAYANAN ===")
	listService := "SELECT service_id, service_name, service_price FROM Services"

	rows, err := tx.Query(listService)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var service entity.Service
		err := rows.Scan(&service.ID, &service.Name, &service.Price)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Nama: %s, No. HP: %.2f\n", service.ID, service.Name, service.Price)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	validate(err, "View", tx)

}

func listTransaction(tx *sql.Tx) {
	fmt.Println("=== DAFTAR TRANSAKSI ===")

	rows, err := tx.Query(`
		SELECT t.transaction_id, c.customer_name, s.service_name, t.transaction_date_in, t.transaction_date_done, t.received_by, t.quantity, t.unit, t.price, t.total_price
		FROM Transactions t
		INNER JOIN Customers c ON t.customer_id = c.customer_id
		INNER JOIN Services s ON t.service_id = s.service_id
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction entity.Transaction
		var customer entity.Customer
		var service entity.Service
		err := rows.Scan(&transaction.Transaction_Id, &customer.Name, &service.Name, &transaction.Transaction_In, &transaction.Transaction_Done, &transaction.Received_By, &transaction.Quantity, &transaction.Unit, &transaction.Price, &transaction.Total_Price)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Pelanggan: %s, Layanan: %s, Tanggal Masuk: %s, Tanggal Selesai: %s, Diterima Oleh: %s, Jumlah: %d %s, Harga: %.2f, Total Harga: %.2f\n",
		transaction.Transaction_Id, customer.Name, service.Name, transaction.Transaction_In, transaction.Transaction_Done, transaction.Received_By, transaction.Quantity, transaction.Unit, transaction.Price, transaction.Total_Price)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	validate(err, "View", tx)
}

func addTransaction(tx *sql.Tx) {
	
	var transaction entity.Transaction
	fmt.Println("=== TAMBAH TRANSAKSI BARU ===")

	fmt.Print("ID Pelanggan: ")
	fmt.Scanln(&transaction.Customer_Id)

	fmt.Print("ID Pelayanan: ")
	fmt.Scanln(&transaction.Service_Id)

	fmt.Print("Produk Masuk (YYYY-MM-DD): ")
	fmt.Scanln(&transaction.Transaction_In)

	fmt.Print("Tanggal Selesai (YYYY-MM-DD): ")
	fmt.Scanln(&transaction.Transaction_Done)

	fmt.Print("Diterima Oleh: ")
	fmt.Scanln(&transaction.Received_By)

	fmt.Print("Jumlah: ")
	fmt.Scanln(&transaction.Quantity)

	fmt.Print("Satuan: ")
	fmt.Scanln(&transaction.Unit)

	fmt.Print("Harga: ")
	fmt.Scanln(&transaction.Price)

	if transaction.Customer_Id <= 0 || transaction.Service_Id <= 0 || len(transaction.Transaction_In) != 10 || len(transaction.Transaction_Done) != 10 || len(transaction.Received_By) == 0 || transaction.Quantity <= 0 || len(transaction.Unit) == 0 || transaction.Price <= 0 {
		fmt.Println("Data transaksi tidak valid!")
		return
	}

	_, err := tx.Exec("INSERT INTO Transactions (customer_id, service_id, transaction_date_in, transaction_date_done, received_by, quantity, unit, price, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);", transaction.Customer_Id, transaction.Service_Id, transaction.Transaction_In, transaction.Transaction_Done, transaction.Received_By, transaction.Quantity, transaction.Unit, transaction.Price, transaction.Price*float64(transaction.Quantity))
	if err != nil {
		panic(err)
	}

	validate(err, "Add", tx)
}

func listCustomer(tx *sql.Tx) {
	fmt.Println("=== DAFTAR PELANGGAN ===")

	rows, err := tx.Query("SELECT customer_id, customer_name, phone_number FROM Customers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Phone)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %s, Nama: %s, No. HP: %s\n", customer.ID, customer.Name, customer.Phone)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	validate(err, "View", tx)
}

func addCustomer(tx *sql.Tx) {
	fmt.Println("=== TAMBAH PELANGGAN BARU ===")

	var customer entity.Customer

	fmt.Print("Nama Pelanggan: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("NO HP:")
	fmt.Scanln(&customer.Phone)

	if len(customer.Name) == 0 {
		fmt.Println("Nama harus diisi")
		return
	} else if len(customer.Phone) == 0 {
		fmt.Println("No HP harus diisi")
		return
	}

	_, err := tx.Exec("INSERT INTO Customers (customer_name, phone_number) VALUES ($1, $2);", customer.Name, customer.Phone)
	if err != nil {
		panic(err)
	}

	validate(err, "Add", tx)
}

func deleteCustomer(tx *sql.Tx) {
	fmt.Println("=== HAPUS PELANGGAN ===")

	var customer entity.Customer

	fmt.Print("ID Pelanggan: ")
	fmt.Scanln(&customer.ID)

	if len(customer.ID) == 0 {
		fmt.Println("Harap masukkan id nya")
		return
	}

	_, err := tx.Exec("DELETE FROM Customers WHERE customer_id = $1;", customer.ID)
	if err != nil {
		panic(err)
	}

	validate(err, "Delete", tx)
}

func updateCustomer(tx *sql.Tx) {
	fmt.Println("=== UPDATE DATA PELANGGAN ===")

	var customer entity.Customer

	fmt.Print("ID Pelanggan: ")
	fmt.Scanln(&customer.ID)
	fmt.Print("Nama Pelanggan: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("NO HP: ")
	fmt.Scanln(&customer.Phone)

	if len(customer.ID) == 0 || len(customer.Name) == 0 || len(customer.Phone) == 0 {
		fmt.Println("Data harus diisi!")
		return
	}

	_, err := tx.Exec("UPDATE Customers SET customer_name=$1, phone_number=$2 WHERE customer_id=$3;", customer.Name, customer.Phone, customer.ID)
	if err != nil {
		panic(err)
	}

	validate(err, "Update", tx)
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback")
	} else {
		fmt.Println("Succesfully " + message + " data!")
	}
}

func connectDb() *sql.DB{
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected")
	}

	return db
}