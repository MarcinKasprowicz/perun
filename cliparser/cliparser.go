// Copyright 2017 Appliscale
//
// Maintainers and Contributors:
//
//   - Piotr Figwer (piotr.figwer@appliscale.io)
//   - Wojciech Gawroński (wojciech.gawronski@appliscale.io)
//   - Kacper Patro (kacper.patro@appliscale.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cliparser provides tools and structures for parsing and validating
// perun CLI arguments.
package cliparser

import (
	"errors"
	"strings"
	"gopkg.in/alecthomas/kingpin.v2"
)

const ValidateMode string = "validate"
const ConvertMode string = "convert"
const OfflineValidateMode = "validate_offline"

const JSON string = "json"
const YAML string = "yaml"

type CliArguments struct {
	Mode              *string
	TemplatePath      *string
	OutputFilePath    *string
	OutputFileFormat  *string
	ConfigurationPath *string
	Quiet             *bool
	Yes               *bool
	Verbosity         *string
	MFA               *bool
	Profile           *string
}

// Get and validate CLI arguments. Returns error if validation fails.
func ParseCliArguments() (cliArguments CliArguments, err error) {

	cliArguments.Mode = kingpin.Flag("mode", ValidateMode+"|"+OfflineValidateMode+"|"+ConvertMode).Short('m').String()
	cliArguments.TemplatePath = kingpin.Flag("template", "A path to the template").Short('t').String()
	cliArguments.OutputFilePath = kingpin.Flag("output", "A path, where converted file will be saved").Short('o').String()
	cliArguments.OutputFileFormat = kingpin.Flag("format", "Output format: "+strings.ToUpper(JSON)+"|"+strings.ToUpper(YAML)).Short('x').String()
	cliArguments.ConfigurationPath = kingpin.Flag("config", "A path to the configuration file").Short('c').String()
	cliArguments.Quiet = kingpin.Flag("quiet", "No console output, just return code").Short('q').Bool()
	cliArguments.Yes = kingpin.Flag("yes", "Always say yes").Short('y').Bool()
	cliArguments.Verbosity = kingpin.Flag("verbosity", "TRACE|DEBUG|INFO|ERROR").Short('v').String()
	cliArguments.MFA = kingpin.Flag("mfa", "Enable AWS MFA").Bool()
	cliArguments.Profile = kingpin.Flag("profile", "An AWS profile.").Short('p').String()

	kingpin.Parse()

	if *cliArguments.Mode == "" {
		err = errors.New("You should specify what you want to do with --mode flag")
		return
	}

	if *cliArguments.Mode != ValidateMode && *cliArguments.Mode != ConvertMode && *cliArguments.Mode != OfflineValidateMode {
		err = errors.New("Invalid mode. Use validate, validate_offline or convert")
		return
	}

	if *cliArguments.TemplatePath == "" {
		err = errors.New("You should specify a source of the template file with --template flag")
		return
	}

	if *cliArguments.Mode == ConvertMode {
		if *cliArguments.OutputFilePath == "" {
			err = errors.New("You should specify a output file path with --output flag")
			return
		}

		if *cliArguments.OutputFileFormat == "" {
			err = errors.New("You should specify a output file format with --format flag")
			return
		}

		*cliArguments.OutputFileFormat = strings.ToLower(*cliArguments.OutputFileFormat)
		if *cliArguments.OutputFileFormat != JSON && *cliArguments.OutputFileFormat != YAML {
			err = errors.New("Invalid output file format. Use JSON or YAML")
			return
		}
	}

	return
}
