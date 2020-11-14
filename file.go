// file.go.

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
)

// Tries to create a Folder ignoring an Error if the Folder already exists.
func CreateFolderSafely(
	path string,
	permissions os.FileMode,
) (err error) {

	err = os.Mkdir(path, permissions)
	if err != nil {
		if os.IsExist(err) {
			err = nil
			return
		}
		return
	}

	return
}

// Tries to get Access to a File.
func Exists(
	filePath string,
) (exists bool, err error) {
	_, err = os.Stat(filePath)
	if err == nil {
		exists = true
		return
	}
	if os.IsNotExist(err) {
		err = nil
		exists = false
		return
	}
	// We can not say whether it exists or not while it is not accessible.
	// So, we return an Error.
	return
}
