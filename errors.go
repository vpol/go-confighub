package confighub

import "errors"

var WrongValueErr = errors.New("incorrect value")
var WrongTypeErr = errors.New("unknown type")
var NoValueErr = errors.New("incorrect value")
var TypeValueErr = errors.New("incorrect type")
var FileRefNotFoundErr = errors.New("referenced file not found")
