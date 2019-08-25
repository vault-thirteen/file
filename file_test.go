////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
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

// +build test

package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_CreateFolderSafely(t *testing.T) {

	const TmpFolder = "/tmp"
	const TmpTestFolder = "ads-libraries-temporary"

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
