# arcloud-effects-cli

---

## Running the Program

Execute the program with the required parameters:

```bash
./arcloud-effects-cli -s <path/to/effects.zip> --id <Arcloud folder ID> --api-url <Arcloud URL>
```

### Parameters:

- `-s, --source`: Path to the effects' zip file (required)
- `-d, --destination`: Destination folder to extract the effects (default: effects)
- `--id`: Arcloud folder ID (required)
- `-u, --api-url`: Arcloud URL (required)

### Example

```bash
./effects-zip-extractor -s ./effects.zip --id 12345-abcde-1a2b3c4d --api-url https://example.com/api
```

## Manifest File

The program generates a JSON manifest file containing metadata about the extracted effects. Each entry in the manifest includes the following information:

- URL: The URL to download the effect zip file. 
- Preview: The URL to download the preview image of the effect.
- ETag: The ETag value calculated from the effect's data.
                                                            
## Effects zip archive structure

```
effects.zip
    effect_1/
        preview.png
        ...and other files
    effect_2/
        preview.png
        ...and other files 
```
