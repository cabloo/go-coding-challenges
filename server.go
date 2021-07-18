/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"

	"github.com/microsoft/vscode-remote-try-go/hello"
)

func main() {
	fmt.Println(hello.Hello())
	// portNumber := "9000"
	// http.HandleFunc("/", handle)
	// fmt.Println("Server listening on port ", portNumber)
	// http.ListenAndServe(":"+portNumber, nil)
}
