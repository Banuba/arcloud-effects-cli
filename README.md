# Banuba AR Cloud CLI for effects

Banuba AR Cloud service enables your app to connect with our servers and access AR effects directly from the cloud. This means that the effects are downloaded on-demand, whenever a user requires them, instead of being pre-packaged within your app. As a result, your app remains lightweight while still offering a wealth of content.

AR Cloud service can be deployed on your favorite CDN solution i.e. AWS.

## Process data
Use `.zip` archive with AR effects provided by Banuba representatives to process data before uploading on your CDN.

CLI tools are in ```bin``` folder.  
Run the program with the required parameters in Terminal.

### Mac
```bash
./mac_arcloud-effects-cli -s <path/to/effects.zip> -u <your_arcloud_domain>
```

### Windows
```
.\win_arcloud-effects-cli.exe -s <path\to\effects.zip> -u <your_arcloud_domain>
```

### Parameters

- `-s, --source`: Path to the effects' zip file (required)
- `-d, --destination`: Destination folder to extract the effects (default: effects)
- `-u, --api-url`: Arcloud URL (required)

### Example
```/mac_arcloud-effects-cli -s /Users/username/Work/arcloud-effects-cli/banuba_effects.zip -u 'https://resources.custom/banuba-resources/'```

As a result ```effects``` folder will be created with processed data.  
```effects``` folder includes
1. ```<effect_name>.zip``` and ```<effect_name>.png``` files for each AR effect of the provided `.zip`.
2. ```api_response``` meta file that has json structure for all effects.

Example of ```api_response```
```json
{
  "effects": [
    {
      "URL": "https://resources.custom/banuba-resources/Effect1.zip",
      "Preview": "https://resources.custom/banuba-resources/Effect1.png",
      "ETag": "\"106r4b391ad99eefd9a17b608c64ddae\""
    },
    {
      "URL": "https://resources.custom/banuba-resources/Effect2.zip",
      "Preview": "https://resources.custom/banuba-resources/Effect2.png",
      "ETag": "\"996r4b391ad99eefd9a17b608c64ddae\""
    },
    ...
  ]
}
```
Fields:
- `"URL"` - path to zip archive with effect's files. All effect's files should be placed in root of archive.
- `"Preview"` - path to preview png file. Each effect has `preview.png` file.
- `"ETag"` - MD5 hash of effect `.zip` file.

## Upload to CDN
Use data from ```effects``` folder.

### Manifest file
Upload ```api_response``` file to ```https://<your_arcloud_domain>/v1/effects/api_response```

Manifest file just a JSON file without file extension. Set `Content-Type` as `application/json` before upload file to s3
bucket. The file's name (UUID) should be the same as bucket contained effects.

### Effects
Upload all ```.zip``` and ```.png``` files to ```<your_arcloud_domain>```

- `/` (folder store effects and their preview)
    - `Effect1.zip`
    - `Effect1.png>`
- `v1/`
    - `effects/`
        - `api_response`


