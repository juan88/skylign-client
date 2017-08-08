package main

import (
	//"github.com/juan88/skylign-client/client"
	"flag"
	"fmt"
	"os"
)

func main() {

	//url := "http://skylign.org/"
	
	filenamePtr := flag.String("file", "", "The filename to send to the Skyling server. Formats: Aligned FASTA, Clustal (and Clustal-like), PSI-BLAST, PHYLIP, Selex, GCG/MSF, STOCKHOLM format, UC Santa Cruz A2M (alignment to model)")
	heightPtr := flag.String("height", "info_content_all", "Letter height server param. Available options are: info_content_all (default), info_content_above, score")
	processingPtr := flag.String("processing", "info_content_all", "Processing mode. Available options are: observed (default), weighted, hmm, hmm_all")
	fragPtr := flag.String("frag", "full", "Fragment handling. Available options: full, frag")

	flag.Parse()
	fmt.Println("processing: ", *processingPtr)
	fmt.Println("height: ", *heightPtr)
	fmt.Println("frag: ", *fragPtr)
	fmt.Println("filename: ", *filenamePtr)

	
	if(len(*filenamePtr) == 0) {
		fmt.Println("Error! File is requiered")
		os.Exit(1)
	}
	

	/*
	var response client.UploadedAlignFileResponse;
	params := map[string]string{
		"processing": "observed",
		"letter_height": "info_content_all",
		"frag":	"full",
	}
	*/
	
	//client.UploadData(url, "ADIN0.sto", params, &response)
}