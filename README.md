# Sonicwall Config Scraper

## Description

Decided to create a script to faciliate dumping Sonicwall .cfg exports to plaintext. This tool does NOT work on Sonicwall config exports that have been password protected.

## Usage

Run the binary in the directory containing your Sonicwall .cfg file(s). It will dump all .cfg files found in the current directory to a plaintext file for each one named `sonicwall_config_export(FILENAME).txt`.
