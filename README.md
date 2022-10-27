## AB Invoicer

The purpose of this tool is mainly for learning purposes. I an trying to learn the Go programming language, and decided to work on a project instead of going through tutorial hell.

This is a CLI program that will display a form to fill out some data, ie user data, expenses, and labour to render a PDF document of the invoice.

* It uses the Go Template html files to render an html page of the invoice
* It uses [go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf) to convert the Go Template html file into a PDF
* It uses [Charmbracelet's BubbleTea](https://github.com/charmbracelet/bubbletea) to render the CLI UI along with other libraries from Charmbracelet to style it up

In order to get this to run on your computer, you will need to install [wkhtmltopdf](https://wkhtmltopdf.org/)



# Helpful links
https://www.reddit.com/r/golang/comments/3gtg3i/passing_slice_of_values_as_slice_of_interfaces/
https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces/12754757#12754757
https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
