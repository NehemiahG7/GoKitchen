package main

type Inventory struct{
	func getInv() {
		fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inventory:")
		scanner := bufio.NewScanner(os.Stdin)
		fInv, err := os.Create("fInv.csv")
		if err != nil {
			log.Fatalf("File failed to create %s", err)
		}
}