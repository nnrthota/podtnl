/*
Copyright [2020] [Narendranath Reddy]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
package main
*/

package types

// TunnelStatus that ngrok support
type TunnelStatus uint

// Protocol protocol type
type Protocol uint

// IsCreated tunnel status
type IsCreated bool

// Protocol that ngrok support
const (
	HTTP Protocol = 0 + iota
	TCP
	TLS
)

// Protocols that ngrok support
var Protocols = [...]string{
	"http",
	"tcp",
	"tls",
}

// TunnelStatus that ngrok support
const (
	Ready TunnelStatus = 0 + iota
	Pending
	Failed
	Deleted
)

var status = [...]string{
	"READY",
	"PENDING_CREATE",
	"CREATE_FAILED",
	"DELETED",
}

//ToString convert to string
func (c TunnelStatus) ToString() string {

	if c != Ready && c != Pending && c != Failed && c != Deleted {
		return ""
	}

	return status[c]
}

//ToString convert to string
func (p Protocol) ToString() string {

	if p != HTTP && p != TCP && p != TLS {
		return ""
	}

	return Protocols[p]
}
