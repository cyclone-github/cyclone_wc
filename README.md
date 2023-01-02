# Cyclone Word Count
Alternative to GNU wc -l that is cross compiled for Linux, Windows & Mac. Cyclone Word Count by default will count each line of input file. See usage examples below.

Usage Examples:
- cyclone_wc -help _(displays help info)_
- cyclone_wc -version _(displays version info)_
- cyclone_wc -i input.txt

## Testing

### Linux:
- $ ./cyclone_wc.bin -i cyclone_hk_v2.txt
- Counting lines:  cyclone_hk_v2.txt
- Finished in:     7.599916052s (80.115 million lines/sec)
- Total lines:     608864500

### Windows:
- \> .\cyclone_wc.exe -i cyclone_hk_v2.txt
- Counting lines:  cyclone_hk_v2.txt
- Finished in:     8.2322596s (73.961 million lines/sec)
- Total lines:     608864500

### Mac:
- _untested_
