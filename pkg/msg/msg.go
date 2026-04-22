// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import (
	"net"
	"reflect"
)

const (
	TypeLogin              = 'o'
	TypeLoginResp          = '1'
	TypeNewProxy           = 'p'
	TypeNewProxyResp       = '2'
	TypeCloseProxy         = 'c'
	TypeNewWorkConn        = 'w'
	TypeReqWorkConn        = 'r'
	TypeStartWorkConn      = 's'
	TypeNewVisitorConn     = 'v'
	TypeNewVisitorConnResp = '3'
	TypePing               = 'h'
	TypePong               = '4'
	TypeUDPPacket          = 'u'
	TypeNatHoleVisitor     = 'i'
	TypeNatHoleClient      = 'n'
	TypeNatHoleResp        = 'm'
	TypeNatHoleSid         = '5'
	TypeNatHoleReport      = '6'
)

var msgTypeMap = map[byte]any{
	TypeLogin:              Login{},
	TypeLoginResp:          LoginResp{},
	TypeNewProxy:           NewProxy{},
	TypeNewProxyResp:       NewProxyResp{},
	TypeCloseProxy:         CloseProxy{},
	TypeNewWorkConn:        NewWorkConn{},
	TypeReqWorkConn:        ReqWorkConn{},
	TypeStartWorkConn:      StartWorkConn{},
	TypeNewVisitorConn:     NewVisitorConn{},
	TypeNewVisitorConnResp: NewVisitorConnResp{},
	TypePing:               Ping{},
	TypePong:               Pong{},
	TypeUDPPacket:          UDPPacket{},
	TypeNatHoleVisitor:     NatHoleVisitor{},
	TypeNatHoleClient:      NatHoleClient{},
	TypeNatHoleResp:        NatHoleResp{},
	TypeNatHoleSid:         NatHoleSid{},
	TypeNatHoleReport:      NatHoleReport{},
}

var TypeNameNatHoleResp = reflect.TypeOf(&NatHoleResp{}).Elem().Name()

type ClientSpec struct {
	// Due to the support of VirtualClient, frps needs to know the client type in order to
	// differentiate the processing logic.
	// Optional values: ssh-tunnel
	Type string `json:"m2k9s,omitempty"`
	// If the value is true, the client will not require authentication.
	AlwaysAuthPass bool `json:"q7v1h,omitempty"`
}

// When frpc start, client send this message to login to server.
type Login struct {
	Version      string            `json:"r3c8d,omitempty"`
	Hostname     string            `json:"t6n0x,omitempty"`
	Os           string            `json:"u1p4j,omitempty"`
	Arch         string            `json:"v8e2b,omitempty"`
	User         string            `json:"w5g7l,omitempty"`
	PrivilegeKey string            `json:"x0d9q,omitempty"`
	Timestamp    int64             `json:"y4s1m,omitempty"`
	RunID        string            `json:"z7h3t,omitempty"`
	Metas        map[string]string `json:"a2n8v,omitempty"`

	// Currently only effective for VirtualClient.
	ClientSpec ClientSpec `json:"b6q0c,omitempty"`

	// Some global configures.
	PoolCount int `json:"c9l2f,omitempty"`
}

type LoginResp struct {
	Version string `json:"r3c8d,omitempty"`
	RunID   string `json:"z7h3t,omitempty"`
	Error   string `json:"aa2u6r,omitempty"`
}

// When frpc login success, send this message to frps for running a new proxy.
type NewProxy struct {
	ProxyName          string            `json:"d3p7k,omitempty"`
	ProxyType          string            `json:"e8u1r,omitempty"`
	UseEncryption      bool              `json:"f2m6s,omitempty"`
	UseCompression     bool              `json:"g7t0x,omitempty"`
	BandwidthLimit     string            `json:"h1v5j,omitempty"`
	BandwidthLimitMode string            `json:"i6c9q,omitempty"`
	Group              string            `json:"j0d4m,omitempty"`
	GroupKey           string            `json:"k4s8t,omitempty"`
	Metas              map[string]string `json:"a2n8v,omitempty"`
	Annotations        map[string]string `json:"l9h2v,omitempty"`

	// tcp and udp only
	RemotePort int `json:"m5q7c,omitempty"`

	// http and https only
	CustomDomains     []string          `json:"n0l3f,omitempty"`
	SubDomain         string            `json:"o7p1k,omitempty"`
	Locations         []string          `json:"p2u6r,omitempty"`
	HTTPUser          string            `json:"q8m0s,omitempty"`
	HTTPPwd           string            `json:"r3t7x,omitempty"`
	HostHeaderRewrite string            `json:"s1v4j,omitempty"`
	Headers           map[string]string `json:"t6c9q,omitempty"`
	ResponseHeaders   map[string]string `json:"u0d2m,omitempty"`
	RouteByHTTPUser   string            `json:"v4s8t,omitempty"`

	// stcp, sudp, xtcp
	Sk         string   `json:"w9h1v,omitempty"`
	AllowUsers []string `json:"x5q7c,omitempty"`

	// tcpmux
	Multiplexer string `json:"y0l3f,omitempty"`
}

type NewProxyResp struct {
	ProxyName  string `json:"d3p7k,omitempty"`
	RemoteAddr string `json:"z7p1k,omitempty"`
	Error      string `json:"aa2u6r,omitempty"`
}

type CloseProxy struct {
	ProxyName string `json:"d3p7k,omitempty"`
}

type NewWorkConn struct {
	RunID        string `json:"z7h3t,omitempty"`
	PrivilegeKey string `json:"x0d9q,omitempty"`
	Timestamp    int64  `json:"y4s1m,omitempty"`
}

type ReqWorkConn struct{}

type StartWorkConn struct {
	ProxyName string `json:"d3p7k,omitempty"`
	SrcAddr   string `json:"ab8m0s,omitempty"`
	DstAddr   string `json:"ac3t7x,omitempty"`
	SrcPort   uint16 `json:"ad1v4j,omitempty"`
	DstPort   uint16 `json:"ae6c9q,omitempty"`
	Error     string `json:"aa2u6r,omitempty"`
}

type NewVisitorConn struct {
	RunID          string `json:"z7h3t,omitempty"`
	ProxyName      string `json:"d3p7k,omitempty"`
	SignKey        string `json:"af0d2m,omitempty"`
	Timestamp      int64  `json:"y4s1m,omitempty"`
	UseEncryption  bool   `json:"f2m6s,omitempty"`
	UseCompression bool   `json:"g7t0x,omitempty"`
}

type NewVisitorConnResp struct {
	ProxyName string `json:"d3p7k,omitempty"`
	Error     string `json:"aa2u6r,omitempty"`
}

type Ping struct {
	PrivilegeKey string `json:"x0d9q,omitempty"`
	Timestamp    int64  `json:"y4s1m,omitempty"`
}

type Pong struct {
	Error string `json:"aa2u6r,omitempty"`
}

type UDPPacket struct {
	Content    string       `json:"ag4s8t,omitempty"`
	LocalAddr  *net.UDPAddr `json:"ah9h1v,omitempty"`
	RemoteAddr *net.UDPAddr `json:"ai5q7c,omitempty"`
}

type NatHoleVisitor struct {
	TransactionID string   `json:"aj0l3f,omitempty"`
	ProxyName     string   `json:"d3p7k,omitempty"`
	PreCheck      bool     `json:"ak7p1k,omitempty"`
	Protocol      string   `json:"al2u6r,omitempty"`
	SignKey       string   `json:"af0d2m,omitempty"`
	Timestamp     int64    `json:"y4s1m,omitempty"`
	MappedAddrs   []string `json:"am8m0s,omitempty"`
	AssistedAddrs []string `json:"an3t7x,omitempty"`
}

type NatHoleClient struct {
	TransactionID string   `json:"aj0l3f,omitempty"`
	ProxyName     string   `json:"d3p7k,omitempty"`
	Sid           string   `json:"ao1v4j,omitempty"`
	MappedAddrs   []string `json:"am8m0s,omitempty"`
	AssistedAddrs []string `json:"an3t7x,omitempty"`
}

type PortsRange struct {
	From int `json:"ap6c9q,omitempty"`
	To   int `json:"aq0d2m,omitempty"`
}

type NatHoleDetectBehavior struct {
	Role              string       `json:"ar4s8t,omitempty"` // sender or receiver
	Mode              int          `json:"as9h1v,omitempty"` // 0, 1, 2...
	TTL               int          `json:"at5q7c,omitempty"`
	SendDelayMs       int          `json:"au0l3f,omitempty"`
	ReadTimeoutMs     int          `json:"av7p1k,omitempty"`
	CandidatePorts    []PortsRange `json:"aw2u6r,omitempty"`
	SendRandomPorts   int          `json:"ax8m0s,omitempty"`
	ListenRandomPorts int          `json:"ay3t7x,omitempty"`
}

type NatHoleResp struct {
	TransactionID  string                `json:"aj0l3f,omitempty"`
	Sid            string                `json:"ao1v4j,omitempty"`
	Protocol       string                `json:"al2u6r,omitempty"`
	CandidateAddrs []string              `json:"az1v4j,omitempty"`
	AssistedAddrs  []string              `json:"an3t7x,omitempty"`
	DetectBehavior NatHoleDetectBehavior `json:"ba6c9q,omitempty"`
	Error          string                `json:"aa2u6r,omitempty"`
}

type NatHoleSid struct {
	TransactionID string `json:"aj0l3f,omitempty"`
	Sid           string `json:"ao1v4j,omitempty"`
	Response      bool   `json:"bb0d2m,omitempty"`
	Nonce         string `json:"bc4s8t,omitempty"`
}

type NatHoleReport struct {
	Sid     string `json:"ao1v4j,omitempty"`
	Success bool   `json:"bd9h1v,omitempty"`
}
