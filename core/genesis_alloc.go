// Copyright 2017 The go-ethereum Authors
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

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xf9\x02\x18\xf8o\x946\x1eIX\xbd\x8bM\x03\\h\xc4s\x19>w\xfd\xabj\xea\x9d\xf8X\x8965\u026d\xc5\u07a0\x00\x00\x89\x1b\x1a\xe4\xd6\xe2\xefP\x00\x00\x80\xb8A\x04\f/\a\xcf\"\x96n\xba\xa1\x9d\xd8}\xe1\x1d\x04\xf9\xa6\x06\xe8>u\xa5)Utv\nZM\xd8(\a\xe7\xd9\xf2\x8f\x86\xcd8\x98q\u0753\x1aG|\xacd\xe23\xa6\xf0\x98\u0572W\x87\x1d\x81c\xdc\x1b\xdd\xcf\xee\x94Wei-Ning Sonic Bojie\u0600\x80\x94Wei-Ning Sonic Bojie\x80\xf8o\x94\xb6\xa2\xe2p\xb5O\x19\xb4\x00-e5\xa7\x1c\xfd\xd1\xca\x1f_\xc1\xf8X\x8965\u026d\xc5\u07a0\x00\x00\x89\x1b\x1a\xe4\xd6\xe2\xefP\x00\x00\x80\xb8A\x04c\xf7\x8d9\x8fk\x04)R\x06*&\xf3U\x1aN)J\x11\bK\x10\x06\xc8\xd67d7\xb7$\x88\xd8\xf3\xafk\x84\x991\xc06$\xad-X\x19\x00\xe9\xad_<\xfe\xee\xd0@\u01d9\xe50\x8c\xd8\x01(S\xfb\xf8o\x94\xb6\xbb\xbe9\x0311\xe4(\u0563TXV\xa5\x96\xca\xcbx\x10\xf8X\x8965\u026d\xc5\u07a0\x00\x00\x89\x1b\x1a\xe4\xd6\xe2\xefP\x00\x00\x80\xb8A\x04\xa1\x11\x8c\x8f\x1c.)=\xe5(\xe5\xb4Z\xc5\xe6Qg\xafL\xb4\xd1\b\x87\x92J\x9f\x83\x92\x87\xd0`\xda&AN\x97\x1c\xabu\x01>V\r\xfe\u01c9{\x19\x83\xea\xa61\xae/\x16\xea@bRS\x10\xed\xd8\xea\xf8o\x94\xbcq\xc0\xf2\xae\xcaF\xfa\x00\x01\xe7e\x96P)\xb1j\x06\x04\x9e\xf8X\x8965\u026d\xc5\u07a0\x00\x00\x89\x1b\x1a\xe4\xd6\xe2\xefP\x00\x00\x80\xb8A\x04\xa1*\x01_'\u064eS\x05I\x8d#\xae\xac\xde(\x1eH\xaa\xd9\x18\xae\x87\xd7\xf6\x17\x84\xabC\x82\x89S\u06d9\xbf\x1fDBb\x94\xcb\xc2\x1e2\xcf:\x99\x12/,\xe8\xa3[\xa7\x18\xa7\xc8\x14\xee\u01fa\xe9\xa3\xea\u453f\x8cH\xa6 \xba\xccF\x90\u007f\x9b\x89s-%\xe4z-|\xf7\u038a\x15-\x02\xc7\xe1J\xf6\x80\x00\x00\x80\x80\x80"

const testnetAllocData = mainnetAllocData
