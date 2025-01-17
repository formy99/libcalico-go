// Copyright (c) 2021 Tigera, Inc. All rights reserved.

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

package resources

import (
	log "github.com/sirupsen/logrus"

	libapiv3 "github.com/projectcalico/libcalico-go/lib/apis/v3"
	cnet "github.com/projectcalico/libcalico-go/lib/net"
)

// FindNodeAddress returns node address of the specified type. Type can be one of
// CalicoNodeIP, InternalIP or ExternalIP
func FindNodeAddress(node *libapiv3.Node, ipType string) (*cnet.IP, *cnet.IPNet) {
	for _, addr := range node.Spec.Addresses {
		if addr.Type == ipType {
			ip, cidr, err := cnet.ParseCIDROrIP(addr.Address)
			if err == nil {
				if ip.To16() == nil {
					continue
				}
				log.WithFields(log.Fields{"ip": ip, "cidr": cidr}).Debug("Parsed IPv6 address")
				return ip, cidr
			} else {
				log.WithError(err).WithField("IPv6Address", addr.Address).Warn("Failed to parse IPv6Address")
			}
		}
	}
	return nil, nil
}
