package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "flag"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type Origin int

const (
	GAME Origin = iota
	FRAME
	MODULE
	JTMODULE
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

var parsed []string

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

func append_filelist( dest *[]FileList, src []FileList, other *[]string, origin Origin ) {
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
			each = strings.TrimSpace(each)
			if strings.HasSuffix(each,".yaml") {
				var path string
				switch origin {
					case GAME: path = os.Getenv("CORES")+"/"+newfl.From+"/hdl/"
					case FRAME: path = os.Getenv("JTFRAME")+"/hdl/"+newfl.From+"/"
					default: path = os.Getenv("MODULES")+"/"
				}
				*other = append( *other, path+each )
			} else {
				newfl.Get = append(newfl.Get, each)
			}
		}
		if len(newfl.Get)>0 {
			found := false;
			for k,each:=range(*dest) {
				if each.From == newfl.From {
					(*dest)[k].Get = append((*dest)[k].Get, newfl.Get... )
					found = true
					break
				}
			}
			if !found {
				*dest = append( *dest, newfl )
			}
		}
	}
}

func parse_yaml( filename string, files *JTFiles ) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot open file %s",filename)
	}
	if parsed == nil {
		parsed = make( []string, 0 )
	}
	parsed = append( parsed, filename )
	var aux JTFiles
	err_yaml := yaml.Unmarshal( buf, &aux )
	if err_yaml != nil {
		//fmt.Println(err_yaml)
		log.Fatalf("jtfiles: cannot parse file\n\t%s\n\t%v", filename, err_yaml )
	}
	other := make( []string, 0 )
	// Parse
	append_filelist( &files.Game, aux.Game, &other, GAME )
	append_filelist( &files.JTFrame, aux.JTFrame, &other, FRAME )
	append_filelist( &files.Modules.Other, aux.Modules.Other, &other, MODULE )
	if files.Modules.JT==nil {
		files.Modules.JT = make( []JTModule, 0 )
	}
	for _,each := range(aux.Modules.JT) {
		files.Modules.JT = append( files.Modules.JT, each )
	}
	for _,each := range(other) {
		found := false
		for _,k := range(parsed) {
			if each == k {
				found = true
				break
			}
		}
		if !found {
			parse_yaml( each, files, )
		}
	}
}

func dump_filelist( fl []FileList, origin Origin ) {
	for _,each := range(fl) {
		var path string
		switch( origin ) {
			case GAME: path=os.Getenv("CORES")+"/"+each.From+"/hdl/"
			case FRAME: path=os.Getenv("JTFRAME")+"/hdl/"+each.From+"/"
			default: path=os.Getenv("JTROOT")+"/"
		}
		for _,each := range(each.Get) {
			if len(each)>0 {
				fmt.Println(path+each)
			}
		}
	}
}

func dump_qip( files JTFiles ) {
	dump_filelist( files.Game, GAME )
	dump_filelist( files.JTFrame, FRAME )
}

func main() {
	var args Args
	parse_args(&args)

	var files JTFiles
	parse_yaml( get_filename(args), &files )
	dump_qip(files)
}