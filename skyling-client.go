package main

import (
	"github.com/juan88/skylign-client/client"
	"flag"
	"fmt"
	"os"
	"github.com/juan88/go-rainbow"
)

func main() {
	filenamePtr := flag.String("file", "", "The filename to send to the Skyling server. Formats: Aligned FASTA, Clustal (and Clustal-like), PSI-BLAST, PHYLIP, Selex, GCG/MSF, STOCKHOLM format, UC Santa Cruz A2M (alignment to model)")
	heightPtr := flag.String("height", "info_content_all", "Letter height server param. Available options are: info_content_all (default), info_content_above, score")
	processingPtr := flag.String("processing", "info_content_all", "Processing mode. Available options are: observed (default), weighted, hmm, hmm_all")
	fragPtr := flag.String("frag", "full", "Fragment handling. Available options: full, frag")
	flag.Parse()
	
	if(len(*filenamePtr) == 0) {
		fmt.Println("Error! File is requiered")
		os.Exit(1)
	}

	params := map[string]string{
		"processing": *processingPtr,
		"letter_height": *heightPtr,
		"frag":	*fragPtr,
	}
	
	client.GenerateLogo(*filenamePtr, params)
	fmt.Println(rainbow.Green("Logo generated successfully!"))
}