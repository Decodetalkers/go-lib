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

import (
	"reflect"
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
