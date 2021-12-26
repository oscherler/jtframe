package main

import (
    "fmt"
    "log"
    "os"
    "flag"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type FileList struct {
	From string `yaml:"from"`
	Get []string `yaml:"get"`
}

type JTModule struct {
	Name string `yaml:"name"`
}

type JTFiles struct {
    Game [] FileList `yaml:"game"`
    JTFrame [] FileList `yaml:"jtframe"`
    Modules struct {
    	JT [] JTModule `yaml:"jt"`
	    Other [] FileList `yaml:"other"`
    } `yaml:"modules"`
}

type Args struct {
	Corename string
}

func parse_args( args *Args ) {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s, part of JTFRAME. (c) Jose Tejada 2021.\nUsage:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.StringVar(&args.Corename,"core","","core name")
	flag.Parse()
	if len(args.Corename)==0 {
		log.Fatal("JTFILES: You must specify the core name with argument -core")
	}
}

func get_filename( args Args ) string {
	cores := os.Getenv("CORES")
	if len(cores)==0 {
		log.Fatal("JTFILES: environment variable CORES is not defined")
	}
	fname := cores + "/" + args.Corename + "/hdl/game.yaml"
	return fname
}

func append_filelist( dest *[]FileList, src []FileList ) {
	if src == nil {
		return
	}
	if dest == nil {
		*dest = make( []FileList, 0 )
	}
	for _,each := range(src) {
		//fmt.Println(each)
		var newfl FileList
		newfl.From = each.From
		newfl.Get = make([]string,2)
		for _,each := range(each.Get) {
			newfl.Get = append(newfl.Get, each)
		}
		*dest = append( *dest, newfl )
	}
}

func parse_yaml( filename string, files *JTFiles, root bool ) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot open file %s",filename)
	}
	var aux JTFiles
	err_yaml := yaml.Unmarshal( buf, &aux )
	if err_yaml != nil {
		//fmt.Println(err_yaml)
		log.Fatalf("jtfiles: cannot parse file\n\t%s\n\t%v", filename, err_yaml )
	}
	// Parse
	if( root ) {
		append_filelist( &files.Game, aux.Game )
	}
	append_filelist( &files.JTFrame, aux.JTFrame )
	append_filelist( &files.Modules.Other, aux.Modules.Other )
	if files.Modules.JT==nil {
		files.Modules.JT = make( []JTModule, 0 )
	}
	for _,each := range(aux.Modules.JT) {
		files.Modules.JT = append( files.Modules.JT, each )
	}
}

func main() {
	var args Args
	parse_args(&args)

	var files JTFiles
	parse_yaml( get_filename(args), &files, true )
	fmt.Println(files)
}