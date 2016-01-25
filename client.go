// Copyright © 2016 Gavin Bong <raverunner@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//Package weisshorn provides functionality to parse & visualize
//CSVs from compasscard.ca
package weisshorn

import (
	//"errors"
	"fmt"
	tomb "gopkg.in/tomb.v2"
	"time"
)

func init() {
	fmt.Println("init inside client.go")
}

type Action interface {
	Filename() string
	Stop() error
}

type LineProducer struct {
	Action

	Sink      chan string
	Tombstone tomb.Tomb
}

func NewLineProducer(lines chan string) *LineProducer {
	lr := &LineProducer{
		Sink: lines,
	}
	lr.Tombstone.Go(lr.generate)
	return lr
}

func (c *LineProducer) Stop() error {
	c.Tombstone.Kill(nil)
	fmt.Println("sent killed")
	return c.Tombstone.Wait()
}

func (c *LineProducer) generate() error {
	a := []string{
		"Jan-20-2016 06:32 PM,Oakridge-41st Stn,Tap out at Oakridge-41st Stn,1 Zone Monthly Pass,2.75",
		"Jan-19-2016 01:40 PM,Waterfront Stn,Tap in at Waterfront Stn,1 Zone Monthly Pass,-2.75"}

	fmt.Println("CsvImporter :: Running")

	for _, line := range a {
		fmt.Println(line)
		time.Sleep(800 * time.Millisecond)
		select {
		case c.Sink <- line:
		case <-c.Tombstone.Dying():
			return nil
		}
	}

	//fmt.Println("CsvImporter :: Dying")
	close(c.Sink)
	return nil
}
