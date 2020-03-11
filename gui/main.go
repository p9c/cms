package main

import "C"
import (
	"bufio"
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/p9c/cms/connection"
	"github.com/p9c/cms/gui/js"
	"github.com/p9c/learngio/helpers"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
	Address = ":9999"

	addr = "localhost:4242"

	message = "foobar"
)

var (
	btn = new(widget.Button)
)

func main() {
	gofont.Register()
	th := material.NewTheme()
	js.BuildJS()

	var c net.Conn
	go func() {
		var err error

		fmt.Println("Starting client...")
		c, err = net.Dial("tcp", "localhost:19999")
		if err != nil {
			fmt.Println(err)
		}
		client := &connection.Client{Socket: c}
		go client.Receive()
		for {
			reader := bufio.NewReader(os.Stdin)
			message, _ := reader.ReadString('\n')
			c.Write([]byte(strings.TrimRight(message, "\n")))
		}
	}()
	//go connection.StartClientModule()

	http.Handle("/", http.FileServer(http.Dir("./html")))
	//go panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))

	log.Println("Listening on :3333...")
	go http.ListenAndServe(":3333", nil)
	//err := http.ListenAndServe(":3333", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	go func() {

		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Flexed(0.5, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ffcf30cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					}),
					layout.Flexed(0.25, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
						for btn.Clicked(gtx) {
							c.Write([]byte(strings.TrimRight(message, "\n")))

						}
						th.Button("Click me!").Layout(gtx, btn)
					}),
					layout.Flexed(0.25, func() {
						cs := gtx.Constraints
						helpers.DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					}),
				)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()

	//// map the routes
	//mapRoutes()
	// START OF WEB SERVER CODE
	log.Print("Starting server...")
	// make a http server using the goweb.DefaultHttpHandler()
	//s := &http.Server{
	//	Addr:           Address,
	//	Handler:        goweb.DefaultHttpHandler(),
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//listener, listenErr := net.Listen("tcp", Address)
	//log.Printf("  visit: %s", Address)
	//if listenErr != nil {
	//	log.Fatalf("Could not listen: %s", listenErr)
	//}
	//log.Println("")
	//log.Print("Some things to try in your browser:")
	//log.Printf("\t  http://localhost%s", Address)
	//log.Printf("\t  http://localhost%s/status-code/404", Address)
	//log.Printf("\t  http://localhost%s/people", Address)
	//log.Printf("\t  http://localhost%s/people/123", Address)
	//log.Printf("\t  http://localhost%s/people/anything", Address)
	//log.Printf("\t  http://localhost%s/people/me (will redirect)", Address)
	//log.Printf("\t  http://localhost%s/errortest", Address)
	//log.Printf("\t  http://localhost%s/things (try RESTful actions)", Address)
	//log.Printf("\t  http://localhost%s/123", Address)
	//log.Printf("\t  http://localhost%s/static/simple.html", Address)
	//log.Println("")
	//log.Println("Also try some of these routes:")
	//log.Printf("%s", goweb.DefaultHttpHandler())
	//go func() {
	//	for _ = range c {
	//		// sig is a ^C, handle it
	//		// stop the HTTP server
	//		log.Print("Stopping the server...")
	//		listener.Close()
	//		/*   Tidy up and tear down
	//		*/
	//		log.Print("Tearing down...")
	//		// TODO: tidy code up here
	//		log.Fatal("Finished - bye bye.  ;-)")
	//	}
	//}()
	// begin the server

	//	r := rts.Routes()

}
