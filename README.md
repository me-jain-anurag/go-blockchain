# Go Blockchain
A minimal Proof-of-Work blockchain implementation in Go featuring persistent storage, mining, and a production-style CLI. The project focuses on correctness, modularity, and reproducible local execution.
## 📁 Project Structure
```text
cmd/
└── blockchain/
    └── main.go          # CLI entry point (flag parsing, command routing)
internal/
└── core/
    ├── block.go         # Block structure and serialization
    ├── blockchain.go   # Blockchain management and persistence
    ├── iterator.go     # Custom blockchain iterator
    ├── pow.go           # Proof-of-Work implementation
    └── wallet.go        # Wallet and cryptographic utilities
tmp/
└── blocks/              # BadgerDB persistence directory (auto-created at runtime, ignored in VCS)
```
## 🧠 Design Decisions
  - **Iterator Pattern:** Implemented a custom blockchain iterator to traverse the chain efficiently without loading the entire history into RAM.
  - **Persistence Strategy:** Chose [BadgerDB](https://github.com/dgraph-io/badger) for its high-performance, key-value nature which maps directly to block hash lookups, enabling direct block-hash lookups without relational overhead.
  - **Modular Proof-of-Work:** Decoupled the consensus algorithm from the block structure, allowing difficulty targets and hashing strategies to be modified independently.
  - **CLI Architecture:** Designed a command-line interface to separate the application entry point from the core logic, ensuring the system remains testable and extensible.
## 🚀 Features
  * **Core Blockchain:** Append-only chain where each block references the previous block hash.
  * **Proof of Work (PoW):** Hashcash-style mining algorithm with adjustable difficulty target.
  * **Persistence:** Blockchain state is stored permanently using BadgerDB.
  * **Command Line Interface (CLI):** Interactive tool to manage the blockchain (add blocks, inspect chain).
  * **Iterator:** Efficient blockchain traversal.
## 🛠️ Tech Stack
  * **Language:** [Go (Golang)](https://go.dev/)
  * **Database:** [BadgerDB](https://dgraph.io/docs/badger/) (Embedded Key-Value Store)
  * **Hashing:** SHA-256 / RIPEMD-160
  * **Cryptography:** Elliptic Curve (P-256)
## 📦 Installation & Usage
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/me-jain-anurag/go-blockchain.git
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
## 🚧 Limitations & Future Work
  - No peer-to-peer networking or consensus between nodes.
  - No transaction pool or validation (currently uses raw data strings).
  - Planned extensions include a REST API, basic transaction model (UTXO), and unit tests.
## 📚 References
  - **Tensor Programming:** [Golang Blockchain Tutorial](https://github.com/tensor-programming/golang-blockchain) (architecture inspiration)
  - [Go Documentation](https://pkg.go.dev/) (Standard Library)
  - [BadgerDB Documentation](https://dgraph-io.github.io/badger/)
## 📄 License
This project is open-source and available under the [MIT License](LICENSE).
