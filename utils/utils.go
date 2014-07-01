/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package utils

// #include <stdlib.h>
import "C"
import "unsafe"

import (
	"io/ioutil"
	"os"
	"os/user"
	"reflect"
	"strings"
)

func IsElementEqual(e1, e2 interface{}) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	return reflect.DeepEqual(e1, e2)
}

func IsElementInList(e interface{}, list interface{}) bool {
	if list == nil {
		return false
	}

	v := reflect.ValueOf(list)
	if !v.IsValid() {
		return false
	}

	if v.Type().Kind() == reflect.Slice ||
		v.Type().Kind() == reflect.Array {
		l := v.Len()
		for i := 0; i < l; i++ {
			if IsElementEqual(e, v.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

func CopyFile(src, dest string) bool {
	if ok := IsFileExist(src); !ok && len(dest) < 1 {
		return false
	}

	contents, err := ioutil.ReadFile(src)
	if err != nil {
		return false
	}

	f, err1 := os.Create(dest + "~")
	if err1 != nil {
		return false
	}
	defer f.Close()

	if _, err := f.Write(contents); err != nil {
		return false
	}
	f.Sync()

	if err := os.Rename(dest+"~", dest); err != nil {
		return false
	}

	return true
}

func IsFileExist(filename string) bool {
	if len(filename) < 1 {
		return false
	}

	path := URIToPath(filename)
	if len(path) < 1 {
		return false
	}
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

func IsEnvExists(envName string) (ok bool) {
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, envName+"=") {
			ok = true
			break
		}
	}
	return
}

func UnsetEnv(envName string) (err error) {
	doUnsetEnvC(envName) // call C.unsetenv() is necessary
	envs := os.Environ()
	newEnvsData := make(map[string]string)
	for _, e := range envs {
		a := strings.SplitN(e, "=", 2)
		var name, value string
		if len(a) == 2 {
			name = a[0]
			value = a[1]
		} else {
			name = a[0]
			value = ""
		}
		if name != envName {
			newEnvsData[name] = value
		}
	}
	os.Clearenv()
	for e, v := range newEnvsData {
		err = os.Setenv(e, v)
		if err != nil {
			return
		}
	}
	return
}

func doUnsetEnvC(envName string) {
	cname := C.CString(envName)
	defer C.free(unsafe.Pointer(cname))
	C.unsetenv(cname)
}

func GetUserName() string {
	info, err := user.Current()
	if err != nil {
		return ""
	}

	return info.Username
}
