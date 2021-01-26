package namecase_test

import (
	"strings"
	"testing"

	"github.com/StirlingMarketingGroup/go-namecase"
	. "github.com/stretchr/testify/assert"
)

var formatter = namecase.New()

var names = []string{
	// General
	"Keith", "Yuri's", "Leigh-Williams", "McCarthy",
	"O'Callaghan", "St. John", "von Streit",
	"van Dyke", "Van", "ap Llwyd Dafydd",
	"al Fahd", "Al",
	"el Grecco",
	"ben Gurion", "Ben",
	"da Vinci",
	"di Caprio", "du Pont", "de Legate",
	"del Crond", "der Sind", "van der Post", "van den Thillart",
	"ter Zanden", "ten Brink",
	"von Trapp", "la Poisson", "le Figaro",
	"Mack Knife", "Dougal MacDonald",
	"Yusof bin Ishak",
}

var macNames = []string{
	// Mac exceptions
	"Machin", "Machlin", "Machar",
	"Mackle", "Macklin", "Mackie",
	"Macquarie", "Machado", "Macevicius",
	"Maciulis", "Macias", "MacMurdo",
}

var romanNames = []string{
	// Roman numerals
	"Henry VIII", "Louis III", "Louis XIV",
	"Charles II", "Fred XLIX",
}

func TestNumbersAreUnaffected(t *testing.T) {
	properCased := "42"
	Equal(t, properCased, formatter.NameCase(strings.ToLower(properCased)))
}

func TestInternationalization(t *testing.T) {
	properCased := "Iñtërnâtiônàlizætiøn"
	Equal(t, properCased, formatter.NameCase(strings.ToLower(properCased)))
}

func TestStatic(t *testing.T) {
	for _, name := range names {
		Equal(t, name, formatter.NameCase(strings.ToLower(name)))
	}
}

func TestEmptyString(t *testing.T) {
	properCased := ""
	Equal(t, properCased, formatter.NameCase(strings.ToLower(properCased)))
}

func TestIrishExceptions(t *testing.T) {
	for _, name := range macNames {
		Equal(t, name, formatter.NameCase(strings.ToLower(name)))
	}
}

func TestRomanNumerals(t *testing.T) {
	for _, name := range romanNames {
		Equal(t, name, formatter.NameCase(strings.ToLower(name)))
	}
}

func TestInitials(t *testing.T) {
	properCased := "J. F. Kennedy"
	Equal(t, properCased, formatter.NameCase(strings.ToLower(properCased)))

	properCased = "J.F. Kennedy"
	Equal(t, properCased, formatter.NameCase(strings.ToLower(properCased)))
}

var programmers = []string{
	// General
	"Michael Abrash", "Scott Adams", "Leonard Adleman", "Alfred Aho", "J.J. Allaire",
	"Andrei Alexandrescu", "Paul Allen", "Eric Allman", "Marc Andreessen",
	"Bill Atkinson", "John Backus", "Richard Bartle", "Brian Behlendorf",
	"Kent Beck", "Donald Becker", "Doug Bell", "Fabrice Bellard",
	"Tim Berners-Lee", "Daniel J. Bernstein", "Sabeer Bhatia", "Eric Bina",
	"Marc Blank", "Joshua Bloch", "Grady Booch", "Bert Bos",
	"Stephen Richard Bourne", "David Bradley", "Andrew Braybrook", "Lawrence M. Breed",
	"Jack E. Bresenham", "Dan Bricklin", "Walter Bright", "Richard Brodie",
	"Andries Brouwer", "Danielle Bunten Berry", "Dries Buytaert", "Steve Capps",
	"John D. Carmack", "Vinton Cerf", "Ward Christensen", "Edgar F. Codd",
	"Bram Cohen", "Alain Colmerauer", "Alan Cooper", "Alan Cox",
	"Brad Cox", "Mike Cowlishaw", "Mark Crispin", "Ward Cunningham",
	"William Crowther", "Dave Cutler", "Ole-Johan Dahl", "James Duncan Davidson",
	"L. Peter Deutsch", "Edsger Dijkstra", "Matt Dillon", "Jack Dorsey",
	"Martin Dougiamas", "Adam Dunkels", "Les Earnest", "Brendan Eich",
	"Larry Ellison", "Marc Ewing", "Dan Farmer", "Stuart Feldman",
	"David Filo", "Brad Fitzpatrick", "Andrew Fluegelman", "Brian Fox",
	"Martin Fowler", "Jim Fruchterman", "Elon Gasper", "Bill Gates",
	"Steve Gibson", "John Gilmore", "Adele Goldberg", "Ryan C. Gordon",
	"James Gosling", "Bill Gosper", "Andrew Gower", "Paul Gower",
	"Paul Graham", "John Graham-Cumming", "Ralph Griswold", "Richard Greenblatt",
	"Scott Guthrie", "Andi Gutmans", "Jim Hall", "David Heinemeier Hansson",
	"David Albert Huffman", "Rebecca Heineman", "Anders Hejlsberg", "Ted Henter",
	"Andy Hertzfeld", "Rich Hickey", "D. Richard Hipp", "C. A. R. Hoare",
	"James Holmes", "Grace Hopper", "Dave Hyatt", "Miguel de Icaza",
	"Roberto Ierusalimschy", "Dan Ingalls", "Geir Ivarsøy",
	"Kenneth E. Iverson", "Toru Iwatani", "Bo Jangeborg", "Paul Jardetzky",
	"Lynne Jolitz", "William Jolitz", "Stephen C. Johnson", "Bill Joy",
	"Robert K. Jung", "Poul-Henning Kamp", "Mitch Kapor", "Phil Katz",
	"Alan Kay", "Mel Kaye", "John George Kemeny", "Stan Kelly-Bootle",
	"Brian Kernighan", "Gary Kildall", "Tom Knight", "Jim Knopf",
	"Donald E. Knuth", "Andrew Koenig",
	/* "Andre LaMothe", Combined La names are an issue */
	"Tom Lane", "Leslie Lamport", "Butler Lampson", "Sam Lantinga",
	"Richard H. Lathwell", "Chris Lattner", "Samuel J Leffler", "Rasmus Lerdorf",
	"Michael Lesk", "Gordon Letwin", "Rockford Lhotka", "Håkon Wium Lie", "Robert Love", "Ada Lovelace",
	/* "Al Lowe", Al is an issue */
	"Raphael Manfredi", "Khaled Mardam-Bey", "Yukihiro Matsumoto", "John McCarthy", "Craig McClanahan",
	"Daniel D. McCracken", "Douglas McIlroy", "Shawn McKenzie", "Marshall Kirk McKusick",
	"Bertrand Meyer", "Scott Meyers", "Bob Miner", "Jeff Minter",
	"Lou Montulli", "Bram Moolenaar", "David Moon", "Charles H. Moore",
	"Roger Moore", "Mike Muuss", "Patrick Naughton", "Peter Naur",
	"Fredrik Neij", "Graham Nelson", "Peter Norton", "Kristen Nygaard",
	/* "Ed Oates", Ed is an issue */
	"Martin Odersky", "Jarkko Oikarinen", "Oliver Twins",
	"John Ousterhout", "Onel de Guzman", "Larry Page", "Alexey Pajitnov",
	"Seymour Papert", "Tim Paterson", "Markus Persson", "Jeffrey Peterson",
	"Charles Petzold", "Rob Pike", "Kent Pitman", "Theo de Raadt",
	"Jef Raskin", "Eric S. Raymond", "Hans Reiser", "John Resig",
	"Craig Reynolds", "Dennis Ritchie", "Ron Rivest", "John Romero",
	"Blake Ross", "Guido van Rossum", "Jeff Rulifson", "Rusty Russell",
	"Steve Russell", "Mark Russinovich", "Bob Sabiston", "Muni Sakya",
	"Carl Sassenrath", "Chris Sawyer", "Bill Schelter", "Randal L. Schwartz",
	"Adi Shamir", "Mike Shaver", "Cliff Shaw", "Zed Shaw",
	"Emily Short", "Jacek Sieka", "Ken Silverman", "Charles Simonyi",
	"Colin Simpson", "Rich Skrenta", "Matthew Smith", "Henry Spencer",
	"Joel Spolsky", "Quentin Stafford-Fraser", "Richard Stallman", "Guy Steele",
	"Alexander Stepanov", "Ludvig Strigeus", "Bjarne Stroustrup", "Zeev Suraski",
	"Gerald Jay Sussman", "Herb Sutter", "Gottfrid Svartholm", "Tim Sweeney",
	"Andrew S. Tanenbaum", "Audrey Tang", "Simon Tatham", "Larry Tesler",
	"Jon Stephenson von Tetzchner", "Avie Tevanian", "Ken Thompson",
	"Michael Tiemann", "Linus Torvalds", "Andrew Tridgell", "Roy Trubshaw",
	"Bob Truel", "Wietse Venema", "Pat Villani", "Paul Vixie",
	"Patrick Volkerding", "Larry Wall", "Bob Wallace", "John Walker",
	"John Warnock", "Joseph Weizenbaum", "Robert Watson", "Pei-Yuan Wei",
	"Peter J. Weinberger", "Andrew Welch", "David Wheeler", "Arthur Whitney",
	"Bruce Wilcox", "Evan Williams", "Roberta Williams", "Sophie Wilson",
	"Dave Winer", "Niklaus Wirth", "Stephen Wolfram", "Don Woods",
	"Steve Wozniak", "Will Wright", "Jerry Yang", "Victor Yngve",
	"Jamie Zawinski", "Philip Zimmermann", "Mark Zuckerberg",
}

// Test http://en.wikipedia.org/wiki/List_of_programmers
func TestListOfProgrammers(t *testing.T) {
	for _, name := range programmers {
		Equal(t, name, formatter.NameCase(strings.ToLower(name)))
	}
}

func TestLowerCaseWords(t *testing.T) {
	properCased := "Prince Philip, Duke of Edinburgh"
	Equal(t, properCased, formatter.NameCase(strings.ToLower(properCased)))
}
