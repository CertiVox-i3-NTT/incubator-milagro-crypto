/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

/* Password Based Key Derivation Function Example */

package main

import (
	"encoding/hex"
	"fmt"

	amcl "git.apache.org/incubator-milagro-crypto.git/go/amcl-go"
)

func main() {
	// Seed value for Random Number Generator (RNG)
	seedHex := "9e8b4178790cd57a5761c4a6f164ba72"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := amcl.NewRAND()
	rng.Seed(len(seed), seed)

	// Generate random byte values
	for i := 0; i < 10; i++ {
		val := amcl.GENERATE_RANDOM(rng, 12)
		fmt.Printf("Random byte array %s\n", hex.EncodeToString(val))
	}

}
