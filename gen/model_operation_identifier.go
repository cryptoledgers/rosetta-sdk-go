// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package gen

// OperationIdentifier The `operation_identifier` uniquely identifies an operation within a transaction.
type OperationIdentifier struct {
	// The operation `index` is used to ensure each operation has a unique identifier within a transaction.  To clarify, there may not be any notion of an operation index in the blockchain being described.
	Index int64 `json:"index"`
	// Some blockchains specify an operation index that is essential for client use. For example, Bitcoin uses a `network_index` to identify which UTXO was used in a transaction.  `network_index` should not be populated if there is no notion of an operation index in a blockchain (typically most account-based blockchains).
	NetworkIndex *int64 `json:"network_index,omitempty"`
}