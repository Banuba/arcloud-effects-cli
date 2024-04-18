# arcloud-effects-cli

---

## ARCloud structure

- `<UUID>/` (folder store effects and their preview)
    - `<effect's_files_zip_archive>.zip`
    - `<effect_preview_file.png>`
- `v1/`
    - `effects/`
        - `<json_manifest_file>`

## Prepare effects for arcloud uploading

Execute the program with the required parameters:

#### Mac
```bash
./mac_arcloud-effects-cli -s <path/to/effects.zip> --id <UUID-of-your-bucket-with-effects> --api-url <your_arcloud_domain>
```

#### Windows
```
.\win_arcloud-effects-cli.exe -s <path\to\effects.zip> --id <UUID-of-your-bucket-with-effects> --api-url <your_arcloud_domain>
```

### Parameters:

- `-s, --source`: Path to the effects' zip file (required)
- `-d, --destination`: Destination folder to extract the effects (default: effects)
- `--id`: Arcloud folder ID (required)
- `-u, --api-url`: Arcloud URL (required)

## Json Manifest File

Path structure to file:
```
https://<your_arcloud_domain>/v1/effects/<json_manifest_file>
```

Manifest file just a JSON file without file extension. Set `Content-Type` as `application/json` before upload file to s3
bucket. The file's name (UUID) should be the same as bucket contained effects.

Example:
```
https://api.arcloud.banuba.net/v1/effects/B4E0A9AA-16C7-47DD-9D00-24B5536B2932
```

## Json Manifest file structure

```json
{
  "effects": [
    {
      "URL": "https://api.arcloud.banuba.net/B4E0A9AA-16C7-47DD-9D00-24B553/2_5D_HeadphoneMusic.zip",
      "Preview": "https://api.arcloud.banuba.net/B4E0A9AA-16C7-47DD-9D00-24B553/2_5D_HeadphoneMusic.png",
      "ETag": "\"03623b391ad77eefd9a17b608c64ddae\""
    },
    ...
  ]
}
```

Fields:

- `"URL"` - path to zip archive with effect's files. All effect's files should be placed in root of archive.
- `"Preview"` - path to preview png file. Each effect has `preview.png` file.
- `"ETag"` - MD5 hash of effect `.zip` file.

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
