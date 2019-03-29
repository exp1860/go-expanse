// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package mru

import (
	"crypto/ecdsa"

	"github.com/expanse-org/go-expanse/common"
	"github.com/expanse-org/go-expanse/crypto"
)

// Signs resource updates
type Signer interface {
	Sign(common.Hash) (Signature, error)
}

type GenericSigner struct {
	PrivKey *ecdsa.PrivateKey
}

func (self *GenericSigner) Sign(data common.Hash) (signature Signature, err error) {
	signaturebytes, err := crypto.Sign(data.Bytes(), self.PrivKey)
	if err != nil {
		return
	}
	copy(signature[:], signaturebytes)
	return
}
