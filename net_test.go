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
	LISTENER_LAUNCH_DELAY = time.Millisecond * 10
	LISTENER_SERVER_ADDR  = ":64080"
	LISTENER_SERVER_NET   = "tcp"
	WAIT_TIMEOUT          = time.Millisecond * 250
)

func TestWaitPeer(t *testing.T) {
	go func() {
		time.Sleep(LISTENER_LAUNCH_DELAY)

		l, err := net.Listen(LISTENER_SERVER_NET, LISTENER_SERVER_ADDR)
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
		LISTENER_SERVER_NET,
		LISTENER_SERVER_ADDR,
		WAIT_TIMEOUT,
	) {
		t.Fatal("Could not wait to peer to be ready")
	}

	if WaitPeerListening(
		LISTENER_SERVER_NET,
		LISTENER_SERVER_ADDR,
		WAIT_TIMEOUT,
	) {
		t.Fatal("The peer should not be ready")
	}
}
