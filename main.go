package main

func main() {

	inv := loadInv()

	inv.editList("carrots", true)
	inv.editList("onions", true)
	inv.editList("green beans", true)
	inv.editList("milk", true)
	inv.editList("cream", true)

	groc := createList(*inv)

	inv.printInv()

	groc.Print()

}
