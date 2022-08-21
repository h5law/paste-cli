/*
Copyright © 2022 Harry Law <hrryslw@pm.me>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
*/
package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func CreatePaste() (map[string]string, error) {
	url := viper.Get("url")
	filePath := viper.GetString("file")
	fileType := viper.GetString("filetype")
	expiresIn := viper.GetInt("expiresIn")

	// Set input file depending to either os.Stdin or file flag
	var input *os.File
	if filePath == "" {
		input = os.Stdin
	} else {
		// Check file exists and open it
		exists, err := fileExists(filePath)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, fmt.Errorf("File not found: %s", filePath)
		}
		input, err = os.Open(filePath)
		if err != nil {
			return nil, err
		}
	}

	// Read lines into slice
	var content []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	input.Close()

	// Create request JSON body
	postBody, err := json.Marshal(map[string]interface{}{
		"content":   content,
		"filetype":  fileType,
		"expiresIn": expiresIn,
	})
	if err != nil {
		return nil, err
	}
	responseBody := bytes.NewBuffer(postBody)

	// Send post request and read body
	resp, err := http.Post(url.(string), "application/json", responseBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response into map to return
	m := make(map[string]string)
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}