package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Lexer struct {
	CommandString string
}

/* take a string and searched a command defined, execute a function for the command */
func (tmp *Lexer) GeneralComand(command string) {
	tmp.CommandString = command
	if matched, _ := regexp.Match("(mkdisk)(.*)", []byte(tmp.CommandString)); matched {
		tmp.CommandMkdir()
	} else if matched, _ := regexp.Match("(rmdisk)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando rmdisk")
	} else if matched, _ := regexp.Match("(fdisk)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando fdisk")
	} else if matched, _ := regexp.Match("(mount)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando mount")
	} else if matched, _ := regexp.Match("(unmount)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando unmount")
	} else if matched, _ := regexp.Match("(mkfs)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando mkfs")
	} else if matched, _ := regexp.Match("(rep)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando rep")
	} else if matched, _ := regexp.Match("(pause)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando pause")
	} else if matched, _ := regexp.Match("(login)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando login")
	} else if matched, _ := regexp.Match("(logout)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando Logout")
	} else if matched, _ := regexp.Match("(mkgrp)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando mkgrp")
	} else if matched, _ := regexp.Match("(rmgrp)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando rmgrp")
	} else if matched, _ := regexp.Match("(mkuser)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando mkuser")
	} else if matched, _ := regexp.Match("(rmuser)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando rmuser")
	} else if matched, _ := regexp.Match("(mkfile)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando mkfile")
	} else if matched, _ := regexp.Match("(mkdir)(.*)", []byte(tmp.CommandString)); matched {
		fmt.Println("contiene el comando mkdir")
	}
}

/* This method is used for make file of types binaries, with the structure implemented*/
func (tmp *Lexer) CommandMkdir() {
	pathMkdir := tmp.PathParameter(true)
	size := tmp.SizeParameter(true)
	fit := tmp.fitParameter(false)
	unit := tmp.unitParameter(false)
	if pathMkdir != "" && size > 0 {
		tmp := Mkdisk{Path: pathMkdir, Fit: fit, Unit: unit, Size: size}
		tmp.Execute()
	}

}

/* The parameter path contain the path where create file bin*/
func (tmp *Lexer) PathParameter(obligatory bool) string {
	result1 := ""
	if matchedPaht, _ := regexp.Match(">path=", []byte(tmp.CommandString)); matchedPaht {
		//verificamos que la path con comillas
		if matchedPaht1, _ := regexp.Match("\"(/.*)+/[a-zA-Z0-9]+.dsk\"", []byte(tmp.CommandString)); matchedPaht1 {
			regePath := regexp.MustCompile("\"(/.*)+/[a-zA-Z0-9]+.dsk\"")
			content := regePath.FindAllString(tmp.CommandString, -1)
			if len(content) > 0 {
				result1 = content[0]
			}
		} else if matchedPaht1, _ := regexp.Match("(/[a-zA-Z0-9\\.]+)+", []byte(tmp.CommandString)); matchedPaht1 {
			regePath := regexp.MustCompile("(/[a-zA-Z0-9...]+)+")
			content := regePath.FindAllString(tmp.CommandString, -1)
			if len(content) > 0 {
				result1 = content[0]
			}
		}
	} else if obligatory {
		fmt.Println("El parametro Path es obligatorio")
		result1 = ""
	} else if !obligatory {
		result1 = ""
	}
	return result1
}

/* The parameter size contain the size of file or partition or a diferent object */
func (tmp *Lexer) SizeParameter(obligatory bool) int {
	var size int
	var tmpString string
	if matched, _ := regexp.Match(">size=", []byte(tmp.CommandString)); matched {
		regeSize := regexp.MustCompile(">size=[0-9]+")
		content := regeSize.FindAllString(tmp.CommandString, -1)
		if len(content) > 0 {
			tmpString = content[0]
			tmpString = strings.Trim(tmpString, ">size=")
			size, _ = strconv.Atoi(tmpString)
		}
	} else if obligatory {
		fmt.Println("El parametro size es obligatorio")
		size = -1
	} else {
		size = 0
	}
	return size
}

/* This parameter fit contain the configuration of asignation of disk or partition*/
func (tmp *Lexer) fitParameter(obligatory bool) byte {
	var fit byte
	var tmpString string
	if matched, _ := regexp.Match(">fit=", []byte(tmp.CommandString)); matched {
		regexFit := regexp.MustCompile(">fit=(bf|ff|wf)")
		content := regexFit.FindAllString(tmp.CommandString, -1)
		tmpString = content[0]
		remplace := regexp.MustCompile(">fit=")
		tmpString = remplace.ReplaceAllString(tmpString, "")
		switch {
		case tmpString == "ff":
			fit = 'f'
		case tmpString == "bf":
			fit = 'b'
		case tmpString == "wf":
			fit = 'w'
		}
	} else if obligatory {
		fmt.Println("El parametro fit es obligatorio")
		fit = 'n'
	} else if !obligatory { // is optional, i can asigned orther value
		fit = 'o'
	}
	return fit
}

/* The parameter unit contain de information respect the type of storage in the disk or partition*/
func (tmp *Lexer) unitParameter(obligatory bool) byte {
	var unit byte
	var tmpString string
	if matched, _ := regexp.Match(">unit=", []byte(tmp.CommandString)); matched {
		regexFit := regexp.MustCompile(">unit=(k|m)")
		content := regexFit.FindAllString(tmp.CommandString, -1)
		tmpString = content[0]
		remplace := regexp.MustCompile(">unit=")
		tmpString = remplace.ReplaceAllString(tmpString, "")
		switch {
		case tmpString == "m":
			unit = 'm'
		case tmpString == "k":
			unit = 'k'
		}
	} else if obligatory {
		fmt.Println("El parametro unit es obligatorio")
		unit = 'n'
	} else if !obligatory { // is optional, i can asigned orther value
		unit = 'o'
	}
	return unit
}