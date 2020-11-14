// file_test.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_CreateFolderSafely(t *testing.T) {

	const TmpFolder = "/tmp"
	const TmpTestFolder = "libs-temporary"

	var aTest *tester.Test
	var err error
	var result os.FileInfo

	aTest = tester.New(t)

	// Prepare the Environment.
	testFolder := filepath.Join(TmpFolder, TmpTestFolder)
	err = os.Mkdir(testFolder, 0755)
	aTest.MustBeNoError(err)
	folderX := filepath.Join(testFolder, "x")

	// Test #1. Create a Folder for the first Time.
	err = CreateFolderSafely(
		folderX,
		0755,
	)
	aTest.MustBeNoError(err)
	result, err = os.Stat(folderX)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result.IsDir(), true)

	// Test #2. Create a Folder for the second Time.
	err = CreateFolderSafely(
		folderX,
		0755,
	)
	aTest.MustBeNoError(err)
	result, err = os.Stat(folderX)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result.IsDir(), true)

	// Clear the Environment.
	err = os.RemoveAll(testFolder)
	aTest.MustBeNoError(err)
}

func Test_Exists(t *testing.T) {

	const (
		TestFile             = "test-file.txt"
		FileThatDoesNotExist = "test-file-that-does-not-exist.txt"
	)

	var aTest *tester.Test
	var err error
	aTest = tester.New(t)

	// Prepare the Environment.
	{
		var file *os.File
		file, err = os.Create(TestFile)
		aTest.MustBeNoError(err)
		err = file.Close()
		aTest.MustBeNoError(err)
	}

	// Test #1. Check the existing File.
	var exists bool
	exists, err = Exists(TestFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(exists, true)

	// Test #2. Check the File which does not exist.
	exists, err = Exists(FileThatDoesNotExist)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(exists, false)

	// Clear the Environment.
	err = os.Remove(TestFile)
	aTest.MustBeNoError(err)
}
