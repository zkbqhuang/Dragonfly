/*
 * Copyright The Dragonfly Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package uploader

import (
	"io"
	"sync"
)

var (
	mutex     sync.RWMutex
	uploadMap = map[string]Uploader{}
)

func RegisterUploader(pattern string, up Uploader) {
	mutex.Lock()
	defer mutex.Unlock()

	uploadMap[pattern] = up
}

func GetUploader(pattern string) (Uploader, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	up, ok := uploadMap[pattern]
	return up, ok
}

// Uploader defines how to upload range by path.
type Uploader interface {
	// UploadRange defines how to upload range by path.
	UploadRange(path string, off, size int64, opt interface{}) (io.ReadCloser, error)
}