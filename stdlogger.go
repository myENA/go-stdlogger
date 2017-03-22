package stdlogger

type StdLogger interface {
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})

	Fatalf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalln(v ...interface{})

	Panicf(format string, v ...interface{})
	Panic(v ...interface{})
	Panicln(v ...interface{})
}
