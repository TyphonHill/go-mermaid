package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	erd "github.com/TyphonHill/go-mermaid/diagrams/entityrelationship"
)

func main() {
	diagram := erd.NewDiagram()
	diagram.EnableMarkdownFence()
	diagram.SetTitle("Library Management System ERD")

	// Create Member entity
	member := diagram.AddEntity("MEMBERS").
		SetAlias("Member")
	member.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	member.AddAttribute("first_name", erd.TypeString).SetRequired()
	member.AddAttribute("last_name", erd.TypeString).SetRequired()
	member.AddAttribute("email", erd.TypeString).SetRequired()
	member.AddAttribute("phone", erd.TypeString)
	member.AddAttribute("join_date", erd.TypeDateTime).SetRequired()
	member.AddAttribute("status", erd.TypeString).SetRequired()

	// Create Book entity
	book := diagram.AddEntity("BOOKS").
		SetAlias("Book")
	book.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	book.AddAttribute("isbn", erd.TypeString).SetRequired()
	book.AddAttribute("title", erd.TypeString).SetRequired()
	book.AddAttribute("publication_year", erd.TypeInteger)
	book.AddAttribute("available_copies", erd.TypeInteger).SetRequired()
	book.AddAttribute("total_copies", erd.TypeInteger).SetRequired()

	// Create Author entity
	author := diagram.AddEntity("AUTHORS").
		SetAlias("Author")
	author.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	author.AddAttribute("name", erd.TypeString).SetRequired()
	author.AddAttribute("biography", erd.TypeString)
	author.AddAttribute("birth_date", erd.TypeDateTime)

	// Create Category entity
	category := diagram.AddEntity("CATEGORIES").
		SetAlias("Category")
	category.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	category.AddAttribute("name", erd.TypeString).SetRequired()
	category.AddAttribute("description", erd.TypeString)

	// Create Loan entity
	loan := diagram.AddEntity("LOANS").
		SetAlias("Loan")
	loan.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	loan.AddAttribute("member_id", erd.TypeInteger).SetForeignKey()
	loan.AddAttribute("book_id", erd.TypeInteger).SetForeignKey()
	loan.AddAttribute("checkout_date", erd.TypeDateTime).SetRequired()
	loan.AddAttribute("due_date", erd.TypeDateTime).SetRequired()
	loan.AddAttribute("return_date", erd.TypeDateTime)
	loan.AddAttribute("status", erd.TypeString).SetRequired()

	// Create Fine entity
	fine := diagram.AddEntity("FINES").
		SetAlias("Fine")
	fine.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	fine.AddAttribute("amount", erd.TypeFloat).SetRequired()
	fine.AddAttribute("issue_date", erd.TypeDateTime).SetRequired()
	fine.AddAttribute("paid_date", erd.TypeDateTime)
	fine.AddAttribute("status", erd.TypeString).SetRequired()

	// Create Reservation entity
	reservation := diagram.AddEntity("RESERVATIONS").
		SetAlias("Reservation")
	reservation.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	reservation.AddAttribute("reservation_date", erd.TypeDateTime).SetRequired()
	reservation.AddAttribute("expiry_date", erd.TypeDateTime).SetRequired()
	reservation.AddAttribute("status", erd.TypeString).SetRequired()

	// Create BookAuthor join table
	bookAuthor := diagram.AddEntity("BOOK_AUTHORS").
		SetAlias("BookAuthor")
	bookAuthor.AddAttribute("book_id", erd.TypeInteger).SetPrimaryKey().SetForeignKey()
	bookAuthor.AddAttribute("author_id", erd.TypeInteger).SetPrimaryKey().SetForeignKey()

	// Similarly for book categories
	bookCategory := diagram.AddEntity("BOOK_CATEGORIES").
		SetAlias("BookCategory")
	bookCategory.AddAttribute("book_id", erd.TypeInteger).SetPrimaryKey().SetForeignKey()
	bookCategory.AddAttribute("category_id", erd.TypeInteger).SetPrimaryKey().SetForeignKey()

	// Define relationships
	diagram.AddRelationship(member, loan).
		SetLabel("makes").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(book, loan).
		SetLabel("involved_in").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(loan, fine).
		SetLabel("may_incur").
		SetCardinality(erd.OneToExactlyOne)

	diagram.AddRelationship(member, reservation).
		SetLabel("places").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(book, reservation).
		SetLabel("has").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(book, bookAuthor).
		SetLabel("has").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(author, bookAuthor).
		SetLabel("writes").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(book, bookCategory).
		SetLabel("has").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(category, bookCategory).
		SetLabel("categorizes").
		SetCardinality(erd.OneToZeroOrMore)

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
