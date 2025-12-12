# Go Blockchain

A simplified implementation of a Proof-of-Work Blockchain written in Go. This project was built to understand the fundamental concepts of distributed ledger technology, including cryptographic hashing, mining, persistence, and command-line interfaces.

## 🚀 Features

* **Core Blockchain:** Linked list structure where blocks are cryptographically linked via SHA-256 hashes.
* **Proof of Work (PoW):** Hashcash-style mining algorithm with adjustable difficulty target.
* **Persistence:** Blockchain state is stored permanently using [BadgerDB](https://github.com/dgraph-io/badger), a fast key-value store.
* **Command Line Interface (CLI):** Easy-to-use tools to interact with the blockchain (add blocks, inspect chain).
* **Iterator:** Efficient blockchain traversal allows for parsing large chains without loading the entire history into RAM.

## 🛠️ Tech Stack

* **Language:** [Go (Golang)](https://go.dev/)
* **Database:** [BadgerDB](https://dgraph.io/docs/badger/) (Embedded Key-Value Store)
* **Hashing:** SHA-256

## 📦 Installation & Usage

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/me-jain-anurag/go-blockchain.git](https://github.com/me-jain-anurag/go-blockchain.git)
    cd go-blockchain
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Build the CLI tool:**
    ```bash
    go build -o blockchain ./cmd/blockchain
    ```

4.  **Run Commands:**
    * **Add a new block:**
        ```bash
        ./blockchain addblock -data "Send 50 BTC to Alice"
        ```
    * **Print the entire chain:**
        ```bash
        ./blockchain printchain
        ```

    *(Note: On Windows, use `blockchain.exe` instead of `./blockchain`)*

## 📚 Credits & Acknowledgements

This project was built for educational purposes with the help of several key resources:

* **Tensor Programming:** The core architecture and logic follow the excellent [Golang Blockchain Tutorial](https://github.com/tensor-programming/golang-blockchain) by Tensor Programming. This is a highly recommended resource for anyone learning Go.
* **Gemini:** Development was supported by Google's **Gemini AI**, which served as a pair programmer, explaining concepts, debugging code, and helping structure the project professionally.
* **Documentation:**
    * [Go Documentation](https://pkg.go.dev/) - For standard library references (`crypto/sha256`, `encoding/binary`, `flag`).
    * [BadgerDB Documentation](https://dgraph-io.github.io/badger/) - For understanding the embedded database integration.

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).