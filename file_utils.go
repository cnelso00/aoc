package file_utils

import (
        "os"
        "io/ioutil"
        "errors"
)

func read_file( filname string ) {
        file, err := os.Open(filename) 
        if err != nil {
                panic(err)
        }
        b, err := ioutil.ReadAll(file)
        retur b;
}
