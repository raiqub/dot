/*
 * Copyright 2015 Fabrício Godoy
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dot

import (
	"net"
	"testing"
	"time"
)

const (
	ListenerLaunchDelay = time.Millisecond * 10
	ListenerServerAddr  = ":64080"
	ListenerServerNet   = "tcp"
	WaitTimeout         = time.Millisecond * 250
)

func TestWaitPeer(t *testing.T) {
	go func() {
		time.Sleep(ListenerLaunchDelay)

		l, err := net.Listen(ListenerServerNet, ListenerServerAddr)
		if err != nil {
			t.Skip("Failed to listen tcp/64080")
		}
		defer l.Close()

		conn, err := l.Accept()
		if err != nil {
			t.Skip("Failed to accept connection")
		}
		defer conn.Close()
	}()

	if !WaitPeerListening(
		ListenerServerNet,
		ListenerServerAddr,
		WaitTimeout,
	) {
		t.Fatal("Could not wait to peer to be ready")
	}

	if WaitPeerListening(
		ListenerServerNet,
		ListenerServerAddr,
		WaitTimeout,
	) {
		t.Fatal("The peer should not be ready")
	}
}
