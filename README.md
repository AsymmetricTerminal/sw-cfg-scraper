# Sonicwall Config Scraper

## Description

Decided to create a script to faciliate dumping Sonicwall .exp exports to plaintext. This tool does NOT work on Sonicwall config exports that have been password protected.

## Usage

Run the binary in the directory containing your Sonicwall .exp file(s). It will dump all .exp files found in the current directory to a plaintext file for each one named `sonicwall_config_export(FILENAME).txt`.
