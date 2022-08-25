This is the code repository for my article, **[Create a Blockchain Application in Go: A Beginner's Guide](https://narasimmantech.com/create-a-blockchain-application-in-go-a-beginners-guide)** on my website - **[Narasimman Tech](https://narasimmantech.com)**.

### **To run the program:**

"ADDR" should be set as the server listen port in the .env file. Example: ADDR=8080

```
go run main.go
```

This program will start by listening on port 8080 and then creating the genesis block.

### **To add a new block:**

```
curl -X POST -H "Content-Type: application/json" -d '{"BPM":10}' http://localhost:8080
```

**Output:**

```
{
  "Index": 1,
  "Timestamp": "2022-08-25 21:08:26.13258 +0530 IST m=+1950.938579001",
  "BPM": 10,
  "Hash": "a6f1a856b1a8112cc0bfb45ce24930408ac3cf0c4da18fc18839aa9c28f7cb31",
  "PrevHash": "b1ae03473636acf0969718c00f417aacad6b0d9ec19f21a8bc19b75bafc3a666"
}
```

This will add a new block to the blockchain with the data BPM = 10.

### **Output Explained**

The blockchain is a slice of blocks. Each block contains the following information:

**Index:** The position of the block in the blockchain.

**Timestamp:** The time when the block was created.

**BPM:** The data contained in the block.

**Hash:** The hash of the block.

**PrevHash:** The hash of the previous block in the blockchain.
